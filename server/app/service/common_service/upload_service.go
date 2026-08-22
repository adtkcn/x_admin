package common_service

import (
	"mime/multipart"

	"x_admin/app/schema/common_schema"
	"x_admin/app/schema/queue_schema"
	"x_admin/core"
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

// UploadFile 上传文件
func (upSrv uploadService) UploadFile(file *multipart.FileHeader) (res common_schema.CommonFileHashResp, e error) {
	md5, e := util.ToolsUtil.GetFileMD5(file)
	if e != nil {
		return
	}

	// 秒传命中：已存在相同 MD5 的文件，直接返回已有记录
	if hit, _ := FileHashService.FindByMd5(md5); hit != nil {
		res = common_schema.CommonFileHashResp{}
		convert_util.Copy(&res, hit)
		return res, nil
	}
	uploadFile, e := storage.StorageDriver.Upload(file)
	if e != nil {
		return
	}

	// id 由 Create 内部生成（UUID），FilePath 存磁盘存储 key
	id, e := FileHashService.Create(md5, file.Size, uploadFile.Path, uploadFile.Ext)
	if e != nil {
		return
	}

	res = common_schema.CommonFileHashResp{
		ID:       id,
		FileMd5:  md5,
		FileSize: file.Size,
		FilePath: uploadFile.Path,
		Ext:      uploadFile.Ext,
	}
	return res, nil
}

// MaybeConvertWebp 上传成功后（直传或 S3 分片注册）按需异步转 webp。
// 仅对 jpg/jpeg/png 位图生效（gif 动画编码暂不支持，保持原样）；小于 10KB 压缩收益低，跳过。
// filePath 为真实存储 key，fileHashId 为 x_common_file_hash 主键，ext 为原图扩展名（不含点）。
// 入队失败仅记日志，不影响主流程。payload 字段名对齐 image_webp worker 的 ImageWebpPayload。
func (upSrv uploadService) MaybeConvertWebp(fileHashId, filePath, ext string, fileSize int64) {
	if !util.ToolsUtil.Contains([]string{"jpg", "jpeg", "png"}, ext) || fileSize < 10*1024 {
		return
	}
	payload := queue_schema.ImageWebpPayload{
		FilePath:   filePath,
		Ext:        ext,
		FileHashId: fileHashId,
	}
	if err := core.Queue.Enqueue(queue_schema.QueueImageWebp, payload); err != nil {
		core.Logger.Errorf("MaybeConvertWebp 入队失败 file_hash_id=%s: %v", fileHashId, err)
	}
}
