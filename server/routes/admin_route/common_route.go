package admin_route

import (
	"x_admin/app/controller/admin_ctl/common_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// initUploadRoute 上传路由
func initUploadRoute(rg *gin.RouterGroup) {
	handleUpload := common_controller.UploadHandler{}
	uploadRg := rg.Group("/common", middleware.LoginAuth())
	uploadRg.POST("/upload/preUploadFile", middleware.RecordLog("文件预上传", middleware.RequestFile), handleUpload.PreUploadFile)
	uploadRg.POST("/upload/file", middleware.RecordLog("上传文件", middleware.RequestFile), handleUpload.UploadFile)
}

// initChunkRoute 分片上传路由
func initChunkRoute(rg *gin.RouterGroup) {
	handleChunk := common_controller.UploadChunkHandler{
		UploadPath: "./uploads",
		TmpPath:    "./uploads/.tmp",
	}
	chunkRg := rg.Group("/common")
	chunkRg.GET("/upload_chunk/CheckFileExist", handleChunk.CheckFileExist)
	chunkRg.GET("/upload_chunk/HasChunk", handleChunk.HasChunk)
	chunkRg.POST("/upload_chunk/UploadChunk", handleChunk.UploadChunk)
	chunkRg.POST("/upload_chunk/MergeChunk", handleChunk.MergeChunk)
}

// initAlbumRoute 相册路由
func initAlbumRoute(rg *gin.RouterGroup) {
	handleAlbum := common_controller.AlbumHandler{}
	albumRg := rg.Group("/common", middleware.LoginAuth())
	albumRg.GET("/album/albumList", handleAlbum.AlbumList)
	albumRg.POST("/album/albumRename", middleware.RecordLog("相册文件重命名"), handleAlbum.AlbumRename)
	albumRg.POST("/album/albumMove", middleware.RecordLog("相册文件移动"), handleAlbum.AlbumMove)
	albumRg.POST("/album/albumDel", middleware.RecordLog("相册文件删除"), handleAlbum.AlbumDel)
	albumRg.GET("/album/cateList", handleAlbum.CateList)
	albumRg.POST("/album/cateAdd", middleware.RecordLog("相册分类新增"), handleAlbum.CateAdd)
	albumRg.POST("/album/cateRename", middleware.RecordLog("相册分类重命名"), handleAlbum.CateRename)
	albumRg.POST("/album/cateDel", middleware.RecordLog("相册分类删除"), handleAlbum.CateDel)
}

// initIndexRoute 首页路由
func initIndexRoute(rg *gin.RouterGroup) {
	handleIndex := common_controller.IndexHandler{}
	indexRg := rg.Group("/common")
	indexRg.GET("/index/console", middleware.LoginAuth(), handleIndex.Console)
	indexRg.GET("/index/config", handleIndex.Config)
}

// initGeTuiRoute 个推路由
func initGeTuiRoute(rg *gin.RouterGroup) {
	handleGeTui := common_controller.GeTuiHandler{}
	geTuiRg := rg.Group("/common")
	geTuiRg.GET("/push", handleGeTui.Push)
}

// initCaptchaRoute 验证码路由
func initCaptchaRoute(rg *gin.RouterGroup) {
	handleCaptcha := common_controller.CaptchaHandler{}
	captchaRg := rg.Group("/common/captcha")
	captchaRg.POST("/get", handleCaptcha.Get)
	captchaRg.POST("/check", handleCaptcha.Check)
}

// 通用模块路由入口（上传、分片上传、相册、首页、个推、验证码）
func init() {
	routeHandlers = append(routeHandlers, initUploadRoute, initChunkRoute, initAlbumRoute, initIndexRoute, initGeTuiRoute, initCaptchaRoute)
}
