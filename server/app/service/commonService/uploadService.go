package commonService

import (
	"fmt"
	"mime/multipart"
	"time"

	"x_admin/app/schema/commonSchema"
	"x_admin/plugin"
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
// cid 分类id
// AdminId 用户id
func (upSrv uploadService) UploadFile(file *multipart.FileHeader, cid uint, AdminId uint) (res commonSchema.CommonUploadFileResp, e error) {
	var upRes *plugin.UploadFile
	if upRes, e = plugin.StorageDriver.Upload(file); e != nil {
		return
	}
	var startTime = time.Now()

	// 计算文件MD5
	md5, e := util.ToolsUtil.GetFileMD5(file)
	if e != nil {
		return
	}
	var endTime = time.Now()
	var costTime = endTime.UnixNano() - startTime.UnixNano()
	// 毫秒
	costTime = costTime / 1000000

	fmt.Printf("\n文件大小%d , md5时间： %d毫秒\n", file.Size, costTime)

	var addReq commonSchema.CommonAlbumAddReq
	convert_util.Copy(&addReq, upRes)
	addReq.AdminId = AdminId //管理员
	addReq.Cid = cid         // 分类id
	addReq.Hash = md5

	var albumId uint
	if albumId, e = AlbumService.AlbumAdd(addReq); e != nil {
		return
	}
	convert_util.Copy(&res, addReq)
	res.ID = albumId
	res.Path = upRes.Path
	return res, nil
}
