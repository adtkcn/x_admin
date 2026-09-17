package fabu_service

import (
	"x_admin/app/model/fabu_model"
	"x_admin/app/schema/fabu_schema"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util/convert_util"

	"gorm.io/gorm"
)

var AppService = NewFabuAppService()

func NewFabuAppService() *fabuAppService {
	return &fabuAppService{db: core.GetDB()}
}

type fabuAppService struct {
	db *gorm.DB
}

func (s fabuAppService) List(req fabu_schema.FabuAppListReq) (res map[string]any, e error) {
	var apps []fabu_model.FabuApp
	var total int64
	chain := s.db.Model(&fabu_model.FabuApp{})
	if req.Keyword != "" {
		chain = chain.Where("name like ? or bundle_id like ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if e = response.CheckErr(chain.Count(&total).Error, "统计应用失败"); e != nil {
		return
	}
	pageNo, pageSize := req.PageNo, req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if e = response.CheckErr(chain.Order("create_time desc").Offset((pageNo-1)*pageSize).Limit(pageSize).Find(&apps).Error, "查询应用失败"); e != nil {
		return
	}
	var list []fabu_schema.FabuAppResp
	convert_util.Copy(&list, apps)
	res = map[string]any{"count": total, "lists": list}
	return
}

func (s fabuAppService) Detail(id string) (res fabu_schema.FabuAppResp, e error) {
	var app fabu_model.FabuApp
	if e = response.CheckErr(s.db.Where("id = ?", id).First(&app).Error, "应用不存在"); e != nil {
		return
	}
	convert_util.Copy(&res, app)
	return
}

func (s fabuAppService) Add(req fabu_schema.FabuAppAddReq) (e error) {
	app := fabu_model.FabuApp{
		Name:        req.Name,
		Platform:    req.Platform,
		BundleId:    req.BundleId,
		BundleName:  req.BundleName,
		Version:     req.Version,
		VersionCode: req.VersionCode,
		ShortUrl:    req.ShortUrl,
		Icon:        req.Icon,
	}
	if e = response.CheckErr(s.db.Create(&app).Error, "创建应用失败"); e != nil {
		return
	}
	return
}

func (s fabuAppService) Edit(req fabu_schema.FabuAppEditReq) (e error) {
	updates := map[string]any{
		"name":        req.Name,
		"bundle_name": req.BundleName,
		"short_url":   req.ShortUrl,
		"icon":        req.Icon,
	}
	if e = response.CheckErr(s.db.Model(&fabu_model.FabuApp{}).Where("id = ?", req.ID).Updates(updates).Error, "更新应用失败"); e != nil {
		return
	}
	return
}

func (s fabuAppService) Del(id string) (e error) {
	if e = response.CheckErr(s.db.Where("id = ?", id).Delete(&fabu_model.FabuApp{}).Error, "删除应用失败"); e != nil {
		return
	}
	return
}

func (s fabuAppService) FindByShortUrl(shortUrl string) (app fabu_model.FabuApp, e error) {
	if e = response.CheckErr(s.db.Where("short_url = ?", shortUrl).First(&app).Error, "应用不存在"); e != nil {
		return
	}
	return
}

func (s fabuAppService) FindByBundleId(bundleId string, app *fabu_model.FabuApp) (e error) {
	if e = response.CheckErr(s.db.Where("bundle_id = ?", bundleId).First(app).Error, "应用不存在"); e != nil {
		return
	}
	return
}
