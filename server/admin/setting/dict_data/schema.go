package dict_data

import "x_admin/core"

//SettingDictDataResp 字典数据返回信息
type SettingDictDataResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	TypeId     uint        `json:"typeId" structs:"typeId"`         // 类型
	Name       string      `json:"name" structs:"name"`             // 键
	Value      string      `json:"value" structs:"value"`           // 值
	Color      string      `json:"color" structs:"color"`           // 颜色
	Remark     string      `json:"remark" structs:"remark"`         // 备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 排序
	Status     uint8       `json:"status" structs:"status"`         // 状态: [0=停用, 1=禁用]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

//SettingDictDataListReq 字典数据列表参数
type SettingDictDataListReq struct {
	DictType string `form:"dictType" binding:"max=200"`               // 字典类型
	Name     string `form:"name" binding:"max=100"`                   // 键
	Value    string `form:"value" binding:"max=200"`                  // 值
	Status   int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataDetailReq 字典数据详情参数
type SettingDictDataDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SettingDictDataAddReq 字典数据新增参数
type SettingDictDataAddReq struct {
	TypeId uint   `form:"typeId" binding:"required,gt=0"`           // 类型
	Name   string `form:"name" binding:"required,max=100"`          // 键
	Value  string `form:"value" binding:"required,max=200"`         // 值
	Color  string `form:"color"`                                    // 颜色
	Remark string `form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `form:"sort" binding:"gte=0"`                     // 排序
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataEditReq 字典数据编辑参数
type SettingDictDataEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`               // 主键
	TypeId uint   `form:"typeId" binding:"required,gte=0"`          // 类型
	Name   string `form:"name" binding:"required,max=100"`          // 键
	Value  string `form:"value" binding:"required,max=200"`         // 值
	Color  string `form:"color"`                                    // 颜色
	Remark string `form:"remark" binding:"max=200"`                 // 备注
	Sort   int    `form:"sort" binding:"gte=0"`                     // 排序
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: 0=停用,1=启用
}

//SettingDictDataDelReq 字典数据删除参数
type SettingDictDataDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}
