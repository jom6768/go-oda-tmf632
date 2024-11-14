# PartyOrPartyRoleFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | **string** | unique identifier | 
**Href** | Pointer to **string** | The URI of the referred entity. | [optional] 
**Name** | **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | 
**ReferredType** | Pointer to **string** | The actual type of the target instance when needed for disambiguation. | [optional] 
**PartyId** | Pointer to **string** | The identifier of the engaged party that is linked to the PartyRole object. | [optional] 
**PartyName** | Pointer to **string** | The name of the engaged party that is linked to the PartyRole object. | [optional] 
**ExternalReference** | Pointer to [**[]ExternalIdentifierFVO**](ExternalIdentifierFVO.md) | List of identifiers of the Party in an external system, for example when party information is imported from a commerce system | [optional] 
**PartyCharacteristic** | Pointer to [**[]CharacteristicFVO**](CharacteristicFVO.md) | List of additional characteristics that a Party can take on. | [optional] 
**TaxExemptionCertificate** | Pointer to [**[]TaxExemptionCertificateFVO**](TaxExemptionCertificateFVO.md) | List of tax exemptions granted to the party. For example, a war veteran might have partial exemption from state tax and a full exemption from federal tax | [optional] 
**CreditRating** | Pointer to [**[]PartyCreditProfileFVO**](PartyCreditProfileFVO.md) | List of credit profiles and scores for the party, typically received from an external credit broker | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleFVO**](RelatedPartyOrPartyRoleFVO.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumFVO**](ContactMediumFVO.md) |  | [optional] 
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
**OtherName** | Pointer to [**[]OtherNameOrganizationFVO**](OtherNameOrganizationFVO.md) | List of additional names by which the organization is known | [optional] 
**IndividualIdentification** | Pointer to [**[]IndividualIdentificationFVO**](IndividualIdentificationFVO.md) | List of official identifications issued to the individual, such as passport, driving licence, social security number | [optional] 
**Disability** | Pointer to [**[]Disability**](Disability.md) | List of disabilities suffered by the individual | [optional] 
**LanguageAbility** | Pointer to [**[]LanguageAbility**](LanguageAbility.md) | List of national languages known by the individual | [optional] 
**Skill** | Pointer to [**[]Skill**](Skill.md) | List of skills exhibited by the individual | [optional] 
**FamilyName** | **string** | Contains the non-chosen or inherited name. Also known as last name in the Western context | 
**GivenName** | **string** | First name of the individual | 
**IsLegalEntity** | Pointer to **bool** | If value is true, the organization is a legal entity known by a national referential. | [optional] 
**IsHeadOffice** | Pointer to **bool** | If value is true, the organization is the head office | [optional] 
**OrganizationType** | Pointer to **string** | Type of Organization (company, department...) | [optional] 
**ExistsDuring** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**NameType** | Pointer to **string** | Type of the name : Co, Inc, Ltd, etc. | [optional] 
**OrganizationIdentification** | Pointer to [**[]OrganizationIdentificationFVO**](OrganizationIdentificationFVO.md) | List of official identifiers given to the organization, for example company number in the registry of companies | [optional] 
**OrganizationChildRelationship** | Pointer to [**[]OrganizationChildRelationshipFVO**](OrganizationChildRelationshipFVO.md) | List of organizations that are contained within this organization. For example if this organization is the Legal Department, the child organizations might include Claims, Courts, Contracts | [optional] 
**OrganizationParentRelationship** | Pointer to [**OrganizationParentRelationshipFVO**](OrganizationParentRelationshipFVO.md) |  | [optional] 
**TradingName** | Pointer to **string** | Name that the organization (unit) trades under | [optional] 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | [**PartyRefFVO**](PartyRefFVO.md) |  | 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRefFVO**](PartyRoleSpecificationRefFVO.md) |  | [optional] 
**Characteristic** | Pointer to [**[]CharacteristicFVO**](CharacteristicFVO.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRefFVO**](AccountRefFVO.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRefFVO**](AgreementRefFVO.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRefFVO**](PaymentMethodRefFVO.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfileFVO**](CreditProfileFVO.md) |  | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyOrPartyRoleFVO

