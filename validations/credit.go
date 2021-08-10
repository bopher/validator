package validations

import (
	v "github.com/bopher/validator"
	"github.com/go-playground/validator/v10"
)

func creditCardValidation(fl validator.FieldLevel) bool {
	return v.IsCreditCardNumber(fl.Field().String())
}

// RegisterCreditCardValidation register credit card validator and it translations
func RegisterCreditCardValidation(val v.Validator) {
	val.AddValidation("creditcard", creditCardValidation)
	val.AddTranslation("en", "creditcard", "Must be a valid credit card number")
	val.AddTranslation("fa", "creditcard", "شماره کارت وارد شده معتبر نیست")
}
