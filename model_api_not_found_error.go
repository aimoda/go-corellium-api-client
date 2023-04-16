/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the ApiNotFoundError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNotFoundError{}

// ApiNotFoundError 
type ApiNotFoundError struct {
	// Error text
	Error string `json:"error"`
	// Error ID
	ErrorID string `json:"errorID"`
	// Name of the object requested
	Name NullableString `json:"name,omitempty"`
	// params provided by user
	Params map[string]interface{} `json:"params,omitempty"`
}

// NewApiNotFoundError instantiates a new ApiNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotFoundError(error_ string, errorID string) *ApiNotFoundError {
	this := ApiNotFoundError{}
	this.Error = error_
	this.ErrorID = errorID
	return &this
}

// NewApiNotFoundErrorWithDefaults instantiates a new ApiNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotFoundErrorWithDefaults() *ApiNotFoundError {
	this := ApiNotFoundError{}
	return &this
}

// GetError returns the Error field value
func (o *ApiNotFoundError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ApiNotFoundError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ApiNotFoundError) SetError(v string) {
	o.Error = v
}

// GetErrorID returns the ErrorID field value
func (o *ApiNotFoundError) GetErrorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorID
}

// GetErrorIDOk returns a tuple with the ErrorID field value
// and a boolean to check if the value has been set.
func (o *ApiNotFoundError) GetErrorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorID, true
}

// SetErrorID sets field value
func (o *ApiNotFoundError) SetErrorID(v string) {
	o.ErrorID = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiNotFoundError) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiNotFoundError) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiNotFoundError) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiNotFoundError) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiNotFoundError) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiNotFoundError) UnsetName() {
	o.Name.Unset()
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiNotFoundError) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiNotFoundError) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ApiNotFoundError) HasParams() bool {
	if o != nil && IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *ApiNotFoundError) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o ApiNotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNotFoundError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["errorID"] = o.ErrorID
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableApiNotFoundError struct {
	value *ApiNotFoundError
	isSet bool
}

func (v NullableApiNotFoundError) Get() *ApiNotFoundError {
	return v.value
}

func (v *NullableApiNotFoundError) Set(val *ApiNotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotFoundError(val *ApiNotFoundError) *NullableApiNotFoundError {
	return &NullableApiNotFoundError{value: val, isSet: true}
}

func (v NullableApiNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

