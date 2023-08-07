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

// checks if the ApiToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiToken{}

// ApiToken 
type ApiToken struct {
	// 
	ApiToken string `json:"apiToken"`
}

// NewApiToken instantiates a new ApiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiToken(apiToken string) *ApiToken {
	this := ApiToken{}
	this.ApiToken = apiToken
	return &this
}

// NewApiTokenWithDefaults instantiates a new ApiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenWithDefaults() *ApiToken {
	this := ApiToken{}
	return &this
}

// GetApiToken returns the ApiToken field value
func (o *ApiToken) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *ApiToken) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *ApiToken) SetApiToken(v string) {
	o.ApiToken = v
}

func (o ApiToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiToken"] = o.ApiToken
	return toSerialize, nil
}

type NullableApiToken struct {
	value *ApiToken
	isSet bool
}

func (v NullableApiToken) Get() *ApiToken {
	return v.value
}

func (v *NullableApiToken) Set(val *ApiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableApiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableApiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiToken(val *ApiToken) *NullableApiToken {
	return &NullableApiToken{value: val, isSet: true}
}

func (v NullableApiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


