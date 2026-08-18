package setting_schema

import "github.com/adtkcn/x_null"

// SettingDictDataResp 字典数据返回信息
type SettingDictDataResp struct {
	ID         string      `json:"id"`          // 主键
	TypeId     string      `json:"type_id"`     // 类型
	Name       string      `json:"name"`        // 键
	Value      string      `json:"value"`       // 值
	Color      string      `json:"color"`       // 颜色
	Remark     string      `json:"remark"`      // 备注
	Sort       uint16      `json:"sort"`        // 排序
	Status     uint8       `json:"status"`      // 状态: [0=停用, 1=禁用]
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}

// SettingDictDataListReq 字典数据列表参数
type SettingDictDataListReq struct {
	DictType string `json:"dict_type" form:"dict_type" binding:"max=200"`           // 字典类型
	Name     string `json:"name" form:"name" binding:"max=100"`                     // 键
	Value    string `json:"value" form:"value" binding:"max=200"`                   // 值
	Status   int8   `json:"status" form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

// SettingDictDataDetailReq 字典数据详情参数
type SettingDictDataDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SettingDictDataAddReq 字典数据新增参数
type SettingDictDataAddReq struct {
	TypeId string `json:"type_id" form:"type_id" binding:"required"`              // 类型
	Name   string `json:"name" form:"name" binding:"required,max=100"`            // 键
	Value  string `json:"value" form:"value" binding:"required,max=200"`          // 值
	Color  string `json:"color" form:"color"`                                     // 颜色
	Remark string `json:"remark" form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `json:"sort" form:"sort" binding:"gte=0"`                       // 排序
	Status int8   `json:"status" form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

// SettingDictDataEditReq 字典数据编辑参数
type SettingDictDataEditReq struct {
	ID     string `json:"id" form:"id" binding:"required"`                        // 主键
	TypeId string `json:"type_id" form:"type_id" binding:"required"`              // 类型
	Name   string `json:"name" form:"name" binding:"required,max=100"`            // 键
	Value  string `json:"value" form:"value" binding:"required,max=200"`          // 值
	Color  string `json:"color" form:"color"`                                     // 颜色
	Remark string `json:"remark" form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `json:"sort" form:"sort" binding:"gte=0"`                       // 排序
	Status int8   `json:"status" form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

// SettingDictDataDelReq 字典数据删除参数
type SettingDictDataDelReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 主键列表
}
