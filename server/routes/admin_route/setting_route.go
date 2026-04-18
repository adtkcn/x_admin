package admin_route

import (
	"x_admin/app/controller/admin_ctl/settingController"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// initCopyrightRoute 版权设置路由
func initCopyrightRoute(rg *gin.RouterGroup) {
	handleCopyright := settingController.CopyrightHandler{}
	copyrightRg := rg.Group("/setting")
	copyrightRg.GET("/copyright/detail", handleCopyright.Detail)
	copyrightRg.POST("/copyright/save", handleCopyright.Save)
}

// initDictDataRoute 字典数据路由（all 接口仅需登录）
func initDictDataRoute(rg *gin.RouterGroup) {
	handleDictData := settingController.DictDataHandler{}

	notAuthDictData := rg.Group("/setting", middleware.LoginAuth())
	notAuthDictData.GET("/dict/data/all", handleDictData.All)

	dictDataRg := rg.Group("/setting", middleware.PermAuth())
	dictDataRg.GET("/dict/data/detail", handleDictData.Detail)
	dictDataRg.POST("/dict/data/add", handleDictData.Add)
	dictDataRg.POST("/dict/data/edit", handleDictData.Edit)
	dictDataRg.POST("/dict/data/del", handleDictData.Del)
}

// initDictTypeRoute 字典类型路由（all 接口仅需登录）
func initDictTypeRoute(rg *gin.RouterGroup) {
	handleDictType := settingController.DictTypeHandler{}

	notAuthDictType := rg.Group("/setting", middleware.LoginAuth())
	notAuthDictType.GET("/dict/type/all", handleDictType.All)

	dictTypeRg := rg.Group("/setting", middleware.PermAuth())
	dictTypeRg.GET("/dict/type/list", handleDictType.List)
	dictTypeRg.GET("/dict/type/detail", handleDictType.Detail)
	dictTypeRg.POST("/dict/type/add", handleDictType.Add)
	dictTypeRg.POST("/dict/type/edit", handleDictType.Edit)
	dictTypeRg.POST("/dict/type/del", handleDictType.Del)
}

// initWebsiteRoute 网站设置路由
func initWebsiteRoute(rg *gin.RouterGroup) {
	handleWebsite := settingController.WebsiteHandler{}
	websiteRg := rg.Group("/setting", middleware.PermAuth())
	websiteRg.GET("/website/detail", handleWebsite.Detail)
	websiteRg.POST("/website/save", handleWebsite.Save)
}

// SettingRoute 系统设置模块路由入口（版权、字典数据、字典类型、网站）
func init() {
	routeHandlers = append(routeHandlers, initCopyrightRoute, initDictDataRoute, initDictTypeRoute, initWebsiteRoute)
}
