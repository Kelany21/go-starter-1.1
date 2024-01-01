package pages

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* get faq with answers
 */
func Faqs(g *gin.Context) {
	///// declare variables
	var rows []models.Faq
	_const.Services.DB.Scopes(models.ActiveFaq).Preload("Answers").Find(&rows)
	/// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), transformers.FaqsResponse(rows, true))

}
