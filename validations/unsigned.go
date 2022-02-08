package validations

import (
	"fmt"

	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func unsignedValidation(fl validator.FieldLevel) bool {
	return v.IsUnsigned(fmt.Sprint(fl.Field().Interface()))
}

// RegisterUnsignedValidation register validations with translations
func RegisterUnsignedValidation(val v.Validator) {
	val.AddValidation("unsigned", unsignedValidation)
	val.AddTranslation("en", "unsigned", "Must be a unsigned number")
	val.AddTranslation("fa", "unsigned", "باید یک عدد صحیح مثبت باشد")
}
