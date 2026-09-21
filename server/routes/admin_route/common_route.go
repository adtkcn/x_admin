package admin_route

import (
	"fmt"
	"x_admin/app/controller/admin_ctl/common_controller"
	"x_admin/app/middleware"
	"x_admin/config"
	"x_admin/docs"

	"github.com/gin-gonic/gin"
)

// initUploadRoute 上传路由
func initUploadRoute(rg *gin.RouterGroup) {
	handleUpload := common_controller.UploadHandler{}
	uploadRg := rg.Group("/common", middleware.LoginAuth())

	uploadRg.POST("/upload/file", handleUpload.UploadFile)
	uploadRg.POST("/upload/checkInstant", handleUpload.CheckInstant)
}

// initUploadChunkRoute S3 标准协议路由（与 AWS SDK 兼容）
func initUploadChunkRoute(rg *gin.RouterGroup) {
	h := common_controller.UploadChunkHandler{}
	upload_chunk := rg.Group("/upload_chunk")

	// 核心路由: 统一入口，按 HTTP 方法 + query 参数分发
	// - POST /upload_chunk/{key}?uploads            → CreateMultipartUpload
	// - PUT  /upload_chunk/{key}?partNumber=&uploadId= → UploadPart
	// - POST /upload_chunk/{key}?uploadId=          → CompleteMultipartUpload
	// - DELETE /upload_chunk/{key}?uploadId=        → AbortMultipartUpload
	// - GET  /upload_chunk/{key}?uploadId=          → ListParts
	// - PUT  /upload_chunk/{key}                    → PutObject
	// - HEAD /upload_chunk/{key}                    → HeadObject
	// - POST /upload_chunk/generateKey              → 生成文件 Key（扩展）
	upload_chunk.Any("/*fileKey", h.S3Handler)
}

// initAlbumRoute 相册路由
func initAlbumRoute(rg *gin.RouterGroup) {
	handleAlbum := common_controller.AlbumHandler{}
	albumRg := rg.Group("/common", middleware.LoginAuth())
	albumRg.GET("/album/albumList", handleAlbum.AlbumList)
	albumRg.POST("/album/albumAddFromFile", middleware.RecordLog("相册文件挂载"), handleAlbum.AlbumAddFromFile)
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

// @Summary	swagger文档数据
// @Tags		公共接口
// @Router		/api/admin/swagger/doc.json [get]
func swaggerDoc(rg *gin.RouterGroup) {
	rg.GET("/swagger/doc.json", func(c *gin.Context) {
		// 获取域名和端口号
		host := ""
		docs.SwaggerInfo.Host = fmt.Sprintf("%v", host)
		docs.SwaggerInfo.Title = config.AppConfig.AppName
		docs.SwaggerInfo.Version = config.AppConfig.Version
		c.String(200, docs.SwaggerInfo.ReadDoc())
	})
}

// 通用模块路由入口（上传、分片上传、相册、首页、个推、验证码）
func init() {
	routeHandlers = append(
		routeHandlers,
		initUploadRoute,
		initUploadChunkRoute,
		initAlbumRoute,
		initIndexRoute,
		initGeTuiRoute,
		swaggerDoc,
	)
}
