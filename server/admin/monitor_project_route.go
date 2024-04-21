package admin

import (
	"x_admin/admin/monitor_project"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入MonitorProjectRoute注册路由

3. 后台手动添加菜单和按钮
monitor_project:add
monitor_project:edit
monitor_project:del
monitor_project:list
monitor_project:listAll
monitor_project:detail
*/

// MonitorProjectRoute(rg)
func MonitorProjectRoute(rg *gin.RouterGroup) {
	handle := monitor_project.MonitorProjectHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/monitor_project/list", handle.List)
	rg.GET("/monitor_project/listAll", handle.ListAll)
	rg.GET("/monitor_project/detail", handle.Detail)
	rg.POST("/monitor_project/add", middleware.RecordLog("错误项目新增"), handle.Add)
	rg.POST("/monitor_project/edit", middleware.RecordLog("错误项目编辑"), handle.Edit)
	rg.POST("/monitor_project/del", middleware.RecordLog("错误项目删除"), handle.Del)
	rg.GET("/monitor_project/ExportFile", middleware.RecordLog("错误项目导出"), handle.ExportFile)
	rg.POST("/monitor_project/ImportFile", handle.ImportFile)
}
