package middleware

import (
	"strings"

	"x_admin/app/model/system_model"
	"x_admin/app/service/system_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context) response.RespType {
	token := c.Request.Header.Get("token")
	if token == "" {
		token = c.Request.URL.Query().Get("token")
	}
	if token == "" {
		return response.TokenEmpty
	}

	// 解析 JWT access_token（内嵌 user_id + token_version），使用 admin 独立密钥
	claims, err := util.AdminJWTUtil.ParseAccessToken(token)
	if err != nil {
		return response.TokenInvalid
	}

	// 一次性读取管理员缓存（含 token_version），缓存优先，miss 回源 MySQL
	var adminUser system_model.SystemAuthAdmin
	var tokenVersion int64
	userStr := util.RedisUtil.Get(config.AdminConfig.BackstageAdminKey + ":" + claims.UserID)
	if userStr == "" {
		user, err2 := system_service.AdminService.CacheAdminById(claims.UserID)
		if err2 != nil {
			core.Logger.Errorf("缓存管理员失败: err=[%+v]", err2)
			return response.SystemError
		}
		adminUser = user
		tokenVersion = user.TokenVersion
	} else {
		cached, err := util.ToolsUtil.JsonToObj[system_service.AdminCache](userStr)
		if err != nil {
			core.Logger.Errorf("auth Unmarshal err: err=[%+v]", err)
			return response.SystemError
		}
		adminUser = cached.SystemAuthAdmin
		tokenVersion = cached.TokenVersion
	}

	// token_version 校验（踢人下线检测）
	if claims.TokenVersion != tokenVersion {
		return response.TokenInvalid
	}

	// 校验用户被禁用
	if adminUser.IsDisable == 1 {
		return response.LoginDisableError
	}

	// access_token 剩余<30分钟自动续签，在响应头返回新 token
	if util.AdminJWTUtil.IsAccessExpiringSoon(claims) {
		if newAccess, genErr := util.AdminJWTUtil.GenerateAccessToken(claims.UserID, claims.TokenVersion); genErr == nil {
			c.Header("X-New-Access-Token", newAccess)
		}
	}

	// 单次请求信息保存
	c.Set(config.AdminConfig.ReqAdminIdKey, claims.UserID)
	c.Set(config.AdminConfig.ReqEmailKey, adminUser.Email)
	c.Set(config.AdminConfig.ReqNicknameKey, adminUser.Nickname)

	return response.Success
}

// 仅检查token有效性，获取用户信息，不判断接口权限
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := auth(c)
		if resp != response.Success {
			response.FailWithResp(c, resp)
			c.Abort()
			return
		}
		c.Next()
	}
}

// PermAuth 检查token有效性，获取用户信息、判断接口权限
func PermAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 路由转权限
		ApiAuth := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", ":")

		// 免登录接口
		if util.ToolsUtil.Contains(config.AdminConfig.NotLoginUri, ApiAuth) {
			c.Next()
			return
		}
		resp := auth(c)
		if resp != response.Success {
			response.FailWithResp(c, resp)
			c.Abort()
			return
		}
		// 免权限验证接口
		if util.ToolsUtil.Contains(config.AdminConfig.NotAuthUri, ApiAuth) {
			c.Next()
			return
		}
		adminId := config.AdminConfig.GetAdminId(c)
		if adminId == config.AdminConfig.SuperAdminId {
			c.Next()
			return
		}
		// 获取用户权限
		perms, err := system_service.PermService.GetAdminPerms(adminId)
		if err != nil {
			core.Logger.Errorf("获取用户权限失败: err=[%+v]", err)
			response.FailWithResp(c, response.SystemError)
			c.Abort()
			return
		}
		if !(len(perms) > 0 && util.ToolsUtil.Contains(perms, ApiAuth)) {
			response.FailWithResp(c, response.NoPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
