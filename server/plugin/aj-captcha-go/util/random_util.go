package util

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

// RandomInt 返回 [min, max) 范围内的随机整数（并发安全，使用 crypto/rand）
func RandomInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return min
	}
	return int(n.Int64()) + min
}

// RandString 生成指定长度的随机字符串（并发安全，单次批量读取随机源）
func RandString(codeLen int) string {
	if codeLen <= 0 {
		return ""
	}
	const rawStr = "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
	n := len(rawStr)
	buf := make([]byte, codeLen)
	if _, err := rand.Read(buf); err != nil {
		return ""
	}
	for i := range buf {
		buf[i] = rawStr[int(buf[i])%n]
	}
	return string(buf)
}

// RandUint64 生成随机 uint64（并发安全）
func RandUint64() uint64 {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}
