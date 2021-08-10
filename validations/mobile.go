package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func mobileValidation(fl validator.FieldLevel) bool {
	return v.IsMobile(fl.Field().String())
}

// RegisterMobileValidation register validations with translations
func RegisterMobileValidation(val v.Validator) {
	val.AddValidation("mobile", identifierValidation)
	val.AddTranslation("en", "mobile", "Must be a valid mobile")
	val.AddTranslation("fa", "mobile", "شناسه وارد شده معتبر نیست")
}
