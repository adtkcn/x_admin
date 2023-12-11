package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/core" 
	"x_admin/middleware" 
	"x_admin/admin/flow_history"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入FlowHistoryRoute注册路由

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
	db := core.GetDB()

	server := flow_history.NewFlowHistoryService(db)

	handle := flow_history.FlowHistoryHandler{Service: server}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/flow_history/list", handle.List)
	rg.GET("/flow_history/listAll", handle.ListAll)
	rg.GET("/flow_history/detail", handle.Detail)
	rg.POST("/flow_history/add", handle.Add)
	rg.POST("/flow_history/edit", handle.Edit)
	rg.POST("/flow_history/del", handle.Del)
}