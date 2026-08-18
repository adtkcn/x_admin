package monitor_schema

import "github.com/adtkcn/x_null"

// MonitorErrorListAddReq 错误对应的用户记录新增参数
type MonitorErrorListAddReq struct {
	ErrorId  string `json:"error_id"`  // 错误id
	ClientId string `json:"client_id"` // 客户端id
	UserId   string `json:"user_id"`   // 用户id

	Width  x_null.Int64 `json:"width"`  // 屏幕
	Height x_null.Int64 `json:"height"` // 屏幕高度

	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	Operator string `json:"operator"` // 电信运营商
	Ip       string `json:"ip"`       // ip地址
	// Ua       string // ua信息
	// ProjectKey string // 项目id
}
