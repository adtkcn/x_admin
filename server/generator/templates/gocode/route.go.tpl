package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/core" 
	"x_admin/middleware" 
	"x_admin/admin/{{{ .ModuleName }}}"
)

// 请在 admin/entry.go 目录引入这个函数
// {{{ title (toCamelCase .ModuleName) }}}Route(rg)
func {{{ title (toCamelCase .ModuleName) }}}Route(rg *gin.RouterGroup) {
	db := core.GetDB()

	server := {{{ .ModuleName }}}.New{{{ title (toCamelCase .EntityName) }}}Service(db)

	handle := {{{ .ModuleName}}}.{{{ title (toCamelCase .EntityName) }}}Handler{Service: server}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/{{{ .ModuleName }}}/list", handle.List)
	rg.GET("/{{{ .ModuleName }}}/detail", handle.Detail)
	rg.POST("/{{{ .ModuleName }}}/add", handle.Add)
	rg.POST("/{{{ .ModuleName }}}/edit", handle.Edit)
	rg.POST("/{{{ .ModuleName }}}/del", handle.Del)
}