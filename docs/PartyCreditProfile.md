# PartyCreditProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**CreditAgencyName** | Pointer to **string** | Name of the credit agency giving the score | [optional] 
**CreditAgencyType** | Pointer to **string** | Type of the credit agency giving the score | [optional] 
**RatingReference** | Pointer to **string** | Reference corresponding to the credit rating | [optional] 
**RatingScore** | Pointer to **int32** | A measure of a party&#39;s creditworthiness calculated on the basis of a combination of factors such as their income and credit history | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyCreditProfile

`func NewPartyCreditProfile(type_ string, ) *PartyCreditProfile`

NewPartyCreditProfile instantiates a new PartyCreditProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyCreditProfileWithDefaults

`func NewPartyCreditProfileWithDefaults() *PartyCreditProfile`

NewPartyCreditProfileWithDefaults instantiates a new PartyCreditProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyCreditProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyCreditProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyCreditProfile) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyCreditProfile) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyCreditProfile) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyCreditProfile) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyCreditProfile) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyCreditProfile) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyCreditProfile) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyCreditProfile) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyCreditProfile) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyCreditProfile) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyCreditProfile) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyCreditProfile) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyCreditProfile) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyCreditProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyCreditProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyCreditProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartyCreditProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreditAgencyName

`func (o *PartyCreditProfile) GetCreditAgencyName() string`

GetCreditAgencyName returns the CreditAgencyName field if non-nil, zero value otherwise.

### GetCreditAgencyNameOk

`func (o *PartyCreditProfile) GetCreditAgencyNameOk() (*string, bool)`

GetCreditAgencyNameOk returns a tuple with the CreditAgencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyName

`func (o *PartyCreditProfile) SetCreditAgencyName(v string)`

SetCreditAgencyName sets CreditAgencyName field to given value.

### HasCreditAgencyName

`func (o *PartyCreditProfile) HasCreditAgencyName() bool`

HasCreditAgencyName returns a boolean if a field has been set.

### GetCreditAgencyType

`func (o *PartyCreditProfile) GetCreditAgencyType() string`

GetCreditAgencyType returns the CreditAgencyType field if non-nil, zero value otherwise.

### GetCreditAgencyTypeOk

`func (o *PartyCreditProfile) GetCreditAgencyTypeOk() (*string, bool)`

GetCreditAgencyTypeOk returns a tuple with the CreditAgencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAgencyType

`func (o *PartyCreditProfile) SetCreditAgencyType(v string)`

SetCreditAgencyType sets CreditAgencyType field to given value.

### HasCreditAgencyType

`func (o *PartyCreditProfile) HasCreditAgencyType() bool`

HasCreditAgencyType returns a boolean if a field has been set.

### GetRatingReference

`func (o *PartyCreditProfile) GetRatingReference() string`

GetRatingReference returns the RatingReference field if non-nil, zero value otherwise.

### GetRatingReferenceOk

`func (o *PartyCreditProfile) GetRatingReferenceOk() (*string, bool)`

GetRatingReferenceOk returns a tuple with the RatingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingReference

`func (o *PartyCreditProfile) SetRatingReference(v string)`

SetRatingReference sets RatingReference field to given value.

### HasRatingReference

`func (o *PartyCreditProfile) HasRatingReference() bool`

HasRatingReference returns a boolean if a field has been set.

### GetRatingScore

`func (o *PartyCreditProfile) GetRatingScore() int32`

GetRatingScore returns the RatingScore field if non-nil, zero value otherwise.

### GetRatingScoreOk

`func (o *PartyCreditProfile) GetRatingScoreOk() (*int32, bool)`

GetRatingScoreOk returns a tuple with the RatingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingScore

`func (o *PartyCreditProfile) SetRatingScore(v int32)`

SetRatingScore sets RatingScore field to given value.

### HasRatingScore

`func (o *PartyCreditProfile) HasRatingScore() bool`

HasRatingScore returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyCreditProfile) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyCreditProfile) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyCreditProfile) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyCreditProfile) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


