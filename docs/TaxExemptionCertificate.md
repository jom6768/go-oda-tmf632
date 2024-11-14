# TaxExemptionCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Identifier of the tax exemption within list of the exemptions | [optional] 
**TaxDefinition** | Pointer to [**[]TaxDefinition**](TaxDefinition.md) | A list of taxes that are covered by the exemption, e.g. City Tax, State Tax. The definition would include the exemption (e.g. for a rate exemption 0% would be a full exemption, 5% could be a partial exemption if the actual rate was 10%). | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**CertificateNumber** | Pointer to **string** | Identifier of a document that shows proof of exemption from taxes for the taxing jurisdiction | [optional] 
**IssuingJurisdiction** | Pointer to **string** | Name of the jurisdiction that issued the exemption | [optional] 
**Reason** | Pointer to **string** | Reason for the tax exemption | [optional] 
**Attachment** | Pointer to [**AttachmentRefOrValue**](AttachmentRefOrValue.md) |  | [optional] 

## Methods

### NewTaxExemptionCertificate

`func NewTaxExemptionCertificate(type_ string, ) *TaxExemptionCertificate`

NewTaxExemptionCertificate instantiates a new TaxExemptionCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxExemptionCertificateWithDefaults

`func NewTaxExemptionCertificateWithDefaults() *TaxExemptionCertificate`

NewTaxExemptionCertificateWithDefaults instantiates a new TaxExemptionCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxExemptionCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxExemptionCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxExemptionCertificate) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *TaxExemptionCertificate) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TaxExemptionCertificate) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TaxExemptionCertificate) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TaxExemptionCertificate) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *TaxExemptionCertificate) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *TaxExemptionCertificate) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *TaxExemptionCertificate) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *TaxExemptionCertificate) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *TaxExemptionCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxExemptionCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxExemptionCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxExemptionCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxDefinition

`func (o *TaxExemptionCertificate) GetTaxDefinition() []TaxDefinition`

GetTaxDefinition returns the TaxDefinition field if non-nil, zero value otherwise.

### GetTaxDefinitionOk

`func (o *TaxExemptionCertificate) GetTaxDefinitionOk() (*[]TaxDefinition, bool)`

GetTaxDefinitionOk returns a tuple with the TaxDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDefinition

`func (o *TaxExemptionCertificate) SetTaxDefinition(v []TaxDefinition)`

SetTaxDefinition sets TaxDefinition field to given value.

### HasTaxDefinition

`func (o *TaxExemptionCertificate) HasTaxDefinition() bool`

HasTaxDefinition returns a boolean if a field has been set.

### GetValidFor

`func (o *TaxExemptionCertificate) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *TaxExemptionCertificate) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *TaxExemptionCertificate) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *TaxExemptionCertificate) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *TaxExemptionCertificate) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *TaxExemptionCertificate) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *TaxExemptionCertificate) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *TaxExemptionCertificate) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.

### GetIssuingJurisdiction

`func (o *TaxExemptionCertificate) GetIssuingJurisdiction() string`

GetIssuingJurisdiction returns the IssuingJurisdiction field if non-nil, zero value otherwise.

### GetIssuingJurisdictionOk

`func (o *TaxExemptionCertificate) GetIssuingJurisdictionOk() (*string, bool)`

GetIssuingJurisdictionOk returns a tuple with the IssuingJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingJurisdiction

`func (o *TaxExemptionCertificate) SetIssuingJurisdiction(v string)`

SetIssuingJurisdiction sets IssuingJurisdiction field to given value.

### HasIssuingJurisdiction

`func (o *TaxExemptionCertificate) HasIssuingJurisdiction() bool`

HasIssuingJurisdiction returns a boolean if a field has been set.

### GetReason

`func (o *TaxExemptionCertificate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxExemptionCertificate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxExemptionCertificate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxExemptionCertificate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAttachment

`func (o *TaxExemptionCertificate) GetAttachment() AttachmentRefOrValue`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TaxExemptionCertificate) GetAttachmentOk() (*AttachmentRefOrValue, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TaxExemptionCertificate) SetAttachment(v AttachmentRefOrValue)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TaxExemptionCertificate) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


