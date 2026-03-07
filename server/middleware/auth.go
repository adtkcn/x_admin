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
	// Token是否为空
	token := c.Request.Header.Get("token")
	if token == "" { // 从url获取token
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

	// 用户信息缓存
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
		util.RedisUtil.Expire(tokenKey, 7200)
	}

	// 单次请求信息保存
	c.Set(config.AdminConfig.ReqAdminIdKey, uid)
	c.Set(config.AdminConfig.ReqRoleIdKey, adminUser.RoleId)
	c.Set(config.AdminConfig.ReqUsernameKey, adminUser.Username)
	c.Set(config.AdminConfig.ReqNicknameKey, adminUser.Nickname)

	// 校验角色的权限，redis没有就重新查询
	roleId := adminUser.RoleId
	if roleId != "" {
		if !util.RedisUtil.HExists(config.AdminConfig.BackstageRolesKey, roleId) {

			err = systemService.PermService.CacheRoleMenusByRoleId(roleId)
			if err != nil {
				core.Logger.Errorf("Auth 缓存角色失败 err: [%+v]", err)
				return response.SystemError
			}
		}
	}

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
		// 超管权限
		if config.AdminConfig.GetAdminId(c) == config.AdminConfig.SuperAdminId {
			c.Next()
			return
		}
		// 验证是否有权限操作
		menus := util.RedisUtil.HGet(config.AdminConfig.BackstageRolesKey, config.AdminConfig.GetRoleId(c))
		if !(menus != "" && util.ToolsUtil.Contains(strings.Split(menus, ","), ApiAuth)) {
			response.FailWithResp(c, response.NoPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
