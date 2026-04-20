package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DeptHandler 部门控制器
type DeptHandler struct{}

// All 部门所有
func (dh DeptHandler) All(c *gin.Context) {
	res, err := system_service.DeptService.All()
	response.CheckAndRespWithData(c, res, err)
}

// List 部门列表
func (dh DeptHandler) List(c *gin.Context) {
	var listReq system_schema.SystemAuthDeptListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.DeptService.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 部门详情
func (dh DeptHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthDeptDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.DeptService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 部门新增
func (dh DeptHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthDeptAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := system_service.DeptService.Add(addReq)
	response.CheckAndRespWithData(c, nil, err)
}

// Edit 部门编辑
func (dh DeptHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthDeptEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := system_service.DeptService.Edit(editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// Del 部门删除
func (dh DeptHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthDeptDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := system_service.DeptService.Del(delReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}
