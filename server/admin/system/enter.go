package system

import (
	"x_admin/admin/system/admin"
	"x_admin/admin/system/role"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()

	// permSrv := role.NewSystemAuthPermService(db)
	// roleSrv := role.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := admin.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := NewSystemLoginService(db, adminSrv)

	handle := admin.AdminHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())

	rg.GET("/admin/self", handle.Self)
	rg.GET("/admin/list", handle.List)
	rg.GET("/admin/ListByDeptId", handle.ListByDeptId)
	rg.GET("/admin/detail", handle.Detail)
	rg.POST("/admin/add", middleware.RecordLog("管理员新增"), handle.Add)
	rg.POST("/admin/edit", middleware.RecordLog("管理员编辑"), handle.Edit)
	rg.POST("/admin/upInfo", middleware.RecordLog("管理员更新"), handle.UpInfo)
	rg.POST("/admin/del", middleware.RecordLog("管理员删除"), handle.Del)
	rg.POST("/admin/disable", middleware.RecordLog("管理员状态切换"), handle.Disable)

	rg.GET("/admin/ExportFile", middleware.RecordLog("管理员导出"), handle.ExportFile)

}
func RoleRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := role.NewSystemAuthPermService(db)
	// roleSrv := NewSystemAuthRoleService(db, permSrv)
	// adminSrv := NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := NewSystemLoginService(db, adminSrv)

	// server := role.NewSystemAuthRoleService()

	handle := role.RoleHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.GET("/role/all", handle.All)
	rg.GET("/role/list", middleware.RecordLog("角色列表"), handle.List)
	rg.GET("/role/detail", middleware.RecordLog("角色详情"), handle.Detail)
	rg.POST("/role/add", middleware.RecordLog("角色新增"), handle.Add)
	rg.POST("/role/edit", middleware.RecordLog("角色编辑"), handle.Edit)
	rg.POST("/role/del", middleware.RecordLog("角色删除"), handle.Del)
}
