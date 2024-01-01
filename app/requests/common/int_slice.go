package common

import (
	"github.com/thedevsaddam/govalidator"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate store category request
 */
func Bulk(request *_const.IDS) *govalidator.Validator {

	/// Validation rules
	rules := govalidator.MapData{
		"ids": []string{"int_slice"},
	}

	messages := govalidator.MapData{
		"ids": []string{helpers.IntSlice()},
	}

	opts := govalidator.Options{
		Request:         _const.Request(), // request object
		Rules:           rules,            // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	return govalidator.New(opts)
}