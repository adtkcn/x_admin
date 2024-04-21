package admin

import (
	"x_admin/admin/monitor_client"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入MonitorClientRoute注册路由

3. 后台手动添加菜单和按钮
monitor_client:add
monitor_client:edit
monitor_client:del
monitor_client:list
monitor_client:listAll
monitor_client:detail
*/

// MonitorClientRoute(rg)
func MonitorClientRoute(rg *gin.RouterGroup) {
	handle := monitor_client.MonitorClientHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/monitor_client/list", handle.List)
	rg.GET("/monitor_client/listAll", handle.ListAll)
	rg.GET("/monitor_client/detail", handle.Detail)
	rg.POST("/monitor_client/add", middleware.RecordLog("客户端信息新增"), handle.Add)
	rg.POST("/monitor_client/edit", middleware.RecordLog("客户端信息编辑"), handle.Edit)
	rg.POST("/monitor_client/del", middleware.RecordLog("客户端信息删除"), handle.Del)
	rg.GET("/monitor_client/ExportFile", middleware.RecordLog("客户端信息导出"), handle.ExportFile)
	rg.POST("/monitor_client/ImportFile", handle.ImportFile)
}
