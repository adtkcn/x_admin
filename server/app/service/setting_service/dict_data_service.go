package setting_service

import (
	"x_admin/app/model/setting_model"
	"x_admin/app/schema/setting_schema"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

type ISettingDictDataService interface {
	All(allReq setting_schema.SettingDictDataListReq) (res []setting_schema.SettingDictDataResp, e error)
	// List(page request.PageReq, listReq SettingDictDataListReq) (res response.PageResp, e error)
	Detail(id string) (res setting_schema.SettingDictDataResp, e error)
	Add(addReq setting_schema.SettingDictDataAddReq) (e error)
	Edit(editReq setting_schema.SettingDictDataEditReq) (e error)
	Del(delReq setting_schema.SettingDictDataDelReq) (e error)
}

var DictDataService = NewSettingDictDataService()

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
func (ddSrv settingDictDataService) All(allReq setting_schema.SettingDictDataListReq) (res []setting_schema.SettingDictDataResp, e error) {
	var dictType setting_model.DictType
	err := ddSrv.db.Where("dict_type = ?", allReq.DictType).First(&dictType).Error
	if e = response.CheckDBNotRecord(err, "该字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "All First err"); e != nil {
		return
	}
	ddModel := ddSrv.db.Where("type_id = ?", dictType.ID)
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
	res = []setting_schema.SettingDictDataResp{}
	convert_util.Copy(&res, dictDatas)
	return
}

// Detail 字典数据详情
func (ddSrv settingDictDataService) Detail(id string) (res setting_schema.SettingDictDataResp, e error) {
	var dd setting_model.DictData
	err := ddSrv.db.Where("id = ?", id).First(&dd).Error
	if e = response.CheckDBNotRecord(err, "字典数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, dd)
	return
}

// Add 字典数据新增
func (ddSrv settingDictDataService) Add(addReq setting_schema.SettingDictDataAddReq) (e error) {
	if r := ddSrv.db.Where("name = ?", addReq.Name).First(&setting_model.DictData{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典数据已存在！")
	}
	var dd setting_model.DictData
	convert_util.Copy(&dd, addReq)
	err := ddSrv.db.Create(&dd).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 字典数据编辑
func (ddSrv settingDictDataService) Edit(editReq setting_schema.SettingDictDataEditReq) (e error) {
	var dd setting_model.DictData
	err := ddSrv.db.Where("id = ?", editReq.ID).First(&dd).Error
	if e = response.CheckDBNotRecord(err, "字典数据不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	if r := ddSrv.db.Where("id != ? AND name = ?", editReq.ID, editReq.Name).First(&setting_model.DictData{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典数据已存在！")
	}

	convert_util.Copy(&dd, editReq)
	err = ddSrv.db.Save(&dd).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 字典数据删除
func (ddSrv settingDictDataService) Del(delReq setting_schema.SettingDictDataDelReq) (e error) {
	err := ddSrv.db.Model(&setting_model.DictData{}).Where("id IN ?", delReq.Ids).Updates(
		setting_model.DictData{IsDelete: 1, DeleteTime: util.NullTimeUtil.Now()}).Error
	return response.CheckErr(err, "Del Update err")
}
