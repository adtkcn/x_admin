package dept

import "x_admin/core"

//SystemAuthDeptListReq 部门列表参数
type SystemAuthDeptListReq struct {
	Name   string `form:"name"`                                     // 部门名称
	IsStop int8   `form:"isStop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

//SystemAuthDeptDetailReq 部门详情参数
type SystemAuthDeptDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthDeptAddReq 部门新增参数
type SystemAuthDeptAddReq struct {
	Pid    uint   `form:"pid" binding:"gte=0"`                   // 部门父级
	Name   string `form:"name" binding:"required,min=1,max=100"` // 部门名称
	DutyId int    `form:"dutyId" binding:"omitempty"`            // 负责人id
	Duty   string `form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `form:"mobile" binding:"omitempty,len=11"`     // 联系电话
	IsStop uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort   int    `form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

//SystemAuthDeptEditReq 部门编辑参数
type SystemAuthDeptEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`            // 主键
	Pid    uint   `form:"pid" binding:"gte=0"`                   // 部门父级
	Name   string `form:"name" binding:"required,min=1,max=100"` // 部门名称
	DutyId int    `form:"dutyId" binding:"omitempty"`            // 负责人id
	Duty   string `form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `form:"mobile" binding:"omitempty,len=11"`     // 联系电话
	IsStop uint8  `form:"isStop" binding:"oneof=0 1"`            // 是否停用: [0=否, 1=是]
	Sort   int    `form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

//SystemAuthDeptDelReq 部门删除参数
type SystemAuthDeptDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

//SystemAuthDeptResp 系统部门返回信息
type SystemAuthDeptResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Pid        uint        `json:"pid" structs:"pid"`               // 部门父级
	Name       string      `json:"name" structs:"name"`             // 部门名称
	DutyId     int         `json:"dutyId" structs:"dutyId"`         // 负责人id
	Duty       string      `json:"duty" structs:"duty"`             // 负责人
	Mobile     string      `json:"mobile" structs:"mobile"`         // 联系电话
	Sort       uint16      `json:"sort" structs:"sort"`             // 排序编号
	IsStop     uint8       `json:"isStop" structs:"isStop"`         // 是否停用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
