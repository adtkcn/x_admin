package util

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"x_admin/plugin/aj-captcha-go/model/vo"
	"x_admin/util/file_util"

	"github.com/golang/freetype"
)

type ImageUtil struct {
	Src       string
	SrcImage  image.Image
	RgbaImage *image.RGBA
	FontPath  string
	Width     int
	Height    int
}

// NewImageUtil src 为图片相对路径，fontPath 为字体相对路径
func NewImageUtil(src string, fontPath string) *ImageUtil {
	srcImage := OpenPngImage(src)
	if srcImage == nil {
		log.Printf("加载图片失败: %s", src)
		return nil
	}

	return &ImageUtil{
		Src:       src,
		SrcImage:  srcImage,
		RgbaImage: ImageToRGBA(srcImage),
		Width:     srcImage.Bounds().Dx(),
		Height:    srcImage.Bounds().Dy(),
		FontPath:  fontPath,
	}
}

// IsOpacity 该像素是否透明
func (i *ImageUtil) IsOpacity(x, y int) bool {
	A := i.RgbaImage.RGBAAt(x, y).A

	return float32(A) <= 125
}

// SetText 为图片设置右下角水印文字
func (i *ImageUtil) SetText(text string, fontsize int, color color.RGBA) {
	x := float64(i.Width) - float64(GetEnOrChLength(text))
	y := float64(i.Height) - 5

	font := NewFontUtil(i.FontPath)

	fc := freetype.NewContext()
	fc.SetFont(font.GetFont())
	fc.SetFontSize(float64(fontsize))
	fc.SetClip(i.RgbaImage.Bounds())
	fc.SetDst(i.RgbaImage)
	fc.SetSrc(image.NewUniform(color))
	pt := freetype.Pt(int(x), int(y))
	if _, err := fc.DrawString(text, pt); err != nil {
		log.Println("构造水印失败:", err)
	}
}

// SetArtText 为图片设置文字（随机颜色）
func (i *ImageUtil) SetArtText(text string, fontsize int, point vo.PointVO) error {
	return i.SetArtTextWithColor(text, fontsize, point, color.RGBA{
		R: uint8(RandomInt(1, 200)),
		G: uint8(RandomInt(1, 200)),
		B: uint8(RandomInt(1, 200)),
		A: 255,
	})
}

// SetArtTextWithColor 为图片设置文字（指定颜色）
func (i *ImageUtil) SetArtTextWithColor(text string, fontsize int, point vo.PointVO, c color.RGBA) error {
	font := NewFontUtil(i.FontPath)

	fc := freetype.NewContext()
	fc.SetFont(font.GetFont())
	fc.SetFontSize(float64(fontsize))
	fc.SetClip(i.RgbaImage.Bounds())
	fc.SetDst(i.RgbaImage)
	fc.SetSrc(image.NewUniform(c))
	pt := freetype.Pt(point.X, point.Y+fontsize)
	_, err := fc.DrawString(text, pt)
	if err != nil {
		log.Printf("构造水印失败 err: %v", err)
		return err
	}
	return nil
}

// SetPixel 为像素设置颜色
func (i *ImageUtil) SetPixel(rgba color.RGBA, x, y int) {
	i.RgbaImage.SetRGBA(x, y, rgba)
}

// Base64 将图片编码为 base64 字符串
func (i *ImageUtil) Base64() (string, error) {
	// 开辟一个新的空buff
	var buf bytes.Buffer
	// img写入到buff
	err := png.Encode(&buf, i.RgbaImage)

	if err != nil {
		log.Printf("img写入buf失败 err: %v", err)
		return "", err
	}
	// buff转成base64
	dist := base64.StdEncoding.EncodeToString(buf.Bytes())
	return dist, nil
}

// VagueImage 用周围 8 邻域像素的平均值模糊当前像素
func (i *ImageUtil) VagueImage(x int, y int) {
	points := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	var red, green, blue, alpha uint32
	var count uint32
	for _, p := range points {
		px, py := x+p[0], y+p[1]
		if px < 0 || px >= i.Width || py < 0 || py >= i.Height {
			continue
		}
		r, g, b, a := i.RgbaImage.RGBAAt(px, py).RGBA()
		red += r >> 8
		green += g >> 8
		blue += b >> 8
		alpha += a >> 8
		count++
	}
	if count == 0 {
		return
	}
	i.RgbaImage.SetRGBA(x, y, color.RGBA{
		R: uint8(red / count),
		G: uint8(green / count),
		B: uint8(blue / count),
		A: uint8(alpha / count),
	})
}

// OpenPngImage 打开png图片
func OpenPngImage(src string) image.Image {
	ff, err := file_util.Open(src)
	if err != nil {
		log.Printf("打开 %s 图片失败: %v", src, err)
		return nil
	}
	defer ff.Close()

	img, err := png.Decode(ff)
	if err != nil {
		log.Printf("png %s decode 失败: %v", src, err)
		return nil
	}

	return img
}

// ImageToRGBA 图片转rgba
func ImageToRGBA(img image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := img.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return dst
}
