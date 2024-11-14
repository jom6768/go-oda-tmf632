# CreditProfileFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**CreditProfileDate** | Pointer to **time.Time** | The date the profile was established | [optional] 
**CreditRiskRating** | Pointer to **int32** | This is an integer whose value is used to rate the risk | [optional] 
**CreditScore** | Pointer to **int32** | A measure of a person or organizations creditworthiness calculated on the basis of a combination of factors such as their income and credit history | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewCreditProfileFVO

`func NewCreditProfileFVO(type_ string, ) *CreditProfileFVO`

NewCreditProfileFVO instantiates a new CreditProfileFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditProfileFVOWithDefaults

`func NewCreditProfileFVOWithDefaults() *CreditProfileFVO`

NewCreditProfileFVOWithDefaults instantiates a new CreditProfileFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreditProfileFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditProfileFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditProfileFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *CreditProfileFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *CreditProfileFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *CreditProfileFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *CreditProfileFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *CreditProfileFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *CreditProfileFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *CreditProfileFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *CreditProfileFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *CreditProfileFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditProfileFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditProfileFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreditProfileFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreditProfileDate

`func (o *CreditProfileFVO) GetCreditProfileDate() time.Time`

GetCreditProfileDate returns the CreditProfileDate field if non-nil, zero value otherwise.

### GetCreditProfileDateOk

`func (o *CreditProfileFVO) GetCreditProfileDateOk() (*time.Time, bool)`

GetCreditProfileDateOk returns a tuple with the CreditProfileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfileDate

`func (o *CreditProfileFVO) SetCreditProfileDate(v time.Time)`

SetCreditProfileDate sets CreditProfileDate field to given value.

### HasCreditProfileDate

`func (o *CreditProfileFVO) HasCreditProfileDate() bool`

HasCreditProfileDate returns a boolean if a field has been set.

### GetCreditRiskRating

`func (o *CreditProfileFVO) GetCreditRiskRating() int32`

GetCreditRiskRating returns the CreditRiskRating field if non-nil, zero value otherwise.

### GetCreditRiskRatingOk

`func (o *CreditProfileFVO) GetCreditRiskRatingOk() (*int32, bool)`

GetCreditRiskRatingOk returns a tuple with the CreditRiskRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRiskRating

`func (o *CreditProfileFVO) SetCreditRiskRating(v int32)`

SetCreditRiskRating sets CreditRiskRating field to given value.

### HasCreditRiskRating

`func (o *CreditProfileFVO) HasCreditRiskRating() bool`

HasCreditRiskRating returns a boolean if a field has been set.

### GetCreditScore

`func (o *CreditProfileFVO) GetCreditScore() int32`

GetCreditScore returns the CreditScore field if non-nil, zero value otherwise.

### GetCreditScoreOk

`func (o *CreditProfileFVO) GetCreditScoreOk() (*int32, bool)`

GetCreditScoreOk returns a tuple with the CreditScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScore

`func (o *CreditProfileFVO) SetCreditScore(v int32)`

SetCreditScore sets CreditScore field to given value.

### HasCreditScore

`func (o *CreditProfileFVO) HasCreditScore() bool`

HasCreditScore returns a boolean if a field has been set.

### GetValidFor

`func (o *CreditProfileFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *CreditProfileFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *CreditProfileFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *CreditProfileFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


