# PartyRoleSpecificationRefFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | **string** | The identifier of the referred entity. | 
**Href** | Pointer to **string** | The URI of the referred entity. | [optional] 
**Name** | Pointer to **string** | Name of the referred entity. | [optional] 
**ReferredType** | Pointer to **string** | The actual type of the target instance when needed for disambiguation. | [optional] 

## Methods

### NewPartyRoleSpecificationRefFVO

`func NewPartyRoleSpecificationRefFVO(type_ string, id string, ) *PartyRoleSpecificationRefFVO`

NewPartyRoleSpecificationRefFVO instantiates a new PartyRoleSpecificationRefFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleSpecificationRefFVOWithDefaults

`func NewPartyRoleSpecificationRefFVOWithDefaults() *PartyRoleSpecificationRefFVO`

NewPartyRoleSpecificationRefFVOWithDefaults instantiates a new PartyRoleSpecificationRefFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRoleSpecificationRefFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRoleSpecificationRefFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRoleSpecificationRefFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRoleSpecificationRefFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRoleSpecificationRefFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRoleSpecificationRefFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRoleSpecificationRefFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRoleSpecificationRefFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRoleSpecificationRefFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRoleSpecificationRefFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRoleSpecificationRefFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *PartyRoleSpecificationRefFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRoleSpecificationRefFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRoleSpecificationRefFVO) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *PartyRoleSpecificationRefFVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRoleSpecificationRefFVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRoleSpecificationRefFVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRoleSpecificationRefFVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *PartyRoleSpecificationRefFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRoleSpecificationRefFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRoleSpecificationRefFVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRoleSpecificationRefFVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyRoleSpecificationRefFVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyRoleSpecificationRefFVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyRoleSpecificationRefFVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyRoleSpecificationRefFVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


