package faqs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* get all rows with pagination
 */
func GetAllFaqs(g *gin.Context) {
	/// array of rows
	rows := Rows()
	/// query before any thing
	paginator := helpers.Paging(&helpers.Param{
		DB:      _const.Services.DB,
		Page:    helpers.Page(),
		Limit:   helpers.Limit(),
		OrderBy: helpers.Order("id desc"),
		Filters: filter(),
		Preload: preload(),
		ShowSQL: true,
	}, &rows)
	/// transform slice
	response := make(map[string]interface{})
	response["data"] = transformers.FaqsResponse(rows, true)
	// transform slice
	paginator.Records = response
	/// return response
	helpers.OkResponseWithPaging(helpers.DoneGetAllItems(), paginator)
}

/***
get all status
*/
func Status(g *gin.Context) {
	helpers.OkResponse(helpers.DoneGetAllItems(), transformers.ActionsResponse(models.GetActionByModule(models.FaqModule())))
}
func ShowFaq(g *gin.Context) {
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Answers")
	}); row.ID == 0 {
		return
	}
	/// transform row
	m := transformers.FaqResponse(row, true)
	/// transform translation
	m = transformers.TransformTranslationObject(helpers.RowsTranslations(models.FaqTransTable(), row.ID), m)
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), m)
}
