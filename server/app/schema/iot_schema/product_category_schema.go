package iot_schema

import (
	"github.com/adtkcn/x_null"
)

// ============ 产品分类（树） ============
type ProductCategoryAddReq struct {
	Pid    string `json:"pid" form:"pid"`                                    // 上级分类
	Name   string `json:"name" form:"name" binding:"required,min=1,max=100"` // 分类名称
	Sort   uint16 `json:"sort" form:"sort" binding:"lte=9999"`               // 排序
	Remark string `json:"remark" form:"remark" binding:"max=255"`            // 备注
}

type ProductCategoryEditReq struct {
	ID     string `json:"id" form:"id" binding:"required"`                   // 主键
	Pid    string `json:"pid" form:"pid"`                                    // 上级分类
	Name   string `json:"name" form:"name" binding:"required,min=1,max=100"` // 分类名称
	Sort   uint16 `json:"sort" form:"sort" binding:"lte=9999"`               // 排序
	Remark string `json:"remark" form:"remark" binding:"max=255"`            // 备注
}

type ProductCategoryDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type ProductCategoryDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

type ProductCategorySortReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 同级分类id，按拖拽后顺序排列
}

type ProductCategoryResp struct {
	ID         string      `json:"id"`          // 主键
	Pid        string      `json:"pid"`         // 上级分类
	Name       string      `json:"name"`        // 分类名称
	Sort       uint16      `json:"sort"`        // 排序
	Remark     string      `json:"remark"`      // 备注
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
