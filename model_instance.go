/*
Corellium API

REST API to manage your virtual devices.

API version: 5.2.0-17376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package corellium

import (
	"encoding/json"
	"time"
)

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance 
type Instance struct {
	// Instance Identifier
	Id NullableString `json:"id,omitempty"`
	// Instance Name
	Name NullableString `json:"name,omitempty"`
	// Key used to encrypt the Instance
	Key NullableString `json:"key,omitempty"`
	// The type of virtual machine this is
	Flavor NullableString `json:"flavor,omitempty"`
	// 
	Type NullableString `json:"type,omitempty"`
	// The projectId of the project this instance belongs to
	Project NullableString `json:"project,omitempty"`
	State *InstanceState `json:"state,omitempty"`
	// Time the state of the instance last changed
	StateChanged NullableTime `json:"stateChanged,omitempty"`
	// Time the instance was started
	StartedAt NullableString `json:"startedAt,omitempty"`
	// Currently executing User Task
	UserTask NullableString `json:"userTask,omitempty"`
	// Current task state if any
	TaskState NullableString `json:"taskState,omitempty"`
	// Current error state if any
	Error NullableString `json:"error,omitempty"`
	BootOptions *InstanceBootOptions `json:"bootOptions,omitempty"`
	// Services IP Address
	ServiceIp NullableString `json:"serviceIp,omitempty"`
	// LAN IP Address
	WifiIp NullableString `json:"wifiIp,omitempty"`
	// Secondary Inteface LAN IP Address (if supported)
	SecondaryIp NullableString `json:"secondaryIp,omitempty"`
	Services *InstanceServices `json:"services,omitempty"`
	// 
	Panicked NullableBool `json:"panicked,omitempty"`
	// Time instance was created
	Created NullableTime `json:"created,omitempty"`
	// Model of virtual machine device
	Model NullableString `json:"model,omitempty"`
	// URL that package used to create this instance is available at
	Fwpackage NullableString `json:"fwpackage,omitempty"`
	// 
	Os NullableString `json:"os,omitempty"`
	Agent NullableInstanceAgentState `json:"agent,omitempty"`
	Netmon *InstanceNetmonState `json:"netmon,omitempty"`
	Netdump *InstanceNetdumpState `json:"netdump,omitempty"`
	// 
	ExposePort NullableString `json:"exposePort,omitempty"`
	// 
	Fault NullableBool `json:"fault,omitempty"`
	// 
	Patches []string `json:"patches,omitempty"`
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance() *Instance {
	this := Instance{}
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Instance) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Instance) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Instance) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Instance) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Instance) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Instance) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Instance) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Instance) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *Instance) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *Instance) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *Instance) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *Instance) UnsetKey() {
	o.Key.Unset()
}

// GetFlavor returns the Flavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetFlavor() string {
	if o == nil || IsNil(o.Flavor.Get()) {
		var ret string
		return ret
	}
	return *o.Flavor.Get()
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flavor.Get(), o.Flavor.IsSet()
}

// HasFlavor returns a boolean if a field has been set.
func (o *Instance) HasFlavor() bool {
	if o != nil && o.Flavor.IsSet() {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given NullableString and assigns it to the Flavor field.
func (o *Instance) SetFlavor(v string) {
	o.Flavor.Set(&v)
}
// SetFlavorNil sets the value for Flavor to be an explicit nil
func (o *Instance) SetFlavorNil() {
	o.Flavor.Set(nil)
}

// UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
func (o *Instance) UnsetFlavor() {
	o.Flavor.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Instance) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Instance) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Instance) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Instance) UnsetType() {
	o.Type.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *Instance) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *Instance) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *Instance) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *Instance) UnsetProject() {
	o.Project.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Instance) GetState() InstanceState {
	if o == nil || IsNil(o.State) {
		var ret InstanceState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetStateOk() (*InstanceState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Instance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given InstanceState and assigns it to the State field.
func (o *Instance) SetState(v InstanceState) {
	o.State = &v
}

// GetStateChanged returns the StateChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetStateChanged() time.Time {
	if o == nil || IsNil(o.StateChanged.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StateChanged.Get()
}

// GetStateChangedOk returns a tuple with the StateChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetStateChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateChanged.Get(), o.StateChanged.IsSet()
}

// HasStateChanged returns a boolean if a field has been set.
func (o *Instance) HasStateChanged() bool {
	if o != nil && o.StateChanged.IsSet() {
		return true
	}

	return false
}

// SetStateChanged gets a reference to the given NullableTime and assigns it to the StateChanged field.
func (o *Instance) SetStateChanged(v time.Time) {
	o.StateChanged.Set(&v)
}
// SetStateChangedNil sets the value for StateChanged to be an explicit nil
func (o *Instance) SetStateChangedNil() {
	o.StateChanged.Set(nil)
}

// UnsetStateChanged ensures that no value is present for StateChanged, not even an explicit nil
func (o *Instance) UnsetStateChanged() {
	o.StateChanged.Unset()
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetStartedAt() string {
	if o == nil || IsNil(o.StartedAt.Get()) {
		var ret string
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Instance) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableString and assigns it to the StartedAt field.
func (o *Instance) SetStartedAt(v string) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *Instance) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *Instance) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetUserTask returns the UserTask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetUserTask() string {
	if o == nil || IsNil(o.UserTask.Get()) {
		var ret string
		return ret
	}
	return *o.UserTask.Get()
}

// GetUserTaskOk returns a tuple with the UserTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetUserTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserTask.Get(), o.UserTask.IsSet()
}

// HasUserTask returns a boolean if a field has been set.
func (o *Instance) HasUserTask() bool {
	if o != nil && o.UserTask.IsSet() {
		return true
	}

	return false
}

// SetUserTask gets a reference to the given NullableString and assigns it to the UserTask field.
func (o *Instance) SetUserTask(v string) {
	o.UserTask.Set(&v)
}
// SetUserTaskNil sets the value for UserTask to be an explicit nil
func (o *Instance) SetUserTaskNil() {
	o.UserTask.Set(nil)
}

// UnsetUserTask ensures that no value is present for UserTask, not even an explicit nil
func (o *Instance) UnsetUserTask() {
	o.UserTask.Unset()
}

// GetTaskState returns the TaskState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetTaskState() string {
	if o == nil || IsNil(o.TaskState.Get()) {
		var ret string
		return ret
	}
	return *o.TaskState.Get()
}

// GetTaskStateOk returns a tuple with the TaskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetTaskStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskState.Get(), o.TaskState.IsSet()
}

// HasTaskState returns a boolean if a field has been set.
func (o *Instance) HasTaskState() bool {
	if o != nil && o.TaskState.IsSet() {
		return true
	}

	return false
}

// SetTaskState gets a reference to the given NullableString and assigns it to the TaskState field.
func (o *Instance) SetTaskState(v string) {
	o.TaskState.Set(&v)
}
// SetTaskStateNil sets the value for TaskState to be an explicit nil
func (o *Instance) SetTaskStateNil() {
	o.TaskState.Set(nil)
}

// UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
func (o *Instance) UnsetTaskState() {
	o.TaskState.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *Instance) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *Instance) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *Instance) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *Instance) UnsetError() {
	o.Error.Unset()
}

// GetBootOptions returns the BootOptions field value if set, zero value otherwise.
func (o *Instance) GetBootOptions() InstanceBootOptions {
	if o == nil || IsNil(o.BootOptions) {
		var ret InstanceBootOptions
		return ret
	}
	return *o.BootOptions
}

// GetBootOptionsOk returns a tuple with the BootOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetBootOptionsOk() (*InstanceBootOptions, bool) {
	if o == nil || IsNil(o.BootOptions) {
		return nil, false
	}
	return o.BootOptions, true
}

// HasBootOptions returns a boolean if a field has been set.
func (o *Instance) HasBootOptions() bool {
	if o != nil && !IsNil(o.BootOptions) {
		return true
	}

	return false
}

// SetBootOptions gets a reference to the given InstanceBootOptions and assigns it to the BootOptions field.
func (o *Instance) SetBootOptions(v InstanceBootOptions) {
	o.BootOptions = &v
}

// GetServiceIp returns the ServiceIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetServiceIp() string {
	if o == nil || IsNil(o.ServiceIp.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceIp.Get()
}

// GetServiceIpOk returns a tuple with the ServiceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetServiceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceIp.Get(), o.ServiceIp.IsSet()
}

// HasServiceIp returns a boolean if a field has been set.
func (o *Instance) HasServiceIp() bool {
	if o != nil && o.ServiceIp.IsSet() {
		return true
	}

	return false
}

// SetServiceIp gets a reference to the given NullableString and assigns it to the ServiceIp field.
func (o *Instance) SetServiceIp(v string) {
	o.ServiceIp.Set(&v)
}
// SetServiceIpNil sets the value for ServiceIp to be an explicit nil
func (o *Instance) SetServiceIpNil() {
	o.ServiceIp.Set(nil)
}

// UnsetServiceIp ensures that no value is present for ServiceIp, not even an explicit nil
func (o *Instance) UnsetServiceIp() {
	o.ServiceIp.Unset()
}

// GetWifiIp returns the WifiIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetWifiIp() string {
	if o == nil || IsNil(o.WifiIp.Get()) {
		var ret string
		return ret
	}
	return *o.WifiIp.Get()
}

// GetWifiIpOk returns a tuple with the WifiIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetWifiIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WifiIp.Get(), o.WifiIp.IsSet()
}

// HasWifiIp returns a boolean if a field has been set.
func (o *Instance) HasWifiIp() bool {
	if o != nil && o.WifiIp.IsSet() {
		return true
	}

	return false
}

// SetWifiIp gets a reference to the given NullableString and assigns it to the WifiIp field.
func (o *Instance) SetWifiIp(v string) {
	o.WifiIp.Set(&v)
}
// SetWifiIpNil sets the value for WifiIp to be an explicit nil
func (o *Instance) SetWifiIpNil() {
	o.WifiIp.Set(nil)
}

// UnsetWifiIp ensures that no value is present for WifiIp, not even an explicit nil
func (o *Instance) UnsetWifiIp() {
	o.WifiIp.Unset()
}

// GetSecondaryIp returns the SecondaryIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetSecondaryIp() string {
	if o == nil || IsNil(o.SecondaryIp.Get()) {
		var ret string
		return ret
	}
	return *o.SecondaryIp.Get()
}

// GetSecondaryIpOk returns a tuple with the SecondaryIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetSecondaryIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryIp.Get(), o.SecondaryIp.IsSet()
}

// HasSecondaryIp returns a boolean if a field has been set.
func (o *Instance) HasSecondaryIp() bool {
	if o != nil && o.SecondaryIp.IsSet() {
		return true
	}

	return false
}

// SetSecondaryIp gets a reference to the given NullableString and assigns it to the SecondaryIp field.
func (o *Instance) SetSecondaryIp(v string) {
	o.SecondaryIp.Set(&v)
}
// SetSecondaryIpNil sets the value for SecondaryIp to be an explicit nil
func (o *Instance) SetSecondaryIpNil() {
	o.SecondaryIp.Set(nil)
}

// UnsetSecondaryIp ensures that no value is present for SecondaryIp, not even an explicit nil
func (o *Instance) UnsetSecondaryIp() {
	o.SecondaryIp.Unset()
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Instance) GetServices() InstanceServices {
	if o == nil || IsNil(o.Services) {
		var ret InstanceServices
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetServicesOk() (*InstanceServices, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Instance) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given InstanceServices and assigns it to the Services field.
func (o *Instance) SetServices(v InstanceServices) {
	o.Services = &v
}

// GetPanicked returns the Panicked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetPanicked() bool {
	if o == nil || IsNil(o.Panicked.Get()) {
		var ret bool
		return ret
	}
	return *o.Panicked.Get()
}

// GetPanickedOk returns a tuple with the Panicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetPanickedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Panicked.Get(), o.Panicked.IsSet()
}

// HasPanicked returns a boolean if a field has been set.
func (o *Instance) HasPanicked() bool {
	if o != nil && o.Panicked.IsSet() {
		return true
	}

	return false
}

// SetPanicked gets a reference to the given NullableBool and assigns it to the Panicked field.
func (o *Instance) SetPanicked(v bool) {
	o.Panicked.Set(&v)
}
// SetPanickedNil sets the value for Panicked to be an explicit nil
func (o *Instance) SetPanickedNil() {
	o.Panicked.Set(nil)
}

// UnsetPanicked ensures that no value is present for Panicked, not even an explicit nil
func (o *Instance) UnsetPanicked() {
	o.Panicked.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Instance) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Instance) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Instance) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Instance) UnsetCreated() {
	o.Created.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *Instance) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *Instance) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *Instance) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *Instance) UnsetModel() {
	o.Model.Unset()
}

// GetFwpackage returns the Fwpackage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetFwpackage() string {
	if o == nil || IsNil(o.Fwpackage.Get()) {
		var ret string
		return ret
	}
	return *o.Fwpackage.Get()
}

// GetFwpackageOk returns a tuple with the Fwpackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetFwpackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fwpackage.Get(), o.Fwpackage.IsSet()
}

// HasFwpackage returns a boolean if a field has been set.
func (o *Instance) HasFwpackage() bool {
	if o != nil && o.Fwpackage.IsSet() {
		return true
	}

	return false
}

// SetFwpackage gets a reference to the given NullableString and assigns it to the Fwpackage field.
func (o *Instance) SetFwpackage(v string) {
	o.Fwpackage.Set(&v)
}
// SetFwpackageNil sets the value for Fwpackage to be an explicit nil
func (o *Instance) SetFwpackageNil() {
	o.Fwpackage.Set(nil)
}

// UnsetFwpackage ensures that no value is present for Fwpackage, not even an explicit nil
func (o *Instance) UnsetFwpackage() {
	o.Fwpackage.Unset()
}

// GetOs returns the Os field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetOs() string {
	if o == nil || IsNil(o.Os.Get()) {
		var ret string
		return ret
	}
	return *o.Os.Get()
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Os.Get(), o.Os.IsSet()
}

// HasOs returns a boolean if a field has been set.
func (o *Instance) HasOs() bool {
	if o != nil && o.Os.IsSet() {
		return true
	}

	return false
}

// SetOs gets a reference to the given NullableString and assigns it to the Os field.
func (o *Instance) SetOs(v string) {
	o.Os.Set(&v)
}
// SetOsNil sets the value for Os to be an explicit nil
func (o *Instance) SetOsNil() {
	o.Os.Set(nil)
}

// UnsetOs ensures that no value is present for Os, not even an explicit nil
func (o *Instance) UnsetOs() {
	o.Os.Unset()
}

// GetAgent returns the Agent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetAgent() InstanceAgentState {
	if o == nil || IsNil(o.Agent.Get()) {
		var ret InstanceAgentState
		return ret
	}
	return *o.Agent.Get()
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetAgentOk() (*InstanceAgentState, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agent.Get(), o.Agent.IsSet()
}

// HasAgent returns a boolean if a field has been set.
func (o *Instance) HasAgent() bool {
	if o != nil && o.Agent.IsSet() {
		return true
	}

	return false
}

// SetAgent gets a reference to the given NullableInstanceAgentState and assigns it to the Agent field.
func (o *Instance) SetAgent(v InstanceAgentState) {
	o.Agent.Set(&v)
}
// SetAgentNil sets the value for Agent to be an explicit nil
func (o *Instance) SetAgentNil() {
	o.Agent.Set(nil)
}

// UnsetAgent ensures that no value is present for Agent, not even an explicit nil
func (o *Instance) UnsetAgent() {
	o.Agent.Unset()
}

// GetNetmon returns the Netmon field value if set, zero value otherwise.
func (o *Instance) GetNetmon() InstanceNetmonState {
	if o == nil || IsNil(o.Netmon) {
		var ret InstanceNetmonState
		return ret
	}
	return *o.Netmon
}

// GetNetmonOk returns a tuple with the Netmon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetNetmonOk() (*InstanceNetmonState, bool) {
	if o == nil || IsNil(o.Netmon) {
		return nil, false
	}
	return o.Netmon, true
}

// HasNetmon returns a boolean if a field has been set.
func (o *Instance) HasNetmon() bool {
	if o != nil && !IsNil(o.Netmon) {
		return true
	}

	return false
}

// SetNetmon gets a reference to the given InstanceNetmonState and assigns it to the Netmon field.
func (o *Instance) SetNetmon(v InstanceNetmonState) {
	o.Netmon = &v
}

// GetNetdump returns the Netdump field value if set, zero value otherwise.
func (o *Instance) GetNetdump() InstanceNetdumpState {
	if o == nil || IsNil(o.Netdump) {
		var ret InstanceNetdumpState
		return ret
	}
	return *o.Netdump
}

// GetNetdumpOk returns a tuple with the Netdump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetNetdumpOk() (*InstanceNetdumpState, bool) {
	if o == nil || IsNil(o.Netdump) {
		return nil, false
	}
	return o.Netdump, true
}

// HasNetdump returns a boolean if a field has been set.
func (o *Instance) HasNetdump() bool {
	if o != nil && !IsNil(o.Netdump) {
		return true
	}

	return false
}

// SetNetdump gets a reference to the given InstanceNetdumpState and assigns it to the Netdump field.
func (o *Instance) SetNetdump(v InstanceNetdumpState) {
	o.Netdump = &v
}

// GetExposePort returns the ExposePort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetExposePort() string {
	if o == nil || IsNil(o.ExposePort.Get()) {
		var ret string
		return ret
	}
	return *o.ExposePort.Get()
}

// GetExposePortOk returns a tuple with the ExposePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetExposePortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExposePort.Get(), o.ExposePort.IsSet()
}

// HasExposePort returns a boolean if a field has been set.
func (o *Instance) HasExposePort() bool {
	if o != nil && o.ExposePort.IsSet() {
		return true
	}

	return false
}

// SetExposePort gets a reference to the given NullableString and assigns it to the ExposePort field.
func (o *Instance) SetExposePort(v string) {
	o.ExposePort.Set(&v)
}
// SetExposePortNil sets the value for ExposePort to be an explicit nil
func (o *Instance) SetExposePortNil() {
	o.ExposePort.Set(nil)
}

// UnsetExposePort ensures that no value is present for ExposePort, not even an explicit nil
func (o *Instance) UnsetExposePort() {
	o.ExposePort.Unset()
}

// GetFault returns the Fault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetFault() bool {
	if o == nil || IsNil(o.Fault.Get()) {
		var ret bool
		return ret
	}
	return *o.Fault.Get()
}

// GetFaultOk returns a tuple with the Fault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetFaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fault.Get(), o.Fault.IsSet()
}

// HasFault returns a boolean if a field has been set.
func (o *Instance) HasFault() bool {
	if o != nil && o.Fault.IsSet() {
		return true
	}

	return false
}

// SetFault gets a reference to the given NullableBool and assigns it to the Fault field.
func (o *Instance) SetFault(v bool) {
	o.Fault.Set(&v)
}
// SetFaultNil sets the value for Fault to be an explicit nil
func (o *Instance) SetFaultNil() {
	o.Fault.Set(nil)
}

// UnsetFault ensures that no value is present for Fault, not even an explicit nil
func (o *Instance) UnsetFault() {
	o.Fault.Unset()
}

// GetPatches returns the Patches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetPatches() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Patches
}

// GetPatchesOk returns a tuple with the Patches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetPatchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Patches) {
		return nil, false
	}
	return o.Patches, true
}

// HasPatches returns a boolean if a field has been set.
func (o *Instance) HasPatches() bool {
	if o != nil && IsNil(o.Patches) {
		return true
	}

	return false
}

// SetPatches gets a reference to the given []string and assigns it to the Patches field.
func (o *Instance) SetPatches(v []string) {
	o.Patches = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Instance) GetCreatedBy() CreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret CreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetCreatedByOk() (*CreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Instance) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given CreatedBy and assigns it to the CreatedBy field.
func (o *Instance) SetCreatedBy(v CreatedBy) {
	o.CreatedBy = &v
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Flavor.IsSet() {
		toSerialize["flavor"] = o.Flavor.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.StateChanged.IsSet() {
		toSerialize["stateChanged"] = o.StateChanged.Get()
	}
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	if o.UserTask.IsSet() {
		toSerialize["userTask"] = o.UserTask.Get()
	}
	if o.TaskState.IsSet() {
		toSerialize["taskState"] = o.TaskState.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !IsNil(o.BootOptions) {
		toSerialize["bootOptions"] = o.BootOptions
	}
	if o.ServiceIp.IsSet() {
		toSerialize["serviceIp"] = o.ServiceIp.Get()
	}
	if o.WifiIp.IsSet() {
		toSerialize["wifiIp"] = o.WifiIp.Get()
	}
	if o.SecondaryIp.IsSet() {
		toSerialize["secondaryIp"] = o.SecondaryIp.Get()
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if o.Panicked.IsSet() {
		toSerialize["panicked"] = o.Panicked.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Fwpackage.IsSet() {
		toSerialize["fwpackage"] = o.Fwpackage.Get()
	}
	if o.Os.IsSet() {
		toSerialize["os"] = o.Os.Get()
	}
	if o.Agent.IsSet() {
		toSerialize["agent"] = o.Agent.Get()
	}
	if !IsNil(o.Netmon) {
		toSerialize["netmon"] = o.Netmon
	}
	if !IsNil(o.Netdump) {
		toSerialize["netdump"] = o.Netdump
	}
	if o.ExposePort.IsSet() {
		toSerialize["exposePort"] = o.ExposePort.Get()
	}
	if o.Fault.IsSet() {
		toSerialize["fault"] = o.Fault.Get()
	}
	if o.Patches != nil {
		toSerialize["patches"] = o.Patches
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	return toSerialize, nil
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


