package admin

import (
	"x_admin/admin/monitor_web"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入MonitorWebRoute注册路由

3. 后台手动添加菜单和按钮
admin:monitor_web:add
admin:monitor_web:edit
admin:monitor_web:del
admin:monitor_web:list
admin:monitor_web:listAll
admin:monitor_web:detail
admin:monitor_web:ExportFile
admin:monitor_web:ImportFile

*/

// MonitorWebRoute(rg)
func MonitorWebRoute(rg *gin.RouterGroup) {
	handle := monitor_web.MonitorWebHandler{}

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/monitor_web/list", handle.List)
	r.GET("/monitor_web/listAll", handle.ListAll)
	r.GET("/monitor_web/detail", handle.Detail)
	r.POST("/monitor_web/add", middleware.RecordLog("错误收集error新增"), handle.Add)
	r.POST("/monitor_web/edit", middleware.RecordLog("错误收集error编辑"), handle.Edit)
	r.POST("/monitor_web/del", middleware.RecordLog("错误收集error删除"), handle.Del)
	r.GET("/monitor_web/ExportFile", middleware.RecordLog("错误收集error导出"), handle.ExportFile)
	r.POST("/monitor_web/ImportFile", handle.ImportFile)
}
