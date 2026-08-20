package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"x_admin/config"
)

// 本地存储引擎，实现 StorageEngine 接口
type localStorageEngine struct {
	mu sync.Mutex // 防止同一文件并发合并
}

func newLocalStorageEngine() StorageEngine {
	return &localStorageEngine{}
}

// PutObject 上传完整对象
func (e *localStorageEngine) PutObject(key string, data io.Reader, size int64) (string, error) {
	absPath := filepath.Join(config.FileConfig.UploadDirectory, key)
	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}
	out, err := os.Create(absPath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()
	if _, err = io.Copy(out, data); err != nil {
		return "", fmt.Errorf("写入文件失败: %w", err)
	}
	return key, nil
}

// ObjectExists 检查对象是否存在
func (e *localStorageEngine) ObjectExists(key string) (bool, error) {
	absPath := filepath.Join(config.FileConfig.UploadDirectory, key)
	_, err := os.Stat(absPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetObjectURL 获取对象访问 URL
func (e *localStorageEngine) GetObjectURL(key string) (string, error) {
	// 构建相对 URL 路径，如 /api/uploads/images/20240101/xxx.png
	publicURL := filepath.ToSlash(filepath.Join(config.FileConfig.UploadPrefix, key))
	return publicURL, nil
}

// GetObject 读取本地对象内容
func (e *localStorageEngine) GetObject(key string) (io.ReadCloser, error) {
	absPath := filepath.Join(config.FileConfig.UploadDirectory, key)
	f, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	return f, nil
}

// ---- 分片上传相关 ----

// InitMultipartUpload 初始化分片上传（本地：创建临时目录，返回临时目录路径作为 uploadId）
func (e *localStorageEngine) InitMultipartUpload(key string) (string, error) {
	// uploadId 用临时分片目录路径表示
	uploadId := filepath.Join(config.FileConfig.ChunkTmpDir, strings.TrimSuffix(key, filepath.Ext(key)))
	if err := os.MkdirAll(uploadId, 0755); err != nil {
		return "", fmt.Errorf("创建分片临时目录失败: %w", err)
	}
	return uploadId, nil
}

// UploadPart 上传单个分片
func (e *localStorageEngine) UploadPart(key string, uploadId string, partNumber int, data io.Reader, size int64) (string, error) {
	partPath := filepath.Join(uploadId, strconv.Itoa(partNumber))
	partData, err := io.ReadAll(data)
	if err != nil {
		return "", fmt.Errorf("读取分片数据失败: %w", err)
	}
	if err = os.WriteFile(partPath, partData, 0644); err != nil {
		return "", fmt.Errorf("写入分片失败: %w", err)
	}
	// 本地存储用文件名（分片序号）作为 ETag
	return strconv.Itoa(partNumber), nil
}

// CompleteMultipartUpload 完成分片合并
func (e *localStorageEngine) CompleteMultipartUpload(key string, uploadId string, parts []MultipartPart) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	absPath := filepath.Join(config.FileConfig.UploadDirectory, key)
	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	mergedFile, err := os.Create(absPath)
	if err != nil {
		return "", fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer mergedFile.Close()

	// 按 PartNumber 排序
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].PartNumber < parts[j].PartNumber
	})

	for _, part := range parts {
		partPath := filepath.Join(uploadId, strconv.Itoa(part.PartNumber))
		partData, err := os.ReadFile(partPath)
		if err != nil {
			return "", fmt.Errorf("读取分片 %d 失败: %w", part.PartNumber, err)
		}
		if _, err = mergedFile.Write(partData); err != nil {
			return "", fmt.Errorf("写入分片 %d 失败: %w", part.PartNumber, err)
		}
	}

	// 清理临时分片目录
	_ = os.RemoveAll(uploadId)
	return key, nil
}

// AbortMultipartUpload 中止分片上传，清理临时目录
func (e *localStorageEngine) AbortMultipartUpload(key string, uploadId string) error {
	return os.RemoveAll(uploadId)
}

// ListParts 列出已上传的分片
func (e *localStorageEngine) ListParts(key string, uploadId string) ([]MultipartPart, error) {
	entries, err := os.ReadDir(uploadId)
	if err != nil {
		// 目录不存在说明没有已上传分片
		if os.IsNotExist(err) {
			return []MultipartPart{}, nil
		}
		return nil, err
	}
	var parts []MultipartPart
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		num, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}
		parts = append(parts, MultipartPart{PartNumber: num, ETag: entry.Name()})
	}
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].PartNumber < parts[j].PartNumber
	})
	return parts, nil
}
