package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// ProductCategoryHandler 产品分类控制器
type ProductCategoryHandler struct{}

// @Summary		产品分类所有
// @Description	获取所有产品分类(树形,不分页)
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string														true	"token"
// @Success		200		{object}	response.Response{data=[]iot_schema.ProductCategoryResp}	"成功"
// @Router			/api/admin/iot/productCategory/all [get]
func (h ProductCategoryHandler) All(c *gin.Context) {
	res, err := iot_service.ProductCategoryService.All()
	response.JSON(c, res, err)
}

// @Summary		产品分类详情
// @Description	获取产品分类详情
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string													true	"token"
// @Param			id		query		string													true	"主键"
// @Success		200		{object}	response.Response{data=iot_schema.ProductCategoryResp}	"成功"
// @Router			/api/admin/iot/productCategory/detail [get]
func (h ProductCategoryHandler) Detail(c *gin.Context) {
	var req iot_schema.ProductCategoryDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.ProductCategoryService.Detail(req.ID)
	response.JSON(c, res, err)
}

// @Summary		产品分类新增
// @Description	新增产品分类
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.ProductCategoryAddReq	true	"新增参数"
// @Success		200		{object}	response.Response					"成功"
// @Router			/api/admin/iot/productCategory/add [post]
func (h ProductCategoryHandler) Add(c *gin.Context) {
	var req iot_schema.ProductCategoryAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductCategoryService.Add(req))
}

// @Summary		产品分类编辑
// @Description	编辑产品分类
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.ProductCategoryEditReq	true	"编辑参数"
// @Success		200		{object}	response.Response					"成功"
// @Router			/api/admin/iot/productCategory/edit [post]
func (h ProductCategoryHandler) Edit(c *gin.Context) {
	var req iot_schema.ProductCategoryEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductCategoryService.Edit(req))
}

// @Summary		产品分类删除
// @Description	删除产品分类
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.ProductCategoryDelReq	true	"删除参数"
// @Success		200		{object}	response.Response					"成功"
// @Router			/api/admin/iot/productCategory/del [post]
func (h ProductCategoryHandler) Del(c *gin.Context) {
	var req iot_schema.ProductCategoryDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductCategoryService.Del(req.ID))
}

// @Summary		产品分类排序
// @Description	同级产品分类拖拽排序
// @Tags			iot_productCategory-产品分类
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.ProductCategorySortReq	true	"排序参数"
// @Success		200		{object}	response.Response					"成功"
// @Router			/api/admin/iot/productCategory/sort [post]
func (h ProductCategoryHandler) Sort(c *gin.Context) {
	var req iot_schema.ProductCategorySortReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.ProductCategoryService.Sort(req.Ids))
}
