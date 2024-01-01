package general

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* this will be the first api
* call in the hole system
* return all translations
* all setting
 */
//TODO::will cache this response in redis
func Init(g *gin.Context) {
	/// declare variables
	var settings []models.Setting
	var pages []models.Page
	/// queries
	_const.Services.DB.Preload("Images").Find(&pages)
	_const.Services.DB.Find(&settings)
	/// build response
	var response = make(map[string]interface{})
	response["pages"] = transformers.PagesResponse(pages, true)
	response["settings"] = transformers.SettingsResponse(settings)
	/// return with data
	helpers.OkResponse(helpers.T("init_project"), response)
	return
}
