package util

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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

// loca 通常取决于 http 请求头的 'Accept-Language'
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

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e = response.ParamsValidError.MakeData(err.Error())
			return
		}
		e = response.ParamsValidError.MakeData(errs.Translate(trans))
		return
	}
	return
}

func (vu verifyUtil) VerifyJSONArray(c *gin.Context, obj any) (e error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyBody(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBind(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e = response.ParamsValidError.MakeData(err.Error())
			return
		}
		e = response.ParamsValidError.MakeData(errs.Translate(trans))
		return
	}
	return
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindHeader(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e = response.ParamsValidError.MakeData(err.Error())
			return
		}
		e = response.ParamsValidError.MakeData(errs.Translate(trans))
		return
	}
	return
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e = response.ParamsValidError.MakeData(err.Error())
			return
		}
		e = response.ParamsValidError.MakeData(errs.Translate(trans))
		return
	}
	return
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader, e error) {
	file, err := c.FormFile(name)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e = response.ParamsValidError.MakeData(err.Error())
			return
		}
		e = response.ParamsValidError.MakeData(errs.Translate(trans))
		return
	}
	return
}
