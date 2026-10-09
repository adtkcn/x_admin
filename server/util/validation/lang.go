package validation

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// trans 校验错误翻译器，由 registerTranslations 初始化
var trans ut.Translator

// registerTranslations 初始化翻译器并注册默认中文错误文案
func register_translations(v *validator.Validate) {
	uni := ut.New(en.New(), zh.New(), en.New())
	var ok bool
	if trans, ok = uni.GetTranslator("zh"); !ok {
		return
	}
	if err := chTranslations.RegisterDefaultTranslations(v, trans); err != nil {
		fmt.Printf("注册中文校验翻译失败: %v\n", err)
	}
}

// ValidMsg 把字段校验错误翻译为提示文案；map 遍历顺序不定，排序保证同一请求顺序稳定
func ValidMsg(errs validator.ValidationErrors) string {
	translated := errs.Translate(trans)
	msgs := make([]string, 0, len(translated))
	for _, v := range translated {
		msgs = append(msgs, v)
	}
	sort.Strings(msgs)
	return strings.Join(msgs, "; ")
}
