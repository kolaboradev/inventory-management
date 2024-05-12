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

func IsValidUrl(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
		return urlRegex.MatchString(value)
	}
	return false
}

func IsValidURL(url string) bool {
	var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
	return urlRegex.MatchString(url)
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
	case "url_valid":
		return "image not url"
	default:
		return "Tag not implement custom message error"
	}
}
