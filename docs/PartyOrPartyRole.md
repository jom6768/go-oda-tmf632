# PartyOrPartyRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | **string** | unique identifier | 
**Name** | Pointer to **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | [optional] 
**ReferredType** | Pointer to **string** | The actual type of the target instance when needed for disambiguation. | [optional] 
**PartyId** | Pointer to **string** | The identifier of the engaged party that is linked to the PartyRole object. | [optional] 
**PartyName** | Pointer to **string** | The name of the engaged party that is linked to the PartyRole object. | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifier**](ExternalIdentifier.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]Characteristic**](Characteristic.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificate**](TaxExemptionCertificate.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfile**](PartyCreditProfile.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRole**](RelatedPartyOrPartyRole.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMedium**](ContactMedium.md) |  | [optional] 
**Gender** | Pointer to **string** | Gender | [optional] 
**PlaceOfBirth** | Pointer to **string** | Reference to the place where the individual was born | [optional] 
**CountryOfBirth** | Pointer to **string** | Country where the individual was born | [optional] 
**Nationality** | Pointer to **string** | Nationality | [optional] 
**MaritalStatus** | Pointer to **string** | Marital status (married, divorced, widow ...) | [optional] 
**BirthDate** | Pointer to **time.Time** | Birth date | [optional] 
**DeathDate** | Pointer to **time.Time** | Date of death | [optional] 
**Title** | Pointer to **string** | Useful for titles (aristocratic, social,...) Pr, Dr, Sir, ... | [optional] 
**AristocraticTitle** | Pointer to **string** | e.g. Baron, Graf, Earl | [optional] 
**Generation** | Pointer to **string** | e.g.. Sr, Jr, III (the third) | [optional] 
**PreferredGivenName** | Pointer to **string** | Contains the chosen name by which the individual prefers to be addressed. Note: This name may be a name other than a given name, such as a nickname | [optional] 
**FamilyNamePrefix** | Pointer to **string** | Family name prefix | [optional] 
**LegalName** | Pointer to **string** | Legal name or birth name (name one has for official purposes) | [optional] 
**MiddleName** | Pointer to **string** | Middles name or initial | [optional] 
**FormattedName** | Pointer to **string** | A fully formatted name in one string with all of its pieces in their proper place and all of the necessary punctuation. Useful for specific contexts (Chinese, Japanese, Korean) | [optional] 
**Location** | Pointer to **string** | Temporary current location of the individual (may be used if the individual has approved its sharing) | [optional] 
**Status** | Pointer to **string** | Used to track the lifecycle status of the party role. | [optional] 
**OtherName** | Pointer to [**[]OtherNameOrganization**](OtherNameOrganization.md) | List of additional names by which the organization is known | [optional] 
**IndividualIdentification** | Pointer to [**[]IndividualIdentification**](IndividualIdentification.md) | List of official identifications issued to the individual, such as passport, driving licence, social security number | [optional] 
**Disability** | Pointer to [**[]Disability**](Disability.md) | List of disabilities suffered by the individual | [optional] 
**LanguageAbility** | Pointer to [**[]LanguageAbility**](LanguageAbility.md) | List of national languages known by the individual | [optional] 
**Skill** | Pointer to [**[]Skill**](Skill.md) | List of skills exhibited by the individual | [optional] 
**FamilyName** | Pointer to **string** | Contains the non-chosen or inherited name. Also known as last name in the Western context | [optional] 
**GivenName** | Pointer to **string** | First name of the individual | [optional] 
**IsLegalEntity** | Pointer to **bool** | If value is true, the organization is a legal entity known by a national referential. | [optional] 
**IsHeadOffice** | Pointer to **bool** | If value is true, the organization is the head office | [optional] 
**OrganizationType** | Pointer to **string** | Type of Organization (company, department...) | [optional] 
**ExistsDuring** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**NameType** | Pointer to **string** | Type of the name : Co, Inc, Ltd, etc. | [optional] 
**OrganizationIdentification** | Pointer to [**[]OrganizationIdentification**](OrganizationIdentification.md) | List of official identifiers given to the organization, for example company number in the registry of companies | [optional] 
**OrganizationChildRelationship** | Pointer to [**[]OrganizationChildRelationship**](OrganizationChildRelationship.md) | List of organizations that are contained within this organization. For example if this organization is the Legal Department, the child organizations might include Claims, Courts, Contracts | [optional] 
**OrganizationParentRelationship** | Pointer to [**OrganizationParentRelationship**](OrganizationParentRelationship.md) |  | [optional] 
**TradingName** | Pointer to **string** | Name that the organization (unit) trades under | [optional] 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | Pointer to [**PartyRef**](PartyRef.md) |  | [optional] 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRef**](PartyRoleSpecificationRef.md) |  | [optional] 
**Characteristic** | Pointer to [**[]Characteristic**](Characteristic.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRef**](AccountRef.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRef**](AgreementRef.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRef**](PaymentMethodRef.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfile**](CreditProfile.md) |  | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyOrPartyRole

