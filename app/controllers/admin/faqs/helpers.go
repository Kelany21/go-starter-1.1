package faqs

import (
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/admin/faq"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* constructor
 */
func Row() models.Faq {
	return models.Faq{}
}

/**
* constructor with Request
 */
func ConvertRequestRow() models.Faq {
	request := models.ConvertBodyToHashMap()
	faq1 := models.Faq{
		Status: request["status"].(string),
	}
	return faq1
}
func RowAnswer() models.Answer {
	return models.Answer{}
}

/**
* constructor with Request
 */
func ConvertRequestRowAnswer() models.Answer {
	request := models.ConvertBodyToHashMap()

	answer := models.Answer{
		FaqId: int(request["faq_id"].(float64)),
	}
	return answer
}

/**
* constructor Array
 */
func Rows() []models.Faq {
	return []models.Faq{}
}

/**
* filter module with some columns
 */
func filter() []string {
	g := _const.Services.GIN
	var filter []string
	if g.Query("question") != "" {
		filter = append(filter, `question like "%`+g.Query("question")+`%"`)
	}
	if g.Query("status") != "" {
		if g.Query("status") != "all" {
			filter = append(filter, `status = "`+g.Query("status")+`"`)
		}
	}
	return filter
}

/**
* preload module with some preload conditions
 */
func preload() []string {
	return []string{"Answers"}
}

/**
* here we will check if request valid or not
 */
func validateRequest() bool {
	// init struct to validate request

	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := faq.StoreUpdate()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	return true
}
func validateRequestAnswer() bool {
	// init struct to validate request

	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := faq.StoreUpdateAnswer()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return false
	}
	return true
}

func store() (map[string]interface{}, models.Faq) {
	row := ConvertRequestRow()
	_const.Services.DB.Create(&row)
	return storeUpdateData(&row), row
}
