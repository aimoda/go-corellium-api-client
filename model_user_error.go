/*
Corellium API

REST API to manage your virtual devices.

API version: 5.2.0-17376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the UserError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserError{}

// UserError 
type UserError struct {
	// Error text
	Error string `json:"error"`
	// Error ID
	ErrorID string `json:"errorID"`
	// Field associated with error
	Field NullableString `json:"field,omitempty"`
}

// NewUserError instantiates a new UserError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserError(error_ string, errorID string) *UserError {
	this := UserError{}
	this.Error = error_
	this.ErrorID = errorID
	return &this
}

// NewUserErrorWithDefaults instantiates a new UserError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserErrorWithDefaults() *UserError {
	this := UserError{}
	return &this
}

// GetError returns the Error field value
func (o *UserError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *UserError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *UserError) SetError(v string) {
	o.Error = v
}

// GetErrorID returns the ErrorID field value
func (o *UserError) GetErrorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorID
}

// GetErrorIDOk returns a tuple with the ErrorID field value
// and a boolean to check if the value has been set.
func (o *UserError) GetErrorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorID, true
}

// SetErrorID sets field value
func (o *UserError) SetErrorID(v string) {
	o.ErrorID = v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserError) GetField() string {
	if o == nil || IsNil(o.Field.Get()) {
		var ret string
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserError) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *UserError) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableString and assigns it to the Field field.
func (o *UserError) SetField(v string) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *UserError) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *UserError) UnsetField() {
	o.Field.Unset()
}

func (o UserError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["errorID"] = o.ErrorID
	if o.Field.IsSet() {
		toSerialize["field"] = o.Field.Get()
	}
	return toSerialize, nil
}

type NullableUserError struct {
	value *UserError
	isSet bool
}

func (v NullableUserError) Get() *UserError {
	return v.value
}

func (v *NullableUserError) Set(val *UserError) {
	v.value = val
	v.isSet = true
}

func (v NullableUserError) IsSet() bool {
	return v.isSet
}

func (v *NullableUserError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserError(val *UserError) *NullableUserError {
	return &NullableUserError{value: val, isSet: true}
}

func (v NullableUserError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


