package model

//FlowTemplate 流程模板实体
type FlowTemplate struct {
	Id                  int    `gorm:"primarykey;comment:''"` //
	FlowName            string `gorm:"comment:'流程名称'"`        // 流程名称
	FlowGroup           int    `gorm:"comment:'流程分类'"`        // 流程分类
	FlowRemark          string `gorm:"comment:'流程描述'"`        // 流程描述
	FlowFormData        string `gorm:"comment:'表单配置'"`        // 表单配置
	FlowProcessData     string `gorm:"comment:'流程配置'"`        // 流程配置
	FlowProcessDataList string `gorm:"comment:'流程配置list数据'"`  // 流程配置list数据
}
