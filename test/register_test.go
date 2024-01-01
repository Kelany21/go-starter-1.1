package test

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"testing"
)

/**
* check if user has register with email before
 */
func TestRegisterWithValidCase(t *testing.T) {
	w, _ := registerNewUser(true)
	assert.Equal(t, 200, w.Code)
}

/**
* check if user has register with email before
 */
func TestRegisterWithExistEmail(t *testing.T) {
	w, data := registerNewUser(true)
	assert.Equal(t, 200, w.Code)
	k := postWitOutHeader(data, "register", false)
	assert.Equal(t, 409, k.Code)
}

/**
* Test Required inputs
 */
func TestRegisterRequireInputs(t *testing.T) {
	///not send email
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Password: "123457",
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
	///not send name
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "zizo199988@gmail.com",
		Password: "123457",
	}, "register", true)
	///not send password
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email: "zizo199988@gmail.com",
		Name:  "Abdel Aziz Hassan",
	}, "register", true)
}

/**
* Test not valid inputs
 */
func TestRegisterNotValidInputs(t *testing.T) {
	///not valid email
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "sasdasd",
		Password: "123457",
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
}

/**
* Test inputs limitaion
 */
func TestRegisterInputsLimitation(t *testing.T) {
	///min email fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "s@s.i",
		Password: "123457",
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
	///max email fail
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    helpers.RandomString(50) + "@gmail.com",
		Password: "123457",
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
	///max password fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "zizohassan@gmail.com",
		Password: helpers.RandomString(52),
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
	///min password fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "zizohassan@gmail.com",
		Password: helpers.RandomString(5),
		Name:     "Abdel Aziz Hassan",
	}, "register", true)
	///min name fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "zizohassan@gmail.com",
		Password: "123457",
		Name:     helpers.RandomString(3),
	}, "register", true)
	///max name fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email:    "zizohassan@gmail.com",
		Password: "123457",
		Name:     helpers.RandomString(55),
	}, "register", true)
}

func registerNewUser(migrate bool) (*httptest.ResponseRecorder, models.User) {
	data := models.User{
		Name:     "Abdel Aziz Hassan",
		Email:    "zizo199988@gmail.com",
		Password: "123456",
	}
	w := postWitOutHeader(data, "register", migrate)
	return w, data
}

func addNewUser() {
	data := models.User{
		Name:              "Abdel Aziz Hassan",
		Email:             "zizo199988@gmail.com",
		PermissionGroupId: _const.ADMIN_ID,
		Password:          "123456",
		Status:            _const.ACTIVE,
	}
	_const.Services.DB.Create(&data)

}
