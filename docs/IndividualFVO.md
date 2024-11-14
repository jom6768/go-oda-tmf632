# IndividualFVO

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
**IndividualIdentification** | Pointer to [**[]IndividualIdentificationFVO**](IndividualIdentificationFVO.md) | List of official identifications issued to the individual, such as passport, driving licence, social security number | [optional] 
**Disability** | Pointer to [**[]Disability**](Disability.md) | List of disabilities suffered by the individual | [optional] 
**LanguageAbility** | Pointer to [**[]LanguageAbility**](LanguageAbility.md) | List of national languages known by the individual | [optional] 
**Skill** | Pointer to [**[]Skill**](Skill.md) | List of skills exhibited by the individual | [optional] 
**FamilyName** | **string** | Contains the non-chosen or inherited name. Also known as last name in the Western context | 
**GivenName** | **string** | First name of the individual | 

## Methods

### NewIndividualFVO

`func NewIndividualFVO(familyName string, givenName string, ) *IndividualFVO`

NewIndividualFVO instantiates a new IndividualFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualFVOWithDefaults

`func NewIndividualFVOWithDefaults() *IndividualFVO`

NewIndividualFVOWithDefaults instantiates a new IndividualFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGender

`func (o *IndividualFVO) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IndividualFVO) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IndividualFVO) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IndividualFVO) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *IndividualFVO) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *IndividualFVO) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *IndividualFVO) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *IndividualFVO) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *IndividualFVO) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *IndividualFVO) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *IndividualFVO) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *IndividualFVO) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetNationality

`func (o *IndividualFVO) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *IndividualFVO) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *IndividualFVO) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *IndividualFVO) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMaritalStatus

`func (o *IndividualFVO) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *IndividualFVO) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *IndividualFVO) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *IndividualFVO) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### GetBirthDate

`func (o *IndividualFVO) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *IndividualFVO) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *IndividualFVO) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *IndividualFVO) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetDeathDate

`func (o *IndividualFVO) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *IndividualFVO) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *IndividualFVO) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *IndividualFVO) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### GetTitle

`func (o *IndividualFVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualFVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualFVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IndividualFVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *IndividualFVO) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *IndividualFVO) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *IndividualFVO) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *IndividualFVO) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *IndividualFVO) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *IndividualFVO) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *IndividualFVO) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *IndividualFVO) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *IndividualFVO) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *IndividualFVO) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *IndividualFVO) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *IndividualFVO) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *IndividualFVO) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *IndividualFVO) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *IndividualFVO) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *IndividualFVO) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetLegalName

`func (o *IndividualFVO) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *IndividualFVO) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *IndividualFVO) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *IndividualFVO) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *IndividualFVO) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IndividualFVO) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IndividualFVO) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IndividualFVO) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetName

`func (o *IndividualFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndividualFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndividualFVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndividualFVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormattedName

`func (o *IndividualFVO) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *IndividualFVO) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *IndividualFVO) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *IndividualFVO) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetLocation

`func (o *IndividualFVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IndividualFVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IndividualFVO) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IndividualFVO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *IndividualFVO) GetStatus() IndividualStateType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IndividualFVO) GetStatusOk() (*IndividualStateType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IndividualFVO) SetStatus(v IndividualStateType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IndividualFVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *IndividualFVO) GetOtherName() []OtherNameIndividual`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *IndividualFVO) GetOtherNameOk() (*[]OtherNameIndividual, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *IndividualFVO) SetOtherName(v []OtherNameIndividual)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *IndividualFVO) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetIndividualIdentification

`func (o *IndividualFVO) GetIndividualIdentification() []IndividualIdentificationFVO`

GetIndividualIdentification returns the IndividualIdentification field if non-nil, zero value otherwise.

### GetIndividualIdentificationOk

`func (o *IndividualFVO) GetIndividualIdentificationOk() (*[]IndividualIdentificationFVO, bool)`

GetIndividualIdentificationOk returns a tuple with the IndividualIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualIdentification

`func (o *IndividualFVO) SetIndividualIdentification(v []IndividualIdentificationFVO)`

SetIndividualIdentification sets IndividualIdentification field to given value.

### HasIndividualIdentification

`func (o *IndividualFVO) HasIndividualIdentification() bool`

HasIndividualIdentification returns a boolean if a field has been set.

### GetDisability

`func (o *IndividualFVO) GetDisability() []Disability`

GetDisability returns the Disability field if non-nil, zero value otherwise.

### GetDisabilityOk

`func (o *IndividualFVO) GetDisabilityOk() (*[]Disability, bool)`

GetDisabilityOk returns a tuple with the Disability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisability

`func (o *IndividualFVO) SetDisability(v []Disability)`

SetDisability sets Disability field to given value.

### HasDisability

`func (o *IndividualFVO) HasDisability() bool`

HasDisability returns a boolean if a field has been set.

### GetLanguageAbility

`func (o *IndividualFVO) GetLanguageAbility() []LanguageAbility`

GetLanguageAbility returns the LanguageAbility field if non-nil, zero value otherwise.

### GetLanguageAbilityOk

`func (o *IndividualFVO) GetLanguageAbilityOk() (*[]LanguageAbility, bool)`

GetLanguageAbilityOk returns a tuple with the LanguageAbility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageAbility

`func (o *IndividualFVO) SetLanguageAbility(v []LanguageAbility)`

SetLanguageAbility sets LanguageAbility field to given value.

### HasLanguageAbility

`func (o *IndividualFVO) HasLanguageAbility() bool`

HasLanguageAbility returns a boolean if a field has been set.

### GetSkill

`func (o *IndividualFVO) GetSkill() []Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *IndividualFVO) GetSkillOk() (*[]Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *IndividualFVO) SetSkill(v []Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *IndividualFVO) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetFamilyName

`func (o *IndividualFVO) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IndividualFVO) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IndividualFVO) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetGivenName

`func (o *IndividualFVO) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IndividualFVO) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IndividualFVO) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


