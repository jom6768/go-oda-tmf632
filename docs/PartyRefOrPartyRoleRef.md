# PartyRefOrPartyRoleRef

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

### NewPartyRefOrPartyRoleRef

`func NewPartyRefOrPartyRoleRef(type_ string, id string, ) *PartyRefOrPartyRoleRef`

NewPartyRefOrPartyRoleRef instantiates a new PartyRefOrPartyRoleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRefOrPartyRoleRefWithDefaults

`func NewPartyRefOrPartyRoleRefWithDefaults() *PartyRefOrPartyRoleRef`

NewPartyRefOrPartyRoleRefWithDefaults instantiates a new PartyRefOrPartyRoleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRefOrPartyRoleRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRefOrPartyRoleRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRefOrPartyRoleRef) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRefOrPartyRoleRef) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRefOrPartyRoleRef) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRefOrPartyRoleRef) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRefOrPartyRoleRef) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRefOrPartyRoleRef) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRefOrPartyRoleRef) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRefOrPartyRoleRef) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRefOrPartyRoleRef) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyRefOrPartyRoleRef) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRefOrPartyRoleRef) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRefOrPartyRoleRef) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRefOrPartyRoleRef) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyRefOrPartyRoleRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRefOrPartyRoleRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRefOrPartyRoleRef) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyRefOrPartyRoleRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRefOrPartyRoleRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRefOrPartyRoleRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRefOrPartyRoleRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyRefOrPartyRoleRef) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyRefOrPartyRoleRef) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyRefOrPartyRoleRef) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyRefOrPartyRoleRef) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetPartyId

`func (o *PartyRefOrPartyRoleRef) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *PartyRefOrPartyRoleRef) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *PartyRefOrPartyRoleRef) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *PartyRefOrPartyRoleRef) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *PartyRefOrPartyRoleRef) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyRefOrPartyRoleRef) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyRefOrPartyRoleRef) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyRefOrPartyRoleRef) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


