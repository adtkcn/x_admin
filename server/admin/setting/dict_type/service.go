package dict_type

import (
	"time"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/setting_model"

	"gorm.io/gorm"
)

type ISettingDictTypeService interface {
	All() (res []SettingDictTypeResp, e error)
	List(page request.PageReq, listReq SettingDictTypeListReq) (res response.PageResp, e error)
	Detail(id uint) (res SettingDictTypeResp, e error)
	Add(addReq SettingDictTypeAddReq) (e error)
	Edit(editReq SettingDictTypeEditReq) (e error)
	Del(delReq SettingDictTypeDelReq) (e error)
}

var Service = NewSettingDictTypeService()

// NewSettingDictTypeService 初始化
func NewSettingDictTypeService() ISettingDictTypeService {
	db := core.GetDB()
	return &settingDictTypeService{db: db}
}

// settingDictTypeService 字典类型服务实现类
type settingDictTypeService struct {
	db *gorm.DB
}

// All 字典类型所有
func (dtSrv settingDictTypeService) All() (res []SettingDictTypeResp, e error) {
	var dictTypes []setting_model.DictType
	err := dtSrv.db.Where("is_delete = ?", 0).Order("id desc").Find(&dictTypes).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []SettingDictTypeResp{}
	response.Copy(&res, dictTypes)
	return
}

// List 字典类型列表
func (dtSrv settingDictTypeService) List(page request.PageReq, listReq SettingDictTypeListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dtModel := dtSrv.db.Model(&setting_model.DictType{}).Where("is_delete = ?", 0)
	if listReq.DictName != "" {
		dtModel = dtModel.Where("dict_name like ?", "%"+listReq.DictName+"%")
	}
	if listReq.DictType != "" {
		dtModel = dtModel.Where("dict_type like ?", "%"+listReq.DictType+"%")
	}
	if listReq.DictStatus >= 0 {
		dtModel = dtModel.Where("dict_status = ?", listReq.DictStatus)
	}
	var count int64
	err := dtModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var dts []setting_model.DictType
	err = dtModel.Limit(limit).Offset(offset).Order("id desc").Find(&dts).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	dtResp := []SettingDictTypeResp{}
	response.Copy(&dtResp, dts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    dtResp,
	}, nil
}

// Detail 字典类型详情
func (dtSrv settingDictTypeService) Detail(id uint) (res SettingDictTypeResp, e error) {
	var dt setting_model.DictType
	err := dtSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&dt).Error
	if e = response.CheckErrDBNotRecord(err, "字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, dt)
	return
}

// Add 字典类型新增
func (dtSrv settingDictTypeService) Add(addReq SettingDictTypeAddReq) (e error) {
	if r := dtSrv.db.Where("dict_name = ? AND is_delete = ?", addReq.DictName, 0).Limit(1).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("字典名称已存在！")
	}
	if r := dtSrv.db.Where("dict_type = ? AND is_delete = ?", addReq.DictType, 0).Limit(1).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("字典类型已存在！")
	}
	var dt setting_model.DictType
	response.Copy(&dt, addReq)
	err := dtSrv.db.Create(&dt).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 字典类型编辑
func (dtSrv settingDictTypeService) Edit(editReq SettingDictTypeEditReq) (e error) {
	err := dtSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&setting_model.DictType{}).Error
	if e = response.CheckErrDBNotRecord(err, "字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	if r := dtSrv.db.Where("id != ? AND dict_name = ? AND is_delete = ?", editReq.ID, editReq.DictName, 0).Limit(1).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("字典名称已存在！")
	}
	if r := dtSrv.db.Where("id != ? AND dict_type = ? AND is_delete = ?", editReq.ID, editReq.DictType, 0).Limit(1).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("字典类型已存在！")
	}
	var dt setting_model.DictType
	response.Copy(&dt, editReq)
	err = dtSrv.db.Save(&dt).Error
	e = response.CheckErr(err, "Edit Save err")
	return
}

// Del 字典类型删除
func (dtSrv settingDictTypeService) Del(delReq SettingDictTypeDelReq) (e error) {
	err := dtSrv.db.Model(&setting_model.DictType{}).Where("id IN ?", delReq.Ids).Updates(
		setting_model.DictType{IsDelete: 1, DeleteTime: time.Now().Unix()}).Error
	return response.CheckErr(err, "Del Update err")
}
