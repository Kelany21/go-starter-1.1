package pages

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* get all rows with pagination
 */
func GetAllPages(g *gin.Context) {
	// array of rows
	var rows []models.Page
	// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      _const.Services.DB,
		Page:    helpers.Page(),
		Limit:   helpers.Limit(),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	// transform slice
	response := make(map[string]interface{})
	response["data"] = transformers.PagesResponse(rows,true)
	// transform slice
	paginator.Records = response
	// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
}

/***
get all status
*/
func Status(g *gin.Context) {
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.ActionsResponse(models.GetActionByModule(models.PageModule())))
}
