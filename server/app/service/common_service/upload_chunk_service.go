package common_service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"x_admin/util"
)

var UploadChunkService = NewUploadChunkService()

// NewUploadService 初始化
func NewUploadChunkService() *uploadChunkService {
	return &uploadChunkService{}
}

// uploadService 上传服务实现类
type uploadChunkService struct{}

// CheckFileExist 检查文件是否存在
func (upSrv uploadChunkService) CheckFileExist(filePath string) bool {
	return util.ToolsUtil.IsFileExist(filePath)
}

func (upSrv uploadChunkService) UploadChunk(chunkDir string, chunkPath string, chunk *multipart.FileHeader) error {
	os.MkdirAll(chunkDir, 0755)

	chunkData, err := chunk.Open()
	if err != nil {
		return err
	}
	defer chunkData.Close()

	chunkBytes, err := io.ReadAll(chunkData)
	if err != nil {
		return err
	}
	return os.WriteFile(chunkPath, chunkBytes, 0644)
}

// 通过文件列表中的文件名获取最大chunk
func (upSrv uploadChunkService) HasChunk(chunkDir string) []int {
	// chunks, err := os.ReadDir(chunkDir)
	files, err := os.ReadDir(chunkDir)
	var chunks []int
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
		return chunks
	}

	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		filename := file.Name()
		num, err := strconv.Atoi(filename) // 直接转换整个文件名
		if err != nil {
			continue // 跳过非数字文件名
		}
		chunks = append(chunks, num)
	}
	return chunks
}
func (upSrv uploadChunkService) MergeChunk(chunkDir, filePath string, chunkCount int) error {

	// chunkDir := fmt.Sprintf("./tmp/%s", fileMd5)

	chunks, err := os.ReadDir(chunkDir)
	if err != nil {
		return errors.New("分片不存在")
	}
	// 判断切片数量是否一致
	if len(chunks) != chunkCount {
		return errors.New("分片未全部上传，请重试")
	}
	mergedFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer mergedFile.Close()
	for index := range chunkCount {
		chunkData, err := os.ReadFile(filepath.Join(chunkDir, fmt.Sprintf("%d", index)))
		if err != nil {
			return err
		}

		_, err = mergedFile.Write(chunkData)
		if err != nil {
			return err
		}
	}
	return os.RemoveAll(chunkDir) // 清理临时分片
}
