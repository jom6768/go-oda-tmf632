# PartyCreditProfileMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**CreditAgencyName** | Pointer to **string** | Name of the credit agency giving the score | [optional] 
**CreditAgencyType** | Pointer to **string** | Type of the credit agency giving the score | [optional] 
**RatingReference** | Pointer to **string** | Reference corresponding to the credit rating | [optional] 
**RatingScore** | Pointer to **int32** | A measure of a party&#39;s creditworthiness calculated on the basis of a combination of factors such as their income and credit history | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyCreditProfileMVO

`func NewPartyCreditProfileMVO(type_ string, ) *PartyCreditProfileMVO`

NewPartyCreditProfileMVO instantiates a new PartyCreditProfileMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyCreditProfileMVOWithDefaults

`func NewPartyCreditProfileMVOWithDefaults() *PartyCreditProfileMVO`

NewPartyCreditProfileMVOWithDefaults instantiates a new PartyCreditProfileMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyCreditProfileMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyCreditProfileMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyCreditProfileMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyCreditProfileMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyCreditProfileMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyCreditProfileMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyCreditProfileMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyCreditProfileMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyCreditProfileMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyCreditProfileMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyCreditProfileMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetCreditAgencyName

`func (o *PartyCreditProfileMVO) GetCreditAgencyName() string`

GetCreditAgencyName returns the CreditAgencyName field if non-nil, zero value otherwise.

### GetCreditAgencyNameOk

`func (o *PartyCreditProfileMVO) GetCreditAgencyNameOk() (*string, bool)`

GetCreditAgencyNameOk returns a tuple with the CreditAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyName

`func (o *PartyCreditProfileMVO) SetCreditAgencyName(v string)`

SetCreditAgencyName sets CreditAgencyName field to given value.

### HasCreditAgencyName

`func (o *PartyCreditProfileMVO) HasCreditAgencyName() bool`

HasCreditAgencyName returns a boolean if a field has been set.

### GetCreditAgencyType

`func (o *PartyCreditProfileMVO) GetCreditAgencyType() string`

GetCreditAgencyType returns the CreditAgencyType field if non-nil, zero value otherwise.

### GetCreditAgencyTypeOk

`func (o *PartyCreditProfileMVO) GetCreditAgencyTypeOk() (*string, bool)`

GetCreditAgencyTypeOk returns a tuple with the CreditAgencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyType

`func (o *PartyCreditProfileMVO) SetCreditAgencyType(v string)`

SetCreditAgencyType sets CreditAgencyType field to given value.

### HasCreditAgencyType

`func (o *PartyCreditProfileMVO) HasCreditAgencyType() bool`

HasCreditAgencyType returns a boolean if a field has been set.

### GetRatingReference

`func (o *PartyCreditProfileMVO) GetRatingReference() string`

GetRatingReference returns the RatingReference field if non-nil, zero value otherwise.

### GetRatingReferenceOk

`func (o *PartyCreditProfileMVO) GetRatingReferenceOk() (*string, bool)`

GetRatingReferenceOk returns a tuple with the RatingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingReference

`func (o *PartyCreditProfileMVO) SetRatingReference(v string)`

SetRatingReference sets RatingReference field to given value.

### HasRatingReference

`func (o *PartyCreditProfileMVO) HasRatingReference() bool`

HasRatingReference returns a boolean if a field has been set.

### GetRatingScore

`func (o *PartyCreditProfileMVO) GetRatingScore() int32`

GetRatingScore returns the RatingScore field if non-nil, zero value otherwise.

### GetRatingScoreOk

`func (o *PartyCreditProfileMVO) GetRatingScoreOk() (*int32, bool)`

GetRatingScoreOk returns a tuple with the RatingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingScore

`func (o *PartyCreditProfileMVO) SetRatingScore(v int32)`

SetRatingScore sets RatingScore field to given value.

### HasRatingScore

`func (o *PartyCreditProfileMVO) HasRatingScore() bool`

HasRatingScore returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyCreditProfileMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyCreditProfileMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyCreditProfileMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyCreditProfileMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


