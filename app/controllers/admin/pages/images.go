package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/page"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"strconv"
)

/**
* update image form page
 */
func UploadImage(g *gin.Context) {
	// init struct to validate request
	var row models.PageImageRequest
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := page.UploadStoreUpdate(&row)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	///get id
	id, _ := strconv.Atoi(g.Param("id"))
	// find this row or return 404
	page := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &page); page.ID == 0 {
		return
	}
	/// upload images and insert in database
	insertImageInDataBase(row.Images, id)
	/// get the new data with images
	newPage := Row()
	if models.FindOrFail(g.Param("id"), &row, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Translations", "lang = ?", helpers.LangHeader())
	}, func(db *gorm.DB) *gorm.DB {
		return db.Preload("Images")
	}); newPage.ID == 0 {
		return
	}
	helpers.OkResponse(helpers.DoneUpdate(models.PageModel()), transformers.PageResponse(newPage))
}

/***
* Delete images
* Delete image by id
 */
func DeleteImage(g *gin.Context) {
	// find this row or return 404
	row := RowImage()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	_const.Services.DB.Unscoped().Delete(&row)
	// now return ok response
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.PageModel()))
}

/***
* Delete images assign to page by page id
 */
func DeletePageImages(g *gin.Context) {
	// find this row or return 404
	row := Row()
	// check if this id exits , abort if not
	if models.FindOrFail(g.Param("id"), &row); row.ID == 0 {
		return
	}
	deleteAllPageImage(row.ID)
	// now return ok response
	helpers.OkResponseWithOutData(helpers.DoneDelete(models.PageModel()))
	return
}

/***
*  get all page image and delete
 */
func deleteAllPageImage(id uint) {
	rows := RowImage()
	_const.Services.DB.Unscoped().Where("page_id = ? ", id).Delete(&rows)
}

/**
* upload images
* loop and insert images with id
 */
func insertImageInDataBase(images []string, id int) {
	///// loop and insert image in database
	if len(images) > 0 {
		for _, upload := range images {
			_const.Services.DB.Create(&models.PageImage{
				PageId: id,
				Image:  upload,
			})
		}
	}

}
