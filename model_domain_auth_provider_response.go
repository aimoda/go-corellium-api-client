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

// checks if the DomainAuthProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainAuthProviderResponse{}

// DomainAuthProviderResponse 
type DomainAuthProviderResponse struct {
	// Provider ID
	Identifier string `json:"identifier"`
	// Provider ID for backward compatibility with frontend
	Id NullableString `json:"id,omitempty"`
	// Provider Type
	ProviderType string `json:"providerType"`
	// Provider Type for backward compatibility with frontend
	Provider NullableString `json:"provider,omitempty"`
	// Login Button Text
	Label string `json:"label"`
	// Title Text for the Frontend's Provider Configuration Form
	Name NullableString `json:"name,omitempty"`
	// Coordinator endpoint for Authorizing with this provider
	AuthorizationUrl NullableString `json:"authorizationUrl,omitempty"`
	// True: Configured in coordinator.json, False: Custom Domain Provider
	Default bool `json:"default"`
	// Enabled/Disabled
	Enabled bool `json:"enabled"`
	Config *OpenIDConfig `json:"config,omitempty"`
	// Optional User ID of creator
	CreatedBy NullableString `json:"createdBy,omitempty"`
	// Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	CreatedAt string `json:"createdAt"`
	// Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	UpdatedAt string `json:"updatedAt"`
}

// NewDomainAuthProviderResponse instantiates a new DomainAuthProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainAuthProviderResponse(identifier string, providerType string, label string, default_ bool, enabled bool, createdAt string, updatedAt string) *DomainAuthProviderResponse {
	this := DomainAuthProviderResponse{}
	this.Identifier = identifier
	this.ProviderType = providerType
	this.Label = label
	this.Default = default_
	this.Enabled = enabled
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDomainAuthProviderResponseWithDefaults instantiates a new DomainAuthProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainAuthProviderResponseWithDefaults() *DomainAuthProviderResponse {
	this := DomainAuthProviderResponse{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *DomainAuthProviderResponse) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *DomainAuthProviderResponse) SetIdentifier(v string) {
	o.Identifier = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DomainAuthProviderResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DomainAuthProviderResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DomainAuthProviderResponse) UnsetId() {
	o.Id.Unset()
}

// GetProviderType returns the ProviderType field value
func (o *DomainAuthProviderResponse) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *DomainAuthProviderResponse) SetProviderType(v string) {
	o.ProviderType = v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderResponse) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderResponse) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *DomainAuthProviderResponse) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *DomainAuthProviderResponse) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *DomainAuthProviderResponse) UnsetProvider() {
	o.Provider.Unset()
}

// GetLabel returns the Label field value
func (o *DomainAuthProviderResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DomainAuthProviderResponse) SetLabel(v string) {
	o.Label = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DomainAuthProviderResponse) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DomainAuthProviderResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DomainAuthProviderResponse) UnsetName() {
	o.Name.Unset()
}

// GetAuthorizationUrl returns the AuthorizationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderResponse) GetAuthorizationUrl() string {
	if o == nil || IsNil(o.AuthorizationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorizationUrl.Get()
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderResponse) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationUrl.Get(), o.AuthorizationUrl.IsSet()
}

// HasAuthorizationUrl returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasAuthorizationUrl() bool {
	if o != nil && o.AuthorizationUrl.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationUrl gets a reference to the given NullableString and assigns it to the AuthorizationUrl field.
func (o *DomainAuthProviderResponse) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl.Set(&v)
}
// SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil
func (o *DomainAuthProviderResponse) SetAuthorizationUrlNil() {
	o.AuthorizationUrl.Set(nil)
}

// UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
func (o *DomainAuthProviderResponse) UnsetAuthorizationUrl() {
	o.AuthorizationUrl.Unset()
}

// GetDefault returns the Default field value
func (o *DomainAuthProviderResponse) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *DomainAuthProviderResponse) SetDefault(v bool) {
	o.Default = v
}

// GetEnabled returns the Enabled field value
func (o *DomainAuthProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DomainAuthProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *DomainAuthProviderResponse) GetConfig() OpenIDConfig {
	if o == nil || IsNil(o.Config) {
		var ret OpenIDConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetConfigOk() (*OpenIDConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OpenIDConfig and assigns it to the Config field.
func (o *DomainAuthProviderResponse) SetConfig(v OpenIDConfig) {
	o.Config = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainAuthProviderResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainAuthProviderResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DomainAuthProviderResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *DomainAuthProviderResponse) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *DomainAuthProviderResponse) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *DomainAuthProviderResponse) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *DomainAuthProviderResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DomainAuthProviderResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DomainAuthProviderResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DomainAuthProviderResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DomainAuthProviderResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o DomainAuthProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainAuthProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["providerType"] = o.ProviderType
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	toSerialize["label"] = o.Label
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AuthorizationUrl.IsSet() {
		toSerialize["authorizationUrl"] = o.AuthorizationUrl.Get()
	}
	toSerialize["default"] = o.Default
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableDomainAuthProviderResponse struct {
	value *DomainAuthProviderResponse
	isSet bool
}

func (v NullableDomainAuthProviderResponse) Get() *DomainAuthProviderResponse {
	return v.value
}

func (v *NullableDomainAuthProviderResponse) Set(val *DomainAuthProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainAuthProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainAuthProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainAuthProviderResponse(val *DomainAuthProviderResponse) *NullableDomainAuthProviderResponse {
	return &NullableDomainAuthProviderResponse{value: val, isSet: true}
}

func (v NullableDomainAuthProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainAuthProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


