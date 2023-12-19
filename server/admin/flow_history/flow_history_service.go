package flow_history

import (
	"encoding/json"
	"x_admin/admin/flow_apply"
	"x_admin/admin/system/admin"
	"x_admin/core"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/model"
	"x_admin/model/system_model"
	"x_admin/util"

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
func NewFlowHistoryService() *flowHistoryService {
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
func (Service flowHistoryService) ListAll(listReq FlowHistoryListReq) (res []FlowHistoryResp, e error) {

	// 查询
	dbModel := Service.db.Model(&model.FlowHistory{})
	if listReq.ApplyId > 0 {
		dbModel = dbModel.Where("apply_id = ?", listReq.ApplyId)
	}
	// 数据
	var objs []model.FlowHistory
	err := dbModel.Find(&objs).Error
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
* 获取节点的审批用户
 */
func (Service flowHistoryService) GetApprover(node FlowTree) (res []admin.SystemAuthAdminResp, e error) {
	var userId = node.UserId
	var deptId = node.DeptId
	var postId = node.PostId
	adminTbName := core.DBTableName(&system_model.SystemAuthAdmin{})

	adminModel := Service.db.Table(adminTbName+" AS admin").Where("admin.is_delete = ?", 0)

	dept := map[string]interface{}{}
	if deptId > 0 {
		dept["admin.dept_id"] = deptId
		// adminModel.Or("admin.dept_id =?", deptId)
	}
	if postId > 0 {
		dept["admin.post_id"] = postId
		// adminModel.Or("admin.post_id =?", postId)
	}

	var where = Service.db.Where(dept)
	if userId > 0 {
		where.Or("admin.id =?", userId)
	}

	// 数据
	var adminResp []admin.SystemAuthAdminResp
	err := adminModel.Where(where).Find(&adminResp).Error
	if e = response.CheckErr(err, "获取审批用户失败"); e != nil {
		return
	}
	for i := 0; i < len(adminResp); i++ {
		adminResp[i].Avatar = util.UrlUtil.ToAbsoluteUrl(adminResp[i].Avatar)
		if adminResp[i].ID == 1 {
			adminResp[i].Role = "系统管理员"
		}
	}
	return adminResp, nil
}

func (Service flowHistoryService) Pass(nextNode NextNodeReq) (e error) {
	nextNodes, applyDetail, err := Service.GetNextNode(nextNode)

	isEnd := false
	if err == nil {
		for _, v := range nextNodes {
			// if v.Type == "bpmn:exclusiveGateway" {
			// 这里网关不用处理，顶多加一条历史记录
			// }
			if v.Type == "bpmn:serviceTask" {
				// 发邮件之类的，待完善
			} else if v.Type == "bpmn:userTask" {
				var addReq = FlowHistoryAddReq{}
				addReq.ApplyId = applyDetail.Id
				addReq.FormValue = applyDetail.FormValue
				addReq.NodeId = v.Id
				addReq.ApproverId = nextNode.NextNodeAdminId

				// addReq.ApproverNickname = applyDetail.ApproverNickname

				addReq.TemplateId = applyDetail.TemplateId
				addReq.ApplyUserId = applyDetail.ApplyUserId
				addReq.ApplyUserNickname = applyDetail.ApplyUserNickname
				addReq.PassStatus = 1
				err = Service.Add(addReq)
			} else if v.Type == "bpmn:endEvent" {
				isEnd = true
				var addReq = FlowHistoryAddReq{}
				addReq.ApplyId = applyDetail.Id
				addReq.FormValue = applyDetail.FormValue
				addReq.NodeId = v.Id
				addReq.ApproverId = 0

				// addReq.ApproverNickname = applyDetail.ApproverNickname

				addReq.TemplateId = applyDetail.TemplateId
				addReq.ApplyUserId = applyDetail.ApplyUserId
				addReq.ApplyUserNickname = applyDetail.ApplyUserNickname
				addReq.PassStatus = 2
				err = Service.Add(addReq)
			}
		}
	}
	// 待提交或者有结束节点，修改申请状态
	if applyDetail.Status == 1 || isEnd {
		status := 2 //审批中
		if isEnd {
			status = 3 //审批通过
		}
		// 更改状态
		err = flow_apply.Service.Edit(flow_apply.FlowApplyEditReq{
			Id:     nextNode.ApplyId,
			Status: status,
		})
		if err != nil {
			return err
		}
	}
	return err
}

/**
 * 获取下一批流程，直到审批或结束节点
 */
func (Service flowHistoryService) GetNextNode(nextNode NextNodeReq) (res []FlowTree, apply flow_apply.FlowApplyResp, e error) {
	var applyDetail, err = flow_apply.Service.Detail(nextNode.ApplyId)
	if e = response.CheckErr(err, "获取审批申请失败"); e != nil {
		return
	}
	// start
	var flowTree []FlowTree
	json.Unmarshal([]byte(applyDetail.FlowProcessDataList), &flowTree)
	var formValue map[string]interface{}
	json.Unmarshal([]byte(nextNode.FormValue), &formValue)

	var next []FlowTree

	if nextNode.CurrentNodeId == "" {
		for _, v := range flowTree {
			if v.Type == "bpmn:startEvent" {
				next = *v.Children
				break
			}
		}
	} else {
		for _, v := range flowTree {
			if v.Id == nextNode.CurrentNodeId {
				next = *v.Children
				break
			}
		}
	}
	var nextNodes []FlowTree
	res = DeepNextNode(nextNodes, &next, formValue)
	return res, applyDetail, e
}

// 返回节点数组，最后一个节点为用户或结束节点
func DeepNextNode(nextNodes []FlowTree, flowTree *[]FlowTree, formValue map[string]interface{}) []FlowTree {
	for _, v := range *flowTree {
		if v.Type == "bpmn:startEvent" {
			// 开始节点
			child := DeepNextNode(nextNodes, v.Children, formValue)
			nextNodes = append(nextNodes, child...)
			break
		} else if v.Type == "bpmn:exclusiveGateway" {
			// 网关

			// 判断formValue值，决定是不是递归这个网关
			child := DeepNextNode(nextNodes, v.Children, formValue)
			nextNodes = append(nextNodes, v)
			nextNodes = append(nextNodes, child...)
			break
		} else if v.Type == "bpmn:serviceTask" {
			// 系统服务
			child := DeepNextNode(nextNodes, v.Children, formValue)
			nextNodes = append(nextNodes, v)
			nextNodes = append(nextNodes, child...)
		} else if v.Type == "bpmn:userTask" {
			//用户节点
			nextNodes = append(nextNodes, v)
			break
		} else if v.Type == "bpmn:endEvent" {
			// 结束节点
			nextNodes = append(nextNodes, v)
			break
		}
	}
	return nextNodes
}
