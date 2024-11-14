# OrganizationIdentificationMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**IdentificationId** | Pointer to **string** | Identifier | [optional] 
**IssuingAuthority** | Pointer to **string** | Authority which has issued the identifier (chamber of commerce...) | [optional] 
**IssuingDate** | Pointer to **time.Time** | Date at which the identifier was issued | [optional] 
**IdentificationType** | Pointer to **string** | Type of identification information used to identify the company in a country or internationally | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Attachment** | Pointer to [**AttachmentRefOrValueMVO**](AttachmentRefOrValueMVO.md) |  | [optional] 

## Methods

### NewOrganizationIdentificationMVO

`func NewOrganizationIdentificationMVO(type_ string, ) *OrganizationIdentificationMVO`

NewOrganizationIdentificationMVO instantiates a new OrganizationIdentificationMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationIdentificationMVOWithDefaults

`func NewOrganizationIdentificationMVOWithDefaults() *OrganizationIdentificationMVO`

NewOrganizationIdentificationMVOWithDefaults instantiates a new OrganizationIdentificationMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationIdentificationMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationIdentificationMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationIdentificationMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OrganizationIdentificationMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OrganizationIdentificationMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OrganizationIdentificationMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OrganizationIdentificationMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OrganizationIdentificationMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OrganizationIdentificationMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OrganizationIdentificationMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OrganizationIdentificationMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetIdentificationId

`func (o *OrganizationIdentificationMVO) GetIdentificationId() string`

GetIdentificationId returns the IdentificationId field if non-nil, zero value otherwise.

### GetIdentificationIdOk

`func (o *OrganizationIdentificationMVO) GetIdentificationIdOk() (*string, bool)`

GetIdentificationIdOk returns a tuple with the IdentificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationId

`func (o *OrganizationIdentificationMVO) SetIdentificationId(v string)`

SetIdentificationId sets IdentificationId field to given value.

### HasIdentificationId

`func (o *OrganizationIdentificationMVO) HasIdentificationId() bool`

HasIdentificationId returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *OrganizationIdentificationMVO) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *OrganizationIdentificationMVO) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *OrganizationIdentificationMVO) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *OrganizationIdentificationMVO) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetIssuingDate

`func (o *OrganizationIdentificationMVO) GetIssuingDate() time.Time`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *OrganizationIdentificationMVO) GetIssuingDateOk() (*time.Time, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *OrganizationIdentificationMVO) SetIssuingDate(v time.Time)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *OrganizationIdentificationMVO) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetIdentificationType

`func (o *OrganizationIdentificationMVO) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *OrganizationIdentificationMVO) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *OrganizationIdentificationMVO) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *OrganizationIdentificationMVO) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetValidFor

`func (o *OrganizationIdentificationMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *OrganizationIdentificationMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *OrganizationIdentificationMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *OrganizationIdentificationMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetAttachment

`func (o *OrganizationIdentificationMVO) GetAttachment() AttachmentRefOrValueMVO`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *OrganizationIdentificationMVO) GetAttachmentOk() (*AttachmentRefOrValueMVO, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *OrganizationIdentificationMVO) SetAttachment(v AttachmentRefOrValueMVO)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *OrganizationIdentificationMVO) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


