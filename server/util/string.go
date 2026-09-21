package util

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var StringUtil = stringUtil{}

// arrayUtil 数组工具类
type stringUtil struct{}

// 转换为蛇形命名
func (su stringUtil) ToSnakeCase(s string) string {
	buf := bytes.Buffer{}
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// 转换为小驼峰命名
func (su stringUtil) ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	c := cases.Title(language.Und, cases.NoLower)
	for i := 1; i < len(words); i++ {
		words[i] = c.String(words[i])
	}
	// return s
	return strings.Join(words, "")
}

// 转换为大驼峰命名
func (su stringUtil) ToUpperCamelCase(s string) string {
	if s == "id" {
		return "ID"
	}
	words := strings.Split(s, "_")
	c := cases.Title(language.Und, cases.NoLower)
	for i := 0; i < len(words); i++ {
		words[i] = c.String(words[i])
	}
	return strings.Join(words, "")
}

// 检查字符串只能包含字母、数字和下划线
func (su stringUtil) CheckSafeString(s string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return !reg.MatchString(s)
}
