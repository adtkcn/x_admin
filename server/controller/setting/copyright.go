package copyright

import (
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func CopyrightRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// service := NewSettingCopyrightService(db)
	handle := copyrightHandler{}

	rg = rg.Group("/setting")
	rg.GET("/copyright/detail", handle.Detail)
	rg.POST("/copyright/save", handle.save)
}

type copyrightHandler struct {
}

// detail 获取备案信息
func (ch copyrightHandler) Detail(c *gin.Context) {
	res, err := Service.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// save 保存备案信息
func (ch copyrightHandler) save(c *gin.Context) {
	var cReqs []SettingCopyrightItemReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.CheckAndResp(c, Service.Save(cReqs))
}
