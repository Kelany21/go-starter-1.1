package transformers

import (
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

func TransformTranslation(translations []helpers.TranslationScan, m map[string]interface{}) map[string]interface{} {
	u := []interface{}{}
	for _, translation := range translations {
		if _, ok := m[translation.FiledName]; ok {
			u = append(u, translation.Value)
			m[translation.FiledName] = u
		} else {
			u = []interface{}{}
			u = append(u, translation.Value)
			m[translation.FiledName] = u
		}
	}

	return m
}

func GetTranslations(ids []int, table string) map[uint64][]helpers.TranslationScan {
	translationMap := make(map[uint64][]helpers.TranslationScan)
	var translations []helpers.TranslationScan
	_const.Services.DB.Table(table).Where("reference_id in (?) ", ids).
		Order("find_in_set(reference_id,'" + helpers.ArrayToString(ids) + "')").
		Order("filed_name desc").Scan(&translations)
	for i := 0; i < len(translations); i++ {
		if _, ok := translationMap[translations[i].ReferenceId]; ok {
			translationMap[translations[i].ReferenceId] = append(translationMap[translations[i].ReferenceId], translations[i])
		} else {
			translationMap[translations[i].ReferenceId] = []helpers.TranslationScan{translations[i]}
		}
	}
	return translationMap
}

func TransformTranslationObject(translations []helpers.TranslationScan, m map[string]interface{}) map[string]interface{} {
	lang := helpers.GetCurrentLang()
	for _, translation := range translations {
		if translation.Language == lang {
			m[translation.FiledName] = translation.Value
		}
	}

	return m
}
