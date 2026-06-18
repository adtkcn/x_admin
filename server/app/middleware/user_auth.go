package middleware

import (
	"strings"
	"x_admin/app/service/user_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// UserLoginAuth 用户系统JWT鉴权中间件
// 从 Header "Authorization: Bearer <token>" 获取 access_token
func UserLoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.FailWithResp(c, response.TokenEmpty.SetMessage("缺少Authorization头"))
			c.Abort()
			return
		}

		// 解析 Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.FailWithResp(c, response.TokenEmpty.SetMessage("Authorization格式错误，应为: Bearer <token>"))
			c.Abort()
			return
		}
		tokenStr := parts[1]

		// 解析 access_token
		claims, err := util.JWTUtil.ParseAccessToken(tokenStr)
		if err != nil {
			response.FailWithResp(c, response.TokenInvalid.SetMessage("token已失效: "+err.Error()))
			c.Abort()
			return
		}

		// token_version 校验（踢人下线检测）—— 优先读 Redis，miss 才回源 MySQL
		currentVersion, err := user_service.UserService.GetCachedTokenVersion(claims.UserID)
		if err != nil {
			core.Logger.Errorf("UserLoginAuth 获取token_version失败: userID=%s err=%v", claims.UserID, err)
			response.FailWithResp(c, response.SystemError.SetMessage("鉴权查询失败"))
			c.Abort()
			return
		}
		if claims.TokenVersion != currentVersion {
			response.FailWithResp(c, response.TokenInvalid.SetMessage("token已失效，请重新登录"))
			c.Abort()
			return
		}

		// 将用户信息写入 gin.Context（仅 userID）
		c.Set(config.UserIDKey, claims.UserID)

		// 续签：access_token 剩余<30分钟时，在响应头返回新 token
		if util.JWTUtil.IsAccessExpiringSoon(claims) {
			newAccess, err := util.JWTUtil.GenerateAccessToken(claims.UserID, claims.TokenVersion)
			if err == nil {
				c.Header("X-New-Access-Token", newAccess)
			}
		}

		c.Next()
	}
}
