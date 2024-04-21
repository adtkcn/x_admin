package monitor_client

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

type IMonitorClientService interface {
	List(page request.PageReq, listReq MonitorClientListReq) (res response.PageResp, e error)
	ListAll() (res []MonitorClientResp, e error)

	Detail(id int) (res MonitorClientResp, e error)
	Add(addReq MonitorClientAddReq) (e error)
	Edit(editReq MonitorClientEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewMonitorClientService()

// NewMonitorClientService 初始化
func NewMonitorClientService() *monitorClientService {
	db := core.GetDB()
	return &monitorClientService{db: db}
}

// monitorClientService 客户端信息服务实现类
type monitorClientService struct {
	db *gorm.DB
}

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
	if len(listReq.ClientTime) == 2 {
		// dbModel = dbModel.Where("client_time = ?", listReq.ClientTime)
		dbModel = dbModel.Where("client_time >= ?", listReq.ClientTime[0]).Where("client_time <= ?", listReq.ClientTime[1])
	}
	return dbModel
}

// List 客户端信息列表
func (service monitorClientService) List(page request.PageReq, listReq MonitorClientListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	dbModel := service.GetModel(listReq)
	// 查询
	// dbModel := service.db.Model(&model.MonitorClient{})
	// if listReq.ProjectKey != "" {
	// 	dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	// }
	// if listReq.ClientId != "" {
	// 	dbModel = dbModel.Where("client_id = ?", listReq.ClientId)
	// }
	// if listReq.UserId != "" {
	// 	dbModel = dbModel.Where("user_id = ?", listReq.UserId)
	// }
	// if listReq.Os != "" {
	// 	dbModel = dbModel.Where("os = ?", listReq.Os)
	// }
	// if listReq.Browser != "" {
	// 	dbModel = dbModel.Where("browser = ?", listReq.Browser)
	// }
	// if listReq.City != "" {
	// 	dbModel = dbModel.Where("city = ?", listReq.City)
	// }
	// if listReq.Width > 0 {
	// 	dbModel = dbModel.Where("width = ?", listReq.Width)
	// }
	// if listReq.Height > 0 {
	// 	dbModel = dbModel.Where("height = ?", listReq.Height)
	// }
	// if listReq.Ua != "" {
	// 	dbModel = dbModel.Where("ua = ?", listReq.Ua)
	// }
	// if listReq.ClientTime > 0 {
	// 	dbModel = dbModel.Where("client_time = ?", listReq.ClientTime)
	// }
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.MonitorClient
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []MonitorClientResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 客户端信息列表
func (service monitorClientService) ListAll() (res []MonitorClientResp, e error) {
	var objs []model.MonitorClient

	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 客户端信息详情
func (service monitorClientService) Detail(id int) (res MonitorClientResp, e error) {
	var obj model.MonitorClient
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

// Add 客户端信息新增
func (service monitorClientService) Add(addReq MonitorClientAddReq) (e error) {
	var obj model.MonitorClient
	response.Copy(&obj, addReq)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return e
	}
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 客户端信息编辑
func (service monitorClientService) Edit(editReq MonitorClientEditReq) (e error) {
	var obj model.MonitorClient
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

// Del 客户端信息删除
func (service monitorClientService) Del(id int) (e error) {
	var obj model.MonitorClient
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

// ExportFile 客户端信息导出
func (service monitorClientService) ExportFile(listReq MonitorClientListReq) (res []MonitorClientResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// dbModel := service.db.Model(&model.MonitorClient{})
	// if listReq.ProjectKey != "" {
	// 	dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	// }
	// if listReq.ClientId != "" {
	// 	dbModel = dbModel.Where("client_id = ?", listReq.ClientId)
	// }
	// if listReq.UserId != "" {
	// 	dbModel = dbModel.Where("user_id = ?", listReq.UserId)
	// }
	// if listReq.Os != "" {
	// 	dbModel = dbModel.Where("os = ?", listReq.Os)
	// }
	// if listReq.Browser != "" {
	// 	dbModel = dbModel.Where("browser = ?", listReq.Browser)
	// }
	// if listReq.City != "" {
	// 	dbModel = dbModel.Where("city = ?", listReq.City)
	// }
	// if listReq.Width > 0 {
	// 	dbModel = dbModel.Where("width = ?", listReq.Width)
	// }
	// if listReq.Height > 0 {
	// 	dbModel = dbModel.Where("height = ?", listReq.Height)
	// }
	// if listReq.Ua != "" {
	// 	dbModel = dbModel.Where("ua = ?", listReq.Ua)
	// }
	// if listReq.ClientTime > 0 {
	// 	dbModel = dbModel.Where("client_time = ?", listReq.ClientTime)
	// }

	// 数据
	var objs []model.MonitorClient
	err := dbModel.Order("id asc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []MonitorClientResp{}
	response.Copy(&resps, objs)
	return resps, nil
}

// 导入
func (service monitorClientService) ImportFile(importReq []MonitorClientResp) (e error) {
	var importData []model.MonitorClient
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "Add Create err")
	return e
}
