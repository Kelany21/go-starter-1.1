package models

import (
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/***
* model struct here we will build the main
* struct that connect to database
* status 1 active 2 is not active
 */
type Page struct {
	gorm.Model
	Status string      `gorm:"type:varchar(20);" json:"status"`
	Slug   string      `gorm:"type:varchar(20);" json:"slug"`
	Image  []string    `gorm:"-" json:"image"`
	Images []PageImage `gorm:"association_autoupdate:false;association_autocreate:false" json:"images"`
}

type PageI18ns struct {
	gorm.Model
	ReferenceId uint64 `gorm:"type:int(20);" json:"reference_id"`
	FiledName   string `gorm:"type:varchar(20);" json:"filed_name"`
	Language    string `gorm:"type:varchar(5);" json:"language"`
	Value       string `gorm:"type:text;" json:"value"`
}

/*
* return with module name
 */
func PageModule() string {
	return helpers.ModuleName("pages")
}

func PageTable() string {
	return helpers.ModuleTable(PageModule())
}

/*
* return with route name
 */
func PageRoute() string {
	return helpers.ModuleRoute(PageModule())
}

func PageTransTable() string {
	return helpers.ModuleTransTable(PageModule())
}

/*
* return with model name
 */
func PageModel() string {
	return helpers.ModelModel(PageModule())
}

/*
* return with Statuses
 */
func PageStatuses() interface{} {
	return helpers.ModuleStatuses(PageModule())
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func PageMigrate() {
	_const.Services.DB.AutoMigrate(&Page{})
	_const.Services.DB.AutoMigrate(&PageI18ns{})
}

/*
* event run after add Page
 */
func (u *Page) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(PageModule())
	return
}

/*
* event run after delete Faq
 */
func (u *Page) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, PageModule())
	return
}

/**
* update status
 */
func (u *Page) BeforeUpdate() (err error) {
	var page Page
	if u.ID != 0 {
		_const.Services.DB.First(&page, u.ID)
		if page.Status != u.Status {
			DecreaseRow(page.Status, PageModule())
			IncreaseRow(u.Status, PageModule())
		}
	}
	return
}

/**
* you can update these column only
 */
func PageFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["status"] = "string"
	m["slug"] = "string"
	return m
}
func PageTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"

	return m
}

/**
* active Page only
 */
func ActivePage(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + _const.ACTIVE)
}
