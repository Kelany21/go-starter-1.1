package helpers

import (
	"github.com/jinzhu/inflection"
	"reflect"
	_const "starter-golang-new/const"
	"strings"
)

/**
* get table name based on object
 */
func GetTableNameFormStruct(model interface{}) string {
	return strings.ToLower(inflection.Plural(GetModelNameFromStruct(model)))
}

func GetModelNameFromStruct(model interface{}) string {
	m := reflect.TypeOf(model).String()
	m = strings.Replace(m, "*", "", -1)
	stringsKey := strings.Split(m, ".")
	return stringsKey[1]
}

/***
* increase 1 on column or with some conditions
 */
func Increase(tableName string, columnName string, id interface{}, increaseNumber interface{}, where ...string) {
	len := len(where)
	if len > 0 {
		db := "UPDATE " + tableName + " SET " + columnName + " = " + columnName + " + " + increaseNumber.(string) + " WHERE "
		for i, w := range where {
			db += ` ` + w + ` `
			if (len - 1) != i {
				db += ` AND`
			}
		}
		_const.Services.DB.Exec(db)
		return
	}
	_const.Services.DB.Exec("UPDATE "+tableName+" SET "+columnName+" = "+columnName+" + "+increaseNumber.(string)+" WHERE id = ?", id)
}

/***
* Decrease 1 on column or with some conditions
 */
func Decrease(tableName string, columnName string, id interface{}, decreaseNumber interface{}, where ...string) {
	check := columnName + " > 0 AND " + columnName + " >= " + decreaseNumber.(string)
	if len(where) > 0 {
		db := "UPDATE " + tableName + " SET " + columnName + " = " + columnName + " - " + decreaseNumber.(string) + " WHERE "
		for _, w := range where {
			db += ` ` + w + ` `
			db += ` AND `
		}
		db += check
		_const.Services.DB.Exec(db)
		return
	}
	_const.Services.DB.Exec("UPDATE "+tableName+" SET "+columnName+" = "+columnName+" - "+decreaseNumber.(string)+" WHERE id = ? AND WHERE "+check, id)
}
