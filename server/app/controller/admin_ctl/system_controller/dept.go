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

// @Summary		部门所有
// @Description	获取所有部门列表(不分页)
// @Tags			system_dept-部门
// @Param			token	header		string														true	"token"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthDeptResp}	"成功"
// @Router			/api/admin/system/dept/all [get]
func (dh DeptHandler) All(c *gin.Context) {
	res, err := system_service.DeptService.All()
	response.JSON(c, res, err)
}

// @Summary		部门列表
// @Description	获取部门列表
// @Tags			system_dept-部门
// @Param			token	header		string														true	"token"
// @Param			name	query		string														false	"部门名称"
// @Param			is_stop	query		int8														false	"是否停用: [0=否, 1=是]"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthDeptResp}	"成功"
// @Router			/api/admin/system/dept/list [get]
func (dh DeptHandler) List(c *gin.Context) {
	var listReq system_schema.SystemAuthDeptListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.DeptService.List(listReq)
	response.JSON(c, res, err)
}

// @Summary		部门详情
// @Description	获取部门详情
// @Tags			system_dept-部门
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthDeptResp}	"成功"
// @Router			/api/admin/system/dept/detail [get]
func (dh DeptHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthDeptDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.DeptService.Detail(detailReq.ID)
	response.JSON(c, res, err)
}

// @Summary		部门新增
// @Description	新增部门
// @Tags			system_dept-部门
// @Param			token	header		string				true	"token"
// @Param			pid		body		string				false	"部门父级"
// @Param			name	body		string				true	"部门名称"
// @Param			duty_id	body		string				false	"负责人id"
// @Param			duty	body		string				false	"负责人"
// @Param			mobile	body		string				false	"联系电话"
// @Param			is_stop	body		uint8				false	"是否停用: [0=否, 1=是]"
// @Param			sort	body		int					false	"排序编号"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/dept/add [post]
func (dh DeptHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthDeptAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := system_service.DeptService.Add(addReq)
	response.JSON(c, nil, err)
}

// @Summary		部门编辑
// @Description	编辑部门
// @Tags			system_dept-部门
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Param			pid		body		string				false	"部门父级"
// @Param			name	body		string				true	"部门名称"
// @Param			duty_id	body		string				false	"负责人id"
// @Param			duty	body		string				false	"负责人"
// @Param			mobile	body		string				false	"联系电话"
// @Param			is_stop	body		uint8				false	"是否停用: [0=否, 1=是]"
// @Param			sort	body		int					false	"排序编号"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/dept/edit [post]
func (dh DeptHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthDeptEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := system_service.DeptService.Edit(editReq)
	response.JSON(c, nil, err)
}

// @Summary		部门删除
// @Description	删除部门
// @Tags			system_dept-部门
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/dept/del [post]
func (dh DeptHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthDeptDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := system_service.DeptService.Del(delReq.ID)
	response.JSON(c, nil, err)
}

// @Summary		部门拖拽排序
// @Description	同级部门拖拽排序，按传入 id 顺序持久化
// @Tags			system_dept-部门
// @Param			token	header		string		true	"token"
// @Param			ids		body		[]string	true	"部门id顺序"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/dept/sort [post]
func (dh DeptHandler) Sort(c *gin.Context) {
	var sortReq system_schema.SystemAuthDeptSortReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &sortReq)) {
		return
	}
	err := system_service.DeptService.Sort(sortReq.Ids)
	response.JSON(c, nil, err)
}
