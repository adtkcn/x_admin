package index

import (
	"x_admin/core/response"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

func IndexRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// authSrv := system.NewSystemAuthMenuService(db, permSrv)
	// IndexService := NewIndexService()
	handle := indexHandler{}

	rg = rg.Group("/common", middleware.TokenAuth())
	rg.GET("/index/console", handle.console)
	rg.GET("/index/config", handle.config)

}

type indexHandler struct{}

// console 控制台
func (ih indexHandler) console(c *gin.Context) {
	res, err := Service.Console()
	response.CheckAndRespWithData(c, res, err)
}

// config 公共配置
func (ih indexHandler) config(c *gin.Context) {
	res, err := Service.Config()
	response.CheckAndRespWithData(c, res, err)
}
