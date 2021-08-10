package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func telValidation(fl validator.FieldLevel) bool {
	return v.IsTel(fl.Field().String())
}

// RegisterTelValidation register validations with translations
func RegisterTelValidation(val v.Validator) {
	val.AddValidation("tel", identifierValidation)
	val.AddTranslation("en", "tel", "Must be a valid  tel")
	val.AddTranslation("fa", "tel", "شناسه وارد شده معتبر نیست")
}