`func NewPartyOrPartyRole(type_ string, id string, ) *PartyOrPartyRole`

NewPartyOrPartyRole instantiates a new PartyOrPartyRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyOrPartyRoleWithDefaults

`func NewPartyOrPartyRoleWithDefaults() *PartyOrPartyRole`

NewPartyOrPartyRoleWithDefaults instantiates a new PartyOrPartyRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyOrPartyRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyOrPartyRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyOrPartyRole) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyOrPartyRole) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyOrPartyRole) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyOrPartyRole) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyOrPartyRole) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyOrPartyRole) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyOrPartyRole) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyOrPartyRole) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyOrPartyRole) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyOrPartyRole) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyOrPartyRole) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyOrPartyRole) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyOrPartyRole) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyOrPartyRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyOrPartyRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyOrPartyRole) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyOrPartyRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyOrPartyRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyOrPartyRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyOrPartyRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferredType

`func (o *PartyOrPartyRole) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyOrPartyRole) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyOrPartyRole) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyOrPartyRole) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetPartyId

`func (o *PartyOrPartyRole) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *PartyOrPartyRole) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *PartyOrPartyRole) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *PartyOrPartyRole) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *PartyOrPartyRole) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyOrPartyRole) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyOrPartyRole) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyOrPartyRole) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetExternalReference

`func (o *PartyOrPartyRole) GetExternalReference() []ExternalIdentifier`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PartyOrPartyRole) GetExternalReferenceOk() (*[]ExternalIdentifier, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PartyOrPartyRole) SetExternalReference(v []ExternalIdentifier)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PartyOrPartyRole) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *PartyOrPartyRole) GetPartyCharacteristic() []Characteristic`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *PartyOrPartyRole) GetPartyCharacteristicOk() (*[]Characteristic, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *PartyOrPartyRole) SetPartyCharacteristic(v []Characteristic)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *PartyOrPartyRole) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *PartyOrPartyRole) GetTaxExemptionCertificate() []TaxExemptionCertificate`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *PartyOrPartyRole) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificate, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *PartyOrPartyRole) SetTaxExemptionCertificate(v []TaxExemptionCertificate)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *PartyOrPartyRole) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *PartyOrPartyRole) GetCreditRating() []PartyCreditProfile`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *PartyOrPartyRole) GetCreditRatingOk() (*[]PartyCreditProfile, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *PartyOrPartyRole) SetCreditRating(v []PartyCreditProfile)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *PartyOrPartyRole) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyOrPartyRole) GetRelatedParty() []RelatedPartyOrPartyRole`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyOrPartyRole) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRole, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyOrPartyRole) SetRelatedParty(v []RelatedPartyOrPartyRole)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyOrPartyRole) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyOrPartyRole) GetContactMedium() []ContactMedium`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyOrPartyRole) GetContactMediumOk() (*[]ContactMedium, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyOrPartyRole) SetContactMedium(v []ContactMedium)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyOrPartyRole) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetGender

`func (o *PartyOrPartyRole) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PartyOrPartyRole) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PartyOrPartyRole) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PartyOrPartyRole) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *PartyOrPartyRole) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PartyOrPartyRole) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PartyOrPartyRole) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PartyOrPartyRole) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *PartyOrPartyRole) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *PartyOrPartyRole) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *PartyOrPartyRole) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *PartyOrPartyRole) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetNationality

