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

// checks if the InstanceStartOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceStartOptions{}

// InstanceStartOptions 
type InstanceStartOptions struct {
	// Start device paused
	Paused NullableBool `json:"paused,omitempty"`
}

// NewInstanceStartOptions instantiates a new InstanceStartOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceStartOptions() *InstanceStartOptions {
	this := InstanceStartOptions{}
	return &this
}

// NewInstanceStartOptionsWithDefaults instantiates a new InstanceStartOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceStartOptionsWithDefaults() *InstanceStartOptions {
	this := InstanceStartOptions{}
	return &this
}

// GetPaused returns the Paused field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceStartOptions) GetPaused() bool {
	if o == nil || IsNil(o.Paused.Get()) {
		var ret bool
		return ret
	}
	return *o.Paused.Get()
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceStartOptions) GetPausedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paused.Get(), o.Paused.IsSet()
}

// HasPaused returns a boolean if a field has been set.
func (o *InstanceStartOptions) HasPaused() bool {
	if o != nil && o.Paused.IsSet() {
		return true
	}

	return false
}

// SetPaused gets a reference to the given NullableBool and assigns it to the Paused field.
func (o *InstanceStartOptions) SetPaused(v bool) {
	o.Paused.Set(&v)
}
// SetPausedNil sets the value for Paused to be an explicit nil
func (o *InstanceStartOptions) SetPausedNil() {
	o.Paused.Set(nil)
}

// UnsetPaused ensures that no value is present for Paused, not even an explicit nil
func (o *InstanceStartOptions) UnsetPaused() {
	o.Paused.Unset()
}

func (o InstanceStartOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceStartOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Paused.IsSet() {
		toSerialize["paused"] = o.Paused.Get()
	}
	return toSerialize, nil
}

type NullableInstanceStartOptions struct {
	value *InstanceStartOptions
	isSet bool
}

func (v NullableInstanceStartOptions) Get() *InstanceStartOptions {
	return v.value
}

func (v *NullableInstanceStartOptions) Set(val *InstanceStartOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceStartOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceStartOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceStartOptions(val *InstanceStartOptions) *NullableInstanceStartOptions {
	return &NullableInstanceStartOptions{value: val, isSet: true}
}

func (v NullableInstanceStartOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceStartOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

