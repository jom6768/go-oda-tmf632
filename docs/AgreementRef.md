# AgreementRef

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

### NewAgreementRef

`func NewAgreementRef(type_ string, id string, ) *AgreementRef`

NewAgreementRef instantiates a new AgreementRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementRefWithDefaults

`func NewAgreementRefWithDefaults() *AgreementRef`

NewAgreementRefWithDefaults instantiates a new AgreementRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgreementRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementRef) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *AgreementRef) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *AgreementRef) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *AgreementRef) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *AgreementRef) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *AgreementRef) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *AgreementRef) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *AgreementRef) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *AgreementRef) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *AgreementRef) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AgreementRef) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AgreementRef) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AgreementRef) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AgreementRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementRef) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AgreementRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgreementRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgreementRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgreementRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *AgreementRef) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *AgreementRef) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *AgreementRef) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *AgreementRef) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


