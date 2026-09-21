package util

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"sort"
	"strings"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

// local 通常取决于 http 请求头的 'Accept-Language'
func transInit(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		trans, o = uni.GetTranslator(local)
		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		//register translate
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}
func init() {
	if err := transInit("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
}

var VerifyUtil = verifyUtil{}

// verifyUtil 参数验证工具类
type verifyUtil struct{}

// validMsg 把字段校验错误翻译为可直接展示给调用方的提示文案
func validMsg(errs validator.ValidationErrors) string {
	translated := errs.Translate(trans)
	msgs := make([]string, 0, len(translated))
	for _, v := range translated {
		msgs = append(msgs, v)
	}
	// map 遍历顺序不固定，排序保证同一请求的错误提示顺序稳定
	sort.Strings(msgs)
	return strings.Join(msgs, "; ")
}

// paramErr 把入参绑定/校验错误统一翻译成参数校验业务错误。
// 提示文案放在 message 中，不占用 data，保持「失败响应不带数据」的约定。
func paramErr(err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		return response.ParamsValidError.SetMessage(validMsg(errs))
	}
	return response.ParamsValidError.SetMessage(err.Error())
}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) (e error) {
	// var reqInfo any
	if err := c.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		e = paramErr(err)
		return
	}
	return
}

func (vu verifyUtil) VerifyJSONArray(c *gin.Context, obj any) (e error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e = paramErr(err)
		return
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		e = paramErr(err)
		return
	}
	return
}

func (vu verifyUtil) VerifyBody(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBind(obj); err != nil {
		e = paramErr(err)
		return
	}
	return
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindHeader(obj); err != nil {
		e = paramErr(err)
		return
	}
	return
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		e = paramErr(err)
		return
	}
	return
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader, e error) {
	file, err := c.FormFile(name)
	if err != nil {
		e = paramErr(err)
		return
	}
	return
}
