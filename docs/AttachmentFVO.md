# AttachmentFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**Name** | Pointer to **string** | The name of the attachment | [optional] 
**Description** | Pointer to **string** | A narrative text describing the content of the attachment | [optional] 
**Url** | Pointer to **string** | Uniform Resource Locator, is a web page address (a subset of URI) | [optional] 
**Content** | Pointer to **string** | The actual contents of the attachment object, if embedded, encoded as base64 | [optional] 
**Size** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**AttachmentType** | **string** | a business characterization of the purpose of the attachment, for example logo, instructionManual, contractCopy | 
**MimeType** | **string** | a technical characterization of the attachment content format using IETF Mime Types | 

## Methods

### NewAttachmentFVO

`func NewAttachmentFVO(type_ string, attachmentType string, mimeType string, ) *AttachmentFVO`

NewAttachmentFVO instantiates a new AttachmentFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentFVOWithDefaults

`func NewAttachmentFVOWithDefaults() *AttachmentFVO`

NewAttachmentFVOWithDefaults instantiates a new AttachmentFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *AttachmentFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *AttachmentFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *AttachmentFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *AttachmentFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *AttachmentFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *AttachmentFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *AttachmentFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *AttachmentFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *AttachmentFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachmentFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttachmentFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachmentFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AttachmentFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentFVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentFVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AttachmentFVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentFVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentFVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentFVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *AttachmentFVO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AttachmentFVO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AttachmentFVO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AttachmentFVO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContent

`func (o *AttachmentFVO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AttachmentFVO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AttachmentFVO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AttachmentFVO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSize

`func (o *AttachmentFVO) GetSize() Quantity`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AttachmentFVO) GetSizeOk() (*Quantity, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AttachmentFVO) SetSize(v Quantity)`

SetSize sets Size field to given value.

### HasSize

`func (o *AttachmentFVO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetValidFor

`func (o *AttachmentFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *AttachmentFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *AttachmentFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *AttachmentFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetAttachmentType

`func (o *AttachmentFVO) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *AttachmentFVO) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *AttachmentFVO) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.


### GetMimeType

`func (o *AttachmentFVO) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AttachmentFVO) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AttachmentFVO) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


