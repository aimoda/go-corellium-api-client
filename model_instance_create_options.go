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

// checks if the InstanceCreateOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceCreateOptions{}

// InstanceCreateOptions 
type InstanceCreateOptions struct {
	// the name of the device
	Name NullableString `json:"name,omitempty"`
	// Key used to encrypt the Instance
	Key NullableString `json:"key,omitempty"`
	// the flavor id
	Flavor string `json:"flavor"`
	// project UUID
	Project string `json:"project"`
	// OS Version
	Os string `json:"os"`
	// OS Build
	Osbuild NullableString `json:"osbuild,omitempty"`
	// list of patches to apply
	Patches []string `json:"patches,omitempty"`
	// URL or image id
	Fwpackage NullableString `json:"fwpackage,omitempty"`
	// URL that firmware package used to create this instance is available at
	OrigFwPackageUrl NullableString `json:"origFwPackageUrl,omitempty"`
	// 
	Encrypt NullableBool `json:"encrypt,omitempty"`
	// 
	OverrideWifiMAC NullableString `json:"overrideWifiMAC,omitempty"`
	Volume *VolumeOptions `json:"volume,omitempty"`
	// Snapshot ID for this instance to be cloned from if defined
	Snapshot NullableString `json:"snapshot,omitempty"`
	BootOptions *InstanceBootOptions `json:"bootOptions,omitempty"`
	Device *Model `json:"device,omitempty"`
}

// NewInstanceCreateOptions instantiates a new InstanceCreateOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceCreateOptions(flavor string, project string, os string) *InstanceCreateOptions {
	this := InstanceCreateOptions{}
	this.Flavor = flavor
	this.Project = project
	this.Os = os
	return &this
}

// NewInstanceCreateOptionsWithDefaults instantiates a new InstanceCreateOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceCreateOptionsWithDefaults() *InstanceCreateOptions {
	this := InstanceCreateOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InstanceCreateOptions) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InstanceCreateOptions) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InstanceCreateOptions) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *InstanceCreateOptions) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *InstanceCreateOptions) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *InstanceCreateOptions) UnsetKey() {
	o.Key.Unset()
}

// GetFlavor returns the Flavor field value
func (o *InstanceCreateOptions) GetFlavor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flavor, true
}

// SetFlavor sets field value
func (o *InstanceCreateOptions) SetFlavor(v string) {
	o.Flavor = v
}

// GetProject returns the Project field value
func (o *InstanceCreateOptions) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *InstanceCreateOptions) SetProject(v string) {
	o.Project = v
}

// GetOs returns the Os field value
func (o *InstanceCreateOptions) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *InstanceCreateOptions) SetOs(v string) {
	o.Os = v
}

// GetOsbuild returns the Osbuild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetOsbuild() string {
	if o == nil || IsNil(o.Osbuild.Get()) {
		var ret string
		return ret
	}
	return *o.Osbuild.Get()
}

// GetOsbuildOk returns a tuple with the Osbuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetOsbuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Osbuild.Get(), o.Osbuild.IsSet()
}

// HasOsbuild returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasOsbuild() bool {
	if o != nil && o.Osbuild.IsSet() {
		return true
	}

	return false
}

// SetOsbuild gets a reference to the given NullableString and assigns it to the Osbuild field.
func (o *InstanceCreateOptions) SetOsbuild(v string) {
	o.Osbuild.Set(&v)
}
// SetOsbuildNil sets the value for Osbuild to be an explicit nil
func (o *InstanceCreateOptions) SetOsbuildNil() {
	o.Osbuild.Set(nil)
}

// UnsetOsbuild ensures that no value is present for Osbuild, not even an explicit nil
func (o *InstanceCreateOptions) UnsetOsbuild() {
	o.Osbuild.Unset()
}

// GetPatches returns the Patches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetPatches() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetPatchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Patches) {
		return nil, false
	}
	return o.Patches, true
}

// HasPatches returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasPatches() bool {
	if o != nil && IsNil(o.Patches) {
		return true
	}

	return false
}

// SetPatches gets a reference to the given []string and assigns it to the Patches field.
func (o *InstanceCreateOptions) SetPatches(v []string) {
	o.Patches = v
}

// GetFwpackage returns the Fwpackage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetFwpackage() string {
	if o == nil || IsNil(o.Fwpackage.Get()) {
		var ret string
		return ret
	}
	return *o.Fwpackage.Get()
}

