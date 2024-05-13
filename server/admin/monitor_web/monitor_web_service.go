package monitor_web

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

//	type IMonitorWebService interface {
//		List(page request.PageReq, listReq MonitorWebListReq) (res response.PageResp, e error)
//		ListAll() (res []MonitorWebResp, e error)
//
//		Detail(id int) (res MonitorWebResp, e error)
//		Add(addReq MonitorWebAddReq) (e error)
//		Edit(editReq MonitorWebEditReq) (e error)
//		Del(id int) (e error)
//	}
var Service = NewMonitorWebService()

// NewMonitorWebService 初始化
func NewMonitorWebService() *monitorWebService {
	db := core.GetDB()
	return &monitorWebService{db: db}
}

// monitorWebService 错误收集error服务实现类
type monitorWebService struct {
	db *gorm.DB
}

// List 错误收集error列表
func (service monitorWebService) List(page request.PageReq, listReq MonitorWebListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.MonitorWeb{})
	if listReq.ProjectKey != "" {
		dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	}
	if listReq.ClientId != "" {
		dbModel = dbModel.Where("client_id = ?", listReq.ClientId)
	}
	if listReq.EventType != "" {
		dbModel = dbModel.Where("event_type = ?", listReq.EventType)
	}
	if listReq.Page != "" {
		dbModel = dbModel.Where("page = ?", listReq.Page)
	}
	if listReq.Message != "" {
		dbModel = dbModel.Where("message = ?", listReq.Message)
	}
	if listReq.Stack != "" {
		dbModel = dbModel.Where("stack = ?", listReq.Stack)
	}
	if listReq.ClientTimeStart != "" {
		dbModel = dbModel.Where("client_time >= ?", listReq.ClientTimeStart)
	}
	if listReq.ClientTimeEnd != "" {
		dbModel = dbModel.Where("client_time <= ?", listReq.ClientTimeEnd)
	}
	if listReq.CreateTimeStart != "" {
		dbModel = dbModel.Where("create_time >= ?", listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != "" {
		dbModel = dbModel.Where("create_time <= ?", listReq.CreateTimeEnd)
	}

	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.MonitorWeb
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []MonitorWebResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 错误收集error列表
func (service monitorWebService) ListAll() (res []MonitorWebResp, e error) {
	var objs []model.MonitorWeb

	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 错误收集error详情
func (service monitorWebService) Detail(id int) (res MonitorWebResp, e error) {
	var obj model.MonitorWeb
	err := service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 错误收集error新增
func (service monitorWebService) Add(addReq MonitorWebAddReq) (e error) {
	var obj model.MonitorWeb
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return e
	}
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 错误收集error编辑
func (service monitorWebService) Edit(editReq MonitorWebEditReq) (e error) {
	var obj model.MonitorWeb
	err := service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 错误收集error删除
func (service monitorWebService) Del(id int) (e error) {
	var obj model.MonitorWeb
	err := service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	err = service.db.Delete(&obj).Error
	e = response.CheckErr(err, "Del Delete err")
	return
}

// ExportFile 错误收集error导出
func (service monitorWebService) ExportFile(listReq MonitorWebListReq) (res []MonitorWebResp, e error) {
	// 查询
	dbModel := service.db.Model(&model.MonitorWeb{})
	if listReq.ProjectKey != "" {
		dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	}
	if listReq.ClientId != "" {
		dbModel = dbModel.Where("client_id = ?", listReq.ClientId)
	}
	if listReq.EventType != "" {
		dbModel = dbModel.Where("event_type = ?", listReq.EventType)
	}
	if listReq.Page != "" {
		dbModel = dbModel.Where("page = ?", listReq.Page)
	}
	if listReq.Message != "" {
		dbModel = dbModel.Where("message = ?", listReq.Message)
	}
	if listReq.Stack != "" {
		dbModel = dbModel.Where("stack = ?", listReq.Stack)
	}
	// if len(listReq.ClientTime) > 0 {
	// 	dbModel = dbModel.Where("client_time = ?", listReq.ClientTime)
	// }
	// if len(listReq.ClientTime) == 2 {
	// 	dbModel = dbModel.Where("client_time > ?", listReq.ClientTime[0]).Where("client_time < ?", listReq.ClientTime[1])
	// }

	// 数据
	var objs []model.MonitorWeb
	err := dbModel.Order("id asc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []MonitorWebResp{}
	response.Copy(&resps, objs)
	return resps, nil
}

// 导入
func (service monitorWebService) ImportFile(importReq []MonitorWebResp) (e error) {
	var importData []model.MonitorWeb
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "Add Create err")
	return e
}
