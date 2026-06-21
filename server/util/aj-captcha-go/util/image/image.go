package image

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"x_admin/util/aj-captcha-go/captcha_config"
	"x_admin/util/aj-captcha-go/util"
)

var (
	backgroundImageArr      []string
	clickBackgroundImageArr []string
	templateImageArr        []string
	setupOnce               sync.Once
)

// SetUp 初始化图片资源目录（通过 sync.Once 保证只执行一次）
func SetUp() {
	setupOnce.Do(func() {
		backgroundImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultBackgroundImageDirectory
		templateImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultTemplateImageDirectory
		clickBackgroundImageRoot := captcha_config.DefaultResourceRoot + captcha_config.DefaultClickBackgroundImageDirectory

		backgroundImageArr = walkDir(backgroundImageRoot)
		templateImageArr = walkDir(templateImageRoot)
		clickBackgroundImageArr = walkDir(clickBackgroundImageRoot)

		if len(backgroundImageArr) == 0 {
			log.Printf("警告: 背景图片目录为空: %s", backgroundImageRoot)
		}
		if len(templateImageArr) == 0 {
			log.Printf("警告: 模板图片目录为空: %s", templateImageRoot)
		}
		if len(clickBackgroundImageArr) == 0 {
			log.Printf("警告: 点击背景图目录为空: %s", clickBackgroundImageRoot)
		}
	})
}

// walkDir 遍历目录收集所有文件路径
func walkDir(root string) []string {
	var result []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("初始化resource目录失败: %s, err: %v", root, err)
	}
	return result
}

// GetBackgroundImage 随机获取一张背景图
func GetBackgroundImage() (*util.ImageUtil, error) {
	path := randomPick(backgroundImageArr)
	if path == "" {
		return nil, fmt.Errorf("背景图片目录为空")
	}
	img := util.NewImageUtil(path, captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
	if img == nil {
		return nil, fmt.Errorf("加载背景图片失败: %s", path)
	}
	return img, nil
}

// GetTemplateImage 随机获取一张模板图
func GetTemplateImage() (*util.ImageUtil, error) {
	path := randomPick(templateImageArr)
	if path == "" {
		return nil, fmt.Errorf("模板图片目录为空")
	}
	img := util.NewImageUtil(path, captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
	if img == nil {
		return nil, fmt.Errorf("加载模板图片失败: %s", path)
	}
	return img, nil
}

// GetClickBackgroundImage 随机获取一张点击背景图
func GetClickBackgroundImage() (*util.ImageUtil, error) {
	path := randomPick(clickBackgroundImageArr)
	if path == "" {
		return nil, fmt.Errorf("点击背景图目录为空")
	}
	img := util.NewImageUtil(path, captcha_config.DefaultResourceRoot+captcha_config.DefaultFont)
	if img == nil {
		return nil, fmt.Errorf("加载点击背景图失败: %s", path)
	}
	return img, nil
}

// randomPick 从数组中随机选取一个元素
func randomPick(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	return arr[util.RandomInt(0, len(arr))]
}
