package ws_util

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	UUID    string // 用于单推
	Uid     string // 用于单推
	RoomID  string // 用于群推
	conn    *websocket.Conn
	send    chan []byte
	Manager *Manager //持有对 Manager 的引用:显式的依赖注入，结构清晰、可测试、可扩展
}

func NewClient(uuid, Uid, roomID string, conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		UUID:    uuid,
		Uid:     Uid,
		RoomID:  roomID,
		conn:    conn,
		send:    make(chan []byte, 256),
		Manager: manager,
	}
}

// 读取消息（通常用于接收客户端消息，此处简化）
func (c *Client) Read() {
	defer func() {
		// 从 Manager 中注销客户端
		c.Manager.UnRegister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)              // 设置最大消息大小
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) // 设置读取超时时间
	// 设置 Pong 处理函数，用于响应客户端的 Ping 消息
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait)) // 设置读取超时时间
		return nil
	})

	for {
		messageType, data, err := c.conn.ReadMessage()
		if err != nil {
			// 处理浏览器主动关闭的情况
			if websocket.IsCloseError(err,
				websocket.CloseGoingAway,         // 客户端正在关闭连接
				websocket.CloseAbnormalClosure) { // 异常关闭
				log.Printf("WebSocket client %s closed: %v", c.UUID, err)
			} else {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}
		msg := string(data)
		if msg == "ping" {
			// 回复客户端的 Ping 消息
			c.conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			continue
		}
		log.Printf("WebSocket message:%d, %s", messageType, msg)
		// 可在此处理客户端发来的消息（如聊天内容）
	}
}

// 写入消息（由 Manager 触发）
func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// 消息，不合并
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
			// w, err := c.conn.NextWriter(websocket.TextMessage)
			// if err != nil {
			// 	return
			// }
			// w.Write(message)

			// // 支持多条消息合并发送（可选）
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	w.Write([]byte("\n"))
			// 	w.Write(<-c.send)
			// }

			// if err := w.Close(); err != nil {
			// 	return
			// }

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 主动关闭连接
func (c *Client) Close() {
	c.Manager.UnRegister <- c
	c.conn.Close()
}
