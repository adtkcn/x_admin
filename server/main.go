package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/docs"
	"x_admin/middleware"
	"x_admin/routes"

	_ "x_admin/app/corn"
	// _ "x_admin/docs"

	"github.com/gin-gonic/gin"
)

// // go:embed public/static
// var staticFs embed.FS

// initRouter 初始化router
func initRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.AppConfig.GinMode)
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 设置上传文件的静态路径路由
	r.Static(config.FileConfig.PublicPrefix, config.FileConfig.UploadDirectory)

	// staticHttpFs := http.FS(staticFs)
	// r.GET("/api/static/*filepath", func(c *gin.Context) {
	// 	filepath := c.Param("filepath")
	// 	fmt.Println(filepath)

	// 	c.FileFromFS("public/static"+filepath, staticHttpFs)
	// })

	// 静态文件路由
	r.Static("/api/static", "./public/static")

	// 设置中间件
	r.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())

	// 演示模式
	if config.AppConfig.DisallowModify {
		r.Use(middleware.ShowMode())
	}
	// 特殊异常处理
	r.NoMethod(response.NoMethod)
	// r.NoRoute(response.NoRoute)
	// 注册路由
	apiGroup := r.Group("/api")

	routes.RegisterRoute(apiGroup, r)

	return r
}

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,  //从连接建立到读取完整请求头和 body的最大时间
		WriteTimeout:   100 * time.Second, // 从读取完请求到写完响应的最大时间
		MaxHeaderBytes: 8192,              // 8KB,请求头最大字节数
	}
}

//	@title			x_admin文档
//	@version		0.0.1
//	@description	x_admin是一个完整的后台管理系统
//	@termsOfService	http://x.adtk.cn

//	@contact.name	API Support
//	@contact.url	http://x.adtk.cn
//	@contact.email	11675084@qq.com

//	@license.name	MIT License
//	@license.url	https://gitee.com/xiangheng/x_admin/blob/main/LICENSE
//	@BasePath	/
//	@securityDefinitions.basic	BasicAuth
//
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer db.Close()
	}

	// 初始化router
	router := initRouter()
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/api/swagger/doc.json", func(c *gin.Context) {
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", config.AppConfig.Port)
		c.String(200, docs.SwaggerInfo.ReadDoc())
	})
	fmt.Println("格式化文档注释:", "swag fmt")
	fmt.Println("生成文档:", "swag init")
	// fmt.Printf("文档: http://localhost:%v/swagger/index.html", config.AppConfig.Port)
	fmt.Printf("文档: http://localhost:%v/api/static/api/index.html", config.AppConfig.Port)

	// 初始化server
	s := initServer(router)
	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())

}
