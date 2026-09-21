package img_util

import (
	"bytes"
	"fmt"
	"image"
	"io"

	_ "image/jpeg" // 注册 jpeg 解码器，供 image.Decode 使用
	_ "image/png"  // 注册 png 解码器，供 image.Decode 使用

	webp "github.com/SeriousBug/webp-go-pure/std"
	"github.com/kovidgoyal/imaging"
)

func EmptyGif() []byte {

	return []byte{
		0x47, 0x49, 0x46, 0x38, 0x39, 0x61, // GIF89a
		0x01, 0x00, 0x01, 0x00, // width=1, height=1
		0x80, 0x00, 0x00, // global color table
		0xff, 0xff, 0xff, // background color (white)
		0x00, 0x00, 0x00, // palette: black
		0x2c,                   // image descriptor
		0x00, 0x00, 0x00, 0x00, // x=0, y=0
		0x01, 0x00, 0x01, 0x00, // width=1, height=1
		0x00,       // no local color table
		0x02,       // LZW min code size
		0x02,       // image data size
		0x4c, 0x01, // image data
		0x3b, // trailer
	}

}

// ConvertToWebp 将 jpg/png 等图片转换为 webp（有损压缩）。
// quality: 有损质量 0-100，越大越清晰、体积越大；越界返回错误。
// 实现即官方示例：解码 jpg/png 得到 image.Image，再 webp.Encode 直接编码，中间无需额外转换。
func ConvertToWebp(src io.Reader, quality int, targetW, targetH int) ([]byte, error) {
	if quality > 100 || quality < 0 {
		return nil, fmt.Errorf("img_util: webp quality %d out of range 0-100", quality)
	}

	// 解码 jpg/png（image/jpeg、image/png 已在本文件 import 注册解码器）。
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("img_util: decode source image: %w", err)
	}
	if targetW > 0 || targetH > 0 {
		img = ScaleImage(img, targetW, targetH)
	}

	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{Quality: quality}); err != nil {
		return nil, fmt.Errorf("img_util: encode webp: %w", err)
	}
	return buf.Bytes(), nil
}

func ScaleImage(img image.Image, targetW, targetH int) image.Image {
	bounds := img.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()

	if targetW > 0 && srcW > targetW || targetH > 0 && srcH > targetH {
		// 限制最大宽高缩放
		if targetW > 0 && targetH > 0 {
			return imaging.Fit(img, targetW, targetH, imaging.Lanczos)

		}
		// 执行缩放，imaging.Resize 会正确处理 targetW 或 targetH 为0的情况
		return imaging.Resize(img, targetW, targetH, imaging.Lanczos)
	} else {
		return img
	}

}
