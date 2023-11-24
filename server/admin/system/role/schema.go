package role

import "x_admin/core"

//SystemAuthRoleSimpleResp 系统角色返回简单信息
type SystemAuthRoleSimpleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

//SystemAuthRoleResp 系统角色返回信息
type SystemAuthRoleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	Remark     string      `json:"remark" structs:"remark"`         // 角色备注
	Menus      []uint      `json:"menus" structs:"menus"`           // 关联菜单
	Member     int64       `json:"member" structs:"member"`         // 成员数量
	Sort       uint16      `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

//

//SystemAuthRoleDetailReq 角色详情参数
type SystemAuthRoleDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthRoleAddReq 新增角色参数
type SystemAuthRoleAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

//SystemAuthRoleEditReq 编辑角色参数
type SystemAuthRoleEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色备注
	MenuIds   string `form:"menuIds"`                              // 关联菜单
}

//SystemAuthRoleDelReq 删除角色参数
type SystemAuthRoleDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}
