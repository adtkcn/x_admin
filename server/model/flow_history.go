package model

//FlowHistory 流程历史实体
type FlowHistory struct {
	Id int `gorm:"primarykey;comment:'历史id'"` // 历史id

	ApplyId int `gorm:"comment:'申请id'"` // 申请id

	TemplateId int `gorm:"comment:'模板id'"` // 模板id

	ApplyUserId int `gorm:"comment:'申请人id'"` // 申请人id

	ApplyUserNickname string `gorm:"comment:'申请人昵称'"` // 申请人昵称

	ApproverId int `gorm:"comment:'审批人id'"` // 审批人id

	ApproverNickname string `gorm:"comment:'审批用户昵称'"` // 审批用户昵称

	NodeId    string `gorm:"comment:'节点'"`   // 节点
	NodeType  string `gorm:"comment:'节点类型'"` // 节点类型
	NodeLabel string `gorm:"comment:'节点名称'"` //节点名称

	FormValue string `gorm:"comment:'表单值'"` // 表单值

	PassStatus int `gorm:"comment:'通过状态：1待处理，2通过，3拒绝'"` // 通过状态：1待处理，2通过，3拒绝

	PassRemark string `gorm:"comment:'通过备注'"` // 通过备注

	UpdateTime int64 `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间

	CreateTime int64 `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间

	DeleteTime int64 `gorm:"comment:'删除时间'"` // 删除时间

}
