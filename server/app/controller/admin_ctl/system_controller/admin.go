package system_controller

import (
	"time"
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"

	"x_admin/util/excel2"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// AdminHandler 管理员控制器
type AdminHandler struct{}

// @Summary		管理员信息
// @Description	获取当前管理员信息
// @Tags			system_admin-管理员
// @Param			token	header		string															true	"token"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthAdminSelfResp}	"成功"
// @Router			/api/admin/system/admin/self [get]
func (ah AdminHandler) Self(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := system_service.AdminService.Self(adminId)
	response.JSON(c, res, err)
}

// @Summary		导出管理员文件
// @Description	导出管理员列表到Excel
// @Tags			system_admin-管理员
// @Param			token		header	string	true	"token"
// @Param			email		query	string	false	"邮箱(账号)"
// @Param			nickname	query	string	false	"昵称"
// @Param			roleId		query	string	false	"角色ID"
// @Success		200			"文件流"
// @Router			/api/admin/system/admin/export [get]
func (ah AdminHandler) ExportFile(c *gin.Context) {
	var listReq system_schema.SystemAuthAdminListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.ExportFile(listReq)

	if err != nil {
		response.Fail(c, response.CheckErr(err, "查询导出失败"))
		return
	}
	f, err := excel2.Export(res, system_service.AdminService.GetExcelCol(), "Sheet1", "用户信息")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "导出失败"))
		return
	}
	excel2.DownLoadExcel("用户信息"+time.Now().Format("20060102-150405"), c.Writer, f)
}

// @Summary		导入管理员文件
// @Description	从Excel导入管理员数据
// @Tags			system_admin-管理员
// @Param			token	header		string				true	"token"
// @Param			file	formData	file				true	"Excel文件"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/import [post]
func (ah AdminHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件不存在"))
		return
	}
	defer file.Close()
	importList := []system_schema.SystemAuthAdminResp{}
	err = excel2.GetExcelData(file, &importList, system_service.AdminService.GetExcelCol())
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件解析失败"))
		return
	}
	err = system_service.AdminService.ImportFile(importList)
	response.JSON(c, nil, err)
}

// @Summary		管理员列表
// @Description	获取管理员列表
// @Tags			system_admin-管理员
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Param			email		query		string																				false	"邮箱(账号)"
// @Param			nickname	query		string																				false	"昵称"
// @Param			roleId		query		string																				false	"角色ID"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemAuthAdminResp}}	"成功"
// @Router			/api/admin/system/admin/list [get]
func (ah AdminHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq system_schema.SystemAuthAdminListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary		所有管理员列表
// @Description	获取所有管理员列表(不分页)
// @Tags			system_admin-管理员
// @Param			token		header		string														true	"token"
// @Param			email		query		string														false	"邮箱(账号)"
// @Param			nickname	query		string														false	"昵称"
// @Param			role_id		query		string														false	"角色ID"
// @Success		200			{object}	response.Response{data=[]system_schema.SystemAuthAdminResp}	"成功"
// @Router			/api/admin/system/admin/list_all [get]
func (ah AdminHandler) ListAll(c *gin.Context) {

	var listReq system_schema.SystemAuthAdminListReq

	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.ListAll(listReq)
	response.JSON(c, res, err)
}

// @Summary		管理员详情
// @Description	获取管理员详情
// @Tags			system_admin-管理员
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthAdminResp}	"成功"
// @Router			/api/admin/system/admin/detail [get]
func (ah AdminHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthAdminDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.AdminService.Detail(detailReq.ID)
	response.JSON(c, res, err)
}

