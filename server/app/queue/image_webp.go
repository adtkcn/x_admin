package queue

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json/v2"
	"fmt"
	"io"

	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/plugin/storage"
	"x_admin/util"
	"x_admin/util/img_util"
)

// // ProcessImageWebp 消费转 webp 任务：
// 1) 读取原图字节；2) img_util.ConvertToWebp 转 webp（jpg/png 有损压缩）；
// 3) 另存为同名 .webp 文件（保留原图）；4) 回写 x_common_file_hash 的路径/扩展/大小。
func ProcessImageWebp(ctx context.Context, body []byte) error {
	var payload queue_schema.ImageWebpPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		core.Logger.Errorf("解析 image_webp 任务失败: %v", err)
		return nil
	}
	if payload.FileHashId == "" {
		core.Logger.Warnf("image_webp 任务字段缺失: %+v", payload)
		return nil
	}
	fileHash, err := common_service.FileHashService.FindById(payload.FileHashId)
	if err != nil {
		core.Logger.Errorf("查询主图记录失败 file_hash_id=%s: %v", payload.FileHashId, err)
		return nil
	}

	engine := storage.GetStorageEngine()

	// 1) 读取原图全部字节（用于后续与原图大小比较）
	src, err := engine.GetObject(fileHash.FilePath)
	if err != nil {
		core.Logger.Errorf("读取原图失败 path=%s: %v", fileHash.FilePath, err)
		return nil
	}
	defer src.Close()
	srcData, err := io.ReadAll(src)
	if err != nil {
		core.Logger.Errorf("读取原图内容失败 path=%s: %v", fileHash.FilePath, err)
		return nil
	}

	// 2) 转 webp
	quality := payload.Quality
	if quality == 0 {
		quality = config.FileConfig.WebpQuality
	}
	webpData, err := img_util.ConvertToWebp(bytes.NewReader(srcData), quality, payload.ScaleWidth, payload.ScaleHeight)
	if err != nil {
		core.Logger.Errorf("转 webp 失败 path=%s: %v", fileHash.FilePath, err)
		return nil
	}

	// // 3) 体积未优化则忽略（保留原图，不写回、不更新记录）
	// if len(webpData) >= len(srcData) {
	// 	core.Logger.Infof("转 webp 体积未减小，忽略 path=%s: orig=%d webp=%d",
	// 		fileHash.FilePath, len(srcData), len(webpData))
	// 	return nil
	// }
	// 体积未优化：不再忽略，而是新增一条"占位派生记录"，
	// 该记录复用主图的路径/格式/大小（即请求派生参数时仍返回原图），
	// 避免后续重复转码与反复回退主文件。FileMd5 用主图MD5+派生参数组合，
	// 避开 x_common_file_hash.FileMd5 唯一索引（不可复用主图MD5本身）。
	if len(webpData) >= len(srcData) {
		core.Logger.Infof("转 webp 体积未减小，新增占位派生记录 path=%s: 原图大小=%d webp大小=%d", fileHash.FilePath, len(srcData), len(webpData))

		if _, err := common_service.FileHashService.CreateDerived(
			payload.FileHashId,
			fileHash.FilePath, // 复用主图路径
			fileHash.Ext,      // 复用主图格式
			fileHash.FileMd5,
			fileHash.FileSize, // 复用主图大小
			quality,
			payload.ScaleWidth,
			payload.ScaleHeight,
		); err != nil {
			core.Logger.Errorf("新增占位派生记录失败 file_hash_id=%s: %v", payload.FileHashId, err)
			return nil
		}
		return nil
	} else {
		var joinPath = fmt.Sprintf("q%d_w%d_h%d", quality, payload.ScaleWidth, payload.ScaleHeight)
		// 3) 另存为同名 .webp（保留原图），新 key 替换扩展名
		neFilePath := util.UrlUtil.ReplaceExt(fileHash.FilePath, joinPath+".webp")
		if _, err := engine.PutObject(neFilePath, bytes.NewReader(webpData), int64(len(webpData))); err != nil {
			core.Logger.Errorf("写回 webp 失败 key=%s: %v", neFilePath, err)
			return nil
		}

		// 4) 新增 webp 派生记录（pid 指向原图主记录，主记录始终保留原图路径）
		webpMd5 := md5.Sum(webpData)
		webpFileMd5 := hex.EncodeToString(webpMd5[:])
		if _, err := common_service.FileHashService.CreateDerived(
			payload.FileHashId,
			neFilePath,
			"webp",
			webpFileMd5,
			int64(len(webpData)),
			quality,
			payload.ScaleWidth,
			payload.ScaleHeight,
		); err != nil {
			core.Logger.Errorf("新增 webp 派生记录失败 file_hash_id=%s: %v", payload.FileHashId, err)
			return nil
		}
		return nil
	}

}
