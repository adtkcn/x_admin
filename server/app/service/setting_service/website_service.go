package setting_service

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/core"
	"x_admin/core/response"

	"gorm.io/gorm"
)

type ISettingWebsiteService interface {
	Detail() (res map[string]string, e error)
	Save(wsReq setting_schema.SettingWebsiteReq) (e error)
}

var WebsiteService = NewSettingWebsiteService()

// NewSettingWebsiteService 初始化
func NewSettingWebsiteService() ISettingWebsiteService {
	db := core.GetDB()
	return &settingWebsiteService{db: db}
}

// settingWebsiteService 网站信息配置服务实现类
type settingWebsiteService struct {
	db *gorm.DB
}

// Detail 获取网站信息
func (wSrv settingWebsiteService) Detail() (res map[string]string, e error) {
	data, err := SystemConfigService.Get(wSrv.db, "website")
	if e = response.CheckErr(err, "Detail Get err"); e != nil {
		return
	}
	return map[string]string{
		"name":      data["name"],
		// 业务表存完整访问地址（/api/uploads/<id>），无需再转换
		"logo":      data["logo"],
		"favicon":   data["favicon"],
		"backdrop":  data["backdrop"],
		"shop_name": data["shop_name"],
		"shop_logo": data["shop_logo"],
	}, nil
}

// Save 保存网站信息
func (wSrv settingWebsiteService) Save(wsReq setting_schema.SettingWebsiteReq) (e error) {
	err := SystemConfigService.Set(wSrv.db, "website", "name", wsReq.Name)
	if e = response.CheckErr(err, "Save Set name err"); e != nil {
		return
	}
	err = SystemConfigService.Set(wSrv.db, "website", "logo", wsReq.Logo)
	if e = response.CheckErr(err, "Save Set logo err"); e != nil {
		return
	}
	err = SystemConfigService.Set(wSrv.db, "website", "favicon", wsReq.Favicon)
	if e = response.CheckErr(err, "Save Set favicon err"); e != nil {
		return
	}
	err = SystemConfigService.Set(wSrv.db, "website", "backdrop", wsReq.Backdrop)
	if e = response.CheckErr(err, "Save Set backdrop err"); e != nil {
		return
	}
	err = SystemConfigService.Set(wSrv.db, "website", "shop_name", wsReq.ShopName)
	if e = response.CheckErr(err, "Save Set shop_name err"); e != nil {
		return
	}
	err = SystemConfigService.Set(wSrv.db, "website", "shop_logo", wsReq.ShopLogo)
	e = response.CheckErr(err, "Save Set shop_logo err")
	return
}