// GetFwpackageOk returns a tuple with the Fwpackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetFwpackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fwpackage.Get(), o.Fwpackage.IsSet()
}

// HasFwpackage returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasFwpackage() bool {
	if o != nil && o.Fwpackage.IsSet() {
		return true
	}

	return false
}

// SetFwpackage gets a reference to the given NullableString and assigns it to the Fwpackage field.
func (o *InstanceCreateOptions) SetFwpackage(v string) {
	o.Fwpackage.Set(&v)
}
// SetFwpackageNil sets the value for Fwpackage to be an explicit nil
func (o *InstanceCreateOptions) SetFwpackageNil() {
	o.Fwpackage.Set(nil)
}

// UnsetFwpackage ensures that no value is present for Fwpackage, not even an explicit nil
func (o *InstanceCreateOptions) UnsetFwpackage() {
	o.Fwpackage.Unset()
}

// GetOrigFwPackageUrl returns the OrigFwPackageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetOrigFwPackageUrl() string {
	if o == nil || IsNil(o.OrigFwPackageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OrigFwPackageUrl.Get()
}

// GetOrigFwPackageUrlOk returns a tuple with the OrigFwPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetOrigFwPackageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrigFwPackageUrl.Get(), o.OrigFwPackageUrl.IsSet()
}

// HasOrigFwPackageUrl returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasOrigFwPackageUrl() bool {
	if o != nil && o.OrigFwPackageUrl.IsSet() {
		return true
	}

	return false
}

// SetOrigFwPackageUrl gets a reference to the given NullableString and assigns it to the OrigFwPackageUrl field.
func (o *InstanceCreateOptions) SetOrigFwPackageUrl(v string) {
	o.OrigFwPackageUrl.Set(&v)
}
// SetOrigFwPackageUrlNil sets the value for OrigFwPackageUrl to be an explicit nil
func (o *InstanceCreateOptions) SetOrigFwPackageUrlNil() {
	o.OrigFwPackageUrl.Set(nil)
}

// UnsetOrigFwPackageUrl ensures that no value is present for OrigFwPackageUrl, not even an explicit nil
func (o *InstanceCreateOptions) UnsetOrigFwPackageUrl() {
	o.OrigFwPackageUrl.Unset()
}

// GetEncrypt returns the Encrypt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetEncrypt() bool {
	if o == nil || IsNil(o.Encrypt.Get()) {
		var ret bool
		return ret
	}
	return *o.Encrypt.Get()
}

// GetEncryptOk returns a tuple with the Encrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetEncryptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encrypt.Get(), o.Encrypt.IsSet()
}

// HasEncrypt returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasEncrypt() bool {
	if o != nil && o.Encrypt.IsSet() {
		return true
	}

	return false
}

// SetEncrypt gets a reference to the given NullableBool and assigns it to the Encrypt field.
func (o *InstanceCreateOptions) SetEncrypt(v bool) {
	o.Encrypt.Set(&v)
}
// SetEncryptNil sets the value for Encrypt to be an explicit nil
func (o *InstanceCreateOptions) SetEncryptNil() {
	o.Encrypt.Set(nil)
}

// UnsetEncrypt ensures that no value is present for Encrypt, not even an explicit nil
func (o *InstanceCreateOptions) UnsetEncrypt() {
	o.Encrypt.Unset()
}

// GetOverrideWifiMAC returns the OverrideWifiMAC field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetOverrideWifiMAC() string {
	if o == nil || IsNil(o.OverrideWifiMAC.Get()) {
		var ret string
		return ret
	}
	return *o.OverrideWifiMAC.Get()
}

// GetOverrideWifiMACOk returns a tuple with the OverrideWifiMAC field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetOverrideWifiMACOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverrideWifiMAC.Get(), o.OverrideWifiMAC.IsSet()
}

// HasOverrideWifiMAC returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasOverrideWifiMAC() bool {
	if o != nil && o.OverrideWifiMAC.IsSet() {
		return true
	}

	return false
}

// SetOverrideWifiMAC gets a reference to the given NullableString and assigns it to the OverrideWifiMAC field.
func (o *InstanceCreateOptions) SetOverrideWifiMAC(v string) {
	o.OverrideWifiMAC.Set(&v)
}
// SetOverrideWifiMACNil sets the value for OverrideWifiMAC to be an explicit nil
func (o *InstanceCreateOptions) SetOverrideWifiMACNil() {
	o.OverrideWifiMAC.Set(nil)
}

