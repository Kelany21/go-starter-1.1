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
type PageImage struct {
	gorm.Model
	Image  string `gorm:"type:varchar(255);" json:"image"`
	PageId int    `gorm:"type:int" json:"page_id"`
}

/**
* request images
 */
type PageImageRequest struct {
	Images []string `json:"images"`
}

/**
* migration function must be the file name concat with Migrate
* key word Example : user will be UserMigrate
 */
func PageImageMigrate() {
	_const.Services.DB.AutoMigrate(&PageImage{})
}

/**
* you can update these column only
 */
func PageImageFillAbleColumn() []string {
	return []string{"image", "page_id"}
}
