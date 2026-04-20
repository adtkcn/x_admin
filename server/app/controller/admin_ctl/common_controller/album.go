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

// AlbumList 相册文件列表
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

// AlbumRename 相册文件重命名
func (ah AlbumHandler) AlbumRename(c *gin.Context) {
	var rnReq common_schema.CommonAlbumRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumRename(rnReq.ID, rnReq.Name))
}

// AlbumMove 相册文件移动
func (ah AlbumHandler) AlbumMove(c *gin.Context) {
	var mvReq common_schema.CommonAlbumMoveReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &mvReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumMove(mvReq.Ids, mvReq.Cid))
}

// AlbumDel 相册文件删除
func (ah AlbumHandler) AlbumDel(c *gin.Context) {
	var delReq common_schema.CommonAlbumDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.AlbumDel(delReq.Ids))
}

// CateList 类目列表
func (ah AlbumHandler) CateList(c *gin.Context) {
	var listReq common_schema.CommonCateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	res, err := common_service.AlbumService.CateList(adminId, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// CateAdd 类目新增
func (ah AlbumHandler) CateAdd(c *gin.Context) {
	var addReq common_schema.CommonCateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateAdd(adminId, addReq))
}

// CateRename 类目命名
func (ah AlbumHandler) CateRename(c *gin.Context) {
	var rnReq common_schema.CommonCateRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateRename(rnReq.ID, rnReq.Name))
}

// CateDel 类目删除
func (ah AlbumHandler) CateDel(c *gin.Context) {
	var delReq common_schema.CommonCateDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, common_service.AlbumService.CateDel(delReq.ID))
}
