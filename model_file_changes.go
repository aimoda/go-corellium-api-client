/*
Corellium API

REST API to manage your virtual devices.

API version: 5.5.0-18750
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the FileChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileChanges{}

// FileChanges 
type FileChanges struct {
	// Optional New path
	Path NullableString `json:"path,omitempty"`
	// Optional New Mode
	Mode NullableFloat32 `json:"mode,omitempty"`
	// Optional New Owner UID
	Uid NullableFloat32 `json:"uid,omitempty"`
	// Optional New Group GID
	Gid NullableFloat32 `json:"gid,omitempty"`
}

// NewFileChanges instantiates a new FileChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileChanges() *FileChanges {
	this := FileChanges{}
	return &this
}

// NewFileChangesWithDefaults instantiates a new FileChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileChangesWithDefaults() *FileChanges {
	this := FileChanges{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileChanges) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileChanges) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *FileChanges) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *FileChanges) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *FileChanges) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *FileChanges) UnsetPath() {
	o.Path.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileChanges) GetMode() float32 {
	if o == nil || IsNil(o.Mode.Get()) {
		var ret float32
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileChanges) GetModeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *FileChanges) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableFloat32 and assigns it to the Mode field.
func (o *FileChanges) SetMode(v float32) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *FileChanges) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *FileChanges) UnsetMode() {
	o.Mode.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileChanges) GetUid() float32 {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret float32
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileChanges) GetUidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *FileChanges) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableFloat32 and assigns it to the Uid field.
func (o *FileChanges) SetUid(v float32) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *FileChanges) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *FileChanges) UnsetUid() {
	o.Uid.Unset()
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileChanges) GetGid() float32 {
	if o == nil || IsNil(o.Gid.Get()) {
		var ret float32
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileChanges) GetGidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *FileChanges) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableFloat32 and assigns it to the Gid field.
func (o *FileChanges) SetGid(v float32) {
	o.Gid.Set(&v)
}
// SetGidNil sets the value for Gid to be an explicit nil
func (o *FileChanges) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *FileChanges) UnsetGid() {
	o.Gid.Unset()
}

func (o FileChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	return toSerialize, nil
}

type NullableFileChanges struct {
	value *FileChanges
	isSet bool
}

func (v NullableFileChanges) Get() *FileChanges {
	return v.value
}

func (v *NullableFileChanges) Set(val *FileChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableFileChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableFileChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileChanges(val *FileChanges) *NullableFileChanges {
	return &NullableFileChanges{value: val, isSet: true}
}

func (v NullableFileChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


