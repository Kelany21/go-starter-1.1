package seeders

import (
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func PageSeeder() {
	ar := pagesAr()
	for i, page := range pagesEn() {
		data := models.Page{
			Status: _const.ACTIVE,
			Slug:   page,
		}
		_const.Services.DB.Create(&data)
		dataTranslate := models.PageI18ns{
			ReferenceId: uint64(data.ID),
			FiledName:   "name",
			Language:    "en",
			Value:       page,
		}
		dataTranslateAr := models.PageI18ns{
			ReferenceId: uint64(data.ID),
			FiledName:   "name",
			Language:    "en",
			Value:       ar[i],
		}
		_const.Services.DB.Create(&dataTranslate)
		_const.Services.DB.Create(&dataTranslateAr)
	}
}

/***
* list of pages
 */
func pagesEn() []string {
	return []string{
		"home",
		"about",
		"contact",
		"terms",
		"police",
	}
}

/***
* list of pages
 */
func pagesAr() []string {
	return []string{
		"الرءسيه",
		"من نحن ",
		"تواصل معانا ",
		"شروط",
		"الخصوصيات",
	}
}
