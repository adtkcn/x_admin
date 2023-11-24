package dict_type

import "x_admin/core"

//SettingDictTypeListReq 字典类型新增参数
type SettingDictTypeListReq struct {
	DictName   string `form:"dictName" binding:"max=200"`                   // 字典名称
	DictType   string `form:"dictType" binding:"max=200"`                   // 字典类型
	DictStatus int8   `form:"dictStatus,default=-1" binding:"oneof=-1 0 1"` // 字典状态: 0/1
}

//SettingDictTypeDetailReq 字典类型详情参数
type SettingDictTypeDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SettingDictTypeAddReq 字典类型新增参数
type SettingDictTypeAddReq struct {
	DictName   string `form:"dictName" binding:"required,max=200"`     // 字典名称
	DictType   string `form:"dictType" binding:"required,max=200"`     // 字典类型
	DictRemark string `form:"dictRemark" binding:"max=200"`            // 字典备注
	DictStatus int8   `form:"dictStatus" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

//SettingDictTypeEditReq 字典类型编辑参数
type SettingDictTypeEditReq struct {
	ID         uint   `form:"id" binding:"required,gt=0"`              // 主键
	DictName   string `form:"dictName" binding:"required,max=200"`     // 字典名称
	DictType   string `form:"dictType" binding:"required,max=200"`     // 字典类型
	DictRemark string `form:"dictRemark" binding:"max=200"`            // 字典备注
	DictStatus int8   `form:"dictStatus" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

//SettingDictTypeDelReq 字典类型删除参数
type SettingDictTypeDelReq struct {
	Ids []uint `form:"ids" binding:"required"` // 主键列表
}

//SettingDictTypeResp 字典类型返回信息
type SettingDictTypeResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	DictName   string      `json:"dictName" structs:"dictName"`     // 字典名称
	DictType   string      `json:"dictType" structs:"dictType"`     // 字典类型
	DictRemark string      `json:"dictRemark" structs:"dictRemark"` // 字典备注
	DictStatus uint8       `json:"dictStatus" structs:"dictStatus"` // 字典状态
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