// @Summary		管理员新增
// @Description	新增管理员
// @Tags			system_admin-管理员
// @Param			token		header		string				true	"token"
// @Param			dept_id		body		string				false	"部门ID"
// @Param			post_id		body		string				false	"岗位ID"
// @Param			role_ids	body		[]string			false	"角色ID列表"
// @Param			email		body		string				true	"邮箱(账号)"
// @Param			nickname	body		string				true	"昵称"
// @Param			password	body		string				true	"密码"
// @Param			avatar		body		string				false	"头像"
// @Param			sort		body		int					false	"排序"
// @Param			is_disable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/add [post]
func (ah AdminHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthAdminAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := system_service.AdminService.Add(addReq)
	response.JSON(c, nil, err)
}

// @Summary		管理员编辑
// @Description	编辑管理员信息
// @Tags			system_admin-管理员
// @Param			token		header		string				true	"token"
// @Param			id			body		string				true	"主键"
// @Param			dept_id		body		string				false	"部门ID"
// @Param			post_id		body		string				false	"岗位ID"
// @Param			role_ids	body		[]string			false	"角色ID列表"
// @Param			email		body		string				true	"邮箱(账号)"
// @Param			nickname	body		string				true	"昵称"
// @Param			password	body		string				false	"密码"
// @Param			avatar		body		string				false	"头像"
// @Param			sort		body		int					false	"排序"
// @Param			is_disable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/edit [post]
func (ah AdminHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthAdminEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := system_service.AdminService.Edit(c, editReq)
	response.JSON(c, nil, err)
}

// @Summary		发送邮箱验证码
// @Description	修改邮箱时发送验证码到新邮箱
// @Tags			system_admin-管理员
// @Param			token	header		string				true	"token"
// @Param			email	body		string				true	"目标邮箱"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/sendEmailCode [post]
func (ah AdminHandler) SendEmailCode(c *gin.Context) {
	var req system_schema.SystemAuthAdminSendEmailCodeReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := system_service.AdminService.SendBindEmailCode(config.AdminConfig.GetAdminId(c), req.Email)
	response.JSON(c, nil, err)
}

// @Summary		管理员更新信息
// @Description	当前管理员更新自己的信息
// @Tags			system_admin-管理员
// @Param			token			header		string				true	"token"
// @Param			nickname		body		string				true	"昵称"
// @Param			avatar			body		string				false	"头像"
// @Param			email			body		string				false	"邮箱"
// @Param			email_code		body		string				false	"邮箱验证码"
// @Param			password		body		string				false	"密码"
// @Param			curr_password	body		string				false	"当前密码"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/upInfo [post]
func (ah AdminHandler) UpInfo(c *gin.Context) {
	var updateReq system_schema.SystemAuthAdminUpdateReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	err := system_service.AdminService.Update(
		c, updateReq, config.AdminConfig.GetAdminId(c))
	response.JSON(c, nil, err)
}

// @Summary		管理员删除
// @Description	删除管理员
// @Tags			system_admin-管理员
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/del [post]
func (ah AdminHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthAdminDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := system_service.AdminService.Del(c, delReq.ID)
	response.JSON(c, nil, err)
}

// @Summary		管理员状态切换
// @Description	切换管理员禁用状态
// @Tags			system_admin-管理员
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/admin/disable [post]
func (ah AdminHandler) Disable(c *gin.Context) {
	var disableReq system_schema.SystemAuthAdminDisableReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	err := system_service.AdminService.Disable(c, disableReq.ID)
	response.JSON(c, nil, err)
}

// @Summary		获取部门的用户
// @Description	获取部门的用户
// @Tags			system_admin-管理员
// @Param			token	header		string														true	"token"
// @Param			dept_id	query		string														true	"部门id"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthAdminResp}	"{"code": 200, "data": []}"
// @Router			/api/admin/system/admin/ListByDeptId [get]
func (ah AdminHandler) ListByDeptId(c *gin.Context) {
	dept_id, bool := c.GetQuery("dept_id")
	if !bool {
		response.FailMsg(c, "deptId不能为空")
		return
	}

	res, err := system_service.AdminService.ListByDeptId(dept_id)
	response.JSON(c, res, err)
}
