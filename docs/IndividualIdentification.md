# IndividualIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**IdentificationId** | Pointer to **string** | Identifier | [optional] 
**IssuingAuthority** | Pointer to **string** | Authority which has issued the identifier, such as: social security, town hall | [optional] 
**IssuingDate** | Pointer to **time.Time** | Date at which the identifier was issued | [optional] 
**IdentificationType** | Pointer to **string** | Identification type (passport, national identity card, drivers license, social security number, birth certificate) | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Attachment** | Pointer to [**AttachmentRefOrValue**](AttachmentRefOrValue.md) |  | [optional] 

## Methods

### NewIndividualIdentification

`func NewIndividualIdentification(type_ string, ) *IndividualIdentification`

NewIndividualIdentification instantiates a new IndividualIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualIdentificationWithDefaults

`func NewIndividualIdentificationWithDefaults() *IndividualIdentification`

NewIndividualIdentificationWithDefaults instantiates a new IndividualIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IndividualIdentification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IndividualIdentification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IndividualIdentification) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *IndividualIdentification) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *IndividualIdentification) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *IndividualIdentification) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *IndividualIdentification) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *IndividualIdentification) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *IndividualIdentification) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *IndividualIdentification) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *IndividualIdentification) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetIdentificationId

`func (o *IndividualIdentification) GetIdentificationId() string`

GetIdentificationId returns the IdentificationId field if non-nil, zero value otherwise.

### GetIdentificationIdOk

`func (o *IndividualIdentification) GetIdentificationIdOk() (*string, bool)`

GetIdentificationIdOk returns a tuple with the IdentificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationId

`func (o *IndividualIdentification) SetIdentificationId(v string)`

SetIdentificationId sets IdentificationId field to given value.

### HasIdentificationId

`func (o *IndividualIdentification) HasIdentificationId() bool`

HasIdentificationId returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *IndividualIdentification) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *IndividualIdentification) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *IndividualIdentification) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *IndividualIdentification) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetIssuingDate

`func (o *IndividualIdentification) GetIssuingDate() time.Time`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *IndividualIdentification) GetIssuingDateOk() (*time.Time, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *IndividualIdentification) SetIssuingDate(v time.Time)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *IndividualIdentification) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetIdentificationType

`func (o *IndividualIdentification) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *IndividualIdentification) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *IndividualIdentification) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *IndividualIdentification) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetValidFor

`func (o *IndividualIdentification) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *IndividualIdentification) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *IndividualIdentification) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *IndividualIdentification) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetAttachment

`func (o *IndividualIdentification) GetAttachment() AttachmentRefOrValue`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *IndividualIdentification) GetAttachmentOk() (*AttachmentRefOrValue, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *IndividualIdentification) SetAttachment(v AttachmentRefOrValue)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *IndividualIdentification) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


