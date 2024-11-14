# CreditProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**CreditProfileDate** | Pointer to **time.Time** | The date the profile was established | [optional] 
**CreditRiskRating** | Pointer to **int32** | This is an integer whose value is used to rate the risk | [optional] 
**CreditScore** | Pointer to **int32** | A measure of a person or organizations creditworthiness calculated on the basis of a combination of factors such as their income and credit history | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewCreditProfile

`func NewCreditProfile(type_ string, ) *CreditProfile`

NewCreditProfile instantiates a new CreditProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditProfileWithDefaults

`func NewCreditProfileWithDefaults() *CreditProfile`

NewCreditProfileWithDefaults instantiates a new CreditProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreditProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditProfile) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *CreditProfile) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *CreditProfile) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *CreditProfile) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *CreditProfile) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *CreditProfile) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *CreditProfile) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *CreditProfile) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *CreditProfile) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *CreditProfile) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CreditProfile) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CreditProfile) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CreditProfile) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *CreditProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreditProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreditProfileDate

`func (o *CreditProfile) GetCreditProfileDate() time.Time`

GetCreditProfileDate returns the CreditProfileDate field if non-nil, zero value otherwise.

### GetCreditProfileDateOk

`func (o *CreditProfile) GetCreditProfileDateOk() (*time.Time, bool)`

GetCreditProfileDateOk returns a tuple with the CreditProfileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfileDate

`func (o *CreditProfile) SetCreditProfileDate(v time.Time)`

SetCreditProfileDate sets CreditProfileDate field to given value.

### HasCreditProfileDate

`func (o *CreditProfile) HasCreditProfileDate() bool`

HasCreditProfileDate returns a boolean if a field has been set.

### GetCreditRiskRating

`func (o *CreditProfile) GetCreditRiskRating() int32`

GetCreditRiskRating returns the CreditRiskRating field if non-nil, zero value otherwise.

### GetCreditRiskRatingOk

`func (o *CreditProfile) GetCreditRiskRatingOk() (*int32, bool)`

GetCreditRiskRatingOk returns a tuple with the CreditRiskRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRiskRating

`func (o *CreditProfile) SetCreditRiskRating(v int32)`

SetCreditRiskRating sets CreditRiskRating field to given value.

### HasCreditRiskRating

`func (o *CreditProfile) HasCreditRiskRating() bool`

HasCreditRiskRating returns a boolean if a field has been set.

### GetCreditScore

`func (o *CreditProfile) GetCreditScore() int32`

GetCreditScore returns the CreditScore field if non-nil, zero value otherwise.

### GetCreditScoreOk

`func (o *CreditProfile) GetCreditScoreOk() (*int32, bool)`

GetCreditScoreOk returns a tuple with the CreditScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScore

`func (o *CreditProfile) SetCreditScore(v int32)`

SetCreditScore sets CreditScore field to given value.

### HasCreditScore

`func (o *CreditProfile) HasCreditScore() bool`

HasCreditScore returns a boolean if a field has been set.

### GetValidFor

`func (o *CreditProfile) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *CreditProfile) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *CreditProfile) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *CreditProfile) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


