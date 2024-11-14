# ContactMediumMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Identifier for this contact medium. | [optional] 
**Preferred** | Pointer to **bool** | If true, indicates that is the preferred contact medium | [optional] 
**ContactType** | Pointer to **string** | Type of the contact medium to qualifiy it like pro email / personal email. This is not used to define the contact medium used. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewContactMediumMVO

`func NewContactMediumMVO(type_ string, ) *ContactMediumMVO`

NewContactMediumMVO instantiates a new ContactMediumMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactMediumMVOWithDefaults

`func NewContactMediumMVOWithDefaults() *ContactMediumMVO`

NewContactMediumMVOWithDefaults instantiates a new ContactMediumMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactMediumMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactMediumMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactMediumMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *ContactMediumMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *ContactMediumMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *ContactMediumMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *ContactMediumMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *ContactMediumMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *ContactMediumMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *ContactMediumMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *ContactMediumMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *ContactMediumMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactMediumMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactMediumMVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactMediumMVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPreferred

`func (o *ContactMediumMVO) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *ContactMediumMVO) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *ContactMediumMVO) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *ContactMediumMVO) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetContactType

`func (o *ContactMediumMVO) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactMediumMVO) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactMediumMVO) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ContactMediumMVO) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetValidFor

`func (o *ContactMediumMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *ContactMediumMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *ContactMediumMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *ContactMediumMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


