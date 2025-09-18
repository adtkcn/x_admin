package captcha_config

import (
	"image/color"
)

// WatermarkConfig 水印设置
type WatermarkConfig struct {
	FontSize int
	Color    color.RGBA
	Text     string // 水印文字
}

// 滑块配置
type BlockPuzzleConfig struct {
	Offset int // 校验时 容错偏移量
}

// 点击文字配置
type ClickWordConfig struct {
	FontSize   int // 点击文字字体大小
	FontNum    int // 点击文字数量
	AllFontNum int // 点击文字显示数量
	XOffset    int // 点击文字X轴偏移量
	YOffset    int // 点击文字Y轴偏移量
}

type Config struct {
	CacheType      string // 验证码使用的缓存类型
	CacheExpireSec int    // 缓存有效时间

	Watermark   *WatermarkConfig   // 水印配置
	ClickWord   *ClickWordConfig   // 点击文字配置
	BlockPuzzle *BlockPuzzleConfig // 滑动模块配置
}

// 默认验证码配置
func NewMemCacheConfig() *Config {
	return &Config{
		//可以为redis类型缓存RedisCacheKey，也可以为内存MemCacheKey
		CacheType:      MemCacheKey,
		CacheExpireSec: 2 * 60, // 缓存有效时间

		// 水印配置
		Watermark: &WatermarkConfig{
			FontSize: 12,
			Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			Text:     "x_admin",
		},
		// 点击文字配置（参数可从业务系统自定义）
		ClickWord: &ClickWordConfig{
			FontSize:   25,
			FontNum:    4,
			AllFontNum: 10,
			XOffset:    10,
			YOffset:    10,
		},
		// 滑动模块配置（参数可从业务系统自定义）
		BlockPuzzle: &BlockPuzzleConfig{Offset: 10},
	}
}
