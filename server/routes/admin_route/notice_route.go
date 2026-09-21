package admin_route

import (
	"x_admin/app/controller/admin_ctl/system_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routeHandlers = append(routeHandlers, initNoticeRoute)
}

func initNoticeRoute(rg *gin.RouterGroup) {
	handle := system_controller.NoticeHandler{}

	// 仅需登录
	notAuth := rg.Group("/system", middleware.LoginAuth())
	notAuth.GET("/notice/list", handle.List)
	notAuth.GET("/notice/unread_count", handle.UnreadCount)
	notAuth.GET("/notice/setting", handle.GetSetting)

	// 需要权限
	notAuth.POST("/notice/read", handle.Read)
	notAuth.POST("/notice/read_all", handle.ReadAll)
	notAuth.POST("/notice/del", handle.Del)
	notAuth.POST("/notice/setting/save", handle.SaveSetting)
}
