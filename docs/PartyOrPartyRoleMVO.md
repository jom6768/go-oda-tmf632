# PartyOrPartyRoleMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | The URI of the referred entity. | [optional] 
**Id** | **string** | The identifier of the referred entity. | 
**Name** | **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | 
**ReferredType** | Pointer to **string** | The actual type of the target instance when needed for disambiguation. | [optional] 
**PartyId** | Pointer to **string** | The identifier of the engaged party that is linked to the PartyRole object. | [optional] 
**PartyName** | Pointer to **string** | The name of the engaged party that is linked to the PartyRole object. | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifierMVO**](ExternalIdentifierMVO.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]CharacteristicMVO**](CharacteristicMVO.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificateMVO**](TaxExemptionCertificateMVO.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfileMVO**](PartyCreditProfileMVO.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleMVO**](RelatedPartyOrPartyRoleMVO.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumMVO**](ContactMediumMVO.md) |  | [optional] 
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
**OtherName** | Pointer to [**[]OtherNameOrganizationMVO**](OtherNameOrganizationMVO.md) | List of additional names by which the organization is known | [optional] 
**IndividualIdentification** | Pointer to [**[]IndividualIdentificationMVO**](IndividualIdentificationMVO.md) | List of official identifications issued to the individual, such as passport, driving licence, social security number | [optional] 
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
**OrganizationIdentification** | Pointer to [**[]OrganizationIdentificationMVO**](OrganizationIdentificationMVO.md) | List of official identifiers given to the organization, for example company number in the registry of companies | [optional] 
**OrganizationChildRelationship** | Pointer to [**[]OrganizationChildRelationshipMVO**](OrganizationChildRelationshipMVO.md) | List of organizations that are contained within this organization. For example if this organization is the Legal Department, the child organizations might include Claims, Courts, Contracts | [optional] 
**OrganizationParentRelationship** | Pointer to [**OrganizationParentRelationshipMVO**](OrganizationParentRelationshipMVO.md) |  | [optional] 
**TradingName** | Pointer to **string** | Name that the organization (unit) trades under | [optional] 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | [**PartyRefMVO**](PartyRefMVO.md) |  | 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRefMVO**](PartyRoleSpecificationRefMVO.md) |  | [optional] 
**Characteristic** | Pointer to [**[]CharacteristicMVO**](CharacteristicMVO.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRefMVO**](AccountRefMVO.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRefMVO**](AgreementRefMVO.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRefMVO**](PaymentMethodRefMVO.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfileMVO**](CreditProfileMVO.md) |  | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyOrPartyRoleMVO

`func NewPartyOrPartyRoleMVO(type_ string, id string, name string, engagedParty PartyRefMVO, ) *PartyOrPartyRoleMVO`

NewPartyOrPartyRoleMVO instantiates a new PartyOrPartyRoleMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyOrPartyRoleMVOWithDefaults

`func NewPartyOrPartyRoleMVOWithDefaults() *PartyOrPartyRoleMVO`

