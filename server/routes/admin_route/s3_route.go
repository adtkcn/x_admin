package admin_route

import (
	"x_admin/app/controller/admin_ctl/common_controller"

	"github.com/gin-gonic/gin"
)

// initS3Route S3 标准协议路由（与 AWS SDK 兼容）
// 所有操作通过单一 wildcard 路由 + S3Handler 内部分发
func initS3Route(rg *gin.RouterGroup) {
	h := common_controller.S3UploadHandler{}
	s3 := rg.Group("/s3")

	// 核心路由: 统一入口，按 HTTP 方法 + query 参数分发
	// - POST /s3/{key}?uploads            → CreateMultipartUpload
	// - PUT  /s3/{key}?partNumber=&uploadId= → UploadPart
	// - POST /s3/{key}?uploadId=          → CompleteMultipartUpload
	// - DELETE /s3/{key}?uploadId=        → AbortMultipartUpload
	// - GET  /s3/{key}?uploadId=          → ListParts
	// - PUT  /s3/{key}                    → PutObject
	// - HEAD /s3/{key}                    → HeadObject
	// - POST /s3/generateKey              → 生成文件 Key（扩展）
	s3.Any("/*fileKey", h.S3Handler)
}

func init() {
	routeHandlers = append(routeHandlers, initS3Route)
}
