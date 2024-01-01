package helpers

import (
	"encoding/json"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
	"os"
	_const "starter-golang-new/const"
)

/**
* conflict
 */
func ReturnBadRequest() {
	var errors map[string]string
	var data map[string]interface{}
	var msg = gotrans.Tr(GetCurrentLang(), "400")
	response(msg, data, errors, 400, 400, false)
	return
}

/**
* Duplicate data
 */
func ReturnDuplicateData(inputName string) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T("duplicate_data_part_one", inputName, "duplicate_data_part_two")
	response(msg, data, errors, 409, 409, false)
	return
}

/**
* Duplicate data
 */
func YouCantTrashOrDeleteYourSelf() {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T("you_can_not_trash_or_delete_your_self")
	response(msg, data, errors, 409, 409, false)
	return
}

/**
* upload error
 */
func UploadError() {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T("upload_error_code")
	response(msg, data, errors, 415, 415, false)
	return
}

/**
* multi upload error
 */
func MultiUploadError() {
	var errors map[string]string
	var data map[string]interface{}
	var msg = T("upload_multi_images_error_code")
	response(msg, data, errors, 415, 415, false)
	return
}

/**
* NotValidRequest response
 */

func ReturnNotValidRequest(error url.Values) bool {
	if len(error) > 0 {
		_const.Services.GIN.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(), "400"),
				"errors":  error,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotValidRequest response
 */

func TestReturnNotValidRequest(e url.Values) bool {
	if len(e) > 0 {
		_const.Services.GIN.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(), "400"),
				"errors":  e,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotValidFile response
 */
func ReturnNotValidRequestFile(error *govalidator.Validator) bool {
	e := error.Validate()
	if len(e) > 0 {
		_const.Services.GIN.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(), "400"),
				"errors":  e,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotValidRequest response
 */
func ReturnNotValidRequestFormData(error *govalidator.Validator) bool {
	e := error.Validate()
	if len(e) > 0 {
		_const.Services.GIN.JSON(
			http.StatusBadRequest, gin.H{
				"status":  false,
				"message": gotrans.Tr(GetCurrentLang(), "400"),
				"errors":  e,
				"code":    400,
				"payload": nil,
			})
		return true
	}
	return false
}

/**
* NotFound response
 */
func ReturnNotFound(msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(msg, data, errors, http.StatusNotFound, 404, false)
	return
}

/**
* NotFound response
 */
func ReturnNotFoundUser(msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(msg, data, errors, http.StatusBadRequest, 400, false)
	return
}

/**
* item Action response
 */
func ActionItem(msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* Forbidden response
 */
func ReturnForbidden(msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(msg, data, errors, http.StatusForbidden, 403, false)
	return
}

/**
* ok response with data
 */
func OkResponse(msg string, data interface{}) {
	var errors map[string]string
	response(msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* ok response without data
 */
func OkResponseWithOutData(msg string) {
	var errors map[string]string
	var data map[string]interface{}
	response(msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* Not Authorize
 */
func ReturnYouAreNotAuthorize() {
	var errors map[string]string
	var data map[string]interface{}
	var msg = gotrans.Tr(GetCurrentLang(), "401")
	response(msg, data, errors, 401, 401, true)
	return
}

/**
* ok with paging
 */
func OkResponseWithPaging(msg string, data *Paginator) {
	var errors map[string]string
	response(msg, data, errors, http.StatusOK, 200, true)
	return
}

/**
* stander response
 */
func response(msg string, data interface{}, errors map[string]string, httpStatus int, code int, status bool) {
	m := make(map[string]interface{})
	jData := []byte{}
	m["status"] = status
	m["message"] = msg
	m["errors"] = errors
	m["code"] = status
	m["payload"] = data
	if os.Getenv("APP_ENV") != "local" {
		jData, _ = json.Marshal(&m)
		_const.Services.GIN.JSON(httpStatus, jData)
		return
	}

	_const.Services.GIN.JSON(httpStatus, m)

	return
}

/**
* NotValidRequest file
 */
func ReturnNotValidFile(err error) {
	_const.Services.GIN.JSON(
		http.StatusBadRequest, gin.H{
			"status":  false,
			"message": gotrans.Tr(GetCurrentLang(), "400"),
			"errors":  err,
			"code":    400,
			"payload": nil,
		})
}

/**
*  global response
 */
func ReturnResponseWithMessageAndStatus(statusHttp int, message string, status bool) {
	var errors map[string]string
	var data map[string]interface{}
	var msg = gotrans.Tr(GetCurrentLang(), message)
	response(msg, data, errors, statusHttp, statusHttp, status)
	return
}
