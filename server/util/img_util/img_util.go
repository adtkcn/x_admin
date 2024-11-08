package img_util

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
)

func EmptyGif() []byte {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))

	// 设置像素颜色为黑色
	img.Set(0, 0, color.Black)

	// 创建GIF动画
	g := &gif.GIF{
		Image: []*image.Paletted{imageToPaletted(img)},
		Delay: []int{0}, // 延迟时间，单位是10毫秒
	}
	var buffer bytes.Buffer

	// 编码GIF到缓冲区
	gif.EncodeAll(&buffer, g)
	return buffer.Bytes()
}

// 将image.Image转换为*image.Paletted
func imageToPaletted(img image.Image) *image.Paletted {
	b := img.Bounds()
	pm := image.NewPaletted(b, color.Palette{color.Black})
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pm.Set(x, y, img.At(x, y))
		}
	}
	return pm
}
