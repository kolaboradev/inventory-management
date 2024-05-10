package helper

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func MustPhoneNumber(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		phoneNumberRegex := `^\+[0-9]{1,4}-?[0-9]{1,15}$`
		matched, _ := regexp.MatchString(phoneNumberRegex, value)
		if matched {
			return true
		}
	}
	return false
}

func MatchCategoryProduct(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		switch value {
		case "Clothing":
			return true
		case "Accessories":
			return true
		case "Footwear":
			return true
		case "Beverages":
			return true
		default:
			return false
		}
	}
	return false
}

func CustomMessageValidation(valErr validator.FieldError) string {
	switch valErr.ActualTag() {
	case "phone_number":
		return "Phone Number Invalid"
	case "required":
		return fmt.Sprintf("%s required", valErr.Field())
	case "min":
		minParam := valErr.Param()
		return fmt.Sprintf("%s cannot be less than %s", valErr.Field(), minParam)
	case "max":
		maxParam := valErr.Param()
		return fmt.Sprintf("%s can't be more than %s", valErr.Field(), maxParam)
	case "category":
		return "The product does not match the existing type"
	default:
		return "Tag not implement custom message error"
	}
}
