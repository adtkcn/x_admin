package systemSchema

import "x_admin/core"

// SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Username string `form:"username"` // 账号
	Nickname string `form:"nickname"` // 昵称
	RoleId   string `form:"roleId"`   // 角色ID
}

// SystemAuthAdminDetailReq 管理员详情参数
type SystemAuthAdminDetailReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId string `form:"deptId" binding:""` // 部门ID
	PostId string `form:"postId" binding:""` // 岗位ID
	RoleId string `form:"roleId" binding:""` // 角色ID

	Username string `form:"username" binding:"required,min=2,max=32"` // 账号
	Nickname string `form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Password string `form:"password" binding:"required"`              // 密码
	Avatar   string `form:"avatar" binding:""`                        // 头像

	Sort      int   `form:"sort" binding:"gte=0"`          // 排序
	IsDisable uint8 `form:"isDisable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]

}

// SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID       string `form:"id" binding:"required"`                    // 主键
	DeptId   string `form:"deptId" binding:""`                        // 部门ID
	PostId   string `form:"postId" binding:""`                        // 岗位ID
	RoleId   string `form:"roleId" binding:""`                        // 角色ID
	Username string `form:"username" binding:"required,min=2,max=32"` // 账号
	Nickname string `form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Password string `form:"password"`                                 // 密码
	Avatar   string `form:"avatar"`                                   // 头像

	Sort      int   `form:"sort" binding:"gte=0"`          // 排序
	IsDisable uint8 `form:"isDisable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]
}

// SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Avatar       string `form:"avatar"`                                   // 头像
	Password     string `form:"password" binding:""`                      // 密码
	CurrPassword string `form:"currPassword" binding:""`                  // 密码
}

// SystemAuthAdminDelReq 管理员删除参数
type SystemAuthAdminDelReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemAuthAdminDisableReq 管理员状态切换参数
type SystemAuthAdminDisableReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemAuthAdminResp 管理员返回信息
type SystemAuthAdminResp struct {
	ID            string        `json:"id"`                                 // 主键
	Username      string        `json:"username" excel:"name:账号;"`          // 账号
	Nickname      string        `json:"nickname" excel:"name:昵称;"`          // 昵称
	Avatar        string        `json:"avatar" excel:"name:头像;"`            // 头像
	Role          string        `json:"role" excel:"name:角色;"`              // 角色
	DeptId        string        `json:"deptId" excel:"name:部门ID;"`          // 部门ID
	PostId        string        `json:"postId" excel:"name:岗位ID;"`          // 岗位ID
	RoleId        string        `json:"roleId" excel:"name:角色ID;"`          // 角色ID
	Dept          string        `json:"dept" excel:"name:部门;"`              // 部门
	IsDisable     uint8         `json:"isDisable" excel:"name:是否禁用;"`       // 是否禁用: [0=否, 1=是]
	LastLoginIp   string        `json:"lastLoginIp" excel:"name:最后登录IP;"`   // 最后登录IP
	LastLoginTime core.NullTime `json:"lastLoginTime" excel:"name:最后登录时间;"` // 最后登录时间
	CreateTime    core.NullTime `json:"createTime" excel:"name:创建时间;"`      // 创建时间
	UpdateTime    core.NullTime `json:"updateTime" excel:"name:更新时间;"`      // 更新时间
}

// SystemAuthAdminSelfOneResp 当前管理员返回部分信息
type SystemAuthAdminSelfOneResp struct {
	ID            string        `json:"id"`            // 主键
	Username      string        `json:"username"`      // 账号
	Nickname      string        `json:"nickname"`      // 昵称
	Avatar        string        `json:"avatar"`        // 头像
	Role          string        `json:"role"`          // 角色
	Dept          string        `json:"dept"`          // 部门
	IsDisable     uint8         `json:"isDisable"`     // 是否禁用: [0=否, 1=是]
	LastLoginIp   string        `json:"lastLoginIp"`   // 最后登录IP
	LastLoginTime core.NullTime `json:"lastLoginTime"` // 最后登录时间
	CreateTime    core.NullTime `json:"createTime"`    // 创建时间
	UpdateTime    core.NullTime `json:"updateTime"`    // 更新时间
}

// SystemAuthAdminSelfResp 当前系统管理员返回信息
type SystemAuthAdminSelfResp struct {
	User        SystemAuthAdminSelfOneResp `json:"user"`        // 用户信息
	Permissions []string                   `json:"permissions"` // 权限集合: [[*]=>所有权限, ['article:add']=>部分权限]
}

type SystemAuthAdminSimpleInfo struct {
	ID string `json:"id"` // 主键
	// Username string `json:"username"` // 账号
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
}
