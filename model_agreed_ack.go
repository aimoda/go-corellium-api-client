/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"time"
)

// checks if the AgreedAck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgreedAck{}

// AgreedAck 
type AgreedAck struct {
	// Date agreed
	AgreedToTerms NullableTime `json:"agreedToTerms,omitempty"`
}

// NewAgreedAck instantiates a new AgreedAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreedAck() *AgreedAck {
	this := AgreedAck{}
	return &this
}

// NewAgreedAckWithDefaults instantiates a new AgreedAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreedAckWithDefaults() *AgreedAck {
	this := AgreedAck{}
	return &this
}

// GetAgreedToTerms returns the AgreedToTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgreedAck) GetAgreedToTerms() time.Time {
	if o == nil || IsNil(o.AgreedToTerms.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AgreedToTerms.Get()
}

// GetAgreedToTermsOk returns a tuple with the AgreedToTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgreedAck) GetAgreedToTermsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgreedToTerms.Get(), o.AgreedToTerms.IsSet()
}

// HasAgreedToTerms returns a boolean if a field has been set.
func (o *AgreedAck) HasAgreedToTerms() bool {
	if o != nil && o.AgreedToTerms.IsSet() {
		return true
	}

	return false
}

// SetAgreedToTerms gets a reference to the given NullableTime and assigns it to the AgreedToTerms field.
func (o *AgreedAck) SetAgreedToTerms(v time.Time) {
	o.AgreedToTerms.Set(&v)
}
// SetAgreedToTermsNil sets the value for AgreedToTerms to be an explicit nil
func (o *AgreedAck) SetAgreedToTermsNil() {
	o.AgreedToTerms.Set(nil)
}

// UnsetAgreedToTerms ensures that no value is present for AgreedToTerms, not even an explicit nil
func (o *AgreedAck) UnsetAgreedToTerms() {
	o.AgreedToTerms.Unset()
}

func (o AgreedAck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgreedAck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgreedToTerms.IsSet() {
		toSerialize["agreedToTerms"] = o.AgreedToTerms.Get()
	}
	return toSerialize, nil
}

type NullableAgreedAck struct {
	value *AgreedAck
	isSet bool
}

func (v NullableAgreedAck) Get() *AgreedAck {
	return v.value
}

func (v *NullableAgreedAck) Set(val *AgreedAck) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreedAck) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreedAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreedAck(val *AgreedAck) *NullableAgreedAck {
	return &NullableAgreedAck{value: val, isSet: true}
}

func (v NullableAgreedAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreedAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

