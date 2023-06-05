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

// checks if the AgentAppReadyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentAppReadyResponse{}

// AgentAppReadyResponse 
type AgentAppReadyResponse struct {
	// 
	Ready bool `json:"ready"`
}

// NewAgentAppReadyResponse instantiates a new AgentAppReadyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentAppReadyResponse(ready bool) *AgentAppReadyResponse {
	this := AgentAppReadyResponse{}
	this.Ready = ready
	return &this
}

// NewAgentAppReadyResponseWithDefaults instantiates a new AgentAppReadyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentAppReadyResponseWithDefaults() *AgentAppReadyResponse {
	this := AgentAppReadyResponse{}
	return &this
}

// GetReady returns the Ready field value
func (o *AgentAppReadyResponse) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *AgentAppReadyResponse) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *AgentAppReadyResponse) SetReady(v bool) {
	o.Ready = v
}

func (o AgentAppReadyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentAppReadyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ready"] = o.Ready
	return toSerialize, nil
}

type NullableAgentAppReadyResponse struct {
	value *AgentAppReadyResponse
	isSet bool
}

func (v NullableAgentAppReadyResponse) Get() *AgentAppReadyResponse {
	return v.value
}

func (v *NullableAgentAppReadyResponse) Set(val *AgentAppReadyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentAppReadyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentAppReadyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentAppReadyResponse(val *AgentAppReadyResponse) *NullableAgentAppReadyResponse {
	return &NullableAgentAppReadyResponse{value: val, isSet: true}
}

func (v NullableAgentAppReadyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentAppReadyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


