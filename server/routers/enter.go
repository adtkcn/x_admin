package routers

import (
	"x_admin/admin"
	"x_admin/generator"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {
	admin.RegisterGroup(rg)
	generator.RegisterGroup(rg)
}
