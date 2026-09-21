package web_route

import (
	"x_admin/app/controller/web_ctl"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoute 注册用户路由
// 路由前缀: /api/web/user
func UserRoute(rg *gin.RouterGroup) {
	userRg := rg.Group("/user")
	handle := web_ctl.UserController{}
	authHandle := web_ctl.AuthController{}

	// 免登录接口
	userRg.POST("/sendCode", handle.SendEmailCode)                // 发送邮箱验证码
	userRg.POST("/sendSmsCode", handle.SendSmsCode)               // 发送短信验证码
	userRg.POST("/register", handle.Register)                     // 邮箱注册
	userRg.POST("/login", handle.Login)                           // 邮箱+密码登录
	userRg.POST("/phoneLogin", handle.PhoneLogin)                 // 手机号+密码登录
	userRg.POST("/phoneCodeLogin", handle.PhoneCodeLogin)         // 手机号+短信验证码登录
	userRg.POST("/refresh", handle.RefreshToken)                  // 刷新token
	userRg.POST("/resetPassword", handle.ResetPassword)           // 邮箱重置密码
	userRg.POST("/resetPhonePassword", handle.ResetPhonePassword) // 手机号重置密码

	// 微信登录（免登录）
	userRg.POST("/wechatMiniLogin", authHandle.WechatMiniLogin) // 小程序登录
	userRg.POST("/wechatMpLogin", authHandle.WechatMpLogin)     // 公众号登录

	// 需要登录的接口
	auth := userRg.Group("/", middleware.UserLoginAuth())
	{
		auth.GET("/info", handle.GetUserInfo)         // 获取用户信息
		auth.POST("/info", handle.UpdateUserInfo)     // 更新用户信息
		auth.POST("/kickOffline", handle.KickOffline) // 踢人下线

		// 手机绑定/解绑
		auth.POST("/bindPhone", authHandle.BindPhone)     // 绑定手机号（需短信验证码）
		auth.POST("/unbindPhone", authHandle.UnbindPhone) // 解绑手机号（需邮箱验证码）

		// 微信绑定/解绑
		auth.POST("/bindWechatMini", authHandle.BindWechatMini) // 绑定小程序
		auth.POST("/bindWechatMp", authHandle.BindWechatMp)     // 绑定公众号
		auth.POST("/unbindWechat", authHandle.UnbindWechat)     // 解绑微信

		// 第三方绑定列表
		auth.GET("/authList", authHandle.GetUserAuthList) // 获取绑定列表
		auth.PUT("/changePassword", handle.ChangePassword)
	}
}
func init() {
	webRouteHandlers = append(webRouteHandlers, UserRoute)
}
