package commonService

import (
	"image/color"
	"x_admin/app/schema/commonSchema"
	"x_admin/core"
	"x_admin/util/aj-captcha-go/captcha_config"
	"x_admin/util/aj-captcha-go/captcha_service"
)

// var captcha_config = config.NewMemCacheConfig()// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
var captchaConfig = captcha_config.Config{
	CacheType: captcha_config.RedisCacheKey,
	Watermark: &captcha_config.WatermarkConfig{
		FontSize: 12,
		Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Text:     "",
	},
	ClickWord: &captcha_config.ClickWordConfig{
		FontSize:   24,
		FontNum:    3,
		AllFontNum: 7,
		XOffset:    8,
		YOffset:    8,
	},
	BlockPuzzle:    &captcha_config.BlockPuzzleConfig{Offset: 8},
	CacheExpireSec: 2 * 60, // 缓存有效时间
}

// 服务工厂，主要用户注册 获取 缓存和验证服务
var factory = captcha_service.NewCaptchaServiceFactory(&captchaConfig)

func init() {
	// 这里默认是注册了 内存缓存，但是不足以应对生产环境
	// factory.RegisterCache(captcha_config.MemCacheKey, captcha_service.NewMemCacheService(200000)) // 这里20指的是缓存阈值

	// //注册自定义配置redis数据库
	factory.RegisterCache(captcha_config.RedisCacheKey, captcha_service.NewConfigRedisCacheService(core.Redis))

	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(captcha_config.ClickWordCaptcha, captcha_service.NewClickWordCaptchaService(factory))
	factory.RegisterService(captcha_config.BlockPuzzleCaptcha, captcha_service.NewBlockPuzzleCaptchaService(factory))
}

func CaptchaGet(captchaType string) (interface{}, error) {
	// 根据参数类型获取不同服务即可
	data, err := factory.GetService(captchaType).Get()

	return data, err
}

// 检查是否正确
func CaptchaCheck(params commonSchema.ClientParams) error {
	ser := factory.GetService(params.CaptchaType)
	// 登录验证并删除
	err := ser.Check(params.Token, params.PointJson)
	return err
}

// 登录等场景验证，并删除
func CaptchaVerify(params commonSchema.ClientParams) error {
	ser := factory.GetService(params.CaptchaType)
	// 登录验证并删除
	err := ser.Verification(params.Token, params.PointJson)
	return err
}
