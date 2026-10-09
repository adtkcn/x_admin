package iot_schema

import (
	"x_admin/core/request"

	"github.com/adtkcn/x_null"
)

// ============ 产品 ============

// 节点类型常量（对齐阿里云/腾讯云 IoT 惯例，存于产品 device_type 字段）：
// 网关身份由产品节点类型定义，设备本身不再隐式充当网关
const (
	NodeTypeDevice  = "DEVICE"  // 直连设备：只能作为子设备接入网关，不能挂子设备
	NodeTypeGateway = "GATEWAY" // 网关设备：可挂载子设备
	NodeTypeEdge    = "EDGE"    // 边缘主机：兼具网关能力
)

// NodeTypeSupportsGateway 判断节点类型是否具备网关能力（空值为历史遗留数据，按直连设备处理）
func NodeTypeSupportsGateway(nodeType string) bool {
	return nodeType == NodeTypeGateway || nodeType == NodeTypeEdge
}

type ProductListReq struct {
	Name       x_null.String `json:"name" form:"name"`               // 产品名称(模糊)
	CategoryId x_null.String `json:"category_id" form:"category_id"` // 分类ID(含子分类)
	IsStop     x_null.String `json:"is_stop" form:"is_stop"`         // 是否停用: 0正常 1停用，不传=全部
	request.PageReq
}

type ProductAddReq struct {
	Name           string `json:"name" form:"name" binding:"required,min=1,max=100"`                           // 产品名称
	CategoryId     string `json:"category_id" form:"category_id" binding:"required"`                           // 分类ID
	DeviceType     string `json:"device_type" form:"device_type" binding:"required,oneof=DEVICE GATEWAY EDGE"` // 节点类型
	Protocol       string `json:"protocol" form:"protocol" binding:"required,oneof=mqtt tcp http ws"`          // 接入协议
	OfflineTimeout int    `json:"offline_timeout" form:"offline_timeout" binding:"required,min=1,max=86400"`   // 离线判定时长(秒)，产品级必填
	ThingModel     string `json:"thing_model" form:"thing_model"`                                              // 物模型JSON
	RuleChainId    string `json:"rule_chain_id" form:"rule_chain_id"`                                          // 引用的可视化规则链ID
	Remark         string `json:"remark" form:"remark" binding:"max=255"`                                      // 备注
	IsStop         uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`                                  // 是否停用
}

type ProductEditReq struct {
	ID             string `json:"id" form:"id" binding:"required"`                                             // 主键
	Name           string `json:"name" form:"name" binding:"required,min=1,max=100"`                           // 产品名称
	CategoryId     string `json:"category_id" form:"category_id" binding:"required"`                           // 分类ID
	DeviceType     string `json:"device_type" form:"device_type" binding:"required,oneof=DEVICE GATEWAY EDGE"` // 节点类型
	Protocol       string `json:"protocol" form:"protocol" binding:"required,oneof=mqtt tcp http ws"`          // 接入协议
	OfflineTimeout int    `json:"offline_timeout" form:"offline_timeout" binding:"required,min=1,max=86400"`   // 离线判定时长(秒)，产品级必填
	ThingModel     string `json:"thing_model" form:"thing_model"`                                              // 物模型JSON
	RuleChainId    string `json:"rule_chain_id" form:"rule_chain_id"`                                          // 引用的可视化规则链ID
	Remark         string `json:"remark" form:"remark" binding:"max=255"`                                      // 备注
	IsStop         uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`                                  // 是否停用
}

type ProductDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type ProductDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type ProductResp struct {
	ID             string      `json:"id"`              // 主键
	Name           string      `json:"name"`            // 产品名称
	CategoryId     string      `json:"category_id"`     // 分类ID
	CategoryName   string      `json:"category_name"`   // 分类名称
	DeviceType     string      `json:"device_type"`     // 节点类型: DEVICE=直连设备, GATEWAY=网关设备, EDGE=边缘主机
	Protocol       string      `json:"protocol"`        // 接入协议
	OfflineTimeout int         `json:"offline_timeout"` // 离线判定时长(秒)，作为其下设备未自定义时的取值
	ThingModel     string      `json:"thing_model"`     // 物模型JSON
	RuleChainId    string      `json:"rule_chain_id"`   // 引用的可视化规则链ID
	RuleChainName  string      `json:"rule_chain_name"` // 规则链名称(回填)
	Remark         string      `json:"remark"`          // 备注
	IsStop         uint8       `json:"is_stop"`         // 是否停用
	CreateTime     x_null.Time `json:"create_time"`     // 创建时间
	UpdateTime     x_null.Time `json:"update_time"`     // 更新时间
}
