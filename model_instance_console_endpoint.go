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

// checks if the InstanceConsoleEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceConsoleEndpoint{}

// InstanceConsoleEndpoint 
type InstanceConsoleEndpoint struct {
	// Console Websocket URL
	Url NullableString `json:"url,omitempty"`
}

// NewInstanceConsoleEndpoint instantiates a new InstanceConsoleEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceConsoleEndpoint() *InstanceConsoleEndpoint {
	this := InstanceConsoleEndpoint{}
	return &this
}

// NewInstanceConsoleEndpointWithDefaults instantiates a new InstanceConsoleEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceConsoleEndpointWithDefaults() *InstanceConsoleEndpoint {
	this := InstanceConsoleEndpoint{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceConsoleEndpoint) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceConsoleEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *InstanceConsoleEndpoint) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *InstanceConsoleEndpoint) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *InstanceConsoleEndpoint) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *InstanceConsoleEndpoint) UnsetUrl() {
	o.Url.Unset()
}

func (o InstanceConsoleEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceConsoleEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableInstanceConsoleEndpoint struct {
	value *InstanceConsoleEndpoint
	isSet bool
}

func (v NullableInstanceConsoleEndpoint) Get() *InstanceConsoleEndpoint {
	return v.value
}

func (v *NullableInstanceConsoleEndpoint) Set(val *InstanceConsoleEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceConsoleEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceConsoleEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceConsoleEndpoint(val *InstanceConsoleEndpoint) *NullableInstanceConsoleEndpoint {
	return &NullableInstanceConsoleEndpoint{value: val, isSet: true}
}

func (v NullableInstanceConsoleEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceConsoleEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

