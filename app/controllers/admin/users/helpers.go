package users

import (
	"fmt"
	"net/url"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/user"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.User {
	return models.User{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.User {
	request := models.ConvertBodyToHashMap()

	user1 := models.User{
		Name:              request["name"].(string),
		Email:             request["email"].(string),
		PermissionGroupId: int(request["permission_group_id"].(float64)),
		Password:          request["password"].(string),
		Status:            request["status"].(string),
		Image:             request["image"].(string),
	}

	return user1
}

/**
* constructor Array
 */
func Rows() []models.User {
	return []models.User{}
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
	if g.Query("email") != "" {
		filter = append(filter, `email like "%`+g.Query("email")+`%"`)
	}
	if g.Query("role") != "" {
		filter = append(filter, `role like "%`+g.Query("role")+`%"`)
	}
	if g.Query("status") != "" {
		if g.Query("status") == "all" {
			filter = append(filter, "status <> "+_const.TRASH)
		}else {
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
func validateRequest(action string) bool {
	err := url.Values{}
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	if action == "store" {
		err = user.Store()
	} else {
		err = user.Update()
	}
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	fmt.Println("1111111111111111111111")
	return true
}

/**
* update row make sure you used UpdateOnlyAllowColumns to update allow columns
* use fill able method to only update what you need
 */
func updateColumns(data *models.User, oldRow *models.User) {
	////check if password not empty
	if data.Password != "" {
		password, _ := helpers.HashPassword(data.Password)
		data.Password = password
	}
	// update based on fill able data and assign the new data
	// the new data will set in the same pointer
	//models.Update(data, oldRow, models.UserFillAbleColumn())
}
