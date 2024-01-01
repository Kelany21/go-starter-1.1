package helpers

import (
	"encoding/json"
	"net/http"
	_const "starter-golang-new/const"
)

/**
* decode body then return interface
 */
func DecodeAndReturn(req *http.Request, structBind interface{}) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&structBind)
	if err != nil {
		return structBind, err
	}
	return structBind, nil
}

/**
* get lang header
 */
func LangHeader() string {
	s := ""
	s = _const.Services.GIN.Request.Header.Get("Accept-Language")
	return s
}
