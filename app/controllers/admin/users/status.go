package users

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"strconv"
)

/***
* active un de active
 */
func Active(g *gin.Context) {
	row := Row()
	if checkIfUserTakeActionOnHimSelf(){
		return
	}
	helpers.Active(g.Param("id"), &row)
}

/***
* de active un active
 */
func DeActive(g *gin.Context) {
	row := Row()
	if checkIfUserTakeActionOnHimSelf(){
		return
	}
	helpers.DeActive(g.Param("id"), &row)
}

/***
* trash item
 */
func Trash(g *gin.Context) {
	row := Row()
	if checkIfUserTakeActionOnHimSelf(){
		return
	}
	helpers.Trash(g.Param("id"), &row)
}

/**
* restore trash
 */
func Block(g *gin.Context) {
	row := Row()
	if checkIfUserTakeActionOnHimSelf(){
		return
	}
	helpers.Block(g.Param("id"), &row)
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
* bulk trash item
 */
func BulkBlock(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.BLOCK, &row)
}

/***
* bulk delete item
 */
func BulkDelete(g *gin.Context) {
	row := Row()
	helpers.BulkInit(_const.DELETE, &row)
}

func checkIfUserTakeActionOnHimSelf() bool {
	authUser := models.Auth()
	id, _ := strconv.Atoi(_const.Services.GIN.Param("id"))
	if int(authUser.ID) == id {
		helpers.YouCantTrashOrDeleteYourSelf()
		return true
	}
	return false
}
