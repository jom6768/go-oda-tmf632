# CharacteristicMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Unique identifier of the characteristic | [optional] 
**Name** | Pointer to **string** | Name of the characteristic | [optional] 
**ValueType** | Pointer to **string** | Data type of the value of the characteristic | [optional] 
**CharacteristicRelationship** | Pointer to [**[]CharacteristicRelationshipMVO**](CharacteristicRelationshipMVO.md) |  | [optional] 

## Methods

### NewCharacteristicMVO

`func NewCharacteristicMVO(type_ string, ) *CharacteristicMVO`

NewCharacteristicMVO instantiates a new CharacteristicMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacteristicMVOWithDefaults

`func NewCharacteristicMVOWithDefaults() *CharacteristicMVO`

NewCharacteristicMVOWithDefaults instantiates a new CharacteristicMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CharacteristicMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacteristicMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacteristicMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *CharacteristicMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *CharacteristicMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *CharacteristicMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *CharacteristicMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *CharacteristicMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *CharacteristicMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *CharacteristicMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *CharacteristicMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *CharacteristicMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharacteristicMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharacteristicMVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CharacteristicMVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CharacteristicMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacteristicMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacteristicMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CharacteristicMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueType

`func (o *CharacteristicMVO) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CharacteristicMVO) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CharacteristicMVO) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CharacteristicMVO) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetCharacteristicRelationship

`func (o *CharacteristicMVO) GetCharacteristicRelationship() []CharacteristicRelationshipMVO`

GetCharacteristicRelationship returns the CharacteristicRelationship field if non-nil, zero value otherwise.

### GetCharacteristicRelationshipOk

`func (o *CharacteristicMVO) GetCharacteristicRelationshipOk() (*[]CharacteristicRelationshipMVO, bool)`

GetCharacteristicRelationshipOk returns a tuple with the CharacteristicRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicRelationship

`func (o *CharacteristicMVO) SetCharacteristicRelationship(v []CharacteristicRelationshipMVO)`

SetCharacteristicRelationship sets CharacteristicRelationship field to given value.

### HasCharacteristicRelationship

`func (o *CharacteristicMVO) HasCharacteristicRelationship() bool`

HasCharacteristicRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


