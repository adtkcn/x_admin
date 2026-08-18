package admin_route

import (
	"x_admin/app/controller/admin_ctl/flow_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func FlowApplyRoute(rg *gin.RouterGroup) {

	handle := flow_controller.FlowApplyHandler{}

	rg = rg.Group("/flow", middleware.PermAuth())
	rg.GET("/flow_apply/list", handle.List)
	rg.GET("/flow_apply/detail", handle.Detail)
	rg.POST("/flow_apply/add", handle.Add)
	rg.POST("/flow_apply/edit", handle.Edit)
	rg.POST("/flow_apply/del", handle.Del)

}

// FlowHistoryRoute(rg)
func FlowHistoryRoute(rg *gin.RouterGroup) {

	handle := flow_controller.FlowHistoryHandler{}

	rg = rg.Group("/flow", middleware.PermAuth())
	rg.GET("/flow_history/list", handle.List)
	rg.GET("/flow_history/list_all", handle.ListAll)
	rg.GET("/flow_history/detail", handle.Detail)
	rg.POST("/flow_history/add", handle.Add)
	rg.POST("/flow_history/edit", handle.Edit)
	rg.POST("/flow_history/del", handle.Del)
	rg.POST("/flow_history/done_hidden", handle.DoneHidden)

	rg.POST("/flow_history/pass", handle.Pass)
	rg.POST("/flow_history/back", handle.Back)

	rg.POST("/flow_history/next_node", handle.NextNode)
	rg.POST("/flow_history/get_approver", handle.GetApprover)
}

// FlowTemplateRoute(rg)
func FlowTemplateRoute(rg *gin.RouterGroup) {

	handle := flow_controller.FlowTemplateHandler{}

	rg = rg.Group("/flow", middleware.PermAuth())
	rg.GET("/flow_template/list", handle.List)
	rg.GET("/flow_template/list_all", handle.ListAll)
	rg.GET("/flow_template/detail", handle.Detail)
	rg.POST("/flow_template/add", handle.Add)
	rg.POST("/flow_template/edit", handle.Edit)
	rg.POST("/flow_template/del", handle.Del)
}
func init() {
	routeHandlers = append(routeHandlers, FlowApplyRoute, FlowHistoryRoute, FlowTemplateRoute)
}
