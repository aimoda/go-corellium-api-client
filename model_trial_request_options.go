/*
Corellium API

REST API to manage your virtual devices.

API version: 5.0.0-17089
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
)

// checks if the TrialRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrialRequestOptions{}

// TrialRequestOptions 
type TrialRequestOptions struct {
	Address *Address `json:"address,omitempty"`
	// The user's email address.
	Email NullableString `json:"email,omitempty"`
	// The user's full name.
	Name NullableString `json:"name,omitempty"`
	Metadata *TrialRequestMetadata `json:"metadata,omitempty"`
	// If true, create an enterprise domain.
	Enterprise NullableBool `json:"enterprise,omitempty"`
	// Stripe payment token.
	Token NullableString `json:"token,omitempty"`
}

// NewTrialRequestOptions instantiates a new TrialRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialRequestOptions() *TrialRequestOptions {
	this := TrialRequestOptions{}
	return &this
}

// NewTrialRequestOptionsWithDefaults instantiates a new TrialRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialRequestOptionsWithDefaults() *TrialRequestOptions {
	this := TrialRequestOptions{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TrialRequestOptions) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialRequestOptions) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *TrialRequestOptions) SetAddress(v Address) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestOptions) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestOptions) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *TrialRequestOptions) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *TrialRequestOptions) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *TrialRequestOptions) UnsetEmail() {
	o.Email.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestOptions) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TrialRequestOptions) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TrialRequestOptions) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TrialRequestOptions) UnsetName() {
	o.Name.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TrialRequestOptions) GetMetadata() TrialRequestMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret TrialRequestMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialRequestOptions) GetMetadataOk() (*TrialRequestMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given TrialRequestMetadata and assigns it to the Metadata field.
func (o *TrialRequestOptions) SetMetadata(v TrialRequestMetadata) {
	o.Metadata = &v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestOptions) GetEnterprise() bool {
	if o == nil || IsNil(o.Enterprise.Get()) {
		var ret bool
		return ret
	}
	return *o.Enterprise.Get()
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestOptions) GetEnterpriseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enterprise.Get(), o.Enterprise.IsSet()
}

// HasEnterprise returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasEnterprise() bool {
	if o != nil && o.Enterprise.IsSet() {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given NullableBool and assigns it to the Enterprise field.
func (o *TrialRequestOptions) SetEnterprise(v bool) {
	o.Enterprise.Set(&v)
}
// SetEnterpriseNil sets the value for Enterprise to be an explicit nil
func (o *TrialRequestOptions) SetEnterpriseNil() {
	o.Enterprise.Set(nil)
}

// UnsetEnterprise ensures that no value is present for Enterprise, not even an explicit nil
func (o *TrialRequestOptions) UnsetEnterprise() {
	o.Enterprise.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrialRequestOptions) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrialRequestOptions) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *TrialRequestOptions) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *TrialRequestOptions) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *TrialRequestOptions) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *TrialRequestOptions) UnsetToken() {
	o.Token.Unset()
}

func (o TrialRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrialRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Enterprise.IsSet() {
		toSerialize["enterprise"] = o.Enterprise.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableTrialRequestOptions struct {
	value *TrialRequestOptions
	isSet bool
}

func (v NullableTrialRequestOptions) Get() *TrialRequestOptions {
	return v.value
}

func (v *NullableTrialRequestOptions) Set(val *TrialRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialRequestOptions(val *TrialRequestOptions) *NullableTrialRequestOptions {
	return &NullableTrialRequestOptions{value: val, isSet: true}
}

func (v NullableTrialRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


