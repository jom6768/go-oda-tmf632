# IndividualMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Name** | Pointer to **string** | Full name flatten (first, middle, and last names) - this is the name that is expected to be presented in reference data types such as PartyRef, RelatedParty, etc. that refer to Individual | [optional] 
**FormattedName** | Pointer to **string** | A fully formatted name in one string with all of its pieces in their proper place and all of the necessary punctuation. Useful for specific contexts (Chinese, Japanese, Korean) | [optional] 
**Location** | Pointer to **string** | Temporary current location of the individual (may be used if the individual has approved its sharing) | [optional] 
**Status** | Pointer to [**IndividualStateType**](IndividualStateType.md) |  | [optional] 
**OtherName** | Pointer to [**[]OtherNameIndividual**](OtherNameIndividual.md) | List of other names by which this individual is known | [optional] 
**IndividualIdentification** | Pointer to [**[]IndividualIdentificationMVO**](IndividualIdentificationMVO.md) | List of official identifications issued to the individual, such as passport, driving licence, social security number | [optional] 
**Disability** | Pointer to [**[]Disability**](Disability.md) | List of disabilities suffered by the individual | [optional] 
**LanguageAbility** | Pointer to [**[]LanguageAbility**](LanguageAbility.md) | List of national languages known by the individual | [optional] 
**Skill** | Pointer to [**[]Skill**](Skill.md) | List of skills exhibited by the individual | [optional] 
**FamilyName** | Pointer to **string** | Contains the non-chosen or inherited name. Also known as last name in the Western context | [optional] 
**GivenName** | Pointer to **string** | First name of the individual | [optional] 

## Methods

### NewIndividualMVO

`func NewIndividualMVO() *IndividualMVO`

NewIndividualMVO instantiates a new IndividualMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualMVOWithDefaults

`func NewIndividualMVOWithDefaults() *IndividualMVO`

NewIndividualMVOWithDefaults instantiates a new IndividualMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGender

`func (o *IndividualMVO) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IndividualMVO) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IndividualMVO) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IndividualMVO) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *IndividualMVO) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *IndividualMVO) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *IndividualMVO) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *IndividualMVO) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *IndividualMVO) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *IndividualMVO) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *IndividualMVO) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *IndividualMVO) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetNationality

`func (o *IndividualMVO) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *IndividualMVO) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *IndividualMVO) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *IndividualMVO) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *IndividualMVO) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *IndividualMVO) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *IndividualMVO) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *IndividualMVO) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetBirthDate

`func (o *IndividualMVO) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *IndividualMVO) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *IndividualMVO) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *IndividualMVO) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetDeathDate

`func (o *IndividualMVO) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *IndividualMVO) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *IndividualMVO) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *IndividualMVO) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### GetTitle

`func (o *IndividualMVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualMVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualMVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IndividualMVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *IndividualMVO) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *IndividualMVO) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *IndividualMVO) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *IndividualMVO) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *IndividualMVO) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *IndividualMVO) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *IndividualMVO) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *IndividualMVO) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *IndividualMVO) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *IndividualMVO) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *IndividualMVO) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *IndividualMVO) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *IndividualMVO) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *IndividualMVO) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *IndividualMVO) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *IndividualMVO) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetLegalName

`func (o *IndividualMVO) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *IndividualMVO) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *IndividualMVO) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *IndividualMVO) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *IndividualMVO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IndividualMVO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IndividualMVO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IndividualMVO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetName

`func (o *IndividualMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndividualMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndividualMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndividualMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormattedName

`func (o *IndividualMVO) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *IndividualMVO) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *IndividualMVO) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *IndividualMVO) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetLocation

`func (o *IndividualMVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IndividualMVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IndividualMVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IndividualMVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *IndividualMVO) GetStatus() IndividualStateType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndividualMVO) GetStatusOk() (*IndividualStateType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndividualMVO) SetStatus(v IndividualStateType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndividualMVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *IndividualMVO) GetOtherName() []OtherNameIndividual`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *IndividualMVO) GetOtherNameOk() (*[]OtherNameIndividual, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *IndividualMVO) SetOtherName(v []OtherNameIndividual)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *IndividualMVO) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetIndividualIdentification

`func (o *IndividualMVO) GetIndividualIdentification() []IndividualIdentificationMVO`

GetIndividualIdentification returns the IndividualIdentification field if non-nil, zero value otherwise.

### GetIndividualIdentificationOk

`func (o *IndividualMVO) GetIndividualIdentificationOk() (*[]IndividualIdentificationMVO, bool)`

GetIndividualIdentificationOk returns a tuple with the IndividualIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentification

`func (o *IndividualMVO) SetIndividualIdentification(v []IndividualIdentificationMVO)`

SetIndividualIdentification sets IndividualIdentification field to given value.

### HasIndividualIdentification

`func (o *IndividualMVO) HasIndividualIdentification() bool`

HasIndividualIdentification returns a boolean if a field has been set.

### GetDisability

`func (o *IndividualMVO) GetDisability() []Disability`

GetDisability returns the Disability field if non-nil, zero value otherwise.

### GetDisabilityOk

`func (o *IndividualMVO) GetDisabilityOk() (*[]Disability, bool)`

GetDisabilityOk returns a tuple with the Disability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisability

`func (o *IndividualMVO) SetDisability(v []Disability)`

SetDisability sets Disability field to given value.

### HasDisability

`func (o *IndividualMVO) HasDisability() bool`

HasDisability returns a boolean if a field has been set.

### GetLanguageAbility

`func (o *IndividualMVO) GetLanguageAbility() []LanguageAbility`

GetLanguageAbility returns the LanguageAbility field if non-nil, zero value otherwise.

### GetLanguageAbilityOk

`func (o *IndividualMVO) GetLanguageAbilityOk() (*[]LanguageAbility, bool)`

GetLanguageAbilityOk returns a tuple with the LanguageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageAbility

`func (o *IndividualMVO) SetLanguageAbility(v []LanguageAbility)`

SetLanguageAbility sets LanguageAbility field to given value.

### HasLanguageAbility

`func (o *IndividualMVO) HasLanguageAbility() bool`

HasLanguageAbility returns a boolean if a field has been set.

### GetSkill

`func (o *IndividualMVO) GetSkill() []Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *IndividualMVO) GetSkillOk() (*[]Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *IndividualMVO) SetSkill(v []Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *IndividualMVO) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetFamilyName

`func (o *IndividualMVO) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IndividualMVO) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IndividualMVO) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *IndividualMVO) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *IndividualMVO) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IndividualMVO) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IndividualMVO) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *IndividualMVO) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


