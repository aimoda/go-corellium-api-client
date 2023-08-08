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

// checks if the InstanceServices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceServices{}

// InstanceServices 
type InstanceServices struct {
	Vpn *VpnDefinition `json:"vpn,omitempty"`
}

// NewInstanceServices instantiates a new InstanceServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceServices() *InstanceServices {
	this := InstanceServices{}
	return &this
}

// NewInstanceServicesWithDefaults instantiates a new InstanceServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceServicesWithDefaults() *InstanceServices {
	this := InstanceServices{}
	return &this
}

// GetVpn returns the Vpn field value if set, zero value otherwise.
func (o *InstanceServices) GetVpn() VpnDefinition {
	if o == nil || IsNil(o.Vpn) {
		var ret VpnDefinition
		return ret
	}
	return *o.Vpn
}

// GetVpnOk returns a tuple with the Vpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceServices) GetVpnOk() (*VpnDefinition, bool) {
	if o == nil || IsNil(o.Vpn) {
		return nil, false
	}
	return o.Vpn, true
}

// HasVpn returns a boolean if a field has been set.
func (o *InstanceServices) HasVpn() bool {
	if o != nil && !IsNil(o.Vpn) {
		return true
	}

	return false
}

// SetVpn gets a reference to the given VpnDefinition and assigns it to the Vpn field.
func (o *InstanceServices) SetVpn(v VpnDefinition) {
	o.Vpn = &v
}

func (o InstanceServices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceServices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vpn) {
		toSerialize["vpn"] = o.Vpn
	}
	return toSerialize, nil
}

type NullableInstanceServices struct {
	value *InstanceServices
	isSet bool
}

func (v NullableInstanceServices) Get() *InstanceServices {
	return v.value
}

func (v *NullableInstanceServices) Set(val *InstanceServices) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceServices) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceServices(val *InstanceServices) *NullableInstanceServices {
	return &NullableInstanceServices{value: val, isSet: true}
}

func (v NullableInstanceServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


