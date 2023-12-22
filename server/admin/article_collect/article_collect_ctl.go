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

//	@Summary	article_collect列表
//	@Tags		article_collect文章
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		PageNo		query		int					true	"页码"
//	@Param		PageSize	query		int					true	"每页大小"
//	@Param		userId		query		int					false	"用户ID"
//	@Param		articleId	query		int					false	"文章ID"
//	@Success	200			{object}	response.PageResp	"成功"
//	@Failure	400			{object}	string				"请求错误"
//	@Failure	500			{object}	string				"内部错误"
//	@Router		/api/article_collect/list [get]
//
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
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}
