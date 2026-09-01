package file_util

import "os"

// IsFileExist 判断文件或目录是否存在
func IsFileExist(path string) bool {
	_, err := Stat(path)
	return err == nil || os.IsExist(err)
}
