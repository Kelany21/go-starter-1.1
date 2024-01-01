package permission_groups

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* get all permission group only
 */
func GetAllPermissionGroups(g *gin.Context)  {
	rows := Rows()
	_const.Services.DB.Find(&rows)
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.PermissionGroupsResponse(rows,true))
}

/***
get all status
*/
func Status(g *gin.Context) {
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.ActionsResponse(models.GetActionByModule(models.PermissionGroupModule())))
}