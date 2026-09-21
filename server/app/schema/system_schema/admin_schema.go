package system_schema

import "github.com/adtkcn/x_null"

// SystemAuthAdminListReq 管理员列表参数
type SystemAuthAdminListReq struct {
	Email    string `json:"email" form:"email"`       // 邮箱(账号)
	Nickname string `json:"nickname" form:"nickname"` // 昵称
	RoleId   string `json:"role_id" form:"role_id"`   // 角色ID(用于筛选)
}

// SystemAuthAdminDetailReq 管理员详情参数
type SystemAuthAdminDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthAdminAddReq 管理员新增参数
type SystemAuthAdminAddReq struct {
	DeptId  string   `json:"dept_id" form:"dept_id" binding:""`   // 部门ID
	PostId  string   `json:"post_id" form:"post_id" binding:""`   // 岗位ID
	RoleIds []string `json:"role_ids" form:"role_ids" binding:""` // 角色ID列表

	Email    string `json:"email" form:"email" binding:"required,min=5,max=200"`      // 邮箱(账号)
	Nickname string `json:"nickname" form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Password string `json:"password" form:"password" binding:"required"`              // 密码
	Avatar   string `json:"avatar" form:"avatar" binding:""`                          // 头像

	Sort      int   `json:"sort" form:"sort" binding:"gte=0"`                 // 排序
	IsDisable uint8 `json:"is_disable" form:"is_disable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]

}

// SystemAuthAdminTodayCountResp 管理员今日新增用户数量和总用户数量
type SystemAuthAdminTodayCountResp struct {
	TotalUsers int64 `json:"total_users"` // 总用户数量
	TodayUsers int64 `json:"today_users"` // 今日新增用户数量
}

// SystemAuthAdminEditReq 管理员编辑参数
type SystemAuthAdminEditReq struct {
	ID       string   `json:"id" form:"id" binding:"required"`                          // 主键
	DeptId   string   `json:"dept_id" form:"dept_id" binding:""`                        // 部门ID
	PostId   string   `json:"post_id" form:"post_id" binding:""`                        // 岗位ID
	RoleIds  []string `json:"role_ids" form:"role_ids" binding:""`                      // 角色ID列表
	Email    string   `json:"email" form:"email" binding:"required,min=5,max=200"`      // 邮箱(账号)
	Nickname string   `json:"nickname" form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Password string   `json:"password" form:"password"`                                 // 密码
	Avatar   string   `json:"avatar" form:"avatar"`                                     // 头像

	Sort      int   `json:"sort" form:"sort" binding:"gte=0"`                 // 排序
	IsDisable uint8 `json:"is_disable" form:"is_disable" binding:"oneof=0 1"` // 是否禁用: [0=否, 1=是]
}

// SystemAuthAdminSendEmailCodeReq 发送邮箱验证码参数
type SystemAuthAdminSendEmailCodeReq struct {
	Email string `json:"email" form:"email" binding:"required,email,min=5,max=200"` // 目标邮箱
}

// SystemAuthAdminUpdateReq 管理员更新参数
type SystemAuthAdminUpdateReq struct {
	Nickname     string `json:"nickname" form:"nickname" binding:"required,min=2,max=32"` // 昵称
	Avatar       string `json:"avatar" form:"avatar"`                                     // 头像
	Email        string `json:"email" form:"email" binding:""`                            // 邮箱
	EmailCode    string `json:"email_code" form:"email_code" binding:""`                  // 邮箱验证码（改邮箱时必填）
	Password     string `json:"password" form:"password" binding:""`                      // 密码
	CurrPassword string `json:"curr_password" form:"curr_password" binding:""`            // 密码
}

// SystemAuthAdminDelReq 管理员删除参数
type SystemAuthAdminDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthAdminDisableReq 管理员状态切换参数
type SystemAuthAdminDisableReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthAdminResp 管理员返回信息
type SystemAuthAdminResp struct {
	ID       string `json:"id"`                        // 主键
	Email    string `json:"email" excel:"name:账号;"`    // 邮箱(账号)
	Nickname string `json:"nickname" excel:"name:昵称;"` // 昵称
	Avatar   string `json:"avatar" excel:"name:头像;"`   // 头像

	DeptId string `json:"dept_id" excel:"name:部门ID;"` // 部门ID
	Dept   string `json:"dept" excel:"name:部门;"`      // 部门

	PostId string `json:"post_id" excel:"name:岗位ID;"` // 岗位ID
	Post   string `json:"post" excel:"name:岗位;"`      // 岗位名称

	RoleIds []string `json:"role_ids" gorm:"-"`     // 角色ID列表
	Role    string   `json:"role" excel:"name:角色;"` // 角色名称(逗号分隔)

	IsDisable     uint8       `json:"is_disable" excel:"name:是否禁用;"`        // 是否禁用: [0=否, 1=是]
	LastLoginIp   string      `json:"last_login_ip" excel:"name:最后登录IP;"`   // 最后登录IP
	LastLoginTime x_null.Time `json:"last_login_time" excel:"name:最后登录时间;"` // 最后登录时间
	CreateTime    x_null.Time `json:"create_time" excel:"name:创建时间;"`       // 创建时间
	UpdateTime    x_null.Time `json:"update_time" excel:"name:更新时间;"`       // 更新时间
}

// SystemAuthAdminSelfOneResp 当前管理员返回部分信息
type SystemAuthAdminSelfOneResp struct {
	ID            string      `json:"id"`              // 主键
	Email         string      `json:"email"`           // 邮箱(账号)
	Nickname      string      `json:"nickname"`        // 昵称
	Avatar        string      `json:"avatar"`          // 头像
	Role          string      `json:"role"`            // 角色
	Dept          string      `json:"dept"`            // 部门
	IsDisable     uint8       `json:"is_disable"`      // 是否禁用: [0=否, 1=是]
	LastLoginIp   string      `json:"last_login_ip"`   // 最后登录IP
	LastLoginTime x_null.Time `json:"last_login_time"` // 最后登录时间
	CreateTime    x_null.Time `json:"create_time"`     // 创建时间
	UpdateTime    x_null.Time `json:"update_time"`     // 更新时间
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
