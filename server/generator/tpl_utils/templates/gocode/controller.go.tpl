package {{{ .ModuleName }}}

import (
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
)

 
type {{{ title (toCamelCase .ModuleName) }}}Handler struct {
	Service I{{{ title (toCamelCase .EntityName) }}}Service
}

//list {{{ .ModuleName }}}列表
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) List(c *gin.Context) {
	var page request.PageReq
	var listReq {{{ title (toCamelCase .EntityName) }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//detail {{{ .ModuleName }}}详情
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Detail(c *gin.Context) {
	var detailReq {{{ title (toCamelCase .EntityName) }}}DetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := hd.Service.Detail(detailReq.{{{ title (toCamelCase .PrimaryKey) }}})
	response.CheckAndRespWithData(c, res, err)
}

//add {{{ .ModuleName }}}新增
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Add(c *gin.Context) {
	var addReq {{{ title (toCamelCase .EntityName) }}}AddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Add(addReq))
}

//edit {{{ .ModuleName }}}编辑
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Edit(c *gin.Context) {
	var editReq {{{ title (toCamelCase .EntityName) }}}EditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Edit(editReq))
}

//del {{{ .ModuleName }}}删除
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Del(c *gin.Context) {
	var delReq {{{ title (toCamelCase .EntityName) }}}DelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Del(delReq.{{{ title (toCamelCase .PrimaryKey) }}}))
}
