package validations

import "github.com/bopher/validator"

// RegisterExtraValidations register extra validations to validator
func RegisterExtraValidations(v validator.Validator) {
	RegisterCreditCardValidation(v)
	RegisterIdentifierValidation(v)
	RegisterIDNumberValidation(v)
	RegisterIPPortValidation(v)
	RegisterJalaliValidation(v)
	RegisterMobileValidation(v)
	RegisterNationalCodeValidation(v)
	RegisterPostalCodeValidation(v)
	RegisterTelValidation(v)
	RegisterUnsignedValidation(v)
	RegisterUsernameValidation(v)
	RegisterUUIDValidation(v)
}
