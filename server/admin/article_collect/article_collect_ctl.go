package article_collect

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type ArticleCollectHandler struct {
	Service IArticleCollectService
}

// list article_collect列表
func (hd ArticleCollectHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq ArticleCollectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := hd.Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail article_collect详情
func (hd ArticleCollectHandler) Detail(c *gin.Context) {
	var detailReq ArticleCollectDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := hd.Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

// add article_collect新增
func (hd ArticleCollectHandler) Add(c *gin.Context) {
	var addReq ArticleCollectAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Add(addReq))
}

// edit article_collect编辑
func (hd ArticleCollectHandler) Edit(c *gin.Context) {
	var editReq ArticleCollectEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Edit(editReq))
}

// del article_collect删除
func (hd ArticleCollectHandler) Del(c *gin.Context) {
	var delReq ArticleCollectDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, hd.Service.Del(delReq.Id))
}
