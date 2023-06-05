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

// checks if the Trial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trial{}

// Trial Trial options
type Trial struct {
	// Duration in days
	Duration float32 `json:"duration"`
}

// NewTrial instantiates a new Trial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrial(duration float32) *Trial {
	this := Trial{}
	this.Duration = duration
	return &this
}

// NewTrialWithDefaults instantiates a new Trial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialWithDefaults() *Trial {
	this := Trial{}
	return &this
}

// GetDuration returns the Duration field value
func (o *Trial) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *Trial) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *Trial) SetDuration(v float32) {
	o.Duration = v
}

func (o Trial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	return toSerialize, nil
}

type NullableTrial struct {
	value *Trial
	isSet bool
}

func (v NullableTrial) Get() *Trial {
	return v.value
}

func (v *NullableTrial) Set(val *Trial) {
	v.value = val
	v.isSet = true
}

func (v NullableTrial) IsSet() bool {
	return v.isSet
}

func (v *NullableTrial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrial(val *Trial) *NullableTrial {
	return &NullableTrial{value: val, isSet: true}
}

func (v NullableTrial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


