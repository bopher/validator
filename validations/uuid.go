package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func uuidValidation(fl validator.FieldLevel) bool {
	return v.IsUUID(fl.Field().String())
}

// RegisterUUIDValidation register validations with translations
func RegisterUUIDValidation(val v.Validator) {
	val.AddValidation("uuid", uuidValidation)
	val.AddTranslation("en", "uuid", "Must be a valid uuid")
	val.AddTranslation("fa", "uuid", "مقدار ورودی باید یه شناسه UUID معتبر باشد")
}
