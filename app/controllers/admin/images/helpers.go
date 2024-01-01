package images

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/image"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.ImageRequest {
	return models.ImageRequest{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.ImageRequest {
	request := models.ConvertBodyToHashMap()
	pageImageRequest1 := models.ImageRequest{
		Images: helpers.ConvertArrayInterfacesToArrayStrings(request["images"].([]interface{})),
	}
	return pageImageRequest1
}
func RowImage() models.PageImage {
	return models.PageImage{}
}

/**
* upload images
* loop and insert images with id
 */
func UploadImages(images []string) []string {
	return helpers.MultiDecodeImage(images)
}

/**

/**
* here we will check if request valid or not
*/
func validateRequest() bool {

	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := image.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	return true
}
