package admin_route

import (
	"x_admin/app/controller/admin_ctl/settingController"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// SettingRoute 系统设置模块路由（版权、字典数据、字典类型、网站）
func SettingRoute(rg *gin.RouterGroup) {

	// 版权设置
	handleCopyright := settingController.CopyrightHandler{}
	copyrightRg := rg.Group("/setting")
	copyrightRg.GET("/copyright/detail", handleCopyright.Detail)
	copyrightRg.POST("/copyright/save", handleCopyright.Save)

	// 字典数据
	handleDictData := settingController.DictDataHandler{}
	notAuthDictData := rg.Group("/setting", middleware.LoginAuth())
	notAuthDictData.GET("/dict/data/all", handleDictData.All)

	dictDataRg := rg.Group("/setting", middleware.PermAuth())
	dictDataRg.GET("/dict/data/detail", handleDictData.Detail)
	dictDataRg.POST("/dict/data/add", handleDictData.Add)
	dictDataRg.POST("/dict/data/edit", handleDictData.Edit)
	dictDataRg.POST("/dict/data/del", handleDictData.Del)

	// 字典类型
	handleDictType := settingController.DictTypeHandler{}
	notAuthDictType := rg.Group("/setting", middleware.LoginAuth())
	notAuthDictType.GET("/dict/type/all", handleDictType.All)

	dictTypeRg := rg.Group("/setting", middleware.PermAuth())
	dictTypeRg.GET("/dict/type/list", handleDictType.List)
	dictTypeRg.GET("/dict/type/detail", handleDictType.Detail)
	dictTypeRg.POST("/dict/type/add", handleDictType.Add)
	dictTypeRg.POST("/dict/type/edit", handleDictType.Edit)
	dictTypeRg.POST("/dict/type/del", handleDictType.Del)

	// 网站设置
	handleWebsite := settingController.WebsiteHandler{}
	websiteRg := rg.Group("/setting", middleware.PermAuth())
	websiteRg.GET("/website/detail", handleWebsite.Detail)
	websiteRg.POST("/website/save", handleWebsite.Save)
}

func init() {
	routeHandlers = append(routeHandlers, SettingRoute)
}
