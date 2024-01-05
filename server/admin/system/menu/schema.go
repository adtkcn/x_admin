package menu

import "x_admin/core"

//SystemAuthMenuDetailReq 菜单详情参数
type SystemAuthMenuDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthMenuAddReq 新增菜单参数
type SystemAuthMenuAddReq struct {
	Pid       uint   `form:"pid" binding:"gte=0"`                      // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"gte=0"`                 // 菜单排序
	Perms     string `form:"perms" binding:"max=10000"`                // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   uint8  `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

//SystemAuthMenuEditReq 编辑菜单参数
type SystemAuthMenuEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`               // 主键
	Pid       uint   `form:"pid" binding:"gte=0"`                      // 上级菜单
	MenuType  string `form:"menuType" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `form:"menuName" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `form:"menuIcon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `form:"menuSort" binding:"gte=0"`                 // 菜单排序
	Perms     string `form:"perms" binding:"max=10000"`                // 权限标识
	Paths     string `form:"paths" binding:"max=200"`                  // 路由地址
	Component string `form:"component" binding:"max=200"`              // 前端组件
	Selected  string `form:"selected" binding:"max=200"`               // 选中路径
	Params    string `form:"params" binding:"max=200"`                 // 路由参数
	IsCache   uint8  `form:"isCache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `form:"isShow" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

//SystemAuthMenuDelReq 删除菜单参数
type SystemAuthMenuDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthMenuResp 系统菜单返回信息
type SystemAuthMenuResp struct {
	ID         uint                 `json:"id" structs:"id"`                       // 主键
	Pid        uint                 `json:"pid" structs:"pid"`                     // 上级菜单
	MenuType   string               `json:"menuType" structs:"menuType"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName   string               `json:"menuName" structs:"menuName"`           // 菜单名称
	MenuIcon   string               `json:"menuIcon" structs:"menuIcon"`           // 菜单图标
	MenuSort   uint16               `json:"menuSort" structs:"menuSort"`           // 菜单排序
	Perms      string               `json:"perms" structs:"perms"`                 // 权限标识
	Paths      string               `json:"paths" structs:"paths"`                 // 路由地址
	Component  string               `json:"component" structs:"component"`         // 前端组件
	Selected   string               `json:"selected" structs:"selected"`           // 选中路径
	Params     string               `json:"params" structs:"params"`               // 路由参数
	IsCache    uint8                `json:"isCache" structs:"isCache"`             // 是否缓存: [0=否, 1=是]
	IsShow     uint8                `json:"isShow" structs:"isShow"`               // 是否显示: [0=否, 1=是]
	IsDisable  uint8                `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime          `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime core.TsTime          `json:"updateTime" structs:"updateTime"`       // 更新时间
	Children   []SystemAuthMenuResp `json:"children,omitempty" structs:"children"` // 子集
}
