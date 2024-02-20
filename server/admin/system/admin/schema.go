package admin

import "x_admin/core"

//SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Username string `form:"username"`        // 账号
	Nickname string `form:"nickname"`        // 昵称
	Role     int    `form:"role,default=-1"` // 角色ID
}

//SystemAuthAdminDetailReq 管理员详情参数
type SystemAuthAdminDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId       uint   `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       uint   `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password" binding:"required"`              // 密码
	Avatar       string `form:"avatar" binding:"required"`                // 头像
	Role         uint   `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

//SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID           uint   `form:"id" binding:"required,gt=0"`               // 主键
	DeptId       uint   `form:"deptId" binding:"required,gt=0"`           // 部门ID
	PostId       uint   `form:"postId" binding:"required,gt=0"`           // 岗位ID
	Username     string `form:"username" binding:"required,min=2,max=20"` // 账号
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Password     string `form:"password"`                                 // 密码
	Avatar       string `form:"avatar"`                                   // 头像
	Role         uint   `form:"role" binding:"gte=0"`                     // 角色
	Sort         int    `form:"sort" binding:"gte=0"`                     // 排序
	IsDisable    uint8  `form:"isDisable" binding:"oneof=0 1"`            // 是否禁用: [0=否, 1=是]
	IsMultipoint uint8  `form:"isMultipoint" binding:"oneof=0 1"`         // 多端登录: [0=否, 1=是]
}

//SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `form:"nickname" binding:"required,min=2,max=30"` // 昵称
	Avatar       string `form:"avatar"`                                   // 头像
	Password     string `form:"password" binding:""`                      // 密码
	CurrPassword string `form:"currPassword" binding:""`                  // 密码
}

//SystemAuthAdminDelReq 管理员删除参数
type SystemAuthAdminDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthAdminDisableReq 管理员状态切换参数
type SystemAuthAdminDisableReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthAdminResp 管理员返回信息
type SystemAuthAdminResp struct {
	ID            uint        `json:"id" structs:"id"`                                            // 主键
	Username      string      `json:"username" structs:"username" excel:"name:账号;"`               // 账号
	Nickname      string      `json:"nickname" structs:"nickname" excel:"name:昵称;"`               // 昵称
	Avatar        string      `json:"avatar" structs:"avatar" excel:"name:头像;"`                   // 头像
	Role          string      `json:"role" structs:"role" excel:"name:角色;"`                       // 角色
	DeptId        uint        `json:"deptId" structs:"deptId" excel:"name:部门ID;"`                 // 部门ID
	PostId        uint        `json:"postId" structs:"postId" excel:"name:岗位ID;"`                 // 岗位ID
	Dept          string      `json:"dept" structs:"dept" excel:"name:部门;"`                       // 部门
	IsMultipoint  uint8       `json:"isMultipoint" structs:"isMultipoint" excel:"name:多端登录;"`     // 多端登录: [0=否, 1=是]
	IsDisable     uint8       `json:"isDisable" structs:"isDisable" excel:"name:是否禁用;"`           // 是否禁用: [0=否, 1=是]
	LastLoginIp   string      `json:"lastLoginIp" structs:"lastLoginIp" excel:"name:最后登录IP;"`     // 最后登录IP
	LastLoginTime core.TsTime `json:"lastLoginTime" structs:"lastLoginTime" excel:"name:最后登录时间;"` // 最后登录时间
	CreateTime    core.TsTime `json:"createTime" structs:"createTime" excel:"name:创建时间;"`         // 创建时间
	UpdateTime    core.TsTime `json:"updateTime" structs:"updateTime" excel:"name:更新时间;"`         // 更新时间
}

//SystemAuthAdminSelfOneResp 当前管理员返回部分信息
type SystemAuthAdminSelfOneResp struct {
	ID            uint        `json:"id" structs:"id"`                       // 主键
	Username      string      `json:"username" structs:"username"`           // 账号
	Nickname      string      `json:"nickname" structs:"nickname"`           // 昵称
	Avatar        string      `json:"avatar" structs:"avatar"`               // 头像
	Role          string      `json:"role" structs:"role"`                   // 角色
	Dept          string      `json:"dept" structs:"dept"`                   // 部门
	IsMultipoint  uint8       `json:"isMultipoint" structs:"isMultipoint"`   // 多端登录: [0=否, 1=是]
	IsDisable     uint8       `json:"isDisable" structs:"isDisable"`         // 是否禁用: [0=否, 1=是]
	LastLoginIp   string      `json:"lastLoginIp" structs:"lastLoginIp"`     // 最后登录IP
	LastLoginTime core.TsTime `json:"lastLoginTime" structs:"lastLoginTime"` // 最后登录时间
	CreateTime    core.TsTime `json:"createTime" structs:"createTime"`       // 创建时间
	UpdateTime    core.TsTime `json:"updateTime" structs:"updateTime"`       // 更新时间
}

//SystemAuthAdminSelfResp 当前系统管理员返回信息
type SystemAuthAdminSelfResp struct {
	User        SystemAuthAdminSelfOneResp `json:"user" structs:"user"`               // 用户信息
	Permissions []string                   `json:"permissions" structs:"permissions"` // 权限集合: [[*]=>所有权限, ['article:add']=>部分权限]
}
