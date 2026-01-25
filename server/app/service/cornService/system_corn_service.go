package cornService

import (
	"errors"
	"x_admin/app/schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"gorm.io/gorm"
)

var SystemCornService = NewSystemCornService()

// NewSystemCornService 初始化
func NewSystemCornService() *systemCornService {
	return &systemCornService{
		db: core.GetDB(),
		CacheUtil: util.CacheUtil{
			Name: "systemCorn",
		},
	}
}

// systemCornService 定时任务服务实现类
type systemCornService struct {
	db        *gorm.DB
	CacheUtil util.CacheUtil
}

// List 定时任务列表
func (service systemCornService) GetModel(listReq schema.SystemCornListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.SystemCorn{}).Joins("CreatedByUser")
	tableName := core.DBTableName(&model.SystemCorn{})
	if listReq.TaskName.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".task_name like ?", "%"+listReq.TaskName.ValueOrZero()+"%")
	}

	if listReq.TaskCode.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".task_code = ?", listReq.TaskCode.ValueOrZero())
	}
	if listReq.CornExpr.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".corn_expr = ?", listReq.CornExpr.ValueOrZero())
	}
	if listReq.Status.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".Status = ?", listReq.Status.ValueOrZero())
	}
	if listReq.CreatedBy.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".created_by = ?", listReq.CreatedBy.ValueOrZero())
	}
	if listReq.Nickname.IsExistsAndNotNull() {
		dbModel = dbModel.Where("CreatedByUser.nickname like ?", "%"+listReq.Nickname.ValueOrZero()+"%")
	}

	if listReq.CreateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".create_time >= ?", listReq.CreateTimeStart.ValueOrZero())
	}
	if listReq.CreateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".create_time <= ?", listReq.CreateTimeEnd.ValueOrZero())
	}
	if listReq.UpdateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".update_time >= ?", listReq.UpdateTimeStart.ValueOrZero())
	}
	if listReq.UpdateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where(tableName+".update_time <= ?", listReq.UpdateTimeEnd.ValueOrZero())
	}
	// dbModel = dbModel.Where("is_delete = ?", 0)
	return dbModel
}

// 获取更新map
func (service systemCornService) GetUpdateMap(editReq schema.SystemCornEditReq) map[string]interface{} {
	updateMap := make(map[string]interface{})
	if editReq.Id != "" {
		updateMap["id"] = editReq.Id
	}
	if editReq.TaskName.IsExists() {
		updateMap["task_name"] = editReq.TaskName.GetValue()
	}
	if editReq.TaskCode.IsExists() {
		updateMap["task_code"] = editReq.TaskCode.GetValue()
	}
	if editReq.CornExpr.IsExists() {
		updateMap["corn_expr"] = editReq.CornExpr.GetValue()
	}
	if editReq.Status.IsExists() {
		updateMap["Status"] = editReq.Status.GetValue()
	}
	return updateMap
}

// List 定时任务列表
func (service systemCornService) List(page request.PageReq, listReq schema.SystemCornListReq) (res response.PageResp, e error) {
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
	var modelList []model.SystemCorn
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []schema.SystemCornResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 定时任务列表
func (service systemCornService) ListAll(listReq schema.SystemCornListReq) (res []schema.SystemCornResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.SystemCorn

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 定时任务详情
func (service systemCornService) Detail(Id string) (res schema.SystemCornResp, e error) {
	var obj = model.SystemCorn{}
	err := service.CacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ? AND is_delete = ?", Id, 0).Preload("CreatedByUser").Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		service.CacheUtil.SetCache(obj.Id, obj)
	}

	convert_util.Copy(&res, obj)
	return
}

// Add 定时任务新增
func (service systemCornService) Add(addReq schema.SystemCornAddReq, adminId string) (createId string, e error) {
	var obj model.SystemCorn
	convert_util.Copy(&obj, addReq)
	obj.CreatedBy.SetValue(adminId)
	err := service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return "", e
	}
	service.CacheUtil.SetCache(obj.Id, obj)
	createId = obj.Id
	return
}

// Edit 定时任务编辑
func (service systemCornService) Edit(editReq schema.SystemCornEditReq) (e error) {
	var obj model.SystemCorn
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	// convert_util.Copy(&obj, editReq)
	updateMap := service.GetUpdateMap(editReq)
	if len(updateMap) == 0 {
		return errors.New("没有可更新的字段")
	}

	err = service.db.Model(&obj).Updates(updateMap).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	service.CacheUtil.RemoveCache(obj.Id)
	// service.Detail(obj.Id)
	return
}

// Del 定时任务删除
func (service systemCornService) Del(Id string) (e error) {
	var obj model.SystemCorn
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
	service.CacheUtil.RemoveCache(obj.Id)
	return
}

// DelBatch 用户协议-批量删除
func (service systemCornService) DelBatch(Ids []string) (e error) {
	var obj model.SystemCorn
	err := service.db.Where("id in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	service.CacheUtil.RemoveCache(Ids)
	return nil
}

// 获取Excel的列
func (service systemCornService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "任务名称", Key: "TaskName", Width: 15, Decode: core.DecodeString},
		{Name: "任务编码", Key: "TaskCode", Width: 15, Decode: core.DecodeString},
		{Name: "corn表达式", Key: "CornExpr", Width: 15, Decode: core.DecodeString},
		{Name: "禁用", Key: "Status", Width: 15, Decode: core.DecodeInt},
		{Name: "创建人", Key: "CreatedBy", Width: 15, Decode: core.DecodeString},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 定时任务导出
func (service systemCornService) ExportFile(listReq schema.SystemCornListReq) (res []schema.SystemCornResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.SystemCorn
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []schema.SystemCornResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service systemCornService) ImportFile(importReq []schema.SystemCornResp) (e error) {
	var importData []model.SystemCorn
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}

// 获取任务列表
func (service systemCornService) GetTaskList() (list []map[string]interface{}) {
	// var list []map[string]interface{}
	for _, task := range TaskInfoList {
		list = append(list, map[string]interface{}{
			"Lock":     task.Lock,
			"LockTTL":  task.LockTTL.Seconds(),
			"TaskCode": task.TaskCode,
			"TaskDesc": task.TaskDesc,
		})
	}
	return list
}
