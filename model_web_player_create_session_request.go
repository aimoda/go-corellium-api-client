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

// checks if the WebPlayerCreateSessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebPlayerCreateSessionRequest{}

// WebPlayerCreateSessionRequest 
type WebPlayerCreateSessionRequest struct {
	// Project Identifier
	ProjectId string `json:"projectId"`
	// VM Instance Identifier
	InstanceId string `json:"instanceId"`
	// Number of seconds token remains valid
	ExpiresIn float32 `json:"expiresIn"`
	Features Features `json:"features"`
	// Optional string value supplied by client for tracking purposes
	ClientId NullableString `json:"clientId,omitempty"`
}

// NewWebPlayerCreateSessionRequest instantiates a new WebPlayerCreateSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebPlayerCreateSessionRequest(projectId string, instanceId string, expiresIn float32, features Features) *WebPlayerCreateSessionRequest {
	this := WebPlayerCreateSessionRequest{}
	this.ProjectId = projectId
	this.InstanceId = instanceId
	this.ExpiresIn = expiresIn
	this.Features = features
	return &this
}

// NewWebPlayerCreateSessionRequestWithDefaults instantiates a new WebPlayerCreateSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebPlayerCreateSessionRequestWithDefaults() *WebPlayerCreateSessionRequest {
	this := WebPlayerCreateSessionRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *WebPlayerCreateSessionRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *WebPlayerCreateSessionRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *WebPlayerCreateSessionRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetInstanceId returns the InstanceId field value
func (o *WebPlayerCreateSessionRequest) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *WebPlayerCreateSessionRequest) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *WebPlayerCreateSessionRequest) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *WebPlayerCreateSessionRequest) GetExpiresIn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *WebPlayerCreateSessionRequest) GetExpiresInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *WebPlayerCreateSessionRequest) SetExpiresIn(v float32) {
	o.ExpiresIn = v
}

// GetFeatures returns the Features field value
func (o *WebPlayerCreateSessionRequest) GetFeatures() Features {
	if o == nil {
		var ret Features
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *WebPlayerCreateSessionRequest) GetFeaturesOk() (*Features, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Features, true
}

// SetFeatures sets field value
func (o *WebPlayerCreateSessionRequest) SetFeatures(v Features) {
	o.Features = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebPlayerCreateSessionRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPlayerCreateSessionRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *WebPlayerCreateSessionRequest) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *WebPlayerCreateSessionRequest) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *WebPlayerCreateSessionRequest) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *WebPlayerCreateSessionRequest) UnsetClientId() {
	o.ClientId.Unset()
}

func (o WebPlayerCreateSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebPlayerCreateSessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["instanceId"] = o.InstanceId
	toSerialize["expiresIn"] = o.ExpiresIn
	toSerialize["features"] = o.Features
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	return toSerialize, nil
}

type NullableWebPlayerCreateSessionRequest struct {
	value *WebPlayerCreateSessionRequest
	isSet bool
}

func (v NullableWebPlayerCreateSessionRequest) Get() *WebPlayerCreateSessionRequest {
	return v.value
}

func (v *NullableWebPlayerCreateSessionRequest) Set(val *WebPlayerCreateSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebPlayerCreateSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebPlayerCreateSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebPlayerCreateSessionRequest(val *WebPlayerCreateSessionRequest) *NullableWebPlayerCreateSessionRequest {
	return &NullableWebPlayerCreateSessionRequest{value: val, isSet: true}
}

func (v NullableWebPlayerCreateSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebPlayerCreateSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

