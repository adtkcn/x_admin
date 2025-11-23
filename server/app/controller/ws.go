package controller

import (
	"log"
	"net/http"
	"x_admin/core"
	"x_admin/util"
	"x_admin/util/ws_util"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产环境应严格校验
	},
}

func WsHandler(c *gin.Context) {
	uuid := util.ToolsUtil.MakeUuidV7()
	// 从查询参数获取用户ID和房间ID（实际项目中应通过认证获取）
	uid := c.Query("uid")
	roomID := c.Query("room")
	if uid == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	client := ws_util.NewClient(uuid, uid, roomID, conn, core.Ws)
	core.Ws.Register <- client

	// 启动读写协程
	go client.Write()
	go client.Read()
}
