package system_schema

import "github.com/adtkcn/x_null"

// SystemAuthDeptListReq 部门列表参数
type SystemAuthDeptListReq struct {
	Name   string `json:"name" form:"name"`                                       // 部门名称
	IsStop int8   `json:"isStop" form:"isStop,default=-1" binding:"oneof=-1 0 1"` // 是否停用: [0=否, 1=是]
}

// SystemAuthDeptDetailReq 部门详情参数
type SystemAuthDeptDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthDeptAddReq 部门新增参数
type SystemAuthDeptAddReq struct {
	Pid    string `json:"pid" form:"pid" binding:""`                         // 部门父级
	Name   string `json:"name" form:"name" binding:"required,min=1,max=100"` // 部门名称
	DutyId string `json:"duty_id" form:"duty_id" binding:"omitempty"`        // 负责人id
	Duty   string `json:"duty" form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `json:"mobile" form:"mobile" binding:"omitempty,len=11"`   // 联系电话
	IsStop uint8  `json:"isStop" form:"isStop" binding:"oneof=0 1"`          // 是否停用: [0=否, 1=是]
	Sort   int    `json:"sort" form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

// SystemAuthDeptEditReq 部门编辑参数
type SystemAuthDeptEditReq struct {
	ID     string `json:"id" form:"id" binding:"required"`                   // 主键
	Pid    string `json:"pid" form:"pid" binding:"gte=0"`                    // 部门父级
	Name   string `json:"name" form:"name" binding:"required,min=1,max=100"` // 部门名称
	DutyId string `json:"duty_id" form:"duty_id" binding:"omitempty"`        // 负责人id
	Duty   string `json:"duty" form:"duty" binding:"omitempty,min=1,max=30"` // 负责人
	Mobile string `json:"mobile" form:"mobile" binding:"omitempty,len=11"`   // 联系电话
	IsStop uint8  `json:"isStop" form:"isStop" binding:"oneof=0 1"`          // 是否停用: [0=否, 1=是]
	Sort   int    `json:"sort" form:"sort" binding:"gte=0,lte=9999"`         // 排序编号
}

// SystemAuthDeptDelReq 部门删除参数
type SystemAuthDeptDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemAuthDeptSortReq 部门拖拽排序参数
type SystemAuthDeptSortReq struct {
	Ids []string `json:"ids" form:"ids" binding:"required"` // 同层级部门id，按拖拽后顺序排列
}

// SystemAuthDeptResp 系统部门返回信息
type SystemAuthDeptResp struct {
	ID         string      `json:"id"`          // 主键
	Pid        string      `json:"pid"`         // 部门父级
	Name       string      `json:"name"`        // 部门名称
	DutyId     string      `json:"duty_id"`     // 负责人id
	Duty       string      `json:"duty"`        // 负责人
	Mobile     string      `json:"mobile"`      // 联系电话
	Sort       uint16      `json:"sort"`        // 排序编号
	IsStop     uint8       `json:"is_stop"`     // 是否停用: [0=否, 1=是]
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}
