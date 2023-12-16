package admin

import (
	"x_admin/admin/flow_apply"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入FlowApplyRoute注册路由

3. 后台手动添加菜单和按钮
flow_apply:add
flow_apply:edit
flow_apply:del
flow_apply:list
flow_apply:detail
*/

// FlowApplyRoute(rg)
func FlowApplyRoute(rg *gin.RouterGroup) {

	handle := flow_apply.FlowApplyHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/flow_apply/list", handle.List)
	rg.GET("/flow_apply/detail", handle.Detail)
	rg.POST("/flow_apply/add", handle.Add)
	rg.POST("/flow_apply/edit", handle.Edit)
	rg.POST("/flow_apply/del", handle.Del)

}
