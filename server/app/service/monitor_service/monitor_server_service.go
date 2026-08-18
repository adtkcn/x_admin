package monitor_service

import (
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
			Name: "monitorServer",
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

	service.CacheUtil.SetCache(computerIp, serverInfo)
	// 注意：此处不要对 :ips 集合设置过期。定时任务每次刷新都会重新 SAdd/Expire，
	// 只要服务在跑，TTL 会持续被续期而永不过期，等同未设置；
	// 且 SAdd 相同 ip 成员会被 set 去重，集合成员数恒定（=机器数），不会无限增长。
	util.RedisUtil.SAdd(service.CacheUtil.Name+":ips", computerIp)
	return nil
}

// GetServerInfoFromRedis 从Redis获取服务器信息
func (service *monitorServerService) GetServerInfoFromRedis(computerIp string) (map[string]any, error) {
	data := map[string]any{}
	err := service.CacheUtil.GetCache(computerIp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAllServerLatestInfo 获取所有服务器最新信息
func (service *monitorServerService) GetAllServerLatestInfo() (map[string]map[string]any, error) {

	ips := util.RedisUtil.SGet(service.CacheUtil.Name + ":ips")

	serverInfos := make(map[string]map[string]any)

	for _, computerIp := range ips {

		// 获取最新的服务器信息
		info, err := service.GetServerInfoFromRedis(computerIp)
		if err != nil {
			core.Logger.Errorf("获取服务器 %s 的监控信息失败: %v", computerIp, err)
			continue
		}

		serverInfos[computerIp] = info
	}

	return serverInfos, nil
}
