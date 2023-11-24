package copyright

//SettingCopyrightItemReq 保存备案信息参数
type SettingCopyrightItemReq struct {
	Name string `form:"name" json:"name"`  // 名称
	Link string `form:"link"  json:"link"` // 链接
}
