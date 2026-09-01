package util

import (
	"log"
	"sync"
	"unicode"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"x_admin/util/file_util"
)

type FontUtil struct {
	Src string
}

// fontCache 按字体路径缓存已解析的字体对象，避免每次生成验证码都重复读盘与解析 TTF
var fontCache sync.Map // src(string) -> *truetype.Font

func NewFontUtil(src string) *FontUtil {
	return &FontUtil{Src: src}
}

// GetFont 获取字体对象（按 Src 缓存，解析一次后复用）
func (f *FontUtil) GetFont() *truetype.Font {
	if cached, ok := fontCache.Load(f.Src); ok {
		return cached.(*truetype.Font)
	}

	fontSourceBytes, err := file_util.ReadFile(f.Src)
	if err != nil {
		log.Println("读取字体失败:", err)
		return nil
	}
	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		log.Println("解析字体失败:", err)
		return nil
	}

	fontCache.Store(f.Src, trueTypeFont)
	return trueTypeFont
}

func GetEnOrChLength(text string) int {
	enCount, zhCount := 0, 0

	for _, t := range text {
		if unicode.Is(unicode.Han, t) {
			zhCount++
		} else {
			enCount++
		}
	}

	chOffset := 12*zhCount + 5
	enOffset := enCount * 8

	return chOffset + enOffset
}
