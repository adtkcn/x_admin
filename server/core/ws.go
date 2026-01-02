package core

import "x_admin/util/ws_util"

// Ws WebSocket 管理器,作为基础设施方便全局使用
var Ws = ws_util.NewManager()

func init() {
	go Ws.Start()
}
