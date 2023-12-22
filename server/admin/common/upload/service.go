package upload

import (
	"mime/multipart"
	"x_admin/admin/common/album"
	"x_admin/core/response"
	"x_admin/plugin"
)

type IUploadService interface {
	UploadImage(file *multipart.FileHeader, cid uint, aid uint) (res album.CommonUploadFileResp, e error)
	UploadVideo(file *multipart.FileHeader, cid uint, aid uint) (res album.CommonUploadFileResp, e error)
}

var Service = NewUploadService()

// NewUploadService 初始化
func NewUploadService() *uploadService {
	return &uploadService{}
}

// uploadService 上传服务实现类
type uploadService struct{}

// UploadImage 上传图片
func (upSrv uploadService) UploadImage(file *multipart.FileHeader, cid uint, aid uint) (res album.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "image", 10, cid, aid)
}

// UploadVideo 上传视频
func (upSrv uploadService) UploadVideo(file *multipart.FileHeader, cid uint, aid uint) (res album.CommonUploadFileResp, e error) {
	return upSrv.uploadFile(file, "video", 20, cid, aid)
}

// uploadFile 上传文件
func (upSrv uploadService) uploadFile(file *multipart.FileHeader, folder string, fileType int, cid uint, aid uint) (res album.CommonUploadFileResp, e error) {
	var upRes *plugin.UploadFile
	if upRes, e = plugin.StorageDriver.Upload(file, folder, fileType); e != nil {
		return
	}
	var addReq album.CommonAlbumAddReq
	response.Copy(&addReq, upRes)
	addReq.Aid = aid
	addReq.Cid = cid
	var albumId uint
	if albumId, e = album.Service.AlbumAdd(addReq); e != nil {
		return
	}
	response.Copy(&res, addReq)
	res.ID = albumId
	res.Path = upRes.Path
	return res, nil
}
