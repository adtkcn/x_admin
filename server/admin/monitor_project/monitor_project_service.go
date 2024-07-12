package monitor_project

import (
	"errors"
	"strconv"
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

// 设置缓存
func (service monitorProjectService) SetCache(obj model.MonitorProject) bool {
	str, e := util.ToolsUtil.ObjToJson(obj)
	if e != nil {
		return false
	}
	return util.RedisUtil.HSet("MonitorProject", strconv.Itoa(obj.Id), str, 3600)
}

// 获取缓存
func (service monitorProjectService) GetCache(id int) (model.MonitorProject, error) {
	var obj model.MonitorProject
	str := util.RedisUtil.HGet("MonitorProject", strconv.Itoa(id))
	if str == "" {
		return obj, errors.New("获取缓存失败")
	}
	err := util.ToolsUtil.JsonToObj(str, &obj)

	if err != nil {
		return obj, errors.New("获取缓存失败")
	}
	return obj, nil
}

// 删除缓存
func (service monitorProjectService) RemoveCache(obj model.MonitorProject) bool {
	return util.RedisUtil.HDel("MonitorProject", strconv.Itoa(obj.Id))
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
	if listReq.CreateTimeStart != "" {
		dbModel = dbModel.Where("create_time >= ?", listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != "" {
		dbModel = dbModel.Where("create_time <= ?", listReq.CreateTimeEnd)
	}
	if listReq.UpdateTimeStart != "" {
		dbModel = dbModel.Where("update_time >= ?", listReq.UpdateTimeStart)
	}
	if listReq.UpdateTimeEnd != "" {
		dbModel = dbModel.Where("update_time <= ?", listReq.UpdateTimeEnd)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.MonitorProject
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	result := []MonitorProjectResp{}
	response.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 错误项目列表
func (service monitorProjectService) ListAll() (res []MonitorProjectResp, e error) {
	var modelList []model.MonitorProject

	err := service.db.Find(&modelList).Error
	if e = response.CheckErr(err, "获取列表失败"); e != nil {
		return
	}
	response.Copy(&res, modelList)
	return res, nil
}

// Detail 错误项目详情
func (service monitorProjectService) Detail(id int) (res MonitorProjectResp, e error) {

	var obj, err = service.GetCache(id)
	if err != nil {
		err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "详情获取失败"); e != nil {
			return
		}
		service.SetCache(obj)
	}

	response.Copy(&res, obj)
	return
}

// Add 错误项目新增
func (service monitorProjectService) Add(addReq MonitorProjectAddReq) (createId int, e error) {
	var obj model.MonitorProject
	response.Copy(&obj, addReq)
	obj.ProjectKey = util.ToolsUtil.MakeUuid()
	err := service.db.Create(&obj).Error

	if e = response.CheckMysqlErr(err); e != nil {
		return 0, e
	}
	service.SetCache(obj)
	createId = obj.Id
	e = response.CheckErr(err, "添加失败")
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
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)

	err = service.db.Model(&obj).Select("*").Updates(obj).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	service.SetCache(obj)
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
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	// 删除
	obj.IsDelete = 1
	err = service.db.Save(&obj).Error
	service.RemoveCache(obj)
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
	var modelList []model.MonitorProject
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	result := []MonitorProjectResp{}
	response.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorProjectService) ImportFile(importReq []MonitorProjectResp) (e error) {
	var importData []model.MonitorProject
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
