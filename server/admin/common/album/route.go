package album

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func AlbumRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewAlbumService(db)

	handle := albumHandler{}

	rg = rg.Group("/common", middleware.TokenAuth())

	rg.GET("/album/albumList", handle.albumList)
	rg.POST("/album/albumRename", middleware.RecordLog("相册文件重命名"), handle.albumRename)
	rg.POST("/album/albumMove", middleware.RecordLog("相册文件移动"), handle.albumMove)
	rg.POST("/album/albumDel", middleware.RecordLog("相册文件删除"), handle.albumDel)
	rg.GET("/album/cateList", handle.cateList)
	rg.POST("/album/cateAdd", middleware.RecordLog("相册分类新增"), handle.cateAdd)
	rg.POST("/album/cateRename", middleware.RecordLog("相册分类重命名"), handle.cateRename)
	rg.POST("/album/cateDel", middleware.RecordLog("相册分类删除"), handle.cateDel)
}

type albumHandler struct{}

// albumList 相册文件列表
func (ah albumHandler) albumList(c *gin.Context) {
	var page request.PageReq
	var listReq CommonAlbumListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.AlbumList(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// albumRename 相册文件重命名
func (ah albumHandler) albumRename(c *gin.Context) {
	var rnReq CommonAlbumRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndResp(c, Service.AlbumRename(rnReq.ID, rnReq.Name))
}

// albumMove 相册文件移动
func (ah albumHandler) albumMove(c *gin.Context) {
	var mvReq CommonAlbumMoveReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &mvReq)) {
		return
	}
	response.CheckAndResp(c, Service.AlbumMove(mvReq.Ids, mvReq.Cid))
}

// albumDel 相册文件删除
func (ah albumHandler) albumDel(c *gin.Context) {
	var delReq CommonAlbumDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.AlbumDel(delReq.Ids))
}

// cateList 类目列表
func (ah albumHandler) cateList(c *gin.Context) {
	var listReq CommonCateListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.CateList(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// cateAdd 类目新增
func (ah albumHandler) cateAdd(c *gin.Context) {
	var addReq CommonCateAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.CateAdd(addReq))
}

// cateRename 类目命名
func (ah albumHandler) cateRename(c *gin.Context) {
	var rnReq CommonCateRenameReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &rnReq)) {
		return
	}
	response.CheckAndResp(c, Service.CateRename(rnReq.ID, rnReq.Name))
}

// cateDel 类目删除
func (ah albumHandler) cateDel(c *gin.Context) {
	var delReq CommonCateDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.CateDel(delReq.ID))
}
