package commonService

import (
	"image/color"
	"x_admin/core"
	"x_admin/schema/commonSchema"
	config2 "x_admin/util/aj-captcha-go/config"
	constant "x_admin/util/aj-captcha-go/const"
	"x_admin/util/aj-captcha-go/service"
)

// var config = config2.NewConfig()// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
// var config = config2.WatermarkConfig//自定义配置
// 水印配置（参数可从业务系统自定义）
var watermarkConfig = &config2.WatermarkConfig{
	FontSize: 12,
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Text:     "x_admin",
}

// 点击文字配置（参数可从业务系统自定义）
var clickWordConfig = &config2.ClickWordConfig{
	FontSize: 25,
	FontNum:  5,
}

// 滑动模块配置（参数可从业务系统自定义）
var blockPuzzleConfig = &config2.BlockPuzzleConfig{Offset: 10}

// 行为校验配置模块（具体参数可从业务系统配置文件自定义）
var captcha_config = config2.BuildConfig(constant.RedisCacheKey, constant.DefaultResourceRoot, watermarkConfig,
	clickWordConfig, blockPuzzleConfig, 2*60)

// 服务工厂，主要用户注册 获取 缓存和验证服务
var factory = service.NewCaptchaServiceFactory(captcha_config)

func init() {
	// 这里默认是注册了 内存缓存，但是不足以应对生产环境，希望自行注册缓存驱动 实现缓存接口即可替换（CacheType就是注册进去的 key）
	// factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(200000)) // 这里20指的是缓存阈值

	// //注册自定义配置redis数据库
	factory.RegisterCache(constant.RedisCacheKey, service.NewConfigRedisCacheService(core.Redis))

	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))
}

// 登录等场景验证，并删除
func CaptchaVerify(params commonSchema.ClientParams) error {
	ser := factory.GetService(params.CaptchaType)
	// 登录验证并删除
	err := ser.Verification(params.Token, params.PointJson)
	return err
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
