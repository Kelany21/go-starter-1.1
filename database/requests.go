package database

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func GetAll(moduleName string, object interface{}, paginate int) {
	url := os.Getenv("DATABASE_URL") + "/" + os.Getenv("DATABASE_NAME") + "/get/" + moduleName + "?paginate=" + strconv.Itoa(paginate)
	jsonValue, _ := json.Marshal(object)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	client := &http.Client{}
	_, _ = client.Do(req)
	return
}

func GetByid(moduleName string, object interface{}, paginate int) {
	url := os.Getenv("DATABASE_URL") + "/" + os.Getenv("DATABASE_NAME") + "/get/" + moduleName + "?paginate=" + strconv.Itoa(paginate)
	jsonValue, _ := json.Marshal(object)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	client := &http.Client{}
	_, _ = client.Do(req)
	return
}
