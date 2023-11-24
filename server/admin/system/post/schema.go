package post

import "x_admin/core"

//SystemAuthPostListReq 岗位列表参数
type SystemAuthPostListReq struct {
	Code   string `form:"code"`                                     // 岗位编码
	Name   string `form:"name"`                                     // 岗位名称
	IsStop int8   `form:"isStop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

//SystemAuthPostDetailReq 岗位详情参数
type SystemAuthPostDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthPostAddReq 岗位新增参数
type SystemAuthPostAddReq struct {
	Code    string `form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort    int    `form:"sort" binding:"gte=0"`                  // 排序编号
}

//SystemAuthPostEditReq 岗位编辑参数
type SystemAuthPostEditReq struct {
	ID      uint   `form:"id" binding:"required,gt=0"`            // 主键
	Code    string `form:"code" binding:"omitempty,min=1,max=30"` // 岗位编码
	Name    string `form:"name" binding:"required,min=1,max=30"`  // 岗位名称
	Remarks string `form:"remarks" binding:"max=250"`             // 岗位备注
	IsStop  uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort    int    `form:"sort" binding:"gte=0"`                  // 排序编号
}

//SystemAuthPostDelReq 岗位删除参数
type SystemAuthPostDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthPostResp 系统岗位返回信息
type SystemAuthPostResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Code       string      `json:"code" structs:"code"`             // 岗位编号
	Name       string      `json:"name" structs:"name"`             // 岗位名称
	Remarks    string      `json:"remarks" structs:"remarks"`       // 岗位备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 岗位排序
	IsStop     uint8       `json:"isStop" structs:"isStop"`         // 是否停用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
