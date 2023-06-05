/*
Corellium API

REST API to manage your virtual devices.

API version: 5.0.0-17089
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the AgentIcons type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentIcons{}

// AgentIcons 
type AgentIcons struct {
	// The data for an icon
	Icon NullableString `json:"icon,omitempty"`
}

// NewAgentIcons instantiates a new AgentIcons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentIcons() *AgentIcons {
	this := AgentIcons{}
	return &this
}

// NewAgentIconsWithDefaults instantiates a new AgentIcons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentIconsWithDefaults() *AgentIcons {
	this := AgentIcons{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentIcons) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentIcons) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AgentIcons) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AgentIcons) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *AgentIcons) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AgentIcons) UnsetIcon() {
	o.Icon.Unset()
}

func (o AgentIcons) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentIcons) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

type NullableAgentIcons struct {
	value *AgentIcons
	isSet bool
}

func (v NullableAgentIcons) Get() *AgentIcons {
	return v.value
}

func (v *NullableAgentIcons) Set(val *AgentIcons) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentIcons) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentIcons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentIcons(val *AgentIcons) *NullableAgentIcons {
	return &NullableAgentIcons{value: val, isSet: true}
}

func (v NullableAgentIcons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentIcons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


