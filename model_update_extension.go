/*
Corellium API

REST API to manage your virtual devices.

API version: 5.5.0-18750
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the UpdateExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExtension{}

// UpdateExtension 
type UpdateExtension struct {
	// If true, instances requiring this extension can be created or started
	Enabled NullableBool `json:"enabled,omitempty"`
}

// NewUpdateExtension instantiates a new UpdateExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExtension() *UpdateExtension {
	this := UpdateExtension{}
	return &this
}

// NewUpdateExtensionWithDefaults instantiates a new UpdateExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExtensionWithDefaults() *UpdateExtension {
	this := UpdateExtension{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExtension) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExtension) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateExtension) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateExtension) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateExtension) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateExtension) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o UpdateExtension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

type NullableUpdateExtension struct {
	value *UpdateExtension
	isSet bool
}

func (v NullableUpdateExtension) Get() *UpdateExtension {
	return v.value
}

func (v *NullableUpdateExtension) Set(val *UpdateExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExtension(val *UpdateExtension) *NullableUpdateExtension {
	return &NullableUpdateExtension{value: val, isSet: true}
}

func (v NullableUpdateExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


