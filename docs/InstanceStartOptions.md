# InstanceStartOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paused** | Pointer to **NullableBool** | Start device paused | [optional] 
**Sockcap** | Pointer to **NullableBool** | Start device with Process Network Monitoring enabled | [optional] 

## Methods

### NewInstanceStartOptions

`func NewInstanceStartOptions() *InstanceStartOptions`

NewInstanceStartOptions instantiates a new InstanceStartOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStartOptionsWithDefaults

`func NewInstanceStartOptionsWithDefaults() *InstanceStartOptions`

NewInstanceStartOptionsWithDefaults instantiates a new InstanceStartOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaused

`func (o *InstanceStartOptions) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *InstanceStartOptions) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *InstanceStartOptions) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *InstanceStartOptions) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### SetPausedNil

`func (o *InstanceStartOptions) SetPausedNil(b bool)`

 SetPausedNil sets the value for Paused to be an explicit nil

### UnsetPaused
`func (o *InstanceStartOptions) UnsetPaused()`

UnsetPaused ensures that no value is present for Paused, not even an explicit nil
### GetSockcap

`func (o *InstanceStartOptions) GetSockcap() bool`

GetSockcap returns the Sockcap field if non-nil, zero value otherwise.

### GetSockcapOk

`func (o *InstanceStartOptions) GetSockcapOk() (*bool, bool)`

GetSockcapOk returns a tuple with the Sockcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockcap

`func (o *InstanceStartOptions) SetSockcap(v bool)`

SetSockcap sets Sockcap field to given value.

### HasSockcap

`func (o *InstanceStartOptions) HasSockcap() bool`

HasSockcap returns a boolean if a field has been set.

### SetSockcapNil

`func (o *InstanceStartOptions) SetSockcapNil(b bool)`

 SetSockcapNil sets the value for Sockcap to be an explicit nil

### UnsetSockcap
`func (o *InstanceStartOptions) UnsetSockcap()`

UnsetSockcap ensures that no value is present for Sockcap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


