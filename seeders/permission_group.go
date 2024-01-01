package seeders

import (
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"syreclabs.com/go/faker"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func PermissionGroupSeeder() {
	for i := 0; i < 10; i++ {
		newPermissionGroups()
	}
}


/**
* fake data and create data base
 */
func newPermissionGroups() {
	data := models.PermissionGroup{
		Status: _const.ACTIVE,
	}
	_const.Services.DB.Create(&data)
	for key, _ := range _const.SupportedLang() {
		for key2, _ := range models.PermissionGroupTranslateFillable() {
			value := faker.Internet().UserName() + key
			categoryI18ns := models.PermissionGroupI18ns{
				FiledName:   key2,
				Value:       value,
				Language:    key,
				ReferenceId: uint64(data.ID),
			}
			_const.Services.DB.Create(&categoryI18ns)
		}

	}

}

