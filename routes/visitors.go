package routes

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/controllers/admin/settings"
	"starter-golang-new/app/controllers/visitor/auth"
	"starter-golang-new/app/controllers/visitor/general"
	"starter-golang-new/app/controllers/visitor/pages"
)

/***
* any route here will add after /
* anyone will have access this routes
 */
func Visitor(r *gin.RouterGroup) *gin.RouterGroup {
	general.Routes(r)
	auth.Routes(r)
	pages.Routes(r)
	r.GET("make-backup", settings.MakeBackup)
	/// serve static files like images
	r.Static("/public" , "./public")

	return r
}