`func (o *PartyOrPartyRole) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PartyOrPartyRole) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PartyOrPartyRole) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PartyOrPartyRole) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *PartyOrPartyRole) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *PartyOrPartyRole) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *PartyOrPartyRole) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *PartyOrPartyRole) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetBirthDate

`func (o *PartyOrPartyRole) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *PartyOrPartyRole) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *PartyOrPartyRole) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *PartyOrPartyRole) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetDeathDate

`func (o *PartyOrPartyRole) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *PartyOrPartyRole) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *PartyOrPartyRole) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *PartyOrPartyRole) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### GetTitle

`func (o *PartyOrPartyRole) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PartyOrPartyRole) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PartyOrPartyRole) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PartyOrPartyRole) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *PartyOrPartyRole) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *PartyOrPartyRole) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *PartyOrPartyRole) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *PartyOrPartyRole) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *PartyOrPartyRole) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *PartyOrPartyRole) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *PartyOrPartyRole) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *PartyOrPartyRole) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *PartyOrPartyRole) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *PartyOrPartyRole) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *PartyOrPartyRole) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *PartyOrPartyRole) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *PartyOrPartyRole) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *PartyOrPartyRole) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *PartyOrPartyRole) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *PartyOrPartyRole) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetLegalName

`func (o *PartyOrPartyRole) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *PartyOrPartyRole) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *PartyOrPartyRole) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *PartyOrPartyRole) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *PartyOrPartyRole) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PartyOrPartyRole) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PartyOrPartyRole) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PartyOrPartyRole) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetFormattedName

`func (o *PartyOrPartyRole) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *PartyOrPartyRole) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *PartyOrPartyRole) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *PartyOrPartyRole) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetLocation

`func (o *PartyOrPartyRole) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PartyOrPartyRole) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PartyOrPartyRole) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PartyOrPartyRole) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *PartyOrPartyRole) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyOrPartyRole) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyOrPartyRole) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyOrPartyRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *PartyOrPartyRole) GetOtherName() []OtherNameOrganization`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *PartyOrPartyRole) GetOtherNameOk() (*[]OtherNameOrganization, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *PartyOrPartyRole) SetOtherName(v []OtherNameOrganization)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *PartyOrPartyRole) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetIndividualIdentification

`func (o *PartyOrPartyRole) GetIndividualIdentification() []IndividualIdentification`

GetIndividualIdentification returns the IndividualIdentification field if non-nil, zero value otherwise.

### GetIndividualIdentificationOk

`func (o *PartyOrPartyRole) GetIndividualIdentificationOk() (*[]IndividualIdentification, bool)`

GetIndividualIdentificationOk returns a tuple with the IndividualIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentification

`func (o *PartyOrPartyRole) SetIndividualIdentification(v []IndividualIdentification)`

SetIndividualIdentification sets IndividualIdentification field to given value.

### HasIndividualIdentification

`func (o *PartyOrPartyRole) HasIndividualIdentification() bool`

HasIndividualIdentification returns a boolean if a field has been set.

### GetDisability

`func (o *PartyOrPartyRole) GetDisability() []Disability`

GetDisability returns the Disability field if non-nil, zero value otherwise.

### GetDisabilityOk

`func (o *PartyOrPartyRole) GetDisabilityOk() (*[]Disability, bool)`

GetDisabilityOk returns a tuple with the Disability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisability

`func (o *PartyOrPartyRole) SetDisability(v []Disability)`

SetDisability sets Disability field to given value.

### HasDisability

`func (o *PartyOrPartyRole) HasDisability() bool`

HasDisability returns a boolean if a field has been set.

### GetLanguageAbility

`func (o *PartyOrPartyRole) GetLanguageAbility() []LanguageAbility`

GetLanguageAbility returns the LanguageAbility field if non-nil, zero value otherwise.

### GetLanguageAbilityOk

`func (o *PartyOrPartyRole) GetLanguageAbilityOk() (*[]LanguageAbility, bool)`

GetLanguageAbilityOk returns a tuple with the LanguageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageAbility

`func (o *PartyOrPartyRole) SetLanguageAbility(v []LanguageAbility)`

SetLanguageAbility sets LanguageAbility field to given value.

