package model

//FlowApply 申请流程实体
type FlowApply struct {
	Id                  int    `gorm:"primarykey;comment:''"`              //
	TemplateId          int    `gorm:"comment:'模板'"`                       // 模板
	ApplyUserId         int    `gorm:"comment:'申请人id'"`                    // 申请人id
	ApplyUserNickname   string `gorm:"comment:'申请人昵称'"`                    // 申请人昵称
	FlowName            string `gorm:"comment:'流程名称'"`                     // 流程名称
	FlowGroup           int    `gorm:"comment:'流程分类'"`                     // 流程分类
	FlowRemark          string `gorm:"comment:'流程描述'"`                     // 流程描述
	FlowFormData        string `gorm:"comment:'表单配置'"`                     // 表单配置
	FlowProcessData     string `gorm:"comment:'流程配置'"`                     // 流程配置
	FlowProcessDataList string `gorm:"comment:'流程配置list数据'"`               // 流程配置list数据
	FormValue           string `gorm:"comment:'表单值'"`                      // 表单值
	Status              int    `gorm:"comment:'状态：1待提交，2审批中，3审批完成，4审批失败'"` // 状态：0待提交，1审批中，2审批完成，3审批失败
	IsDelete            int    `gorm:"comment:'是否删除'"`                     // 是否删除
	UpdateTime          int64  `gorm:"autoUpdateTime;comment:'更新时间'"`      // 更新时间
	CreateTime          int64  `gorm:"autoCreateTime;comment:'创建时间'"`      // 创建时间
	DeleteTime          int64  `gorm:"comment:'删除时间'"`                     // 删除时间
}
