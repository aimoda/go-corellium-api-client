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

// checks if the VolumeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeOptions{}

// VolumeOptions 
type VolumeOptions struct {
	// GB
	Allocate NullableFloat32 `json:"allocate,omitempty"`
	// 
	Partitions []map[string]interface{} `json:"partitions,omitempty"`
	// 
	ComputeNode NullableString `json:"computeNode,omitempty"`
}

// NewVolumeOptions instantiates a new VolumeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeOptions() *VolumeOptions {
	this := VolumeOptions{}
	return &this
}

// NewVolumeOptionsWithDefaults instantiates a new VolumeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeOptionsWithDefaults() *VolumeOptions {
	this := VolumeOptions{}
	return &this
}

// GetAllocate returns the Allocate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VolumeOptions) GetAllocate() float32 {
	if o == nil || IsNil(o.Allocate.Get()) {
		var ret float32
		return ret
	}
	return *o.Allocate.Get()
}

// GetAllocateOk returns a tuple with the Allocate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeOptions) GetAllocateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Allocate.Get(), o.Allocate.IsSet()
}

// HasAllocate returns a boolean if a field has been set.
func (o *VolumeOptions) HasAllocate() bool {
	if o != nil && o.Allocate.IsSet() {
		return true
	}

	return false
}

// SetAllocate gets a reference to the given NullableFloat32 and assigns it to the Allocate field.
func (o *VolumeOptions) SetAllocate(v float32) {
	o.Allocate.Set(&v)
}
// SetAllocateNil sets the value for Allocate to be an explicit nil
func (o *VolumeOptions) SetAllocateNil() {
	o.Allocate.Set(nil)
}

// UnsetAllocate ensures that no value is present for Allocate, not even an explicit nil
func (o *VolumeOptions) UnsetAllocate() {
	o.Allocate.Unset()
}

// GetPartitions returns the Partitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VolumeOptions) GetPartitions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeOptions) GetPartitionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *VolumeOptions) HasPartitions() bool {
	if o != nil && IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []map[string]interface{} and assigns it to the Partitions field.
func (o *VolumeOptions) SetPartitions(v []map[string]interface{}) {
	o.Partitions = v
}

// GetComputeNode returns the ComputeNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VolumeOptions) GetComputeNode() string {
	if o == nil || IsNil(o.ComputeNode.Get()) {
		var ret string
		return ret
	}
	return *o.ComputeNode.Get()
}

// GetComputeNodeOk returns a tuple with the ComputeNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeOptions) GetComputeNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeNode.Get(), o.ComputeNode.IsSet()
}

// HasComputeNode returns a boolean if a field has been set.
func (o *VolumeOptions) HasComputeNode() bool {
	if o != nil && o.ComputeNode.IsSet() {
		return true
	}

	return false
}

// SetComputeNode gets a reference to the given NullableString and assigns it to the ComputeNode field.
func (o *VolumeOptions) SetComputeNode(v string) {
	o.ComputeNode.Set(&v)
}
// SetComputeNodeNil sets the value for ComputeNode to be an explicit nil
func (o *VolumeOptions) SetComputeNodeNil() {
	o.ComputeNode.Set(nil)
}

// UnsetComputeNode ensures that no value is present for ComputeNode, not even an explicit nil
func (o *VolumeOptions) UnsetComputeNode() {
	o.ComputeNode.Unset()
}

func (o VolumeOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Allocate.IsSet() {
		toSerialize["allocate"] = o.Allocate.Get()
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	if o.ComputeNode.IsSet() {
		toSerialize["computeNode"] = o.ComputeNode.Get()
	}
	return toSerialize, nil
}

type NullableVolumeOptions struct {
	value *VolumeOptions
	isSet bool
}

func (v NullableVolumeOptions) Get() *VolumeOptions {
	return v.value
}

func (v *NullableVolumeOptions) Set(val *VolumeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeOptions(val *VolumeOptions) *NullableVolumeOptions {
	return &NullableVolumeOptions{value: val, isSet: true}
}

func (v NullableVolumeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


