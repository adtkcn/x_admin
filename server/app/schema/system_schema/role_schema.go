package system_schema

import "github.com/adtkcn/x_null"

// SystemAuthRoleSimpleResp 系统角色返回简单信息
type SystemAuthRoleSimpleResp struct {
	ID         string      `json:"id"`          // 主键
	Name       string      `json:"name"`        // 角色名称
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}

// SystemAuthRoleResp 系统角色返回信息
type SystemAuthRoleResp struct {
	ID         string      `json:"id"`          // 主键
	Name       string      `json:"name"`        // 角色名称
	Remark     string      `json:"remark"`      // 角色备注
	Menus      []string    `json:"menus"`       // 关联菜单
	Member     int64       `json:"member"`      // 成员数量
	Sort       uint16      `json:"sort"`        // 角色排序
	IsDisable  uint8       `json:"is_disable"`  // 是否禁用: [0=否, 1=是]
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}

//

// SystemAuthRoleDetailReq 角色详情参数
type SystemAuthRoleDetailReq struct {
	ID string `json:"id" form:"id" binding:"required,gt=0"` // 主键
}

// SystemAuthRoleAddReq 新增角色参数
type SystemAuthRoleAddReq struct {
	Name      string `json:"name" form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `json:"sort" form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `json:"is_disable" form:"is_disable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]
	Remark    string `json:"remark" form:"remark" binding:"max=200"`           // 角色备注
	MenuIds   string `json:"menuIds" form:"menuIds"`                           // 关联菜单
}

// SystemAuthRoleEditReq 编辑角色参数
type SystemAuthRoleEditReq struct {
	ID        string `json:"id" form:"id" binding:"required,gt=0"`             // 主键
	Name      string `json:"name" form:"name" binding:"required,min=1,max=30"` // 角色名称
	Sort      int    `json:"sort" form:"sort" binding:"gte=0"`                 // 角色排序
	IsDisable uint8  `json:"is_disable" form:"is_disable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]
	Remark    string `json:"remark" form:"remark" binding:"max=200"`           // 角色备注
	MenuIds   string `json:"menuIds" form:"menuIds"`                           // 关联菜单
}

// SystemAuthRoleDelReq 删除角色参数
type SystemAuthRoleDelReq struct {
	ID string `json:"id" form:"id" binding:"required,gt=0"` // 主键
}
