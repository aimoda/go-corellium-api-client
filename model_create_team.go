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

// checks if the CreateTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTeam{}

// CreateTeam 
type CreateTeam struct {
	// Team name
	Name string `json:"name"`
}

// NewCreateTeam instantiates a new CreateTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTeam(name string) *CreateTeam {
	this := CreateTeam{}
	this.Name = name
	return &this
}

// NewCreateTeamWithDefaults instantiates a new CreateTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTeamWithDefaults() *CreateTeam {
	this := CreateTeam{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTeam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTeam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTeam) SetName(v string) {
	o.Name = v
}

func (o CreateTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCreateTeam struct {
	value *CreateTeam
	isSet bool
}

func (v NullableCreateTeam) Get() *CreateTeam {
	return v.value
}

func (v *NullableCreateTeam) Set(val *CreateTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTeam(val *CreateTeam) *NullableCreateTeam {
	return &NullableCreateTeam{value: val, isSet: true}
}

func (v NullableCreateTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


