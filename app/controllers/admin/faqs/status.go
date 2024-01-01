package faqs

import (
	"github.com/gin-gonic/gin"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* active un deactive
 */
func Active(g *gin.Context) {
	row := Row()
	helpers.Active(g.Param("id"), &row)
}

/***
* deactive un active
 */
func DeActive(g *gin.Context) {
	row := Row()
	helpers.DeActive(g.Param("id"), &row)
}

/***
* trash item
 */
func Trash(g *gin.Context) {
	row := Row()
	helpers.Trash(g.Param("id"), &row)
}

/***
* bulk active item
 */
func BulkActive(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.ACTIVE, &row)
}

/***
* bulk de-active item
 */
func BulkDeActive(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.DEACTIVATE, &row)
}

/***
* bulk trash item
 */
func BulkTrash(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.TRASH, &row)
}

/***
* bulk delete item
 */
func BulkDelete(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.DELETE, &row)
}
