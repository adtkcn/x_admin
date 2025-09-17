package commonController

import (
	"errors"
	"x_admin/app/schema/commonSchema"
	"x_admin/app/service/commonService"

	"github.com/gin-gonic/gin"
)

func CaptchaRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/common/captcha")
	rg.POST("/get", func(c *gin.Context) {
		var captchaGet commonSchema.CaptchaGetParams
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
		data, err := commonService.CaptchaGet(captchaGet.CaptchaType)
		if err != nil {
			c.JSON(200, errorRes(err))
			return
		}
		//输出json结果给调用方
		c.JSON(200, successRes(data))
	})
	rg.POST("/check", func(c *gin.Context) {
		var params commonSchema.ClientParams
		if err := c.ShouldBind(&params); err != nil {
			// 返回错误信息
			c.JSON(200, errorRes(err))
			return
		}
		err := commonService.CaptchaCheck(params)

		if err != nil {
			c.JSON(200, errorRes(err))
			return
		}
		//输出json结果给调用方
		c.JSON(200, successRes(nil))
	})

}

func successRes(data interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = false
	ret["repCode"] = "0000"
	ret["repData"] = data
	ret["repMsg"] = nil
	ret["successRes"] = true

	return ret
}
func errorRes(err error) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = true
	ret["repCode"] = "0001"
	ret["repData"] = nil
	ret["repMsg"] = err.Error()
	ret["successRes"] = false
	return ret
}
