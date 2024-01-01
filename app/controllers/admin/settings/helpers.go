package settings

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/setting"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.Setting {
	return models.Setting{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.Setting {
	request := models.ConvertBodyToHashMap()

	setting1 := models.Setting{
		Name:        request["name"].(string),
		Value:       request["value"].(string),
		Status:      request["status"].(string),
		SettingType: request["setting_type"].(string),
		Slug:        request["slug"].(string),
	}
	return setting1
}

/**
* constructor Array
 */
func Rows() []models.Setting {
	return []models.Setting{}
}

/**
* filter module with some columns
 */
func filter() []string {
	g := _const.Services.GIN
	var filter []string
	if g.Query("value") != "" {
		filter = append(filter, `value like "%`+g.Query("value")+`%"`)
	}
	if g.Query("name") != "" {
		filter = append(filter, `name like "%`+g.Query("name")+`%"`)
	}
	if g.Query("slug") != "" {
		filter = append(filter, `slug like "%`+g.Query("slug")+`%"`)
	}
	if g.Query("setting_type") != "" {
		filter = append(filter, `setting_type like "%`+g.Query("setting_type")+`%"`)
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
}

/**
* here we will check if request valid or not
 */
func validateRequest() bool {
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := setting.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.TestReturnNotValidRequest(err) {
		return false
	}
	return true
}
