# TaxExemptionCertificateFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Identifier of the tax exemption within list of the exemptions | [optional] 
**TaxDefinition** | Pointer to [**[]TaxDefinitionFVO**](TaxDefinitionFVO.md) | A list of taxes that are covered by the exemption, e.g. City Tax, State Tax. The definition would include the exemption (e.g. for a rate exemption 0% would be a full exemption, 5% could be a partial exemption if the actual rate was 10%). | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**CertificateNumber** | Pointer to **string** | Identifier of a document that shows proof of exemption from taxes for the taxing jurisdiction | [optional] 
**IssuingJurisdiction** | Pointer to **string** | Name of the jurisdiction that issued the exemption | [optional] 
**Reason** | Pointer to **string** | Reason for the tax exemption | [optional] 
**Attachment** | Pointer to [**AttachmentRefOrValueFVO**](AttachmentRefOrValueFVO.md) |  | [optional] 

## Methods

### NewTaxExemptionCertificateFVO

`func NewTaxExemptionCertificateFVO(type_ string, ) *TaxExemptionCertificateFVO`

NewTaxExemptionCertificateFVO instantiates a new TaxExemptionCertificateFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxExemptionCertificateFVOWithDefaults

`func NewTaxExemptionCertificateFVOWithDefaults() *TaxExemptionCertificateFVO`

NewTaxExemptionCertificateFVOWithDefaults instantiates a new TaxExemptionCertificateFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxExemptionCertificateFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxExemptionCertificateFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxExemptionCertificateFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *TaxExemptionCertificateFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TaxExemptionCertificateFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TaxExemptionCertificateFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TaxExemptionCertificateFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *TaxExemptionCertificateFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *TaxExemptionCertificateFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *TaxExemptionCertificateFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *TaxExemptionCertificateFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *TaxExemptionCertificateFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxExemptionCertificateFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxExemptionCertificateFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxExemptionCertificateFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxDefinition

`func (o *TaxExemptionCertificateFVO) GetTaxDefinition() []TaxDefinitionFVO`

GetTaxDefinition returns the TaxDefinition field if non-nil, zero value otherwise.

### GetTaxDefinitionOk

`func (o *TaxExemptionCertificateFVO) GetTaxDefinitionOk() (*[]TaxDefinitionFVO, bool)`

GetTaxDefinitionOk returns a tuple with the TaxDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDefinition

`func (o *TaxExemptionCertificateFVO) SetTaxDefinition(v []TaxDefinitionFVO)`

SetTaxDefinition sets TaxDefinition field to given value.

### HasTaxDefinition

`func (o *TaxExemptionCertificateFVO) HasTaxDefinition() bool`

HasTaxDefinition returns a boolean if a field has been set.

### GetValidFor

`func (o *TaxExemptionCertificateFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *TaxExemptionCertificateFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *TaxExemptionCertificateFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *TaxExemptionCertificateFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *TaxExemptionCertificateFVO) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *TaxExemptionCertificateFVO) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *TaxExemptionCertificateFVO) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *TaxExemptionCertificateFVO) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.

### GetIssuingJurisdiction

`func (o *TaxExemptionCertificateFVO) GetIssuingJurisdiction() string`

GetIssuingJurisdiction returns the IssuingJurisdiction field if non-nil, zero value otherwise.

### GetIssuingJurisdictionOk

`func (o *TaxExemptionCertificateFVO) GetIssuingJurisdictionOk() (*string, bool)`

GetIssuingJurisdictionOk returns a tuple with the IssuingJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingJurisdiction

`func (o *TaxExemptionCertificateFVO) SetIssuingJurisdiction(v string)`

SetIssuingJurisdiction sets IssuingJurisdiction field to given value.

### HasIssuingJurisdiction

`func (o *TaxExemptionCertificateFVO) HasIssuingJurisdiction() bool`

HasIssuingJurisdiction returns a boolean if a field has been set.

### GetReason

`func (o *TaxExemptionCertificateFVO) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxExemptionCertificateFVO) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxExemptionCertificateFVO) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxExemptionCertificateFVO) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAttachment

`func (o *TaxExemptionCertificateFVO) GetAttachment() AttachmentRefOrValueFVO`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *TaxExemptionCertificateFVO) GetAttachmentOk() (*AttachmentRefOrValueFVO, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *TaxExemptionCertificateFVO) SetAttachment(v AttachmentRefOrValueFVO)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *TaxExemptionCertificateFVO) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


