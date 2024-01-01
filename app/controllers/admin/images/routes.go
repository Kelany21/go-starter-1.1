package images

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	name := models.ImageRoute()

	//// images
	r.POST(name+"/upload", Store)
	r.POST(name+"/delete", Delete)
	return r
}
