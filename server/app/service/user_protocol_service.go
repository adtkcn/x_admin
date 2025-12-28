package service

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

var UserProtocolService = NewUserProtocolService()

// NewUserProtocolService 初始化
func NewUserProtocolService() *userProtocolService {
	return &userProtocolService{
		db: core.GetDB(),
		CacheUtil: util.CacheUtil{
			Name: "userProtocol",
		},
	}
}

// userProtocolService 用户协议服务实现类
type userProtocolService struct {
	db *gorm.DB

	CacheUtil util.CacheUtil
}

// List 用户协议列表
func (service userProtocolService) GetModel(listReq schema.UserProtocolListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.UserProtocol{}).Preload("CreatedByUser")
	if listReq.Title.GetValue() != nil {
		dbModel = dbModel.Where("title like ?", "%"+*listReq.Title.GetValue()+"%")
	}
	if listReq.Content.GetValue() != nil {
		dbModel = dbModel.Where("content = ?", *listReq.Content.GetValue())
	}
	if listReq.Version.GetValue() != nil {
		dbModel = dbModel.Where("version = ?", *listReq.Version.GetValue())
	}

	if listReq.CreateTimeStart.GetValue() != nil {
		dbModel = dbModel.Where("create_time >= ?", *listReq.CreateTimeStart.GetValue())
	}
	if listReq.CreateTimeEnd.GetValue() != nil {
		dbModel = dbModel.Where("create_time <= ?", *listReq.CreateTimeEnd.GetValue())
	}
	if listReq.UpdateTimeStart.GetValue() != nil {
		dbModel = dbModel.Where("update_time >= ?", *listReq.UpdateTimeStart.GetValue())
	}
	if listReq.UpdateTimeEnd.GetValue() != nil {
		dbModel = dbModel.Where("update_time <= ?", *listReq.UpdateTimeEnd.GetValue())
	}
	dbModel = dbModel.Where("is_delete = ?", 0)
	return dbModel
}

// 获取更新map,原因是可以如果字段为空,则不更新
func (service userProtocolService) GetUpdateMap(editReq schema.UserProtocolEditReq) map[string]interface{} {
	updateMap := make(map[string]interface{})

	if editReq.Tag.IsExists() {
		updateMap["tag"] = editReq.Tag.GetValue()
	}
	if editReq.Title.IsExists() {
		updateMap["title"] = editReq.Title.GetValue()
	}
	if editReq.Content.IsExists() {
		updateMap["content"] = editReq.Content.GetValue()
	}
	if editReq.Version.IsExists() {
		updateMap["version"] = editReq.Version.GetValue()
	}
	return updateMap
}

// List 用户协议列表
func (service userProtocolService) List(page request.PageReq, listReq schema.UserProtocolListReq) (res response.PageResp, e error) {
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
	var modelList []model.UserProtocol
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []schema.UserProtocolResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 用户协议列表
func (service userProtocolService) ListAll(listReq schema.UserProtocolListReq) (res []schema.UserProtocolResp, e error) {
	dbModel := service.GetModel(listReq)

	var modelList []model.UserProtocol

	err := dbModel.Find(&modelList).Error
	if e = response.CheckErr(err, "查询全部失败"); e != nil {
		return
	}
	convert_util.Copy(&res, modelList)
	return res, nil
}

// Detail 用户协议详情
func (service userProtocolService) Detail(Id string) (res schema.UserProtocolResp, e error) {
	var obj = model.UserProtocol{}
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

// Add 用户协议新增
func (service userProtocolService) Add(addReq schema.UserProtocolAddReq, adminId string) (createId string, e error) {
	var obj model.UserProtocol
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

// Edit 用户协议编辑
func (service userProtocolService) Edit(editReq schema.UserProtocolEditReq) (e error) {
	var obj model.UserProtocol
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	// 不使用结构体是因为没法区分前端是否有值,都会清空，如果全传就无所谓
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

// Del 用户协议删除
func (service userProtocolService) Del(Id string) (e error) {
	var obj model.UserProtocol
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
func (service userProtocolService) DelBatch(Ids []string) (e error) {
	var obj model.UserProtocol
	err := service.db.Where("id in (?)", Ids).Delete(&obj).Error
	if err != nil {
		return err
	}
	// 删除缓存
	for _, v := range Ids {
		service.CacheUtil.RemoveCache(v)
	}
	return nil
}

// 获取Excel的列
func (service userProtocolService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "标识", Key: "Tag", Width: 15, Decode: core.DecodeString},
		{Name: "版本", Key: "Version", Width: 15, Decode: core.DecodeInt},
		{Name: "标题", Key: "Title", Width: 15, Decode: core.DecodeString},
		{Name: "协议内容", Key: "Content", Width: 15, Decode: core.DecodeString},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 用户协议导出
func (service userProtocolService) ExportFile(listReq schema.UserProtocolListReq) (res []schema.UserProtocolResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.UserProtocol
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []schema.UserProtocolResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service userProtocolService) ImportFile(importReq []schema.UserProtocolResp) (e error) {
	var importData []model.UserProtocol
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
