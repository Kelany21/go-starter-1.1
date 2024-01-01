package test

import (
	"starter-golang-new/app/models"
	"starter-golang-new/config"
	_const "starter-golang-new/const"
)

/**
* register new user
* return token as header
*/
func getTokenAsHeader(migrate bool) map[string]string {
	token := addAdminUser(migrate)
	var authToken = make(map[string]string)
	authToken["Authorization"] = token

	return authToken
}

/***
* create new admin
* return with admin token
*/
func addAdminUser(migrate bool) string {
	// connect database
	config.ConnectToDatabase()
	/// drop data base
	if migrate {
		models.MigrateAllTable()
	}
	data := models.User{
		Name:              "Abdel Aziz",
		Email:             "zizo199988@gmail.com",
		PermissionGroupId: 1,
		Password:          "1234567",
		Token:             "",
		Status:            "activate",
	}
	_const.Services.DB.Create(&data)

	return data.Token
}

