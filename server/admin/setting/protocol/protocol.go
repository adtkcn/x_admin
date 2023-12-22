package protocol

import (
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func ProtocolRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewSettingProtocolService(db)

	handle := protocolHandler{}

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/protocol/detail", handle.Detail)
	rg.POST("/protocol/save", handle.save)
}

type protocolHandler struct {
}

// detail 获取政策信息
func (ph protocolHandler) Detail(c *gin.Context) {
	res, err := Service.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// save 保存政策信息
func (ph protocolHandler) save(c *gin.Context) {
	var pReq SettingProtocolReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &pReq)) {
		return
	}
	response.CheckAndResp(c, Service.Save(pReq))
}
