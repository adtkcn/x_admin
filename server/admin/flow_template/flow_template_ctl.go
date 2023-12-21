package flow_template

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type FlowTemplateHandler struct{}

// @Summary	流程模板列表
// @Tags		flow_template-流程模板
// @Produce	json
// @Param		Token			header		string				true	"token"
// @Param		PageNo			query		int					true	"页码"
// @Param		PageSize		query		int					true	"每页数量"
// @Param		flowName		query		string				false	"流程名称"
// @Param		flowGroup		query		int					false	"流程分类"
// @Param		flowRemark		query		string				false	"流程描述"
// @Param		flowFormData	query		string				false	"表单配置"
// @Param		flowProcessData	query		string				false	"流程配置"
// @Success	200				{object}	[]FlowTemplateResp	"成功"
// @Failure	400				{object}	string				"请求错误"
// @Router		/api/flow_template/list [get]
func (hd FlowTemplateHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq FlowTemplateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程模板列表-所有
// @Tags		flow_template-流程模板
// @Router		/api/flow_template/listAll [get]
func (hd FlowTemplateHandler) ListAll(c *gin.Context) {
	res, err := Service.ListAll()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程模板详情
// @Tags		flow_template-流程模板
// @Produce	json
// @Param		Token	header		string				true	"token"
// @Param		id		query		int					false	"历史id"
// @Success	200		{object}	FlowTemplateResp	"成功"
// @Router		/api/flow_template/detail [get]
func (hd FlowTemplateHandler) Detail(c *gin.Context) {
	var detailReq FlowTemplateDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	流程模板新增
// @Tags		flow_template-流程模板
// @Produce	json
// @Param		Token			header		string				true	"token"
// @Param		flowName		body		string				false	"流程名称"
// @Param		flowGroup		body		int					false	"流程分类"
// @Param		flowRemark		body		string				false	"流程描述"
// @Param		flowFormData	body		string				false	"表单配置"
// @Param		flowProcessData	body		string				false	"流程配置"
// @Success	200				{object}	response.RespType	"成功"
// @Router		/api/flow_template/add [post]
func (hd FlowTemplateHandler) Add(c *gin.Context) {
	var addReq FlowTemplateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// @Summary	流程模板编辑
// @Tags		flow_template-流程模板
// @Produce	json
// @Param		Token			header		string				true	"token"
// @Param		id				body		int					false	"."
// @Param		flowName		body		string				false	"流程名称"
// @Param		flowGroup		body		int					false	"流程分类"
// @Param		flowRemark		body		string				false	"流程描述"
// @Param		flowFormData	body		string				false	"表单配置"
// @Param		flowProcessData	body		string				false	"流程配置"
// @Success	200				{object}	response.RespType	"成功"
// @Router		/api/flow_template/edit [post]
func (hd FlowTemplateHandler) Edit(c *gin.Context) {
	var editReq FlowTemplateEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// @Summary	流程模板删除
// @Tags		flow_template-流程模板
// @Produce	json
// @Param		Token	header		string				true	"token"
// @Param		id		body		int					false	"历史id"
// @Success	200		{object}	response.RespType	"成功"
// @Router		/api/flow_template/del [post]
func (hd FlowTemplateHandler) Del(c *gin.Context) {
	var delReq FlowTemplateDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}
