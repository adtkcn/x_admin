package util

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var VerifyUtil = verifyUtil{}

// verifyUtil 参数验证工具类
type verifyUtil struct{}

func (vu verifyUtil) VerifyJSON(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
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
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyHeader(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindHeader(obj); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyQuery(c *gin.Context, obj any) (e error) {
	if err := c.ShouldBindQuery(obj); err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}

func (vu verifyUtil) VerifyFile(c *gin.Context, name string) (file *multipart.FileHeader, e error) {
	file, err := c.FormFile(name)
	if err != nil {
		e = response.ParamsValidError.MakeData(err.Error())
		return
	}
	return
}
