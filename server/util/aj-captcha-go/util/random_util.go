package util

import (
	"bytes"
	"math/rand"
	"time"
)

// 使用全局变量存储随机数生成器实例
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	return rng.Intn(max-min) + min
}

func RandString(codeLen int) string {
	// 1. 定义原始字符串
	rawStr := "jkwangagDGFHGSERKILMJHSNOPQR546413890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := RandomInt(0, rawStrLen)

		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}
