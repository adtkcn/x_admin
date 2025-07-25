package config

import (
	"image/color"
	"x_admin/util/aj-captcha-go/constant"
)

// WatermarkConfig 水印设置
type WatermarkConfig struct {
	FontSize int        `yaml:"fontSize"`
	Color    color.RGBA `yaml:"color"`
	Text     string     `yaml:"text"`
}

type BlockPuzzleConfig struct {
	// 校验时 容错偏移量
	Offset int `yaml:"offset"`
}

type ClickWordConfig struct {
	FontSize int `yaml:"fontSize"`
	FontNum  int `yaml:"fontNum"`
}

// RedisConfig redis配置选项
type RedisConfig struct {
	//redis单机或者集群访问地址
	DBAddress []string `yaml:"dbAddress"`
	//最大空闲连接数
	DBMaxIdle int `yaml:"dbMaxIdle"`
	//最大连接数
	DBMaxActive int `yaml:"dbMaxActive"`
	//redis表示空闲连接保活时间
	DBIdleTimeout int `yaml:"dbIdleTimeout"`
	//redis用户
	DBUserName string `yaml:"dbUserName"`
	//redis密码
	DBPassWord string `yaml:"dbPassWord"`
	//是否使用redis集群
	EnableCluster bool `yaml:"enableCluster"`
	//单机模式下使用redis的指定库，比如：0，1，2，3等等，默认为0
	DB int `yaml:"db"`
}

type Config struct {
	// 验证码使用的缓存类型
	CacheType      string `yaml:"cacheType"`
	CacheExpireSec int    `yaml:"cacheExpireSec"`

	Watermark   *WatermarkConfig   `yaml:"watermark"`
	ClickWord   *ClickWordConfig   `yaml:"clickWord"`
	BlockPuzzle *BlockPuzzleConfig `yaml:"blockPuzzle"`
}

// 默认验证码配置
func NewConfig() *Config {
	return &Config{
		//可以为redis类型缓存RedisCacheKey，也可以为内存MemCacheKey
		CacheType: constant.MemCacheKey,
		// 水印配置
		Watermark: &WatermarkConfig{
			FontSize: 12,
			Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			Text:     "我的水印",
		},
		// 点击文字配置（参数可从业务系统自定义）
		ClickWord: &ClickWordConfig{
			FontSize: 25,
			FontNum:  4,
		},
		// 滑动模块配置（参数可从业务系统自定义）
		BlockPuzzle:    &BlockPuzzleConfig{Offset: 10},
		CacheExpireSec: 2 * 60, // 缓存有效时间
	}
}
