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

// checks if the CouponOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponOptions{}

// CouponOptions Coupon options
type CouponOptions struct {
	// 
	Type string `json:"type"`
	// Amount
	Amount float32 `json:"amount"`
	// Is Used. If true, this coupon has been used and cannot be used again
	Used bool `json:"used"`
}

// NewCouponOptions instantiates a new CouponOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponOptions(type_ string, amount float32, used bool) *CouponOptions {
	this := CouponOptions{}
	this.Type = type_
	this.Amount = amount
	this.Used = used
	return &this
}

// NewCouponOptionsWithDefaults instantiates a new CouponOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponOptionsWithDefaults() *CouponOptions {
	this := CouponOptions{}
	return &this
}

// GetType returns the Type field value
func (o *CouponOptions) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponOptions) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponOptions) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *CouponOptions) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CouponOptions) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CouponOptions) SetAmount(v float32) {
	o.Amount = v
}

// GetUsed returns the Used field value
func (o *CouponOptions) GetUsed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *CouponOptions) GetUsedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *CouponOptions) SetUsed(v bool) {
	o.Used = v
}

func (o CouponOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["amount"] = o.Amount
	toSerialize["used"] = o.Used
	return toSerialize, nil
}

type NullableCouponOptions struct {
	value *CouponOptions
	isSet bool
}

func (v NullableCouponOptions) Get() *CouponOptions {
	return v.value
}

func (v *NullableCouponOptions) Set(val *CouponOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponOptions(val *CouponOptions) *NullableCouponOptions {
	return &NullableCouponOptions{value: val, isSet: true}
}

func (v NullableCouponOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


