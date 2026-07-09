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
//   - Read goroutine: 读取消息，心跳内部处理，业务消息回调 OnMessage
//   - Write goroutine: 发送消息，定期 Ping 保活
type Client struct {
	UUID string // 客户端唯一标识（UUID v7）
	Uid  string // 用户 ID

	conn      *websocket.Conn
	send      chan []byte
	Manager   *Manager
	closeOnce sync.Once
	closed    chan struct{}
	OnMessage MessageHandler // 业务消息回调，由 Controller 层设置
}

// NewClient 创建 WebSocket 客户端实例。
func NewClient(uuid, uid string, conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		UUID:    uuid,
		Uid:     uid,
		conn:    conn,
		send:    make(chan []byte, sendChanSize),
		Manager: manager,
		closed:  make(chan struct{}),
	}
}

// closeConn 关闭连接并清理资源，保证只执行一次。
// 直接调用 unregisterClient，避免 channel 满时资源泄漏。
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
			if websocket.IsCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Printf("[ws] client %s closed: %v", c.UUID, err)
			} else {
				log.Printf("[ws] read error for %s: %v", c.UUID, err)
			}
			return
		}

		// 心跳消息：基础设施层内部处理
		if string(data) == "ping" {
			select {
			case c.send <- []byte("pong"):
			default:
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

	for {
		select {
		case <-c.closed:
			return
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
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
