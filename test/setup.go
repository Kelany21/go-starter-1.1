package test

import (
	"bytes"
	"encoding/json"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"go-backend/provider"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests"
	"starter-golang-new/config"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"starter-golang-new/providers"
)

var baseTestUrl = "http://localhost:9090/"
var check = true

/**
* define struct that carry the arg of request
 */
type RequestData struct {
	Migrate     bool
	RequestType string
	Url         string
	Data        interface{}
	Header      map[string]string
}

/**
* init gin and return gin engine
 */
func setupRouter(migrate bool) *gin.Engine {

	helpers.LoadEnv()
	_, b, _, _ := runtime.Caller(0)
	if os.Getenv("APP_ENV") == "local" {
		_const.DIR = filepath.Dir(b)
	} else {
		_const.DIR = os.Getenv("PUBLIC_PATH")
	}
	err := gotrans.InitLocales("public/trans")
	if err != nil {
		panic(err)
	}
	providers.StartContainer()

	requests.Init()

	config.ConnectToDatabase()
	if migrate {
		models.MigrateAllTable()
	}
	provider.Seed()
	/// start gin
	r := providers.Gin()
	return providers.Routing(r)
}

/**
* post request
 */
func post(data interface{}, url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "POST",
		Data:        data,
		Header:      headers,
	})
}

func postWitOutHeader(data interface{}, url string, migrate bool) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "POST",
		Data:        data,
	})
}

/**
* Put request
 */
func put(data interface{}, url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "PUT",
		Data:        data,
		Header:      headers,
	})
}

func putWithOutHeader(data interface{}, url string, migrate bool) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "PUT",
		Data:        data,
	})
}

/**
* Get request
 */
func get(url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "GET",
		Header:      headers,
	})
}

func getWithOutHeader(url string, migrate bool) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "GET",
	})
}

/**
* Get request
 */
func deleter(url string, migrate bool, headers map[string]string) *httptest.ResponseRecorder {
	url = baseTestUrl + url
	return request(RequestData{
		Url:         url,
		Migrate:     migrate,
		RequestType: "DELETE",
		Header:      headers,
	})
}

/**
* Make new request
 */
func request(request RequestData) *httptest.ResponseRecorder {
	router := setupRouter(request.Migrate)
	w := httptest.NewRecorder()
	sendData, _ := json.Marshal(&request.Data)
	req, _ := http.NewRequest(request.RequestType, request.Url, bytes.NewReader(sendData))
	if len(request.Header) > 0 {
		for headerName, headerValue := range request.Header {
			req.Header.Set(headerName, headerValue)
		}
	}
	router.ServeHTTP(w, req)
	return w
}

/**
* return response as json
 */
func responseData(c io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c)
	return buf.String()
}
