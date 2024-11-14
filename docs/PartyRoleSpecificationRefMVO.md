# PartyRoleSpecificationRefMVO

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

### NewPartyRoleSpecificationRefMVO

`func NewPartyRoleSpecificationRefMVO(type_ string, id string, ) *PartyRoleSpecificationRefMVO`

NewPartyRoleSpecificationRefMVO instantiates a new PartyRoleSpecificationRefMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleSpecificationRefMVOWithDefaults

`func NewPartyRoleSpecificationRefMVOWithDefaults() *PartyRoleSpecificationRefMVO`

NewPartyRoleSpecificationRefMVOWithDefaults instantiates a new PartyRoleSpecificationRefMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRoleSpecificationRefMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRoleSpecificationRefMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRoleSpecificationRefMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRoleSpecificationRefMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRoleSpecificationRefMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRoleSpecificationRefMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRoleSpecificationRefMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRoleSpecificationRefMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRoleSpecificationRefMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRoleSpecificationRefMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRoleSpecificationRefMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyRoleSpecificationRefMVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRoleSpecificationRefMVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRoleSpecificationRefMVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRoleSpecificationRefMVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyRoleSpecificationRefMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRoleSpecificationRefMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRoleSpecificationRefMVO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyRoleSpecificationRefMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRoleSpecificationRefMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRoleSpecificationRefMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRoleSpecificationRefMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyRoleSpecificationRefMVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyRoleSpecificationRefMVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyRoleSpecificationRefMVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyRoleSpecificationRefMVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


