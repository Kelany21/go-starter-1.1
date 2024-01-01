package models

import (
	"os"
	_const "starter-golang-new/const"
	"strconv"
)

func MigrateAllTable() {
	deleteTables, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
	if deleteTables {
		DbTruncate()
	}
	AnswerMigrate()
	CategoryMigrate()
	FaqMigrate()
	PageMigrate()
	PageImageMigrate()
	SettingMigrate()
	UserMigrate()
	StatusMigrate()
	PermissionGroupMigrate()
	RoleMigrate()
}

/**
* drop all tables
 */

type query struct {
	Query string
}

func DbTruncate() {
	var query []query
	_const.Services.DB.Table("information_schema.tables").Select("concat('DROP TABLE IF EXISTS `', table_name, '`;') as query").Where("table_schema = ? ", "starter").Find(&query)
	for _, q := range query {
		_const.Services.DB.Exec(q.Query)
	}
}
