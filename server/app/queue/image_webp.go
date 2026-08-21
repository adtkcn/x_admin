package queue

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"io"
	"path"
	"strings"

	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/plugin/storage"
	"x_admin/util/img_util"
)

// // ProcessImageWebp 消费转 webp 任务：
// 1) 读取原图字节；2) img_util.ConvertToWebp 转 webp（jpg/png 有损压缩）；
// 3) 另存为同名 .webp 文件（保留原图）；4) 回写 x_common_file_hash 的路径/扩展/大小。
func ProcessImageWebp(ctx context.Context, body []byte) error {
	var p queue_schema.ImageWebpPayload
	if err := json.Unmarshal(body, &p); err != nil {
		core.Logger.Errorf("解析 image_webp 任务失败: %v", err)
		return nil
	}
	if p.FilePath == "" || p.FileHashId == "" {
		core.Logger.Warnf("image_webp 任务字段缺失: %+v", p)
		return nil
	}

	engine := storage.GetStorageEngine()

	// 1) 读取原图全部字节（用于后续与原图大小比较）
	src, err := engine.GetObject(p.FilePath)
	if err != nil {
		core.Logger.Errorf("读取原图失败 path=%s: %v", p.FilePath, err)
		return nil
	}
	defer src.Close()
	srcData, err := io.ReadAll(src)
	if err != nil {
		core.Logger.Errorf("读取原图内容失败 path=%s: %v", p.FilePath, err)
		return nil
	}

	// 2) 转 webp
	quality := config.FileConfig.WebpQuality
	if quality == 0 {
		quality = 80
	}
	webpData, err := img_util.ConvertToWebp(bytes.NewReader(srcData), quality)
	if err != nil {
		core.Logger.Errorf("转 webp 失败 path=%s: %v", p.FilePath, err)
		return nil
	}

	// 3) 体积未优化则忽略（保留原图，不写回、不更新记录）
	if len(webpData) >= len(srcData) {
		core.Logger.Infof("转 webp 体积未减小，忽略 path=%s: orig=%d webp=%d",
			p.FilePath, len(srcData), len(webpData))
		return nil
	}

	// 3) 另存为同名 .webp（保留原图），新 key 替换扩展名
	webpKey := replaceExt(p.FilePath, "webp")
	if _, err := engine.PutObject(webpKey, bytes.NewReader(webpData), int64(len(webpData))); err != nil {
		core.Logger.Errorf("写回 webp 失败 key=%s: %v", webpKey, err)
		return nil
	}

	// 4) 回写文件哈希记录
	if err := common_service.FileHashService.UpdateWebp(p.FileHashId, webpKey, "webp", int64(len(webpData))); err != nil {
		core.Logger.Errorf("回写 webp 记录失败 file_hash_id=%s: %v", p.FileHashId, err)
	}
	return nil
}

// replaceExt 将 path 的扩展名替换为 newExt（保留目录与文件名主体）。
func replaceExt(filePath, newExt string) string {
	ext := path.Ext(filePath)
	if ext == "" {
		return filePath + "." + newExt
	}
	return strings.TrimSuffix(filePath, ext) + "." + newExt
}
