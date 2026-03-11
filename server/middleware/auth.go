package middleware

import (
	"strings"

	"x_admin/app/service/systemService"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) response.RespType {
	token := c.Request.Header.Get("token")
	if token == "" {
		token = c.Request.URL.Query().Get("token")
	}
	if token == "" {
		return response.TokenEmpty
	}

	// Token是否过期
	tokenKey := config.AdminConfig.BackstageTokenKey + token
	existCnt := util.RedisUtil.Exists(tokenKey)
	if existCnt < 0 {
		return response.SystemError
	} else if existCnt == 0 {
		return response.TokenInvalid
	}
	// 用户token获取用户id
	uid := util.RedisUtil.Get(tokenKey)
	if uid == "" {
		return response.TokenInvalid
	}

	// redis管理员信息不存在时缓存
	if !util.RedisUtil.HExists(config.AdminConfig.BackstageManageKey, uid) {
		err := systemService.AdminService.CacheAdminUserByUid(uid) //缓存管理员信息
		if err != nil {
			core.Logger.Errorf("缓存管理员失败: err=[%+v]", err)
			return response.SystemError
		}
	}

	// 校验用户被删除
	var adminUser system_model.SystemAuthAdmin
	err := util.ToolsUtil.JsonToObj(util.RedisUtil.HGet(config.AdminConfig.BackstageManageKey, uid), &adminUser)
	if err != nil {
		core.Logger.Errorf("TokenAuth Unmarshal err: err=[%+v]", err)
		return response.SystemError
	}
	if adminUser.IsDelete == 1 {
		util.RedisUtil.Del(tokenKey)
		util.RedisUtil.HDel(config.AdminConfig.BackstageManageKey, uid)
		return response.TokenInvalid
	}

	// 校验用户被禁用
	if adminUser.IsDisable == 1 {
		return response.LoginDisableError
	}

	// 令牌剩余30分钟自动续签
	if util.RedisUtil.TTL(tokenKey) < 1800 {
		util.RedisUtil.Expire(tokenKey, config.AdminConfig.TokenExpire)
	}

	// 单次请求信息保存
	c.Set(config.AdminConfig.ReqAdminIdKey, uid)
	c.Set(config.AdminConfig.ReqUsernameKey, adminUser.Username)
	c.Set(config.AdminConfig.ReqNicknameKey, adminUser.Nickname)

	return response.Success
}

// 仅检查token有效性，获取用户信息，不判断接口权限
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := Auth(c)
		if resp != response.Success {
			response.FailWithResp(c, resp)
			c.Abort()
			return
		}
		c.Next()
	}
}

// TokenAuth 检查token有效性，获取用户信息，并判断接口权限
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 路由转权限
		ApiAuth := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/", ":")

		// 免登录接口
		if util.ToolsUtil.Contains(config.AdminConfig.NotLoginUri, ApiAuth) {
			c.Next()
			return
		}
		resp := Auth(c)
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
		perms, err := systemService.PermService.GetAdminPerms(adminId)
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
