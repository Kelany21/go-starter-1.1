package helpers

import (
	"github.com/bykovme/gotrans"
)

func Required() string {
	return "required:" + gotrans.Tr(GetCurrentLang(), "required")
}

func StringsSlice() string {
	return "strings_slice:" + gotrans.Tr(GetCurrentLang(), "strings_slice")
}

func LenSlice(number string) string {
	return "len:" + gotrans.Tr(GetCurrentLang(), "len") + " " + number
}
func IntSlice() string {
	return "int_slice:" + gotrans.Tr(GetCurrentLang(), "strings_slice")
}

func Email() string {
	return "email:" + gotrans.Tr(GetCurrentLang(), "email_not_valid")
}

func Min(number string) string {
	return "min:" + gotrans.Tr(GetCurrentLang(), "min") + " " + number
}
func LangMin(number string) string {
	return "lang_min:" + gotrans.Tr(GetCurrentLang(), "min") + " " + number
}
func LangMax(number string) string {
	return "lang_max:" + gotrans.Tr(GetCurrentLang(), "max") + " " + number
}

func Status(status ...string) string {
	s := GetStatusSeparateWithComma()
	if len(status) > 0 {
		for _, m := range status {
			s += "," + m
		}

	}
	return "in:" + gotrans.Tr(GetCurrentLang(), "status_must_be") + " " + s
}

func In(ableStrings ...string) string {
	returnString := "( "
	for i := 0; i < len(ableStrings)-1; i++ {
		returnString += ableStrings[i] + ", "
	}
	returnString += ableStrings[len(ableStrings)-1] + " )"
	return "in:" + gotrans.Tr(GetCurrentLang(), "in") + " " + returnString
}

func Ext(extentions string) string {
	return "ext:" + gotrans.Tr(GetCurrentLang(), "ext") + " " + extentions
}

func Mime(extentions string) string {
	return "mime:" + gotrans.Tr(GetCurrentLang(), "ext") + " " + extentions
}

func Size(size string) string {
	return "size:" + gotrans.Tr(GetCurrentLang(), "size") + " " + size
}

func Numeric() string {
	return "numeric:" + gotrans.Tr(GetCurrentLang(), "numeric")
}

func Digits() string {
	return "digits:" + gotrans.Tr(GetCurrentLang(), "numeric")
}

func Url() string {
	return "url:" + gotrans.Tr(GetCurrentLang(), "url")
}

func Bool() string {
	return "boolean:" + gotrans.Tr(GetCurrentLang(), "boolean")
}

func Max(number string) string {
	return "max:" + gotrans.Tr(GetCurrentLang(), "max") + " " + number
}

func Unique(key ...string) string {
	s := ""
	for _, k := range key {
		s += gotrans.Tr(GetCurrentLang(), k)
	}
	return s
}

func TUnique(action string) string {
	if action == "store" {
		return "unique:" + gotrans.Tr(GetCurrentLang(), "unique")
	}else {
		return "unique_update:" + gotrans.Tr(GetCurrentLang(), "unique")
	}
}

func Between(number string) string {
	return "between:" + gotrans.Tr(GetCurrentLang(), "between") + " " + number
}

func NotValidExt() string {
	return "ext:" + gotrans.Tr(GetCurrentLang(), "error_read_file")
}

func NotPermissionGroup() string {
	return "is_permission_group:" + gotrans.Tr(GetCurrentLang(), "not_permission_group")
}

func Alpha() string {
	return "alpha_space_dash_underscore:" + gotrans.Tr(GetCurrentLang(), "alpha_space_dash_underscore")
}
