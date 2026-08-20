package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json/v2"
	"io"
	"math"
	"math/rand"
	"mime/multipart"
	"os"
	"slices"
	"time"

	"uuid"
)

var (
	ToolsUtil    = toolsUtil{}
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// toolsUtil 常用工具集合类
type toolsUtil struct{}

// Random 返回随机数
func (tu toolsUtil) Random(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomString 返回随机字符串
func (tu toolsUtil) RandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[r.Intn(62)]
	}
	return string(byteList)
}

// MakeUuidV7 制作UUID v7
func (tu toolsUtil) MakeUuidV7() string {
	v7 := uuid.NewV7()
	return v7.String()
}

// MakeMd5 制作MD5
func (tu toolsUtil) MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// GetFileMD5 获取文件MD5
func (tu toolsUtil) GetFileMD5(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Contains 判断list是否包含elem元素（支持任意可比较类型）
func (tu toolsUtil) Contains[T comparable](list []T, elem T) bool {
	return slices.Contains(list, elem)
}

// Round float四舍五入
func (tu toolsUtil) Round(val float64, n int) float64 {
	base := math.Pow(10, float64(n))
	return math.Round(base*val) / base
}

// JsonToObj JSON转Obj
func (tu toolsUtil) JsonToObj[T any](jsonStr string) (t T, err error) {
	err = json.Unmarshal([]byte(jsonStr), &t)
	return t, err
}

// ObjToJson Obj转JSON
func (tu toolsUtil) ObjToJson(data any) (res string, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return res, err
	}
	res = string(b)
	return res, nil
}

// IsFileExist 判断文件或目录是否存在
func (tu toolsUtil) IsFileExist(path string) bool {
	var root, err = os.OpenRoot(".")
	if err != nil {
		return false
	}
	defer root.Close()
	_, err = root.Stat(path)
	return err == nil || os.IsExist(err)
}

// 创建文件夹
func (tu toolsUtil) CreateDir(path string) error {
	var root, err = os.OpenRoot(".")
	if err != nil {
		return err
	}
	defer root.Close()
	return root.Mkdir(path, 0755)
}
func (tu toolsUtil) WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
