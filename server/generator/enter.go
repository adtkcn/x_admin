package generator

import (
	"x_admin/generator/gen"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {
	gen.GenRoute(rg)
}
