# PartyRefMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | The URI of the referred entity. | [optional] 
**Id** | **string** | The identifier of the referred entity. | 
**Name** | Pointer to **string** | Name of the referred entity. | [optional] 
**ReferredType** | Pointer to **string** | The actual type of the target instance when needed for disambiguation. | [optional] 

## Methods

### NewPartyRefMVO

`func NewPartyRefMVO(type_ string, id string, ) *PartyRefMVO`

NewPartyRefMVO instantiates a new PartyRefMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRefMVOWithDefaults

`func NewPartyRefMVOWithDefaults() *PartyRefMVO`

NewPartyRefMVOWithDefaults instantiates a new PartyRefMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRefMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRefMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRefMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRefMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRefMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRefMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRefMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRefMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRefMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRefMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRefMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyRefMVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRefMVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRefMVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRefMVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyRefMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRefMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRefMVO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyRefMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRefMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRefMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRefMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyRefMVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyRefMVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyRefMVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyRefMVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


