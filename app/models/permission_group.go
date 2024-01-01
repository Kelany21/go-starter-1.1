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
type PermissionGroup struct {
	gorm.Model
	Status string `gorm:"type:varchar(20);" json:"status"`
}

type PermissionGroupI18ns struct {
	gorm.Model
	ReferenceId uint64 `gorm:"type:int(20);" json:"reference_id"`
	FiledName   string `gorm:"type:varchar(20);" json:"filed_name"`
	Language    string `gorm:"type:varchar(5);" json:"language"`
	Value       string `gorm:"type:text;" json:"value"`
}

/*
* return with module name
 */
func PermissionGroupModule() string {
	return helpers.ModuleName("permission_groups")
}

func PermissionGroupTransTable() string {
	return helpers.ModuleTransTable(PermissionGroupModule())
}

func PermissionGroupTranslateFillable() map[string]string {
	var m = make(map[string]string)
	m["name"] = "string"
	return m
}

/*
* return with model name
 */
func PermissionGroupModel() string {
	return helpers.ModelModel(PermissionGroupModule())
}

/*
* return with route name
 */
func PermissionGroupRoute() string {
	return helpers.ModuleRoute(PermissionGroupModule())
}

/*
* return with Statuses
 */
func PermissionGroupStatuses() interface{} {
	return helpers.ModuleStatuses(PermissionGroupModule())
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func PermissionGroupMigrate() {
	_const.Services.DB.AutoMigrate(&PermissionGroup{})
	_const.Services.DB.AutoMigrate(&PermissionGroupI18ns{})
}

/*
* event run after add Page
 */
func (u *PermissionGroup) AfterCreate(scope *gorm.Scope) (err error) {
	IncreaseOnCreate(PermissionGroupModule())
	return
}

/*
* event run after delete Faq
 */
func (u *PermissionGroup) AfterDelete(tx *gorm.DB) (err error) {
	DecreaseOnDelete(u.Status, PermissionGroupModule())
	return
}

/**
* update status
 */
func (u *PermissionGroup) BeforeUpdate() (err error) {
	var permissionGroup PermissionGroup
	if u.ID != 0 {
		_const.Services.DB.First(&permissionGroup, u.ID)
		if permissionGroup.Status != u.Status {
			DecreaseRow(permissionGroup.Status, PermissionGroupModule())
			IncreaseRow(u.Status, PermissionGroupModule())
		}
	}
	return
}

/**
* you can update these column only
 */
func PermissionGroupFillAbleColumn() map[string]string {
	var m = make(map[string]string)
	m["status"] = "string"

	return m
}

/**
* active Page only
 */
func ActivePermissionGroup(db *gorm.DB) *gorm.DB {
	return db.Where("status = " + _const.ACTIVE)
}
