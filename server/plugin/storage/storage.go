package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
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
	fileExt := util.UrlUtil.GetFileExt(file.Filename)

	if fileExt == "" {
		return nil, response.AssertArgumentError.SetMessage("文件类型错误！")
	}

	if e = sd.checkFile(file.Filename, file.Size); e != nil {
		return
	}
	// var folder string = fileExt

	savePash := util.UrlUtil.BuildFileSavePath(file.Filename)
	engine := "local"
	if engine == "local" {
		if e = sd.localSaveFile(file, savePash); e != nil {
			return
		}
	} else {
		core.Logger.Errorf("storageDriver.Upload engine err: err=[unsupported engine]")
		return nil, response.Failed.SetMessage(fmt.Sprintf("engine:%s 暂时不支持", engine))
	}

	return &UploadFile{
		Name: file.Filename,
		// Type: int(fileType),
		Size: file.Size,
		Ext:  fileExt,
		Uri:  util.UrlUtil.ToAbsoluteUrl(savePash),
		Path: savePash,
	}, nil
}

// localSaveFile 本地存储
func (sd storageDriver) localSaveFile(file *multipart.FileHeader, saveName string) (e error) {
	// TODO: 临时方法，后续调整
	// 映射目录
	directory := config.FileConfig.UploadDirectory
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		core.Logger.Errorf("storageDriver.localSaveFile Open err: err=[%+v]", err)
		return response.Failed.SetMessage("打开文件失败!")
	}
	defer src.Close()
	// 文件信息
	saveDir := path.Join(directory, path.Dir(saveName))
	saveFilePath := path.Join(directory, saveName)
	// 创建目录
	err = os.MkdirAll(saveDir, 0755)
	if err != nil && !os.IsExist(err) {
		core.Logger.Errorf(
			"storageDriver.localSaveFile MkdirAll err: path=[%s], err=[%+v]", saveDir, err)
		return response.Failed.SetMessage("创建上传目录失败!")
	}
	// 创建目标文件
	out, err := os.Create(saveFilePath)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localSaveFile Create err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.SetMessage("创建文件失败!")
	}
	defer out.Close()
	// 写入目标文件
	_, err = io.Copy(out, src)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localSaveFile Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.SetMessage("上传文件失败: " + err.Error())
	}

	return nil
}

// checkFile 文件验证
func (sd storageDriver) checkFile(fileName string, fileSize int64) (e error) {

	fileExt := util.UrlUtil.GetFileExt(fileName)

	if util.ToolsUtil.Contains(config.FileConfig.UploadImageExt, fileExt) {
		// 图片文件
		if fileSize > config.FileConfig.UploadImageSize {
			return response.Failed.SetMessage("上传图片不能超出限制: " + strconv.FormatInt(config.FileConfig.UploadImageSize/1024/1024, 10) + "M")
		}
	} else if util.ToolsUtil.Contains(config.FileConfig.UploadVideoExt, fileExt) {
		// 视频文件
		if fileSize > config.FileConfig.UploadVideoSize {
			return response.Failed.SetMessage("上传音视频不能超出限制: " + strconv.FormatInt(config.FileConfig.UploadVideoSize/1024/1024, 10) + "M")
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
