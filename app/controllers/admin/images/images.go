package images

import (
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* update image form page
 */
func Store(g *gin.Context) {
	// validate Request
	if !validateRequest() {
		return
	}
	// init struct to validate request
	var row = ConvertRequestRow()

	result := UploadImages(row.Images)
	if len(result) == 0 {
		result = make([]string, 0)
	}
	// return uploaded images links
	helpers.OkResponse(helpers.DoneCreateItem(models.ImageModel()), result)
}

/***
* Delete images
* Delete image by id
 */
func Delete(g *gin.Context) {
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
