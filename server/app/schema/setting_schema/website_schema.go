package setting_schema

// SettingWebsiteReq 保存网站信息参数
type SettingWebsiteReq struct {
	Name     string `json:"name" form:"name"`           // 网站名称
	Logo     string `json:"logo" form:"logo"`           // 网站图标
	Favicon  string `json:"favicon" form:"favicon"`     // 网站LOGO
	Backdrop string `json:"backdrop" form:"backdrop"`   // 登录页广告图
	ShopName string `json:"shop_name" form:"shop_name"` // 商城名称
	ShopLogo string `json:"shop_logo" form:"shop_logo"` // 商城Logo
}
