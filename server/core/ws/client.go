package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
	sendChanSize   = 256
)

// Client WebSocket 客户端连接
type Client struct {
	UUID    string // 客户端唯一标识（用于单推）
	Uid     string // 用户 ID（用于按用户推送）
	RoomID  string // 房间 ID（用于群推）
	conn    *websocket.Conn
	send    chan []byte
	Manager *Manager

	closeOnce sync.Once // 保证 closeConn 只执行一次
	closed    chan struct{}
}

// NewClient 创建客户端实例
func NewClient(uuid, uid, roomID string, conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		UUID:    uuid,
		Uid:     uid,
		RoomID:  roomID,
		conn:    conn,
		send:    make(chan []byte, sendChanSize),
		Manager: manager,
		closed:  make(chan struct{}),
	}
}

// closeConn 保证只执行一次：注销 + 关闭连接 + 关闭 send channel
func (c *Client) closeConn() {
	c.closeOnce.Do(func() {
		close(c.closed)
		c.conn.Close()
		// 通知 Manager 注销（非阻塞，防止死锁）
		select {
		case c.Manager.UnRegister <- c:
		default:
		}
	})
}

// Read 读取客户端消息（goroutine 运行）
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
		msg := string(data)
		if msg == "ping" {
			c.conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			continue
		}
		log.Printf("[ws] message from %s: %s", c.UUID, msg)
	}
}

// Write 向客户端写入消息（goroutine 运行）
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.closeConn()
	}()

	for {
		select {
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

// Close 主动关闭连接（外部调用，幂等安全）
func (c *Client) Close() {
	c.closeConn()
}

// Send 非阻塞发送消息到 send channel，返回是否成功
func (c *Client) Send(data []byte) bool {
	select {
	case <-c.closed:
		return false
	case c.send <- data:
		return true
	default:
		// channel 满，丢弃消息并记录
		log.Printf("[ws] client %s send buffer full, message dropped", c.UUID)
		return false
	}
}
