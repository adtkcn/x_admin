package common_schema

type CaptchaGetParams struct {
	CaptchaType string `json:"captchaType"`
}

// 客户端参数 看自身业务构建即可
type ClientParams struct {
	Token       string `json:"token"`
	PointJson   string `json:"pointJson"`
	CaptchaType string `json:"captchaType"`
}
