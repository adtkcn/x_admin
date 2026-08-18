package flow_controller

import (
	"fmt"
	"x_admin/app/schema/flow_schema"
	"x_admin/app/service/flow_service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type FlowHistoryHandler struct {
}

// @Summary	流程历史列表
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token				header		string																			true	"token"
// @Param		pageNo				query		int																				true	"页码"
// @Param		pageSize			query		int																				true	"每页数量"
// @Param		apply_id			query		string																			false	"申请id"
// @Param		template_id			query		string																			false	"模板id"
// @Param		apply_user_id		query		string																			false	"申请人id"
// @Param		apply_user_nickname	query		string																			false	"申请人昵称"
// @Param		approver_id			query		string																			false	"审批人id"
// @Param		approver_nickname	query		string																			false	"审批用户昵称"
// @Param		node_id				query		string																			false	"节点"
// @Param		node_label			query		string																			false	"节点名称"
// @Param		node_type			query		string																			false	"节点类型"
// @Param		form_value			query		string																			false	"表单值"
// @Param		pass_status			query		int																				false	"通过状态：1待处理，2通过，3拒绝"
// @Param		pass_remark			query		string																			false	"通过备注"
// @Success	200					{object}	response.Response{data=response.PageResp{lists=[]flow_schema.FlowHistoryResp}}	"成功"
// @Router		/api/admin/flow/flow_history/list [get]
func (hd FlowHistoryHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq = flow_schema.FlowHistoryListReq{
		PassStatus: -9999,
		IsShow:     -9999,
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := flow_service.HistoryService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史列表-所有
// @Tags		flow_history-流程历史
// @Produce	json
// @Success	200	{object}	response.Response{data=flow_schema.FlowHistoryResp}	"成功"
// @Router		/api/admin/flow/flow_history/list_all [get]
func (hd FlowHistoryHandler) ListAll(c *gin.Context) {
	var listReq flow_schema.FlowHistoryListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := flow_service.HistoryService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史详情
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token	header		string												true	"token"
// @Param		id		query		string												false	"历史id"
// @Success	200		{object}	response.Response{data=flow_schema.FlowHistoryResp}	"成功"
// @Router		/api/admin/flow/flow_history/detail [get]
func (hd FlowHistoryHandler) Detail(c *gin.Context) {
	var detailReq flow_schema.FlowHistoryDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := flow_service.HistoryService.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史新增
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token				header		string				true	"token"
// @Param		apply_id			body		string				false	"申请id"
// @Param		template_id			body		string				false	"模板id"
// @Param		apply_user_id		body		string				false	"申请人id"
// @Param		apply_user_nickname	body		string				false	"申请人昵称"
// @Param		approver_id			body		string				false	"审批人id"
// @Param		approver_nickname	body		string				false	"审批用户昵称"
// @Param		node_id				body		string				false	"节点"
// @Param		node_label			body		string				false	"节点名称"
// @Param		node_type			body		string				false	"节点类型"
// @Param		form_value			body		string				false	"表单值"
// @Param		pass_status			body		int					false	"通过状态：1待处理，2通过，3拒绝"
// @Param		pass_remark			body		string				false	"通过备注"
// @Success	200					{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_history/add [post]
func (hd FlowHistoryHandler) Add(c *gin.Context) {
	var addReq flow_schema.FlowHistoryAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, flow_service.HistoryService.Add(addReq))
}

// @Summary	流程历史编辑
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token				header		string				true	"token"
// @Param		id					body		string				false	"历史id"
// @Param		apply_id			body		string				false	"申请id"
// @Param		template_id			body		string				false	"模板id"
// @Param		apply_user_id		body		string				false	"申请人id"
// @Param		apply_user_nickname	body		string				false	"申请人昵称"
// @Param		approver_id			body		string				false	"审批人id"
// @Param		approver_nickname	body		string				false	"审批用户昵称"
// @Param		node_id				body		string				false	"节点"
// @Param		node_label			body		string				false	"节点名称"
// @Param		node_type			body		string				false	"节点类型"
// @Param		form_value			body		string				false	"表单值"
// @Param		pass_status			body		int					false	"通过状态：1待处理，2通过，3拒绝"
// @Param		pass_remark			body		string				false	"通过备注"
// @Success	200					{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_history/edit [post]
func (hd FlowHistoryHandler) Edit(c *gin.Context) {
	var editReq flow_schema.FlowHistoryEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, flow_service.HistoryService.Edit(editReq))
}

// @Summary	流程历史删除
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		int					false	"历史id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_history/del [post]
func (hd FlowHistoryHandler) Del(c *gin.Context) {
	var delReq flow_schema.FlowHistoryDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, flow_service.HistoryService.Del(delReq.Id))
}

// @Summary	已处理页面隐藏（软删除）
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		string				true	"历史id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_history/done_hidden [post]
func (hd FlowHistoryHandler) DoneHidden(c *gin.Context) {
	var delReq flow_schema.FlowHistoryDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, flow_service.HistoryService.DoneHidden(delReq.Id))
}

// 提交申请,通过审批
//
//	@Summary	流程历史-通过审批
//	@Tags		flow_history-流程历史
//	@Produce	json
//	@Param		token				header		string				true	"token"
//	@Param		apply_id			body		string				true	"申请id"
//	@Param		next_node_admin_id	body		string				false	"下一个节点的审批用户id"
//	@Param		pass_remark			body		string				false	"通过备注"
//	@Success	200					{object}	response.Response	"成功"
//	@Router		/api/admin/flow/flow_history/pass [post]
func (hd FlowHistoryHandler) Pass(c *gin.Context) {
	var pass flow_schema.PassReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &pass)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	err := flow_service.HistoryService.Pass(pass, adminId)

	response.CheckAndRespWithData(c, nil, err)
}

// 拒绝审批
//
//	@Summary	流程历史-拒绝审批
//	@Tags		flow_history-流程历史
//	@Produce	json
//	@Param		token		header		string				true	"token"
//	@Param		apply_id	body		string				true	"申请id"
//	@Param		history_id	body		string				true	"审批节点id"
//	@Param		remark		body		string				false	"备注"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/flow/flow_history/back [post]
func (hd FlowHistoryHandler) Back(c *gin.Context) {
	var back flow_schema.BackReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &back)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	err := flow_service.HistoryService.Back(back, adminId)
	fmt.Println(err)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary	获取下一个审批节点，中间可能存在通知节点和网关
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		apply_id	body		string				true	"申请id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_history/next_node [post]
func (hd FlowHistoryHandler) NextNode(c *gin.Context) {
	var nextNode flow_schema.NextNodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &nextNode)) {
		return
	}
	res, _, _, err := flow_service.HistoryService.GetNextNode(nextNode.ApplyId)
	response.CheckAndRespWithData(c, res, err)
}

// 获取节点的可审批用户
//
//	@Summary	流程历史-获取节点可审批用户
//	@Tags		flow_history-流程历史
//	@Produce	json
//	@Param		token		header		string				true	"token"
//	@Param		apply_id	body		string				true	"申请id"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/flow/flow_history/get_approver [post]
func (hd FlowHistoryHandler) GetApprover(c *gin.Context) {
	var nextNode flow_schema.NextNodeReq
	// var node FlowTree
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &nextNode)) {
		return
	}
	res, err := flow_service.HistoryService.GetApprover(nextNode.ApplyId)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Ok(c, res)
}
