package user_route

import (
	"x_admin/app/controller/user_ctl"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoute 注册用户路由
// 路由前缀: /api/user
func UserRoute(rg *gin.RouterGroup) {
	handle := user_ctl.UserController{}
	authHandle := user_ctl.AuthController{}

	// 免登录接口
	rg.POST("/sendCode", handle.SendEmailCode)                // 发送邮箱验证码
	rg.POST("/sendSmsCode", handle.SendSmsCode)               // 发送短信验证码
	rg.POST("/register", handle.Register)                     // 邮箱注册
	rg.POST("/login", handle.Login)                           // 邮箱+密码登录
	rg.POST("/phoneLogin", handle.PhoneLogin)                 // 手机号+密码登录
	rg.POST("/phoneCodeLogin", handle.PhoneCodeLogin)         // 手机号+短信验证码登录
	rg.POST("/refresh", handle.RefreshToken)                  // 刷新token
	rg.POST("/resetPassword", handle.ResetPassword)           // 邮箱重置密码
	rg.POST("/resetPhonePassword", handle.ResetPhonePassword) // 手机号重置密码

	// 需要登录的接口
	auth := rg.Group("/", middleware.UserLoginAuth())
	{
		auth.GET("/info", handle.GetUserInfo)         // 获取用户信息
		auth.POST("/info", handle.UpdateUserInfo)     // 更新用户信息
		auth.POST("/kickOffline", handle.KickOffline) // 踢人下线

		// 手机绑定/解绑
		auth.POST("/bindPhone", authHandle.BindPhone)     // 绑定手机号（需短信验证码）
		auth.POST("/unbindPhone", authHandle.UnbindPhone) // 解绑手机号（需邮箱验证码）

		// 第三方绑定列表（微信/QQ等，手机号不在其中）
		auth.GET("/authList", authHandle.GetUserAuthList) // 获取绑定列表
	}
}
