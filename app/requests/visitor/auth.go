package visitor

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate login request
 */
func Login() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"required", "min:6", "max:50"},
	}

	messages := govalidator.MapData{
		"email":    []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.Email()},
		"password": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50")},
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

/**
* validate register request
 */
func Register() url.Values {

	/// Validation rules
	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:50", "email", "unique:users"},
		"name":     []string{"required", "min:4", "max:50"},
		"password": []string{"required", "min:6", "max:50"},
	}

	messages := govalidator.MapData{
		"email":    []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.Email()},
		"name":     []string{helpers.Required(), helpers.Min("6"), helpers.Max("50")},
		"password": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50")},
	}
	data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Data:     &data,
		Request:  _const.Request(), // request object
		Rules:    rules,            // rules map
		Messages: messages,
	}

	vd := govalidator.New(opts)

	return vd.ValidateJSON()
}

/**
* validate Reset request
 */
func Reset() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"email": []string{"required", "min:6", "max:50", "email"},
	}
	data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Request: _const.Request(), // request object
		Rules:   rules,            // rules map
		Data:    &data,
	}

	vd := govalidator.New(opts)

	return vd.ValidateJSON()
}

/**
* validate Recover request
 */
func Recover() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"token":    []string{"required"},
		"password": []string{"required", "between:6,20"},
	}
	data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Request: _const.Request(), // request object
		Rules:   rules,            // rules map
		Data:    &data,
	}

	vd := govalidator.New(opts)

	return vd.ValidateJSON()
}
