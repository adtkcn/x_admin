package monitor_error_list

import (
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model"
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
func (service monitorErrorListService) Add(addReq MonitorErrorListAddReq) (createId int, e error) {
	var obj model.MonitorErrorList
	convert_util.StructToStruct(addReq, &obj)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0, e
	}

	createId = obj.Id
	return
}
