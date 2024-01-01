package helpers

import (
	_const "starter-golang-new/const"
)

/**
* get module configuration
 */
func Module(name string) map[string]interface{} {
	return _const.Services.Modules[name]
}

/**
* get module Name
 */
func ModuleName(name string) string {
	return _const.Services.Modules[name]["name"].(string)
}

/**
* get module table name
 */
func ModuleTable(name string) string {
	return Module(name)["table"].(string)
}

/**
* get module Translate table name
 */
func ModuleTransTable(name string) string {
	return Module(name)["trans_table"].(string)
}

/**
* get module table name
 */
func ModuleRoute(name string) string {
	return Module(name)["route"].(string)
}

/**
* get module table name
 */
func ModelModel(name string) string {
	return Module(name)["model"].(string)
}

/**
* get module permissions
 */
func ModulePermissions(name string) interface{} {
	return Module(name)["permissions"]
}

/**
* get module statuses
 */
func ModuleStatuses(name string) interface{} {
	return Module(name)["statuses"]
}
