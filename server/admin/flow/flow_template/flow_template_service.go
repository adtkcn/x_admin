package flow_template

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

type IFlowTemplateService interface {
	List(page request.PageReq, listReq FlowTemplateListReq) (res response.PageResp, e error)
	ListAll() (res []FlowTemplateResp, e error)
	Detail(id int) (res FlowTemplateResp, e error)
	Add(addReq FlowTemplateAddReq) (e error)
	Edit(editReq FlowTemplateEditReq) (e error)
	Del(id int) (e error)
}

var Service = NewFlowTemplateService()

// NewFlowTemplateService 初始化
func NewFlowTemplateService() *flowTemplateService {
	db := core.GetDB()
	return &flowTemplateService{db: db}
}

// flowTemplateService 流程模板服务实现类
type flowTemplateService struct {
	db *gorm.DB
}

// List 流程模板列表
func (Service flowTemplateService) List(page request.PageReq, listReq FlowTemplateListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := Service.db.Model(&model.FlowTemplate{})
	if listReq.FlowName != "" {
		dbModel = dbModel.Where("flow_name like ?", "%"+listReq.FlowName+"%")
	}
	if listReq.FlowGroup > 0 {
		dbModel = dbModel.Where("flow_group = ?", listReq.FlowGroup)
	}
	if listReq.FlowRemark != "" {
		dbModel = dbModel.Where("flow_remark = ?", listReq.FlowRemark)
	}
	if listReq.FlowFormData != "" {
		dbModel = dbModel.Where("flow_form_data = ?", listReq.FlowFormData)
	}
	if listReq.FlowProcessData != "" {
		dbModel = dbModel.Where("flow_process_data = ?", listReq.FlowProcessData)
	}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.FlowTemplate
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []FlowTemplateResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 流程模板列表
func (Service flowTemplateService) ListAll() (res []FlowTemplateResp, e error) {
	var objs []model.FlowTemplate
	err := Service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 流程模板详情
func (Service flowTemplateService) Detail(id int) (res FlowTemplateResp, e error) {
	var obj model.FlowTemplate
	err := Service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 流程模板新增
func (Service flowTemplateService) Add(addReq FlowTemplateAddReq) (e error) {
	var obj model.FlowTemplate
	response.Copy(&obj, addReq)
	err := Service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 流程模板编辑
func (Service flowTemplateService) Edit(editReq FlowTemplateEditReq) (e error) {
	var obj model.FlowTemplate
	err := Service.db.Where("id = ?", editReq.Id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = Service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "Edit Updates err")
	return
}

// Del 流程模板删除
func (Service flowTemplateService) Del(id int) (e error) {
	var obj model.FlowTemplate
	err := Service.db.Where("id = ?", id).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 删除
	err = Service.db.Delete(&obj).Error
	e = response.CheckErr(err, "Del Delete err")
	return
}
