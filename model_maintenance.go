/*
Corellium API

REST API to manage your virtual devices.

API version: 5.4.1-18421
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the Maintenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Maintenance{}

// Maintenance 
type Maintenance struct {
	// Maintenance message
	Message NullableString `json:"message,omitempty"`
	// Maintenance header
	Header NullableString `json:"header,omitempty"`
}

// NewMaintenance instantiates a new Maintenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenance() *Maintenance {
	this := Maintenance{}
	return &this
}

// NewMaintenanceWithDefaults instantiates a new Maintenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWithDefaults() *Maintenance {
	this := Maintenance{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Maintenance) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Maintenance) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *Maintenance) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *Maintenance) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *Maintenance) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *Maintenance) UnsetMessage() {
	o.Message.Unset()
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Maintenance) GetHeader() string {
	if o == nil || IsNil(o.Header.Get()) {
		var ret string
		return ret
	}
	return *o.Header.Get()
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Maintenance) GetHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Header.Get(), o.Header.IsSet()
}

// HasHeader returns a boolean if a field has been set.
func (o *Maintenance) HasHeader() bool {
	if o != nil && o.Header.IsSet() {
		return true
	}

	return false
}

// SetHeader gets a reference to the given NullableString and assigns it to the Header field.
func (o *Maintenance) SetHeader(v string) {
	o.Header.Set(&v)
}
// SetHeaderNil sets the value for Header to be an explicit nil
func (o *Maintenance) SetHeaderNil() {
	o.Header.Set(nil)
}

// UnsetHeader ensures that no value is present for Header, not even an explicit nil
func (o *Maintenance) UnsetHeader() {
	o.Header.Unset()
}

func (o Maintenance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Maintenance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Header.IsSet() {
		toSerialize["header"] = o.Header.Get()
	}
	return toSerialize, nil
}

type NullableMaintenance struct {
	value *Maintenance
	isSet bool
}

func (v NullableMaintenance) Get() *Maintenance {
	return v.value
}

func (v *NullableMaintenance) Set(val *Maintenance) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenance) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenance(val *Maintenance) *NullableMaintenance {
	return &NullableMaintenance{value: val, isSet: true}
}

func (v NullableMaintenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


