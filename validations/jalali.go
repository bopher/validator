package validations

import (
	"github.com/bopher/jalali"
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func jalaliValidation(fl validator.FieldLevel) bool {
	if d := jalali.Parse(fl.Field().String()); d != nil {
		return true
	}
	return false
}

// RegisterJalaliValidation register validations with translations
func RegisterJalaliValidation(val v.Validator) {
	val.AddValidation("jalali", jalaliValidation)
	val.AddTranslation("en", "jalali", "Must be a valid jalali date")
	val.AddTranslation("fa", "jalali", "تاریخ وارد شده معتبر نیست")
}
