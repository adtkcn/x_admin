package setting_service

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util"

	"gorm.io/gorm"
)

type ISettingCopyrightService interface {
	Detail() (res []map[string]any, e error)
	Save(cReqs []setting_schema.SettingCopyrightItemReq) (e error)
}

var CopyrightService = NewSettingCopyrightService()

// NewSettingCopyrightService 初始化
func NewSettingCopyrightService() *settingCopyrightService {
	db := core.GetDB()
	return &settingCopyrightService{db: db}
}

// settingCopyrightService 网站备案服务实现类
type settingCopyrightService struct {
	db *gorm.DB
}

// Detail 获取网站备案信息
func (cSrv settingCopyrightService) Detail() (res []map[string]any, e error) {
	data, err := SystemConfigService.GetVal(cSrv.db, "website", "copyright", "[]")
	if e = response.CheckErr(err, "Detail GetVal err"); e != nil {
		return
	}
	res, e = util.ToolsUtil.JsonToObj[[]map[string]any](data)
	e = response.CheckErr(e, "Detail JsonToObj err")
	return
}

// Save 保存网站备案信息
func (cSrv settingCopyrightService) Save(cReqs []setting_schema.SettingCopyrightItemReq) (e error) {
	json, err := util.ToolsUtil.ObjToJson(cReqs)
	if e = response.CheckErr(err, "Save ObjToJson err"); e != nil {
		return
	}
	err = SystemConfigService.Set(cSrv.db, "website", "copyright", json)
	e = response.CheckErr(err, "保存失败")
	return
}
