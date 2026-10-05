package config

// IotConfigStruct IoT 设备接入层与数据处理层配置
//
// 离线判定由到期表（ZSET）驱动：报文/心跳 ZADD 刷新到期时间，app/corn 的
// iotOfflineScan 快路径弹到期、iotOfflineSweep 慢路径全扫兜底并重建；
// 判定时长只取设备 offline_timeout 字段（秒，新增/编辑时已冗余产品值），
// 为 0 时落到兜底值 iot_service.DefaultOfflineTimeout
type IotConfigStruct struct {
	// Enable 接入层总开关：false 时不注册设备侧路由、不启动处理流水线
	Enable bool `mapstructure:"Enable"`

	// DeviceCacheTTL 设备鉴权信息 Redis 缓存时长（秒），承受高频上报时避免每报文查库
	DeviceCacheTTL int `mapstructure:"DeviceCacheTTL"`
	// ThingCacheTTL 物模型缓存时长（秒）
	ThingCacheTTL int `mapstructure:"ThingCacheTTL"`
	// StateTTL 设备最新属性快照在 Redis 中的保留时长（秒）
	StateTTL int `mapstructure:"StateTTL"`
	// MaxValues 单次上报允许的最大属性条数，防报文放大
	MaxValues int `mapstructure:"MaxValues"`
	// PushStateChange 在线状态变化时是否经 core.Ws 广播（大屏/房间）
	PushStateChange bool `mapstructure:"PushStateChange"`
	// LogRaw 是否记录原始报文日志（联调期打开）
	LogRaw bool `mapstructure:"LogRaw"`

	// DataBatchSize 时序历史批量写入的单批行数上限（攒满或到 2 秒落库一次）
	DataBatchSize int `mapstructure:"DataBatchSize"`
	// DataKeepDays 时序历史保留天数，每天凌晨清理更早的数据；0=永久保留不清理
	DataKeepDays int `mapstructure:"DataKeepDays"`
}

var IotConfig = IotConfigStruct{
	Enable: true,

	DeviceCacheTTL:  300,
	ThingCacheTTL:   600,
	StateTTL:        86400,
	MaxValues:       50,
	PushStateChange: true,
	LogRaw:          false,
	DataBatchSize:   200,
	DataKeepDays:    30,
}

// DeviceCtxKey 设备鉴权中间件写入 gin.Context 的键（值为 *iot_schema.DeviceAuthInfo），
// 读取请用 middleware.DeviceFromCtx，避免各处散落裸 c.Get。
const DeviceCtxKey = "iot_device"

// Normalize 校正非法/零值配置，必须在 viper 载入之后调用。
//
// 这些字段被直接当作 Redis TTL 使用，写成 0 不会报错而是语义反转：
//   - MaxValues=0 → 任何上报都算超上限，接入层全部拒绝；
//   - StateTTL=0  → 属性快照写完即失效，管理端回读为空。
//
// 设备的 offline_timeout 不受本函数约束：0 是合法值（含义为未指定，
// 新增/编辑时已把产品值冗余到设备，仍为 0 则判定落兜底 300 秒），
// 到期表窗口在 iot.pipeline.offlineWindow 里解析。
//
// 故非正数一律回退到包级默认值。
func (c *IotConfigStruct) Normalize() {
	d := IotConfig

	if c.DeviceCacheTTL <= 0 {
		c.DeviceCacheTTL = d.DeviceCacheTTL
	}
	if c.ThingCacheTTL <= 0 {
		c.ThingCacheTTL = d.ThingCacheTTL
	}
	if c.StateTTL <= 0 {
		c.StateTTL = d.StateTTL
	}
	if c.MaxValues <= 0 {
		c.MaxValues = d.MaxValues
	}
	if c.DataBatchSize <= 0 {
		c.DataBatchSize = d.DataBatchSize
	}
	// DataKeepDays 的 0 是合法值（永久保留），只校正负数
	if c.DataKeepDays < 0 {
		c.DataKeepDays = d.DataKeepDays
	}
}
