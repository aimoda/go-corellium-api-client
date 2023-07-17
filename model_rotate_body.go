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

// checks if the RotateBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotateBody{}

// RotateBody 
type RotateBody struct {
	// Desired orientation
	Orientation float32 `json:"orientation"`
}

// NewRotateBody instantiates a new RotateBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotateBody(orientation float32) *RotateBody {
	this := RotateBody{}
	this.Orientation = orientation
	return &this
}

// NewRotateBodyWithDefaults instantiates a new RotateBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotateBodyWithDefaults() *RotateBody {
	this := RotateBody{}
	return &this
}

// GetOrientation returns the Orientation field value
func (o *RotateBody) GetOrientation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value
// and a boolean to check if the value has been set.
func (o *RotateBody) GetOrientationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orientation, true
}

// SetOrientation sets field value
func (o *RotateBody) SetOrientation(v float32) {
	o.Orientation = v
}

func (o RotateBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotateBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orientation"] = o.Orientation
	return toSerialize, nil
}

type NullableRotateBody struct {
	value *RotateBody
	isSet bool
}

func (v NullableRotateBody) Get() *RotateBody {
	return v.value
}

func (v *NullableRotateBody) Set(val *RotateBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRotateBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRotateBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotateBody(val *RotateBody) *NullableRotateBody {
	return &NullableRotateBody{value: val, isSet: true}
}

func (v NullableRotateBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotateBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


