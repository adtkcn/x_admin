package util

import (
	"sync"
	"x_admin/config"
	"x_admin/core"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
)

var (
	miniProgramApp     *miniProgram.MiniProgram
	officialAccountApp *officialAccount.OfficialAccount
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
					Addrs: []string{config.RedisConfig.Url},
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
					Addrs: []string{config.RedisConfig.Url},
				}),
			})
			if err != nil {
				core.Logger.Errorf("InitWechatClients OfficialAccount err: %v", err)
			} else {
				officialAccountApp = app
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
