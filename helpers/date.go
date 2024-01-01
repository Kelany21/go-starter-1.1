package helpers

import (
	"regexp"
	_const "starter-golang-new/const"
	"strconv"
	"time"
)

func StringDateReformat(date time.Time) string {
	year, month, day := date.Date()
	lang := GetCurrentLang()
	var m, stringDate string
	if lang == "ar" {
		m = monthAr()[int(month)]
		stringDate = ToPersianDigits(strconv.Itoa(day)) + " " + m + " " + ToPersianDigits(strconv.Itoa(year))
	} else {
		m = monthEn()[int(month)]
		stringDate = strconv.Itoa(day) + " " + m + " " + strconv.Itoa(year)
	}
	return stringDate
}

func StringTimeReformat(date time.Time) string {
	date = date.In(_const.Services.TimeLocation)
	lang := GetCurrentLang()
	if lang == "ar" {
		return ToPersianDigits(PmAm(date.Format(time.Kitchen)))
	}
	return date.Format(time.Kitchen)
}

func PmAm(text string) string {
	var checker = map[string]string{
		"M": "",
		"A": " صباحا",
		"P": " مساء",
	}
	re := regexp.MustCompile("[A-Z]+")
	out := re.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		out := ""
		ss := string(s)
		for _, ch := range ss {
			o := checker[string(ch)]
			out = out + o
		}
		return []byte(out)
	})
	return string(out)
}

/**
* convert english number to arabic number
 */
func ToPersianDigits(text string) string {
	var checker = map[string]string{
		"0": "۰",
		"1": "١",
		"2": "٢",
		"3": "٣",
		"4": "٤",
		"5": "٥",
		"6": "٦",
		"7": "٧",
		"8": "٨",
		"9": "٩",
	}
	re := regexp.MustCompile("[0-9]+")
	out := re.ReplaceAllFunc([]byte(text), func(s []byte) []byte {
		out := ""
		ss := string(s)
		for _, ch := range ss {
			o := checker[string(ch)]
			out = out + o
		}
		return []byte(out)
	})
	return string(out)
}

/**
* map months
 */
func monthEn() map[int]string {
	var m = make(map[int]string)
	m[1] = "January"
	m[2] = "February"
	m[3] = "March"
	m[4] = "April"
	m[5] = "May"
	m[6] = "June"
	m[7] = "July"
	m[8] = "August"
	m[9] = "September"
	m[10] = "October"
	m[11] = "November"
	m[12] = "December"

	return m
}

/**
* map arabic months
 */
func monthAr() map[int]string {
	var m = make(map[int]string)
	m[1] = "يناير"
	m[2] = "فبراير"
	m[3] = "مارس"
	m[4] = "ابريل"
	m[5] = "مايو"
	m[6] = "ينيو"
	m[7] = "يليو"
	m[8] = "اغسطس"
	m[9] = "ستمبر"
	m[10] = "اكتوبر"
	m[11] = "نوفمبر"
	m[12] = "ديسمبر"

	return m
}