// UnsetOverrideWifiMAC ensures that no value is present for OverrideWifiMAC, not even an explicit nil
func (o *InstanceCreateOptions) UnsetOverrideWifiMAC() {
	o.OverrideWifiMAC.Unset()
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *InstanceCreateOptions) GetVolume() VolumeOptions {
	if o == nil || IsNil(o.Volume) {
		var ret VolumeOptions
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetVolumeOk() (*VolumeOptions, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given VolumeOptions and assigns it to the Volume field.
func (o *InstanceCreateOptions) SetVolume(v VolumeOptions) {
	o.Volume = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceCreateOptions) GetSnapshot() string {
	if o == nil || IsNil(o.Snapshot.Get()) {
		var ret string
		return ret
	}
	return *o.Snapshot.Get()
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceCreateOptions) GetSnapshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Snapshot.Get(), o.Snapshot.IsSet()
}

// HasSnapshot returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasSnapshot() bool {
	if o != nil && o.Snapshot.IsSet() {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given NullableString and assigns it to the Snapshot field.
func (o *InstanceCreateOptions) SetSnapshot(v string) {
	o.Snapshot.Set(&v)
}
// SetSnapshotNil sets the value for Snapshot to be an explicit nil
func (o *InstanceCreateOptions) SetSnapshotNil() {
	o.Snapshot.Set(nil)
}

// UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
func (o *InstanceCreateOptions) UnsetSnapshot() {
	o.Snapshot.Unset()
}

// GetBootOptions returns the BootOptions field value if set, zero value otherwise.
func (o *InstanceCreateOptions) GetBootOptions() InstanceBootOptions {
	if o == nil || IsNil(o.BootOptions) {
		var ret InstanceBootOptions
		return ret
	}
	return *o.BootOptions
}

// GetBootOptionsOk returns a tuple with the BootOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetBootOptionsOk() (*InstanceBootOptions, bool) {
	if o == nil || IsNil(o.BootOptions) {
		return nil, false
	}
	return o.BootOptions, true
}

// HasBootOptions returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasBootOptions() bool {
	if o != nil && !IsNil(o.BootOptions) {
		return true
	}

	return false
}

// SetBootOptions gets a reference to the given InstanceBootOptions and assigns it to the BootOptions field.
func (o *InstanceCreateOptions) SetBootOptions(v InstanceBootOptions) {
	o.BootOptions = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InstanceCreateOptions) GetDevice() Model {
	if o == nil || IsNil(o.Device) {
		var ret Model
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceCreateOptions) GetDeviceOk() (*Model, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InstanceCreateOptions) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Model and assigns it to the Device field.
func (o *InstanceCreateOptions) SetDevice(v Model) {
	o.Device = &v
}

func (o InstanceCreateOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceCreateOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	toSerialize["flavor"] = o.Flavor
	toSerialize["project"] = o.Project
	toSerialize["os"] = o.Os
	if o.Osbuild.IsSet() {
		toSerialize["osbuild"] = o.Osbuild.Get()
	}
	if o.Patches != nil {
		toSerialize["patches"] = o.Patches
	}
	if o.Fwpackage.IsSet() {
		toSerialize["fwpackage"] = o.Fwpackage.Get()
	}
	if o.OrigFwPackageUrl.IsSet() {
		toSerialize["origFwPackageUrl"] = o.OrigFwPackageUrl.Get()
	}
	if o.Encrypt.IsSet() {
		toSerialize["encrypt"] = o.Encrypt.Get()
	}
	if o.OverrideWifiMAC.IsSet() {
		toSerialize["overrideWifiMAC"] = o.OverrideWifiMAC.Get()
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if o.Snapshot.IsSet() {
		toSerialize["snapshot"] = o.Snapshot.Get()
	}
	if !IsNil(o.BootOptions) {
		toSerialize["bootOptions"] = o.BootOptions
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableInstanceCreateOptions struct {
	value *InstanceCreateOptions
	isSet bool
}

func (v NullableInstanceCreateOptions) Get() *InstanceCreateOptions {
	return v.value
}

func (v *NullableInstanceCreateOptions) Set(val *InstanceCreateOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceCreateOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceCreateOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceCreateOptions(val *InstanceCreateOptions) *NullableInstanceCreateOptions {
	return &NullableInstanceCreateOptions{value: val, isSet: true}
}

func (v NullableInstanceCreateOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceCreateOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

