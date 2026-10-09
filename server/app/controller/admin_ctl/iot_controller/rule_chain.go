package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// RuleChainHandler 可视化规则链控制器
type RuleChainHandler struct{}

// @Summary		规则链列表
// @Description	分页获取可视化规则链列表
// @Tags			iot_rule_chain-规则链
// @Param			token		header		string																		true	"token"
// @Param			pageNo		query		int																			false	"页码"
// @Param			pageSize	query		int																			false	"每页数量"
// @Param			name		query		string																		false	"规则链名称"
// @Param			is_stop		query		string																		false	"是否停用: 0正常 1停用，不传=全部"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=iot_schema.RuleChainResp}}	"成功"
// @Router			/api/admin/iot/ruleChain/list [get]
func (h RuleChainHandler) List(c *gin.Context) {
	var page request.PageReq
	var req iot_schema.RuleChainListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.RuleChainService.List(page, req)
	response.JSON(c, res, err)
}

// @Summary		规则链所有
// @Description	获取全部启用规则链(供产品表单下拉)
// @Tags			iot_rule_chain-规则链
// @Param			token	header		string												true	"token"
// @Success		200		{object}	response.Response{data=[]iot_schema.RuleChainResp}	"成功"
// @Router			/api/admin/iot/ruleChain/all [get]
func (h RuleChainHandler) All(c *gin.Context) {
	res, err := iot_service.RuleChainService.All()
	response.JSON(c, res, err)
}

// @Summary		规则链详情
// @Description	获取规则链详情(含图JSON)
// @Tags			iot_rule_chain-规则链
// @Param			token	header		string												true	"token"
// @Param			id		query		string												true	"主键"
// @Success		200		{object}	response.Response{data=iot_schema.RuleChainResp}	"成功"
// @Router			/api/admin/iot/ruleChain/detail [get]
func (h RuleChainHandler) Detail(c *gin.Context) {
	var req iot_schema.RuleChainDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.RuleChainService.Detail(req.ID)
	response.JSON(c, res, err)
}

// @Summary		规则链新增
// @Description	新增可视化规则链
// @Tags			iot_rule_chain-规则链
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.RuleChainAddReq	true	"新增参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/ruleChain/add [post]
func (h RuleChainHandler) Add(c *gin.Context) {
	var req iot_schema.RuleChainAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.RuleChainService.Add(req))
}

// @Summary		规则链编辑
// @Description	编辑可视化规则链
// @Tags			iot_rule_chain-规则链
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.RuleChainEditReq	true	"编辑参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/ruleChain/edit [post]
func (h RuleChainHandler) Edit(c *gin.Context) {
	var req iot_schema.RuleChainEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.RuleChainService.Edit(req))
}

// @Summary		规则链删除
// @Description	删除可视化规则链(被产品引用时禁止)
// @Tags			iot_rule_chain-规则链
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.RuleChainDelReq	true	"删除参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/ruleChain/del [post]
func (h RuleChainHandler) Del(c *gin.Context) {
	var req iot_schema.RuleChainDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.RuleChainService.Del(req.ID))
}
