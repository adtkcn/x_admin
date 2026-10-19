package common_service

import (
	"strconv"
	"strings"
	"x_admin/app/service/setting_service"
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

// OnlineTrend 在线人数趋势曲线初始数据（从 Redis onlineCount 列表取，后续增量由 WS 推送）
func (iSrv indexService) OnlineTrend() (res map[string]any, e error) {
	// 在线用户数量
	onlineRecords := util.RedisUtil.LRange("onlineCount", 0, -1)

	// 从 Redis 数据中提取时间和在线数（格式: "15:04:05,count"）
	var dateList []string
	var countList []int
	for _, record := range onlineRecords {
		parts := strings.Split(record, ",")
		if len(parts) != 2 {
			continue
		}
		dateList = append(dateList, parts[0])
		if c, err := strconv.Atoi(parts[1]); err == nil {
			countList = append(countList, c)
		} else {
			countList = append(countList, 0)
		}
	}

	return map[string]any{
		"date": dateList,
		"list": countList,
	}, nil
}

// Config 公共配置
func (iSrv indexService) Config() (res map[string]any, e error) {
	const cacheKey = "Index:Config"
	// 先读缓存
	if cacheStr := util.RedisUtil.Get(cacheKey); cacheStr != "" {
		if res, e = util.ToolsUtil.JsonToObj[map[string]any](cacheStr); e != nil {
			core.Logger.Errorf("Config cache JsonToObj err: %v", e)
		} else {
			return res, nil
		}
	}
	website, err := setting_service.SystemConfigService.Get(iSrv.db, "website")
	if e = response.CheckErr(err, "Config Get err"); e != nil {
		return
	}
	var copyright []map[string]string
	if copyrightStr := website["copyright"]; copyrightStr != "" {
		if copyright, e = util.ToolsUtil.JsonToObj[[]map[string]string](copyrightStr); e != nil {
			e = response.CheckErr(e, "Config JsonToObj err")
			return
		}
	} else {
		copyright = []map[string]string{}
	}
	res = map[string]any{
		"webName": website["name"],
		// 业务表存完整访问地址（/api/uploads/<id>），无需再转换
		"webLogo":     website["logo"],
		"webFavicon":  website["favicon"],
		"webBackdrop": website["backdrop"],
		"ossDomain":   config.AppConfig.OssDomain,
		"copyright":   copyright,
	}
	// 写入缓存
	if cacheStr, jErr := util.ToolsUtil.ObjToJson(res); jErr == nil {
		util.RedisUtil.Set(cacheKey, cacheStr, 20)
	}
	return res, nil
}
