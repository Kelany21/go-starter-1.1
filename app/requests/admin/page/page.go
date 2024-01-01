package page

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* validate update page request
 */
func StoreUpdate() url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		"name":   []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
		"status": []string{"required", "in:" + helpers.GetStatusSeparateWithComma()},
		"slug":   []string{"required", "min:3", "max:20"},
	}

	messages := govalidator.MapData{
		"name":   []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
		"status": []string{helpers.Required(), helpers.Status(helpers.GetStatusSeparateWithComma())},
		"slug":   []string{helpers.Required(), helpers.Min("3"), helpers.Max("20")},
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
* validate upload Image
 */
func UploadStoreUpdate(request *models.PageImageRequest) url.Values {
	/// Validation rules
	rules := govalidator.MapData{
		//"file:images": []string{"ext:jpg,png,jpeg", "size:100000", "mime:image/jpg,image/png,image/jpeg", "required"},
		"images": []string{"required"},
	}
	messages := govalidator.MapData{
		//"file:images":   []string{helpers.Required(lang), helpers.Ext(lang, "jpg,png,jpeg") , helpers.Mime(lang, "image/jpg,image/png,image/jpeg"), helpers.Size(lang, "100000")},
		"images": []string{helpers.Required()},
	}

	//data := make(map[string]interface{}, 0)
	opts := govalidator.Options{
		Request:  _const.Request(), // request object
		Rules:    rules,            // rules map
		Messages: messages,         // custom message map (Optional)
		Data:     request,
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
		"name": []string{"required", "lang", "len:" + _const.Services.SupportedLanguageCountString},
	}

	messages := govalidator.MapData{
		"name": []string{helpers.Required(), helpers.Min("6"), helpers.Max("50"), helpers.LenSlice(_const.Services.SupportedLanguageCountString)},
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
