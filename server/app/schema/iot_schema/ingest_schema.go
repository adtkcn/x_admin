package iot_schema

import "github.com/adtkcn/x_null"

// ============ 设备接入层（HTTP/MQTT/TCP/WS 共用报文契约） ============

// 设备消息类型：接入层只负责组装与投递，具体处理由 iot/pipeline 按类型分派
const (
	MsgTypeTelemetry = "telemetry" // 属性/遥测上报
	MsgTypeHeartbeat = "heartbeat" // 心跳（仅刷新在线态）
)

// 设备接入请求头：X-Device-Sn + X-Access-Token（设备凭证见 x_iot_device.access_token）
const (
	HeaderDeviceSn    = "X-Device-Sn"
	HeaderAccessToken = "X-Access-Token"
)

// DeviceAuthInfo 设备鉴权结果：接入层与处理层共用的最小上下文，缓存于 Redis 以免高频查库
type DeviceAuthInfo struct {
	DeviceId       string `json:"device_id"`
	Sn             string `json:"sn"`
	Name           string `json:"name"`
	ProductId      string `json:"product_id"`
	SpaceId        string `json:"space_id"`
	Status         uint8  `json:"status"`
	Protocol       string `json:"protocol"`        // 产品接入协议: mqtt/tcp/http/ws
	ProductStop    bool   `json:"product_stop"`    // 产品是否停用
	OfflineTimeout int    `json:"offline_timeout"` // 生效的离线判定时长(秒)，已按 设备>产品>兜底值 回退
}

// DeviceTelemetryReq 属性上报请求体
//
//	{"values":{"temperature":26.3,"humidity":55},"ts":1730000000000}
type DeviceTelemetryReq struct {
	Values map[string]any `json:"values" binding:"required,min=1"` // 属性键值：identifier -> 值
	Ts     int64          `json:"ts"`                              // 设备侧采集时间(毫秒)，缺省取服务端接收时间
}

// DeviceHeartbeatReq 心跳请求体
type DeviceHeartbeatReq struct {
	Ts int64 `json:"ts"`
}

// IngestMsg 投递到 core.Queue 的标准报文信封：接入层与处理层解耦的唯一契约。
// 后续 MQTT/TCP/WS 接入只需产出同样的信封。
type IngestMsg struct {
	MsgType        string         `json:"msg_type"`        // telemetry/heartbeat
	DeviceId       string         `json:"device_id"`       //
	Sn             string         `json:"sn"`              //
	ProductId      string         `json:"product_id"`      //
	SpaceId        string         `json:"space_id"`        //
	Protocol       string         `json:"protocol"`        // 接入协议 http/mqtt/tcp/ws
	OfflineTimeout int            `json:"offline_timeout"` // 设备生效的离线判定时长(秒)，随鉴权信息带下
	Values         map[string]any `json:"values,omitempty"`
	Ts             int64          `json:"ts"`          // 采集时间(毫秒)
	ReceivedAt     int64          `json:"received_at"` // 服务端接收时间(毫秒)
}

// DeviceStateReq 运行态回读请求
type DeviceStateReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 设备主键
}

// DeviceDataReq 时序历史回读请求：按时间区间取数，供趋势图与链路验证
type DeviceDataReq struct {
	ID      string `json:"id" form:"id" binding:"required"` // 设备主键
	StartMs int64  `json:"start_ms" form:"start_ms"`        // 开始时间(毫秒)，0=不限下界
	EndMs   int64  `json:"end_ms" form:"end_ms"`            // 结束时间(毫秒)，0=不限上界
	Limit   int    `json:"limit" form:"limit"`              // 条数，<=0 表示不限制（区间内全量回读）
}

// DeviceDataResp 一条时序历史记录（属性已反序列化为 map，便于前端直接渲染）
type DeviceDataResp struct {
	Ts     int64          `json:"ts"`     // 采集时间(毫秒)
	Values map[string]any `json:"values"` // 属性键值（规则链加工后，已按物模型约束）
}

// DeviceStateResp 设备运行态快照（管理端回读，数据源为 Redis）
type DeviceStateResp struct {
	DeviceId       string            `json:"device_id"`
	Sn             string            `json:"sn"`
	Name           string            `json:"name"`            // 设备名称（由控制器从库侧回填，Redis 不存）
	Online         bool              `json:"online"`          // 在线判定窗口内是否收到过报文
	OfflineTimeout int               `json:"offline_timeout"` // 生效的离线判定时长(秒)
	LastSeenMs     int64             `json:"last_seen_ms"`    // 最后一条报文的服务端接收时间(毫秒)
	Values         map[string]string `json:"values"`          // 最新属性快照 identifier -> 值(字符串化)
	UpdatedAt      string            `json:"updated_at"`      // 最后报文时间的可读格式（与 last_seen_ms 同源）
	DBStatus       uint8             `json:"db_status"`       // 库中在线状态（由定时扫描与上报共同维护）
	LastOnline     x_null.Time       `json:"last_online_time"`
	LastOffline    x_null.Time       `json:"last_offline_time"`
}

// ScopeValues 按物模型 identifier 白名单裁剪属性，是“落库/存快照只保留已建模属性”约束的唯一实现，
// 供时序存储、运行态快照、告警记录三处共用。allow 为空表示未配物模型（或旧形态解析不出），
// 此时不约束、直接返回原 map，避免配置缺失反而误丢数据。
func ScopeValues(values map[string]any, allow map[string]bool) map[string]any {
	if len(allow) == 0 {
		return values
	}
	out := make(map[string]any, len(values))
	for k, v := range values {
		if allow[k] {
			out[k] = v
		}
	}
	return out
}
