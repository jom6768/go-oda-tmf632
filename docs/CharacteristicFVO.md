# CharacteristicFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Unique identifier of the characteristic | [optional] 
**Name** | **string** | Name of the characteristic | 
**ValueType** | Pointer to **string** | Data type of the value of the characteristic | [optional] 
**CharacteristicRelationship** | Pointer to [**[]CharacteristicRelationshipFVO**](CharacteristicRelationshipFVO.md) |  | [optional] 

## Methods

### NewCharacteristicFVO

`func NewCharacteristicFVO(type_ string, name string, ) *CharacteristicFVO`

NewCharacteristicFVO instantiates a new CharacteristicFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacteristicFVOWithDefaults

`func NewCharacteristicFVOWithDefaults() *CharacteristicFVO`

NewCharacteristicFVOWithDefaults instantiates a new CharacteristicFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CharacteristicFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CharacteristicFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CharacteristicFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *CharacteristicFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *CharacteristicFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *CharacteristicFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *CharacteristicFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *CharacteristicFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *CharacteristicFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *CharacteristicFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *CharacteristicFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *CharacteristicFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CharacteristicFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CharacteristicFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CharacteristicFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CharacteristicFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CharacteristicFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CharacteristicFVO) SetName(v string)`

SetName sets Name field to given value.


### GetValueType

`func (o *CharacteristicFVO) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CharacteristicFVO) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CharacteristicFVO) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *CharacteristicFVO) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetCharacteristicRelationship

`func (o *CharacteristicFVO) GetCharacteristicRelationship() []CharacteristicRelationshipFVO`

GetCharacteristicRelationship returns the CharacteristicRelationship field if non-nil, zero value otherwise.

### GetCharacteristicRelationshipOk

`func (o *CharacteristicFVO) GetCharacteristicRelationshipOk() (*[]CharacteristicRelationshipFVO, bool)`

GetCharacteristicRelationshipOk returns a tuple with the CharacteristicRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicRelationship

`func (o *CharacteristicFVO) SetCharacteristicRelationship(v []CharacteristicRelationshipFVO)`

SetCharacteristicRelationship sets CharacteristicRelationship field to given value.

### HasCharacteristicRelationship

`func (o *CharacteristicFVO) HasCharacteristicRelationship() bool`

HasCharacteristicRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


