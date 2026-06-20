package core

import "x_admin/core/ws"

// Ws WebSocket 管理器,作为基础设施方便全局使用
var Ws = ws.NewManager()

func init() {
	go Ws.Start()
}
