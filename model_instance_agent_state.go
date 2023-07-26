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

// checks if the InstanceAgentState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAgentState{}

// InstanceAgentState 
type InstanceAgentState struct {
	// 
	Hash NullableString `json:"hash,omitempty"`
	// 
	Info NullableString `json:"info,omitempty"`
}

// NewInstanceAgentState instantiates a new InstanceAgentState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAgentState() *InstanceAgentState {
	this := InstanceAgentState{}
	return &this
}

// NewInstanceAgentStateWithDefaults instantiates a new InstanceAgentState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAgentStateWithDefaults() *InstanceAgentState {
	this := InstanceAgentState{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAgentState) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAgentState) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *InstanceAgentState) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *InstanceAgentState) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *InstanceAgentState) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *InstanceAgentState) UnsetHash() {
	o.Hash.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceAgentState) GetInfo() string {
	if o == nil || IsNil(o.Info.Get()) {
		var ret string
		return ret
	}
	return *o.Info.Get()
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAgentState) GetInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info.Get(), o.Info.IsSet()
}

// HasInfo returns a boolean if a field has been set.
func (o *InstanceAgentState) HasInfo() bool {
	if o != nil && o.Info.IsSet() {
		return true
	}

	return false
}

// SetInfo gets a reference to the given NullableString and assigns it to the Info field.
func (o *InstanceAgentState) SetInfo(v string) {
	o.Info.Set(&v)
}
// SetInfoNil sets the value for Info to be an explicit nil
func (o *InstanceAgentState) SetInfoNil() {
	o.Info.Set(nil)
}

// UnsetInfo ensures that no value is present for Info, not even an explicit nil
func (o *InstanceAgentState) UnsetInfo() {
	o.Info.Unset()
}

func (o InstanceAgentState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAgentState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	if o.Info.IsSet() {
		toSerialize["info"] = o.Info.Get()
	}
	return toSerialize, nil
}

type NullableInstanceAgentState struct {
	value *InstanceAgentState
	isSet bool
}

func (v NullableInstanceAgentState) Get() *InstanceAgentState {
	return v.value
}

func (v *NullableInstanceAgentState) Set(val *InstanceAgentState) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAgentState) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAgentState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAgentState(val *InstanceAgentState) *NullableInstanceAgentState {
	return &NullableInstanceAgentState{value: val, isSet: true}
}

func (v NullableInstanceAgentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAgentState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


