package user_protocol

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type UserProtocolHandler struct {
	requestGroup singleflight.Group
}

//	 @Summary	用户协议列表
//	 @Tags		user_protocol-用户协议
//	 @Produce	json
//	 @Param		Token		header		string				true	"token"
//	 @Param		PageNo		query		int					true	"页码"
//	 @Param		PageSize	query		int					true	"每页数量"
//		@Param Title query string false "标题"
//		@Param Content query string false "协议内容"
//		@Param Sort query number false "排序"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//
// @Success 200	{object} response.Response{ data=response.PageResp{ lists=[]UserProtocolResp}}	"成功"
// @Router	/api/admin/user_protocol/list [get]
func (hd *UserProtocolHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq UserProtocolListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := UserProtocolService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	用户协议列表-所有
//		@Tags		user_protocol-用户协议
//	 @Produce	json
//		@Param Title query string false "标题"
//		@Param Content query string false "协议内容"
//		@Param Sort query number false "排序"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//		@Success	200			{object}	response.Response{ data=[]UserProtocolResp}	"成功"
//		@Router		/api/admin/user_protocol/listAll [get]
func (hd *UserProtocolHandler) ListAll(c *gin.Context) {
	var listReq UserProtocolListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := UserProtocolService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	用户协议详情
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		query		number				false	""
// @Success	200			{object}	response.Response{ data=UserProtocolResp}	"成功"
// @Router		/api/admin/user_protocol/detail [get]
func (hd *UserProtocolHandler) Detail(c *gin.Context) {
	var detailReq UserProtocolDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("UserProtocol:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := UserProtocolService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

// @Summary	用户协议新增
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Title		body		string				false	"标题"
// @Param		Content		body		string				false	"协议内容"
// @Param		Sort		body		number				false	"排序"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/add [post]
func (hd *UserProtocolHandler) Add(c *gin.Context) {
	var addReq UserProtocolAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := UserProtocolService.Add(addReq)
	response.CheckAndRespWithData(c, createId, e)
}

// @Summary	用户协议编辑
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	""
// @Param		Title		body		string				false	"标题"
// @Param		Content		body		string				false	"协议内容"
// @Param		Sort		body		number				false	"排序"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/edit [post]
func (hd *UserProtocolHandler) Edit(c *gin.Context) {
	var editReq UserProtocolEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, UserProtocolService.Edit(editReq))
}

// @Summary	用户协议删除
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	""
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/del [post]
func (hd *UserProtocolHandler) Del(c *gin.Context) {
	var delReq UserProtocolDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, UserProtocolService.Del(delReq.Id))
}

//	@Summary	用户协议删除-批量
//	@Tags		user_protocol-用户协议
//
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Ids		body		string				false	"逗号分割的id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/user_protocol/delBatch [post]
func (hd *UserProtocolHandler) DelBatch(c *gin.Context) {
	var delReq UserProtocolDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, UserProtocolService.DelBatch(Ids))
}

// @Summary	用户协议导出
// @Tags		user_protocol-用户协议
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param Title query string false "标题"
// @Param Content query string false "协议内容"
// @Param Sort query number false "排序"
// @Param CreateTimeStart  query string false "创建时间"
// @Param CreateTimeEnd  query string false "创建时间"
// @Param UpdateTimeStart  query string false "更新时间"
// @Param UpdateTimeEnd  query string false "更新时间"
// @Router		/api/admin/user_protocol/ExportFile [get]
func (hd *UserProtocolHandler) ExportFile(c *gin.Context) {
	var listReq UserProtocolListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := UserProtocolService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, UserProtocolService.GetExcelCol(), "Sheet1", "用户协议")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("用户协议"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	 @Summary	用户协议导入
//	 @Tags		user_protocol-用户协议
//	 @Produce	json
//		@Router		/api/admin/user_protocol/ImportFile [post]
func (hd *UserProtocolHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []UserProtocolResp{}
	err = excel2.GetExcelData(file, &importList, UserProtocolService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = UserProtocolService.ImportFile(importList)
	response.CheckAndResp(c, err)
}
