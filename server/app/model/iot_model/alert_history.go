package iot_model

import (
	"time"
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

// 告警记录状态：一条记录代表一次“告警事件”（从触发到恢复的完整生命周期），
// 持续期间的重复命中只维持 Redis 活动态、不再新增/更新记录（避免每条报文写库）；
// 恢复时置 Recovered 并回填 RecoveredAt。告警持续时长由 recovered_at - triggered_at 得出。
const (
	AlertStatusFiring    uint8 = 1 // 告警中
	AlertStatusRecovered uint8 = 0 // 已恢复
)

// 告警级别：由规则链告警节点脚本 alert(values, sn) 的返回值决定（返回 0 即不告警）。
// 预留可扩展：新增级别只需在此加常量、前端 SeverityMap 加一项，脚本直接返回对应数字。
const (
	AlertSevMinor   uint8 = 1 // 一级轻微告警
	AlertSevSerious uint8 = 2 // 二级严重报警
	AlertSevFault   uint8 = 3 // 故障
)

// DefaultAlertSeverity 脚本返回 true(布尔) 时兑底的默认级别（兼容旧契约）
const DefaultAlertSeverity = AlertSevMinor

// IotAlertHistory 设备告警历史（iot.md：告警记录落 MySQL，运行态在 Redis）
//
// 由规则链告警节点脚本判定命中，处理层的告警状态机按“持续时长/静默/恢复”
// 收敛成事件级记录：触发时建一条 status=1，恢复时更新 status=0 + recovered_at。
// 属性快照以 JSON 文本整包存，与设备时序表一致，避免为不固定的物模型频繁加列。
type IotAlertHistory struct {
	ID          string      `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	DeviceId    string      `gorm:"type:char(36);not null;default:'';index:idx_iot_alert_dev_ts,priority:1;comment:'设备ID'" json:"device_id"`
	Sn          string      `gorm:"type:varchar(64);not null;default:'';comment:'设备编号'" json:"sn"`
	ProductId   string      `gorm:"type:char(36);not null;default:'';index;comment:'产品ID'" json:"product_id"`
	SpaceId     string      `gorm:"type:char(36);not null;default:'';index;comment:'区域ID'" json:"space_id"`
	Status      uint8       `gorm:"not null;default:1;comment:'状态: 1=告警中, 0=已恢复'" json:"status"`
	Severity    uint8       `gorm:"not null;default:1;index;comment:'告警级别: 1=轻微, 2=严重, 3=故障'" json:"severity"`
	Values      string      `gorm:"type:text;comment:'触发/最近命中属性快照JSON'" json:"values"`
	TriggeredAt time.Time   `gorm:"not null;index:idx_iot_alert_dev_ts,priority:2;comment:'首次触发时间'" json:"triggered_at"`
	RecoveredAt x_null.Time `gorm:"default:null;comment:'恢复时间(未恢复为null)'" json:"recovered_at"`
	CreateTime  x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime  x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
}

func (IotAlertHistory) TableName() string { return "x_iot_alert_history" }

func (m *IotAlertHistory) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
