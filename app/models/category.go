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
type Category struct {
	gorm.Model
	Status        string          `gorm:"type:varchar(20);" json:"status"`
}

type CategoryI18ns struct {
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
func CategoryMigrate() {
	_const.Services.DB.AutoMigrate(&Category{})
	_const.Services.DB.AutoMigrate(&CategoryI18ns{})
}

/*
* return with module name
 */
func CategoryModule() string {
	return helpers.ModuleName("categories")
}

func CategoryTable() string {
	return helpers.ModuleTable(CategoryModule())
}

func CategoryTransTable() string {
	return helpers.ModuleTransTable(CategoryModule())
}

/*
* return with route name
 */
func CategoryRoute() string {
	return helpers.ModuleRoute(CategoryModule())
}

/*
* return with model name
 */
func CategoryModel() string {
	return helpers.ModelModel(CategoryModule())
}

/*
* return with Statuses
 */
func CategoryStatuses() interface{} {
	return helpers.ModuleStatuses(CategoryModule())
}

/*
* event run after add Category
 */
func (u *Category) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(CategoryModule())
	return
}

/*
* event run after delete Category
 */
func (u *Category) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, CategoryModule())
	return
}

/**
* update status
 */
func (u *Category) BeforeUpdate() (err error) {

	var category Category
	if u.ID != 0 {
		_const.Services.DB.First(&category, u.ID)
		if category.Status != u.Status {
			DecreaseRow(category.Status, CategoryModule())
			IncreaseRow(u.Status, CategoryModule())
		}
	}
	return
}

/**
* you can update these column only
 */
func CategoryFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["status"] = "string"

	return m
}

func CategoryTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"
	m["description"] = "text"

	return m
}

/**
* active category only
 */
func ActiveCategory(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + _const.ACTIVE)
}
