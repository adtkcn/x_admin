package util

import (
	"bytes"
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

// RandString 生成指定长度的随机字符串（并发安全）
func RandString(codeLen int) string {
	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
	rawStrLen := big.NewInt(int64(len(rawStr)))

	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	for codeLen > 0 {
		n, err := rand.Int(rand.Reader, rawStrLen)
		if err != nil {
			codeLen--
			continue
		}
		b.WriteByte(rawStr[n.Int64()])
		codeLen--
	}
	return b.String()
}

// RandUint64 生成随机 uint64（并发安全）
func RandUint64() uint64 {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(b)
}
