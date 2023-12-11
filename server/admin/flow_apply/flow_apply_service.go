package flow_apply

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

type IFlowApplyService interface {
	List(page request.PageReq, listReq FlowApplyListReq) (res response.PageResp, e error)
	Detail(id int) (res FlowApplyResp, e error)
	Add(addReq FlowApplyAddReq) (e error)
	Edit(editReq FlowApplyEditReq) (e error)
	Del(id int) (e error)
}

// NewFlowApplyService 初始化
func NewFlowApplyService(db *gorm.DB) IFlowApplyService {
	return &flowApplyService{db: db}
}

// flowApplyService 申请流程服务实现类
type flowApplyService struct {
	db *gorm.DB
}

// List 申请流程列表
func (Service flowApplyService) List(page request.PageReq, listReq FlowApplyListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := Service.db.Model(&model.FlowApply{})
	if listReq.TemplateId > 0 {
		dbModel = dbModel.Where("template_id = ?", listReq.TemplateId)
	}
	if listReq.ApplyUserId > 0 {
		dbModel = dbModel.Where("apply_user_id = ?", listReq.ApplyUserId)
	}
	if listReq.ApplyUserNickname != "" {
		dbModel = dbModel.Where("apply_user_nickname like ?", "%"+listReq.ApplyUserNickname+"%")
	}
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
	if listReq.Status > 0 {
		dbModel = dbModel.Where("status = ?", listReq.Status)
	}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.FlowApply
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []FlowApplyResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// Detail 申请流程详情
func (Service flowApplyService) Detail(id int) (res FlowApplyResp, e error) {
	var obj model.FlowApply
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

// Add 申请流程新增
func (Service flowApplyService) Add(addReq FlowApplyAddReq) (e error) {
	var obj model.FlowApply

	response.Copy(&obj, addReq)
	err := Service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 申请流程编辑
func (Service flowApplyService) Edit(editReq FlowApplyEditReq) (e error) {
	var obj model.FlowApply
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

// Del 申请流程删除
func (Service flowApplyService) Del(id int) (e error) {
	var obj model.FlowApply
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
