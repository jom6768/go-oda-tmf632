# Characteristic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Unique identifier of the characteristic | [optional] 
**Name** | Pointer to **string** | Name of the characteristic | [optional] 
**ValueType** | Pointer to **string** | Data type of the value of the characteristic | [optional] 
**CharacteristicRelationship** | Pointer to [**[]CharacteristicRelationship**](CharacteristicRelationship.md) |  | [optional] 

## Methods

### NewCharacteristic

`func NewCharacteristic(type_ string, ) *Characteristic`

NewCharacteristic instantiates a new Characteristic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCharacteristicWithDefaults

`func NewCharacteristicWithDefaults() *Characteristic`

NewCharacteristicWithDefaults instantiates a new Characteristic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Characteristic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Characteristic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Characteristic) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *Characteristic) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *Characteristic) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *Characteristic) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *Characteristic) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *Characteristic) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *Characteristic) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *Characteristic) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *Characteristic) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *Characteristic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Characteristic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Characteristic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Characteristic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Characteristic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Characteristic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Characteristic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Characteristic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueType

`func (o *Characteristic) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *Characteristic) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *Characteristic) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *Characteristic) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetCharacteristicRelationship

`func (o *Characteristic) GetCharacteristicRelationship() []CharacteristicRelationship`

GetCharacteristicRelationship returns the CharacteristicRelationship field if non-nil, zero value otherwise.

### GetCharacteristicRelationshipOk

`func (o *Characteristic) GetCharacteristicRelationshipOk() (*[]CharacteristicRelationship, bool)`

GetCharacteristicRelationshipOk returns a tuple with the CharacteristicRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicRelationship

`func (o *Characteristic) SetCharacteristicRelationship(v []CharacteristicRelationship)`

SetCharacteristicRelationship sets CharacteristicRelationship field to given value.

### HasCharacteristicRelationship

`func (o *Characteristic) HasCharacteristicRelationship() bool`

HasCharacteristicRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


