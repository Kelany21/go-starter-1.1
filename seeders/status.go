package seeders

import (
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func StatusSeeder() {
	for moduleName, module := range _const.Services.Modules {
		loop := module["statuses"].([]interface{})
		for _, name := range loop {
			newStatus(models.Status{
				Noun:       name.(string),
				Verb:       name.(string),
				Count:      0,
				Slug:       name.(string) + "_" + moduleName,
				ModuleName: moduleName,
			})
		}
	}
}

/**
* fake data and create data base
 */
func newStatus(data models.Status) {
	_const.Services.DB.Create(&data)
}
