package permission_groups

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/permission_group"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.PermissionGroup {
	return models.PermissionGroup{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.PermissionGroup {
	request := models.ConvertBodyToHashMap()

	permissionGroup1 := models.PermissionGroup{
		Status: request["status"].(string),
	}
	return permissionGroup1
}

/**
* constructor Array
 */
func Rows() []models.PermissionGroup {
	return []models.PermissionGroup{}
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
}

/**
* here we will check if request valid or not
 */
func validateRequest() bool {

	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := permission_group.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	return true
}
