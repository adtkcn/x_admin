package captcha

import (
	"fmt"
	"image/color"
	"x_admin/config"

	config2 "github.com/TestsLing/aj-captcha-go/config"
	constant "github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

type CaptchaGetParams struct {
	CaptchaType string `json:"captchaType"`
}

// 客户端参数 看自身业务构建即可
type ClientParams struct {
	Token       string `json:"token"`
	PointJson   string `json:"pointJson"`
	CaptchaType string `json:"captchaType"`
}

// **********************默认配置***************************************************
// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
// var config = config2.NewConfig()

// *********************自定义配置**************************************************
// 水印配置（参数可从业务系统自定义）
var watermarkConfig = &config2.WatermarkConfig{
	FontSize: 12,
	Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Text:     "x_admin",
}

// 点击文字配置（参数可从业务系统自定义）
var clickWordConfig = &config2.ClickWordConfig{
	FontSize: 25,
	FontNum:  4,
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
	opt, _ := redis.ParseURL(config.Config.RedisUrl)
	fmt.Printf("%#v", opt)
	factory.RegisterCache(constant.RedisCacheKey, service.NewConfigRedisCacheService([]string{opt.Addr}, opt.Username, opt.Password, false, 0))

	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))
}
func CaptchaRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/common/captcha")
	rg.POST("/get", func(c *gin.Context) {
		var captchaGet CaptchaGetParams
		if err := c.ShouldBind(&captchaGet); err != nil {
			// 返回错误信息
			c.JSON(200, errorRes(err))
			return
		}

		// 根据参数类型获取不同服务即可
		data, _ := factory.GetService(captchaGet.CaptchaType).Get()
		//输出json结果给调用方
		c.JSON(200, successRes(data))
	})
	rg.POST("/check", func(c *gin.Context) {
		var params ClientParams
		if err := c.ShouldBind(&params); err != nil {
			// 返回错误信息
			c.JSON(200, errorRes(err))
			return
		}
		ser := factory.GetService(params.CaptchaType)
		err := ser.Check(params.Token, params.PointJson)

		if err != nil {
			c.JSON(200, errorRes(err))
			return
		}
		//输出json结果给调用方
		c.JSON(200, successRes(nil))
	})

}

// 登录等验证并删除
func Verify(params ClientParams) error {
	ser := factory.GetService(params.CaptchaType)
	// 登录验证并删除
	err := ser.Verification(params.Token, params.PointJson)
	return err
}
func successRes(data interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = false
	ret["repCode"] = "0000"
	ret["repData"] = data
	ret["repMsg"] = nil
	ret["successRes"] = true

	return ret
}
func errorRes(err error) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = true
	ret["repCode"] = "0001"
	ret["repData"] = nil
	ret["repMsg"] = err.Error()
	ret["successRes"] = false
	return ret
}
