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
type Faq struct {
	gorm.Model
	Status  string   `gorm:"type:varchar(20);" json:"status"`
	Answer  []string `gorm:"-" json:"answer"`
	Answers []Answer `gorm:"association_autoupdate:false;association_autocreate:false" json:"answers"`
}

type FaqI18ns struct {
	gorm.Model
	ReferenceId uint64 `gorm:"type:int(20);" json:"reference_id"`
	FiledName   string `gorm:"type:varchar(20);" json:"filed_name"`
	Language    string `gorm:"type:varchar(5);" json:"language"`
	Value       string `gorm:"type:text;" json:"value"`
}

/*
* return with module name
 */
func FaqModule() string {
	return helpers.ModuleName("faqs")
}
func FaqTable() string {
	return helpers.ModuleTable(FaqModule())
}

/*
* return with route name
 */
func FaqRoute() string {
	return helpers.ModuleRoute(FaqModule())
}

/*
* return with model name
 */
func FaqModel() string {
	return helpers.ModelModel(FaqModule())
}

/*
* return with Statuses
 */
func FaqStatuses() interface{} {
	return helpers.ModuleStatuses(FaqModule())
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func FaqMigrate() {
	_const.Services.DB.AutoMigrate(&Faq{})
	_const.Services.DB.AutoMigrate(&FaqI18ns{})
}

/*
* event run after add Category
 */
func (u *Faq) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(FaqModule())
	return
}

/*
* event run after delete Faq
 */
func (u *Faq) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, FaqModule())
	return
}

/**
* update status
 */
func (u *Faq) BeforeUpdate() (err error) {
	var faq Faq
	if u.ID != 0 {
		_const.Services.DB.First(&faq, u.ID)
		if faq.Status != u.Status {
			DecreaseRow(faq.Status, FaqModule())
			IncreaseRow(u.Status, FaqModule())
		}
	}
	return
}

/**
* you can update these column only
 */
func FaqFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["status"] = "string"

	return m

}
func FaqTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["question"] = "string"
	return m
}

func FaqTransTable() string {
	return helpers.ModuleTransTable(FaqModule())
}
/**
* active questions only
 */
func ActiveFaq(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + _const.ACTIVE)
}
