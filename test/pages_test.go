package test

import (
	"github.com/stretchr/testify/assert"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"testing"
)

///// show all case
func TestPagesShowAll(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get("admin/pages", false, token)
	assert.EqualValues(t, 0, returnResponseKey(k, "data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestPagesFilter(t *testing.T) {
	pageUrl := "admin/pages"
	token := getTokenAsHeader(true)
	w := newPage()
	filter(t, pageUrl, w.Status, "status", "equal", token)
	//filter(t, pageUrl, w.Name, "name", "equal", token)
	filter(t, pageUrl, w.Status, "status", "not-equal", token)
}

///// show function cases
func TestPagesShowWithValidId(t *testing.T) {
	pageUrl := "admin/pages"
	token := getTokenAsHeader(true)
	//	w := newPage()
	k := get(pageUrl+"/1", false, token)
	//assert.Equal(t, w.Name, returnResponseKey(k, "data.name"))
	assert.Equal(t, 200, k.Code)
}

func TestPagesShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get("admin/pages/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestPagesUpdateWithValidData(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newPage()
	var oldRow models.Page
	_const.Services.DB.First(&oldRow)
	data := models.Page{
		//Name:   "New Data",
		Status: "activate",
		Image:  images(5),
	}
	k := put(data, "admin/pages/1", false, token)
	//	assert.Equal(t, data.Name, returnResponseKey(k, "data.name"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestPagesRequireInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	url := "admin/pages/1"
	///not send name
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		//Name: helpers.RandomString(10),
	}, url, false, token)
	///not send status
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		Status: "activate",
	}, url, false, token)
}

/**
* Test inputs limitaion
 */
func TestPagesInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(true)
	///create new category
	newPage()
	url := "admin/pages/1"
	///min name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		//Name:   helpers.RandomString(2),
		Status: "activate",
	}, url, false, token)
	///max name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		//Name:   helpers.RandomString(80),
		Status: "activate",
	}, url, false, token)
	///max status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		//Name:   helpers.RandomString(10),
		//Status: 3,
	}, url, false, token)
	///min status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Page{
		//Name:   helpers.RandomString(10),
		Status: "",
	}, url, false, token)

}

func newPage() models.Page {
	data := models.Page{
		//:   "Home",
		Status: "activate",
	}
	_const.Services.DB.Create(&data)
	return data
}
