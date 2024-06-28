package monitor_client

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

var Service = NewMonitorClientService()

// NewMonitorClientService 初始化
func NewMonitorClientService() *monitorClientService {
	db := core.GetDB()
	return &monitorClientService{db: db}
}

// monitorClientService 监控-客户端信息服务实现类
type monitorClientService struct {
	db *gorm.DB
}

// List 监控-客户端信息列表
func (service monitorClientService) GetModel(listReq MonitorClientListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorClient{})
	if listReq.ProjectKey != "" {
		dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	}
	if listReq.ClientId != "" {
		dbModel = dbModel.Where("client_id = ?", listReq.ClientId)
	}
	if listReq.UserId != "" {
		dbModel = dbModel.Where("user_id = ?", listReq.UserId)
	}
	if listReq.Os != "" {
		dbModel = dbModel.Where("os = ?", listReq.Os)
	}
	if listReq.Browser != "" {
		dbModel = dbModel.Where("browser = ?", listReq.Browser)
	}
	if listReq.City != "" {
		dbModel = dbModel.Where("city = ?", listReq.City)
	}
	if listReq.Width > 0 {
		dbModel = dbModel.Where("width = ?", listReq.Width)
	}
	if listReq.Height > 0 {
		dbModel = dbModel.Where("height = ?", listReq.Height)
	}
	if listReq.Ua != "" {
		dbModel = dbModel.Where("ua = ?", listReq.Ua)
	}
	if listReq.CreateTimeStart != "" {
		dbModel = dbModel.Where("create_time >= ?", listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != "" {
		dbModel = dbModel.Where("create_time <= ?", listReq.CreateTimeEnd)
	}
	if listReq.ClientTimeStart != "" {
		dbModel = dbModel.Where("client_time >= ?", listReq.ClientTimeStart)
	}
	if listReq.ClientTimeEnd != "" {
		dbModel = dbModel.Where("client_time <= ?", listReq.ClientTimeEnd)
	}
	return dbModel
}

// List 监控-客户端信息列表
func (service monitorClientService) List(page request.PageReq, listReq MonitorClientListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dbModel := service.GetModel(listReq)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.MonitorClient
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorClientResp{}
	response.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控-客户端信息列表
func (service monitorClientService) ListAll(listReq MonitorClientListReq) (res []MonitorClientResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorClient

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	response.Copy(&res, modelList)
	return res, nil
}

// Detail 监控-客户端信息详情
func (service monitorClientService) Detail(id int) (res MonitorClientResp, e error) {
	var obj model.MonitorClient
	err := service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "获取详情失败"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 监控-客户端信息新增
func (service monitorClientService) Add(addReq MonitorClientAddReq) (e error) {
	var obj model.MonitorClient
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return e
	}
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 监控-客户端信息编辑
func (service monitorClientService) Edit(editReq MonitorClientEditReq) (e error) {
	var obj model.MonitorClient
	err := service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Select("*").Updates(obj).Error
	e = response.CheckErr(err, "更新失败")
	return
}

// Del 监控-客户端信息删除
func (service monitorClientService) Del(id int) (e error) {
	var obj model.MonitorClient
	err := service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询数据失败"); e != nil {
		return
	}
	// 删除
	err = service.db.Delete(&obj).Error
	e = response.CheckErr(err, "删除失败")
	return
}

// ExportFile 监控-客户端信息导出
func (service monitorClientService) ExportFile(listReq MonitorClientListReq) (res []MonitorClientResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorClient
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorClientResp{}
	response.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorClientService) ImportFile(importReq []MonitorClientResp) (e error) {
	var importData []model.MonitorClient
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
