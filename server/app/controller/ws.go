package controller

import (
	"encoding/json/v2"
	"log"
	"net/http"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/ws"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产环境应严格校验
	},
}

// clientMessage 客户端发送的业务消息结构
type clientMessage struct {
	Type string `json:"type"` // 消息类型: "join_room" | "leave_room" | ...
	Data string `json:"data"` // 数据部分
}

// handleWsMessage 处理客户端发来的业务消息。
// 由 ws.Client 的 OnMessage 回调触发，负责消息解析和业务路由。
func handleWsMessage(client *ws.Client, data []byte) {
	var msg clientMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Printf("[ws] invalid message from %s: %s", client.UUID, string(data))
		return
	}

	switch msg.Type {
	case "join_room":
		if msg.Data != "" {
			client.Manager.JoinRoom(client.UUID, msg.Data)
		}
	case "leave_room":
		if msg.Data != "" {
			client.Manager.LeaveRoom(client.UUID, msg.Data)
		}
	default:
		log.Printf("[ws] unknown message type %q from %s", msg.Type, client.UUID)
	}
}

// @Summary	websocket连接
// @Tags		公共接口
// @Router		/api/ws [get]
// @Param		token	header	string	true	"token"
// @Schemes	ws
func WsHandler(c *gin.Context) {
	uuid := util.ToolsUtil.MakeUuidV7()
	var adminId = config.AdminConfig.GetAdminId(c)
	if adminId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "adminId is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	client := ws.NewClient(uuid, adminId, conn, core.Ws)
	// 设置业务消息回调，在 Controller 层处理业务逻辑
	client.OnMessage = handleWsMessage
	core.Ws.Register <- client

	// 启动读写协程
	go client.Write()
	go client.Read()
}
