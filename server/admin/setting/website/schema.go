package website

//SettingWebsiteReq 保存网站信息参数
type SettingWebsiteReq struct {
	Name     string `form:"name"`     // 网站名称
	Logo     string `form:"logo"`     // 网站图标
	Favicon  string `form:"favicon"`  // 网站LOGO
	Backdrop string `form:"backdrop"` // 登录页广告图
	ShopName string `form:"shopName"` // 商城名称
	ShopLogo string `form:"shopLogo"` // 商城Logo
}
