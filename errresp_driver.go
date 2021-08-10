package validator

import (
	"encoding/json"
	"fmt"
)

// errorResponseDriver error response driver
type errorResponseDriver struct {
	errors map[string]map[string]string
}

// init structure
func (er *errorResponseDriver) init() {
	if er.errors == nil {
		er.errors = make(map[string]map[string]string)
	}
}

// AddError add new error
func (er *errorResponseDriver) AddError(field, tag, message string) {
	if _, exists := er.errors[field]; !exists {
		er.errors[field] = make(map[string]string)
	}

	er.errors[field][tag] = message
}

// HasError check if response has error
func (er *errorResponseDriver) HasError() bool {
	return len(er.errors) > 0
}

// Failed check if field has error
// @example:
// resp.Failed("firstname")
func (er *errorResponseDriver) Failed(field string) bool {
	_, exists := er.errors[field]
	return exists
}

// FailedOn check if field has special error
// @example:
// resp.FailedOn("firstname", "required")
func (er *errorResponseDriver) FailedOn(field, err string) bool {
	if errs, exists := er.errors[field]; exists {
		if _, hasErr := errs[err]; hasErr {
			return true
		}
	}
	return false
}

// Errors get errors list
func (er *errorResponseDriver) Errors() map[string]map[string]string {
	return er.errors
}

// String convert to string
func (er *errorResponseDriver) String() string {
	return fmt.Sprintf("%v", er.errors)
}

// MarshalJSON convert to json
func (er *errorResponseDriver) MarshalJSON() ([]byte, error) {
	return json.Marshal(er.errors)
}
