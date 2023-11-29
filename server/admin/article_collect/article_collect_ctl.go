package article_collect

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type ArticleCollectHandler struct {
	Service ArticleCollectService
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
