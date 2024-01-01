package pages

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/page"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.Page {
	return models.Page{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.Page {
	request := models.ConvertBodyToHashMap()

	page1 := models.Page{
		Status: request["status"].(string),
		Slug:   request["slug"].(string),
	}
	return page1
}
func RowImage() models.PageImage {
	return models.PageImage{}
}

/**
* constructor Array
 */
func Rows() []models.Page {
	return []models.Page{}
}

/**
* filter module with some columns
 */
func filter() []string {
	g := _const.Services.GIN
	var filter []string
	if g.Query("slug") != "" {
		filter = append(filter, `slug like "%`+g.Query("slug")+`%"`)
	}
	if g.Query("status") != "" {
		if g.Query("status") != "all" {
			filter = append(filter, `status = "`+g.Query("status")+`"`)
		}
	}
	return filter
}

/**
* preload module with some preload conditions
 */
func preload() []string {
	return []string{"Images"}
}

/**
* preload function when findOrFail
 */

/**
* here we will check if request valid or not
 */
func validateRequest() bool {

	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := page.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	return true
}

func getImages(page *models.Page) {
	_const.Services.DB.Where("page_id = ?", page.ID).Find(&page.Images)
}
