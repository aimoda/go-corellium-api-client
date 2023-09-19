/*
Corellium API

REST API to manage your virtual devices.

API version: 5.4.1-18421
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the ValidationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationError{}

// ValidationError 
type ValidationError struct {
	// Error text
	Error string `json:"error"`
	// Error ID
	ErrorID string `json:"errorID"`
	// Field associated with error
	Field NullableString `json:"field,omitempty"`
}

// NewValidationError instantiates a new ValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationError(error_ string, errorID string) *ValidationError {
	this := ValidationError{}
	this.Error = error_
	this.ErrorID = errorID
	return &this
}

// NewValidationErrorWithDefaults instantiates a new ValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorWithDefaults() *ValidationError {
	this := ValidationError{}
	return &this
}

// GetError returns the Error field value
func (o *ValidationError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ValidationError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ValidationError) SetError(v string) {
	o.Error = v
}

// GetErrorID returns the ErrorID field value
func (o *ValidationError) GetErrorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorID
}

// GetErrorIDOk returns a tuple with the ErrorID field value
// and a boolean to check if the value has been set.
func (o *ValidationError) GetErrorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorID, true
}

// SetErrorID sets field value
func (o *ValidationError) SetErrorID(v string) {
	o.ErrorID = v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidationError) GetField() string {
	if o == nil || IsNil(o.Field.Get()) {
		var ret string
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationError) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *ValidationError) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableString and assigns it to the Field field.
func (o *ValidationError) SetField(v string) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *ValidationError) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *ValidationError) UnsetField() {
	o.Field.Unset()
}

func (o ValidationError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["errorID"] = o.ErrorID
	if o.Field.IsSet() {
		toSerialize["field"] = o.Field.Get()
	}
	return toSerialize, nil
}

type NullableValidationError struct {
	value *ValidationError
	isSet bool
}

func (v NullableValidationError) Get() *ValidationError {
	return v.value
}

func (v *NullableValidationError) Set(val *ValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationError(val *ValidationError) *NullableValidationError {
	return &NullableValidationError{value: val, isSet: true}
}

func (v NullableValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


