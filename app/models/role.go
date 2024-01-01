package models

import (
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
)

/***
* model struct here we will build the main
* struct that connect to database
* status 1 active 2 is not active
 */
type Role struct {
	gorm.Model
	FuncName          string `gorm:"type:varchar(50);" json:"func_name"`
	PermissionGroupId int    ` json:"permission_group_id"`
	Module            string `json:"module"`
	Slug              string `json:"slug"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func RoleMigrate() {
	_const.Services.DB.AutoMigrate(&Role{})
}

/**
* you can update these column only
 */
func RoleFillAbleColumn() []string {
	return []string{"func_name", "permission_group_id", "module", "slug"}
}
