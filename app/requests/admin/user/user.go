package user

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"strings"
)

/**
* validate store user request
 */
func Store() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"name":                []string{"required", "alpha_space_dash_underscore", "min:2", "max:50",},
		"email":               []string{"required", "min:6", "max:50", "email", "unique:users"},
		"password":            []string{"required", "min:6", "max:50"},
		"image":               []string{"max:255", "url"},
		"status":              []string{"required", "in:" + helpers.GetStatusSeparateWithComma()},
		"permission_group_id": []string{"required", "numeric", "min:1", "is_permission_group"},
	}

	messages := govalidator.MapData{
		"name":                []string{helpers.Required(), helpers.Alpha(), helpers.Min("2"), helpers.Max("50")},
		"email":               []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.Email(), helpers.Unique()},
		"password":            []string{helpers.Required(), helpers.Min("6"), helpers.Max("50")},
		"image":               []string{helpers.Max("255"), helpers.Url()},
		"status":              []string{helpers.Required(), helpers.Status()},
		"permission_group_id": []string{helpers.Required(), helpers.Numeric(), helpers.Min("1"), helpers.NotPermissionGroup()},
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
* validate update user request
 */
func Update() url.Values {
	splitUrl := strings.Split(_const.Request().RequestURI, "/")
	/// Validation rules
	rules := govalidator.MapData{
		"name":                []string{"required", "alpha_space_dash_underscore", "min:2", "max:50"},
		"email":               []string{"required", "min:6", "max:50", "email", "unique_update:users," + splitUrl[len(splitUrl)-1]},
		"password":            []string{"max:50"},
		"image":               []string{"max:255", "url"},
		"status":              []string{"required", "in:" + helpers.GetStatusSeparateWithComma() + "," + _const.BLOCK},
		"permission_group_id": []string{"required"},
	}

	messages := govalidator.MapData{
		"name":                []string{helpers.Required(), helpers.Alpha(), helpers.Min("2"), helpers.Max("50")},
		"email":               []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.Email(), helpers.TUnique("update")},
		"password":            []string{helpers.Max("50")},
		"image":               []string{helpers.Max("255"), helpers.Url()},
		"status":              []string{helpers.Required(), helpers.Status(_const.BLOCK)},
		"permission_group_id": []string{helpers.Required()},
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
* validate update user request
 */
func QuickEdit() url.Values {
	splitUrl := strings.Split(_const.Request().RequestURI, "/")
	/// Validation rules
	rules := govalidator.MapData{
		"name":  []string{"required", "alpha_space_dash_underscore", "min:2", "max:50"},
		"email": []string{"required", "min:6", "max:50", "email", "unique_update:users," + splitUrl[len(splitUrl)-1]},
	}

	messages := govalidator.MapData{
		"name":  []string{helpers.Required(), helpers.Alpha(), helpers.Min("2"), helpers.Max("50")},
		"email": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.Email(), helpers.TUnique("update")},
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
