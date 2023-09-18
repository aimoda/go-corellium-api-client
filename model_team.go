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

// checks if the Team type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Team{}

// Team 
type Team struct {
	// Team ID
	Id string `json:"id"`
	// Team Label
	Label string `json:"label"`
	// Users
	Users []User `json:"users,omitempty"`
}

// NewTeam instantiates a new Team object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeam(id string, label string) *Team {
	this := Team{}
	this.Id = id
	this.Label = label
	return &this
}

// NewTeamWithDefaults instantiates a new Team object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamWithDefaults() *Team {
	this := Team{}
	return &this
}

// GetId returns the Id field value
func (o *Team) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Team) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Team) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *Team) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Team) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Team) SetLabel(v string) {
	o.Label = v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Team) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Team) HasUsers() bool {
	if o != nil && IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *Team) SetUsers(v []User) {
	o.Users = v
}

func (o Team) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Team) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableTeam struct {
	value *Team
	isSet bool
}

func (v NullableTeam) Get() *Team {
	return v.value
}

func (v *NullableTeam) Set(val *Team) {
	v.value = val
	v.isSet = true
}

func (v NullableTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeam(val *Team) *NullableTeam {
	return &NullableTeam{value: val, isSet: true}
}

func (v NullableTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


