# OrganizationIdentification

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
**Attachment** | Pointer to [**AttachmentRefOrValue**](AttachmentRefOrValue.md) |  | [optional] 

## Methods

### NewOrganizationIdentification

`func NewOrganizationIdentification(type_ string, ) *OrganizationIdentification`

NewOrganizationIdentification instantiates a new OrganizationIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationIdentificationWithDefaults

`func NewOrganizationIdentificationWithDefaults() *OrganizationIdentification`

NewOrganizationIdentificationWithDefaults instantiates a new OrganizationIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationIdentification) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OrganizationIdentification) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OrganizationIdentification) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OrganizationIdentification) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OrganizationIdentification) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OrganizationIdentification) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OrganizationIdentification) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OrganizationIdentification) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OrganizationIdentification) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetIdentificationId

`func (o *OrganizationIdentification) GetIdentificationId() string`

GetIdentificationId returns the IdentificationId field if non-nil, zero value otherwise.

### GetIdentificationIdOk

`func (o *OrganizationIdentification) GetIdentificationIdOk() (*string, bool)`

GetIdentificationIdOk returns a tuple with the IdentificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationId

`func (o *OrganizationIdentification) SetIdentificationId(v string)`

SetIdentificationId sets IdentificationId field to given value.

### HasIdentificationId

`func (o *OrganizationIdentification) HasIdentificationId() bool`

HasIdentificationId returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *OrganizationIdentification) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *OrganizationIdentification) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *OrganizationIdentification) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *OrganizationIdentification) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetIssuingDate

`func (o *OrganizationIdentification) GetIssuingDate() time.Time`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *OrganizationIdentification) GetIssuingDateOk() (*time.Time, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *OrganizationIdentification) SetIssuingDate(v time.Time)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *OrganizationIdentification) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetIdentificationType

`func (o *OrganizationIdentification) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *OrganizationIdentification) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *OrganizationIdentification) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *OrganizationIdentification) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetValidFor

`func (o *OrganizationIdentification) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *OrganizationIdentification) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *OrganizationIdentification) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *OrganizationIdentification) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetAttachment

`func (o *OrganizationIdentification) GetAttachment() AttachmentRefOrValue`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *OrganizationIdentification) GetAttachmentOk() (*AttachmentRefOrValue, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *OrganizationIdentification) SetAttachment(v AttachmentRefOrValue)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *OrganizationIdentification) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


