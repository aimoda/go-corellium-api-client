/*
Corellium API

REST API to manage your virtual devices.

API version: 5.4.0-18254
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the PatchInstanceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchInstanceOptions{}

// PatchInstanceOptions 
type PatchInstanceOptions struct {
	// the name of the device
	Name NullableString `json:"name,omitempty"`
	// the desired state of the device
	State NullableString `json:"state,omitempty"`
	BootOptions *InstanceBootOptions `json:"bootOptions,omitempty"`
	// 
	Proxy []ProxyConfig `json:"proxy,omitempty"`
}

// NewPatchInstanceOptions instantiates a new PatchInstanceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchInstanceOptions() *PatchInstanceOptions {
	this := PatchInstanceOptions{}
	return &this
}

// NewPatchInstanceOptionsWithDefaults instantiates a new PatchInstanceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchInstanceOptionsWithDefaults() *PatchInstanceOptions {
	this := PatchInstanceOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstanceOptions) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstanceOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchInstanceOptions) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchInstanceOptions) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchInstanceOptions) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchInstanceOptions) UnsetName() {
	o.Name.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstanceOptions) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstanceOptions) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *PatchInstanceOptions) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *PatchInstanceOptions) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *PatchInstanceOptions) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *PatchInstanceOptions) UnsetState() {
	o.State.Unset()
}

// GetBootOptions returns the BootOptions field value if set, zero value otherwise.
func (o *PatchInstanceOptions) GetBootOptions() InstanceBootOptions {
	if o == nil || IsNil(o.BootOptions) {
		var ret InstanceBootOptions
		return ret
	}
	return *o.BootOptions
}

// GetBootOptionsOk returns a tuple with the BootOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchInstanceOptions) GetBootOptionsOk() (*InstanceBootOptions, bool) {
	if o == nil || IsNil(o.BootOptions) {
		return nil, false
	}
	return o.BootOptions, true
}

// HasBootOptions returns a boolean if a field has been set.
func (o *PatchInstanceOptions) HasBootOptions() bool {
	if o != nil && !IsNil(o.BootOptions) {
		return true
	}

	return false
}

// SetBootOptions gets a reference to the given InstanceBootOptions and assigns it to the BootOptions field.
func (o *PatchInstanceOptions) SetBootOptions(v InstanceBootOptions) {
	o.BootOptions = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchInstanceOptions) GetProxy() []ProxyConfig {
	if o == nil {
		var ret []ProxyConfig
		return ret
	}
	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchInstanceOptions) GetProxyOk() ([]ProxyConfig, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *PatchInstanceOptions) HasProxy() bool {
	if o != nil && IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given []ProxyConfig and assigns it to the Proxy field.
func (o *PatchInstanceOptions) SetProxy(v []ProxyConfig) {
	o.Proxy = v
}

func (o PatchInstanceOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchInstanceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.BootOptions) {
		toSerialize["bootOptions"] = o.BootOptions
	}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	return toSerialize, nil
}

type NullablePatchInstanceOptions struct {
	value *PatchInstanceOptions
	isSet bool
}

func (v NullablePatchInstanceOptions) Get() *PatchInstanceOptions {
	return v.value
}

func (v *NullablePatchInstanceOptions) Set(val *PatchInstanceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchInstanceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchInstanceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchInstanceOptions(val *PatchInstanceOptions) *NullablePatchInstanceOptions {
	return &NullablePatchInstanceOptions{value: val, isSet: true}
}

func (v NullablePatchInstanceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchInstanceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


