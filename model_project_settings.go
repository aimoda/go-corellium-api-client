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

// checks if the ProjectSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSettings{}

// ProjectSettings 
type ProjectSettings struct {
	// 
	InternetAccess NullableBool `json:"internet-access,omitempty"`
	// 
	Dhcp NullableBool `json:"dhcp,omitempty"`
}

// NewProjectSettings instantiates a new ProjectSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSettings() *ProjectSettings {
	this := ProjectSettings{}
	return &this
}

// NewProjectSettingsWithDefaults instantiates a new ProjectSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSettingsWithDefaults() *ProjectSettings {
	this := ProjectSettings{}
	return &this
}

// GetInternetAccess returns the InternetAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSettings) GetInternetAccess() bool {
	if o == nil || IsNil(o.InternetAccess.Get()) {
		var ret bool
		return ret
	}
	return *o.InternetAccess.Get()
}

// GetInternetAccessOk returns a tuple with the InternetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSettings) GetInternetAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternetAccess.Get(), o.InternetAccess.IsSet()
}

// HasInternetAccess returns a boolean if a field has been set.
func (o *ProjectSettings) HasInternetAccess() bool {
	if o != nil && o.InternetAccess.IsSet() {
		return true
	}

	return false
}

// SetInternetAccess gets a reference to the given NullableBool and assigns it to the InternetAccess field.
func (o *ProjectSettings) SetInternetAccess(v bool) {
	o.InternetAccess.Set(&v)
}
// SetInternetAccessNil sets the value for InternetAccess to be an explicit nil
func (o *ProjectSettings) SetInternetAccessNil() {
	o.InternetAccess.Set(nil)
}

// UnsetInternetAccess ensures that no value is present for InternetAccess, not even an explicit nil
func (o *ProjectSettings) UnsetInternetAccess() {
	o.InternetAccess.Unset()
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSettings) GetDhcp() bool {
	if o == nil || IsNil(o.Dhcp.Get()) {
		var ret bool
		return ret
	}
	return *o.Dhcp.Get()
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSettings) GetDhcpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dhcp.Get(), o.Dhcp.IsSet()
}

// HasDhcp returns a boolean if a field has been set.
func (o *ProjectSettings) HasDhcp() bool {
	if o != nil && o.Dhcp.IsSet() {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given NullableBool and assigns it to the Dhcp field.
func (o *ProjectSettings) SetDhcp(v bool) {
	o.Dhcp.Set(&v)
}
// SetDhcpNil sets the value for Dhcp to be an explicit nil
func (o *ProjectSettings) SetDhcpNil() {
	o.Dhcp.Set(nil)
}

// UnsetDhcp ensures that no value is present for Dhcp, not even an explicit nil
func (o *ProjectSettings) UnsetDhcp() {
	o.Dhcp.Unset()
}

func (o ProjectSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InternetAccess.IsSet() {
		toSerialize["internet-access"] = o.InternetAccess.Get()
	}
	if o.Dhcp.IsSet() {
		toSerialize["dhcp"] = o.Dhcp.Get()
	}
	return toSerialize, nil
}

type NullableProjectSettings struct {
	value *ProjectSettings
	isSet bool
}

func (v NullableProjectSettings) Get() *ProjectSettings {
	return v.value
}

func (v *NullableProjectSettings) Set(val *ProjectSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSettings(val *ProjectSettings) *NullableProjectSettings {
	return &NullableProjectSettings{value: val, isSet: true}
}

func (v NullableProjectSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


