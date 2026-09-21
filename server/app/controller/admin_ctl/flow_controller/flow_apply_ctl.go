package flow_controller

import (
	. "x_admin/app/schema/flow_schema"
	"x_admin/app/service/flow_service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type FlowApplyHandler struct{}

// @Summary	申请流程列表
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		token					header		string																true	"token"
// @Param		pageNo					query		int																	true	"页码"
// @Param		pageSize				query		int																	true	"每页数量"
// @Param		template_id				query		string																false	"模板"
// @Param		apply_user_id			query		string																false	"申请人id"
// @Param		apply_user_nickname		query		string																false	"申请人昵称"
// @Param		flow_name				query		string																false	"流程名称"
// @Param		flow_group				query		int																	false	"流程分类"
// @Param		flow_remark				query		string																false	"流程描述"
// @Param		flow_form_data			query		string																false	"表单配置"
// @Param		flow_process_data		query		string																false	"流程配置"
// @Param		flow_process_data_list	query		string																false	"流程配置list数据"
// @Param		form_value				query		string																false	"表单值"
// @Param		status					query		int																	false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200						{object}	response.Response{data=response.PageResp{lists=[]FlowApplyResp}}	"成功"
// @Router		/api/admin/flow/flow_apply/list [get]
func (hd FlowApplyHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq FlowApplyListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := flow_service.ApplyService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary	申请流程详情
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		token	header		string									true	"token"
// @Param		id		query		string									false	"申请id"
// @Success	200		{object}	response.Response{data=FlowApplyResp}	"成功"
// @Router		/api/admin/flow/flow_apply/detail [get]
func (hd FlowApplyHandler) Detail(c *gin.Context) {
	var detailReq FlowApplyDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := flow_service.ApplyService.Detail(detailReq.Id)
	response.JSON(c, res, err)
}

// @Summary	申请流程新增
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		token				header		string				true	"token"
// @Param		template_id			body		string				false	"模板"
// @Param		apply_user_id		body		string				false	"申请人id"
// @Param		apply_user_nickname	body		string				false	"申请人昵称"
// @Param		flow_name			body		string				false	"流程名称"
// @Param		form_value			body		string				false	"表单值"
// @Param		status				body		int					false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200					{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_apply/add [post]
func (hd FlowApplyHandler) Add(c *gin.Context) {
	var addReq FlowApplyAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}

	var nickname = config.AdminConfig.GetNickname(c)
	var adminId = config.AdminConfig.GetAdminId(c)
	addReq.ApplyUserNickname = nickname
	addReq.ApplyUserId = adminId
	addReq.Status = 1

	response.JSON(c, nil, flow_service.ApplyService.Add(addReq))
}

// @Summary	申请流程编辑
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		id			body		string				false	"申请id"
// @Param		flow_name	body		string				false	"流程名称"
// @Param		form_value	body		string				false	"表单值"
// @Param		status		body		int					false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_apply/edit [post]
func (hd FlowApplyHandler) Edit(c *gin.Context) {
	var editReq FlowApplyEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.JSON(c, nil, flow_service.ApplyService.Edit(editReq))
}

// @Summary	申请流程删除
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		string				true	"申请id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/flow/flow_apply/del [post]
func (hd FlowApplyHandler) Del(c *gin.Context) {
	var delReq FlowApplyDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.JSON(c, nil, flow_service.ApplyService.Del(delReq.Id))
}
