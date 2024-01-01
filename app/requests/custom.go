package requests

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"reflect"
	"regexp"
	"starter-golang-new/app/models"
	_const "starter-golang-new/const"
	"strconv"
	"strings"
)

func Init() {
	/**
	* this role check if slice of strings is not empty
	 */
	govalidator.AddCustomRule("strings_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		values := value.([]interface{})
		if len(values) == 0 {
			if message != "" {
				return err
			}
			for i := 0; i < len(values); i++ {
				if reflect.TypeOf(values[i]) != reflect.TypeOf(reflect.String) {
					return err
				}
			}
		}
		return nil
	})

	/**
	* this role check if slice of int is not empty
	 */
	govalidator.AddCustomRule("int_slice", func(field string, rule string, message string, value interface{}) error {
		if value == nil {
			return fmt.Errorf("The field is required")
		}
		if len(value.([]interface{})) <= 0 {
			if message == "" {
				return errors.New(message)
			}
			return fmt.Errorf("The field is required")
		}
		return nil
	})

	///**
	//* this role check if slice of int is not empty
	// */
	//govalidator.AddCustomRule("permission_slice", func(field string, rule string, message string, value interface{}) error {
	//	err := fmt.Errorf("The %s field is required", field)
	//	if message != "" {
	//		err = errors.New(message)
	//	}
	//	if value == nil {
	//		return err
	//	}
	//	permissions := value.([]models.PermissionForm)
	//	if len(permissions) == 0 {
	//		return err
	//	}
	//	return nil
	//})

	govalidator.AddCustomRule("int_array_slice", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var i []int
		if reflect.TypeOf(value) == reflect.TypeOf(i) {
			i = value.([]int)
			if len(i) == 0 {
				return err
			}
		}
		return nil
	})

	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is unique", field)
		table := strings.Split(rule, ":")
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var count int
		_const.Services.DB.Table(table[1]).Where(field+" = ?", value.(string)).Count(&count)
		if count != 0 {
			return err
		}
		return nil
	})

	govalidator.AddCustomRule("unique_update", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is unique", field)
		rules := strings.Split(rule, ":")
		table := strings.Split(rules[1], ",")
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var count int
		_const.Services.DB.Table(table[0]).Where(field+" = ?", value.(string)).Where("id != ?", table[1]).Count(&count)
		if count != 0 {
			return err
		}
		return nil
	})

	govalidator.AddCustomRule("lang", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("%s must be multi language value", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		var i []interface{}
		if reflect.TypeOf(value) == reflect.TypeOf(i) {
			i = value.([]interface{})
			if len(i) == 0 {
				return err
			}
			length := len(_const.SupportedLang())
			if len(i) != length {
				err := fmt.Errorf("%s must have "+strconv.Itoa(length)+" language", field)
				return err
			}
		}
		return nil
	})

	govalidator.AddCustomRule("lang_min", func(field string, rule string, message string, value interface{}) error {
		ruleParam := strings.Split(rule, ":")
		err := fmt.Errorf("The %s length must be more than %s ", field, ruleParam[1])
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		min, _ := strconv.Atoi(ruleParam[1])
		var values []interface{}
		if reflect.TypeOf(value) == reflect.TypeOf(values) {
			values = value.([]interface{})
			if len(values) == 0 {
				return err
			}
			for i := 0; i < len(values); i++ {
				if len(strings.TrimSpace(values[i].(string))) < min {
					return err
				}
			}
		}
		return nil
	})
	govalidator.AddCustomRule("lang_max", func(field string, rule string, message string, value interface{}) error {
		ruleParam := strings.Split(rule, ":")
		err := fmt.Errorf("The %s length must be less than %s ", field, ruleParam[1])
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		max, _ := strconv.Atoi(ruleParam[1])
		var values []interface{}
		if reflect.TypeOf(value) == reflect.TypeOf(values) {
			values = value.([]interface{})
			if len(values) == 0 {
				return err
			}
			for i := 0; i < len(values); i++ {
				if len(strings.TrimSpace(values[i].(string))) > max {
					return err
				}
			}
		}
		return nil
	})
	govalidator.AddCustomRule("answers", func(field string, rule string, message string, value interface{}) error {
		ruleParam := strings.Split(rule, ":")
		err := fmt.Errorf("The %s length must be less than %s ", field, ruleParam[1])
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		langLength := strings.Split(ruleParam[1], ",")
		min, _ := strconv.Atoi(langLength[0])
		max, _ := strconv.Atoi(langLength[1])
		var values []interface{}
		if reflect.TypeOf(value) == reflect.TypeOf(values) {
			values = value.([]interface{})
			if len(values) == 0 {
				return err
			}
			for i := 0; i < len(values); i++ {
				texts := values[i].(map[string]interface{})
				for _, v := range texts {
					var textInterfaces []interface{}
					if reflect.TypeOf(v) != reflect.TypeOf(textInterfaces) {
						return err
					}
					textValues := v.([]interface{})
					if len(textValues) != _const.Services.SupportedLanguageCount {
						return err
					}
					for i := 0; i < len(textValues); i++ {
						length := len(strings.TrimSpace(textValues[i].(string)))
						if length > max || length < min {
							return err
						}
					}
				}

			}
		}
		return nil
	})
	govalidator.AddCustomRule("is_permission_group", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s not found ", field)
		if message != "" {
			err = errors.New(message)
		}
		if value == nil {
			return err
		}
		count := 0
		_const.Services.DB.Model(&models.PermissionGroup{}).Where("id = ?", value).Count(&count)
		if count == 0 {
			return err
		}
		return nil
	})
	govalidator.AddCustomRule("alpha_space_dash_underscore", func(field string, rule string, message string, value interface{}) error {
		err := fmt.Errorf("The %s field is required", field)
		if message != "" {
			err = errors.New(message)
		}
		str := value.(string)
		var AlphaSpace = "^[-a-zA-Z_ ]+$"
		regexAlphaSpace   := regexp.MustCompile(AlphaSpace)
		if !regexAlphaSpace.MatchString(str) {
			return err
		}
		return nil
	})
}
