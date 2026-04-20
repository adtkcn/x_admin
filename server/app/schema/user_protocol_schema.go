package schema

import (
	"x_admin/app/schema/system_schema"

	"github.com/adtkcn/x_null"
)

type UserProtocolPrimarykey struct {
	Id string //
}

// UserProtocolListReq 用户协议列表参数
type UserProtocolListReq struct {
	Title   x_null.String // 标题
	Content x_null.String // 协议内容
	Version x_null.Int64  // 排序

	CreateTimeStart x_null.String // 开始创建时间
	CreateTimeEnd   x_null.String // 结束创建时间
	UpdateTimeStart x_null.String // 开始更新时间
	UpdateTimeEnd   x_null.String // 结束更新时间
}

// UserProtocolAddReq 用户协议新增参数
type UserProtocolAddReq struct {
	Tag     x_null.String // 标识
	Version x_null.Int64  // 版本
	Title   x_null.String // 标题
	Content x_null.String // 协议内容
}

// UserProtocolEditReq 用户协议编辑参数
type UserProtocolEditReq struct {
	UserProtocolPrimarykey
	UserProtocolAddReq
}

// // UserProtocolBatchReq 用户协议批量删除参数
type UserProtocolDelBatchReq struct {
	Ids string
}

// UserProtocolResp 用户协议返回信息
type UserProtocolResp struct {
	UserProtocolPrimarykey
	Tag           x_null.String                          `swaggertype:"string"` // 标识
	Version       x_null.Int64                           `swaggertype:"number"` // 版本
	Title         x_null.String                          `swaggertype:"string"` // 标题
	Content       x_null.String                          `swaggertype:"string"` // 协议内容
	CreateTime    x_null.Time                            `swaggertype:"string"` // 创建时间
	UpdateTime    x_null.Time                            `swaggertype:"string"` // 更新时间
	CreatedBy     x_null.String                          `swaggertype:"string"` // 创建人id
	CreatedByUser system_schema.SystemAuthAdminSimpleInfo `swaggertype:"string"` // 创建人
}
