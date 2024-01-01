package _const

/**
* this is supported language
* please make sure you add add what language
* you need here
 */
func SupportedLang() map[string]string {
	return map[string]string{
		"ar": "العربية",
		"en": "English",
	}
}
func SupportedLangSlice() []string {
	return []string{
		"ar",
		"en",
	}
}
