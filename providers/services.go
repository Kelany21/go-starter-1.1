package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	_const "starter-golang-new/const"
	"strings"
)

func StartContainer() {
	/// start server container
	_const.Services = _const.ServicesInit()
	/// inject modules in service container
	InitModules()
}

/**
* init modules
 */
func InitModules() {
	container := make(map[string]map[string]interface{}, 0)
	listAllModules := listAllModules()
	for moduleName, modulePath := range listAllModules {
		name := strings.Replace(moduleName, ".json", "", -1)
		container[name] = readFiles(modulePath)
	}
	_const.Services.Modules = container
}

/**
* open single file and get content
 */
func readFiles(filePath string) map[string]interface{} {
	var m = make(map[string]interface{})
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("error reading files in file : ", filePath, err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &m)
	return m
}

/**
* return list of module files
 */
func listAllModules() map[string]string {
	var files = make(map[string]string)
	root := _const.DIR + "/public/modules"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".json") {
			files[info.Name()] = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
