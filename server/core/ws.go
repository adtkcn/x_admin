package core

import "x_admin/util/ws_util"

var Ws = ws_util.NewManager()

func init() {
	go Ws.Start()
}
