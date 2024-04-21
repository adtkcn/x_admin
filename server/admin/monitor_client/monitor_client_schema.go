package monitor_client

import "x_admin/core"

//MonitorClientListReq 客户端信息列表参数
type MonitorClientListReq struct {
	ProjectKey string `form:"projectKey"` // 项目key
	ClientId   string `form:"clientId"`   // sdk生成的客户端id
	UserId     string `form:"userId"`     // 用户id
	Os         string `form:"os"`         // 系统
	Browser    string `form:"browser"`    // 浏览器
	City       string `form:"city"`       // 城市
	Width      int    `form:"width"`      // 屏幕
	Height     int    `form:"height"`     // 屏幕高度
	Ua         string `form:"ua"`         // ua记录
	ClientTime string `form:"clientTime"` // 客户端时间
}

//MonitorClientDetailReq 客户端信息详情参数
type MonitorClientDetailReq struct {
	Id int `form:"id"` // uuid
}

//MonitorClientAddReq 客户端信息新增参数
type MonitorClientAddReq struct {
	ProjectKey string      `form:"projectKey"` // 项目key
	ClientId   string      `form:"clientId"`   // sdk生成的客户端id
	UserId     string      `form:"userId"`     // 用户id
	Os         string      `form:"os"`         // 系统
	Browser    string      `form:"browser"`    // 浏览器
	City       string      `form:"city"`       // 城市
	Width      int         `form:"width"`      // 屏幕
	Height     int         `form:"height"`     // 屏幕高度
	Ua         string      `form:"ua"`         // ua记录
	ClientTime core.TsTime `form:"clientTime"` // 客户端时间
}

//MonitorClientEditReq 客户端信息编辑参数
type MonitorClientEditReq struct {
	Id         int         `form:"id"`         // uuid
	ProjectKey string      `form:"projectKey"` // 项目key
	ClientId   string      `form:"clientId"`   // sdk生成的客户端id
	UserId     string      `form:"userId"`     // 用户id
	Os         string      `form:"os"`         // 系统
	Browser    string      `form:"browser"`    // 浏览器
	City       string      `form:"city"`       // 城市
	Width      int         `form:"width"`      // 屏幕
	Height     int         `form:"height"`     // 屏幕高度
	Ua         string      `form:"ua"`         // ua记录
	ClientTime core.TsTime `form:"clientTime"` // 客户端时间
}

//MonitorClientDelReq 客户端信息新增参数
type MonitorClientDelReq struct {
	Id int `form:"id"` // uuid
}

//MonitorClientResp 客户端信息返回信息
type MonitorClientResp struct {
	Id         int         `json:"id" structs:"id" excel:"name:uuid;"`                    // uuid
	ProjectKey string      `json:"projectKey" structs:"projectKey" excel:"name:项目key;"`   // 项目key
	ClientId   string      `json:"clientId" structs:"clientId" excel:"name:sdk生成的客户端id;"` // sdk生成的客户端id
	UserId     string      `json:"userId" structs:"userId" excel:"name:用户id;"`            // 用户id
	Os         string      `json:"os" structs:"os" excel:"name:系统;"`                      // 系统
	Browser    string      `json:"browser" structs:"browser" excel:"name:浏览器;"`           // 浏览器
	City       string      `json:"city" structs:"city" excel:"name:城市;"`                  // 城市
	Width      int         `json:"width" structs:"width" excel:"name:屏幕;"`                // 屏幕
	Height     int         `json:"height" structs:"height" excel:"name:屏幕高度;"`            // 屏幕高度
	Ua         string      `json:"ua" structs:"ua" excel:"name:ua记录;"`                    // ua记录
	ClientTime int         `json:"clientTime" structs:"clientTime" excel:"name:客户端时间;"`   // 客户端时间
	CreateTime core.TsTime `json:"createTime" structs:"createTime" excel:"name:创建时间;"`    // 创建时间
}
