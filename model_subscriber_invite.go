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

// checks if the SubscriberInvite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberInvite{}

// SubscriberInvite Subscriber Invite
type SubscriberInvite struct {
	CouponOptions *CouponOptions `json:"coupon_options,omitempty"`
	Plan *Plan `json:"plan,omitempty"`
	Trial *Trial `json:"trial,omitempty"`
	// Name
	Name NullableString `json:"name,omitempty"`
	// Email
	Email NullableString `json:"email,omitempty"`
	// Coupon code
	CouponCode string `json:"coupon_code"`
	// Domain
	Domain NullableString `json:"domain,omitempty"`
	// Aggregate
	Aggregate string `json:"aggregate"`
	// Used By
	UseBy NullableString `json:"use_by,omitempty"`
	// Is Active flag
	Active bool `json:"active"`
	// Is reusable flag
	Reusable bool `json:"reusable"`
	// Created Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	CreatedAt string `json:"createdAt"`
	// Updated Date in ISO-8601 format e.g. 2022-05-06T02:39:23.000Z
	UpdatedAt string `json:"updatedAt"`
}

// NewSubscriberInvite instantiates a new SubscriberInvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberInvite(couponCode string, aggregate string, active bool, reusable bool, createdAt string, updatedAt string) *SubscriberInvite {
	this := SubscriberInvite{}
	this.CouponCode = couponCode
	this.Aggregate = aggregate
	this.Active = active
	this.Reusable = reusable
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriberInviteWithDefaults instantiates a new SubscriberInvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberInviteWithDefaults() *SubscriberInvite {
	this := SubscriberInvite{}
	return &this
}

// GetCouponOptions returns the CouponOptions field value if set, zero value otherwise.
func (o *SubscriberInvite) GetCouponOptions() CouponOptions {
	if o == nil || IsNil(o.CouponOptions) {
		var ret CouponOptions
		return ret
	}
	return *o.CouponOptions
}

// GetCouponOptionsOk returns a tuple with the CouponOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetCouponOptionsOk() (*CouponOptions, bool) {
	if o == nil || IsNil(o.CouponOptions) {
		return nil, false
	}
	return o.CouponOptions, true
}

// HasCouponOptions returns a boolean if a field has been set.
func (o *SubscriberInvite) HasCouponOptions() bool {
	if o != nil && !IsNil(o.CouponOptions) {
		return true
	}

	return false
}

// SetCouponOptions gets a reference to the given CouponOptions and assigns it to the CouponOptions field.
func (o *SubscriberInvite) SetCouponOptions(v CouponOptions) {
	o.CouponOptions = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SubscriberInvite) GetPlan() Plan {
	if o == nil || IsNil(o.Plan) {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetPlanOk() (*Plan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SubscriberInvite) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *SubscriberInvite) SetPlan(v Plan) {
	o.Plan = &v
}

// GetTrial returns the Trial field value if set, zero value otherwise.
func (o *SubscriberInvite) GetTrial() Trial {
	if o == nil || IsNil(o.Trial) {
		var ret Trial
		return ret
	}
	return *o.Trial
}

// GetTrialOk returns a tuple with the Trial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetTrialOk() (*Trial, bool) {
	if o == nil || IsNil(o.Trial) {
		return nil, false
	}
	return o.Trial, true
}

// HasTrial returns a boolean if a field has been set.
func (o *SubscriberInvite) HasTrial() bool {
	if o != nil && !IsNil(o.Trial) {
		return true
	}

	return false
}

// SetTrial gets a reference to the given Trial and assigns it to the Trial field.
func (o *SubscriberInvite) SetTrial(v Trial) {
	o.Trial = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriberInvite) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriberInvite) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SubscriberInvite) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SubscriberInvite) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SubscriberInvite) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SubscriberInvite) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriberInvite) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriberInvite) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SubscriberInvite) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *SubscriberInvite) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *SubscriberInvite) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *SubscriberInvite) UnsetEmail() {
	o.Email.Unset()
}

// GetCouponCode returns the CouponCode field value
func (o *SubscriberInvite) GetCouponCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetCouponCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CouponCode, true
}

// SetCouponCode sets field value
func (o *SubscriberInvite) SetCouponCode(v string) {
	o.CouponCode = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriberInvite) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriberInvite) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *SubscriberInvite) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *SubscriberInvite) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *SubscriberInvite) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *SubscriberInvite) UnsetDomain() {
	o.Domain.Unset()
}

// GetAggregate returns the Aggregate field value
func (o *SubscriberInvite) GetAggregate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregate
}

// GetAggregateOk returns a tuple with the Aggregate field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetAggregateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregate, true
}

// SetAggregate sets field value
func (o *SubscriberInvite) SetAggregate(v string) {
	o.Aggregate = v
}

// GetUseBy returns the UseBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriberInvite) GetUseBy() string {
	if o == nil || IsNil(o.UseBy.Get()) {
		var ret string
		return ret
	}
	return *o.UseBy.Get()
}

// GetUseByOk returns a tuple with the UseBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriberInvite) GetUseByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseBy.Get(), o.UseBy.IsSet()
}

// HasUseBy returns a boolean if a field has been set.
func (o *SubscriberInvite) HasUseBy() bool {
	if o != nil && o.UseBy.IsSet() {
		return true
	}

	return false
}

// SetUseBy gets a reference to the given NullableString and assigns it to the UseBy field.
func (o *SubscriberInvite) SetUseBy(v string) {
	o.UseBy.Set(&v)
}
// SetUseByNil sets the value for UseBy to be an explicit nil
func (o *SubscriberInvite) SetUseByNil() {
	o.UseBy.Set(nil)
}

// UnsetUseBy ensures that no value is present for UseBy, not even an explicit nil
func (o *SubscriberInvite) UnsetUseBy() {
	o.UseBy.Unset()
}

// GetActive returns the Active field value
func (o *SubscriberInvite) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *SubscriberInvite) SetActive(v bool) {
	o.Active = v
}

// GetReusable returns the Reusable field value
func (o *SubscriberInvite) GetReusable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Reusable
}

// GetReusableOk returns a tuple with the Reusable field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetReusableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reusable, true
}

// SetReusable sets field value
func (o *SubscriberInvite) SetReusable(v bool) {
	o.Reusable = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SubscriberInvite) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SubscriberInvite) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SubscriberInvite) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriberInvite) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SubscriberInvite) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o SubscriberInvite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberInvite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CouponOptions) {
		toSerialize["coupon_options"] = o.CouponOptions
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Trial) {
		toSerialize["trial"] = o.Trial
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["coupon_code"] = o.CouponCode
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	toSerialize["aggregate"] = o.Aggregate
	if o.UseBy.IsSet() {
		toSerialize["use_by"] = o.UseBy.Get()
	}
	toSerialize["active"] = o.Active
	toSerialize["reusable"] = o.Reusable
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableSubscriberInvite struct {
	value *SubscriberInvite
	isSet bool
}

func (v NullableSubscriberInvite) Get() *SubscriberInvite {
	return v.value
}

func (v *NullableSubscriberInvite) Set(val *SubscriberInvite) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberInvite(val *SubscriberInvite) *NullableSubscriberInvite {
	return &NullableSubscriberInvite{value: val, isSet: true}
}

func (v NullableSubscriberInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

