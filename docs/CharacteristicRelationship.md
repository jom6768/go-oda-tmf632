# CharacteristicRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Unique identifier of the characteristic | [optional] 
**RelationshipType** | Pointer to **string** | The type of relationship | [optional] 

## Methods

### NewCharacteristicRelationship

`func NewCharacteristicRelationship(type_ string, ) *CharacteristicRelationship`

NewCharacteristicRelationship instantiates a new CharacteristicRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacteristicRelationshipWithDefaults

`func NewCharacteristicRelationshipWithDefaults() *CharacteristicRelationship`

NewCharacteristicRelationshipWithDefaults instantiates a new CharacteristicRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CharacteristicRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacteristicRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacteristicRelationship) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *CharacteristicRelationship) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *CharacteristicRelationship) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *CharacteristicRelationship) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *CharacteristicRelationship) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *CharacteristicRelationship) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *CharacteristicRelationship) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *CharacteristicRelationship) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *CharacteristicRelationship) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *CharacteristicRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharacteristicRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharacteristicRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CharacteristicRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *CharacteristicRelationship) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *CharacteristicRelationship) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *CharacteristicRelationship) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *CharacteristicRelationship) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


