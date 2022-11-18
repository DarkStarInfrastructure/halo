package routers

import (
	"halo/services"

	"github.com/gin-gonic/gin"
)

func DiskRouter(r *gin.RouterGroup) {
	rg := r.Group("/disks")
	rg.GET("/getalldisks", services.ListAllDisks)
}
