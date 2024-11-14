# AttachmentRef

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
**Description** | Pointer to **string** | A narrative text describing the content of the attachment | [optional] 
**Url** | Pointer to **string** | Link to the attachment media/content | [optional] 

## Methods

### NewAttachmentRef

`func NewAttachmentRef(type_ string, id string, ) *AttachmentRef`

NewAttachmentRef instantiates a new AttachmentRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentRefWithDefaults

`func NewAttachmentRefWithDefaults() *AttachmentRef`

NewAttachmentRefWithDefaults instantiates a new AttachmentRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentRef) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *AttachmentRef) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *AttachmentRef) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *AttachmentRef) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *AttachmentRef) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *AttachmentRef) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *AttachmentRef) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *AttachmentRef) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *AttachmentRef) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *AttachmentRef) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AttachmentRef) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AttachmentRef) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AttachmentRef) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *AttachmentRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentRef) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AttachmentRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *AttachmentRef) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *AttachmentRef) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *AttachmentRef) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *AttachmentRef) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetDescription

`func (o *AttachmentRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentRef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentRef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *AttachmentRef) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AttachmentRef) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AttachmentRef) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AttachmentRef) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

