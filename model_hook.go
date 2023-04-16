/*
Corellium API

REST API to manage your virtual devices.

API version: 4.3.1-16664
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the Hook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hook{}

// Hook 
type Hook struct {
	// Identifier
	Identifier NullableString `json:"identifier,omitempty"`
	// Label
	Label NullableString `json:"label,omitempty"`
	// Address
	Address NullableString `json:"address,omitempty"`
	// Patch
	Patch NullableString `json:"patch,omitempty"`
	// Patch Type
	PatchType NullableString `json:"patchType,omitempty"`
	// If true, instances can call required hooks
	Enabled NullableBool `json:"enabled,omitempty"`
	// Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	CreatedAt NullableString `json:"createdAt,omitempty"`
	// Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	UpdatedAt NullableString `json:"updatedAt,omitempty"`
	// Instance Id
	InstanceId NullableString `json:"instanceId,omitempty"`
}

// NewHook instantiates a new Hook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHook() *Hook {
	this := Hook{}
	return &this
}

// NewHookWithDefaults instantiates a new Hook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookWithDefaults() *Hook {
	this := Hook{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Hook) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *Hook) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *Hook) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *Hook) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *Hook) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *Hook) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *Hook) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *Hook) UnsetLabel() {
	o.Label.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *Hook) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *Hook) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *Hook) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *Hook) UnsetAddress() {
	o.Address.Unset()
}

// GetPatch returns the Patch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetPatch() string {
	if o == nil || IsNil(o.Patch.Get()) {
		var ret string
		return ret
	}
	return *o.Patch.Get()
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetPatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Patch.Get(), o.Patch.IsSet()
}

// HasPatch returns a boolean if a field has been set.
func (o *Hook) HasPatch() bool {
	if o != nil && o.Patch.IsSet() {
		return true
	}

	return false
}

// SetPatch gets a reference to the given NullableString and assigns it to the Patch field.
func (o *Hook) SetPatch(v string) {
	o.Patch.Set(&v)
}
// SetPatchNil sets the value for Patch to be an explicit nil
func (o *Hook) SetPatchNil() {
	o.Patch.Set(nil)
}

// UnsetPatch ensures that no value is present for Patch, not even an explicit nil
func (o *Hook) UnsetPatch() {
	o.Patch.Unset()
}

// GetPatchType returns the PatchType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetPatchType() string {
	if o == nil || IsNil(o.PatchType.Get()) {
		var ret string
		return ret
	}
	return *o.PatchType.Get()
}

// GetPatchTypeOk returns a tuple with the PatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetPatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PatchType.Get(), o.PatchType.IsSet()
}

// HasPatchType returns a boolean if a field has been set.
func (o *Hook) HasPatchType() bool {
	if o != nil && o.PatchType.IsSet() {
		return true
	}

	return false
}

// SetPatchType gets a reference to the given NullableString and assigns it to the PatchType field.
func (o *Hook) SetPatchType(v string) {
	o.PatchType.Set(&v)
}
// SetPatchTypeNil sets the value for PatchType to be an explicit nil
func (o *Hook) SetPatchTypeNil() {
	o.PatchType.Set(nil)
}

// UnsetPatchType ensures that no value is present for PatchType, not even an explicit nil
func (o *Hook) UnsetPatchType() {
	o.PatchType.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *Hook) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *Hook) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *Hook) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *Hook) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Hook) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Hook) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Hook) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Hook) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Hook) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *Hook) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Hook) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Hook) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Hook) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId.Get()) {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Hook) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *Hook) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *Hook) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}
// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *Hook) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *Hook) UnsetInstanceId() {
	o.InstanceId.Unset()
}

func (o Hook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Patch.IsSet() {
		toSerialize["patch"] = o.Patch.Get()
	}
	if o.PatchType.IsSet() {
		toSerialize["patchType"] = o.PatchType.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["instanceId"] = o.InstanceId.Get()
	}
	return toSerialize, nil
}

type NullableHook struct {
	value *Hook
	isSet bool
}

func (v NullableHook) Get() *Hook {
	return v.value
}

func (v *NullableHook) Set(val *Hook) {
	v.value = val
	v.isSet = true
}

func (v NullableHook) IsSet() bool {
	return v.isSet
}

func (v *NullableHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHook(val *Hook) *NullableHook {
	return &NullableHook{value: val, isSet: true}
}

func (v NullableHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

