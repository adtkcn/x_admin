package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/app/service/flow_service"
	"x_admin/app/service/iot_service"
	"x_admin/app/service/system_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// IndexHandler 首页控制器
type IndexHandler struct {
	// 防止缓存击穿场景，将大量重复数据库查询合并为一次。
	requestGroup singleflight.Group
}

// @Summary		控制台数据
// @Description	获取控制台统计数据
// @Tags			common_index-首页
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/index/console [get]
func (ih *IndexHandler) Console(c *gin.Context) {
	res, err, _ := ih.requestGroup.Do("Index:Console", func() (any, error) {
		// 工作台各域统计在 controller 直接编排：用户/待办审批/设备/告警
		adminCount, _ := system_service.AdminService.GetTodayCount()
		flowTodo, e := flow_service.ApplyService.GetTodoCount()
		if e != nil {
			return nil, e
		}
		deviceTotal, deviceOnline, e := iot_service.DeviceService.GetConsoleCount()
		if e != nil {
			return nil, e
		}
		alertMinor, alertSerious, alertFault, e := iot_service.AlertService.GetConsoleCount()
		if e != nil {
			return nil, e
		}
		return map[string]any{
			"time":         util.NullTimeUtil.Now(),
			"todayUsers":   adminCount.TodayUsers, // 新增用户
			"totalUsers":   adminCount.TotalUsers, // 总访用户
			"flow_todo":    flowTodo,              // 待办审批
			"deviceTotal":  deviceTotal,           // 设备总量
			"deviceOnline": deviceOnline,          // 在线设备
			"alertMinor":   alertMinor,            // 今日轻微告警
			"alertSerious": alertSerious,          // 今日严重告警
			"alertFault":   alertFault,            // 今日故障
		}, nil
	})
	response.JSON(c, res, err)
}

// @Summary		在线人数趋势
// @Description	获取在线人数曲线初始数据（后续增量由 WS 推送）
// @Tags			common_index-首页
// @Param			token	header		string		true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/index/online-trend [get]
func (ih *IndexHandler) OnlineTrend(c *gin.Context) {
	res, err := common_service.IndexService.OnlineTrend()
	response.JSON(c, res, err)
}

// @Summary		公共配置
// @Description	获取系统公共配置
// @Tags			common_index-首页
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/index/config [get]
func (ih *IndexHandler) Config(c *gin.Context) {
	res, err, _ := ih.requestGroup.Do("Index:Config", func() (any, error) {
		return common_service.IndexService.Config()
	})
	response.JSON(c, res, err)
}
