package plugin

import (
	"sync"
	"x_admin/config"
	"x_admin/core"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/redis/go-redis/v9"
)

// redisAddrs 将配置中的 redis URI（redis://user:pass@host:port/db）
// 转换为 PowerWeChat 缓存所需的 host:port 格式，避免 "too many colons in address"
func redisAddrs() []string {
	opt, err := redis.ParseURL(config.RedisConfig.Url)
	if err != nil {
		core.Logger.Errorf("redis url 解析失败: %v，回退到 127.0.0.1:6379", err)
		return []string{"127.0.0.1:6379"}
	}
	return []string{opt.Addr}
}

var (
	miniProgramApp     *miniProgram.MiniProgram
	officialAccountApp *officialAccount.OfficialAccount
	paymentApp         *payment.Payment
	wechatOnce         sync.Once
)

// InitWechatClients 初始化微信 SDK 客户端（在 main.go 中调用一次）
func InitWechatClients() {
	wechatOnce.Do(func() {
		// 小程序
		if config.WechatConfig.MiniAppID != "" {
			app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
				AppID:     config.WechatConfig.MiniAppID,
				Secret:    config.WechatConfig.MiniSecret,
				HttpDebug: false,
				Log: miniProgram.Log{
					Level:  "error",
					Stdout: false,
				},
				Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
					Addrs: redisAddrs(),
				}),
			})
			if err != nil {
				core.Logger.Errorf("InitWechatClients MiniProgram err: %v", err)
			} else {
				miniProgramApp = app
			}
		}

		// 公众号
		if config.WechatConfig.MpAppID != "" {
			app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
				AppID:     config.WechatConfig.MpAppID,
				Secret:    config.WechatConfig.MpSecret,
				HttpDebug: false,
				Log: officialAccount.Log{
					Level:  "error",
					Stdout: false,
				},
				Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
					Addrs: redisAddrs(),
				}),
			})
			if err != nil {
				core.Logger.Errorf("InitWechatClients OfficialAccount err: %v", err)
			} else {
				officialAccountApp = app
			}
		}

		// 微信支付（v3）
		if config.WechatConfig.Payment.MchID != "" {
			app, err := payment.NewPayment(&payment.UserConfig{
				AppID:       config.WechatConfig.MiniAppID,
				MchID:       config.WechatConfig.Payment.MchID,
				MchApiV3Key: config.WechatConfig.Payment.MchApiV3Key,
				KeyPath:     config.WechatConfig.Payment.KeyPath,
				SerialNo:    config.WechatConfig.Payment.SerialNo,
				NotifyURL:   config.WechatConfig.Payment.NotifyURL,
				Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
					Addrs: redisAddrs(),
				}),
			})
			if err != nil {
				core.Logger.Errorf("InitWechatClients Payment err: %v", err)
			} else {
				paymentApp = app
			}
		}
	})
}

// GetMiniProgramApp 获取小程序客户端
func GetMiniProgramApp() *miniProgram.MiniProgram {
	return miniProgramApp
}

// GetOfficialAccountApp 获取公众号客户端
func GetOfficialAccountApp() *officialAccount.OfficialAccount {
	return officialAccountApp
}

// GetPaymentApp 获取微信支付客户端（未配置时返回 nil）
func GetPaymentApp() *payment.Payment {
	return paymentApp
}
