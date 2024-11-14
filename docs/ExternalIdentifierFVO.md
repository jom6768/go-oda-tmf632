# ExternalIdentifierFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Owner** | Pointer to **string** | Name of the external system that owns the entity. | [optional] 
**ExternalIdentifierType** | Pointer to **string** | Type of the identification, typically would be the type of the entity within the external system | [optional] 
**Id** | **string** | identification of the entity within the external system. | 

## Methods

### NewExternalIdentifierFVO

`func NewExternalIdentifierFVO(type_ string, id string, ) *ExternalIdentifierFVO`

NewExternalIdentifierFVO instantiates a new ExternalIdentifierFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIdentifierFVOWithDefaults

`func NewExternalIdentifierFVOWithDefaults() *ExternalIdentifierFVO`

NewExternalIdentifierFVOWithDefaults instantiates a new ExternalIdentifierFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalIdentifierFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalIdentifierFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalIdentifierFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *ExternalIdentifierFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *ExternalIdentifierFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *ExternalIdentifierFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *ExternalIdentifierFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *ExternalIdentifierFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *ExternalIdentifierFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *ExternalIdentifierFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *ExternalIdentifierFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetOwner

`func (o *ExternalIdentifierFVO) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExternalIdentifierFVO) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExternalIdentifierFVO) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExternalIdentifierFVO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetExternalIdentifierType

`func (o *ExternalIdentifierFVO) GetExternalIdentifierType() string`

GetExternalIdentifierType returns the ExternalIdentifierType field if non-nil, zero value otherwise.

### GetExternalIdentifierTypeOk

`func (o *ExternalIdentifierFVO) GetExternalIdentifierTypeOk() (*string, bool)`

GetExternalIdentifierTypeOk returns a tuple with the ExternalIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifierType

`func (o *ExternalIdentifierFVO) SetExternalIdentifierType(v string)`

SetExternalIdentifierType sets ExternalIdentifierType field to given value.

### HasExternalIdentifierType

`func (o *ExternalIdentifierFVO) HasExternalIdentifierType() bool`

HasExternalIdentifierType returns a boolean if a field has been set.

### GetId

`func (o *ExternalIdentifierFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIdentifierFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIdentifierFVO) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


