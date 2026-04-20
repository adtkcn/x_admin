package corn_service

import (
	"errors"
	"x_admin/app/model"
	"x_admin/app/schema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"
	"x_admin/util/excel2"

	"github.com/adtkcn/x_null"
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
	tableName := core.DBTableName(&model.SystemCorn{})

	dbModel := service.db.Model(&model.SystemCorn{}).Table(tableName + " as a").Joins("CreatedByUser")

	if listReq.TaskName.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.task_name like ?", "%"+listReq.TaskName.ValueOrZero()+"%")
	}

	if listReq.TaskCode.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.task_code = ?", listReq.TaskCode.ValueOrZero())
	}
	if listReq.CornExpr.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.corn_expr = ?", listReq.CornExpr.ValueOrZero())
	}
	if listReq.Status.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.Status = ?", listReq.Status.ValueOrZero())
	}
	if listReq.CreatedBy.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.created_by = ?", listReq.CreatedBy.ValueOrZero())
	}
	if listReq.Nickname.IsExistsAndNotNull() {
		dbModel = dbModel.Where("CreatedByUser.nickname like ?", "%"+listReq.Nickname.ValueOrZero()+"%")
	}

	if listReq.CreateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.create_time >= ?", listReq.CreateTimeStart.ValueOrZero())
	}
	if listReq.CreateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.create_time <= ?", listReq.CreateTimeEnd.ValueOrZero())
	}
	if listReq.UpdateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.update_time >= ?", listReq.UpdateTimeStart.ValueOrZero())
	}
	if listReq.UpdateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where("a.update_time <= ?", listReq.UpdateTimeEnd.ValueOrZero())
	}

	return dbModel
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
		err := service.db.Where("id = ?", Id).Preload("CreatedByUser").First(&obj).Error
		if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
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

	result := service.db.Model(&model.SystemCorn{}).Where("id = ?", editReq.Id).Updates(editReq)

	if result.Error != nil {
		// 这里处理真正的数据库错误（如连接失败、SQL语法错误、约束冲突等）
		core.Logger.Errorf("数据库错误: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		// 这里处理“找不到数据”的情况
		core.Logger.Errorf("未找到 ID 为 %v 的记录，更新失败", editReq.Id)
		return errors.New("记录不存在")
	}
	service.CacheUtil.RemoveCache(editReq.Id)
	// service.Detail(obj.Id)
	return
}

// Del 定时任务删除
func (service systemCornService) Del(Id string) (e error) {
	var obj model.SystemCorn
	err := service.db.Where("id = ?", Id).First(&obj).Error
	// 校验
	if e = response.CheckDBNotRecord(err, "数据不存在!"); e != nil {
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
	service.CacheUtil.RemoveCache(Ids...)
	return nil
}

// 获取Excel的列
func (service systemCornService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "任务名称", Key: "TaskName", Width: 15, Decode: x_null.DecodeString},
		{Name: "任务编码", Key: "TaskCode", Width: 15, Decode: x_null.DecodeString},
		{Name: "corn表达式", Key: "CornExpr", Width: 15, Decode: x_null.DecodeString},
		{Name: "禁用", Key: "Status", Width: 15, Decode: x_null.DecodeInt64},
		{Name: "创建人", Key: "CreatedBy", Width: 15, Decode: x_null.DecodeString},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: x_null.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: x_null.DecodeTime},
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
func (service systemCornService) GetTaskList() (list []map[string]any) {
	// var list []map[string]any
	for _, task := range TaskInfoList {
		list = append(list, map[string]any{
			"Lock":     task.Lock,
			"LockTTL":  task.LockTTL.Seconds(),
			"TaskCode": task.TaskCode,
			"TaskDesc": task.TaskDesc,
		})
	}
	return list
}