`func NewPartyOrPartyRoleFVO(type_ string, id string, name string, familyName string, givenName string, engagedParty PartyRefFVO, ) *PartyOrPartyRoleFVO`

NewPartyOrPartyRoleFVO instantiates a new PartyOrPartyRoleFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyOrPartyRoleFVOWithDefaults

`func NewPartyOrPartyRoleFVOWithDefaults() *PartyOrPartyRoleFVO`

NewPartyOrPartyRoleFVOWithDefaults instantiates a new PartyOrPartyRoleFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyOrPartyRoleFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyOrPartyRoleFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyOrPartyRoleFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyOrPartyRoleFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyOrPartyRoleFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyOrPartyRoleFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyOrPartyRoleFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyOrPartyRoleFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyOrPartyRoleFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyOrPartyRoleFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyOrPartyRoleFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *PartyOrPartyRoleFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyOrPartyRoleFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyOrPartyRoleFVO) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *PartyOrPartyRoleFVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyOrPartyRoleFVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyOrPartyRoleFVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyOrPartyRoleFVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *PartyOrPartyRoleFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyOrPartyRoleFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyOrPartyRoleFVO) SetName(v string)`

SetName sets Name field to given value.


### GetReferredType

`func (o *PartyOrPartyRoleFVO) GetReferredType() string`

GetReferredType returns the ReferredType field if non-nil, zero value otherwise.

### GetReferredTypeOk

`func (o *PartyOrPartyRoleFVO) GetReferredTypeOk() (*string, bool)`

GetReferredTypeOk returns a tuple with the ReferredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredType

`func (o *PartyOrPartyRoleFVO) SetReferredType(v string)`

SetReferredType sets ReferredType field to given value.

### HasReferredType

`func (o *PartyOrPartyRoleFVO) HasReferredType() bool`

HasReferredType returns a boolean if a field has been set.

### GetPartyId

`func (o *PartyOrPartyRoleFVO) GetPartyId() string`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *PartyOrPartyRoleFVO) GetPartyIdOk() (*string, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *PartyOrPartyRoleFVO) SetPartyId(v string)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *PartyOrPartyRoleFVO) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPartyName

`func (o *PartyOrPartyRoleFVO) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyOrPartyRoleFVO) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyOrPartyRoleFVO) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyOrPartyRoleFVO) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetExternalReference

`func (o *PartyOrPartyRoleFVO) GetExternalReference() []ExternalIdentifierFVO`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PartyOrPartyRoleFVO) GetExternalReferenceOk() (*[]ExternalIdentifierFVO, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PartyOrPartyRoleFVO) SetExternalReference(v []ExternalIdentifierFVO)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *PartyOrPartyRoleFVO) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetPartyCharacteristic

`func (o *PartyOrPartyRoleFVO) GetPartyCharacteristic() []CharacteristicFVO`

GetPartyCharacteristic returns the PartyCharacteristic field if non-nil, zero value otherwise.

### GetPartyCharacteristicOk

`func (o *PartyOrPartyRoleFVO) GetPartyCharacteristicOk() (*[]CharacteristicFVO, bool)`

GetPartyCharacteristicOk returns a tuple with the PartyCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyCharacteristic

`func (o *PartyOrPartyRoleFVO) SetPartyCharacteristic(v []CharacteristicFVO)`

SetPartyCharacteristic sets PartyCharacteristic field to given value.

### HasPartyCharacteristic

`func (o *PartyOrPartyRoleFVO) HasPartyCharacteristic() bool`

HasPartyCharacteristic returns a boolean if a field has been set.

### GetTaxExemptionCertificate

`func (o *PartyOrPartyRoleFVO) GetTaxExemptionCertificate() []TaxExemptionCertificateFVO`

GetTaxExemptionCertificate returns the TaxExemptionCertificate field if non-nil, zero value otherwise.

### GetTaxExemptionCertificateOk

`func (o *PartyOrPartyRoleFVO) GetTaxExemptionCertificateOk() (*[]TaxExemptionCertificateFVO, bool)`

GetTaxExemptionCertificateOk returns a tuple with the TaxExemptionCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExemptionCertificate

