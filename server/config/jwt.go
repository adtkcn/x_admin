package config

import "github.com/gin-gonic/gin"

// JWTConfig web 端用户 JWT 配置（保持不变，web 端继续引用）
var JWTConfig = JwtConfig{
	AccessSecret:     "x_admin_jwt_web_access_secret_change_me",
	RefreshSecret:    "x_admin_jwt_web_refresh_secret_change_me",
	AccessExpireSec:  7200,   // access_token 有效期 2小时
	RefreshExpireSec: 604800, // refresh_token 有效期 7天
}

// AdminJWTConfig 后台管理 JWT 配置（与 web 端分离，独立密钥/过期）
var AdminJWTConfig = JwtConfig{
	AccessSecret:     "x_admin_jwt_admin_access_secret_change_me",
	RefreshSecret:    "x_admin_jwt_admin_refresh_secret_change_me",
	AccessExpireSec:  7200,   // access_token 有效期 2小时
	RefreshExpireSec: 604800, // refresh_token 有效期 7天
}

type JwtConfig struct {
	AccessSecret     string `mapstructure:"AccessSecret"`     // access_token 签名密钥
	RefreshSecret    string `mapstructure:"RefreshSecret"`    // refresh_token 签名密钥
	AccessExpireSec  int    `mapstructure:"AccessExpireSec"`  // access_token 有效期(秒)
	RefreshExpireSec int    `mapstructure:"RefreshExpireSec"` // refresh_token 有效期(秒)
}

// Context key 常量
const (
	UserIDKey = "user_id"
)

// GetUserID 从 gin.Context 获取当前用户ID
func (cnf JwtConfig) GetUserID(c *gin.Context) string {
	v, ok := c.Get(UserIDKey)
	if !ok {
		return ""
	}
	return v.(string)
}
