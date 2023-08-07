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

// checks if the ProxyConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfig{}

// ProxyConfig Represents a Proxy configuration object.
type ProxyConfig struct {
	// The device port to use for proxying.
	DevicePort NullableFloat32 `json:"devicePort,omitempty"`
	// If `true`, the first available port will be used if `devicePort` is not available.
	FirstAvailable NullableBool `json:"firstAvailable,omitempty"`
	// If `true`, the proxy will be exposed to the external interface.
	Expose NullableBool `json:"expose,omitempty"`
	// The router port to use for proxying.
	RouterPort NullableFloat32 `json:"routerPort,omitempty"`
}

// NewProxyConfig instantiates a new ProxyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfig() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// NewProxyConfigWithDefaults instantiates a new ProxyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigWithDefaults() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// GetDevicePort returns the DevicePort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfig) GetDevicePort() float32 {
	if o == nil || IsNil(o.DevicePort.Get()) {
		var ret float32
		return ret
	}
	return *o.DevicePort.Get()
}

// GetDevicePortOk returns a tuple with the DevicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfig) GetDevicePortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevicePort.Get(), o.DevicePort.IsSet()
}

// HasDevicePort returns a boolean if a field has been set.
func (o *ProxyConfig) HasDevicePort() bool {
	if o != nil && o.DevicePort.IsSet() {
		return true
	}

	return false
}

// SetDevicePort gets a reference to the given NullableFloat32 and assigns it to the DevicePort field.
func (o *ProxyConfig) SetDevicePort(v float32) {
	o.DevicePort.Set(&v)
}
// SetDevicePortNil sets the value for DevicePort to be an explicit nil
func (o *ProxyConfig) SetDevicePortNil() {
	o.DevicePort.Set(nil)
}

// UnsetDevicePort ensures that no value is present for DevicePort, not even an explicit nil
func (o *ProxyConfig) UnsetDevicePort() {
	o.DevicePort.Unset()
}

// GetFirstAvailable returns the FirstAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfig) GetFirstAvailable() bool {
	if o == nil || IsNil(o.FirstAvailable.Get()) {
		var ret bool
		return ret
	}
	return *o.FirstAvailable.Get()
}

// GetFirstAvailableOk returns a tuple with the FirstAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfig) GetFirstAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstAvailable.Get(), o.FirstAvailable.IsSet()
}

// HasFirstAvailable returns a boolean if a field has been set.
func (o *ProxyConfig) HasFirstAvailable() bool {
	if o != nil && o.FirstAvailable.IsSet() {
		return true
	}

	return false
}

// SetFirstAvailable gets a reference to the given NullableBool and assigns it to the FirstAvailable field.
func (o *ProxyConfig) SetFirstAvailable(v bool) {
	o.FirstAvailable.Set(&v)
}
// SetFirstAvailableNil sets the value for FirstAvailable to be an explicit nil
func (o *ProxyConfig) SetFirstAvailableNil() {
	o.FirstAvailable.Set(nil)
}

// UnsetFirstAvailable ensures that no value is present for FirstAvailable, not even an explicit nil
func (o *ProxyConfig) UnsetFirstAvailable() {
	o.FirstAvailable.Unset()
}

// GetExpose returns the Expose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfig) GetExpose() bool {
	if o == nil || IsNil(o.Expose.Get()) {
		var ret bool
		return ret
	}
	return *o.Expose.Get()
}

// GetExposeOk returns a tuple with the Expose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfig) GetExposeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expose.Get(), o.Expose.IsSet()
}

// HasExpose returns a boolean if a field has been set.
func (o *ProxyConfig) HasExpose() bool {
	if o != nil && o.Expose.IsSet() {
		return true
	}

	return false
}

// SetExpose gets a reference to the given NullableBool and assigns it to the Expose field.
func (o *ProxyConfig) SetExpose(v bool) {
	o.Expose.Set(&v)
}
// SetExposeNil sets the value for Expose to be an explicit nil
func (o *ProxyConfig) SetExposeNil() {
	o.Expose.Set(nil)
}

// UnsetExpose ensures that no value is present for Expose, not even an explicit nil
func (o *ProxyConfig) UnsetExpose() {
	o.Expose.Unset()
}

// GetRouterPort returns the RouterPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxyConfig) GetRouterPort() float32 {
	if o == nil || IsNil(o.RouterPort.Get()) {
		var ret float32
		return ret
	}
	return *o.RouterPort.Get()
}

// GetRouterPortOk returns a tuple with the RouterPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxyConfig) GetRouterPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RouterPort.Get(), o.RouterPort.IsSet()
}

// HasRouterPort returns a boolean if a field has been set.
func (o *ProxyConfig) HasRouterPort() bool {
	if o != nil && o.RouterPort.IsSet() {
		return true
	}

	return false
}

// SetRouterPort gets a reference to the given NullableFloat32 and assigns it to the RouterPort field.
func (o *ProxyConfig) SetRouterPort(v float32) {
	o.RouterPort.Set(&v)
}
// SetRouterPortNil sets the value for RouterPort to be an explicit nil
func (o *ProxyConfig) SetRouterPortNil() {
	o.RouterPort.Set(nil)
}

// UnsetRouterPort ensures that no value is present for RouterPort, not even an explicit nil
func (o *ProxyConfig) UnsetRouterPort() {
	o.RouterPort.Unset()
}

func (o ProxyConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DevicePort.IsSet() {
		toSerialize["devicePort"] = o.DevicePort.Get()
	}
	if o.FirstAvailable.IsSet() {
		toSerialize["firstAvailable"] = o.FirstAvailable.Get()
	}
	if o.Expose.IsSet() {
		toSerialize["expose"] = o.Expose.Get()
	}
	if o.RouterPort.IsSet() {
		toSerialize["routerPort"] = o.RouterPort.Get()
	}
	return toSerialize, nil
}

type NullableProxyConfig struct {
	value *ProxyConfig
	isSet bool
}

func (v NullableProxyConfig) Get() *ProxyConfig {
	return v.value
}

func (v *NullableProxyConfig) Set(val *ProxyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfig(val *ProxyConfig) *NullableProxyConfig {
	return &NullableProxyConfig{value: val, isSet: true}
}

func (v NullableProxyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