`func (o *PartyOrPartyRoleFVO) SetTaxExemptionCertificate(v []TaxExemptionCertificateFVO)`

SetTaxExemptionCertificate sets TaxExemptionCertificate field to given value.

### HasTaxExemptionCertificate

`func (o *PartyOrPartyRoleFVO) HasTaxExemptionCertificate() bool`

HasTaxExemptionCertificate returns a boolean if a field has been set.

### GetCreditRating

`func (o *PartyOrPartyRoleFVO) GetCreditRating() []PartyCreditProfileFVO`

GetCreditRating returns the CreditRating field if non-nil, zero value otherwise.

### GetCreditRatingOk

`func (o *PartyOrPartyRoleFVO) GetCreditRatingOk() (*[]PartyCreditProfileFVO, bool)`

GetCreditRatingOk returns a tuple with the CreditRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRating

`func (o *PartyOrPartyRoleFVO) SetCreditRating(v []PartyCreditProfileFVO)`

SetCreditRating sets CreditRating field to given value.

### HasCreditRating

`func (o *PartyOrPartyRoleFVO) HasCreditRating() bool`

HasCreditRating returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyOrPartyRoleFVO) GetRelatedParty() []RelatedPartyOrPartyRoleFVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyOrPartyRoleFVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleFVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyOrPartyRoleFVO) SetRelatedParty(v []RelatedPartyOrPartyRoleFVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyOrPartyRoleFVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyOrPartyRoleFVO) GetContactMedium() []ContactMediumFVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyOrPartyRoleFVO) GetContactMediumOk() (*[]ContactMediumFVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyOrPartyRoleFVO) SetContactMedium(v []ContactMediumFVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyOrPartyRoleFVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetGender

`func (o *PartyOrPartyRoleFVO) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PartyOrPartyRoleFVO) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PartyOrPartyRoleFVO) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PartyOrPartyRoleFVO) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *PartyOrPartyRoleFVO) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PartyOrPartyRoleFVO) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PartyOrPartyRoleFVO) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PartyOrPartyRoleFVO) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *PartyOrPartyRoleFVO) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *PartyOrPartyRoleFVO) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *PartyOrPartyRoleFVO) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *PartyOrPartyRoleFVO) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetNationality

`func (o *PartyOrPartyRoleFVO) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PartyOrPartyRoleFVO) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PartyOrPartyRoleFVO) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PartyOrPartyRoleFVO) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *PartyOrPartyRoleFVO) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *PartyOrPartyRoleFVO) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *PartyOrPartyRoleFVO) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *PartyOrPartyRoleFVO) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetBirthDate

`func (o *PartyOrPartyRoleFVO) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *PartyOrPartyRoleFVO) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *PartyOrPartyRoleFVO) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *PartyOrPartyRoleFVO) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetDeathDate

`func (o *PartyOrPartyRoleFVO) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *PartyOrPartyRoleFVO) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *PartyOrPartyRoleFVO) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *PartyOrPartyRoleFVO) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### GetTitle

`func (o *PartyOrPartyRoleFVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PartyOrPartyRoleFVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PartyOrPartyRoleFVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PartyOrPartyRoleFVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *PartyOrPartyRoleFVO) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *PartyOrPartyRoleFVO) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *PartyOrPartyRoleFVO) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *PartyOrPartyRoleFVO) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *PartyOrPartyRoleFVO) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *PartyOrPartyRoleFVO) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *PartyOrPartyRoleFVO) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *PartyOrPartyRoleFVO) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *PartyOrPartyRoleFVO) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *PartyOrPartyRoleFVO) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *PartyOrPartyRoleFVO) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *PartyOrPartyRoleFVO) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *PartyOrPartyRoleFVO) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *PartyOrPartyRoleFVO) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *PartyOrPartyRoleFVO) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *PartyOrPartyRoleFVO) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetLegalName

`func (o *PartyOrPartyRoleFVO) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *PartyOrPartyRoleFVO) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *PartyOrPartyRoleFVO) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *PartyOrPartyRoleFVO) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *PartyOrPartyRoleFVO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PartyOrPartyRoleFVO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PartyOrPartyRoleFVO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PartyOrPartyRoleFVO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetFormattedName

