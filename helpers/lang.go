package helpers

import (
	"github.com/bykovme/gotrans"
	_const "starter-golang-new/const"
)

func GetCurrentLang() string {
	return gotrans.DetectLanguage(_const.Services.GIN.GetHeader("Accept-Language"))
}

func GetCurrentLangFromHttp() string {
	return _const.Services.GIN.Request.Header.Get("Accept-Language")
}

func T(key ...string) string {
	s := ""
	for _, k := range key {
		s += gotrans.Tr(GetCurrentLang(), k)
	}
	return s
}
