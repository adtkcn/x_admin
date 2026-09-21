package service

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
	if listReq.Title.IsExistsAndNotNull() {
		dbModel = dbModel.Where("title like ?", "%"+listReq.Title.ValueOrZero()+"%")
	}
	if listReq.Content.IsExistsAndNotNull() {
		dbModel = dbModel.Where("content = ?", listReq.Content.ValueOrZero())
	}
	if listReq.Version.IsExistsAndNotNull() {
		dbModel = dbModel.Where("version = ?", listReq.Version.ValueOrZero())
	}

	if listReq.CreateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where("create_time >= ?", listReq.CreateTimeStart.ValueOrZero())
	}
	if listReq.CreateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where("create_time <= ?", listReq.CreateTimeEnd.ValueOrZero())
	}
	if listReq.UpdateTimeStart.IsExistsAndNotNull() {
		dbModel = dbModel.Where("update_time >= ?", listReq.UpdateTimeStart.ValueOrZero())
	}
	if listReq.UpdateTimeEnd.IsExistsAndNotNull() {
		dbModel = dbModel.Where("update_time <= ?", listReq.UpdateTimeEnd.ValueOrZero())
	}
	return dbModel
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
		err := service.db.Where("id = ?", Id).Preload("CreatedByUser").First(&obj).Error
		if e = response.CheckDBErr(err, "数据不存在!", "获取详情失败"); e != nil {
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
	if err != nil {
		return "", err
	}
	service.CacheUtil.SetCache(obj.Id, obj)
	createId = obj.Id
	return
}

// Edit 用户协议编辑
func (service userProtocolService) Edit(editReq schema.UserProtocolEditReq) (e error) {
	result := service.db.Model(model.UserProtocol{}).Where("id = ?", editReq.Id).Updates(editReq)
	// MySQL 错误码（唯一键冲突、外键约束等）由统一出口翻译成业务文案，
	// 未知数据库错误由出口统一记录日志，此处不再重复记录
	if result.Error != nil {
		return result.Error
	}

	service.CacheUtil.RemoveCache(editReq.Id)
	// service.Detail(editReq.Id)
	return
}

// Del 用户协议删除
func (service userProtocolService) Del(Id string) (e error) {
	result := service.db.Where("id = ?", Id).Delete(&model.UserProtocol{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	service.CacheUtil.RemoveCache(Id)

	return
}

// DelBatch 用户协议-批量删除
func (service userProtocolService) DelBatch(Ids []string) (e error) {

	result := service.db.Where("id in (?)", Ids).Delete(&model.UserProtocol{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "批量删除失败")
	}

	if result.RowsAffected == 0 {
		return errors.New("没有找到可删除的记录")
	}

	for _, v := range Ids {
		service.CacheUtil.RemoveCache(v)
	}
	return
}

// 获取Excel的列
func (service userProtocolService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "标识", JsonTag: "tag", Width: 15, Decode: x_null.DecodeString},
		{Name: "版本", JsonTag: "version", Width: 15, Decode: x_null.DecodeInt64},
		{Name: "标题", JsonTag: "title", Width: 15, Decode: x_null.DecodeString},
		{Name: "协议内容", JsonTag: "content", Width: 15, Decode: x_null.DecodeString},
		{Name: "创建人", JsonTag: "created_by", Width: 15, Decode: x_null.DecodeString},
		{Name: "创建时间", JsonTag: "create_time", Width: 15, Decode: x_null.DecodeTime},
		{Name: "更新时间", JsonTag: "update_time", Width: 15, Decode: x_null.DecodeTime},
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
