package monitorSchema

import "github.com/adtkcn/x_null"

// MonitorErrorListAddReq 错误对应的用户记录新增参数
type MonitorErrorListAddReq struct {
	ErrorId  string // 错误id
	ClientId string // 客户端id
	UserId   string // 用户id

	Width  x_null.Int64 // 屏幕
	Height x_null.Int64 // 屏幕高度

	Country  string // 国家
	Province string // 省份
	City     string // 城市
	Operator string // 电信运营商
	Ip       string // ip地址
	// Ua       string // ua信息
	// ProjectKey string // 项目id
}
