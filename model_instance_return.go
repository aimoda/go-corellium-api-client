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

// checks if the InstanceReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceReturn{}

// InstanceReturn 
type InstanceReturn struct {
	// Instance ID
	Id string `json:"id"`
	State InstanceState `json:"state"`
}

// NewInstanceReturn instantiates a new InstanceReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceReturn(id string, state InstanceState) *InstanceReturn {
	this := InstanceReturn{}
	this.Id = id
	this.State = state
	return &this
}

// NewInstanceReturnWithDefaults instantiates a new InstanceReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceReturnWithDefaults() *InstanceReturn {
	this := InstanceReturn{}
	return &this
}

// GetId returns the Id field value
func (o *InstanceReturn) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InstanceReturn) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InstanceReturn) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *InstanceReturn) GetState() InstanceState {
	if o == nil {
		var ret InstanceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstanceReturn) GetStateOk() (*InstanceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InstanceReturn) SetState(v InstanceState) {
	o.State = v
}

func (o InstanceReturn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableInstanceReturn struct {
	value *InstanceReturn
	isSet bool
}

func (v NullableInstanceReturn) Get() *InstanceReturn {
	return v.value
}

func (v *NullableInstanceReturn) Set(val *InstanceReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceReturn(val *InstanceReturn) *NullableInstanceReturn {
	return &NullableInstanceReturn{value: val, isSet: true}
}

func (v NullableInstanceReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


