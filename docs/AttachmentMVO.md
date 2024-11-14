# AttachmentMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Name** | Pointer to **string** | The name of the attachment | [optional] 
**Description** | Pointer to **string** | A narrative text describing the content of the attachment | [optional] 
**Url** | Pointer to **string** | Uniform Resource Locator, is a web page address (a subset of URI) | [optional] 
**Content** | Pointer to **string** | The actual contents of the attachment object, if embedded, encoded as base64 | [optional] 
**Size** | Pointer to [**Quantity**](Quantity.md) |  | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**AttachmentType** | Pointer to **string** | a business characterization of the purpose of the attachment, for example logo, instructionManual, contractCopy | [optional] 
**MimeType** | Pointer to **string** | a technical characterization of the attachment content format using IETF Mime Types | [optional] 

## Methods

### NewAttachmentMVO

`func NewAttachmentMVO(type_ string, ) *AttachmentMVO`

NewAttachmentMVO instantiates a new AttachmentMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentMVOWithDefaults

`func NewAttachmentMVOWithDefaults() *AttachmentMVO`

NewAttachmentMVOWithDefaults instantiates a new AttachmentMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *AttachmentMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *AttachmentMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *AttachmentMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *AttachmentMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *AttachmentMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *AttachmentMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *AttachmentMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *AttachmentMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetName

`func (o *AttachmentMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttachmentMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AttachmentMVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentMVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentMVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentMVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *AttachmentMVO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AttachmentMVO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AttachmentMVO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AttachmentMVO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContent

`func (o *AttachmentMVO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AttachmentMVO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AttachmentMVO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AttachmentMVO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSize

`func (o *AttachmentMVO) GetSize() Quantity`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AttachmentMVO) GetSizeOk() (*Quantity, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AttachmentMVO) SetSize(v Quantity)`

SetSize sets Size field to given value.

### HasSize

`func (o *AttachmentMVO) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetValidFor

`func (o *AttachmentMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *AttachmentMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *AttachmentMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *AttachmentMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetAttachmentType

`func (o *AttachmentMVO) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *AttachmentMVO) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *AttachmentMVO) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *AttachmentMVO) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### GetMimeType

`func (o *AttachmentMVO) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AttachmentMVO) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AttachmentMVO) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AttachmentMVO) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


