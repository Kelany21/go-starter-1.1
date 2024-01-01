package settings

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	name := models.SettingRoute()
	/// normal route
	r.GET(name+_const.LIST_URL, Index)
	r.GET(name+_const.Status, Status)
	r.PUT(name, Update)
	r.GET(name+_const.SHOW_URL, Show)
	/// this module can not have create or delete option its critical values
	/// must use to init the system

	return r
}
