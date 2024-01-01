package models

import (
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/*
setting_type values
  | "text"
  | "textarea"
  | "email"
  | "date"
  | "time"
  | "select"
  | "multiselect"
  | "url"
  | "video"
  | "password"
  | "switch"
  | "checkbox"
  | "checkbox_group"
  | "radio_group"
  | "map"
  | "file"
  | "image"
  | "color";
*/

/***
* model struct here we will build the main
* struct that connect to database
 */
type Setting struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);" json:"name"`
	Value       string `gorm:"type:varchar(255);" json:"value"`
	Status      string `gorm:"type:varchar(20);" json:"status"`
	SettingType string `gorm:"type:varchar(20);" json:"setting_type"`
	Slug        string `gorm:"type:varchar(50);" json:"slug"`
}

/***
* when update setting value
 */
type UpdateSetting struct {
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Youtube  string `json:"youtube"`
	Linkedin string `json:"linkedin"`
}

/*
* return with module name
 */
func SettingModule() string {
	return helpers.ModuleName("settings")
}

func SettingTable() string {
	return helpers.ModuleTable(SettingModule())
}

/*
* return with model name
 */
func SettingModel() string {
	return helpers.ModelModel(SettingModule())
}

/*
* return with route name
 */
func SettingRoute() string {
	return helpers.ModuleRoute(SettingModule())
}

/*
* return with Statuses
 */
func SettingStatuses() interface{} {
	return helpers.ModuleStatuses(SettingModule())
}

/*
* event run after add Setting
 */
func (u *Setting) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(SettingModule())
	return
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func SettingMigrate() {
	_const.Services.DB.AutoMigrate(&Setting{})
}

/**
* you can update these column only
 */
func SettingFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"
	m["value"] = "string"
	m["setting_type"] = "string"
	m["slug"] = "string"

	return m
}
