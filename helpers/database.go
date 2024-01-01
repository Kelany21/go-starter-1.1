package helpers

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	_const "starter-golang-new/const"
	"strconv"
	"strings"
	"time"
)

/***
* truncate tables
 */

type TranslationScan struct {
	gorm.Model
	ReferenceId uint64 `gorm:"type:int(20);" json:"reference_id"`
	FiledName   string `gorm:"type:varchar(20);" json:"filed_name"`
	Language    string `gorm:"type:varchar(5);" json:"language"`
	Value       string `gorm:"type:text;" json:"value"`
}

func DbTruncate(tableName ...string) {
	for _, table := range tableName {
		_const.Services.DB.Exec("TRUNCATE " + table)
	}
}

/**
* this function get struct and return with only
* Available column that allow to updated depend on FillAbleColumn function
* this for security
* map struct to update
 */
func UpdateOnlyAllowColumns(structNeedToMap interface{}, fillAble []string) interface{} {
	row := structs.Map(structNeedToMap)
	var data = make(map[string]interface{})
	for _, value := range fillAble {
		if row[strings.Title(value)] != "" {
			data[value] = row[strcase.ToCamel(value)]
		}
	}
	return data
}

/**
* add preload dynamic
* this will allow to add more than one preload
 */
func PreloadD(db *gorm.DB, preload []string, preloadConditions map[string][]string) *gorm.DB {
	if len(preload) > 0 {
		for _, p := range preload {
			if val, ok := preloadConditions[p]; ok {
				s := make([]interface{}, len(val))
				for i, v := range val {
					s[i] = v
				}
				db = db.Preload(p, s...)
			} else {
				db = db.Preload(p)
			}
		}
		return db
	}
	return db
}

/**
* get translate rows depend on table name
 */
func RowsTranslations(table string, id interface{}) []TranslationScan {
	results := []TranslationScan{}
	_const.Services.DB.Table(table).Where("reference_id = ?", id).Order("filed_name desc").Scan(&results)
	return results
}

/**
* get translate rows depend on table name
 */
func DeleteRowsTranslations(table string, id ...interface{}) {
	results := []TranslationScan{}
	_const.Services.DB.Table(table).Where("reference_id in (?)", id).Unscoped().Delete(&results)
}

/**
* build insert query
 */
func InsertQuery(table string, data interface{}) bool {
	d := data.(map[string]interface{})
	if len(d) > 0 {
		userInsert := `INSERT INTO ` + table + ` SET `
		for name, value := range d {
			v := ""
			if name == "reference_id" {
				v := strconv.FormatUint(uint64(value.(uint)), 10)
				userInsert += name + `=` + v + `, `
			} else if strings.Contains(name, "id") {
				v := strconv.Itoa(int(value.(float64)))
				userInsert += name + `=` + v + `, `
			} else {
				v = fmt.Sprintf("%s", value)
				userInsert += name + `="` + v + `", `
			}

		}
		userInsert += ` created_at ="` + time.Now().Format("2006-01-02 15:04:05") + `",`
		userInsert += ` updated_at ="` + time.Now().Format("2006-01-02 15:04:05") + `"`
		_, _ = _const.Services.SQL.Exec(userInsert)
		return true
	}
	return false
}
