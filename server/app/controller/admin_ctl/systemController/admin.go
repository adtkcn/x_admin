package systemController

import (
	"net/http"
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"

	"x_admin/util/excel"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func AdminRoute(rg *gin.RouterGroup) {

	handle := AdminHandler{}
	notAuth := rg.Group("/system", middleware.LoginAuth())
	notAuth.GET("/admin/self", handle.Self)
	notAuth.POST("/admin/upInfo", middleware.RecordLog("管理员更新"), handle.UpInfo)

	auth := rg.Group("/system", middleware.TokenAuth())

	auth.GET("/admin/list", handle.List)
	auth.GET("/admin/listAll", handle.ListAll)
	auth.GET("/admin/ListByDeptId", handle.ListByDeptId)
	auth.GET("/admin/detail", handle.Detail)
	auth.POST("/admin/add", middleware.RecordLog("管理员新增"), handle.Add)
	auth.POST("/admin/edit", middleware.RecordLog("管理员编辑"), handle.Edit)

	auth.POST("/admin/del", middleware.RecordLog("管理员删除"), handle.Del)
	auth.POST("/admin/disable", middleware.RecordLog("管理员状态切换"), handle.Disable)

	auth.GET("/admin/ExportFile", middleware.RecordLog("管理员导出"), handle.ExportFile)

	auth.POST("/admin/ImportFile", handle.ImportFile)

}

type AdminHandler struct {
	// Service ISystemAuthAdminService
}

// self 管理员信息
func (ah AdminHandler) Self(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := systemService.AdminService.Self(adminId)
	response.CheckAndRespWithData(c, res, err)
}

func (ah AdminHandler) ExportFile(c *gin.Context) {
	var listReq systemSchema.SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := systemService.AdminService.ExportFile(listReq)

	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询导出失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "用户信息", nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("用户信息", c.Writer, f)
	// c.Header("Content-Type", "application/octet-stream")
	// c.Header("Content-Disposition", "attachment; filename="+"用户信息.xlsx")
	// c.Header("Content-Transfer-Encoding", "binary")
	// f.Write(c.Writer)
}

// 导入文件
func (ah AdminHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []systemSchema.SystemAuthAdminResp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = systemService.AdminService.ImportFile(importList)
	response.CheckAndResp(c, err)
}

// list 管理员列表
func (ah AdminHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq systemSchema.SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := systemService.AdminService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// ListAll 所有管理员列表
func (ah AdminHandler) ListAll(c *gin.Context) {

	var listReq systemSchema.SystemAuthAdminListReq

	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := systemService.AdminService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 管理员详细
func (ah AdminHandler) Detail(c *gin.Context) {
	var detailReq systemSchema.SystemAuthAdminDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.AdminService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 管理员新增
func (ah AdminHandler) Add(c *gin.Context) {
	var addReq systemSchema.SystemAuthAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, systemService.AdminService.Add(addReq))
}

// edit 管理员编辑
func (ah AdminHandler) Edit(c *gin.Context) {
	var editReq systemSchema.SystemAuthAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, systemService.AdminService.Edit(c, editReq))
}

// upInfo 管理员更新
func (ah AdminHandler) UpInfo(c *gin.Context) {
	var updateReq systemSchema.SystemAuthAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	response.CheckAndResp(c, systemService.AdminService.Update(
		c, updateReq, config.AdminConfig.GetAdminId(c)))
}

// del 管理员删除
func (ah AdminHandler) Del(c *gin.Context) {
	var delReq systemSchema.SystemAuthAdminDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, systemService.AdminService.Del(c, delReq.ID))
}

// disable 管理员状态切换
func (ah AdminHandler) Disable(c *gin.Context) {
	var disableReq systemSchema.SystemAuthAdminDisableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	response.CheckAndResp(c, systemService.AdminService.Disable(c, disableReq.ID))
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
		response.FailWithMsg(c, response.Failed, "deptId不能为空")
		return
	}

	res, err := systemService.AdminService.ListByUserIdOrDeptIdPostId("", deptId, "")
	response.CheckAndRespWithData(c, res, err)
}
