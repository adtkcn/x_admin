package monitor_service

import (
	"encoding/json/v2"
	"errors"
	"time"
	"x_admin/core"
	"x_admin/util"

	"gorm.io/gorm"
)

var MonitorServerService = NewMonitorServerService()

// NewMonitorServerService 初始化服务器监控服务
func NewMonitorServerService() *monitorServerService {
	return &monitorServerService{
		db: core.GetDB(),
		CacheUtil: util.CacheUtil{
			Name: "{monitorServer}",
		},
	}
}

// monitorServerService 服务器监控服务
type monitorServerService struct {
	db        *gorm.DB
	CacheUtil util.CacheUtil
}

// CollectAndPushServerInfo 收集服务器信息并推送到Redis,通过定时任务调用
func (service *monitorServerService) CollectAndPushServerInfo() error {
	var sys = util.ServerUtil.GetSysInfo()
	var computerIp = sys["computerIp"].(string)

	serverInfo := map[string]any{
		"cpu":       util.ServerUtil.GetCpuInfo(),
		"mem":       util.ServerUtil.GetMemInfo(),
		"sys":       sys,
		"disk":      util.ServerUtil.GetDiskInfo(),
		"go":        util.ServerUtil.GetGoInfo(),
		"timestamp": time.Now().Unix(),
	}

	// 通过 CacheUtil 写入单台机器的监控信息,移除旧的独立 :ips 集合。
	if !service.CacheUtil.SetCache(computerIp, serverInfo) {
		return errors.New("写入服务器监控信息失败")
	}
	return nil
}

// GetServerInfoFromRedis 从Redis获取服务器信息
// func (service *monitorServerService) GetServerInfoFromRedis(computerIp string) (map[string]any, error) {
// 	data := map[string]any{}
// 	if err := service.CacheUtil.GetCache(computerIp, &data); err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// GetAllServerLatestInfo 获取所有服务器最新信息
func (service *monitorServerService) GetAllServerLatestInfo() (map[string]map[string]any, error) {

	// 通过 Hash Tag 前缀扫描得到所有机器的监控数据,不再依赖独立 :ips 集合。
	all := util.RedisUtil.ValuesByPrefix(service.CacheUtil.Name)

	serverInfos := make(map[string]map[string]any)

	for computerIp, raw := range all {

		// 直接解析 ValuesByPrefix 已返回的 JSON,避免重复查询 Redis
		data := map[string]any{}
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			core.Logger.Errorf("解析服务器 %s 的监控信息失败: %v", computerIp, err)
			continue
		}

		serverInfos[computerIp] = data
	}

	return serverInfos, nil
}
