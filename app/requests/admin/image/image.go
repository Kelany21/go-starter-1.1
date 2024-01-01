package image

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate store image request
 */
func StoreUpdate() url.Values {
	/// Validation rules

	rules := govalidator.MapData{
		"images": []string{"required", "strings_slice"},
	}
	messages := govalidator.MapData{
		"images": []string{helpers.Required(), helpers.StringsSlice()},
	}

	data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Request:  _const.Request(), // request object
		Rules:    rules,            // rules map
		Messages: messages,         // custom message map (Optional)
		Data:     &data,
	}
	vd := govalidator.New(opts)

	return vd.ValidateJSON()
}
