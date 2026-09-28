package iot_schema

import (
	"github.com/adtkcn/x_null"
)

// ============ 设备区域（树） ============
type DeviceSpaceAddReq struct {
	Pid     string `json:"pid" form:"pid"`                                    // 上级区域
	Name    string `json:"name" form:"name" binding:"required,min=1,max=100"` // 区域名称
	Sort    uint16 `json:"sort" form:"sort" binding:"lte=9999"`               // 排序
	GeoJson string `json:"geojson" form:"geojson"`                            // 区域范围GeoJSON
	Remark  string `json:"remark" form:"remark" binding:"max=255"`            // 备注
}

type DeviceSpaceEditReq struct {
	ID      string `json:"id" form:"id" binding:"required"`                   // 主键
	Pid     string `json:"pid" form:"pid"`                                    // 上级区域
	Name    string `json:"name" form:"name" binding:"required,min=1,max=100"` // 区域名称
	Sort    uint16 `json:"sort" form:"sort" binding:"lte=9999"`               // 排序
	GeoJson string `json:"geojson" form:"geojson"`                            // 区域范围GeoJSON
	Remark  string `json:"remark" form:"remark" binding:"max=255"`            // 备注
}

type DeviceSpaceDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type DeviceSpaceDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type DeviceSpaceSortReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 同级区域id，按拖拽后顺序排列
}

type DeviceSpaceResp struct {
	ID         string      `json:"id"`          // 主键
	Pid        string      `json:"pid"`         // 上级区域
	Name       string      `json:"name"`        // 区域名称
	Sort       uint16      `json:"sort"`        // 排序
	GeoJson    string      `json:"geojson"`     // 区域范围GeoJSON
	Remark     string      `json:"remark"`      // 备注
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
