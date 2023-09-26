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

// checks if the ProjectQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectQuota{}

// ProjectQuota 
type ProjectQuota struct {
	// 
	Cores NullableFloat32 `json:"cores,omitempty"`
	// 
	Instances NullableFloat32 `json:"instances,omitempty"`
	// 
	Ram NullableFloat32 `json:"ram,omitempty"`
}

// NewProjectQuota instantiates a new ProjectQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectQuota() *ProjectQuota {
	this := ProjectQuota{}
	return &this
}

// NewProjectQuotaWithDefaults instantiates a new ProjectQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectQuotaWithDefaults() *ProjectQuota {
	this := ProjectQuota{}
	return &this
}

// GetCores returns the Cores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectQuota) GetCores() float32 {
	if o == nil || IsNil(o.Cores.Get()) {
		var ret float32
		return ret
	}
	return *o.Cores.Get()
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectQuota) GetCoresOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cores.Get(), o.Cores.IsSet()
}

// HasCores returns a boolean if a field has been set.
func (o *ProjectQuota) HasCores() bool {
	if o != nil && o.Cores.IsSet() {
		return true
	}

	return false
}

// SetCores gets a reference to the given NullableFloat32 and assigns it to the Cores field.
func (o *ProjectQuota) SetCores(v float32) {
	o.Cores.Set(&v)
}
// SetCoresNil sets the value for Cores to be an explicit nil
func (o *ProjectQuota) SetCoresNil() {
	o.Cores.Set(nil)
}

// UnsetCores ensures that no value is present for Cores, not even an explicit nil
func (o *ProjectQuota) UnsetCores() {
	o.Cores.Unset()
}

// GetInstances returns the Instances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectQuota) GetInstances() float32 {
	if o == nil || IsNil(o.Instances.Get()) {
		var ret float32
		return ret
	}
	return *o.Instances.Get()
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectQuota) GetInstancesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances.Get(), o.Instances.IsSet()
}

// HasInstances returns a boolean if a field has been set.
func (o *ProjectQuota) HasInstances() bool {
	if o != nil && o.Instances.IsSet() {
		return true
	}

	return false
}

// SetInstances gets a reference to the given NullableFloat32 and assigns it to the Instances field.
func (o *ProjectQuota) SetInstances(v float32) {
	o.Instances.Set(&v)
}
// SetInstancesNil sets the value for Instances to be an explicit nil
func (o *ProjectQuota) SetInstancesNil() {
	o.Instances.Set(nil)
}

// UnsetInstances ensures that no value is present for Instances, not even an explicit nil
func (o *ProjectQuota) UnsetInstances() {
	o.Instances.Unset()
}

// GetRam returns the Ram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectQuota) GetRam() float32 {
	if o == nil || IsNil(o.Ram.Get()) {
		var ret float32
		return ret
	}
	return *o.Ram.Get()
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectQuota) GetRamOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ram.Get(), o.Ram.IsSet()
}

// HasRam returns a boolean if a field has been set.
func (o *ProjectQuota) HasRam() bool {
	if o != nil && o.Ram.IsSet() {
		return true
	}

	return false
}

// SetRam gets a reference to the given NullableFloat32 and assigns it to the Ram field.
func (o *ProjectQuota) SetRam(v float32) {
	o.Ram.Set(&v)
}
// SetRamNil sets the value for Ram to be an explicit nil
func (o *ProjectQuota) SetRamNil() {
	o.Ram.Set(nil)
}

// UnsetRam ensures that no value is present for Ram, not even an explicit nil
func (o *ProjectQuota) UnsetRam() {
	o.Ram.Unset()
}

func (o ProjectQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cores.IsSet() {
		toSerialize["cores"] = o.Cores.Get()
	}
	if o.Instances.IsSet() {
		toSerialize["instances"] = o.Instances.Get()
	}
	if o.Ram.IsSet() {
		toSerialize["ram"] = o.Ram.Get()
	}
	return toSerialize, nil
}

type NullableProjectQuota struct {
	value *ProjectQuota
	isSet bool
}

func (v NullableProjectQuota) Get() *ProjectQuota {
	return v.value
}

func (v *NullableProjectQuota) Set(val *ProjectQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectQuota(val *ProjectQuota) *NullableProjectQuota {
	return &NullableProjectQuota{value: val, isSet: true}
}

func (v NullableProjectQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