NewPartyOrPartyRoleMVOWithDefaults instantiates a new PartyOrPartyRoleMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyOrPartyRoleMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyOrPartyRoleMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyOrPartyRoleMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyOrPartyRoleMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyOrPartyRoleMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyOrPartyRoleMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyOrPartyRoleMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyOrPartyRoleMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyOrPartyRoleMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyOrPartyRoleMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyOrPartyRoleMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyOrPartyRoleMVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyOrPartyRoleMVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyOrPartyRoleMVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyOrPartyRoleMVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyOrPartyRoleMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyOrPartyRoleMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyOrPartyRoleMVO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyOrPartyRoleMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyOrPartyRoleMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyOrPartyRoleMVO) SetName(v string)`

SetName sets Name field to given value.


### GetReferredType

`func (o *PartyOrPartyRoleMVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyOrPartyRoleMVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyOrPartyRoleMVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyOrPartyRoleMVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetPartyId

`func (o *PartyOrPartyRoleMVO) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *PartyOrPartyRoleMVO) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *PartyOrPartyRoleMVO) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *PartyOrPartyRoleMVO) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *PartyOrPartyRoleMVO) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyOrPartyRoleMVO) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyOrPartyRoleMVO) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyOrPartyRoleMVO) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetExternalReference

`func (o *PartyOrPartyRoleMVO) GetExternalReference() []ExternalIdentifierMVO`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PartyOrPartyRoleMVO) GetExternalReferenceOk() (*[]ExternalIdentifierMVO, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PartyOrPartyRoleMVO) SetExternalReference(v []ExternalIdentifierMVO)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PartyOrPartyRoleMVO) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *PartyOrPartyRoleMVO) GetPartyCharacteristic() []CharacteristicMVO`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *PartyOrPartyRoleMVO) GetPartyCharacteristicOk() (*[]CharacteristicMVO, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *PartyOrPartyRoleMVO) SetPartyCharacteristic(v []CharacteristicMVO)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *PartyOrPartyRoleMVO) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *PartyOrPartyRoleMVO) GetTaxExemptionCertificate() []TaxExemptionCertificateMVO`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *PartyOrPartyRoleMVO) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificateMVO, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *PartyOrPartyRoleMVO) SetTaxExemptionCertificate(v []TaxExemptionCertificateMVO)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *PartyOrPartyRoleMVO) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *PartyOrPartyRoleMVO) GetCreditRating() []PartyCreditProfileMVO`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *PartyOrPartyRoleMVO) GetCreditRatingOk() (*[]PartyCreditProfileMVO, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *PartyOrPartyRoleMVO) SetCreditRating(v []PartyCreditProfileMVO)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *PartyOrPartyRoleMVO) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyOrPartyRoleMVO) GetRelatedParty() []RelatedPartyOrPartyRoleMVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyOrPartyRoleMVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleMVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyOrPartyRoleMVO) SetRelatedParty(v []RelatedPartyOrPartyRoleMVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyOrPartyRoleMVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyOrPartyRoleMVO) GetContactMedium() []ContactMediumMVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyOrPartyRoleMVO) GetContactMediumOk() (*[]ContactMediumMVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyOrPartyRoleMVO) SetContactMedium(v []ContactMediumMVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyOrPartyRoleMVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetGender

`func (o *PartyOrPartyRoleMVO) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PartyOrPartyRoleMVO) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PartyOrPartyRoleMVO) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PartyOrPartyRoleMVO) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *PartyOrPartyRoleMVO) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PartyOrPartyRoleMVO) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PartyOrPartyRoleMVO) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PartyOrPartyRoleMVO) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *PartyOrPartyRoleMVO) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *PartyOrPartyRoleMVO) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *PartyOrPartyRoleMVO) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *PartyOrPartyRoleMVO) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetNationality

`func (o *PartyOrPartyRoleMVO) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PartyOrPartyRoleMVO) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PartyOrPartyRoleMVO) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PartyOrPartyRoleMVO) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *PartyOrPartyRoleMVO) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *PartyOrPartyRoleMVO) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *PartyOrPartyRoleMVO) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *PartyOrPartyRoleMVO) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetBirthDate

`func (o *PartyOrPartyRoleMVO) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *PartyOrPartyRoleMVO) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *PartyOrPartyRoleMVO) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *PartyOrPartyRoleMVO) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetDeathDate

`func (o *PartyOrPartyRoleMVO) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *PartyOrPartyRoleMVO) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *PartyOrPartyRoleMVO) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *PartyOrPartyRoleMVO) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### GetTitle

`func (o *PartyOrPartyRoleMVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PartyOrPartyRoleMVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PartyOrPartyRoleMVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PartyOrPartyRoleMVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *PartyOrPartyRoleMVO) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *PartyOrPartyRoleMVO) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *PartyOrPartyRoleMVO) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *PartyOrPartyRoleMVO) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *PartyOrPartyRoleMVO) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *PartyOrPartyRoleMVO) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *PartyOrPartyRoleMVO) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *PartyOrPartyRoleMVO) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *PartyOrPartyRoleMVO) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *PartyOrPartyRoleMVO) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *PartyOrPartyRoleMVO) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *PartyOrPartyRoleMVO) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *PartyOrPartyRoleMVO) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *PartyOrPartyRoleMVO) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *PartyOrPartyRoleMVO) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *PartyOrPartyRoleMVO) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetLegalName

`func (o *PartyOrPartyRoleMVO) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *PartyOrPartyRoleMVO) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *PartyOrPartyRoleMVO) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *PartyOrPartyRoleMVO) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *PartyOrPartyRoleMVO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PartyOrPartyRoleMVO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PartyOrPartyRoleMVO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PartyOrPartyRoleMVO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetFormattedName

`func (o *PartyOrPartyRoleMVO) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *PartyOrPartyRoleMVO) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *PartyOrPartyRoleMVO) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *PartyOrPartyRoleMVO) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetLocation

`func (o *PartyOrPartyRoleMVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PartyOrPartyRoleMVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PartyOrPartyRoleMVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PartyOrPartyRoleMVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *PartyOrPartyRoleMVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyOrPartyRoleMVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyOrPartyRoleMVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyOrPartyRoleMVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *PartyOrPartyRoleMVO) GetOtherName() []OtherNameOrganizationMVO`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *PartyOrPartyRoleMVO) GetOtherNameOk() (*[]OtherNameOrganizationMVO, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *PartyOrPartyRoleMVO) SetOtherName(v []OtherNameOrganizationMVO)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *PartyOrPartyRoleMVO) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetIndividualIdentification

`func (o *PartyOrPartyRoleMVO) GetIndividualIdentification() []IndividualIdentificationMVO`

GetIndividualIdentification returns the IndividualIdentification field if non-nil, zero value otherwise.

### GetIndividualIdentificationOk

`func (o *PartyOrPartyRoleMVO) GetIndividualIdentificationOk() (*[]IndividualIdentificationMVO, bool)`

GetIndividualIdentificationOk returns a tuple with the IndividualIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentification

`func (o *PartyOrPartyRoleMVO) SetIndividualIdentification(v []IndividualIdentificationMVO)`

SetIndividualIdentification sets IndividualIdentification field to given value.

### HasIndividualIdentification

`func (o *PartyOrPartyRoleMVO) HasIndividualIdentification() bool`

HasIndividualIdentification returns a boolean if a field has been set.

### GetDisability

`func (o *PartyOrPartyRoleMVO) GetDisability() []Disability`

GetDisability returns the Disability field if non-nil, zero value otherwise.

### GetDisabilityOk

`func (o *PartyOrPartyRoleMVO) GetDisabilityOk() (*[]Disability, bool)`

GetDisabilityOk returns a tuple with the Disability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisability

`func (o *PartyOrPartyRoleMVO) SetDisability(v []Disability)`

SetDisability sets Disability field to given value.

### HasDisability

`func (o *PartyOrPartyRoleMVO) HasDisability() bool`

HasDisability returns a boolean if a field has been set.

### GetLanguageAbility

`func (o *PartyOrPartyRoleMVO) GetLanguageAbility() []LanguageAbility`

GetLanguageAbility returns the LanguageAbility field if non-nil, zero value otherwise.

### GetLanguageAbilityOk

`func (o *PartyOrPartyRoleMVO) GetLanguageAbilityOk() (*[]LanguageAbility, bool)`

GetLanguageAbilityOk returns a tuple with the LanguageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageAbility

`func (o *PartyOrPartyRoleMVO) SetLanguageAbility(v []LanguageAbility)`

SetLanguageAbility sets LanguageAbility field to given value.

### HasLanguageAbility

`func (o *PartyOrPartyRoleMVO) HasLanguageAbility() bool`

HasLanguageAbility returns a boolean if a field has been set.

### GetSkill

`func (o *PartyOrPartyRoleMVO) GetSkill() []Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PartyOrPartyRoleMVO) GetSkillOk() (*[]Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PartyOrPartyRoleMVO) SetSkill(v []Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PartyOrPartyRoleMVO) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetFamilyName

`func (o *PartyOrPartyRoleMVO) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PartyOrPartyRoleMVO) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PartyOrPartyRoleMVO) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *PartyOrPartyRoleMVO) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *PartyOrPartyRoleMVO) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PartyOrPartyRoleMVO) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PartyOrPartyRoleMVO) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *PartyOrPartyRoleMVO) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetIsLegalEntity

`func (o *PartyOrPartyRoleMVO) GetIsLegalEntity() bool`

GetIsLegalEntity returns the IsLegalEntity field if non-nil, zero value otherwise.

### GetIsLegalEntityOk

`func (o *PartyOrPartyRoleMVO) GetIsLegalEntityOk() (*bool, bool)`

GetIsLegalEntityOk returns a tuple with the IsLegalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalEntity

`func (o *PartyOrPartyRoleMVO) SetIsLegalEntity(v bool)`

SetIsLegalEntity sets IsLegalEntity field to given value.

### HasIsLegalEntity

`func (o *PartyOrPartyRoleMVO) HasIsLegalEntity() bool`

HasIsLegalEntity returns a boolean if a field has been set.

### GetIsHeadOffice

`func (o *PartyOrPartyRoleMVO) GetIsHeadOffice() bool`

GetIsHeadOffice returns the IsHeadOffice field if non-nil, zero value otherwise.

### GetIsHeadOfficeOk

`func (o *PartyOrPartyRoleMVO) GetIsHeadOfficeOk() (*bool, bool)`

GetIsHeadOfficeOk returns a tuple with the IsHeadOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadOffice

`func (o *PartyOrPartyRoleMVO) SetIsHeadOffice(v bool)`

SetIsHeadOffice sets IsHeadOffice field to given value.

### HasIsHeadOffice

`func (o *PartyOrPartyRoleMVO) HasIsHeadOffice() bool`

HasIsHeadOffice returns a boolean if a field has been set.

### GetOrganizationType

`func (o *PartyOrPartyRoleMVO) GetOrganizationType() string`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *PartyOrPartyRoleMVO) GetOrganizationTypeOk() (*string, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *PartyOrPartyRoleMVO) SetOrganizationType(v string)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *PartyOrPartyRoleMVO) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetExistsDuring

`func (o *PartyOrPartyRoleMVO) GetExistsDuring() TimePeriod`

GetExistsDuring returns the ExistsDuring field if non-nil, zero value otherwise.

### GetExistsDuringOk

`func (o *PartyOrPartyRoleMVO) GetExistsDuringOk() (*TimePeriod, bool)`

GetExistsDuringOk returns a tuple with the ExistsDuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsDuring

`func (o *PartyOrPartyRoleMVO) SetExistsDuring(v TimePeriod)`

SetExistsDuring sets ExistsDuring field to given value.

### HasExistsDuring

`func (o *PartyOrPartyRoleMVO) HasExistsDuring() bool`

HasExistsDuring returns a boolean if a field has been set.

### GetNameType

`func (o *PartyOrPartyRoleMVO) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *PartyOrPartyRoleMVO) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *PartyOrPartyRoleMVO) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *PartyOrPartyRoleMVO) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetOrganizationIdentification

`func (o *PartyOrPartyRoleMVO) GetOrganizationIdentification() []OrganizationIdentificationMVO`

GetOrganizationIdentification returns the OrganizationIdentification field if non-nil, zero value otherwise.

### GetOrganizationIdentificationOk

`func (o *PartyOrPartyRoleMVO) GetOrganizationIdentificationOk() (*[]OrganizationIdentificationMVO, bool)`

GetOrganizationIdentificationOk returns a tuple with the OrganizationIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentification

`func (o *PartyOrPartyRoleMVO) SetOrganizationIdentification(v []OrganizationIdentificationMVO)`

SetOrganizationIdentification sets OrganizationIdentification field to given value.

### HasOrganizationIdentification

`func (o *PartyOrPartyRoleMVO) HasOrganizationIdentification() bool`

HasOrganizationIdentification returns a boolean if a field has been set.

### GetOrganizationChildRelationship

`func (o *PartyOrPartyRoleMVO) GetOrganizationChildRelationship() []OrganizationChildRelationshipMVO`

GetOrganizationChildRelationship returns the OrganizationChildRelationship field if non-nil, zero value otherwise.

### GetOrganizationChildRelationshipOk

`func (o *PartyOrPartyRoleMVO) GetOrganizationChildRelationshipOk() (*[]OrganizationChildRelationshipMVO, bool)`

GetOrganizationChildRelationshipOk returns a tuple with the OrganizationChildRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationChildRelationship

`func (o *PartyOrPartyRoleMVO) SetOrganizationChildRelationship(v []OrganizationChildRelationshipMVO)`

