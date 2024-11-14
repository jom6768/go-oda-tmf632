# Party

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifier**](ExternalIdentifier.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]Characteristic**](Characteristic.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificate**](TaxExemptionCertificate.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfile**](PartyCreditProfile.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRole**](RelatedPartyOrPartyRole.md) | List of parties and/or party roles related to this party | [optional] 
**ContactMedium** | Pointer to [**[]ContactMedium**](ContactMedium.md) | List of means for contacting the party, e.g. mobile phone, email address | [optional] 

## Methods

### NewParty

`func NewParty(type_ string, ) *Party`

NewParty instantiates a new Party object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyWithDefaults

`func NewPartyWithDefaults() *Party`

NewPartyWithDefaults instantiates a new Party object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Party) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Party) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Party) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *Party) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *Party) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *Party) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *Party) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *Party) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *Party) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *Party) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *Party) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *Party) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Party) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Party) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Party) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Party) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Party) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Party) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Party) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalReference

`func (o *Party) GetExternalReference() []ExternalIdentifier`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *Party) GetExternalReferenceOk() (*[]ExternalIdentifier, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *Party) SetExternalReference(v []ExternalIdentifier)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *Party) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *Party) GetPartyCharacteristic() []Characteristic`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *Party) GetPartyCharacteristicOk() (*[]Characteristic, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *Party) SetPartyCharacteristic(v []Characteristic)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *Party) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *Party) GetTaxExemptionCertificate() []TaxExemptionCertificate`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *Party) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificate, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *Party) SetTaxExemptionCertificate(v []TaxExemptionCertificate)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *Party) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *Party) GetCreditRating() []PartyCreditProfile`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *Party) GetCreditRatingOk() (*[]PartyCreditProfile, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *Party) SetCreditRating(v []PartyCreditProfile)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *Party) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *Party) GetRelatedParty() []RelatedPartyOrPartyRole`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *Party) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRole, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *Party) SetRelatedParty(v []RelatedPartyOrPartyRole)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *Party) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *Party) GetContactMedium() []ContactMedium`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *Party) GetContactMediumOk() (*[]ContactMedium, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *Party) SetContactMedium(v []ContactMedium)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *Party) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


