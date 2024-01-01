package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"reflect"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

type ModelFunction func(db *gorm.DB) *gorm.DB

/**
* id interface you can assign id as string or unit or int
* structBind interface you can assign any type of model
* we select the first id based on struct with gorm
* then  retrieve id from struct with refection
* if we not found id we will abort gin and return
 */
func FindOrFail(id interface{}, structBind interface{}, appendFunction ...ModelFunction) {
	g := _const.Services.GIN
	appendFunctionsToQuery(appendFunction).Where("id = ? ", id).First(structBind)
	findId := reflect.ValueOf(structBind).Elem().FieldByName("ID").Uint()
	if findId == 0 {
		helpers.ReturnNotFound(helpers.ItemNotFound())
		g.Abort()
		return
	}
}

/**
* loop and handel where cases
* and preload
 */
func appendFunctionsToQuery(functions []ModelFunction) *gorm.DB {
	db := _const.Services.DB
	if len(functions) > 0 {
		for _, function := range functions {
			db = function(db)
		}
	}
	return db
}

/***
* short hand for find by struct
* it will be useful when you want to update
 */
func FindS(structBind interface{}, appendFunction ...ModelFunction) {
	appendFunctionsToQuery(appendFunction).First(structBind)
}

/***
* short hand for find by id
 */
func Find(id interface{}, structBind interface{}, appendFunction ...ModelFunction) {
	appendFunctionsToQuery(appendFunction).Where("id = ? ", id).First(structBind)
}

/**
* short hand to update data with fill able data
* then return with the new data
 */
func Update(data interface{}, row interface{}, allowColumns []string, preloads ...string) {
	onlyAllowData := helpers.UpdateOnlyAllowColumns(data, allowColumns)
	db := _const.Services.DB.Model(row).Updates(onlyAllowData)
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
	}
	db.Scan(row)
}

func GetModuleStatuses(funcName interface{}) {
	funcName.(func())()
}

/**
* this will take columns , go throw request and insert in
* database
 */
func Store(table string, columns map[string]string, object interface{}, data ...map[string]interface{}) interface{} {
	var d = make(map[string]interface{})
	if len(data) > 0 {
		d = data[0]
	} else {
		body := _const.GetBodyB()
		_ = json.Unmarshal(body, &d)
	}
	if len(columns) > 0 {
		m := make(map[string]interface{})
		for name, value := range d {
			//check if key are lang or normal attr
			if _, ok := columns[name]; ok {
				m[name] = value
			}
		}
		if helpers.InsertQuery(table, m) {
			t := object
			_const.Services.DB.Table(table).Last(t)
			return t
		}
	}
	return nil
}

/**
* convert body to hash map
 */
func ConvertBodyToHashMap() map[string]interface{} {
	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	return d
}

/**
* this will take columns , go throw request and insert in
* database
 */
func UpdateTest(table string, columns map[string]string, object interface{}, id interface{}, data ...map[string]interface{}) interface{} {
	var d = make(map[string]interface{})
	if len(data) > 0 {
		d = data[0]
	} else {
		body := _const.GetBodyB()
		_ = json.Unmarshal(body, &d)
	}
	if len(columns) > 0 {
		m := make(map[string]interface{})
		for name, value := range d {
			//check if key are lang or normal attr
			if _, ok := columns[name]; ok {
				m[name] = value
			}
		}
		_const.Services.DB.Table(table).Where("id = ?", id).UpdateColumns(m)
		_const.Services.DB.Table(table).Where("id = ?", id).First(object)
		return object
	}

	return nil
}

/**
* this will take all fields and go throw request
* figure out this columns and insert in translate table
 */
func StoreTranslateColumn(object map[string]interface{}, table string, columns map[string]string, data ...map[string]interface{}) []helpers.TranslationScan {
	var d = make(map[string]interface{})
	if len(data) > 0 {
		d = data[0]
	} else {
		body := _const.GetBodyB()
		_ = json.Unmarshal(body, &d)
	}
	id := object["id"]
	helpers.DeleteRowsTranslations(table, id)
	for name, _ := range columns {
		//check if key are lang or normal attr
		if _, ok := d[name]; ok {
			// if key is lang we expect that value will be array
			v := d[name].([]interface{})
			if len(v) > 0 {
				for index, filedValue := range v {
					m := make(map[string]interface{})
					m["value"] = filedValue.(string)
					m["language"] = _const.SupportedLangSlice()[index]
					m["reference_id"] = id
					m["filed_name"] = name
					helpers.InsertQuery(table, m)
				}
			}
		}
	}

	return helpers.RowsTranslations(table, id)
}
