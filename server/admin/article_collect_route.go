package admin

import (
	"x_admin/admin/article_collect"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

// 请在 admin/entry.go 目录引入这个函数
// ArticleCollectRoute(rg)
func ArticleCollectRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// server := article_collect.NewArticleCollectService(db)

	handle := article_collect.ArticleCollectHandler{}

	rg = rg.Group("/", middleware.TokenAuth())

	rg.GET("/article_collect/list", handle.List)
}
