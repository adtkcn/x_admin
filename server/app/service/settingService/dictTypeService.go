package settingService

import (
	"errors"
	"x_admin/app/schema/settingSchema"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model/setting_model"

	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

type ISettingDictTypeService interface {
	All() (res []settingSchema.SettingDictTypeResp, e error)
	List(page request.PageReq, listReq settingSchema.SettingDictTypeListReq) (res response.PageResp, e error)
	Detail(id string) (res settingSchema.SettingDictTypeResp, e error)
	Add(addReq settingSchema.SettingDictTypeAddReq) (e error)
	Edit(editReq settingSchema.SettingDictTypeEditReq) (e error)
	Del(delReq settingSchema.SettingDictTypeDelReq) (e error)
}

var DictTypeService = NewSettingDictTypeService()

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
func (dtSrv settingDictTypeService) All() (res []settingSchema.SettingDictTypeResp, e error) {
	var dictTypes []setting_model.DictType
	err := dtSrv.db.Order("id desc").Find(&dictTypes).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []settingSchema.SettingDictTypeResp{}
	convert_util.Copy(&res, dictTypes)
	return
}

// List 字典类型列表
func (dtSrv settingDictTypeService) List(page request.PageReq, listReq settingSchema.SettingDictTypeListReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	dtModel := dtSrv.db.Model(&setting_model.DictType{})
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
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	var dts []setting_model.DictType
	err = dtModel.Limit(limit).Offset(offset).Order("id desc").Find(&dts).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	dtResp := []settingSchema.SettingDictTypeResp{}
	convert_util.Copy(&dtResp, dts)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    dtResp,
	}, nil
}

// Detail 字典类型详情
func (dtSrv settingDictTypeService) Detail(id string) (res settingSchema.SettingDictTypeResp, e error) {
	var dt setting_model.DictType
	err := dtSrv.db.Where("id = ?", id).First(&dt).Error
	if e = response.CheckDBNotRecord(err, "字典类型不存在！"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	convert_util.Copy(&res, dt)
	return
}

// Add 字典类型新增
func (dtSrv settingDictTypeService) Add(addReq settingSchema.SettingDictTypeAddReq) (e error) {
	if r := dtSrv.db.Where("dict_name = ?", addReq.DictName).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典名称已存在！")
	}
	if r := dtSrv.db.Where("dict_type = ?", addReq.DictType).First(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典类型已存在！")
	}
	var dt setting_model.DictType
	convert_util.Copy(&dt, addReq)
	err := dtSrv.db.Create(&dt).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 字典类型编辑
func (dtSrv settingDictTypeService) Edit(editReq settingSchema.SettingDictTypeEditReq) (e error) {
	// 检查字典类型是否存在
	var dt setting_model.DictType
	err := dtSrv.db.Where("id = ?", editReq.ID).First(&dt).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("字典类型不存在")
		}
		return response.CheckErr(err, "查询字典类型失败")
	}

	// 检查名称和类型是否重复
	if r := dtSrv.db.Where("id != ? AND (dict_name = ? OR dict_type = ?)", editReq.ID, editReq.DictName, editReq.DictType).Limit(1).Find(&setting_model.DictType{}); r.RowsAffected > 0 {
		return response.AssertArgumentError.SetMessage("字典名称或类型已存在！")
	}

	convert_util.Copy(&dt, editReq)
	result := dtSrv.db.Model(&dt).Select("*").Updates(dt)
	if result.Error != nil {
		return response.CheckErr(result.Error, "编辑失败")
	}
	return
}

// Del 字典类型删除
func (dtSrv settingDictTypeService) Del(delReq settingSchema.SettingDictTypeDelReq) (e error) {
	if len(delReq.Ids) == 0 {
		return errors.New("删除ID列表不能为空")
	}

	result := dtSrv.db.Where("id in (?)", delReq.Ids).Delete(&setting_model.DictType{})
	if result.Error != nil {
		return response.CheckErr(result.Error, "删除失败")
	}
	if result.RowsAffected == 0 {
		return errors.New("没有找到可删除的记录")
	}
	return
}