SetOrganizationChildRelationship sets OrganizationChildRelationship field to given value.

### HasOrganizationChildRelationship

`func (o *PartyOrPartyRoleMVO) HasOrganizationChildRelationship() bool`

HasOrganizationChildRelationship returns a boolean if a field has been set.

### GetOrganizationParentRelationship

`func (o *PartyOrPartyRoleMVO) GetOrganizationParentRelationship() OrganizationParentRelationshipMVO`

GetOrganizationParentRelationship returns the OrganizationParentRelationship field if non-nil, zero value otherwise.

### GetOrganizationParentRelationshipOk

`func (o *PartyOrPartyRoleMVO) GetOrganizationParentRelationshipOk() (*OrganizationParentRelationshipMVO, bool)`

GetOrganizationParentRelationshipOk returns a tuple with the OrganizationParentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationParentRelationship

`func (o *PartyOrPartyRoleMVO) SetOrganizationParentRelationship(v OrganizationParentRelationshipMVO)`

SetOrganizationParentRelationship sets OrganizationParentRelationship field to given value.

### HasOrganizationParentRelationship

`func (o *PartyOrPartyRoleMVO) HasOrganizationParentRelationship() bool`

HasOrganizationParentRelationship returns a boolean if a field has been set.

### GetTradingName

`func (o *PartyOrPartyRoleMVO) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *PartyOrPartyRoleMVO) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *PartyOrPartyRoleMVO) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *PartyOrPartyRoleMVO) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetDescription

`func (o *PartyOrPartyRoleMVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyOrPartyRoleMVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyOrPartyRoleMVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyOrPartyRoleMVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyOrPartyRoleMVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyOrPartyRoleMVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyOrPartyRoleMVO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyOrPartyRoleMVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyOrPartyRoleMVO) GetEngagedParty() PartyRefMVO`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyOrPartyRoleMVO) GetEngagedPartyOk() (*PartyRefMVO, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyOrPartyRoleMVO) SetEngagedParty(v PartyRefMVO)`

SetEngagedParty sets EngagedParty field to given value.


### GetPartyRoleSpecification

`func (o *PartyOrPartyRoleMVO) GetPartyRoleSpecification() PartyRoleSpecificationRefMVO`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyOrPartyRoleMVO) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRefMVO, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyOrPartyRoleMVO) SetPartyRoleSpecification(v PartyRoleSpecificationRefMVO)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyOrPartyRoleMVO) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyOrPartyRoleMVO) GetCharacteristic() []CharacteristicMVO`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyOrPartyRoleMVO) GetCharacteristicOk() (*[]CharacteristicMVO, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyOrPartyRoleMVO) SetCharacteristic(v []CharacteristicMVO)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyOrPartyRoleMVO) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyOrPartyRoleMVO) GetAccount() []AccountRefMVO`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyOrPartyRoleMVO) GetAccountOk() (*[]AccountRefMVO, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyOrPartyRoleMVO) SetAccount(v []AccountRefMVO)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyOrPartyRoleMVO) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyOrPartyRoleMVO) GetAgreement() []AgreementRefMVO`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyOrPartyRoleMVO) GetAgreementOk() (*[]AgreementRefMVO, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyOrPartyRoleMVO) SetAgreement(v []AgreementRefMVO)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyOrPartyRoleMVO) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyOrPartyRoleMVO) GetPaymentMethod() []PaymentMethodRefMVO`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyOrPartyRoleMVO) GetPaymentMethodOk() (*[]PaymentMethodRefMVO, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyOrPartyRoleMVO) SetPaymentMethod(v []PaymentMethodRefMVO)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyOrPartyRoleMVO) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyOrPartyRoleMVO) GetCreditProfile() []CreditProfileMVO`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyOrPartyRoleMVO) GetCreditProfileOk() (*[]CreditProfileMVO, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyOrPartyRoleMVO) SetCreditProfile(v []CreditProfileMVO)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyOrPartyRoleMVO) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyOrPartyRoleMVO) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyOrPartyRoleMVO) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyOrPartyRoleMVO) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyOrPartyRoleMVO) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyOrPartyRoleMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyOrPartyRoleMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyOrPartyRoleMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyOrPartyRoleMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


