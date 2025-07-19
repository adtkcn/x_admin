package commonService

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
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
func (upSrv uploadChunkService) MergeChunk(fileMd5, filePath string, chunkCount int) error {

	chunkDir := fmt.Sprintf("./tmp/%s", fileMd5)

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
