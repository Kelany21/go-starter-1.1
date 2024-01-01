package category

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate store category request
 */
func StoreUpdate() url.Values {
	/// Validation rules

	rules := govalidator.MapData{
		"name":        []string{"required", "lang","lang_min:6","lang_max:50", "len:" + _const.Services.SupportedLanguageCountString},
		"description": []string{"required", "lang", "lang_min:6","lang_max:50","len:" + _const.Services.SupportedLanguageCountString},
		"status":      []string{"required", "in:" + helpers.GetStatusSeparateWithComma()},
	}

	messages := govalidator.MapData{
		"name":        []string{helpers.Required(), helpers.LangMin("6"), helpers.LangMax("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"description": []string{helpers.Required(), helpers.LangMin("6"), helpers.LangMax("500"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"status":      []string{helpers.Required(), helpers.Status()},
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
* validate quick-edit category request
 */
func QuickEdit() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"name":        []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
		"description": []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
	}

	messages := govalidator.MapData{
		"name":        []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"description": []string{helpers.Required(), helpers.Min("6"), helpers.Max("500"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
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
