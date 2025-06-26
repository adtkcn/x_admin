package commonService

import (
	"image/color"
	"x_admin/core"
	"x_admin/schema/commonSchema"
	config2 "x_admin/util/aj-captcha-go/config"
	constant "x_admin/util/aj-captcha-go/const"
	"x_admin/util/aj-captcha-go/service"
)

// var captcha_config = config2.NewConfig()// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
var captcha_config = config2.Config{
	CacheType: constant.RedisCacheKey,
	Watermark: &config2.WatermarkConfig{
		FontSize: 12,
		Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Text:     "admin",
	},
	ClickWord: &config2.ClickWordConfig{
		FontSize: 25,
		FontNum:  4,
	},
	BlockPuzzle:    &config2.BlockPuzzleConfig{Offset: 10},
	CacheExpireSec: 2 * 60, // 缓存有效时间
}

// 服务工厂，主要用户注册 获取 缓存和验证服务
var factory = service.NewCaptchaServiceFactory(&captcha_config)

func init() {
	// 这里默认是注册了 内存缓存，但是不足以应对生产环境，希望自行注册缓存驱动 实现缓存接口即可替换（CacheType就是注册进去的 key）
	// factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(200000)) // 这里20指的是缓存阈值

	// //注册自定义配置redis数据库
	factory.RegisterCache(constant.RedisCacheKey, service.NewConfigRedisCacheService(core.Redis))

	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))
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