### HasLanguageAbility

`func (o *PartyOrPartyRole) HasLanguageAbility() bool`

HasLanguageAbility returns a boolean if a field has been set.

### GetSkill

`func (o *PartyOrPartyRole) GetSkill() []Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PartyOrPartyRole) GetSkillOk() (*[]Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PartyOrPartyRole) SetSkill(v []Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PartyOrPartyRole) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetFamilyName

`func (o *PartyOrPartyRole) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PartyOrPartyRole) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PartyOrPartyRole) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PartyOrPartyRole) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *PartyOrPartyRole) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PartyOrPartyRole) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PartyOrPartyRole) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PartyOrPartyRole) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetIsLegalEntity

`func (o *PartyOrPartyRole) GetIsLegalEntity() bool`

GetIsLegalEntity returns the IsLegalEntity field if non-nil, zero value otherwise.

### GetIsLegalEntityOk

`func (o *PartyOrPartyRole) GetIsLegalEntityOk() (*bool, bool)`

GetIsLegalEntityOk returns a tuple with the IsLegalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalEntity

`func (o *PartyOrPartyRole) SetIsLegalEntity(v bool)`

SetIsLegalEntity sets IsLegalEntity field to given value.

### HasIsLegalEntity

`func (o *PartyOrPartyRole) HasIsLegalEntity() bool`

HasIsLegalEntity returns a boolean if a field has been set.

### GetIsHeadOffice

`func (o *PartyOrPartyRole) GetIsHeadOffice() bool`

GetIsHeadOffice returns the IsHeadOffice field if non-nil, zero value otherwise.

### GetIsHeadOfficeOk

`func (o *PartyOrPartyRole) GetIsHeadOfficeOk() (*bool, bool)`

GetIsHeadOfficeOk returns a tuple with the IsHeadOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadOffice

`func (o *PartyOrPartyRole) SetIsHeadOffice(v bool)`

SetIsHeadOffice sets IsHeadOffice field to given value.

### HasIsHeadOffice

`func (o *PartyOrPartyRole) HasIsHeadOffice() bool`

HasIsHeadOffice returns a boolean if a field has been set.

### GetOrganizationType

`func (o *PartyOrPartyRole) GetOrganizationType() string`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *PartyOrPartyRole) GetOrganizationTypeOk() (*string, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *PartyOrPartyRole) SetOrganizationType(v string)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *PartyOrPartyRole) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetExistsDuring

`func (o *PartyOrPartyRole) GetExistsDuring() TimePeriod`

GetExistsDuring returns the ExistsDuring field if non-nil, zero value otherwise.

### GetExistsDuringOk

`func (o *PartyOrPartyRole) GetExistsDuringOk() (*TimePeriod, bool)`

GetExistsDuringOk returns a tuple with the ExistsDuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsDuring

`func (o *PartyOrPartyRole) SetExistsDuring(v TimePeriod)`

SetExistsDuring sets ExistsDuring field to given value.

### HasExistsDuring

`func (o *PartyOrPartyRole) HasExistsDuring() bool`

HasExistsDuring returns a boolean if a field has been set.

### GetNameType

