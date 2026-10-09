package util

import (
	"encoding/json/v2"
	"errors"
	"io"
	"mime/multipart"
	"x_admin/core/response"
	"x_admin/util/validation"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var VerifyUtil = verifyUtil{}

// verifyUtil 参数验证工具类
type verifyUtil struct{}

// paramErr 把入参绑定/校验错误统一翻译成参数校验业务错误。
// 提示文案放在 message 中，不占用 data，保持「失败响应不带数据」的约定。
func paramErr(err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		return response.ParamsValidError.SetMessage(validation.ValidMsg(errs))
	}
	return response.ParamsValidError.SetMessage(err.Error())
}

// VerifyJSON 绑定并校验 JSON body（用 ShouldBindBodyWith 缓存 body，允许后续重复读取）
func (verifyUtil) VerifyJSON(c *gin.Context, obj any) error {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		return paramErr(err)
	}
	return nil
}

// VerifyJSONArray 读取并解析 JSON 数组 body 到 obj（obj 需为切片指针）
func (verifyUtil) VerifyJSONArray(c *gin.Context, obj any) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return paramErr(err)
	}
	if err = json.Unmarshal(body, obj); err != nil {
		return paramErr(err)
	}
	return nil
}

// VerifyBody 按请求 Content-Type 自动绑定并校验（表单/JSON/query 等）
func (verifyUtil) VerifyBody(c *gin.Context, obj any) error {
	if err := c.ShouldBind(obj); err != nil {
		return paramErr(err)
	}
	return nil
}

// VerifyQuery 绑定并校验 URL query 参数
func (verifyUtil) VerifyQuery(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return paramErr(err)
	}
	return nil
}

// VerifyFile 绑定并校验 multipart 上传文件
func (verifyUtil) VerifyFile(c *gin.Context, name string) (*multipart.FileHeader, error) {
	file, err := c.FormFile(name)
	if err != nil {
		return nil, paramErr(err)
	}
	return file, nil
}
