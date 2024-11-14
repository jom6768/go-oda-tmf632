# ContactMediumFVO

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

### NewContactMediumFVO

`func NewContactMediumFVO(type_ string, ) *ContactMediumFVO`

NewContactMediumFVO instantiates a new ContactMediumFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactMediumFVOWithDefaults

`func NewContactMediumFVOWithDefaults() *ContactMediumFVO`

NewContactMediumFVOWithDefaults instantiates a new ContactMediumFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactMediumFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactMediumFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactMediumFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *ContactMediumFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *ContactMediumFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *ContactMediumFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *ContactMediumFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *ContactMediumFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *ContactMediumFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *ContactMediumFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *ContactMediumFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *ContactMediumFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactMediumFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactMediumFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactMediumFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPreferred

`func (o *ContactMediumFVO) GetPreferred() bool`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *ContactMediumFVO) GetPreferredOk() (*bool, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *ContactMediumFVO) SetPreferred(v bool)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *ContactMediumFVO) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetContactType

`func (o *ContactMediumFVO) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactMediumFVO) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactMediumFVO) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ContactMediumFVO) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetValidFor

`func (o *ContactMediumFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *ContactMediumFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *ContactMediumFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *ContactMediumFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


