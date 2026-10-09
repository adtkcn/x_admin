package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// ProductHandler 产品控制器
type ProductHandler struct{}

// @Summary		产品列表
// @Description	分页获取产品列表
// @Tags			iot_product-产品
// @Param			token		header		string																	true	"token"
// @Param			pageNo		query		int																		false	"页码"
// @Param			pageSize	query		int																		false	"每页数量"
// @Param			name		query		string																	false	"产品名称"
// @Param			category_id	query		string																	false	"分类ID(含子分类)"
// @Param			is_stop		query		string																	false	"是否停用: 0正常 1停用，不传=全部"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=iot_schema.ProductResp}}	"成功"
// @Router			/api/admin/iot/product/list [get]
func (h ProductHandler) List(c *gin.Context) {
	var page request.PageReq
	var req iot_schema.ProductListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.ProductService.List(page, req)
	response.JSON(c, res, err)
}

// @Summary		产品所有
// @Description	获取所有启用产品(供设备表单下拉)
// @Tags			iot_product-产品
// @Param			token	header		string												true	"token"
// @Success		200		{object}	response.Response{data=[]iot_schema.ProductResp}	"成功"
// @Router			/api/admin/iot/product/all [get]
func (h ProductHandler) All(c *gin.Context) {
	res, err := iot_service.ProductService.All()
	response.JSON(c, res, err)
}

// @Summary		产品详情
// @Description	获取产品详情
// @Tags			iot_product-产品
// @Param			token	header		string											true	"token"
// @Param			id		query		string											true	"主键"
// @Success		200		{object}	response.Response{data=iot_schema.ProductResp}	"成功"
// @Router			/api/admin/iot/product/detail [get]
func (h ProductHandler) Detail(c *gin.Context) {
	var req iot_schema.ProductDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.ProductService.Detail(req.ID)
	response.JSON(c, res, err)
}

// @Summary		产品新增
// @Description	新增产品
// @Tags			iot_product-产品
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.ProductAddReq	true	"新增参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/product/add [post]
func (h ProductHandler) Add(c *gin.Context) {
	var req iot_schema.ProductAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductService.Add(req))
}

// @Summary		产品编辑
// @Description	编辑产品
// @Tags			iot_product-产品
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.ProductEditReq	true	"编辑参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/product/edit [post]
func (h ProductHandler) Edit(c *gin.Context) {
	var req iot_schema.ProductEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductService.Edit(req))
}

// @Summary		产品删除
// @Description	删除产品
// @Tags			iot_product-产品
// @Param			token	header		string						true	"token"
// @Param			data	body		iot_schema.ProductDelReq	true	"删除参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/product/del [post]
func (h ProductHandler) Del(c *gin.Context) {
	var req iot_schema.ProductDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductService.Del(req.ID))
}
