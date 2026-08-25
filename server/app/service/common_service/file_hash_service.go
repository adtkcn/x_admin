package common_service

import (
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

// fileHashCacheKey Redis 缓存前缀（RedisUtil.Set 内部会再拼 RedisConfig.RedisPrefix）
const fileHashCacheKey = "file_hash:"

// fileHashCacheTTLSec 缓存有效期 24h
// 写路径（Create / UpdateWebp）完成后主动覆盖，保证一致性。
const fileHashCacheTTLSec = 24 * 3600

var FileHashService = newFileHashService()

type fileHashService struct{}

func newFileHashService() *fileHashService {
	return &fileHashService{}
}

// Create 创建文件哈希记录，主键 id 内部生成（UUID）。
// filePath 为磁盘存储 key（含扩展名），ext 为扩展名。返回生成的 id。
func (s *fileHashService) Create(fileMd5 string, fileSize int64, filePath, ext string) (string, error) {

	record := common_model.CommonFileHash{
		// ID:       id,
		FileMd5:  fileMd5,
		FileSize: fileSize,
		FilePath: filePath,
		Ext:      ext,
	}
	db := core.GetDB()
	err := db.Create(&record).Error
	if err != nil {
		return "", response.CheckErr(err, "创建文件哈希记录失败")
	}
	// 预热缓存：避免首次访问穿透到 DB
	s.cacheFilePath(record.ID, filePath)
	return record.ID, nil
}

// GetFilePath 根据文件哈希ID获取磁盘存储相对路径（存储 key）。
// 优先读 Redis（key=file_hash:<id>，TTL 24h），miss 时查 DB 并回填。
// 用于文件流路由按 id 定位物理文件。记录不存在返回空串。
func (s *fileHashService) GetFilePath(id string) string {
	if id == "" {
		return ""
	}
	cacheKey := fileHashCacheKey + id
	if v := util.RedisUtil.Get(cacheKey); v != "" {
		return v
	}
	hash, err := s.FindById(id)
	if err != nil || hash == nil {
		return ""
	}
	s.cacheFilePath(id, hash.FilePath)
	return hash.FilePath
}

// FindById 根据主键ID查找文件哈希记录
func (s *fileHashService) FindById(id string) (*common_model.CommonFileHash, error) {
	var record common_model.CommonFileHash
	db := core.GetDB()
	err := db.Where("id = ?", id).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		core.Logger.Errorf("FileHashService.FindById err: id=%s, err=%+v", id, err)
		return nil, err
	}
	// 读取即更新 LastAccessTime，利用 redis限流不需要频繁更新
	db.Model(&common_model.CommonFileHash{}).
		Where("id = ?", id).
		Update("last_access_time", time.Now())
	return &record, nil
}

// cacheFilePath 写入 Redis 缓存（TTL 24h）。失败仅记日志，不影响主流程。
func (s *fileHashService) cacheFilePath(id string, filePath string) {
	if id == "" || filePath == "" {
		return
	}
	util.RedisUtil.Set(fileHashCacheKey+id, filePath, fileHashCacheTTLSec)
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

// UpdateWebp 异步转 webp 完成后回写文件哈希记录：仅更新路径、扩展名与大小，
// 文件 MD5 保持不变（原图未变，webp 为派生文件）。
// 成功后刷新 Redis 缓存，避免后续访问命中旧路径（同一 id 但 ext 已变）。
func (s *fileHashService) UpdateWebp(id string, filePath string, ext string, fileSize int64) error {
	db := core.GetDB()
	err := db.Model(&common_model.CommonFileHash{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"file_path": filePath,
			"ext":       ext,
			"file_size": fileSize,
		}).Error
	if err != nil {
		core.Logger.Errorf("FileHashService.UpdateWebp err: id=%s, err=%+v", id, err)
		return err
	}
	// 强制覆盖缓存（旧值可能仍存在，TTL 未到）
	util.RedisUtil.Set(fileHashCacheKey+id, filePath, fileHashCacheTTLSec)
	return nil
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
			// 清理 Redis 缓存，避免后续访问命中已删除的物理文件
			util.RedisUtil.Del(fileHashCacheKey + file.ID)
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
