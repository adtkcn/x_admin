package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"x_admin/app/model/user_model"
	"x_admin/config"
	"x_admin/core"
	"x_admin/plugin"
	"x_admin/routes"

	_ "x_admin/app/corn"

	"github.com/gin-gonic/gin"
)

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

//	@description	x_admin后台管理系统
//	@termsOfService	http://x.adtk.cn

//	@contact.name	xh
//	@contact.url	http://x.adtk.cn
//	@contact.email	x@adtk.cn

// @license.name	MIT License
// @license.url	https://gitee.com/xiangheng/x_admin/blob/main/LICENSE
// @BasePath		/
func main() {
	plugin.RegisterNullValidator()
	// 刷新日志缓冲
	defer core.Logger.Sync()
	// 程序结束前关闭数据库连接
	if core.GetDB() != nil {
		db, _ := core.GetDB().DB()
		defer db.Close()
	}

	// 自动迁移用户表
	core.AutoMigrate(&user_model.User{}, &user_model.UserAuth{})

	// 初始化router
	router := routes.InitRouter()
	// 初始化server
	s := initServer(router)

	fmt.Println("格式化文档注释:", "swag fmt")
	fmt.Println("生成文档:", "swag init")
	fmt.Printf("文档: http://localhost:%v/api/static/api/index.html\n", config.AppConfig.Port)

	// 运行服务
	log.Fatalln(s.ListenAndServe().Error())

}
