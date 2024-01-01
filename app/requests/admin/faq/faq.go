package faq

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate store faq request
 */
func StoreUpdate() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"question": []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
		"status":   []string{"required", "in:" + helpers.GetStatusSeparateWithComma()},
	}

	messages := govalidator.MapData{
		"question": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"status":   []string{helpers.Required(), helpers.Status(helpers.GetStatusSeparateWithComma())},
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

func StoreUpdateAnswer() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"text":   []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
		"faq_id": []string{"required", "min:1"},
	}

	messages := govalidator.MapData{
		"question": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"faq_id":   []string{helpers.Required(), helpers.Min("1")},
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
* validate QuickEdit faq request
 */
func QuickEdit() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"question": []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
	}

	messages := govalidator.MapData{
		"question": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
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
* validate store faq request
 */
func StoreUpdateFaqWithArray() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"question": []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
		"status":   []string{"required", "in:" + helpers.GetStatusSeparateWithComma()},
		"answers":  []string{"required", "answers:2,50"},
	}

	messages := govalidator.MapData{
		"question": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"status":   []string{helpers.Required(), helpers.Status(helpers.GetStatusSeparateWithComma())},
		"answers":  []string{helpers.Required()},
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
