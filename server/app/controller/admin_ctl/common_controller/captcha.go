package common_controller

import (
	"errors"
	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"

	"github.com/gin-gonic/gin"
)

func CaptchaRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/common/captcha")
	rg.POST("/get", func(c *gin.Context) {
		var captchaGet common_schema.CaptchaGetParams
		if err := c.ShouldBind(&captchaGet); err != nil {
			// 返回错误信息
			c.JSON(200, errorRes(err))
			return
		}
		if captchaGet.CaptchaType == "" {
			c.JSON(200, errorRes(errors.New("验证码类型不能为空")))
			return
		}
		// 根据参数类型获取不同服务即可
		data, err := common_service.CaptchaGet(captchaGet.CaptchaType)
		if err != nil {
			c.JSON(200, errorRes(err))
			return
		}
		//输出json结果给调用方
		c.JSON(200, successRes(data))
	})
	rg.POST("/check", func(c *gin.Context) {
		var params common_schema.ClientParams
		if err := c.ShouldBind(&params); err != nil {
			// 返回错误信息
			c.JSON(200, errorRes(err))
			return
		}
		err := common_service.CaptchaCheck(params)

		if err != nil {
			c.JSON(200, errorRes(err))
			return
		}
		//输出json结果给调用方
		c.JSON(200, successRes(nil))
	})

}

func successRes(data any) map[string]any {
	ret := make(map[string]any)
	ret["error"] = false
	ret["repCode"] = "0000"
	ret["repData"] = data
	ret["repMsg"] = nil
	ret["successRes"] = true

	return ret
}
func errorRes(err error) map[string]any {
	ret := make(map[string]any)
	ret["error"] = true
	ret["repCode"] = "0001"
	ret["repData"] = nil
	ret["repMsg"] = err.Error()
	ret["successRes"] = false
	return ret
}
