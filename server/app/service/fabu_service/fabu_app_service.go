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

// Del 删除应用：事务内级联软删其下版本与热更新包，避免孤儿数据
func (s fabuAppService) Del(id string) (e error) {
	e = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("app_id = ?", id).Delete(&fabu_model.FabuWgt{}).Error; err != nil {
			return response.CheckErr(err, "删除热更新包失败")
		}
		if err := tx.Where("app_id = ?", id).Delete(&fabu_model.FabuAppVersion{}).Error; err != nil {
			return response.CheckErr(err, "删除版本失败")
		}
		if err := tx.Where("id = ?", id).Delete(&fabu_model.FabuApp{}).Error; err != nil {
			return response.CheckErr(err, "删除应用失败")
		}
		return nil
	})
	return
}

func (s fabuAppService) FindByShortUrl(shortUrl string) (app fabu_model.FabuApp, e error) {
	if e = response.CheckErr(s.db.Where("short_url = ?", shortUrl).First(&app).Error, "应用不存在"); e != nil {
		return
	}
	return
}

// FindByBundleId 按 BundleId+平台定位应用（安卓/iOS 包名可能相同，必须区分平台）
func (s fabuAppService) FindByBundleId(bundleId, platform string, app *fabu_model.FabuApp) (e error) {
	if e = response.CheckErr(s.db.Where("bundle_id = ? AND platform = ?", bundleId, platform).First(app).Error, "应用不存在"); e != nil {
		return
	}
	return
}
