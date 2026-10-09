package iot_model

import (
	"math"
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

// 告警级别：由规则链告警节点脚本 alert(values, sn) 的返回值决定（返回 0 即不告警/恢复，0 不是级别）。
// 数值本身就是危害升序：故障(-1) < 轻微(1) < 严重(2)，故障危害最轻故下沉到负数档，
// 正数 3 及以上预留给未来可能的高级告警；级别比较一律走 AlertPriority，不要直接比数值。
// 预留可扩展：新增级别只需在此加常量、前端 SeverityMap 加一项，脚本直接返回对应数字。
const (
	AlertSevFault   int8 = -1 // 故障（危害最轻，但覆盖优先级最高）
	AlertSevMinor   int8 = 1  // 一级轻微告警
	AlertSevSerious int8 = 2  // 二级严重报警
)

// AlertPriority 告警的覆盖优先序：故障代表设备自身失能，其期间的阈值告警都不可信，
// 故一旦命中就以它为准（优先于其他告警）；其余级别仍按危害数值比较。
// 优先序只决定“本次事件记哪个级别”，不改变危害档（故障仍是最低档，展示/筛选用原值）。
func AlertPriority(level int8) int {
	if level == AlertSevFault {
		return math.MaxInt16 // 远高于任何合法级别（int8 至多 127）
	}
	return int(level)
}

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
	Severity    int8        `gorm:"not null;default:1;index;comment:'告警级别: -1=故障, 1=轻微, 2=严重(数值即危害序,3+预留)'" json:"severity"` // 故障危害最轻但优先接管，比较一律过 AlertPriority
	Values      string      `gorm:"type:text;comment:'触发/最近命中属性快照JSON'" json:"values"`
	TriggeredAt x_null.Time `gorm:"not null;index:idx_iot_alert_dev_ts,priority:2;comment:'首次触发时间'" json:"triggered_at"`
	RecoveredAt x_null.Time `gorm:"default:null;comment:'恢复时间(未恢复为null)'" json:"recovered_at"`
	CreateTime  x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime  x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
}

func (IotAlertHistory) TableName() string { return "x_iot_alert_history" }

func (m *IotAlertHistory) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
