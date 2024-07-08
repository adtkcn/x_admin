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

	r := rg.Group("/", middleware.TokenAuth())

	r.GET("/monitor_project/list", handle.List)
	r.GET("/monitor_project/listAll", handle.ListAll)
	r.GET("/monitor_project/detail", handle.Detail)
	r.POST("/monitor_project/add", middleware.RecordLog("错误项目新增"), handle.Add)
	r.POST("/monitor_project/edit", middleware.RecordLog("错误项目编辑"), handle.Edit)
	r.POST("/monitor_project/del", middleware.RecordLog("错误项目删除"), handle.Del)
	r.GET("/monitor_project/ExportFile", middleware.RecordLog("错误项目导出"), handle.ExportFile)
	r.POST("/monitor_project/ImportFile", handle.ImportFile)
}
