package main

import (
	"database/sql/driver"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/routes"

	_ "x_admin/app/corn"
	// _ "x_admin/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

// ValidateValuer 将 NullInt等类型 转换为底层值（int64 或 nil）
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, _ := valuer.Value()
		return val // 返回 int64 或 nil
	}
	return nil
}

//	@description	x_admin是一个完整的后台管理系统
//	@termsOfService	http://x.adtk.cn

//	@contact.name	xh
//	@contact.url	http://x.adtk.cn
//	@contact.email	x@adtk.cn

// @license.name				MIT License
// @license.url				https://gitee.com/xiangheng/x_admin/blob/main/LICENSE
// @BasePath					/
//
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	// 注册自定义类型的验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(ValidateValuer, core.NullString{}, core.NullInt{}, core.NullFloat{}, core.NullTime{})
	}

	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer db.Close()
	}

	// 初始化router
	router := initRouter()

	fmt.Println("格式化文档注释:", "swag fmt")
	fmt.Println("生成文档:", "swag init")
	// fmt.Printf("文档: http://localhost:%v/swagger/index.html", config.AppConfig.Port)
	fmt.Printf("文档: http://localhost:%v/api/static/api/index.html\n", config.AppConfig.Port)

	// 初始化server
	s := initServer(router)
	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())

}
