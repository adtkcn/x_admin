package util

import (
	"testing"
)

func TestCheckSafeString(t *testing.T) {
	// 测试正常字符串
	if StringUtil.CheckSafeString("abc123") {
		t.Log("正常字符串")
	}
	// 测试包含特殊字符的字符串
	if !StringUtil.CheckSafeString("abc123!") {
		t.Log("包含特殊字符的字符串")
	}
}