`func (o *PartyOrPartyRoleFVO) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *PartyOrPartyRoleFVO) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *PartyOrPartyRoleFVO) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *PartyOrPartyRoleFVO) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetLocation

`func (o *PartyOrPartyRoleFVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PartyOrPartyRoleFVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PartyOrPartyRoleFVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PartyOrPartyRoleFVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *PartyOrPartyRoleFVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyOrPartyRoleFVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyOrPartyRoleFVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyOrPartyRoleFVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *PartyOrPartyRoleFVO) GetOtherName() []OtherNameOrganizationFVO`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *PartyOrPartyRoleFVO) GetOtherNameOk() (*[]OtherNameOrganizationFVO, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *PartyOrPartyRoleFVO) SetOtherName(v []OtherNameOrganizationFVO)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *PartyOrPartyRoleFVO) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetIndividualIdentification

`func (o *PartyOrPartyRoleFVO) GetIndividualIdentification() []IndividualIdentificationFVO`

GetIndividualIdentification returns the IndividualIdentification field if non-nil, zero value otherwise.

### GetIndividualIdentificationOk

`func (o *PartyOrPartyRoleFVO) GetIndividualIdentificationOk() (*[]IndividualIdentificationFVO, bool)`

GetIndividualIdentificationOk returns a tuple with the IndividualIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentification

`func (o *PartyOrPartyRoleFVO) SetIndividualIdentification(v []IndividualIdentificationFVO)`

SetIndividualIdentification sets IndividualIdentification field to given value.

### HasIndividualIdentification

`func (o *PartyOrPartyRoleFVO) HasIndividualIdentification() bool`

HasIndividualIdentification returns a boolean if a field has been set.

### GetDisability

`func (o *PartyOrPartyRoleFVO) GetDisability() []Disability`

GetDisability returns the Disability field if non-nil, zero value otherwise.

### GetDisabilityOk

`func (o *PartyOrPartyRoleFVO) GetDisabilityOk() (*[]Disability, bool)`

GetDisabilityOk returns a tuple with the Disability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisability

`func (o *PartyOrPartyRoleFVO) SetDisability(v []Disability)`

SetDisability sets Disability field to given value.

### HasDisability

`func (o *PartyOrPartyRoleFVO) HasDisability() bool`

HasDisability returns a boolean if a field has been set.

### GetLanguageAbility

`func (o *PartyOrPartyRoleFVO) GetLanguageAbility() []LanguageAbility`

GetLanguageAbility returns the LanguageAbility field if non-nil, zero value otherwise.

### GetLanguageAbilityOk

`func (o *PartyOrPartyRoleFVO) GetLanguageAbilityOk() (*[]LanguageAbility, bool)`

GetLanguageAbilityOk returns a tuple with the LanguageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageAbility

`func (o *PartyOrPartyRoleFVO) SetLanguageAbility(v []LanguageAbility)`

SetLanguageAbility sets LanguageAbility field to given value.

### HasLanguageAbility

`func (o *PartyOrPartyRoleFVO) HasLanguageAbility() bool`

HasLanguageAbility returns a boolean if a field has been set.

### GetSkill

`func (o *PartyOrPartyRoleFVO) GetSkill() []Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PartyOrPartyRoleFVO) GetSkillOk() (*[]Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PartyOrPartyRoleFVO) SetSkill(v []Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PartyOrPartyRoleFVO) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetFamilyName

`func (o *PartyOrPartyRoleFVO) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *PartyOrPartyRoleFVO) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *PartyOrPartyRoleFVO) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetGivenName

`func (o *PartyOrPartyRoleFVO) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *PartyOrPartyRoleFVO) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *PartyOrPartyRoleFVO) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetIsLegalEntity

`func (o *PartyOrPartyRoleFVO) GetIsLegalEntity() bool`

GetIsLegalEntity returns the IsLegalEntity field if non-nil, zero value otherwise.

### GetIsLegalEntityOk

`func (o *PartyOrPartyRoleFVO) GetIsLegalEntityOk() (*bool, bool)`

