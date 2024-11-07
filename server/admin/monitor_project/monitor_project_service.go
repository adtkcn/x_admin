package monitor_project

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"gorm.io/gorm"
)

var MonitorProjectService = NewMonitorProjectService()
var cacheUtil = util.CacheUtil{
	Name: MonitorProjectService.Name,
}

// NewMonitorProjectService 初始化
func NewMonitorProjectService() *monitorProjectService {
	return &monitorProjectService{
		db:   core.GetDB(),
		Name: "monitorProject",
	}
}

// monitorProjectService 监控项目服务实现类
type monitorProjectService struct {
	db   *gorm.DB
	Name string
}

// List 监控项目列表
func (service monitorProjectService) GetModel(listReq MonitorProjectListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorProject{})
	if listReq.ProjectKey != nil {
		dbModel = dbModel.Where("project_key = ?", *listReq.ProjectKey)
	}
	if listReq.ProjectName != nil {
		dbModel = dbModel.Where("project_name like ?", "%"+*listReq.ProjectName+"%")
	}
	if listReq.ProjectType != nil {
		dbModel = dbModel.Where("project_type = ?", *listReq.ProjectType)
	}
	if listReq.Status != nil {
		dbModel = dbModel.Where("status = ?", *listReq.Status)
	}
	if listReq.CreateTimeStart != nil {
		dbModel = dbModel.Where("create_time >= ?", *listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != nil {
		dbModel = dbModel.Where("create_time <= ?", *listReq.CreateTimeEnd)
	}
	if listReq.UpdateTimeStart != nil {
		dbModel = dbModel.Where("update_time >= ?", *listReq.UpdateTimeStart)
	}
	if listReq.UpdateTimeEnd != nil {
		dbModel = dbModel.Where("update_time <= ?", *listReq.UpdateTimeEnd)
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	return dbModel
}

// List 监控项目列表
func (service monitorProjectService) List(page request.PageReq, listReq MonitorProjectListReq) (res response.PageResp, e error) {
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
	var modelList []model.MonitorProject
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorProjectResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控项目列表
func (service monitorProjectService) ListAll(listReq MonitorProjectListReq) (res []MonitorProjectResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorProject

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 监控项目详情
func (service monitorProjectService) Detail(Id int) (res MonitorProjectResp, e error) {
	var obj = model.MonitorProject{}
	err := cacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ? AND is_delete = ?", Id, 0).Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		cacheUtil.SetCache(obj.Id, obj)
	}

	convert_util.Copy(&res, obj)
	return
}

// Add 监控项目新增
func (service monitorProjectService) Add(addReq MonitorProjectAddReq) (createId int, e error) {
	var obj model.MonitorProject
	convert_util.StructToStruct(addReq, &obj)
	obj.ProjectKey = util.ToolsUtil.MakeUuid()
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0, e
	}
	cacheUtil.SetCache(obj.Id, obj)
	createId = obj.Id
	return
}

// Edit 监控项目编辑
func (service monitorProjectService) Edit(editReq MonitorProjectEditReq) (e error) {
	var obj model.MonitorProject
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	convert_util.Copy(&obj, editReq)

	err = service.db.Model(&obj).Select("*").Updates(obj).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	cacheUtil.RemoveCache(obj.Id)
	service.Detail(obj.Id)
	return
}

// Del 监控项目删除
func (service monitorProjectService) Del(Id int) (e error) {
	var obj model.MonitorProject
	err := service.db.Where("id = ? AND is_delete = ?", Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询数据失败"); e != nil {
		return
	}
	// 删除
	obj.IsDelete = 1
	obj.DeleteTime = util.NullTimeUtil.Now()
	err = service.db.Save(&obj).Error
	e = response.CheckErr(err, "删除失败")
	cacheUtil.RemoveCache(obj.Id)
	return
}

// DelBatch 用户协议-批量删除
func (service monitorProjectService) DelBatch(Ids []string) (e error) {
	var obj model.MonitorProject
	err := service.db.Where("id in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	// for _, v := range Ids {
	// 	cacheUtil.RemoveCache(v)
	// }
	cacheUtil.RemoveCache(Ids)
	return nil
}

// 获取Excel的列
func (service monitorProjectService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "项目uuid", Key: "ProjectKey", Width: 15},
		{Name: "项目名称", Key: "ProjectName", Width: 15},
		{Name: "项目类型go java web node php 等", Key: "ProjectType", Width: 15},
		{Name: "是否启用: 0=否, 1=是", Key: "Status", Width: 15, Decode: core.DecodeInt},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 监控项目导出
func (service monitorProjectService) ExportFile(listReq MonitorProjectListReq) (res []MonitorProjectResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorProject
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorProjectResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorProjectService) ImportFile(importReq []MonitorProjectResp) (e error) {
	var importData []model.MonitorProject
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
