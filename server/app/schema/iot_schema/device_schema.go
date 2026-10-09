package iot_schema

import (
	"x_admin/core/request"

	"github.com/adtkcn/x_null"
)

// ============ 设备 ============
// 查询条件统一用 x_null.String：不传/空串均视为“不筛选”，无需 -1 之类的哨兵值
type DeviceListReq struct {
	Name      x_null.String `json:"name" form:"name"`             // 设备名称(模糊)
	Sn        x_null.String `json:"sn" form:"sn"`                 // 设备编号(模糊)
	ProductId x_null.String `json:"product_id" form:"product_id"` // 产品ID
	SpaceId   x_null.String `json:"space_id" form:"space_id"`     // 区域ID(含子区域)
	Status    x_null.String `json:"status" form:"status"`         // 在线状态: 0离线 1在线
	StartTime x_null.String `json:"start_time" form:"start_time"` // 创建时间起(2006-01-02 15:04:05)
	EndTime   x_null.String `json:"end_time" form:"end_time"`     // 创建时间止(2006-01-02 15:04:05)
	request.PageReq
}

type DeviceAddReq struct {
	Name           string  `json:"name" form:"name" binding:"required,min=1,max=100"`          // 设备名称
	Sn             string  `json:"sn" form:"sn" binding:"required,min=1,max=64"`               // 设备编号
	ProductId      string  `json:"product_id" form:"product_id" binding:"required"`            // 产品ID
	SpaceId        string  `json:"space_id" form:"space_id"`                                   // 区域ID
	GatewayId      string  `json:"gateway_id" form:"gateway_id"`                               // 网关设备ID
	OfflineTimeout int     `json:"offline_timeout" form:"offline_timeout" binding:"max=86400"` // 离线判定时长(秒)，0=自动取所属产品值
	Longitude      float64 `json:"longitude" form:"longitude"`                                 // 经度
	Latitude       float64 `json:"latitude" form:"latitude"`                                   // 纬度
	Remark         string  `json:"remark" form:"remark" binding:"max=255"`                     // 备注
}

type DeviceEditReq struct {
	ID             string  `json:"id" form:"id" binding:"required"`                            // 主键
	Name           string  `json:"name" form:"name" binding:"required,min=1,max=100"`          // 设备名称
	Sn             string  `json:"sn" form:"sn" binding:"required,min=1,max=64"`               // 设备编号
	ProductId      string  `json:"product_id" form:"product_id" binding:"required"`            // 产品ID
	SpaceId        string  `json:"space_id" form:"space_id"`                                   // 区域ID
	GatewayId      string  `json:"gateway_id" form:"gateway_id"`                               // 网关设备ID
	OfflineTimeout int     `json:"offline_timeout" form:"offline_timeout" binding:"max=86400"` // 离线判定时长(秒)，0=自动取所属产品值
	Longitude      float64 `json:"longitude" form:"longitude"`                                 // 经度
	Latitude       float64 `json:"latitude" form:"latitude"`                                   // 纬度
	Remark         string  `json:"remark" form:"remark" binding:"max=255"`                     // 备注
}

type DeviceDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type DeviceDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// DeviceMapReq 大屏地图取数请求：按区域过滤（逗号分隔多个 id，选父区域时由前端把子树 id 集齐传入；空=全部设备）
type DeviceMapReq struct {
	SpaceIds string `json:"space_ids" form:"space_ids"` // 区域ID列表(逗号分隔)，缺省不限
}

// DeviceMapItem 大屏地图一个设备点：坐标 + 在线状态 + 活动告警级别
type DeviceMapItem struct {
	ID        string  `json:"id"`
	Sn        string  `json:"sn"`
	Name      string  `json:"name"`
	SpaceId   string  `json:"space_id"`
	ProductId string  `json:"product_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Status    uint8   `json:"status"`   // 0=离线 1=在线
	Severity  int8    `json:"severity"` // 活动告警级别 0=无告警 -1=故障 1=轻微 2=严重（数值即危害序，3+ 预留；!=0 即告警中，故障优先于其他告警）
}

type DeviceResp struct {
	ID              string      `json:"id"`                // 主键
	Name            string      `json:"name"`              // 设备名称
	Sn              string      `json:"sn"`                // 设备编号
	ProductId       string      `json:"product_id"`        // 产品ID
	ProductName     string      `json:"product_name"`      // 产品名称
	ProductNodeType string      `json:"product_node_type"` // 所属产品节点类型: DEVICE/GATEWAY/EDGE
	SpaceId         string      `json:"space_id"`          // 区域ID
	SpaceName       string      `json:"space_name"`        // 区域名称
	GatewayId       string      `json:"gateway_id"`        // 网关设备ID
	GatewayName     string      `json:"gateway_name"`      // 网关设备名称
	Status          uint8       `json:"status"`            // 在线状态（系统自动维护）
	OfflineTimeout  int         `json:"offline_timeout"`   // 离线判定时长(秒)，以设备记录为准（冗余产品值+可自定义），0=兜底300秒
	LastOnlineTime  x_null.Time `json:"last_online_time"`  // 最后上线时间
	LastOfflineTime x_null.Time `json:"last_offline_time"` // 最后离线时间
	Longitude       float64     `json:"longitude"`         // 经度
	Latitude        float64     `json:"latitude"`          // 纬度
	AccessToken     string      `json:"access_token"`      // 接入凭证
	Remark          string      `json:"remark"`            // 备注
	CreateTime      x_null.Time `json:"create_time"`       // 创建时间
	UpdateTime      x_null.Time `json:"update_time"`       // 更新时间
}
