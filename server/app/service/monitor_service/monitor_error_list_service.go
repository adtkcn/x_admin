package monitor_service

import (
	"time"
	"x_admin/app/model"
	"x_admin/app/schema/monitor_schema"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var MonitorErrorListService = NewMonitorErrorListService()

// NewMonitorErrorListService 初始化
func NewMonitorErrorListService() *monitorErrorListService {
	return &monitorErrorListService{
		db:   core.GetDB(),
		Name: "monitorErrorList",
	}
}

// monitorErrorListService 错误对应的用户记录服务实现类
type monitorErrorListService struct {
	db   *gorm.DB
	Name string
}

// Add 错误对应的用户记录新增
func (service monitorErrorListService) Add(addReq monitor_schema.MonitorErrorListAddReq) (createId string, e error) {
	var obj model.MonitorErrorList
	convert_util.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return "", e
	}

	createId = obj.Id
	return
}

// 删除三个月前的数据
func (service monitorErrorListService) DelThreeMonthAgo() error {
	err := service.db.Where("create_time < ?", time.Now().AddDate(0, -3, 0)).Delete(&model.MonitorErrorList{}).Error

	return err
}
