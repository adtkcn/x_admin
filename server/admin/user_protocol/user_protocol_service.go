package user_protocol

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

var UserProtocolService = NewUserProtocolService()
var cacheUtil = util.CacheUtil{
	Name: UserProtocolService.Name,
}

// NewUserProtocolService 初始化
func NewUserProtocolService() *userProtocolService {
	return &userProtocolService{
		db:   core.GetDB(),
		Name: "userProtocol",
	}
}

// userProtocolService 用户协议服务实现类
type userProtocolService struct {
	db   *gorm.DB
	Name string
}

// List 用户协议列表
func (service userProtocolService) GetModel(listReq UserProtocolListReq) *gorm.DB {
	// 查询
	dbModel := service.db.Model(&model.UserProtocol{})
	if listReq.Title != nil {
		dbModel = dbModel.Where("title like ?", "%"+*listReq.Title+"%")
	}
	if listReq.Content != nil {
		dbModel = dbModel.Where("content = ?", *listReq.Content)
	}
	if listReq.Sort != nil {
		dbModel = dbModel.Where("sort = ?", *listReq.Sort)
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

// List 用户协议列表
func (service userProtocolService) List(page request.PageReq, listReq UserProtocolListReq) (res response.PageResp, e error) {
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
	result := []UserProtocolResp{}
	convert_util.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// ListAll 用户协议列表
func (service userProtocolService) ListAll(listReq UserProtocolListReq) (res []UserProtocolResp, e error) {
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
func (service userProtocolService) Detail(Id int) (res UserProtocolResp, e error) {
	var obj = model.UserProtocol{}
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

// Add 用户协议新增
func (service userProtocolService) Add(addReq UserProtocolAddReq) (createId int, e error) {
	var obj model.UserProtocol
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

// Edit 用户协议编辑
func (service userProtocolService) Edit(editReq UserProtocolEditReq) (e error) {
	var obj model.UserProtocol
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

// Del 用户协议删除
func (service userProtocolService) Del(Id int) (e error) {
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
	cacheUtil.RemoveCache(obj.Id)
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
		cacheUtil.RemoveCache(v)
	}
	return nil
}

// 获取Excel的列
func (service userProtocolService) GetExcelCol() []excel2.Col {
	var cols = []excel2.Col{
		{Name: "标题", Key: "Title", Width: 15},
		{Name: "协议内容", Key: "Content", Width: 15},
		{Name: "排序", Key: "Sort", Width: 15, Decode: core.DecodeFloat},
		{Name: "创建时间", Key: "CreateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
		{Name: "更新时间", Key: "UpdateTime", Width: 15, Decode: util.NullTimeUtil.DecodeTime},
	}
	// 还可以考虑字典，请求下来加上 Replace 实现替换导出
	return cols
}

// ExportFile 用户协议导出
func (service userProtocolService) ExportFile(listReq UserProtocolListReq) (res []UserProtocolResp, e error) {
	// 查询
	dbModel := service.GetModel(listReq)

	// 数据
	var modelList []model.UserProtocol
	err := dbModel.Order("id asc").Find(&modelList).Error
	if e = response.CheckErr(err, "查询失败"); e != nil {
		return
	}
	result := []UserProtocolResp{}
	convert_util.Copy(&result, modelList)
	return result, nil
}

// 导入
func (service userProtocolService) ImportFile(importReq []UserProtocolResp) (e error) {
	var importData []model.UserProtocol
	convert_util.Copy(&importData, importReq)
	err := service.db.Create(&importData).Error
	e = response.CheckErr(err, "添加失败")
	return e
}
