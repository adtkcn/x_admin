package common_service

import (
	"mime/multipart"

	"x_admin/app/schema/common_schema"
	"x_admin/plugin/storage"
	"x_admin/util"
	"x_admin/util/convert_util"
)

var UploadService = NewUploadService()

// NewUploadService 初始化
func NewUploadService() *uploadService {
	return &uploadService{}
}

// uploadService 上传服务实现类
type uploadService struct{}

// UploadFile 上传
// 仅负责落盘 + 登记文件哈希（x_common_file_hash）。
func (upSrv uploadService) UploadFile(file *multipart.FileHeader) (res common_schema.CommonFileHashResp, e error) {
	var upRes *storage.UploadFile
	if upRes, e = storage.StorageDriver.Upload(file); e != nil {
		return
	}

	// 计算文件MD5
	md5, e := util.ToolsUtil.GetFileMD5(file)
	if e != nil {
		return
	}

	// 登记文件哈希（秒传去重），返回哈希记录
	hash, e := FileHashService.CreateOrGet(md5, upRes.Size, upRes.Uri, upRes.Ext)
	if e != nil {
		return
	}
	res = common_schema.CommonFileHashResp{}
	convert_util.Copy(&res, hash)
	return res, nil
}
