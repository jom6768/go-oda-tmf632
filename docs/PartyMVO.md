# PartyMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifierMVO**](ExternalIdentifierMVO.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]CharacteristicMVO**](CharacteristicMVO.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificateMVO**](TaxExemptionCertificateMVO.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfileMVO**](PartyCreditProfileMVO.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleMVO**](RelatedPartyOrPartyRoleMVO.md) | List of parties and/or party roles related to this party | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumMVO**](ContactMediumMVO.md) | List of means for contacting the party, e.g. mobile phone, email address | [optional] 

## Methods

### NewPartyMVO

`func NewPartyMVO(type_ string, ) *PartyMVO`

NewPartyMVO instantiates a new PartyMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyMVOWithDefaults

`func NewPartyMVOWithDefaults() *PartyMVO`

NewPartyMVOWithDefaults instantiates a new PartyMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetExternalReference

`func (o *PartyMVO) GetExternalReference() []ExternalIdentifierMVO`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PartyMVO) GetExternalReferenceOk() (*[]ExternalIdentifierMVO, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PartyMVO) SetExternalReference(v []ExternalIdentifierMVO)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PartyMVO) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *PartyMVO) GetPartyCharacteristic() []CharacteristicMVO`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *PartyMVO) GetPartyCharacteristicOk() (*[]CharacteristicMVO, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *PartyMVO) SetPartyCharacteristic(v []CharacteristicMVO)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *PartyMVO) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *PartyMVO) GetTaxExemptionCertificate() []TaxExemptionCertificateMVO`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *PartyMVO) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificateMVO, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *PartyMVO) SetTaxExemptionCertificate(v []TaxExemptionCertificateMVO)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *PartyMVO) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *PartyMVO) GetCreditRating() []PartyCreditProfileMVO`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *PartyMVO) GetCreditRatingOk() (*[]PartyCreditProfileMVO, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *PartyMVO) SetCreditRating(v []PartyCreditProfileMVO)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *PartyMVO) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyMVO) GetRelatedParty() []RelatedPartyOrPartyRoleMVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyMVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleMVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyMVO) SetRelatedParty(v []RelatedPartyOrPartyRoleMVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyMVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyMVO) GetContactMedium() []ContactMediumMVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyMVO) GetContactMediumOk() (*[]ContactMediumMVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyMVO) SetContactMedium(v []ContactMediumMVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyMVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


