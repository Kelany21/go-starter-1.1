package categories

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/category"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.Category {
	return models.Category{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.Category {
	request := models.ConvertBodyToHashMap()

	category1 := models.Category{
		Status: request["status"].(string),
	}
	return category1
}

/**
* constructor Array
 */
func Rows() []models.Category {
	return []models.Category{}
}

/**
* filter module with some columns
 */
func filter() []string {
	g := _const.Services.GIN
	var filter []string
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name")+`%"`)
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
	return []string{}
} /**
* preload module with some preload conditions
 */
func preloadConditions() map[string][]string {
	conditions := make(map[string][]string)
	conditions["Translations"] = []string{"`category_i18ns`.language = ?", _const.Services.GIN.GetHeader("Accept-Language")}
	return conditions
}

/**
* here we will check if request valid or not
 */
func validateRequest() bool {
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := category.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.TestReturnNotValidRequest(err) {
		return false
	}
	return true
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(data *models.Category, oldRow *models.Category) {
	//models.Update(data, oldRow, models.CategoryFillAbleColumn())

}
