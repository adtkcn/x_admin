package common_service

import (
	"time"
	"x_admin/app/service/setting_service"
	"x_admin/app/service/system_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"gorm.io/gorm"
)

var IndexService = NewIndexService()

// NewIndexService 初始化
func NewIndexService() *indexService {
	db := core.GetDB()
	return &indexService{db: db}
}

// indexService 主页服务实现类
type indexService struct {
	db *gorm.DB
}

// Console 控制台数据
func (iSrv indexService) Console() (res map[string]any, e error) {
	// 版本信息
	name, err := setting_service.SystemConfigService.GetVal(iSrv.db, "website", "name", "x_admin-Go")
	if e = response.CheckErr(err, "Console Get err"); e != nil {
		return
	}
	version := map[string]any{
		"name":    name,
		"version": config.AppConfig.Version,
	}
	adminCount, err := system_service.AdminService.GetTodayCount()
	// 今日数据
	today := map[string]any{
		"time":        util.NullTimeUtil.Now(),
		"todayVisits": 0,                     // 访问量(人)
		"totalVisits": 0,                     // 总访问量
		"flow_todo":   0,                     // 待办审批
		"todayOrder":  0,                     // 订单量(笔)
		"totalOrder":  0,                     // 总订单量
		"todayUsers":  adminCount.TodayUsers, // 新增用户
		"totalUsers":  adminCount.TotalUsers, // 总访用户
	}

	// 在线用户
	onlineCount := util.RedisUtil.LRange("onlineCount", 0, -1)

	// 访客图表
	now := time.Now()
	var date []string
	for i := 14; i >= 0; i-- {
		date = append(date, now.AddDate(0, 0, -i).Format(config.ConstantConfig.DateFormat))
	}
	visitor := map[string]any{
		"date": date,
		"list": onlineCount,
	}
	return map[string]any{
		"version": version,
		"today":   today,
		"visitor": visitor,
	}, nil
}

// Config 公共配置
func (iSrv indexService) Config() (res map[string]any, e error) {
	website, err := setting_service.SystemConfigService.Get(iSrv.db, "website")
	if e = response.CheckErr(err, "Config Get err"); e != nil {
		return
	}
	copyrightStr, err := setting_service.SystemConfigService.GetVal(iSrv.db, "website", "copyright", "")
	if e = response.CheckErr(err, "Config GetVal err"); e != nil {
		return
	}
	var copyright []map[string]string
	if copyrightStr != "" {
		err = util.ToolsUtil.JsonToObj(copyrightStr, &copyright)
		if e = response.CheckErr(err, "Config JsonToObj err"); e != nil {
			return
		}
	} else {
		copyright = []map[string]string{}
	}
	return map[string]any{
		"webName":     website["name"],
		"webLogo":     util.UrlUtil.ToAbsoluteUrl(website["logo"]),
		"webFavicon":  util.UrlUtil.ToAbsoluteUrl(website["favicon"]),
		"webBackdrop": util.UrlUtil.ToAbsoluteUrl(website["backdrop"]),
		"ossDomain":   config.AppConfig.OssDomain,
		"copyright":   copyright,
	}, nil
}
