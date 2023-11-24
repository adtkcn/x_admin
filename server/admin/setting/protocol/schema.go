package protocol

//SettingProtocolItem 政策通用参数
type SettingProtocolItem struct {
	Name    string `form:"name" json:"name"`        // 名称
	Content string `form:"content"  json:"content"` // 内容
}

//SettingProtocolReq 保存政策信息参数
type SettingProtocolReq struct {
	Service SettingProtocolItem `form:"service" json:"service"`  // 服务协议
	Privacy SettingProtocolItem `form:"privacy"  json:"privacy"` // 隐私协议
}
