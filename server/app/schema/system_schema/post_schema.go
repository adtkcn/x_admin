package system_schema

import "github.com/adtkcn/x_null"

// SystemAuthPostListReq 岗位列表参数
type SystemAuthPostListReq struct {
	Code   string `json:"code" form:"code"`                                      // 岗位编码
	Name   string `json:"name" form:"name"`                                      // 岗位名称
	IsStop int8   `json:"is_stop" form:"is_stop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

// SystemAuthPostDetailReq 岗位详情参数
type SystemAuthPostDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthPostAddReq 岗位新增参数
type SystemAuthPostAddReq struct {
	Code    string `json:"code" form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `json:"name" form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `json:"remarks" form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`           // 是否停用: [0=否, 1=是]
	Sort    int    `json:"sort" form:"sort" binding:"gte=0"`                  // 排序编号
}

// SystemAuthPostEditReq 岗位编辑参数
type SystemAuthPostEditReq struct {
	ID      string `json:"id" form:"id" binding:"required"`                 // 主键
	Code    string `json:"code" form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `json:"name" form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `json:"remarks" form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`           // 是否停用: [0=否, 1=是]
	Sort    int    `json:"sort" form:"sort" binding:"gte=0"`                  // 排序编号
}

// SystemAuthPostDelReq 岗位删除参数
type SystemAuthPostDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthPostResp 系统岗位返回信息
type SystemAuthPostResp struct {
	ID         string      `json:"id"`          // 主键
	Code       string      `json:"code"`        // 岗位编号
	Name       string      `json:"name"`        // 岗位名称
	Remarks    string      `json:"remarks"`     // 岗位备注
	Sort       uint16      `json:"sort"`        // 岗位排序
	IsStop     uint8       `json:"is_stop"`     // 是否停用: [0=否, 1=是]
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
