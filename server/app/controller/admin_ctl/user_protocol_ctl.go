package admin_ctl

import (
	"fmt"
	"strings"
	"time"
	"x_admin/app/schema"
	"x_admin/app/service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type UserProtocolHandler struct {
	// 防止缓存击穿场景，将大量重复数据库查询合并为一次。
	requestGroup singleflight.Group
}

// @Summary	用户协议列表
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token				header		string																		true	"token"
// @Param		pageNo				query		int																			true	"页码"
// @Param		pageSize			query		int																			true	"每页数量"
// @Param		title				query		string																		false	"标题"
// @Param		content				query		string																		false	"协议内容"
// @Param		version				query		number																		false	"版本"
// @Param		create_time_start	query		string																		false	"创建时间"
// @Param		create_time_end		query		string																		false	"创建时间"
// @Param		update_time_start	query		string																		false	"更新时间"
// @Param		update_time_end		query		string																		false	"更新时间"
//
// @Success	200					{object}	response.Response{data=response.PageResp{lists=[]schema.UserProtocolResp}}	"成功"
// @Router		/api/admin/user_protocol/list [get]
func (hd *UserProtocolHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq schema.UserProtocolListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}

	res, err := service.UserProtocolService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary	用户协议列表-所有
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token				header		string												true	"token"
// @Param		title				query		string												false	"标题"
// @Param		content				query		string												false	"协议内容"
// @Param		version				query		number												false	"版本"
// @Param		create_time_start	query		string												false	"创建时间"
// @Param		create_time_end		query		string												false	"创建时间"
// @Param		update_time_start	query		string												false	"更新时间"
// @Param		update_time_end		query		string												false	"更新时间"
// @Success	200					{object}	response.Response{data=[]schema.UserProtocolResp}	"成功"
// @Router		/api/admin/user_protocol/list_all [get]
func (hd *UserProtocolHandler) ListAll(c *gin.Context) {
	var listReq schema.UserProtocolListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := service.UserProtocolService.ListAll(listReq)
	response.JSON(c, res, err)
}

// @Summary	用户协议详情
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token	header		string											true	"token"
// @Param		id		query		string											false	"id"
// @Success	200		{object}	response.Response{data=schema.UserProtocolResp}	"成功"
// @Router		/api/admin/user_protocol/detail [get]
func (hd *UserProtocolHandler) Detail(c *gin.Context) {
	var detailReq schema.UserProtocolPrimarykey
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do(fmt.Sprintf("UserProtocol:Detail:%v", detailReq.Id), func() (any, error) {
		v, err := service.UserProtocolService.Detail(detailReq.Id)
		return v, err
	})

	response.JSON(c, res, err)
}

// @Summary	用户协议新增
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		tag		body		string				false	"标识"
// @Param		version	body		number				false	"版本"
// @Param		title	body		string				false	"标题"
// @Param		content	body		string				false	"协议内容"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/add [post]
func (hd *UserProtocolHandler) Add(c *gin.Context) {
	var addReq schema.UserProtocolAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	// 添加创建人
	var adminId = config.AdminConfig.GetAdminId(c)

	createId, err := service.UserProtocolService.Add(addReq, adminId)
	response.JSON(c, createId, err)
}

// @Summary	用户协议编辑
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		string				false	"id"
// @Param		tag		body		string				false	"标识"
// @Param		version	body		number				false	"版本"
// @Param		title	body		string				false	"标题"
// @Param		content	body		string				false	"协议内容"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/edit [post]
func (hd *UserProtocolHandler) Edit(c *gin.Context) {
	var editReq schema.UserProtocolEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := service.UserProtocolService.Edit(editReq)
	response.JSON(c, editReq.Id, err)
}

// @Summary	用户协议删除
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		string				false	"id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/del [post]
func (hd *UserProtocolHandler) Del(c *gin.Context) {
	var delReq schema.UserProtocolPrimarykey
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := service.UserProtocolService.Del(delReq.Id)
	response.JSON(c, nil, err)
}

// @Summary	用户协议删除-批量
// @Tags		user_protocol-用户协议
//
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		ids		body		string				false	"逗号分割的id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/del_batch [post]
func (hd *UserProtocolHandler) DelBatch(c *gin.Context) {
	var delReq schema.UserProtocolDelBatchReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailMsg(c, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	err := service.UserProtocolService.DelBatch(Ids)
	response.JSON(c, nil, err)
}

// @Summary	用户协议导出
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token				header		string				true	"token"
// @Param		title				query		string				false	"标题"
// @Param		content				query		string				false	"协议内容"
// @Param		version				query		number				false	"版本"
// @Param		create_time_start	query		string				false	"创建时间"
// @Param		create_time_end		query		string				false	"创建时间"
// @Param		update_time_start	query		string				false	"更新时间"
// @Param		update_time_end		query		string				false	"更新时间"
// @Success	200					{file}		string				"成功"
// @Failure	500					{object}	response.Response	"失败"
// @Router		/api/admin/user_protocol/export_file [get]
func (hd *UserProtocolHandler) ExportFile(c *gin.Context) {
	var listReq schema.UserProtocolListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := service.UserProtocolService.ExportFile(listReq)
	if err != nil {
		response.Fail(c, response.CheckErr(err, "查询信息失败"))
		return
	}
	f, err := excel2.Export(res, service.UserProtocolService.GetExcelCol(), "Sheet1", "用户协议")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "导出失败"))
		return
	}
	excel2.DownLoadExcel("用户协议"+time.Now().Format("20060102-150405"), c.Writer, f)
}

// @Summary	用户协议导入
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		file	formData	file				true	"导入文件"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/import_file [post]
func (hd *UserProtocolHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件不存在"))
		return
	}
	defer file.Close()
	importList := []schema.UserProtocolResp{}
	err = excel2.GetExcelData(file, &importList, service.UserProtocolService.GetExcelCol())
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件解析失败"))
		return
	}

	err = service.UserProtocolService.ImportFile(importList)
	response.JSON(c, nil, err)
}
