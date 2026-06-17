package common_service

import (
	"x_admin/app/model/common_model"
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
	db := core.GetDB()
	// 已存在则跳过
	exist, err := s.FindByMd5(fileMd5)
	if err != nil {
		return err
	}
	if exist != nil {
		return nil
	}
	record := common_model.CommonFileHash{
		ID:       util.ToolsUtil.MakeUuidV7(),
		FileMd5:  fileMd5,
		FileSize: fileSize,
		FilePath: filePath,
		Ext:      ext,
	}
	err = db.Create(&record).Error
	if err != nil {
		core.Logger.Errorf("FileHashService.Create err: %+v", err)
		return response.CheckErr(err, "创建文件哈希记录失败")
	}
	return nil
}
