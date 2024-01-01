package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	_const "starter-golang-new/const"
)

/**
* here we set the default language
* if we have Accept-Language in header we will work on it
* else we will pass the default language
 */
func Language() gin.HandlerFunc {
	return func(g *gin.Context) {

		_const.Services.GIN = g
		_const.Services.GinCopy = g.Copy()

		//// set gin object and get copy
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(g.Request.Body)


		//// get body
		_const.Services.GIN.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		_const.Services.GinCopy.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		_const.Services.RequestBody = bodyBytes

		lang := g.GetHeader("Accept-Language")
		if _, ok := _const.SupportedLang()[lang]; !ok {
			g.Request.Header.Set("Accept-Language", os.Getenv("DEFAULT_LANG"))
		}
		if lang == "" {
			g.Request.Header.Set("Accept-Language", os.Getenv("DEFAULT_LANG"))
		}
		g.Next()
	}
}
