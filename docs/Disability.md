# Disability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisabilityCode** | Pointer to **string** | Code of the disability | [optional] 
**DisabilityName** | Pointer to **string** | Name of the disability | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewDisability

`func NewDisability() *Disability`

NewDisability instantiates a new Disability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisabilityWithDefaults

`func NewDisabilityWithDefaults() *Disability`

NewDisabilityWithDefaults instantiates a new Disability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabilityCode

`func (o *Disability) GetDisabilityCode() string`

GetDisabilityCode returns the DisabilityCode field if non-nil, zero value otherwise.

### GetDisabilityCodeOk

`func (o *Disability) GetDisabilityCodeOk() (*string, bool)`

GetDisabilityCodeOk returns a tuple with the DisabilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityCode

`func (o *Disability) SetDisabilityCode(v string)`

SetDisabilityCode sets DisabilityCode field to given value.

### HasDisabilityCode

`func (o *Disability) HasDisabilityCode() bool`

HasDisabilityCode returns a boolean if a field has been set.

### GetDisabilityName

`func (o *Disability) GetDisabilityName() string`

GetDisabilityName returns the DisabilityName field if non-nil, zero value otherwise.

### GetDisabilityNameOk

`func (o *Disability) GetDisabilityNameOk() (*string, bool)`

GetDisabilityNameOk returns a tuple with the DisabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabilityName

`func (o *Disability) SetDisabilityName(v string)`

SetDisabilityName sets DisabilityName field to given value.

### HasDisabilityName

`func (o *Disability) HasDisabilityName() bool`

HasDisabilityName returns a boolean if a field has been set.

### GetValidFor

`func (o *Disability) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *Disability) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *Disability) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *Disability) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


