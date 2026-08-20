package common_service

import (
	"errors"
	"os"
	"path"
	"time"

	"x_admin/app/model/common_model"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

var FileHashService = newFileHashService()

type fileHashService struct{}

func newFileHashService() *fileHashService {
	return &fileHashService{}
}

// FindByMd5 根据 MD5 查找文件哈希记录（用于秒传）
func (s *fileHashService) FindByMd5(fileMd5 string) (*common_model.CommonFileHash, error) {
	var record common_model.CommonFileHash
	db := core.GetDB()
	err := db.Where("file_md5 = ?", fileMd5).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 记录不存在，秒传未命中（正常情况）
			return nil, nil
		}
		// 真实数据库错误
		core.Logger.Errorf("FileHashService.FindByMd5 err: md5=%s, err=%+v", fileMd5, err)
		return nil, err
	}
	return &record, nil
}

// Create 创建文件哈希记录（秒传：记录已上传文件的 hash 与路径映射）
func (s *fileHashService) Create(fileMd5 string, fileSize int64, filePath string, ext string) error {
	_, err := s.CreateOrGet(fileMd5, fileSize, filePath, ext)
	return err
}

// UpdateWebp 异步转 webp 完成后回写文件哈希记录：仅更新路径、扩展名与大小，
// 文件 MD5 保持不变（原图未变，webp 为派生文件）。
func (s *fileHashService) UpdateWebp(id string, filePath string, ext string, fileSize int64) error {
	db := core.GetDB()
	err := db.Model(&common_model.CommonFileHash{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"file_path":  filePath,
			"ext":        ext,
			"file_size":  fileSize,
		}).Error
	if err != nil {
		core.Logger.Errorf("FileHashService.UpdateWebp err: id=%s, err=%+v", id, err)
	}
	return err
}

// CreateOrGet 按 MD5 查重：存在则返回已有记录，否则新建并返回记录（含 ID）
func (s *fileHashService) CreateOrGet(fileMd5 string, fileSize int64, filePath string, ext string) (hash common_model.CommonFileHash, e error) {
	db := core.GetDB()
	err := db.Where("file_md5 = ?", fileMd5).First(&hash).Error
	if err == nil {
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		core.Logger.Errorf("FileHashService.CreateOrGet Find err: md5=%s, err=%+v", fileMd5, err)
		e = response.Failed.SetMessage("查询文件哈希失败:" + err.Error())
		return
	}
	hash = common_model.CommonFileHash{
		ID:       util.ToolsUtil.MakeUuidV7(),
		FileMd5:  fileMd5,
		FileSize: fileSize,
		FilePath: filePath,
		Ext:      ext,
	}
	err = db.Create(&hash).Error
	if e = response.CheckErr(err, "创建文件哈希记录失败"); e != nil {
		return
	}
	return
}

// CleanOrphanFiles 清理超过 expireDays 天未被访问、且未被相册引用的文件
// 文件访问时间由 last_access_time 标记冷热；未访问时回落到 create_time 判断冷热
func (s *fileHashService) CleanOrphanFiles(expireDays int) {
	if expireDays <= 0 {
		expireDays = 365
	}
	db := core.GetDB()
	cutoff := time.Now().AddDate(0, 0, -expireDays)

	batchSize := 100
	totalCleaned := 0

	for {
		var files []common_model.CommonFileHash
		err := db.Where("COALESCE(last_access_time, create_time) < ?", cutoff).Limit(batchSize).Find(&files).Error
		if err != nil {
			core.Logger.Errorf("fileHashService: 查询冷文件失败: %+v", err)
			return
		}
		if len(files) == 0 {
			break
		}
		for _, file := range files {
			filePath := path.Join(config.FileConfig.UploadDirectory, file.FilePath)
			if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
				core.Logger.Warnf("fileHashService: 删除物理文件失败 %s: %+v", filePath, err)
			}
			if err := db.Where("id = ?", file.ID).Delete(&common_model.CommonFileHash{}).Error; err != nil {
				core.Logger.Errorf("fileHashService: 删除数据库记录失败 file_id=%s: %+v", file.ID, err)
				continue
			}
			totalCleaned++
			core.Logger.Infof("fileHashService: 清理冷文件 id=%s path=%s", file.ID, file.FilePath)
		}
		if len(files) < batchSize {
			break
		}
	}
	if totalCleaned > 0 {
		core.Logger.Infof("fileHashService: 本次共清理 %d 个冷文件", totalCleaned)
	}
}
