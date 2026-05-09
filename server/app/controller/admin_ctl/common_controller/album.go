package common_controller

import (
	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// AlbumHandler 相册控制器
type AlbumHandler struct{}

// @Summary		相册文件列表
// @Description	获取相册文件列表
// @Tags			common_album-相册管理
// @Param			token		header		string									true	"token"
// @Param			pageNo		query		int										true	"页码"
// @Param			pageSize	query		int										true	"每页数量"
// @Param			cid			query		string									false	"分类ID"
// @Param			name		query		string									false	"文件名"
// @Param			ext		query		[]string									false	"文件扩展名"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=common_schema.CommonAlbumListResp}}	"成功"
// @Router			/api/admin/common/album/list [get]
func (ah AlbumHandler) AlbumList(c *gin.Context) {
	var page request.PageReq
	var listReq common_schema.CommonAlbumListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	res, err := common_service.AlbumService.AlbumList(adminId, page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		相册文件重命名
// @Description	相册文件重命名
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			id		body		string						true	"文件ID"
// @Param			name	body		string						true	"新文件名"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/rename [post]
func (ah AlbumHandler) AlbumRename(c *gin.Context) {
	var rnReq common_schema.CommonAlbumRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumRename(rnReq.ID, rnReq.Name))
}

// @Summary		相册文件移动
// @Description	相册文件移动到指定分类
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			ids		body		[]string					true	"文件ID列表"
// @Param			cid		body		string						true	"目标分类ID"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/move [post]
func (ah AlbumHandler) AlbumMove(c *gin.Context) {
	var mvReq common_schema.CommonAlbumMoveReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &mvReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumMove(mvReq.Ids, mvReq.Cid))
}

// @Summary		相册文件删除
// @Description	删除相册文件
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			ids		body		[]string					true	"文件ID列表"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/del [post]
func (ah AlbumHandler) AlbumDel(c *gin.Context) {
	var delReq common_schema.CommonAlbumDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumDel(delReq.Ids))
}

// @Summary		类目列表
// @Description	获取相册类目列表
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			name		query		string									false	"类目名称"
// @Success		200		{object}	response.Response{data=[]common_schema.CommonCateListResp}	"成功"
// @Router			/api/admin/common/album/cateList [get]
func (ah AlbumHandler) CateList(c *gin.Context) {
	var listReq common_schema.CommonCateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	res, err := common_service.AlbumService.CateList(adminId, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		类目新增
// @Description	新增相册类目
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			pid	body		string						false	"父类目ID"
// @Param			name	body		string						true	"类目名称"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/cateAdd [post]
func (ah AlbumHandler) CateAdd(c *gin.Context) {
	var addReq common_schema.CommonCateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateAdd(adminId, addReq))
}

// @Summary		类目重命名
// @Description	相册类目重命名
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			id		body		string						true	"类目ID"
// @Param			name	body		string						true	"新类目名称"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/cateRename [post]
func (ah AlbumHandler) CateRename(c *gin.Context) {
	var rnReq common_schema.CommonCateRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateRename(rnReq.ID, rnReq.Name))
}

// @Summary		类目删除
// @Description	删除相册类目
// @Tags			common_album-相册管理
// @Param			token	header		string						true	"token"
// @Param			id		body		string						true	"类目ID"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/album/cateDel [post]
func (ah AlbumHandler) CateDel(c *gin.Context) {
	var delReq common_schema.CommonCateDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateDel(delReq.ID))
}
