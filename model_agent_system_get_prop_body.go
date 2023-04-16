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

// checks if the AgentSystemGetPropBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentSystemGetPropBody{}

// AgentSystemGetPropBody 
type AgentSystemGetPropBody struct {
	// 
	Property string `json:"property"`
}

// NewAgentSystemGetPropBody instantiates a new AgentSystemGetPropBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSystemGetPropBody(property string) *AgentSystemGetPropBody {
	this := AgentSystemGetPropBody{}
	this.Property = property
	return &this
}

// NewAgentSystemGetPropBodyWithDefaults instantiates a new AgentSystemGetPropBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSystemGetPropBodyWithDefaults() *AgentSystemGetPropBody {
	this := AgentSystemGetPropBody{}
	return &this
}

// GetProperty returns the Property field value
func (o *AgentSystemGetPropBody) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *AgentSystemGetPropBody) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value
func (o *AgentSystemGetPropBody) SetProperty(v string) {
	o.Property = v
}

func (o AgentSystemGetPropBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentSystemGetPropBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property"] = o.Property
	return toSerialize, nil
}

type NullableAgentSystemGetPropBody struct {
	value *AgentSystemGetPropBody
	isSet bool
}

func (v NullableAgentSystemGetPropBody) Get() *AgentSystemGetPropBody {
	return v.value
}

func (v *NullableAgentSystemGetPropBody) Set(val *AgentSystemGetPropBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSystemGetPropBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSystemGetPropBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSystemGetPropBody(val *AgentSystemGetPropBody) *NullableAgentSystemGetPropBody {
	return &NullableAgentSystemGetPropBody{value: val, isSet: true}
}

func (v NullableAgentSystemGetPropBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSystemGetPropBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

