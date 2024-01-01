package main

import (
	"github.com/bykovme/gotrans"
	"os"
	"path/filepath"
	"runtime"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests"
	"starter-golang-new/config"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"starter-golang-new/providers"
	"starter-golang-new/seeders"
)

func main() {
	/**
	* load env file
	 */
	helpers.LoadEnv()
	/**
	* set project base path
	 */
	_, b, _, _ := runtime.Caller(0)
	if os.Getenv("APP_ENV") == "local" {
		_const.DIR = filepath.Dir(b)
	} else {
		_const.DIR = os.Getenv("PUBLIC_PATH")
	}
	/**
	* start multi language
	 */
	err := gotrans.InitLocales("public/trans")
	if err != nil {
		panic(err)
	}
	/**
	* start container will carry all models
	* and database connection
	 */
	providers.StartContainer()
	/**
	* add custom role to validation
	 */
	requests.Init()
	/**
	* connect with data base logic you can edit .env file to
	* change any connection params
	 */
	config.ConnectToDatabase()
	/**
	* drop All tables and migrate
	* to stop delete tables make DROP_ALL_TABLES false in env file
	* if you need to stop auto migration just stop this line
	 */
	models.MigrateAllTable()
	/**
	* this function will open seeders folder look inside all files
	* search for seeders function and seed execute these function
	* if you need to stop seeding you can stop this line
	 */
	seeders.Seed()
	/**
	* Run gin framework
	* add middleware
	* run routing
	* serve app
	 */
	//fmt.Println(helpers.Row("Select * from users where id = 1"))
	providers.Run()
}