GetIsLegalEntityOk returns a tuple with the IsLegalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalEntity

`func (o *PartyOrPartyRoleFVO) SetIsLegalEntity(v bool)`

SetIsLegalEntity sets IsLegalEntity field to given value.

### HasIsLegalEntity

`func (o *PartyOrPartyRoleFVO) HasIsLegalEntity() bool`

HasIsLegalEntity returns a boolean if a field has been set.

### GetIsHeadOffice

`func (o *PartyOrPartyRoleFVO) GetIsHeadOffice() bool`

GetIsHeadOffice returns the IsHeadOffice field if non-nil, zero value otherwise.

### GetIsHeadOfficeOk

`func (o *PartyOrPartyRoleFVO) GetIsHeadOfficeOk() (*bool, bool)`

GetIsHeadOfficeOk returns a tuple with the IsHeadOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadOffice

`func (o *PartyOrPartyRoleFVO) SetIsHeadOffice(v bool)`

SetIsHeadOffice sets IsHeadOffice field to given value.

### HasIsHeadOffice

`func (o *PartyOrPartyRoleFVO) HasIsHeadOffice() bool`

HasIsHeadOffice returns a boolean if a field has been set.

### GetOrganizationType

`func (o *PartyOrPartyRoleFVO) GetOrganizationType() string`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *PartyOrPartyRoleFVO) GetOrganizationTypeOk() (*string, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *PartyOrPartyRoleFVO) SetOrganizationType(v string)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *PartyOrPartyRoleFVO) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetExistsDuring

`func (o *PartyOrPartyRoleFVO) GetExistsDuring() TimePeriod`

GetExistsDuring returns the ExistsDuring field if non-nil, zero value otherwise.

### GetExistsDuringOk

`func (o *PartyOrPartyRoleFVO) GetExistsDuringOk() (*TimePeriod, bool)`

GetExistsDuringOk returns a tuple with the ExistsDuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsDuring

`func (o *PartyOrPartyRoleFVO) SetExistsDuring(v TimePeriod)`

SetExistsDuring sets ExistsDuring field to given value.

### HasExistsDuring

`func (o *PartyOrPartyRoleFVO) HasExistsDuring() bool`

HasExistsDuring returns a boolean if a field has been set.

### GetNameType

`func (o *PartyOrPartyRoleFVO) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *PartyOrPartyRoleFVO) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *PartyOrPartyRoleFVO) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *PartyOrPartyRoleFVO) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetOrganizationIdentification

`func (o *PartyOrPartyRoleFVO) GetOrganizationIdentification() []OrganizationIdentificationFVO`

GetOrganizationIdentification returns the OrganizationIdentification field if non-nil, zero value otherwise.

### GetOrganizationIdentificationOk

`func (o *PartyOrPartyRoleFVO) GetOrganizationIdentificationOk() (*[]OrganizationIdentificationFVO, bool)`

GetOrganizationIdentificationOk returns a tuple with the OrganizationIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentification

`func (o *PartyOrPartyRoleFVO) SetOrganizationIdentification(v []OrganizationIdentificationFVO)`

SetOrganizationIdentification sets OrganizationIdentification field to given value.

### HasOrganizationIdentification

`func (o *PartyOrPartyRoleFVO) HasOrganizationIdentification() bool`

HasOrganizationIdentification returns a boolean if a field has been set.

### GetOrganizationChildRelationship

`func (o *PartyOrPartyRoleFVO) GetOrganizationChildRelationship() []OrganizationChildRelationshipFVO`

GetOrganizationChildRelationship returns the OrganizationChildRelationship field if non-nil, zero value otherwise.

### GetOrganizationChildRelationshipOk

`func (o *PartyOrPartyRoleFVO) GetOrganizationChildRelationshipOk() (*[]OrganizationChildRelationshipFVO, bool)`

GetOrganizationChildRelationshipOk returns a tuple with the OrganizationChildRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationChildRelationship

`func (o *PartyOrPartyRoleFVO) SetOrganizationChildRelationship(v []OrganizationChildRelationshipFVO)`

SetOrganizationChildRelationship sets OrganizationChildRelationship field to given value.

### HasOrganizationChildRelationship

`func (o *PartyOrPartyRoleFVO) HasOrganizationChildRelationship() bool`

HasOrganizationChildRelationship returns a boolean if a field has been set.

### GetOrganizationParentRelationship

`func (o *PartyOrPartyRoleFVO) GetOrganizationParentRelationship() OrganizationParentRelationshipFVO`

GetOrganizationParentRelationship returns the OrganizationParentRelationship field if non-nil, zero value otherwise.

### GetOrganizationParentRelationshipOk

`func (o *PartyOrPartyRoleFVO) GetOrganizationParentRelationshipOk() (*OrganizationParentRelationshipFVO, bool)`

GetOrganizationParentRelationshipOk returns a tuple with the OrganizationParentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationParentRelationship

`func (o *PartyOrPartyRoleFVO) SetOrganizationParentRelationship(v OrganizationParentRelationshipFVO)`

SetOrganizationParentRelationship sets OrganizationParentRelationship field to given value.

### HasOrganizationParentRelationship

`func (o *PartyOrPartyRoleFVO) HasOrganizationParentRelationship() bool`

HasOrganizationParentRelationship returns a boolean if a field has been set.

### GetTradingName

`func (o *PartyOrPartyRoleFVO) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *PartyOrPartyRoleFVO) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *PartyOrPartyRoleFVO) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *PartyOrPartyRoleFVO) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetDescription

`func (o *PartyOrPartyRoleFVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyOrPartyRoleFVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyOrPartyRoleFVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyOrPartyRoleFVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyOrPartyRoleFVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyOrPartyRoleFVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyOrPartyRoleFVO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyOrPartyRoleFVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyOrPartyRoleFVO) GetEngagedParty() PartyRefFVO`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyOrPartyRoleFVO) GetEngagedPartyOk() (*PartyRefFVO, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyOrPartyRoleFVO) SetEngagedParty(v PartyRefFVO)`

SetEngagedParty sets EngagedParty field to given value.


### GetPartyRoleSpecification

`func (o *PartyOrPartyRoleFVO) GetPartyRoleSpecification() PartyRoleSpecificationRefFVO`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyOrPartyRoleFVO) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRefFVO, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyOrPartyRoleFVO) SetPartyRoleSpecification(v PartyRoleSpecificationRefFVO)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyOrPartyRoleFVO) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyOrPartyRoleFVO) GetCharacteristic() []CharacteristicFVO`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyOrPartyRoleFVO) GetCharacteristicOk() (*[]CharacteristicFVO, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyOrPartyRoleFVO) SetCharacteristic(v []CharacteristicFVO)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyOrPartyRoleFVO) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyOrPartyRoleFVO) GetAccount() []AccountRefFVO`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyOrPartyRoleFVO) GetAccountOk() (*[]AccountRefFVO, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyOrPartyRoleFVO) SetAccount(v []AccountRefFVO)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyOrPartyRoleFVO) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyOrPartyRoleFVO) GetAgreement() []AgreementRefFVO`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyOrPartyRoleFVO) GetAgreementOk() (*[]AgreementRefFVO, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyOrPartyRoleFVO) SetAgreement(v []AgreementRefFVO)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyOrPartyRoleFVO) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyOrPartyRoleFVO) GetPaymentMethod() []PaymentMethodRefFVO`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyOrPartyRoleFVO) GetPaymentMethodOk() (*[]PaymentMethodRefFVO, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyOrPartyRoleFVO) SetPaymentMethod(v []PaymentMethodRefFVO)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyOrPartyRoleFVO) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyOrPartyRoleFVO) GetCreditProfile() []CreditProfileFVO`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyOrPartyRoleFVO) GetCreditProfileOk() (*[]CreditProfileFVO, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyOrPartyRoleFVO) SetCreditProfile(v []CreditProfileFVO)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyOrPartyRoleFVO) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyOrPartyRoleFVO) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyOrPartyRoleFVO) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyOrPartyRoleFVO) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyOrPartyRoleFVO) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyOrPartyRoleFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyOrPartyRoleFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyOrPartyRoleFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyOrPartyRoleFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


