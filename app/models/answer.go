package models

import (
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type Answer struct {
	gorm.Model
	Text  string `gorm:"type:text" json:"text"`
	FaqId int    `gorm:"type:int" json:"faq_id"`
}

type AnswerI18ns struct {
	gorm.Model
	ReferenceId uint64 `gorm:"type:int(20);" json:"reference_id"`
	FiledName   string `gorm:"type:varchar(20);" json:"filed_name"`
	Language    string `gorm:"type:varchar(5);" json:"language"`
	Value       string `gorm:"type:text;" json:"value"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func AnswerMigrate() {
	_const.Services.DB.AutoMigrate(&Answer{})
	_const.Services.DB.AutoMigrate(&AnswerI18ns{})
}

/*
* return with module name
 */
func AnswerModule() string {
	return helpers.ModuleName("answers")
}
func AnswerTable() string {
	return helpers.ModuleTable(AnswerModule())
}

/*
* return with route name
 */
func AnswerRoute() string {
	return helpers.ModuleRoute(AnswerModule())
}

/*
* return with model name
 */
func AnswerModel() string {
	return helpers.ModelModel(AnswerModule())
}

/*
* return with Statuses
 */
func AnswerStatuses() interface{} {
	return helpers.ModuleStatuses(AnswerModule())
}

/**
* you can update these column only
 */
func AnswerFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["faq_id"] = "int"
	return m

}
func AnswerTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["text"] = "string"
	return m
}

func AnswerTransTable() string {
	return helpers.ModuleTransTable(AnswerModule())
}
