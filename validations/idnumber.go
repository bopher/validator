package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func idNumberValidation(fl validator.FieldLevel) bool {
	return v.IsIDNumber(fl.Field().String())
}

// RegisterIDNumberValidation register validations with translations
func RegisterIDNumberValidation(val v.Validator) {
	val.AddValidation("idnumber", idNumberValidation)
	val.AddTranslation("en", "idnumber", "Must be a valid id number")
	val.AddTranslation("fa", "idnumber", "شماره شناسنامه وارد شده معتبر نیست")
}
