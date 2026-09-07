package admin_route

import (
	"x_admin/app/controller/admin_ctl/system_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// initLoginRoute 登录路由（无需认证）
func initLoginRoute(rg *gin.RouterGroup) {
	handleLogin := system_controller.LoginHandler{}
	loginRg := rg.Group("/system")
	loginRg.POST("/login", handleLogin.Login)
	loginRg.POST("/logout", handleLogin.Logout)
	// 忘记密码（无需认证）
	loginRg.POST("/forgot-pwd/send-code", handleLogin.ForgotPwdSendCode)
	loginRg.POST("/forgot-pwd/reset", handleLogin.ForgotPwdReset)
}

// initAdminRoute 管理员路由（部分接口仅登录即可，其余需权限认证）
func initAdminRoute(rg *gin.RouterGroup) {
	handleAdmin := system_controller.AdminHandler{}

	notAuthAdmin := rg.Group("/system", middleware.LoginAuth())
	notAuthAdmin.GET("/admin/self", handleAdmin.Self)
	notAuthAdmin.POST("/admin/sendEmailCode", middleware.LimitIP(2, 60), middleware.LimitEmail(2, 60), handleAdmin.SendEmailCode)
	notAuthAdmin.POST("/admin/upInfo", middleware.RecordLog("管理员更新"), handleAdmin.UpInfo)

	authAdmin := rg.Group("/system", middleware.PermAuth())
	authAdmin.GET("/admin/list", handleAdmin.List)
	authAdmin.GET("/admin/list_all", handleAdmin.ListAll)
	authAdmin.GET("/admin/ListByDeptId", handleAdmin.ListByDeptId)
	authAdmin.GET("/admin/detail", handleAdmin.Detail)
	authAdmin.POST("/admin/add", middleware.RecordLog("管理员新增"), handleAdmin.Add)
	authAdmin.POST("/admin/edit", middleware.RecordLog("管理员编辑"), handleAdmin.Edit)
	authAdmin.POST("/admin/del", middleware.RecordLog("管理员删除"), handleAdmin.Del)
	authAdmin.POST("/admin/disable", middleware.RecordLog("管理员状态切换"), handleAdmin.Disable)
	authAdmin.GET("/admin/export_file", middleware.RecordLog("管理员导出"), handleAdmin.ExportFile)
	authAdmin.POST("/admin/import_file", handleAdmin.ImportFile)
}

// initMenuRoute 菜单路由（route 接口仅需登录）
func initMenuRoute(rg *gin.RouterGroup) {
	handleMenu := system_controller.MenuHandler{}

	notAuthMenu := rg.Group("/system", middleware.LoginAuth())
	notAuthMenu.GET("/menu/route", handleMenu.Route)

	authMenu := rg.Group("/system", middleware.PermAuth())
	authMenu.GET("/menu/list", handleMenu.List)
	authMenu.GET("/menu/detail", handleMenu.Detail)
	authMenu.POST("/menu/add", handleMenu.Add)
	authMenu.POST("/menu/edit", handleMenu.Edit)
	authMenu.POST("/menu/del", handleMenu.Del)
}

// initPostRoute 岗位路由（all 接口仅需登录）
func initPostRoute(rg *gin.RouterGroup) {
	handlePost := system_controller.PostHandler{}

	notAuthPost := rg.Group("/system", middleware.LoginAuth())
	notAuthPost.GET("/post/all", handlePost.All)

	authPost := rg.Group("/system", middleware.PermAuth())
	authPost.GET("/post/list", handlePost.List)
	authPost.GET("/post/detail", handlePost.Detail)
	authPost.POST("/post/add", handlePost.Add)
	authPost.POST("/post/edit", handlePost.Edit)
	authPost.POST("/post/del", handlePost.Del)
}

// initDeptRoute 部门路由（list 接口仅需登录）
func initDeptRoute(rg *gin.RouterGroup) {
	handleDept := system_controller.DeptHandler{}

	notAuthDept := rg.Group("/system", middleware.LoginAuth())
	notAuthDept.GET("/dept/list", handleDept.List)

	authDept := rg.Group("/system", middleware.PermAuth())
	authDept.GET("/dept/all", handleDept.All)
	authDept.GET("/dept/detail", handleDept.Detail)
	authDept.POST("/dept/add", handleDept.Add)
	authDept.POST("/dept/edit", handleDept.Edit)
	authDept.POST("/dept/del", handleDept.Del)
	authDept.POST("/dept/sort", handleDept.Sort)
}

// initRoleRoute 角色路由（all 接口仅需登录）
func initRoleRoute(rg *gin.RouterGroup) {
	handleRole := system_controller.RoleHandler{}

	notAuthRole := rg.Group("/system", middleware.LoginAuth())
	notAuthRole.GET("/role/all", handleRole.All)

	authRole := rg.Group("/system", middleware.PermAuth())
	authRole.GET("/role/list", handleRole.List)
	authRole.GET("/role/detail", middleware.RecordLog("角色详情"), handleRole.Detail)
	authRole.POST("/role/add", middleware.RecordLog("角色新增"), handleRole.Add)
	authRole.POST("/role/edit", middleware.RecordLog("角色编辑"), handleRole.Edit)
	authRole.POST("/role/del", middleware.RecordLog("角色删除"), handleRole.Del)
}

// initLogRoute 日志路由
func initLogRoute(rg *gin.RouterGroup) {
	handleLog := system_controller.LogHandler{}
	rgLog := rg.Group("/system", middleware.PermAuth())
	rgLog.GET("/log/operate", handleLog.Operate)
	rgLog.GET("/log/login", handleLog.Login)
}

func init() {
	routeHandlers = append(routeHandlers, initLoginRoute, initAdminRoute, initMenuRoute, initPostRoute, initDeptRoute, initRoleRoute, initLogRoute)
}
