package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/notice_service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// NoticeHandler 通知控制器
type NoticeHandler struct{}

// @Summary		通知列表
// @Description	获取当前管理员的通知列表
// @Tags			system_notice-通知
// @Param			token		header		string																			true	"token"
// @Param			pageNo		query		int																				true	"页码"
// @Param			pageSize	query		int																				true	"每页数量"
// @Param			type		query		string																			false	"通知类型"
// @Param			is_read		query		int																				false	"0未读 1已读 -1全部"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemNoticeResp}}	"成功"
// @Router			/api/admin/system/notice/list [get]
func (h NoticeHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq system_schema.SystemNoticeListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	adminId := config.AdminConfig.GetAdminId(c)
	list, count, err := notice_service.NoticeService.List(adminId, page.PageNo, page.PageSize, &listReq)
	if err != nil {
		response.Fail(c, "获取通知列表失败")
		return
	}
	response.Ok(c, response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    list,
	})
}

// @Summary		未读数量
// @Description	获取当前管理员的未读通知数量
// @Tags			system_notice-通知
// @Param			token	header		string																true	"token"
// @Success		200		{object}	response.Response{data=system_schema.SystemNoticeUnreadCountResp}	"成功"
// @Router			/api/admin/system/notice/unread_count [get]
func (h NoticeHandler) UnreadCount(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	count, err := notice_service.NoticeService.UnreadCount(adminId)
	response.CheckAndRespWithData(c, system_schema.SystemNoticeUnreadCountResp{Count: count}, err)
}

// @Summary		标记已读
// @Description	标记单条通知为已读
// @Tags			system_notice-通知
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"通知ID"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/notice/read [post]
func (h NoticeHandler) Read(c *gin.Context) {
	var req system_schema.SystemNoticeReadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	adminId := config.AdminConfig.GetAdminId(c)
	err := notice_service.NoticeService.Read(req.ID, adminId)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		全部已读
// @Description	将当前管理员所有未读通知标记为已读
// @Tags			system_notice-通知
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/notice/read_all [post]
func (h NoticeHandler) ReadAll(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	err := notice_service.NoticeService.ReadAll(adminId)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		删除通知
// @Description	删除指定通知
// @Tags			system_notice-通知
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"通知ID"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/notice/del [post]
func (h NoticeHandler) Del(c *gin.Context) {
	var req system_schema.SystemNoticeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	adminId := config.AdminConfig.GetAdminId(c)
	err := notice_service.NoticeService.Del(req.ID, adminId)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		获取通知偏好
// @Description	获取当前管理员的渠道通知偏好
// @Tags			system_notice-通知
// @Param			token	header		string															true	"token"
// @Success		200		{object}	response.Response{data=system_schema.SystemNoticeSettingResp}	"成功"
// @Router			/api/admin/system/notice/setting [get]
func (h NoticeHandler) GetSetting(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := notice_service.NoticeService.GetSetting(adminId)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		保存通知偏好
// @Description	保存当前管理员的渠道通知偏好（channel -> is_enabled 映射）
// @Tags			system_notice-通知
// @Param			token				header		string															true	"token"
// @Param			settings				body		system_schema.SystemNoticeSettingSaveReq						true	"渠道开关映射"
// @Success		200					{object}	response.Response												"成功"
// @Router			/api/admin/system/notice/setting/save [post]
func (h NoticeHandler) SaveSetting(c *gin.Context) {
	var req system_schema.SystemNoticeSettingSaveReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	adminId := config.AdminConfig.GetAdminId(c)
	err := notice_service.NoticeService.SaveSetting(adminId, &req)
	response.CheckAndRespWithData(c, nil, err)
}
