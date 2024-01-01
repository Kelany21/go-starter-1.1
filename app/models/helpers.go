package models

import (
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* add to active and all
 */
func IncreaseOnCreate(moduleName string) {
	IncreaseRow(_const.ACTIVE, moduleName)
	helpers.Increase("statuses", "count", nil, "1", `slug = "`+_const.ALL+`_`+moduleName+`"`)
}

/**
* remove from all and status
 */

func DecreaseOnDelete(status string, moduleName string) {
	DecreaseRow(status, moduleName)
	helpers.Decrease("statuses", "count", nil, "1", `slug = "`+_const.ALL+`_user"`)
}

func DecreaseRow(status string, moduleName string) {
	helpers.Decrease("statuses", "count", nil, "1", `module_name =  "`+moduleName+`"`, `verb = "`+status+`"`)
}

func IncreaseRow(status string, moduleName string) {
	helpers.Increase("statuses", "count", nil, "1", `module_name =  "`+moduleName+`"`, `verb = "`+status+`"`)
}

func GetActionByModule(moduleName string) []Status {
	var actions []Status
	_const.Services.DB.Where("module_name = ? ", moduleName).Find(&actions)
	return actions
}

