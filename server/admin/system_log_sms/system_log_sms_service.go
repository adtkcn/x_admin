package system_log_sms

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/util"

	"github.com/duke-git/lancet/v2/convertor"
	"gorm.io/gorm"
)

var SystemLogSmsService = NewSystemLogSmsService()
var cacheUtil = util.CacheUtil{
	Name: SystemLogSmsService.Name,
}

// NewSystemLogSmsService 初始化
func NewSystemLogSmsService() *systemLogSmsService {
	return &systemLogSmsService{
		db:   core.GetDB(),
		Name: "systemLogSms",
	}
}

// systemLogSmsService 系统短信日志服务实现类
type systemLogSmsService struct {
	db   *gorm.DB
	Name string
}

// List 系统短信日志列表
func (service systemLogSmsService) GetModel(listReq SystemLogSmsListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.SystemLogSms{})
	if listReq.Scene != nil {
		dbModel = dbModel.Where("scene = ?", *listReq.Scene)
	}
	if listReq.Mobile != nil {
		dbModel = dbModel.Where("mobile like ?", "%"+*listReq.Mobile+"%")
	}
	if listReq.Content != nil {
		dbModel = dbModel.Where("content = ?", *listReq.Content)
	}
	if listReq.Status != nil {
		dbModel = dbModel.Where("status = ?", *listReq.Status)
	}
	if listReq.Results != nil {
		dbModel = dbModel.Where("results = ?", *listReq.Results)
	}
	if listReq.SendTime != nil {
		dbModel = dbModel.Where("send_time = ?", *listReq.SendTime)
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
	return dbModel
}

// List 系统短信日志列表
func (service systemLogSmsService) List(page request.PageReq, listReq SystemLogSmsListReq) (res response.PageResp, e error) {
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
	var modelList []model.SystemLogSms
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []SystemLogSmsResp{}
	response.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 系统短信日志列表
func (service systemLogSmsService) ListAll(listReq SystemLogSmsListReq) (res []SystemLogSmsResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.SystemLogSms

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	response.Copy(&res, modelList)
	return res, nil
}

// Detail 系统短信日志详情
func (service systemLogSmsService) Detail(Id int) (res SystemLogSmsResp, e error) {
	var obj = model.SystemLogSms{}
	err := cacheUtil.GetCache(Id, &obj)
	if err != nil {
		err := service.db.Where("id = ?", Id).Limit(1).First(&obj).Error
		if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
			return
		}
		if e = response.CheckErr(err, "获取详情失败"); e != nil {
			return
		}
		response.Copy(&res, obj)
		cacheUtil.SetCache(obj.Id, obj)
	}

	return
}

// Add 系统短信日志新增
func (service systemLogSmsService) Add(addReq SystemLogSmsAddReq) (createId int, e error) {
	var obj model.SystemLogSms

	// 无法自动转换不同类型
	// response.Copy(&obj, addReq)

	// addInfo, err := convertor.StructToMap(addReq)
	// if err != nil {
	// 	return 0, err
	// }
	// // map弱类型转换到结构体
	// err = mapstructure.WeakDecode(addInfo, &obj)
	err := util.ConvertUtil.StructToStruct(&addReq, &obj)

	if err != nil {
		return 0, err
	}
	err = service.db.Create(&obj).Error
	e = response.CheckMysqlErr(err)
	if e != nil {
		return 0, e
	}
	cacheUtil.SetCache(obj.Id, obj)
	createId = obj.Id
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 系统短信日志编辑
func (service systemLogSmsService) Edit(editReq SystemLogSmsEditReq) (e error) {
	var obj model.SystemLogSms
	err := service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	// 更新
	// response.Copy(&obj, editReq)
	//
	editInfo, err := convertor.StructToMap(editReq)
	if err != nil {
		return err
	}

	err = service.db.Model(&obj).Updates(editInfo).Error
	if e = response.CheckErr(err, "编辑失败"); e != nil {
		return
	}
	cacheUtil.RemoveCache(obj.Id)
	service.Detail(obj.Id)
	return
}

// Del 系统短信日志删除
func (service systemLogSmsService) Del(Id int) (e error) {
	var obj model.SystemLogSms
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

// ExportFile 系统短信日志导出
func (service systemLogSmsService) ExportFile(listReq SystemLogSmsListReq) (res []SystemLogSmsResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.SystemLogSms
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []SystemLogSmsResp{}
	response.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service systemLogSmsService) ImportFile(importReq []SystemLogSmsResp) (e error) {
	var importData []model.SystemLogSms
	response.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
