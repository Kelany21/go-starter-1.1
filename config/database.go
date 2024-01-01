package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"os"
	_const "starter-golang-new/const"
	"strconv"
)

/**
* connect with data base with env file params
* just edit all data in .env file
 */
func ConnectToDatabase() {
	db := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	if _const.Services.DB == nil {
		_const.Services.DB, _const.Services.DBERR = gorm.Open("mysql", db)
	}
	if _const.Services.SQL == nil {
		_const.Services.SQL, _const.Services.DBERR = sql.Open("mysql", db)
	}
	if _const.Services.SQLX == nil {
		_const.Services.SQLX, _const.Services.DBERR = sqlx.Open("mysql", db)
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if debug {
		_const.Services.DB.LogMode(debug)
	}
}
