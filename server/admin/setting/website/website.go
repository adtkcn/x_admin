package website

import (
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func WebsiteRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewSettingWebsiteService(db)

	handle := websiteHandler{}

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/website/detail", handle.Detail)
	rg.POST("/website/save", handle.save)
}

type websiteHandler struct{}

// detail 获取网站信息
func (wh websiteHandler) Detail(c *gin.Context) {
	res, err := Service.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// save 保存网站信息
func (wh websiteHandler) save(c *gin.Context) {
	var wsReq SettingWebsiteReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
		return
	}
	response.CheckAndResp(c, Service.Save(wsReq))
}
