package common_service

import (
	"strconv"
	"time"

	"x_admin/app/model/common_model"
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
	fileHash, err := s.FindById(id)
	if err != nil || fileHash.ID == "" {
		return ""
	}
	s.cacheFilePath(id, fileHash.FilePath)
	return fileHash.FilePath
}

// FindById 根据主键ID查找文件哈希记录
func (s *fileHashService) FindById(id string) (common_model.CommonFileHash, error) {
	var record common_model.CommonFileHash
	db := core.GetDB()
	err := db.Where("id = ?", id).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return record, nil
		}
		core.Logger.Errorf("FileHashService.FindById err: id=%s, err=%+v", id, err)
		return record, err
	}
	// 读取即更新 LastAccessTime，利用 redis限流不需要频繁更新
	db.Model(&common_model.CommonFileHash{}).
		Where("id = ?", id).
		Update("last_access_time", time.Now())
	return record, nil
}

// FindByMd5 根据 MD5 查找文件哈希记录（用于秒传）
func (s *fileHashService) FindByMd5(fileMd5 string) (common_model.CommonFileHash, error) {
	var record common_model.CommonFileHash
	db := core.GetDB()
	err := db.Where("file_md5 = ?", fileMd5).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 记录不存在，秒传未命中（正常情况）
			return record, nil
		}
		// 真实数据库错误
		core.Logger.Errorf("FileHashService.FindByMd5 err: md5=%s, err=%+v", fileMd5, err)
		return record, err
	}
	return record, nil
}

// cacheFilePath 写入 Redis 缓存（TTL 24h）。失败仅记日志，不影响主流程。
func (s *fileHashService) cacheFilePath(id string, filePath string) {
	if id == "" || filePath == "" {
		return
	}
	util.RedisUtil.Set(fileHashCacheKey+id, filePath, fileHashCacheTTLSec)
}

// CreateDerived 新增派生记录（webp/缩略图等）。pid 指向原图主记录 id，
// quality/scaleWidth/scaleHeight 描述派生参数（0 表示未指定/不缩放）。
// 记录的主键 id 由 BeforeCreate 生成，FilePath/Ext/FileSize 为派生文件自身的信息，
// FileMd5 应填派生文件自身的 MD5（避免与主记录唯一索引冲突）。
// 写入成功后预热 Redis 缓存（GetDerivedPath 使用）。
func (s *fileHashService) CreateDerived(pid, filePath, ext, fileMd5 string, fileSize int64, quality, scaleWidth, scaleHeight int) (string, error) {
	record := common_model.CommonFileHash{
		Pid:         pid,
		Quality:     quality,
		ScaleWidth:  scaleWidth,
		ScaleHeight: scaleHeight,
		FileMd5:     fileMd5,
		FileSize:    fileSize,
		FilePath:    filePath,
		Ext:         ext,
	}
	db := core.GetDB()
	if err := db.Create(&record).Error; err != nil {
		return "", response.CheckErr(err, "创建派生文件记录失败")
	}
	// 预热派生记录缓存，避免首次访问穿透 DB
	s.cacheDerivedPath(pid, quality, scaleWidth, scaleHeight, filePath)
	return record.ID, nil
}

// GetDerivedPath 根据父文件 id 与派生参数查询磁盘存储相对路径（存储 key）。
// 优先读 Redis（key=file_derived:<pid>:<quality>:<w>:<h>，TTL 24h），miss 时查 DB 并回填。
// 不存在返回空串。
func (s *fileHashService) GetDerivedPath(pid string, quality, scaleWidth, scaleHeight int) string {
	if pid == "" {
		return ""
	}
	cacheKey := derivedCacheKey(pid, quality, scaleWidth, scaleHeight)
	if v := util.RedisUtil.Get(cacheKey); v != "" {
		return v
	}
	var record common_model.CommonFileHash
	db := core.GetDB()
	err := db.Where("pid = ? AND quality = ? AND scale_width = ? AND scale_height = ?",
		pid, quality, scaleWidth, scaleHeight).First(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			core.Logger.Errorf("FileHashService.GetDerivedPath err: pid=%s, err=%+v", pid, err)
		}
		return ""
	}
	// 读取即更新 LastAccessTime，利用 redis限流不需要频繁更新
	db.Model(&common_model.CommonFileHash{}).
		Where("id = ?", record.ID).
		Update("last_access_time", time.Now())

	s.cacheDerivedPath(pid, quality, scaleWidth, scaleHeight, record.FilePath)
	return record.FilePath
}

// derivedCacheKey 派生记录缓存键
func derivedCacheKey(pid string, quality, scaleWidth, scaleHeight int) string {
	return "file_derived:" + pid + ":" +
		strconv.Itoa(quality) + ":" +
		strconv.Itoa(scaleWidth) + ":" +
		strconv.Itoa(scaleHeight)
}

// cacheDerivedPath 写入派生记录 Redis 缓存（TTL 24h）
func (s *fileHashService) cacheDerivedPath(pid string, quality, scaleWidth, scaleHeight int, filePath string) {
	if pid == "" || filePath == "" {
		return
	}
	util.RedisUtil.Set(derivedCacheKey(pid, quality, scaleWidth, scaleHeight), filePath, fileHashCacheTTLSec)
}
