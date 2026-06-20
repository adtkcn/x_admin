package common_service

import (
	"os"
	"path"
	"time"
	"x_admin/app/model/common_model"
	"x_admin/config"
	"x_admin/core"

	"gorm.io/gorm"
)

var FileRefService = newFileRefService()

type fileRefService struct {
	db *gorm.DB
}

func newFileRefService() *fileRefService {
	return &fileRefService{
		db: core.GetDB(),
	}
}

// AddRef 新增文件业务关联
// fileHashID: x_common_file_hash.id
// bizType: 业务类型，如 "user_avatar"、"article_cover"、"album"
// bizID: 业务实体 ID
// func (s *fileRefService) AddRef(fileHashID, bizType, bizID string) error {
// 	db := core.GetDB()
// 	// 幂等：已存在相同关联则跳过
// 	var count int64
// 	err := db.Model(&common_model.CommonFileRef{}).
// 		Where("file_hash_id = ? AND biz_type = ? AND biz_id = ?", fileHashID, bizType, bizID).
// 		Count(&count).Error
// 	if err != nil {
// 		core.Logger.Errorf("FileRefService.AddRef count err: %+v", err)
// 		return err
// 	}
// 	if count > 0 {
// 		return nil
// 	}

// 	ref := common_model.CommonFileRef{
// 		FileHashID: fileHashID,
// 		BizType:    bizType,
// 		BizID:      bizID,
// 	}
// 	if err := db.Create(&ref).Error; err != nil {
// 		core.Logger.Errorf("FileRefService.AddRef create err: %+v", err)
// 		return err
// 	}
// 	return nil
// }

// // RemoveRef 移除文件业务关联（业务删除时调用）
// func (s *fileRefService) RemoveRef(fileHashID, bizType, bizID string) error {
// 	db := core.GetDB()
// 	err := db.Where("file_hash_id = ? AND biz_type = ? AND biz_id = ?", fileHashID, bizType, bizID).
// 		Delete(&common_model.CommonFileRef{}).Error
// 	if err != nil {
// 		core.Logger.Errorf("FileRefService.RemoveRef err: %+v", err)
// 		return err
// 	}
// 	return nil
// }

// // RemoveByBiz 按业务类型+业务ID移除所有关联（业务实体整体删除时调用）
// func (s *fileRefService) RemoveByBiz(bizType, bizID string) error {
// 	db := core.GetDB()
// 	err := db.Where("biz_type = ? AND biz_id = ?", bizType, bizID).
// 		Delete(&common_model.CommonFileRef{}).Error
// 	if err != nil {
// 		core.Logger.Errorf("FileRefService.RemoveByBiz err: %+v", err)
// 		return err
// 	}
// 	return nil
// }

// HasRef 检查文件是否有业务引用
func (s *fileRefService) HasRef(fileHashID string) (bool, error) {

	var count int64
	err := s.db.Model(&common_model.CommonFileRef{}).
		Where("file_hash_id = ?", fileHashID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// SaveFileRefs 在同一个事务中：删除旧关联 + 批量插入新关联
// tx: 外部传入的事务对象（由调用方 db.Transaction 创建）
// bizType: 业务类型，如 "article_cover"
// bizID: 业务实体 ID
// newFileHashIDs: 需要关联的新文件哈希 ID 列表
func (s *fileRefService) SaveFileRefs(tx *gorm.DB, bizType, bizID string, newFileHashIDs []string) error {
	// 1. 删除旧关联
	if err := tx.Where("biz_type = ? AND biz_id = ?", bizType, bizID).
		Delete(&common_model.CommonFileRef{}).Error; err != nil {
		core.Logger.Errorf("FileRefService.SaveFileRefs delete err: bizType=%s bizID=%s err=%+v", bizType, bizID, err)
		return err
	}

	// 2. 批量插入新关联
	if len(newFileHashIDs) == 0 {
		return nil
	}
	refs := make([]common_model.CommonFileRef, 0, len(newFileHashIDs))
	for _, hashID := range newFileHashIDs {
		if hashID == "" {
			continue
		}
		refs = append(refs, common_model.CommonFileRef{
			FileHashID: hashID,
			BizType:    bizType,
			BizID:      bizID,
		})
	}
	if len(refs) == 0 {
		return nil
	}
	if err := tx.Create(&refs).Error; err != nil {
		core.Logger.Errorf("FileRefService.SaveFileRefs batch create err: bizType=%s bizID=%s err=%+v", bizType, bizID, err)
		return err
	}
	return nil
}

// CleanOrphanFiles 清理上传超过 expireDays 天且无业务引用的文件
// 1. 查询 x_common_file_hash 中创建时间超过 expireDays 天的记录
// 2. 检查 x_common_file_ref 中是否存在引用
// 3. 无引用则删除物理文件 + 数据库记录
func (s *fileRefService) CleanOrphanFiles(expireDays int) {
	if expireDays <= 0 {
		expireDays = 7
	}

	db := s.db
	cutoff := time.Now().AddDate(0, 0, -expireDays)

	// 分批查询，每批 100 条，避免一次加载太多
	batchSize := 100
	totalCleaned := 0

	for {
		var files []common_model.CommonFileHash
		err := db.Where("create_time < ?", cutoff).Limit(batchSize).Find(&files).Error
		if err != nil {
			core.Logger.Errorf("fileRefService: 查询过期文件失败: %+v", err)
			return
		}
		if len(files) == 0 {
			break
		}

		for _, file := range files {
			// 检查是否有业务引用
			hasRef, err := FileRefService.HasRef(file.ID)
			if err != nil {
				core.Logger.Errorf("fileRefService: 检查引用失败 file_id=%s: %+v", file.ID, err)
				continue
			}
			if hasRef {
				// 有引用，跳过
				continue
			}

			// 无引用，删除物理文件
			filePath := path.Join(config.FileConfig.UploadDirectory, file.FilePath)
			if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
				core.Logger.Warnf("fileRefService: 删除物理文件失败 %s: %+v", filePath, err)
				// 物理文件删除失败也继续删除数据库记录（文件可能已被手动清理）
			}

			// 删除数据库记录
			if err := db.Where("id = ?", file.ID).Delete(&common_model.CommonFileHash{}).Error; err != nil {
				core.Logger.Errorf("fileRefService: 删除数据库记录失败 file_id=%s: %+v", file.ID, err)
				continue
			}

			totalCleaned++
			core.Logger.Infof("fileRefService: 清理无引用文件 id=%s path=%s", file.ID, file.FilePath)
		}

		// 如果本批不足 batchSize 条，说明已无更多数据
		if len(files) < batchSize {
			break
		}
	}

	if totalCleaned > 0 {
		core.Logger.Infof("fileRefService: 本次共清理 %d 个无引用文件", totalCleaned)
	}
}
