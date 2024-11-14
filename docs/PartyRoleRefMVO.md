# PartyRoleRefMVO

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
**PartyId** | Pointer to **string** | The identifier of the engaged party that is linked to the PartyRole object. | [optional] 
**PartyName** | Pointer to **string** | The name of the engaged party that is linked to the PartyRole object. | [optional] 

## Methods

### NewPartyRoleRefMVO

`func NewPartyRoleRefMVO(type_ string, id string, ) *PartyRoleRefMVO`

NewPartyRoleRefMVO instantiates a new PartyRoleRefMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleRefMVOWithDefaults

`func NewPartyRoleRefMVOWithDefaults() *PartyRoleRefMVO`

NewPartyRoleRefMVOWithDefaults instantiates a new PartyRoleRefMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRoleRefMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRoleRefMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRoleRefMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRoleRefMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRoleRefMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRoleRefMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRoleRefMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRoleRefMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRoleRefMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRoleRefMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRoleRefMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyRoleRefMVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRoleRefMVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRoleRefMVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRoleRefMVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyRoleRefMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRoleRefMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRoleRefMVO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyRoleRefMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRoleRefMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRoleRefMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRoleRefMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyRoleRefMVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyRoleRefMVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyRoleRefMVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyRoleRefMVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetPartyId

`func (o *PartyRoleRefMVO) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *PartyRoleRefMVO) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *PartyRoleRefMVO) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *PartyRoleRefMVO) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *PartyRoleRefMVO) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyRoleRefMVO) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyRoleRefMVO) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyRoleRefMVO) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


