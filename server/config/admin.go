package config

import "github.com/gin-gonic/gin"

// AdminConfig 后台公共配置
var AdminConfig = adminConfig{
	// 管理缓存键
	BackstageManageKey: "backstage:manage",
	// 角色缓存键
	BackstageRolesKey: "backstage:roles",
	// 令牌缓存键
	BackstageTokenKey: "backstage:token:",
	// 令牌的集合
	BackstageTokenSet: "backstage:token:set:",
	// #region NotAuth
	// 免登录验证
	NotLoginUri: []string{
		// "admin:system:login",        // 登录接口
		// "admin:common:index:config", // 配置接口
	},

	// 免接口权限验证
	NotAuthUri: []string{
		// "admin:system:logout",     // 退出登录
		// "admin:system:menu:menus", // 系统菜单
		// "admin:system:menu:route", // 菜单路由
		// "admin:system:admin:upInfo", // 管理员更新
		// "admin:system:admin:self",     // 管理员信息
		// "admin:system:role:all", // 所有角色
		// "admin:system:post:all", // 所有岗位
		// "admin:system:dept:list", // 所有部门
		// "admin:setting:dict:type:all", // 所有字典类型
		// "admin:setting:dict:data:all", // 所有字典数据
	},
	// #endregion NotAuth
	// 演示模式白名单
	ShowWhitelistUri: []string{
		// "admin:system:login",  // 登录接口
		// "admin:system:logout", // 退出登录
	},

	// 管理员账号id:1
	SuperAdminId: "1",
	// 管理员账号key
	ReqAdminIdKey: "admin_id",
	// 角色key
	ReqRoleIdKey: "role",
	// 用户名key
	ReqUsernameKey: "username",
	// 昵称key
	ReqNicknameKey: "nickname",
}

type adminConfig struct {
	// 管理缓存键"backstage:manage"
	BackstageManageKey string
	// 角色缓存键"backstage:roles"
	BackstageRolesKey string
	// 令牌缓存键"backstage:token:"
	BackstageTokenKey string
	// 令牌的集合 "backstage:token:set:"
	BackstageTokenSet string
	// 免登录验证
	NotLoginUri []string
	// 免权限验证
	NotAuthUri []string
	// 演示模式白名单
	ShowWhitelistUri []string
	// 管理员账号id:1
	SuperAdminId string
	// 管理员账号key
	ReqAdminIdKey string
	// 角色key
	ReqRoleIdKey string
	// 用户名key
	ReqUsernameKey string
	// 昵称key
	ReqNicknameKey string
}

func (cnf adminConfig) GetAdminId(c *gin.Context) string {
	adminId, ok := c.Get(cnf.ReqAdminIdKey)
	if !ok {
		return ""
	}
	return adminId.(string)
}

func (cnf adminConfig) GetRoleId(c *gin.Context) string {
	roleId, ok := c.Get(cnf.ReqRoleIdKey)
	if !ok {
		return ""
	}
	return roleId.(string)
}

func (cnf adminConfig) GetUsername(c *gin.Context) string {
	username, ok := c.Get(cnf.ReqUsernameKey)
	if !ok {
		return ""
	}
	return username.(string)
}

func (cnf adminConfig) GetNickname(c *gin.Context) string {
	nickname, ok := c.Get(cnf.ReqNicknameKey)
	if !ok {
		return ""
	}
	return nickname.(string)
}
