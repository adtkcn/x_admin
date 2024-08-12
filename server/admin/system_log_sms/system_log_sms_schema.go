package system_log_sms

import (
	"x_admin/core"
)

// SystemLogSmsListReq 系统短信日志列表参数
type SystemLogSmsListReq struct {
	Scene           *int    `form:"scene"`            // 场景编号
	Mobile          *string `form:"mobile"`           // 手机号码
	Content         *string `form:"content"`          // 发送内容
	Status          *int    `form:"status"`           // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results         *string `form:"results"`          // 短信结果
	SendTime        *int    `form:"send_time"`        // 发送时间
	CreateTimeStart *string `form:"create_timeStart"` // 开始创建时间
	CreateTimeEnd   *string `form:"create_timeEnd"`   // 结束创建时间
	UpdateTimeStart *string `form:"update_timeStart"` // 开始更新时间
	UpdateTimeEnd   *string `form:"update_timeEnd"`   // 结束更新时间
}

// SystemLogSmsAddReq 系统短信日志新增参数
type SystemLogSmsAddReq struct {
	Scene    interface{} `mapstructure:"scene" json:"scene" form:"scene"`             // 场景编号
	Mobile   interface{} `mapstructure:"mobile" json:"mobile" form:"mobile"`          // 手机号码
	Content  interface{} `mapstructure:"content" json:"content" form:"content"`       // 发送内容
	Status   interface{} `mapstructure:"status" json:"status" form:"status"`          // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results  interface{} `mapstructure:"results" json:"results" form:"results"`       // 短信结果
	SendTime interface{} `mapstructure:"send_time" json:"send_time" form:"send_time"` // 发送时间
}

// SystemLogSmsEditReq 系统短信日志编辑参数
type SystemLogSmsEditReq struct {
	Id       int         `mapstructure:"id" json:"id" form:"id"`                      // id
	Scene    interface{} `mapstructure:"scene" json:"scene" form:"scene"`             // 场景编号
	Mobile   interface{} `mapstructure:"mobile" json:"mobile" form:"mobile"`          // 手机号码
	Content  interface{} `mapstructure:"content" json:"content" form:"content"`       // 发送内容
	Status   interface{} `mapstructure:"status" json:"status" form:"status"`          // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results  interface{} `mapstructure:"results" json:"results" form:"results"`       // 短信结果
	SendTime interface{} `mapstructure:"send_time" json:"send_time" form:"send_time"` // 发送时间
}

// SystemLogSmsDetailReq 系统短信日志详情参数
type SystemLogSmsDetailReq struct {
	Id int `form:"id"` // id
}

// SystemLogSmsDelReq 系统短信日志删除参数
type SystemLogSmsDelReq struct {
	Id int `form:"id"` // id
}

// SystemLogSmsResp 系统短信日志返回信息
type SystemLogSmsResp struct {
	Id         int         `json:"id"`                             // id
	Scene      int         `json:"scene" excel:"name:场景编号;"`       // 场景编号
	Mobile     string      `json:"mobile" excel:"name:手机号码;"`      // 手机号码
	Content    string      `json:"content" excel:"name:发送内容;"`     // 发送内容
	Status     int         `json:"status" excel:"name:发送状态;"`      // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string      `json:"results" excel:"name:短信结果;"`     // 短信结果
	SendTime   int         `json:"send_time" excel:"name:发送时间;"`   // 发送时间
	CreateTime core.TsTime `json:"create_time" excel:"name:创建时间;"` // 创建时间
	UpdateTime core.TsTime `json:"update_time" excel:"name:更新时间;"` // 更新时间
}
