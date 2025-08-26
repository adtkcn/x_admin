package plugin

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"
)

var StorageDriver = storageDriver{}

// UploadFile 文件对象
type UploadFile struct {
	Name string // 文件名称
	// Type int    // 文件类型
	Size int64  // 文件大小
	Ext  string // 文件扩展
	Uri  string // 文件路径
	Path string // 访问地址
}

// storageDriver 存储引擎
type storageDriver struct{}

// Upload 根据引擎类型上传文件
func (sd storageDriver) Upload(file *multipart.FileHeader) (uf *UploadFile, e error) {
	// TODO: engine默认local

	fileExt := sd.getFileExt(file.Filename)
	if fileExt == "" {
		return nil, response.AssertArgumentError.SetMessage("文件类型错误！")
	}

	if e = sd.checkFile(file.Filename, file.Size); e != nil {
		return
	}
	var folder string = fileExt

	saveName := sd.buildSaveName(file)
	engine := "local"
	if engine == "local" {
		if e = sd.localUpload(file, folder, saveName); e != nil {
			return
		}
	} else {
		core.Logger.Errorf("storageDriver.Upload engine err: err=[unsupported engine]")
		return nil, response.Failed.SetMessage(fmt.Sprintf("engine:%s 暂时不支持", engine))
	}

	fileRelPath := path.Join(folder, saveName)
	return &UploadFile{
		Name: file.Filename,
		// Type: int(fileType),
		Size: file.Size,
		Ext:  strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1)),
		Uri:  fileRelPath,
		Path: util.UrlUtil.ToAbsoluteUrl(fileRelPath),
	}, nil
}

// localUpload 本地上传 (临时方法)
func (sd storageDriver) localUpload(file *multipart.FileHeader, folder string, saveName string) (e error) {
	// TODO: 临时方法，后续调整
	// 映射目录
	directory := config.FileConfig.UploadDirectory
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		core.Logger.Errorf("storageDriver.localUpload Open err: err=[%+v]", err)
		return response.Failed.SetMessage("打开文件失败!")
	}
	defer src.Close()
	// 文件信息
	savePath := path.Join(directory, folder, path.Dir(saveName))
	saveFilePath := path.Join(directory, folder, saveName)
	// 创建目录
	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		core.Logger.Errorf(
			"storageDriver.localUpload MkdirAll err: path=[%s], err=[%+v]", savePath, err)
		return response.Failed.SetMessage("创建上传目录失败!")
	}
	// 创建目标文件
	out, err := os.Create(saveFilePath)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Create err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.SetMessage("创建文件失败!")
	}
	defer out.Close()
	// 写入目标文件
	_, err = io.Copy(out, src)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.SetMessage("上传文件失败: " + err.Error())
	}

	return nil
}

// checkFile 生成文件名称
func (sd storageDriver) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060102")
	return path.Join(date, util.ToolsUtil.MakeUuid()+ext)
}

// getFileExt 获取文件扩展名
func (sd storageDriver) getFileExt(fileName string) string {
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	return fileExt
}

// checkFile 文件验证
func (sd storageDriver) checkFile(fileName string, fileSize int64) (e error) {

	fileExt := sd.getFileExt(fileName)

	if util.ToolsUtil.Contains(config.FileConfig.UploadImageExt, fileExt) {
		// 图片文件
		if fileSize > config.FileConfig.UploadImageSize {
			return response.Failed.SetMessage("上传图片不能超出限制: " + strconv.FormatInt(config.FileConfig.UploadImageSize/1024/1024, 10) + "M")
		}
	} else if util.ToolsUtil.Contains(config.FileConfig.UploadVideoExt, fileExt) {
		// 视频文件
		if fileSize > config.FileConfig.UploadVideoSize {
			return response.Failed.SetMessage("上传视频不能超出限制: " + strconv.FormatInt(config.FileConfig.UploadVideoSize/1024/1024, 10) + "M")
		}
	} else if util.ToolsUtil.Contains(config.FileConfig.UploadFileExt, fileExt) {
		// 文件
		if fileSize > config.FileConfig.UploadFileSize {
			return response.Failed.SetMessage("上传文件不能超出限制: " + strconv.FormatInt(config.FileConfig.UploadFileSize/1024/1024, 10) + "M")
		}
	} else {
		core.Logger.Errorf("storageDriver.checkFile fileType err: err=[unsupported fileType]")
		return response.Failed.SetMessage("上传文件类型错误")
	}

	return nil
}
