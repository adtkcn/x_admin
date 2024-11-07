package monitor_slow

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

var MonitorSlowService = NewMonitorSlowService()
var cacheUtil = util.CacheUtil{
	Name: MonitorSlowService.Name,
}

// NewMonitorSlowService 初始化
func NewMonitorSlowService() *monitorSlowService {
	return &monitorSlowService{
		db:   core.GetDB(),
		Name: "monitorSlow",
	}
}

// monitorSlowService 监控-错误列服务实现类
type monitorSlowService struct {
	db   *gorm.DB
	Name string
}

// List 监控-错误列列表
func (service monitorSlowService) GetModel(listReq MonitorSlowListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.MonitorSlow{})
	if listReq.ProjectKey != nil {
		dbModel = dbModel.Where("project_key = ?", *listReq.ProjectKey)
	}
	if listReq.ClientId != nil {
		dbModel = dbModel.Where("client_id = ?", *listReq.ClientId)
	}
	if listReq.UserId != nil {
		dbModel = dbModel.Where("user_id = ?", *listReq.UserId)
	}
	if listReq.Path != nil {
		dbModel = dbModel.Where("path = ?", *listReq.Path)
	}
	if listReq.Time != nil {
		dbModel = dbModel.Where("time = ?", *listReq.Time)
	}
	if listReq.CreateTimeStart != nil {
		dbModel = dbModel.Where("create_time >= ?", *listReq.CreateTimeStart)
	}
	if listReq.CreateTimeEnd != nil {
		dbModel = dbModel.Where("create_time <= ?", *listReq.CreateTimeEnd)
	}
	return dbModel
}

// List 监控-错误列列表
func (service monitorSlowService) List(page request.PageReq, listReq MonitorSlowListReq) (res response.PageResp, e error) {
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
	var modelList []model.MonitorSlow
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorSlowResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 监控-错误列列表
func (service monitorSlowService) ListAll(listReq MonitorSlowListReq) (res []MonitorSlowResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.MonitorSlow

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 监控-错误列详情
func (service monitorSlowService) Detail(Id int) (res MonitorSlowResp, e error) {
	var obj = model.MonitorSlow{}
	err := cacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ?", Id).Limit(1).First(&obj).Error
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

// Add 监控-错误列新增
func (service monitorSlowService) Add(addReq MonitorSlowAddReq) (createId int, e error) {
	var obj model.MonitorSlow
	convert_util.StructToStruct(addReq, &obj)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0, e
	}
	cacheUtil.SetCache(obj.Id, obj)
	createId = obj.Id
	return
}

// Edit 监控-错误列编辑
func (service monitorSlowService) Edit(editReq MonitorSlowEditReq) (e error) {
	var obj model.MonitorSlow
	err := service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
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

// Del 监控-错误列删除
func (service monitorSlowService) Del(Id int) (e error) {
	var obj model.MonitorSlow
	err := service.db.Where("id = ?", Id).Limit(1).First(&obj).Error
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
	cacheUtil.RemoveCache(obj.Id)
	return
}

// DelBatch 用户协议-批量删除
func (service monitorSlowService) DelBatch(Ids []string) (e error) {
	var obj model.MonitorSlow
	err := service.db.Where("id in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	cacheUtil.RemoveCache(Ids)
	return nil
}

// 获取Excel的列
func (service monitorSlowService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "项目key", Key: "ProjectKey", Width: 15},
		{Name: "sdk生成的客户端id", Key: "ClientId", Width: 15},
		{Name: "用户id", Key: "UserId", Width: 15},
		{Name: "URL地址", Key: "Path", Width: 15},
		{Name: "时间", Key: "Time", Width: 15},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 监控-错误列导出
func (service monitorSlowService) ExportFile(listReq MonitorSlowListReq) (res []MonitorSlowResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.MonitorSlow
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []MonitorSlowResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service monitorSlowService) ImportFile(importReq []MonitorSlowResp) (e error) {
	var importData []model.MonitorSlow
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
