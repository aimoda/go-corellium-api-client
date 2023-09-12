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

// checks if the KernelTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KernelTask{}

// KernelTask 
type KernelTask struct {
	// Kernel Task ID
	KernelId NullableString `json:"kernelId,omitempty"`
	// Thread name
	Name NullableString `json:"name,omitempty"`
	// Process ID of task
	Pid NullableInt32 `json:"pid,omitempty"`
	// 
	Threads []KernelThread `json:"threads,omitempty"`
}

// NewKernelTask instantiates a new KernelTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKernelTask() *KernelTask {
	this := KernelTask{}
	return &this
}

// NewKernelTaskWithDefaults instantiates a new KernelTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKernelTaskWithDefaults() *KernelTask {
	this := KernelTask{}
	return &this
}

// GetKernelId returns the KernelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelTask) GetKernelId() string {
	if o == nil || IsNil(o.KernelId.Get()) {
		var ret string
		return ret
	}
	return *o.KernelId.Get()
}

// GetKernelIdOk returns a tuple with the KernelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelTask) GetKernelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KernelId.Get(), o.KernelId.IsSet()
}

// HasKernelId returns a boolean if a field has been set.
func (o *KernelTask) HasKernelId() bool {
	if o != nil && o.KernelId.IsSet() {
		return true
	}

	return false
}

// SetKernelId gets a reference to the given NullableString and assigns it to the KernelId field.
func (o *KernelTask) SetKernelId(v string) {
	o.KernelId.Set(&v)
}
// SetKernelIdNil sets the value for KernelId to be an explicit nil
func (o *KernelTask) SetKernelIdNil() {
	o.KernelId.Set(nil)
}

// UnsetKernelId ensures that no value is present for KernelId, not even an explicit nil
func (o *KernelTask) UnsetKernelId() {
	o.KernelId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelTask) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelTask) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KernelTask) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KernelTask) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KernelTask) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KernelTask) UnsetName() {
	o.Name.Unset()
}

// GetPid returns the Pid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelTask) GetPid() int32 {
	if o == nil || IsNil(o.Pid.Get()) {
		var ret int32
		return ret
	}
	return *o.Pid.Get()
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelTask) GetPidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pid.Get(), o.Pid.IsSet()
}

// HasPid returns a boolean if a field has been set.
func (o *KernelTask) HasPid() bool {
	if o != nil && o.Pid.IsSet() {
		return true
	}

	return false
}

// SetPid gets a reference to the given NullableInt32 and assigns it to the Pid field.
func (o *KernelTask) SetPid(v int32) {
	o.Pid.Set(&v)
}
// SetPidNil sets the value for Pid to be an explicit nil
func (o *KernelTask) SetPidNil() {
	o.Pid.Set(nil)
}

// UnsetPid ensures that no value is present for Pid, not even an explicit nil
func (o *KernelTask) UnsetPid() {
	o.Pid.Unset()
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KernelTask) GetThreads() []KernelThread {
	if o == nil {
		var ret []KernelThread
		return ret
	}
	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KernelTask) GetThreadsOk() ([]KernelThread, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *KernelTask) HasThreads() bool {
	if o != nil && IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given []KernelThread and assigns it to the Threads field.
func (o *KernelTask) SetThreads(v []KernelThread) {
	o.Threads = v
}

func (o KernelTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KernelTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.KernelId.IsSet() {
		toSerialize["kernelId"] = o.KernelId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Pid.IsSet() {
		toSerialize["pid"] = o.Pid.Get()
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}
	return toSerialize, nil
}

type NullableKernelTask struct {
	value *KernelTask
	isSet bool
}

func (v NullableKernelTask) Get() *KernelTask {
	return v.value
}

func (v *NullableKernelTask) Set(val *KernelTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKernelTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKernelTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKernelTask(val *KernelTask) *NullableKernelTask {
	return &NullableKernelTask{value: val, isSet: true}
}

func (v NullableKernelTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKernelTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


