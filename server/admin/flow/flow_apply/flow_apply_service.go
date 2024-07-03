package flow_apply

import (
	"errors"
	"x_admin/admin/flow/flow_template"
	"x_admin/core"
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

var Service = NewFlowApplyService()

// NewFlowApplyService 初始化
func NewFlowApplyService() *flowApplyService {
	db := core.GetDB()
	return &flowApplyService{db: db}
}

// flowApplyService 申请流程服务实现类
type flowApplyService struct {
	db *gorm.DB
}

// List 申请流程列表
func (service flowApplyService) List(page request.PageReq, listReq FlowApplyListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := service.db.Model(&model.FlowApply{})
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
	dbModel = dbModel.Where("is_delete = ?", 0)
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "列表总数获取失败"); e != nil {
		return
	}
	// 数据
	var modelList []model.FlowApply
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&modelList).Error
	if e = response.CheckErr(err, "列表获取失败"); e != nil {
		return
	}
	result := []FlowApplyResp{}
	response.Copy(&result, modelList)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    result,
	}, nil
}

// Detail 申请流程详情
func (service flowApplyService) Detail(id int) (res FlowApplyResp, e error) {
	var obj model.FlowApply
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "详情获取失败"); e != nil {
		return
	}
	response.Copy(&res, obj)
	return
}

// Add 申请流程新增
func (service flowApplyService) Add(addReq FlowApplyAddReq) (e error) {
	var obj model.FlowApply
	var flow_template_resp, err = flow_template.Service.Detail(addReq.TemplateId)
	if e = response.CheckErrDBNotRecord(err, "模板不存在!"); e != nil {
		return
	}
	response.Copy(&obj, addReq)
	// obj.FlowName = flow_template_resp.FlowName
	obj.FlowGroup = flow_template_resp.FlowGroup
	obj.FlowRemark = flow_template_resp.FlowRemark
	obj.FlowFormData = flow_template_resp.FlowFormData
	obj.FlowProcessData = flow_template_resp.FlowProcessData
	obj.FlowProcessDataList = flow_template_resp.FlowProcessDataList

	err = service.db.Create(&obj).Error
	e = response.CheckErr(err, "添加失败")
	return
}

// Edit 申请流程编辑
func (service flowApplyService) Edit(editReq FlowApplyEditReq) (e error) {
	var obj model.FlowApply
	err := service.db.Where("id = ? AND is_delete = ?", editReq.Id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待编辑数据查找失败"); e != nil {
		return
	}
	// 更新
	response.Copy(&obj, editReq)
	err = service.db.Model(&obj).Updates(obj).Error
	e = response.CheckErr(err, "编辑失败")
	return
}

// Del 申请流程删除
func (service flowApplyService) Del(id int) (e error) {
	var obj model.FlowApply
	err := service.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&obj).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "数据不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "待删除数据查找失败"); e != nil {
		return
	}
	if obj.Status == 2 {
		// 审批中不允许删除
		e = errors.New("审批中不允许删除")
		return
	}
	// 删除
	obj.IsDelete = 1
	obj.DeleteTime = core.NowTime()
	err = service.db.Save(&obj).Error
	e = response.CheckErr(err, "Del Save err")
	return
}
