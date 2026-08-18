package schema

import (
	"x_admin/app/schema/system_schema"

	"github.com/adtkcn/x_null"
)

type UserProtocolPrimarykey struct {
	Id string `json:"id" form:"id"`
}

// UserProtocolListReq 用户协议列表参数
type UserProtocolListReq struct {
	Title   x_null.String `json:"title" form:"title"`                   // 标题
	Content x_null.String `json:"content" form:"content"`               // 协议内容
	Version x_null.Int64  `json:"version" form:"version"`                // 排序

	CreateTimeStart x_null.String `json:"create_time_start" form:"create_time_start"` // 开始创建时间
	CreateTimeEnd   x_null.String `json:"create_time_end" form:"create_time_end"`     // 结束创建时间
	UpdateTimeStart x_null.String `json:"update_time_start" form:"update_time_start"` // 开始更新时间
	UpdateTimeEnd   x_null.String `json:"update_time_end" form:"update_time_end"`     // 结束更新时间
}

// UserProtocolAddReq 用户协议新增参数
type UserProtocolAddReq struct {
	Tag     x_null.String `json:"tag" form:"tag"`         // 标识
	Version x_null.Int64  `json:"version" form:"version"` // 版本
	Title   x_null.String `json:"title" form:"title"`     // 标题
	Content x_null.String `json:"content" form:"content"` // 协议内容
}

// UserProtocolEditReq 用户协议编辑参数
type UserProtocolEditReq struct {
	UserProtocolPrimarykey
	UserProtocolAddReq
}

// // UserProtocolBatchReq 用户协议批量删除参数
type UserProtocolDelBatchReq struct {
	Ids string `json:"ids" form:"ids"`
}

// UserProtocolResp 用户协议返回信息
type UserProtocolResp struct {
	UserProtocolPrimarykey
	Tag           x_null.String                           `json:"tag" swaggertype:"string"`       // 标识
	Version       x_null.Int64                            `json:"version" swaggertype:"number"`   // 版本
	Title         x_null.String                           `json:"title" swaggertype:"string"`     // 标题
	Content       x_null.String                           `json:"content" swaggertype:"string"`   // 协议内容
	CreateTime    x_null.Time                             `json:"create_time" swaggertype:"string"`   // 创建时间
	UpdateTime    x_null.Time                             `json:"update_time" swaggertype:"string"`   // 更新时间
	CreatedBy     x_null.String                           `json:"created_by" swaggertype:"string"`    // 创建人id
	CreatedByUser system_schema.SystemAuthAdminSimpleInfo `json:"created_by_user" swaggertype:"string"` // 创建人
}
