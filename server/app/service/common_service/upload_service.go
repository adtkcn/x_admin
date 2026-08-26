package common_service

import (
	"mime/multipart"
	"strconv"

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
	if hit, _ := FileHashService.FindByMd5(md5); hit.ID != "" {
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

// 仅对 jpg/jpeg/png 位图生效（gif 动画编码暂不支持，保持原样）
// filePath 为真实存储 key，
// fileHashId 为 x_common_file_hash 主键，
func (upSrv uploadService) ConvertImage(fileHashId, fileName string, quality, ScaleWidth, ScaleHeight int) {
	ext := util.UrlUtil.GetFileExt(fileName)
	if !util.ToolsUtil.Contains([]string{"jpg", "jpeg", "png", "webp"}, ext) {
		return
	}
	// 去重：同一派生参数在转码进行中（标记未过期）只入队一次，避免重复转码任务
	dedupKey := "converting:" + fileHashId + ":" +
		strconv.Itoa(quality) + ":" + strconv.Itoa(ScaleWidth) + ":" + strconv.Itoa(ScaleHeight)
	if !util.RedisUtil.SetNX(dedupKey, 1, 600) {
		return // 已有相同转码任务在进行，跳过重复入队
	}
	payload := queue_schema.ImageWebpPayload{
		// FilePath:   filePath,
		FileHashId: fileHashId,
		// Ext:          ext,
		Quality:     quality,
		ScaleWidth:  ScaleWidth,
		ScaleHeight: ScaleHeight,
	}
	if err := core.Queue.Enqueue(queue_schema.QueueImageWebp, payload); err != nil {
		core.Logger.Errorf("ConvertImage 入队失败 file_hash_id=%s: %v", fileHashId, err)
	}
}
