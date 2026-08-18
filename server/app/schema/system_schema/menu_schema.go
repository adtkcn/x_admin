package system_schema

import "github.com/adtkcn/x_null"

// SystemAuthMenuDetailReq 菜单详情参数
type SystemAuthMenuDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthMenuAddReq 新增菜单参数
type SystemAuthMenuAddReq struct {
	Pid       string `json:"pid" form:"pid"`                                       // 上级菜单
	MenuType  string `json:"menu_type" form:"menu_type" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `json:"menu_name" form:"menu_name" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `json:"menu_icon" form:"menu_icon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `json:"menu_sort" form:"menu_sort" binding:"gte=0"`                 // 菜单排序
	Perms     string `json:"perms" form:"perms" binding:"max=10000"`                 // 权限标识
	Paths     string `json:"paths" form:"paths" binding:"max=200"`                   // 路由地址
	Component string `json:"component" form:"component" binding:"max=200"`               // 前端组件
	Selected  string `json:"selected" form:"selected" binding:"max=200"`                // 选中路径
	Params    string `json:"params" form:"params" binding:"max=200"`                  // 路由参数
	IsCache   uint8  `json:"is_cache" form:"is_cache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `json:"is_show" form:"is_show" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `json:"is_disable" form:"is_disable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

// SystemAuthMenuEditReq 编辑菜单参数
type SystemAuthMenuEditReq struct {
	ID        string `json:"id" form:"id" binding:"required"`                     // 主键
	Pid       string `json:"pid" form:"pid"`                                       // 上级菜单
	MenuType  string `json:"menu_type" form:"menu_type" binding:"oneof=M C A"`           // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName  string `json:"menu_name" form:"menu_name" binding:"required,min=1,max=30"` // 菜单名称
	MenuIcon  string `json:"menu_icon" form:"menu_icon" binding:"max=100"`               // 菜单图标
	MenuSort  int    `json:"menu_sort" form:"menu_sort" binding:"gte=0"`                 // 菜单排序
	Perms     string `json:"perms" form:"perms" binding:"max=10000"`                 // 权限标识
	Paths     string `json:"paths" form:"paths" binding:"max=200"`                   // 路由地址
	Component string `json:"component" form:"component" binding:"max=200"`               // 前端组件
	Selected  string `json:"selected" form:"selected" binding:"max=200"`                // 选中路径
	Params    string `json:"params" form:"params" binding:"max=200"`                  // 路由参数
	IsCache   uint8  `json:"is_cache" form:"is_cache" binding:"oneof=0 1"`              // 是否缓存: [0=否, 1=是]
	IsShow    uint8  `json:"is_show" form:"is_show" binding:"oneof=0 1"`               // 是否显示: [0=否, 1=是]
	IsDisable uint8  `json:"is_disable" form:"is_disable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
}

// SystemAuthMenuDelReq 删除菜单参数
type SystemAuthMenuDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthMenuResp 系统菜单返回信息
type SystemAuthMenuResp struct {
	ID         string                `json:"id"`                 // 主键
	Pid        string                `json:"pid"`                // 上级菜单
	MenuType   string                `json:"menu_type"`          // 权限类型: [M=目录, C=菜单, A=按钮]
	MenuName   string                `json:"menu_name"`          // 菜单名称
	MenuIcon   string                `json:"menu_icon"`          // 菜单图标
	MenuSort   uint16                `json:"menu_sort"`          // 菜单排序
	Perms      string                `json:"perms"`              // 权限标识
	Paths      string                `json:"paths"`              // 路由地址
	Component  string                `json:"component"`          // 前端组件
	Selected   string                `json:"selected"`           // 选中路径
	Params     string                `json:"params"`             // 路由参数
	IsCache    uint8                 `json:"is_cache"`           // 是否缓存: [0=否, 1=是]
	IsShow     uint8                 `json:"is_show"`            // 是否显示: [0=否, 1=是]
	IsDisable  uint8                 `json:"is_disable"`         // 是否禁用: [0=否, 1=是]
	CreateTime x_null.Time           `json:"create_time"`        // 创建时间
	UpdateTime x_null.Time           `json:"update_time"`        // 更新时间
	Children   []*SystemAuthMenuResp `json:"children,omitempty"` // 子集
}

func (n *SystemAuthMenuResp) GetID() string       { return n.ID }
func (n *SystemAuthMenuResp) GetParentID() string { return n.Pid }
func (n *SystemAuthMenuResp) SetChildren(children []*SystemAuthMenuResp) {
	n.Children = children
}
