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

// checks if the InviteRevokeParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteRevokeParams{}

// InviteRevokeParams 
type InviteRevokeParams struct {
	Ids NullableInviteRevokeParamsIds `json:"ids,omitempty"`
}

// NewInviteRevokeParams instantiates a new InviteRevokeParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteRevokeParams() *InviteRevokeParams {
	this := InviteRevokeParams{}
	return &this
}

// NewInviteRevokeParamsWithDefaults instantiates a new InviteRevokeParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteRevokeParamsWithDefaults() *InviteRevokeParams {
	this := InviteRevokeParams{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteRevokeParams) GetIds() InviteRevokeParamsIds {
	if o == nil || IsNil(o.Ids.Get()) {
		var ret InviteRevokeParamsIds
		return ret
	}
	return *o.Ids.Get()
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteRevokeParams) GetIdsOk() (*InviteRevokeParamsIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids.Get(), o.Ids.IsSet()
}

// HasIds returns a boolean if a field has been set.
func (o *InviteRevokeParams) HasIds() bool {
	if o != nil && o.Ids.IsSet() {
		return true
	}

	return false
}

// SetIds gets a reference to the given NullableInviteRevokeParamsIds and assigns it to the Ids field.
func (o *InviteRevokeParams) SetIds(v InviteRevokeParamsIds) {
	o.Ids.Set(&v)
}
// SetIdsNil sets the value for Ids to be an explicit nil
func (o *InviteRevokeParams) SetIdsNil() {
	o.Ids.Set(nil)
}

// UnsetIds ensures that no value is present for Ids, not even an explicit nil
func (o *InviteRevokeParams) UnsetIds() {
	o.Ids.Unset()
}

func (o InviteRevokeParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteRevokeParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids.IsSet() {
		toSerialize["ids"] = o.Ids.Get()
	}
	return toSerialize, nil
}

type NullableInviteRevokeParams struct {
	value *InviteRevokeParams
	isSet bool
}

func (v NullableInviteRevokeParams) Get() *InviteRevokeParams {
	return v.value
}

func (v *NullableInviteRevokeParams) Set(val *InviteRevokeParams) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteRevokeParams) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteRevokeParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteRevokeParams(val *InviteRevokeParams) *NullableInviteRevokeParams {
	return &NullableInviteRevokeParams{value: val, isSet: true}
}

func (v NullableInviteRevokeParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteRevokeParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


