package config

import "github.com/gin-gonic/gin"

// JWTConfig JWT 配置
var JWTConfig = jwtConfig{
	AccessSecret:     "x_admin_jwt_access_secret_change_me",
	RefreshSecret:    "x_admin_jwt_refresh_secret_change_me",
	AccessExpireSec:  7200,   // access_token 有效期 2小时
	RefreshExpireSec: 604800, // refresh_token 有效期 7天
}

type jwtConfig struct {
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
func (cnf jwtConfig) GetUserID(c *gin.Context) string {
	v, ok := c.Get(UserIDKey)
	if !ok {
		return ""
	}
	return v.(string)
}
