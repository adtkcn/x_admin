package storage

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
	"time"
	"x_admin/config"
)

// ---- 预签名 URL 工具 ----
// 采用 HMAC-SHA256 签名机制，仅依赖标准库，无需第三方 SDK

var (
	_presignSecret     []byte
	_presignSecretOnce sync.Once
)

// getPresignSecret 获取预签名密钥（懒加载）
// 优先使用 config.FileConfig.PresignSecret，若为空则自动生成随机密钥
func getPresignSecret() []byte {
	_presignSecretOnce.Do(func() {
		secret := config.FileConfig.PresignSecret
		if secret != "" {
			_presignSecret = []byte(secret)
		} else {
			// 自动生成 32 字节随机密钥（应用重启后失效，生产环境请在配置中指定固定密钥）
			buf := make([]byte, 32)
			if _, err := rand.Read(buf); err != nil {
				log.Printf("presign: 生成随机密钥失败: %v, 使用默认密钥", err)
				_presignSecret = []byte("x_admin_presign_fallback_secret")
			} else {
				_presignSecret = buf
				log.Printf("presign: 已自动生成随机预签名密钥（应用重启后失效，建议在配置中指定 PresignSecret）")
			}
		}
	})
	return _presignSecret
}

// sign 计算 HMAC-SHA256 签名
func sign(payload string) string {
	mac := hmac.New(sha256.New, getPresignSecret())
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

// PresignParams 预签名 URL 参数
type PresignParams struct {
	Sig       string `json:"sig"`       // 签名
	FileKey   string `json:"fileKey"`   // 文件存储 key
	ExpiresAt int64  `json:"expiresAt"` // 过期时间戳（秒）
	Action    string `json:"action"`    // 操作：putObject / uploadPart
	// uploadPart 专用
	UploadId   string `json:"uploadId,omitempty"`
	PartNumber int    `json:"partNumber,omitempty"`
}

// GeneratePresignedParams 生成预签名参数
// action: "putObject" 或 "uploadPart"
// fileKey: 文件存储 key
// uploadId/partNumber: 分片上传时需要
func GeneratePresignedParams(action, fileKey, uploadId string, partNumber int) (*PresignParams, error) {
	expiresAt := time.Now().Unix() + int64(config.FileConfig.PresignExpire)
	if config.FileConfig.PresignExpire <= 0 {
		expiresAt = time.Now().Unix() + 3600
	}

	// 签名载荷：action|fileKey|uploadId|partNumber|expiresAt
	payload := fmt.Sprintf("%s|%s|%s|%d|%d", action, fileKey, uploadId, partNumber, expiresAt)
	sig := sign(payload)

	return &PresignParams{
		Sig:        sig,
		FileKey:    fileKey,
		ExpiresAt:  expiresAt,
		Action:     action,
		UploadId:   uploadId,
		PartNumber: partNumber,
	}, nil
}

// VerifyPresign 校验预签名参数
func VerifyPresign(action, fileKey, uploadId string, partNumber int, sig string, expiresAt int64) bool {
	// 检查过期
	if time.Now().Unix() > expiresAt {
		return false
	}
	// 重新计算签名
	payload := fmt.Sprintf("%s|%s|%s|%d|%d", action, fileKey, uploadId, partNumber, expiresAt)
	expectedSig := sign(payload)
	return hmac.Equal([]byte(sig), []byte(expectedSig))
}
