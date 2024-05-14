package monitor_project

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"

	"gorm.io/gorm"
)

type IMonitorProjectService interface {
	List(page request.PageReq, listReq MonitorProjectListReq) (res response.PageResp, e error)
	ListAll() (res []MonitorProjectResp, e error)

	Detail(id int) (res MonitorProjectResp, e error)
	Add(addReq MonitorProjectAddReq) (e error)
	Edit(editReq MonitorProjectEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewMonitorProjectService()

// NewMonitorProjectService 初始化
func NewMonitorProjectService() *monitorProjectService {
	db := core.GetDB()
	return &monitorProjectService{db: db}
}

// monitorProjectService 错误项目服务实现类
type monitorProjectService struct {
	db *gorm.DB
}

// List 错误项目列表
func (service monitorProjectService) List(page request.PageReq, listReq MonitorProjectListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.MonitorProject{})
	if listReq.ProjectKey != "" {
		dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	}
	if listReq.ProjectName != "" {
		dbModel = dbModel.Where("project_name like ?", "%"+listReq.ProjectName+"%")
	}
	if listReq.ProjectType != "" {
		dbModel = dbModel.Where("project_type = ?", listReq.ProjectType)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var objs []model.MonitorProject
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	resps := []MonitorProjectResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 错误项目列表
func (service monitorProjectService) ListAll() (res []MonitorProjectResp, e error) {
	var objs []model.MonitorProject

	err := service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 错误项目详情
func (service monitorProjectService) Detail(id int) (res MonitorProjectResp, e error) {
	var obj model.MonitorProject
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 错误项目新增
func (service monitorProjectService) Add(addReq MonitorProjectAddReq) (e error) {
	var obj model.MonitorProject
	response.Copy(&obj, addReq)
	obj.ProjectKey = util.ToolsUtil.MakeUuid()
	err := service.db.Create(&obj).Error

	if e = response.CheckMysqlErr(err); e != nil {
		return e
	}
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 错误项目编辑
func (service monitorProjectService) Edit(editReq MonitorProjectEditReq) (e error) {
	var obj model.MonitorProject
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
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

// Del 错误项目删除
func (service monitorProjectService) Del(id int) (e error) {
	var obj model.MonitorProject
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	obj.IsDelete = 1
	err = service.db.Save(&obj).Error
	e = response.CheckErr(err, "Del Save err")
	return
}

// ExportFile 错误项目导出
func (service monitorProjectService) ExportFile(listReq MonitorProjectListReq) (res []MonitorProjectResp, e error) {
	// 查询
	dbModel := service.db.Model(&model.MonitorProject{})
	if listReq.ProjectKey != "" {
		dbModel = dbModel.Where("project_key = ?", listReq.ProjectKey)
	}
	if listReq.ProjectName != "" {
		dbModel = dbModel.Where("project_name like ?", "%"+listReq.ProjectName+"%")
	}
	if listReq.ProjectType != "" {
		dbModel = dbModel.Where("project_type = ?", listReq.ProjectType)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)

	// 数据
	var objs []model.MonitorProject
	err := dbModel.Order("id asc").Find(&objs).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	resps := []MonitorProjectResp{}
	response.Copy(&resps, objs)
	return resps, nil
}

// 导入
func (service monitorProjectService) ImportFile(importReq []MonitorProjectResp) (e error) {
	var importData []model.MonitorProject
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "Add Create err")
	return e
}
