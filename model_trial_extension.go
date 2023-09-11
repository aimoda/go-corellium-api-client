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

// checks if the TrialExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrialExtension{}

// TrialExtension 
type TrialExtension struct {
	// how many additional days?
	Duration NullableFloat32 `json:"duration,omitempty"`
	// why does the user want more time?
	Reason NullableString `json:"reason,omitempty"`
	// 
	Denied NullableBool `json:"denied,omitempty"`
	// explanation for why the request was denied
	DeniedReason NullableString `json:"deniedReason,omitempty"`
}

// NewTrialExtension instantiates a new TrialExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialExtension() *TrialExtension {
	this := TrialExtension{}
	return &this
}

// NewTrialExtensionWithDefaults instantiates a new TrialExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialExtensionWithDefaults() *TrialExtension {
	this := TrialExtension{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialExtension) GetDuration() float32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret float32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialExtension) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TrialExtension) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableFloat32 and assigns it to the Duration field.
func (o *TrialExtension) SetDuration(v float32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TrialExtension) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TrialExtension) UnsetDuration() {
	o.Duration.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialExtension) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialExtension) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *TrialExtension) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *TrialExtension) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *TrialExtension) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *TrialExtension) UnsetReason() {
	o.Reason.Unset()
}

// GetDenied returns the Denied field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialExtension) GetDenied() bool {
	if o == nil || IsNil(o.Denied.Get()) {
		var ret bool
		return ret
	}
	return *o.Denied.Get()
}

// GetDeniedOk returns a tuple with the Denied field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialExtension) GetDeniedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Denied.Get(), o.Denied.IsSet()
}

// HasDenied returns a boolean if a field has been set.
func (o *TrialExtension) HasDenied() bool {
	if o != nil && o.Denied.IsSet() {
		return true
	}

	return false
}

// SetDenied gets a reference to the given NullableBool and assigns it to the Denied field.
func (o *TrialExtension) SetDenied(v bool) {
	o.Denied.Set(&v)
}
// SetDeniedNil sets the value for Denied to be an explicit nil
func (o *TrialExtension) SetDeniedNil() {
	o.Denied.Set(nil)
}

// UnsetDenied ensures that no value is present for Denied, not even an explicit nil
func (o *TrialExtension) UnsetDenied() {
	o.Denied.Unset()
}

// GetDeniedReason returns the DeniedReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialExtension) GetDeniedReason() string {
	if o == nil || IsNil(o.DeniedReason.Get()) {
		var ret string
		return ret
	}
	return *o.DeniedReason.Get()
}

// GetDeniedReasonOk returns a tuple with the DeniedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialExtension) GetDeniedReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeniedReason.Get(), o.DeniedReason.IsSet()
}

// HasDeniedReason returns a boolean if a field has been set.
func (o *TrialExtension) HasDeniedReason() bool {
	if o != nil && o.DeniedReason.IsSet() {
		return true
	}

	return false
}

// SetDeniedReason gets a reference to the given NullableString and assigns it to the DeniedReason field.
func (o *TrialExtension) SetDeniedReason(v string) {
	o.DeniedReason.Set(&v)
}
// SetDeniedReasonNil sets the value for DeniedReason to be an explicit nil
func (o *TrialExtension) SetDeniedReasonNil() {
	o.DeniedReason.Set(nil)
}

// UnsetDeniedReason ensures that no value is present for DeniedReason, not even an explicit nil
func (o *TrialExtension) UnsetDeniedReason() {
	o.DeniedReason.Unset()
}

func (o TrialExtension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrialExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Denied.IsSet() {
		toSerialize["denied"] = o.Denied.Get()
	}
	if o.DeniedReason.IsSet() {
		toSerialize["deniedReason"] = o.DeniedReason.Get()
	}
	return toSerialize, nil
}

type NullableTrialExtension struct {
	value *TrialExtension
	isSet bool
}

func (v NullableTrialExtension) Get() *TrialExtension {
	return v.value
}

func (v *NullableTrialExtension) Set(val *TrialExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialExtension(val *TrialExtension) *NullableTrialExtension {
	return &NullableTrialExtension{value: val, isSet: true}
}

func (v NullableTrialExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


