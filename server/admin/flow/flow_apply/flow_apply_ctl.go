package flow_apply

import (
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
// @Param		Token				header		string			true	"token"
// @Param		PageNo				query		int				true	"页码"
// @Param		PageSize			query		int				true	"每页数量"
// @Param		templateId			query		int				false	"模板"
// @Param		applyUserId			query		int				false	"申请人id"
// @Param		applyUserNickname	query		string			false	"申请人昵称"
// @Param		flowName			query		string			false	"流程名称"
// @Param		flowGroup			query		int				false	"流程分类"
// @Param		flowRemark			query		string			false	"流程描述"
// @Param		flowFormData		query		string			false	"表单配置"
// @Param		flowProcessData		query		string			false	"流程配置"
// @Param		status				query		int				false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200					{object}	[]FlowApplyResp	"成功"
// @Failure	400					{object}	string			"请求错误"
// @Router		/api/flow_apply/list [get]
func (hd FlowApplyHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq FlowApplyListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	申请流程详情
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		Token	header		string			true	"token"
// @Param		id		query		int				false	"申请id"
// @Success	200		{object}	FlowApplyResp	"成功"
// @Router		/api/flow_apply/detail [get]
func (hd FlowApplyHandler) Detail(c *gin.Context) {
	var detailReq FlowApplyDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	申请流程新增
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		Token				header		string				true	"token"
// @Param		templateId			body		int					false	"模板"
// @Param		applyUserId			body		int					false	"申请人id"
// @Param		applyUserNickname	body		string				false	"申请人昵称"
// @Param		flowName			body		string				false	"流程名称"
// @Param		flowGroup			body		int					false	"流程分类"
// @Param		flowRemark			body		string				false	"流程描述"
// @Param		flowFormData		body		string				false	"表单配置"
// @Param		flowProcessData		body		string				false	"流程配置"
// @Param		status				body		int					false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200					{object}	response.RespType	"成功"
// @Router		/api/flow_apply/add [post]
func (hd FlowApplyHandler) Add(c *gin.Context) {
	var addReq FlowApplyAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}

	var Nickname = config.AdminConfig.GetNickname(c)
	var AdminId = config.AdminConfig.GetAdminId(c)
	addReq.ApplyUserNickname = Nickname
	addReq.ApplyUserId = int(AdminId)
	addReq.Status = 1

	response.CheckAndResp(c, Service.Add(addReq))
}

// @Summary	申请流程编辑
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		Token				header		string				true	"token"
// @Param		id					body		int					false	"申请id"
// @Param		templateId			body		int					false	"模板"
// @Param		applyUserId			body		int					false	"申请人id"
// @Param		applyUserNickname	body		string				false	"申请人昵称"
// @Param		flowName			body		string				false	"流程名称"
// @Param		flowGroup			body		int					false	"流程分类"
// @Param		flowRemark			body		string				false	"流程描述"
// @Param		flowFormData		body		string				false	"表单配置"
// @Param		flowProcessData		body		string				false	"流程配置"
// @Param		status				body		int					false	"状态：1待提交，2审批中，3审批完成，4审批失败"
// @Success	200					{object}	response.RespType	"成功"
// @Router		/api/flow_apply/edit [post]
func (hd FlowApplyHandler) Edit(c *gin.Context) {
	var editReq FlowApplyEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// @Summary	申请流程删除
// @Tags		flow_apply-申请流程
// @Produce	json
// @Param		Token	header		string				true	"token"
// @Param		id		body		int					false	"申请id"
// @Success	200		{object}	response.RespType	"成功"
// @Router		/api/flow_apply/del [post]
func (hd FlowApplyHandler) Del(c *gin.Context) {
	var delReq FlowApplyDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}
