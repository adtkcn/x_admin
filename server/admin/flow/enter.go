package flow

import (
	"x_admin/admin/flow/flow_apply"
	"x_admin/admin/flow/flow_history"
	"x_admin/admin/flow/flow_template"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

// 3. 后台手动添加菜单和按钮
// flow_apply:add
// flow_apply:edit
// flow_apply:del
// flow_apply:list
// flow_apply:detail

func FlowApplyRoute(rg *gin.RouterGroup) {

	handle := flow_apply.FlowApplyHandler{}

	rg = rg.Group("/flow", middleware.TokenAuth())
	rg.GET("/flow_apply/list", handle.List)
	rg.GET("/flow_apply/detail", handle.Detail)
	rg.POST("/flow_apply/add", handle.Add)
	rg.POST("/flow_apply/edit", handle.Edit)
	rg.POST("/flow_apply/del", handle.Del)

}

/**
3. 后台手动添加菜单和按钮
flow_history:add
flow_history:edit
flow_history:del
flow_history:list
flow_history:listAll
flow_history:detail
*/

// FlowHistoryRoute(rg)
func FlowHistoryRoute(rg *gin.RouterGroup) {

	handle := flow_history.FlowHistoryHandler{}

	rg = rg.Group("/flow", middleware.TokenAuth())
	rg.GET("/flow_history/list", handle.List)
	rg.GET("/flow_history/listAll", handle.ListAll)
	rg.GET("/flow_history/detail", handle.Detail)
	rg.POST("/flow_history/add", handle.Add)
	rg.POST("/flow_history/edit", handle.Edit)
	rg.POST("/flow_history/del", handle.Del)

	rg.POST("/flow_history/pass", handle.Pass)
	rg.POST("/flow_history/back", handle.Back)

	rg.POST("/flow_history/next_node", handle.NextNode)
	rg.POST("/flow_history/get_approver", handle.GetApprover)
}

/**

3. 后台手动添加菜单和按钮
flow_template:add
flow_template:edit
flow_template:del
flow_template:list
flow_template:listAll
flow_template:detail
*/

// FlowTemplateRoute(rg)
func FlowTemplateRoute(rg *gin.RouterGroup) {

	handle := flow_template.FlowTemplateHandler{}

	rg = rg.Group("/flow", middleware.TokenAuth())
	rg.GET("/flow_template/list", handle.List)
	rg.GET("/flow_template/listAll", handle.ListAll)
	rg.GET("/flow_template/detail", handle.Detail)
	rg.POST("/flow_template/add", handle.Add)
	rg.POST("/flow_template/edit", handle.Edit)
	rg.POST("/flow_template/del", handle.Del)
}
