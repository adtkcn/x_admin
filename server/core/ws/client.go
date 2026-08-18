package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket 连接相关常量配置
const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 2048
	sendChanSize   = 256
)

// MessageHandler 业务消息处理回调，由 Controller 层设置。
// 收到非心跳消息时调用，data 为原始消息字节。
type MessageHandler func(client *Client, data []byte)

// Client 表示一个 WebSocket 客户端连接。
//
// 职责边界：
//   - 基础设施层：连接生命周期管理、心跳、消息读写
//   - 不处理业务消息：非心跳消息通过 OnMessage 回调上抛给 Controller 层
//
// 并发模型：
//   - Read goroutine:  读取消息；收到文本 "ping" 时投递 pong 信号，
//     业务消息回调 OnMessage。不直接写连接。
//   - Write goroutine: conn 的唯一写入方，负责业务消息、pong 应答、
//     定期 Ping 保活。单一写入方设计使得无需写锁。
type Client struct {
	UUID string // 客户端唯一标识（UUID v7）
	Uid  string // 用户 ID

	conn      *websocket.Conn
	send      chan []byte
	Manager   *Manager
	closeOnce sync.Once
	closed    chan struct{}
	OnMessage MessageHandler // 业务消息回调，由 Controller 层设置

	// pong 文本心跳应答的专用通道，容量 1。
	//
	// 独立于 send 的原因：心跳对时延敏感（前端 pongTimeout 仅 1s），
	// 若与业务消息共用 send，高峰期会被排队甚至丢弃，导致误判掉线。
	//
	// 容量 1 且投递时非阻塞：同一时刻只需保留一个待发 pong，
	// 客户端心跳间隔 10s，堆积多个没有意义。
	pong chan struct{}
}

// NewClient 创建 WebSocket 客户端实例。
func NewClient(uuid, uid string, conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		UUID:    uuid,
		Uid:     uid,
		conn:    conn,
		send:    make(chan []byte, sendChanSize),
		pong:    make(chan struct{}, 1),
		Manager: manager,
		closed:  make(chan struct{}),
	}
}

// closeConn 关闭连接并清理资源，保证只执行一次。
// 直接调用 unregisterClient，避免 channel 满时资源泄漏。
//
// 关于并发：conn.Close() 可能与 Write goroutine 的 WriteMessage 并发执行。
// gorilla/websocket 明确允许 Close 与读写并发调用（它是唯一可并发调用的方法），
// 此时进行中的写会立即返回错误，Write 循环随即退出，因此无需加锁。
func (c *Client) closeConn() {
	c.closeOnce.Do(func() {
		close(c.closed)
		c.conn.Close()
		c.Manager.unregisterClient(c)
	})
}

// Read 持续读取客户端消息。
//
// 仅处理心跳（ping→pong），其他所有消息通过 OnMessage 回调上抛给 Controller 层。
// 不在基础设施层解析任何业务消息。
func (c *Client) Read() {
	defer c.closeConn()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			// 区分「正常断开」与「异常错误」，避免日志被误报淹没。
			//
			// 以下三种关闭码属于客户端断开的常规场景，非服务端故障：
			//   CloseNormalClosure(1000)   前端主动 close()，如退出登录
			//   CloseGoingAway(1001)       页面刷新、关闭标签页、跳转
			//   CloseAbnormalClosure(1006) 未收到 Close 帧即断开，
			//                              常见于断网、休眠、进程被杀
			//
			// 其余情况（读超时、消息超长、协议错误等）才视为真正的异常。
			if websocket.IsCloseError(err,
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Printf("[ws] client %s closed: %v", c.UUID, err)
			} else {
				log.Printf("[ws] read error for %s: %v", c.UUID, err)
			}
			return
		}

		// 文本心跳：兼容无法发送 Ping 控制帧的浏览器客户端。
		//
		// 浏览器 WebSocket API 不暴露发送 Ping 控制帧的能力，因此前端
		// 采用文本 "ping" 作为心跳，服务端必须回复文本 "pong"。
		//
		// 此处有两点关键处理：
		//  1. 刷新读超时。SetPongHandler 只在收到 Pong 控制帧时触发，
		//     纯文本心跳不会走到那里，若不手动刷新，连接会在 pongWait
		//     后被误判为超时断开。
		//  2. 不在本 goroutine 写连接，而是投递信号给 Write goroutine，
		//     由其统一写出。保证 conn 始终只有一个写入方，无需加锁。
		//     使用独立的 pong 通道（而非 send），避免与业务消息争抢缓冲区。
		if string(data) == "ping" {
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			select {
			case c.pong <- struct{}{}:
			default:
				// 已有待发 pong 尚未写出，本次直接跳过即可
			}
			continue
		}

		// 业务消息：回调给 Controller 层处理
		if c.OnMessage != nil {
			c.OnMessage(c, data)
		}
	}
}

// Write 持续向客户端发送消息。
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.closeConn()
	}()

	// 事件循环：监听三种退出/写入信号，任一触发即处理并退出本循环
	for {
		select {
		// 连接已关闭：收到关闭信号，直接退出循环（defer 中已释放资源）
		case <-c.closed:
			return
		// 业务消息通道：将待发送消息写给客户端
		case message, ok := <-c.send:
			// 设置写超时，避免对端卡死导致写阻塞
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 通道被关闭：发送 Close 控制帧后退出
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// 正常发送文本消息；写入失败说明连接异常，退出循环
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		// 文本心跳应答：Read goroutine 收到 "ping" 后投递信号，此处统一写出
		case <-c.pong:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, []byte("pong")); err != nil {
				return
			}

		// 心跳定时器：定期发送 Ping 保持连接活跃，同时检测对端存活
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// 发送 Ping 失败，说明连接已断开，退出循环
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Close 主动关闭连接。幂等安全。
func (c *Client) Close() {
	c.closeConn()
}

// Send 非阻塞发送消息到 send 通道。
func (c *Client) Send(data []byte) bool {
	select {
	case <-c.closed:
		return false
	case c.send <- data:
		return true
	default:
		log.Printf("[ws] client %s send buffer full, message dropped", c.UUID)
		return false
	}
}
