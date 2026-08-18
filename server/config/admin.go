package config

import "github.com/gin-gonic/gin"

// AdminConfig 后台公共配置
var AdminConfig = adminConfig{
	// 用户缓存键 hash
	BackstageAdminKey: "admin:users",
	// 用户权限缓存键(菜单+按钮) hash
	BackstageAdminPermsKey: "admin:perms",
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

	// 管理员账号id:1
	SuperAdminId: "1",
	// 管理员账号key
	ReqAdminIdKey: "admin_id",

	// 邮箱key(账号)
	ReqEmailKey: "email",
	// 昵称key
	ReqNicknameKey: "nickname",
}

type adminConfig struct {
	// 管理缓存键"backstage:admin:users"，field为管理员id，value为管理员信息
	BackstageAdminKey string

	// 用户权限缓存键"backstage:admin:perms"，field为管理员id，value为权限列表(逗号分隔)
	BackstageAdminPermsKey string
	// 免登录验证
	NotLoginUri []string
	// 免权限验证
	NotAuthUri []string

	// 管理员账号id:1
	SuperAdminId string
	// 管理员账号key
	ReqAdminIdKey string

	// 邮箱key(账号)
	ReqEmailKey string
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

func (cnf adminConfig) GetEmail(c *gin.Context) string {
	email, ok := c.Get(cnf.ReqEmailKey)
	if !ok {
		return ""
	}
	return email.(string)
}

func (cnf adminConfig) GetNickname(c *gin.Context) string {
	nickname, ok := c.Get(cnf.ReqNicknameKey)
	if !ok {
		return ""
	}
	return nickname.(string)
}
