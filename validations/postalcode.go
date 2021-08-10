package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func postalCodeValidation(fl validator.FieldLevel) bool {
	return v.IsPostalcode(fl.Field().String())
}

// RegisterPostalCodeValidation register validations with translations
func RegisterPostalCodeValidation(val v.Validator) {
	val.AddValidation("postalcode", identifierValidation)
	val.AddTranslation("en", "postalcode", "Must be a valid postal-code")
	val.AddTranslation("fa", "postalcode", "کد پستی وارد شده معتبر نیست")
}
