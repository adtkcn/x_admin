package controller

import (
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

// @Summary	websocket连接
// @Tags		公共接口
// @Router		/api/ws [get]
// @Param		token	header	string	true	"token"
// @Param		uid		query	string	true	"用户ID"
// @Param		room	query	string	true	"房间ID"
// @Schemes	ws
func WsHandler(c *gin.Context) {
	uuid := util.ToolsUtil.MakeUuidV7()
	// 从查询参数获取用户ID和房间ID（实际项目中应通过认证获取）
	var adminId = config.AdminConfig.GetAdminId(c)
	roomID := c.Query("room")
	if adminId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "adminId is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	client := ws.NewClient(uuid, adminId, roomID, conn, core.Ws)
	core.Ws.Register <- client

	// 启动读写协程
	go client.Write()
	go client.Read()
}
