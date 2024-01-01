package setting

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate store setting request
 */
func StoreUpdate() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"twitter":  []string{"required", "url"},
		"facebook": []string{"required", "url"},
		"youtube":  []string{"required", "url"},
		"linkedin": []string{"required", "url"},
	}

	messages := govalidator.MapData{
		"twitter":  []string{helpers.Required(), helpers.Url()},
		"facebook": []string{helpers.Required(), helpers.Url()},
		"youtube":  []string{helpers.Required(), helpers.Url()},
		"linkedin": []string{helpers.Required(), helpers.Url()},
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
