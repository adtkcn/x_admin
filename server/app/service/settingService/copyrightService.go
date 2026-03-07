package settingService

import (
	"x_admin/app/schema/settingSchema"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util"

	"gorm.io/gorm"
)

type ISettingCopyrightService interface {
	Detail() (res []map[string]any, e error)
	Save(cReqs []settingSchema.SettingCopyrightItemReq) (e error)
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
	e = response.CheckErr(util.ToolsUtil.JsonToObj(data, &res), "Detail JsonToObj err")
	return
}

// Save 保存网站备案信息
func (cSrv settingCopyrightService) Save(cReqs []settingSchema.SettingCopyrightItemReq) (e error) {
	json, err := util.ToolsUtil.ObjToJson(cReqs)
	if e = response.CheckErr(err, "Save ObjToJson err"); e != nil {
		return
	}
	err = SystemConfigService.Set(cSrv.db, "website", "copyright", json)
	e = response.CheckErr(err, "保存失败")
	return
}
