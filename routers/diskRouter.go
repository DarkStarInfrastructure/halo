package routers

import (
	"halo/controllers/disks"

	"github.com/gin-gonic/gin"
)

func DiskRouters(r *gin.Engine) {
	diskRoutersInit := r.Group("/disks")
	diskRoutersInit.GET("/getalldisks", disks.DiskController{}.ListAllDisks)

}
