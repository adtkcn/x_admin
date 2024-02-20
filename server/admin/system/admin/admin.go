package admin

import (
	"strconv"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util/excel"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// func AdminRoute(rg *gin.RouterGroup) {
// 	db := core.GetDB()

// 	permSrv := role.NewSystemAuthPermService(db)
// 	roleSrv := role.NewSystemAuthRoleService(db, permSrv)
// 	adminSrv := NewSystemAuthAdminService(db, permSrv, roleSrv)
// 	// service := NewSystemLoginService(db, adminSrv)

// 	handle := AdminHandler{Service: adminSrv}

// 	rg = rg.Group("/system", middleware.TokenAuth())

// 	rg.GET("/admin/self", handle.self)
// 	rg.GET("/admin/list", handle.List)
// 	rg.GET("/admin/detail", handle.Detail)
// 	rg.POST("/admin/add", middleware.RecordLog("管理员新增"), handle.Add)
// 	rg.POST("/admin/edit", middleware.RecordLog("管理员编辑"), handle.Edit)
// 	rg.POST("/admin/upInfo", middleware.RecordLog("管理员更新"), handle.upInfo)
// 	rg.POST("/admin/del", middleware.RecordLog("管理员删除"), handle.Del)
// 	rg.POST("/admin/disable", middleware.RecordLog("管理员状态切换"), handle.disable)
// }

type AdminHandler struct {
	// Service ISystemAuthAdminService
}

// self 管理员信息
func (ah AdminHandler) Self(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := Service.Self(adminId)
	response.CheckAndRespWithData(c, res, err)
}

func (ah AdminHandler) ExportFile(c *gin.Context) {
	var listReq SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询导出失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "用户信息", "", true, false, nil)
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

// list 管理员列表
func (ah AdminHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 管理员详细
func (ah AdminHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthAdminDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 管理员新增
func (ah AdminHandler) Add(c *gin.Context) {
	var addReq SystemAuthAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 管理员编辑
func (ah AdminHandler) Edit(c *gin.Context) {
	var editReq SystemAuthAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(c, editReq))
}

// upInfo 管理员更新
func (ah AdminHandler) UpInfo(c *gin.Context) {
	var updateReq SystemAuthAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	response.CheckAndResp(c, Service.Update(
		c, updateReq, config.AdminConfig.GetAdminId(c)))
}

// del 管理员删除
func (ah AdminHandler) Del(c *gin.Context) {
	var delReq SystemAuthAdminDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(c, delReq.ID))
}

// disable 管理员状态切换
func (ah AdminHandler) Disable(c *gin.Context) {
	var disableReq SystemAuthAdminDisableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	response.CheckAndResp(c, Service.Disable(c, disableReq.ID))
}

// @Summary		获取部门的用户
// @Description	获取部门的用户
// @Tags			管理员
// @Param			deptId	path		int					true	"部门id"
// @Success		200		{object}	response.Response	"{"code": 200, "data": []}"
// @Router			/system/admin/ListByDeptId/{deptId} [get]
func (ah AdminHandler) ListByDeptId(c *gin.Context) {
	deptIdStr, bool := c.GetQuery("deptId")
	if bool == false {
		response.FailWithMsg(c, response.Failed, "deptId不能为空")
		return
	}
	deptId, err := strconv.Atoi(deptIdStr)
	if err != nil {
		response.FailWithMsg(c, response.Failed, "deptId参数错误")
		return
	}
	// if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &deptId)) {
	// return
	// }
	res, err := Service.ListByUserIdOrDeptIdPostId(0, deptId, 0)
	response.CheckAndRespWithData(c, res, err)
}
