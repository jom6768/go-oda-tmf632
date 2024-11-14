# PartyCreditProfileFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**CreditAgencyName** | Pointer to **string** | Name of the credit agency giving the score | [optional] 
**CreditAgencyType** | Pointer to **string** | Type of the credit agency giving the score | [optional] 
**RatingReference** | Pointer to **string** | Reference corresponding to the credit rating | [optional] 
**RatingScore** | Pointer to **int32** | A measure of a party&#39;s creditworthiness calculated on the basis of a combination of factors such as their income and credit history | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyCreditProfileFVO

`func NewPartyCreditProfileFVO(type_ string, ) *PartyCreditProfileFVO`

NewPartyCreditProfileFVO instantiates a new PartyCreditProfileFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyCreditProfileFVOWithDefaults

`func NewPartyCreditProfileFVOWithDefaults() *PartyCreditProfileFVO`

NewPartyCreditProfileFVOWithDefaults instantiates a new PartyCreditProfileFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyCreditProfileFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyCreditProfileFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyCreditProfileFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyCreditProfileFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyCreditProfileFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyCreditProfileFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyCreditProfileFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyCreditProfileFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyCreditProfileFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyCreditProfileFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyCreditProfileFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *PartyCreditProfileFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyCreditProfileFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyCreditProfileFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartyCreditProfileFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreditAgencyName

`func (o *PartyCreditProfileFVO) GetCreditAgencyName() string`

GetCreditAgencyName returns the CreditAgencyName field if non-nil, zero value otherwise.

### GetCreditAgencyNameOk

`func (o *PartyCreditProfileFVO) GetCreditAgencyNameOk() (*string, bool)`

GetCreditAgencyNameOk returns a tuple with the CreditAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyName

`func (o *PartyCreditProfileFVO) SetCreditAgencyName(v string)`

SetCreditAgencyName sets CreditAgencyName field to given value.

### HasCreditAgencyName

`func (o *PartyCreditProfileFVO) HasCreditAgencyName() bool`

HasCreditAgencyName returns a boolean if a field has been set.

### GetCreditAgencyType

`func (o *PartyCreditProfileFVO) GetCreditAgencyType() string`

GetCreditAgencyType returns the CreditAgencyType field if non-nil, zero value otherwise.

### GetCreditAgencyTypeOk

`func (o *PartyCreditProfileFVO) GetCreditAgencyTypeOk() (*string, bool)`

GetCreditAgencyTypeOk returns a tuple with the CreditAgencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyType

`func (o *PartyCreditProfileFVO) SetCreditAgencyType(v string)`

SetCreditAgencyType sets CreditAgencyType field to given value.

### HasCreditAgencyType

`func (o *PartyCreditProfileFVO) HasCreditAgencyType() bool`

HasCreditAgencyType returns a boolean if a field has been set.

### GetRatingReference

`func (o *PartyCreditProfileFVO) GetRatingReference() string`

GetRatingReference returns the RatingReference field if non-nil, zero value otherwise.

### GetRatingReferenceOk

`func (o *PartyCreditProfileFVO) GetRatingReferenceOk() (*string, bool)`

GetRatingReferenceOk returns a tuple with the RatingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingReference

`func (o *PartyCreditProfileFVO) SetRatingReference(v string)`

SetRatingReference sets RatingReference field to given value.

### HasRatingReference

`func (o *PartyCreditProfileFVO) HasRatingReference() bool`

HasRatingReference returns a boolean if a field has been set.

### GetRatingScore

`func (o *PartyCreditProfileFVO) GetRatingScore() int32`

GetRatingScore returns the RatingScore field if non-nil, zero value otherwise.

### GetRatingScoreOk

`func (o *PartyCreditProfileFVO) GetRatingScoreOk() (*int32, bool)`

GetRatingScoreOk returns a tuple with the RatingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingScore

`func (o *PartyCreditProfileFVO) SetRatingScore(v int32)`

SetRatingScore sets RatingScore field to given value.

### HasRatingScore

`func (o *PartyCreditProfileFVO) HasRatingScore() bool`

HasRatingScore returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyCreditProfileFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyCreditProfileFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyCreditProfileFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyCreditProfileFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


