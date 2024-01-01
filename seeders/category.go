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
func CategorySeeder() {
	for i := 0; i < 10; i++ {
		newCategory()
	}
}

/**
* fake data and create data base
 */
func newCategory() {
	data := models.Category{
		Status: _const.ACTIVE,
	}
	_const.Services.DB.Create(&data)
	for key, _ := range _const.SupportedLang() {
		for key2, keyName := range models.CategoryTranslateFillable() {
			value := faker.Internet().UserName() + key
			if keyName == "image" {
				value = "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885_960_720.jpg"
			}
			categoryI18ns := models.CategoryI18ns{
				FiledName:   key2,
				Value:       value,
				Language:    key,
				ReferenceId: uint64(data.ID),
			}
			_const.Services.DB.Create(&categoryI18ns)
		}

	}

}
