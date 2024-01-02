package flow_history

import (
	"fmt"
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
// @Param		Token				header		string				true	"token"
// @Param		PageNo				query		int					true	"页码"
// @Param		PageSize			query		int					true	"每页数量"
// @Param		applyId				query		int					false	"申请id"
// @Param		templateId			query		int					false	"模板id"
// @Param		applyUserId			query		int					false	"申请人id"
// @Param		applyUserNickname	query		string				false	"申请人昵称"
// @Param		approverId			query		int					false	"审批人id"
// @Param		approverNickname	query		string				false	"审批用户昵称"
// @Param		nodeId				query		string				false	"节点"
// @Param		formValue			query		string				false	"表单值"
// @Param		passStatus			query		int					false	"通过状态：0待处理，1通过，2拒绝"
// @Param		passRemark			query		string				false	"通过备注"
// @Success	200					{object}	[]FlowHistoryResp	"成功"
// @Failure	400					{object}	string				"请求错误"
// @Router		/api/flow_history/list [get]
func (hd FlowHistoryHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq = FlowHistoryListReq{
		PassStatus: -9999,
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史列表-所有
// @Tags		flow_history-流程历史
// @Produce	json
// @Success	200	{object}	[]FlowHistoryResp	"成功"
// @Router		/api/flow_history/listAll [get]
func (hd FlowHistoryHandler) ListAll(c *gin.Context) {
	var listReq FlowHistoryListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史详情
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		Token	header		string			true	"token"
// @Param		id		query		int				false	"历史id"
// @Success	200		{object}	FlowHistoryResp	"成功"
// @Router		/api/flow_history/detail [get]
func (hd FlowHistoryHandler) Detail(c *gin.Context) {
	var detailReq FlowHistoryDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程历史新增
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		Token				header		string				true	"token"
// @Param		applyId				body		int					false	"申请id"
// @Param		templateId			body		int					false	"模板id"
// @Param		applyUserId			body		int					false	"申请人id"
// @Param		applyUserNickname	body		string				false	"申请人昵称"
// @Param		approverId			body		int					false	"审批人id"
// @Param		approverNickname	body		string				false	"审批用户昵称"
// @Param		nodeId				body		string				false	"节点"
// @Param		formValue			body		string				false	"表单值"
// @Param		passStatus			body		int					false	"通过状态：0待处理，1通过，2拒绝"
// @Param		passRemark			body		string				false	"通过备注"
// @Success	200					{object}	response.RespType	"成功"
// @Router		/api/flow_history/add [post]
func (hd FlowHistoryHandler) Add(c *gin.Context) {
	var addReq FlowHistoryAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// @Summary	流程历史编辑
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		Token				header		string				true	"token"
// @Param		id					body		int					false	"历史id"
// @Param		applyId				body		int					false	"申请id"
// @Param		templateId			body		int					false	"模板id"
// @Param		applyUserId			body		int					false	"申请人id"
// @Param		applyUserNickname	body		string				false	"申请人昵称"
// @Param		approverId			body		int					false	"审批人id"
// @Param		approverNickname	body		string				false	"审批用户昵称"
// @Param		nodeId				body		string				false	"节点"
// @Param		formValue			body		string				false	"表单值"
// @Param		passStatus			body		int					false	"通过状态：0待处理，1通过，2拒绝"
// @Param		passRemark			body		string				false	"通过备注"
// @Success	200					{object}	response.RespType	"成功"
// @Router		/api/flow_history/edit [post]
func (hd FlowHistoryHandler) Edit(c *gin.Context) {
	var editReq FlowHistoryEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// @Summary	流程历史删除
// @Tags		flow_history-流程历史
// @Produce	json
// @Param		Token	header		string				true	"token"
// @Param		id		body		int					false	"历史id"
// @Success	200		{object}	response.RespType	"成功"
// @Router		/api/flow_history/del [post]
func (hd FlowHistoryHandler) Del(c *gin.Context) {
	var delReq FlowHistoryDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}

// 提交申请,通过审批
//
//	@Router	/api/flow_apply/pass [post]
func (hd FlowHistoryHandler) Pass(c *gin.Context) {
	var pass PassReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &pass)) {
		return
	}
	err := Service.Pass(pass)
	fmt.Println(err)
	response.CheckAndResp(c, err)
}

// 拒绝审批
// @Router	/api/flow_apply/back [post]
func (hd FlowHistoryHandler) Back(c *gin.Context) {
	var back BackReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &back)) {
		return
	}
	err := Service.Back(back)
	fmt.Println(err)
	response.CheckAndResp(c, err)
}

// 获取下一个审批节点，中间可能存在系统任务节点和网关
func (hd FlowHistoryHandler) NextNode(c *gin.Context) {
	var nextNode NextNodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &nextNode)) {
		return
	}

	// response.CheckAndResp(c, Service.GetNextNode(nextNode))
	res, _, _, err := Service.GetNextNode(nextNode.ApplyId)
	response.CheckAndRespWithData(c, res, err)
}

// 获取节点的可审批用户
func (hd FlowHistoryHandler) GetApprover(c *gin.Context) {
	var nextNode NextNodeReq
	// var node FlowTree
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &nextNode)) {
		return
	}

	// response.CheckAndResp(c, Service.GetNextNode(node))
	res, err := Service.GetApprover(nextNode.ApplyId)
	if err != nil {
		response.FailWithMsg(c, response.Failed, err.Error())
		return
	}
	response.OkWithData(c, res)
}

// 同意审批(当前nodeId)

// 拒绝审批，驳回审批