`func (o *PartyOrPartyRole) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *PartyOrPartyRole) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *PartyOrPartyRole) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *PartyOrPartyRole) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetOrganizationIdentification

`func (o *PartyOrPartyRole) GetOrganizationIdentification() []OrganizationIdentification`

GetOrganizationIdentification returns the OrganizationIdentification field if non-nil, zero value otherwise.

### GetOrganizationIdentificationOk

`func (o *PartyOrPartyRole) GetOrganizationIdentificationOk() (*[]OrganizationIdentification, bool)`

GetOrganizationIdentificationOk returns a tuple with the OrganizationIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentification

`func (o *PartyOrPartyRole) SetOrganizationIdentification(v []OrganizationIdentification)`

SetOrganizationIdentification sets OrganizationIdentification field to given value.

### HasOrganizationIdentification

`func (o *PartyOrPartyRole) HasOrganizationIdentification() bool`

HasOrganizationIdentification returns a boolean if a field has been set.

### GetOrganizationChildRelationship

`func (o *PartyOrPartyRole) GetOrganizationChildRelationship() []OrganizationChildRelationship`

GetOrganizationChildRelationship returns the OrganizationChildRelationship field if non-nil, zero value otherwise.

### GetOrganizationChildRelationshipOk

`func (o *PartyOrPartyRole) GetOrganizationChildRelationshipOk() (*[]OrganizationChildRelationship, bool)`

GetOrganizationChildRelationshipOk returns a tuple with the OrganizationChildRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationChildRelationship

`func (o *PartyOrPartyRole) SetOrganizationChildRelationship(v []OrganizationChildRelationship)`

SetOrganizationChildRelationship sets OrganizationChildRelationship field to given value.

### HasOrganizationChildRelationship

`func (o *PartyOrPartyRole) HasOrganizationChildRelationship() bool`

HasOrganizationChildRelationship returns a boolean if a field has been set.

### GetOrganizationParentRelationship

`func (o *PartyOrPartyRole) GetOrganizationParentRelationship() OrganizationParentRelationship`

GetOrganizationParentRelationship returns the OrganizationParentRelationship field if non-nil, zero value otherwise.

### GetOrganizationParentRelationshipOk

`func (o *PartyOrPartyRole) GetOrganizationParentRelationshipOk() (*OrganizationParentRelationship, bool)`

GetOrganizationParentRelationshipOk returns a tuple with the OrganizationParentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationParentRelationship

`func (o *PartyOrPartyRole) SetOrganizationParentRelationship(v OrganizationParentRelationship)`

SetOrganizationParentRelationship sets OrganizationParentRelationship field to given value.

### HasOrganizationParentRelationship

`func (o *PartyOrPartyRole) HasOrganizationParentRelationship() bool`

HasOrganizationParentRelationship returns a boolean if a field has been set.

### GetTradingName

`func (o *PartyOrPartyRole) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *PartyOrPartyRole) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *PartyOrPartyRole) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *PartyOrPartyRole) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetDescription

`func (o *PartyOrPartyRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyOrPartyRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyOrPartyRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyOrPartyRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyOrPartyRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyOrPartyRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyOrPartyRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyOrPartyRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyOrPartyRole) GetEngagedParty() PartyRef`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyOrPartyRole) GetEngagedPartyOk() (*PartyRef, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyOrPartyRole) SetEngagedParty(v PartyRef)`

SetEngagedParty sets EngagedParty field to given value.

### HasEngagedParty

`func (o *PartyOrPartyRole) HasEngagedParty() bool`

HasEngagedParty returns a boolean if a field has been set.

### GetPartyRoleSpecification

`func (o *PartyOrPartyRole) GetPartyRoleSpecification() PartyRoleSpecificationRef`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyOrPartyRole) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRef, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyOrPartyRole) SetPartyRoleSpecification(v PartyRoleSpecificationRef)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyOrPartyRole) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyOrPartyRole) GetCharacteristic() []Characteristic`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyOrPartyRole) GetCharacteristicOk() (*[]Characteristic, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyOrPartyRole) SetCharacteristic(v []Characteristic)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyOrPartyRole) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyOrPartyRole) GetAccount() []AccountRef`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyOrPartyRole) GetAccountOk() (*[]AccountRef, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyOrPartyRole) SetAccount(v []AccountRef)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyOrPartyRole) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyOrPartyRole) GetAgreement() []AgreementRef`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyOrPartyRole) GetAgreementOk() (*[]AgreementRef, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyOrPartyRole) SetAgreement(v []AgreementRef)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyOrPartyRole) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyOrPartyRole) GetPaymentMethod() []PaymentMethodRef`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyOrPartyRole) GetPaymentMethodOk() (*[]PaymentMethodRef, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyOrPartyRole) SetPaymentMethod(v []PaymentMethodRef)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyOrPartyRole) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyOrPartyRole) GetCreditProfile() []CreditProfile`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyOrPartyRole) GetCreditProfileOk() (*[]CreditProfile, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyOrPartyRole) SetCreditProfile(v []CreditProfile)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyOrPartyRole) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyOrPartyRole) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyOrPartyRole) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyOrPartyRole) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyOrPartyRole) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyOrPartyRole) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyOrPartyRole) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyOrPartyRole) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyOrPartyRole) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


