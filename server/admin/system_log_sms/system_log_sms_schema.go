package system_log_sms

import (
	"x_admin/core"
)

// SystemLogSmsListReq 系统短信日志列表参数
type SystemLogSmsListReq struct {
	Scene           *int    `` // 场景编号
	Mobile          *string `` // 手机号码
	Content         *string `` // 发送内容
	Status          *int    `` // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results         *string `` // 短信结果
	SendTimeStart   *string `` // 开始发送时间
	SendTimeEnd     *string `` // 结束发送时间
	CreateTimeStart *string `` // 开始创建时间
	CreateTimeEnd   *string `` // 结束创建时间
	UpdateTimeStart *string `` // 开始更新时间
	UpdateTimeEnd   *string `` // 结束更新时间
}

// SystemLogSmsAddReq 系统短信日志新增参数
type SystemLogSmsAddReq struct {
	Scene    core.NullInt  `` // 场景编号
	Mobile   *string       `` // 手机号码
	Content  *string       `` // 发送内容
	Status   core.NullInt  `` // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results  *string       `` // 短信结果
	SendTime core.NullTime `` // 发送时间
}

// SystemLogSmsEditReq 系统短信日志编辑参数
type SystemLogSmsEditReq struct {
	Id       int           `` // id
	Scene    core.NullInt  `` // 场景编号
	Mobile   *string       `` // 手机号码
	Content  *string       `` // 发送内容
	Status   core.NullInt  `` // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results  *string       `` // 短信结果
	SendTime core.NullTime `` // 发送时间
}

// SystemLogSmsDetailReq 系统短信日志详情参数
type SystemLogSmsDetailReq struct {
	Id int `` // id
}

// SystemLogSmsDelReq 系统短信日志删除参数
type SystemLogSmsDelReq struct {
	Id int `` // id
}

// SystemLogSmsResp 系统短信日志返回信息
type SystemLogSmsResp struct {
	Id         int           `` // id
	Scene      core.NullInt  `` // 场景编号
	Mobile     string        `` // 手机号码
	Content    string        `` // 发送内容
	Status     core.NullInt  `` // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string        `` // 短信结果
	SendTime   core.NullTime `` // 发送时间
	CreateTime core.NullTime `` // 创建时间
	UpdateTime core.NullTime `` // 更新时间
}
