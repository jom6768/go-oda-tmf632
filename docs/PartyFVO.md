# PartyFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifierFVO**](ExternalIdentifierFVO.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]CharacteristicFVO**](CharacteristicFVO.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificateFVO**](TaxExemptionCertificateFVO.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfileFVO**](PartyCreditProfileFVO.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleFVO**](RelatedPartyOrPartyRoleFVO.md) | List of parties and/or party roles related to this party | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumFVO**](ContactMediumFVO.md) | List of means for contacting the party, e.g. mobile phone, email address | [optional] 

## Methods

### NewPartyFVO

`func NewPartyFVO(type_ string, ) *PartyFVO`

NewPartyFVO instantiates a new PartyFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyFVOWithDefaults

`func NewPartyFVOWithDefaults() *PartyFVO`

NewPartyFVOWithDefaults instantiates a new PartyFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *PartyFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartyFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalReference

`func (o *PartyFVO) GetExternalReference() []ExternalIdentifierFVO`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PartyFVO) GetExternalReferenceOk() (*[]ExternalIdentifierFVO, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PartyFVO) SetExternalReference(v []ExternalIdentifierFVO)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PartyFVO) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *PartyFVO) GetPartyCharacteristic() []CharacteristicFVO`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *PartyFVO) GetPartyCharacteristicOk() (*[]CharacteristicFVO, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *PartyFVO) SetPartyCharacteristic(v []CharacteristicFVO)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *PartyFVO) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *PartyFVO) GetTaxExemptionCertificate() []TaxExemptionCertificateFVO`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *PartyFVO) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificateFVO, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *PartyFVO) SetTaxExemptionCertificate(v []TaxExemptionCertificateFVO)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *PartyFVO) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *PartyFVO) GetCreditRating() []PartyCreditProfileFVO`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *PartyFVO) GetCreditRatingOk() (*[]PartyCreditProfileFVO, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *PartyFVO) SetCreditRating(v []PartyCreditProfileFVO)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *PartyFVO) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyFVO) GetRelatedParty() []RelatedPartyOrPartyRoleFVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyFVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleFVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyFVO) SetRelatedParty(v []RelatedPartyOrPartyRoleFVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyFVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyFVO) GetContactMedium() []ContactMediumFVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyFVO) GetContactMediumOk() (*[]ContactMediumFVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyFVO) SetContactMedium(v []ContactMediumFVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyFVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


