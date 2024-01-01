package routes

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/controllers/admin/categories"
	"starter-golang-new/app/controllers/admin/faqs"
	"starter-golang-new/app/controllers/admin/images"
	"starter-golang-new/app/controllers/admin/pages"
	"starter-golang-new/app/controllers/admin/permission_groups"
	"starter-golang-new/app/controllers/admin/settings"
	"starter-golang-new/app/controllers/admin/users"
)

/***
* any route here will add after /admin
* admin only  will have access this routes
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	categories.Routes(r)
	settings.Routes(r)
	users.Routes(r)
	pages.Routes(r)
	faqs.Routes(r)
	images.Routes(r)
	permission_groups.Routes(r)

	return r
}
