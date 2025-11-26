package schema

import (
	"x_admin/core"
)

type UserProtocolPrimarykey struct {
	Id string //
}

// UserProtocolListReq 用户协议列表参数
type UserProtocolListReq struct {
	Title   core.NullString // 标题
	Content core.NullString // 协议内容
	Version core.NullInt    // 排序

	CreateTimeStart core.NullString // 开始创建时间
	CreateTimeEnd   core.NullString // 结束创建时间
	UpdateTimeStart core.NullString // 开始更新时间
	UpdateTimeEnd   core.NullString // 结束更新时间
}

// UserProtocolAddReq 用户协议新增参数
type UserProtocolAddReq struct {
	Tag     core.NullString // 标识
	Version core.NullInt    // 版本
	Title   core.NullString // 标题
	Content core.NullString // 协议内容
}

// UserProtocolEditReq 用户协议编辑参数
type UserProtocolEditReq struct {
	UserProtocolPrimarykey
	UserProtocolAddReq
	// Tag     core.NullString // 标识
	// Version core.NullInt    // 版本
	// Title   core.NullString // 标题
	// Content core.NullString // 协议内容
}

// UserProtocolDetailReq 用户协议详情参数
// type UserProtocolDetailReq struct {
// 	Id string //
// }

// // UserProtocolDelReq 用户协议删除参数
// type UserProtocolDelReq struct {
// 	Id string //
// }

// // UserProtocolBatchReq 用户协议批量删除参数
type UserProtocolDelBatchReq struct {
	Ids string
}

// UserProtocolResp 用户协议返回信息
type UserProtocolResp struct {
	UserProtocolPrimarykey
	Tag     core.NullString // 标识
	Version core.NullInt    // 版本
	Title   core.NullString // 标题
	Content core.NullString // 协议内容
	// Sort       core.NullFloat // 排序
	CreateTime core.NullTime // 创建时间
	UpdateTime core.NullTime // 更新时间
}
