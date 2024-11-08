package dict_data

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/setting_model"
	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

type ISettingDictDataService interface {
	All(allReq SettingDictDataListReq) (res []SettingDictDataResp, e error)
	List(page request.PageReq, listReq SettingDictDataListReq) (res response.PageResp, e error)
	Detail(id uint) (res SettingDictDataResp, e error)
	Add(addReq SettingDictDataAddReq) (e error)
	Edit(editReq SettingDictDataEditReq) (e error)
	Del(delReq SettingDictDataDelReq) (e error)
}

var Service = NewSettingDictDataService()

// NewSettingDictDataService 初始化
func NewSettingDictDataService() ISettingDictDataService {
	db := core.GetDB()
	return &settingDictDataService{db: db}
}

// settingDictDataService 字典数据服务实现类
type settingDictDataService struct {
	db *gorm.DB
}

// All 字典数据所有
func (ddSrv settingDictDataService) All(allReq SettingDictDataListReq) (res []SettingDictDataResp, e error) {
	var dictType setting_model.DictType
	err := ddSrv.db.Where("dict_type = ? AND is_delete = ?", allReq.DictType, 0).Limit(1).First(&dictType).Error
	if e = response.CheckErrDBNotRecord(err, "该字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "All First err"); e != nil {
		return
	}
	ddModel := ddSrv.db.Where("type_id = ? AND is_delete = ?", dictType.ID, 0)
	if allReq.Name != "" {
		ddModel = ddModel.Where("name like ?", "%"+allReq.Name+"%")
	}
	if allReq.Value != "" {
		ddModel = ddModel.Where("value like ?", "%"+allReq.Value+"%")
	}
	if allReq.Status >= 0 {
		ddModel = ddModel.Where("status = ?", allReq.Status)
	}
	var dictDatas []setting_model.DictData
	err = ddModel.Order("id asc").Find(&dictDatas).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []SettingDictDataResp{}
	convert_util.Copy(&res, dictDatas)
	return
}

// List 字典数据列表
func (ddSrv settingDictDataService) List(page request.PageReq, listReq SettingDictDataListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	var dictType setting_model.DictType
	err := ddSrv.db.Where("dict_type = ? AND is_delete = ?", listReq.DictType, 0).Limit(1).First(&dictType).Error
	if e = response.CheckErrDBNotRecord(err, "该字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "List First err"); e != nil {
		return
	}
	ddModel := ddSrv.db.Model(&setting_model.DictData{}).Where("type_id = ? AND is_delete = ?", dictType.ID, 0)
	if listReq.Name != "" {
		ddModel = ddModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Value != "" {
		ddModel = ddModel.Where("value like ?", "%"+listReq.Value+"%")
	}
	if listReq.Status >= 0 {
		ddModel = ddModel.Where("status = ?", listReq.Status)
	}
	var count int64
	e = ddModel.Count(&count).Error
	if e = response.CheckErr(e, "列表总数获取失败"); e != nil {
		return
	}
	var dds []setting_model.DictData
	err = ddModel.Limit(limit).Offset(offset).Order("id asc").Find(&dds).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	dtResp := []SettingDictDataResp{}
	convert_util.Copy(&dtResp, dds)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    dtResp,
	}, nil
}

// Detail 字典数据详情
func (ddSrv settingDictDataService) Detail(id uint) (res SettingDictDataResp, e error) {
	var dd setting_model.DictData
	err := ddSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dd).Error
	if e = response.CheckErrDBNotRecord(err, "字典数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, dd)
	return
}

// Add 字典数据新增
func (ddSrv settingDictDataService) Add(addReq SettingDictDataAddReq) (e error) {
	if r := ddSrv.db.Where("name = ? AND is_delete = ?", addReq.Name, 0).Limit(1).First(&setting_model.DictData{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典数据已存在！")
	}
	var dd setting_model.DictData
	convert_util.Copy(&dd, addReq)
	err := ddSrv.db.Create(&dd).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 字典数据编辑
func (ddSrv settingDictDataService) Edit(editReq SettingDictDataEditReq) (e error) {
	var dd setting_model.DictData
	err := ddSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&dd).Error
	if e = response.CheckErrDBNotRecord(err, "字典数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	if r := ddSrv.db.Where("id != ? AND name = ? AND is_delete = ?", editReq.ID, editReq.Name, 0).Limit(1).First(&setting_model.DictData{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典数据已存在！")
	}

	convert_util.Copy(&dd, editReq)
	err = ddSrv.db.Save(&dd).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 字典数据删除
func (ddSrv settingDictDataService) Del(delReq SettingDictDataDelReq) (e error) {
	err := ddSrv.db.Model(&setting_model.DictData{}).Where("id IN ?", delReq.Ids).Updates(
		setting_model.DictData{IsDelete: 1, DeleteTime: util.NullTimeUtil.Now()}).Error
	return response.CheckErr(err, "Del Update err")
}
