package user_schema

// WechatMiniLoginReq 微信小程序登录请求
type WechatMiniLoginReq struct {
	Code string `json:"code" binding:"required" label:"小程序code"`
}

// WechatMpLoginReq 微信公众号登录请求
type WechatMpLoginReq struct {
	Code string `json:"code" binding:"required" label:"公众号OAuth code"`
}

// WechatBindReq 微信绑定请求（小程序/公众号通用）
type WechatBindReq struct {
	Code string `json:"code" binding:"required" label:"微信code"`
}

// WechatUnbindReq 微信解绑请求
type WechatUnbindReq struct {
	IdentityType string `json:"identityType" binding:"required,oneof=wechat_mini wechat_mp" label:"认证类型"`
}
