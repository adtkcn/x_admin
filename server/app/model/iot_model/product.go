package iot_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IotProduct 产品定义
type IotProduct struct {
	ID             string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Name           string                `gorm:"type:varchar(100);not null;default:'';comment:'产品名称'" json:"name"`
	CategoryId     string                `gorm:"type:char(36);not null;default:'';index;comment:'分类ID'" json:"category_id"`
	DeviceType     string                `gorm:"type:varchar(50);not null;default:'';comment:'节点类型: DEVICE=直连设备, GATEWAY=网关设备, EDGE=边缘主机'" json:"device_type"`
	Protocol       string                `gorm:"type:varchar(20);not null;default:'';comment:'接入协议:mqtt/tcp/http/ws'" json:"protocol"`
	OfflineTimeout int                   `gorm:"not null;default:0;comment:'离线判定时长(秒): 新增/编辑必填 1~86400，作为其下设备未自定义时的取值；变更后同步到未自定义的设备'" json:"offline_timeout"`
	ThingModel     string                `gorm:"type:text;comment:'物模型JSON'" json:"thing_model"`
	DecodeScript   string                `gorm:"type:text;comment:'解码脚本(goja): decode(raw)返回物模型对象'" json:"decode_script"`
	FilterScript   string                `gorm:"type:text;comment:'数据过滤脚本(goja): filter(values)返回布尔,false丢弃本条上报'" json:"filter_script"`
	AlertScript    string                `gorm:"type:text;comment:'数据告警脚本(goja): alert(values)返回布尔,true触发告警'" json:"alert_script"`
	AlertDuration  int                   `gorm:"not null;default:0;comment:'告警持续时长(秒): 异常连续满足该时长才触告,0=命中即告'" json:"alert_duration"`
	AlertSilence   int                   `gorm:"not null;default:0;comment:'告警静默时长(秒): 触发后该时长内不再重复触告,0=不静默'" json:"alert_silence"`
	Remark         string                `gorm:"type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	IsStop         uint8                 `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'" json:"is_stop"`
	IsDelete       soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime     x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime     x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime     x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotProduct) TableName() string { return "x_iot_product" }

func (m *IotProduct) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
