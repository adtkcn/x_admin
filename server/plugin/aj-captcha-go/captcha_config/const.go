package captcha_config

const (
	// CodeKeyPrefix 缓存key前缀
	CodeKeyPrefix = "captcha:%s"
	// BlockPuzzleCaptcha 滑动验证码服务标识
	BlockPuzzleCaptcha = "blockPuzzle"

	// ClickWordCaptcha 点击验证码服务标识
	ClickWordCaptcha = "clickWord"

	// MemCacheKey 内存缓存标识
	MemCacheKey = "mem"
	// RedisCacheKey redis缓存标识
	RedisCacheKey = "redis"

	// DefaultResourceRoot 默认根目录
	DefaultResourceRoot = "./public/"
	// DefaultText 默认水印显示文字
	DefaultText = "我的水印"
)

const (
	// DefaultTemplateImageDirectory 滑动模板图文件目录地址
	DefaultTemplateImageDirectory = "resources/captcha/jigsaw/slidingBlock"
	// DefaultBackgroundImageDirectory 背景图片目录地址
	DefaultBackgroundImageDirectory = "resources/captcha/jigsaw/original"
	// DefaultClickBackgroundImageDirectory 点击背景图默认地址
	DefaultClickBackgroundImageDirectory = "resources/captcha/pic-click"
	// DefaultFont 字体文件地址
	DefaultFont = "resources/captcha/fonts/WenQuanZhengHei.ttf"
)
