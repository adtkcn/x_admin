package file_util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
)

// IsFileExist 判断文件或目录是否存在
func IsFileExist(path string) bool {
	_, err := Stat(path)
	return err == nil || os.IsExist(err)
}

// ReaderMd5 计算 io.Reader 流内容的 MD5（十六进制字符串），是文件 MD5 的统一实现
func ReaderMd5(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// FileMd5 按路径计算文件 MD5（复用安全根目录 Open）
func FileMd5(path string) (string, error) {
	f, err := Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return ReaderMd5(f)
}

// GetFileMD5 获取上传文件（multipart）MD5，复用 file_util.ReaderMd5 统一实现
func GetFileMD5(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	return ReaderMd5(f)
}
