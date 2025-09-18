package image

import (
	"log"
	"os"
	"path/filepath"
	"x_admin/util/aj-captcha-go/captcha_config"
	"x_admin/util/aj-captcha-go/util"
)

var backgroundImageArr []string
var clickBackgroundImageArr []string
var templateImageArr []string

func SetUp() {

	backgroundImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultBackgroundImageDirectory
	templateImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultTemplateImageDirectory
	clickBackgroundImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultClickBackgroundImageDirectory

	err1 := filepath.Walk(backgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		backgroundImageArr = append(backgroundImageArr, path)
		return nil
	})

	err2 := filepath.Walk(templateImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		templateImageArr = append(templateImageArr, path)
		return nil
	})

	err3 := filepath.Walk(clickBackgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		clickBackgroundImageArr = append(clickBackgroundImageArr, path)
		return nil
	})

	if err1 != nil {
		log.Printf("初始化resource目录失败，请检查该目录是否存在 err: %v", err1)
	}
	if err2 != nil {
		log.Printf("初始化resource目录失败，请检查该目录是否存在 err: %v", err2)
	}
	if err3 != nil {
		log.Printf("初始化resource目录失败，请检查该目录是否存在 err: %v", err3)
	}

}

func GetBackgroundImage() *util.ImageUtil {
	max := len(backgroundImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(backgroundImageArr[util.RandomInt(0, max)], captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
}

func GetTemplateImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(templateImageArr[util.RandomInt(0, max)], captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
}

func GetClickBackgroundImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(clickBackgroundImageArr[util.RandomInt(0, max)], captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
}
