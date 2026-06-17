package plugin

import (
	"os"
	"path/filepath"
	"time"

	"x_admin/config"
	"x_admin/core"
)

// CleanChunkTmpDir 清理过期临时分片目录
// 扫描 ChunkTmpDir，删除超过 ChunkExpireHour 小时未修改的临时分片目录
func CleanChunkTmpDir() {
	expireHour := config.FileConfig.ChunkExpireHour
	if expireHour <= 0 {
		expireHour = 24 // 默认 24 小时
	}

	cleanupChunkTmpDir(expireHour)
}

// cleanupChunkTmpDir 清理过期临时分片目录
func cleanupChunkTmpDir(expireHour int) {
	tmpDir := config.FileConfig.ChunkTmpDir

	// 检查目录是否存在
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		core.Logger.Debugf("临时分片目录不存在: %s", tmpDir)
		return
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		core.Logger.Errorf("读取临时分片目录失败: %v", err)
		return
	}

	expire := time.Now().Add(-time.Duration(expireHour) * time.Hour)
	cleaned := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dirPath := filepath.Join(tmpDir, entry.Name())
		info, err := entry.Info()
		if err != nil {
			continue
		}
		// 如果目录超过 expire 小时未修改，删除
		if info.ModTime().Before(expire) {
			if err := os.RemoveAll(dirPath); err == nil {
				cleaned++
				core.Logger.Infof("清理过期分片目录: %s", dirPath)
			} else {
				core.Logger.Errorf("清理分片目录失败 %s: %v", dirPath, err)
			}
		}
	}

	if cleaned > 0 {
		core.Logger.Infof("分片临时目录清理完成，共清理 %d 个目录", cleaned)
	}
}
