package corn

import (
	"x_admin/app/schema"
	"x_admin/app/service/corn_service"
	"x_admin/core"

	"github.com/adtkcn/x_null"
)

var DynamicTasks = NewCronManager()

// 从数据库加载任务
func loadTasks() []corn_service.RunTask {

	allList, err := corn_service.SystemCornService.ListAll(schema.SystemCornListReq{
		Status: x_null.NewInt64(1),
	})
	if err != nil {
		core.Logger.Error("加载任务失败", err)
		return nil
	}

	var RunTaskList = []corn_service.RunTask{} // 运行中的任务列表

	for _, task := range allList {
		if task.Status.ValueOrZero() == 0 {
			continue
		}
		for _, info := range corn_service.TaskInfoList {
			if task.TaskCode.ValueOrZero() == info.TaskCode {

				RunTaskList = append(RunTaskList, corn_service.RunTask{
					TaskId:   task.Id,
					TaskName: task.TaskName.ValueOrZero(),
					CronExpr: task.CornExpr.ValueOrZero(),
					Task:     &info,
				})
				break
			}
		}
	}
	return RunTaskList
}

func init() {
	// 动态任务管理器

	DynamicTasks.Start()

}
