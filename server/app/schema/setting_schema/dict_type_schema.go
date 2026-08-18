package setting_schema

import "github.com/adtkcn/x_null"

// SettingDictTypeListReq 字典类型新增参数
type SettingDictTypeListReq struct {
	DictName   string `json:"dict_name" form:"dict_name" binding:"max=200"`                   // 字典名称
	DictType   string `json:"dict_type" form:"dict_type" binding:"max=200"`                   // 字典类型
	DictStatus int8   `json:"dict_status" form:"dict_status,default=-1" binding:"oneof=-1 0 1"` // 字典状态: 0/1
}

// SettingDictTypeDetailReq 字典类型详情参数
type SettingDictTypeDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SettingDictTypeAddReq 字典类型新增参数
type SettingDictTypeAddReq struct {
	DictName   string `json:"dict_name" form:"dict_name" binding:"required,max=200"`     // 字典名称
	DictType   string `json:"dict_type" form:"dict_type" binding:"required,max=200"`     // 字典类型
	DictRemark string `json:"dict_remark" form:"dict_remark" binding:"max=200"`            // 字典备注
	DictStatus int8   `json:"dict_status" form:"dict_status" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

// SettingDictTypeEditReq 字典类型编辑参数
type SettingDictTypeEditReq struct {
	ID         string `json:"id" form:"id" binding:"required"`                    // 主键
	DictName   string `json:"dict_name" form:"dict_name" binding:"required,max=200"`     // 字典名称
	DictType   string `json:"dict_type" form:"dict_type" binding:"required,max=200"`     // 字典类型
	DictRemark string `json:"dict_remark" form:"dict_remark" binding:"max=200"`            // 字典备注
	DictStatus int8   `json:"dict_status" form:"dict_status" binding:"required,oneof=0 1"` // 字典状态: 0/1
}

// SettingDictTypeDelReq 字典类型删除参数
type SettingDictTypeDelReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// SettingDictTypeResp 字典类型返回信息
type SettingDictTypeResp struct {
	ID         string      `json:"id"`          // 主键
	DictName   string      `json:"dict_name"`   // 字典名称
	DictType   string      `json:"dict_type"`   // 字典类型
	DictRemark string      `json:"dict_remark"` // 字典备注
	DictStatus uint8       `json:"dict_status"` // 字典状态
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
