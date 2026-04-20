package system_controller

import (
	"net/http"
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

// Self 管理员信息
func (ah AdminHandler) Self(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := system_service.AdminService.Self(adminId)
	response.CheckAndRespWithData(c, res, err)
}

func (ah AdminHandler) ExportFile(c *gin.Context) {
	var listReq system_schema.SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.ExportFile(listReq)

	if err != nil {
		response.Fail(c, "查询导出失败")
		return
	}
	f, err := excel2.Export(res, system_service.AdminService.GetExcelCol(), "Sheet1", "用户信息")
	if err != nil {
		response.Fail(c, "导出失败")
		return
	}
	excel2.DownLoadExcel("用户信息"+time.Now().Format("20060102-150405"), c.Writer, f)
}

// 导入文件
func (ah AdminHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []system_schema.SystemAuthAdminResp{}
	err = excel2.GetExcelData(file, &importList, system_service.AdminService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = system_service.AdminService.ImportFile(importList)
	response.CheckAndRespWithData(c, nil, err)
}

// list 管理员列表
func (ah AdminHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq system_schema.SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// ListAll 所有管理员列表
func (ah AdminHandler) ListAll(c *gin.Context) {

	var listReq system_schema.SystemAuthAdminListReq

	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.AdminService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 管理员详细
func (ah AdminHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthAdminDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.AdminService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 管理员新增
func (ah AdminHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := system_service.AdminService.Add(addReq)
	response.CheckAndRespWithData(c, nil, err)
}

// edit 管理员编辑
func (ah AdminHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := system_service.AdminService.Edit(c, editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// upInfo 管理员更新
func (ah AdminHandler) UpInfo(c *gin.Context) {
	var updateReq system_schema.SystemAuthAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	err := system_service.AdminService.Update(
		c, updateReq, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, nil, err)
}

// del 管理员删除
func (ah AdminHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthAdminDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := system_service.AdminService.Del(c, delReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}

// disable 管理员状态切换
func (ah AdminHandler) Disable(c *gin.Context) {
	var disableReq system_schema.SystemAuthAdminDisableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	err := system_service.AdminService.Disable(c, disableReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		获取部门的用户
// @Description	获取部门的用户
// @Tags			管理员
// @Param			deptId	path		int					true	"部门id"
// @Success		200		{object}	response.Response	"{"code": 200, "data": []}"
// @Router			/system/admin/ListByDeptId/{deptId} [get]
func (ah AdminHandler) ListByDeptId(c *gin.Context) {
	deptId, bool := c.GetQuery("deptId")
	if !bool {
		response.Fail(c, "deptId不能为空")
		return
	}

	res, err := system_service.AdminService.ListByDeptId(deptId)
	response.CheckAndRespWithData(c, res, err)
}
