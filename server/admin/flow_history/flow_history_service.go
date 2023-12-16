package flow_history

import (
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"

	"gorm.io/gorm"
)

type IFlowHistoryService interface {
	List(page request.PageReq, listReq FlowHistoryListReq) (res response.PageResp, e error)
	ListAll() (res []FlowHistoryResp, e error)

	Detail(id int) (res FlowHistoryResp, e error)
	Add(addReq FlowHistoryAddReq) (e error)
	Edit(editReq FlowHistoryEditReq) (e error)
	Del(id int) (e error)

	GetNextNode(nextNode NextNodeReq) (e error)
}

var Service = NewFlowHistoryService()

// NewFlowHistoryService 初始化
func NewFlowHistoryService() IFlowHistoryService {
	db := core.GetDB()
	return &flowHistoryService{db: db}
}

// flowHistoryService 流程历史服务实现类
type flowHistoryService struct {
	db *gorm.DB
}

// List 流程历史列表
func (Service flowHistoryService) List(page request.PageReq, listReq FlowHistoryListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	dbModel := Service.db.Model(&model.FlowHistory{})
	if listReq.ApplyId > 0 {
		dbModel = dbModel.Where("apply_id = ?", listReq.ApplyId)
	}
	if listReq.TemplateId > 0 {
		dbModel = dbModel.Where("template_id = ?", listReq.TemplateId)
	}
	if listReq.ApplyUserId > 0 {
		dbModel = dbModel.Where("apply_user_id = ?", listReq.ApplyUserId)
	}
	if listReq.ApplyUserNickname != "" {
		dbModel = dbModel.Where("apply_user_nickname like ?", "%"+listReq.ApplyUserNickname+"%")
	}
	if listReq.ApproverId > 0 {
		dbModel = dbModel.Where("approver_id = ?", listReq.ApproverId)
	}
	if listReq.ApproverNickname != "" {
		dbModel = dbModel.Where("approver_nickname like ?", "%"+listReq.ApproverNickname+"%")
	}
	if listReq.NodeId != "" {
		dbModel = dbModel.Where("node_id = ?", listReq.NodeId)
	}
	if listReq.FormValue != "" {
		dbModel = dbModel.Where("form_value = ?", listReq.FormValue)
	}
	if listReq.PassStatus > 0 {
		dbModel = dbModel.Where("pass_status = ?", listReq.PassStatus)
	}
	if listReq.PassRemark != "" {
		dbModel = dbModel.Where("pass_remark = ?", listReq.PassRemark)
	}
	// 总数
	var count int64
	err := dbModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []model.FlowHistory
	err = dbModel.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	resps := []FlowHistoryResp{}
	response.Copy(&resps, objs)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resps,
	}, nil
}

// ListAll 流程历史列表
func (Service flowHistoryService) ListAll() (res []FlowHistoryResp, e error) {
	var objs model.FlowHistory

	err := Service.db.Find(&objs).Error
	if e = response.CheckErr(err, "ListAll Find err"); e != nil {
		return
	}
	response.Copy(&res, objs)
	return res, nil
}

// Detail 流程历史详情
func (Service flowHistoryService) Detail(id int) (res FlowHistoryResp, e error) {
	var obj model.FlowHistory
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

// Add 流程历史新增
func (Service flowHistoryService) Add(addReq FlowHistoryAddReq) (e error) {
	var obj model.FlowHistory
	response.Copy(&obj, addReq)
	err := Service.db.Create(&obj).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 流程历史编辑
func (Service flowHistoryService) Edit(editReq FlowHistoryEditReq) (e error) {
	var obj model.FlowHistory
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

// Del 流程历史删除
func (Service flowHistoryService) Del(id int) (e error) {
	var obj model.FlowHistory
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

/**
 * 获取下一个流程
 */
func (Service flowHistoryService) GetNextNode(nextNode NextNodeReq) (e error) {
	//
	return e
}
