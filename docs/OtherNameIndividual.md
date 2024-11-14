# OtherNameIndividual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Use for titles (aristrocatic, social, ...): Pr, Dr, Sir,.... | [optional] 
**AristocraticTitle** | Pointer to **string** | e.g. Baron, Graf, Earl, etc. | [optional] 
**Generation** | Pointer to **string** | e.g. Sr, Jr, etc. | [optional] 
**GivenName** | Pointer to **string** | First name | [optional] 
**PreferredGivenName** | Pointer to **string** | Contains the chosen name by which the person prefers to be addressed. Note: This name may be a name other than a given name, such as a nickname | [optional] 
**FamilyNamePrefix** | Pointer to **string** | Family name prefix | [optional] 
**FamilyName** | Pointer to **string** | Contains the non-chosen or inherited name. Also known as last name in the Western context | [optional] 
**LegalName** | Pointer to **string** | Legal name or birth name (name one has for official purposes) | [optional] 
**MiddleName** | Pointer to **string** | Middle name or initial | [optional] 
**FullName** | Pointer to **string** | Full name flatten (first, middle, and last names) | [optional] 
**FormattedName** | Pointer to **string** | . A fully formatted name in one string with all of its pieces in their proper place and all of the necessary punctuation. Useful for specific contexts (Chinese, Japanese, Korean, etc.) | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewOtherNameIndividual

`func NewOtherNameIndividual() *OtherNameIndividual`

NewOtherNameIndividual instantiates a new OtherNameIndividual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherNameIndividualWithDefaults

`func NewOtherNameIndividualWithDefaults() *OtherNameIndividual`

NewOtherNameIndividualWithDefaults instantiates a new OtherNameIndividual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *OtherNameIndividual) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OtherNameIndividual) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OtherNameIndividual) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OtherNameIndividual) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAristocraticTitle

`func (o *OtherNameIndividual) GetAristocraticTitle() string`

GetAristocraticTitle returns the AristocraticTitle field if non-nil, zero value otherwise.

### GetAristocraticTitleOk

`func (o *OtherNameIndividual) GetAristocraticTitleOk() (*string, bool)`

GetAristocraticTitleOk returns a tuple with the AristocraticTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAristocraticTitle

`func (o *OtherNameIndividual) SetAristocraticTitle(v string)`

SetAristocraticTitle sets AristocraticTitle field to given value.

### HasAristocraticTitle

`func (o *OtherNameIndividual) HasAristocraticTitle() bool`

HasAristocraticTitle returns a boolean if a field has been set.

### GetGeneration

`func (o *OtherNameIndividual) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *OtherNameIndividual) GetGenerationOk() (*string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *OtherNameIndividual) SetGeneration(v string)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *OtherNameIndividual) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetGivenName

`func (o *OtherNameIndividual) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *OtherNameIndividual) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *OtherNameIndividual) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *OtherNameIndividual) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetPreferredGivenName

`func (o *OtherNameIndividual) GetPreferredGivenName() string`

GetPreferredGivenName returns the PreferredGivenName field if non-nil, zero value otherwise.

### GetPreferredGivenNameOk

`func (o *OtherNameIndividual) GetPreferredGivenNameOk() (*string, bool)`

GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredGivenName

`func (o *OtherNameIndividual) SetPreferredGivenName(v string)`

SetPreferredGivenName sets PreferredGivenName field to given value.

### HasPreferredGivenName

`func (o *OtherNameIndividual) HasPreferredGivenName() bool`

HasPreferredGivenName returns a boolean if a field has been set.

### GetFamilyNamePrefix

`func (o *OtherNameIndividual) GetFamilyNamePrefix() string`

GetFamilyNamePrefix returns the FamilyNamePrefix field if non-nil, zero value otherwise.

### GetFamilyNamePrefixOk

`func (o *OtherNameIndividual) GetFamilyNamePrefixOk() (*string, bool)`

GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyNamePrefix

`func (o *OtherNameIndividual) SetFamilyNamePrefix(v string)`

SetFamilyNamePrefix sets FamilyNamePrefix field to given value.

### HasFamilyNamePrefix

`func (o *OtherNameIndividual) HasFamilyNamePrefix() bool`

HasFamilyNamePrefix returns a boolean if a field has been set.

### GetFamilyName

`func (o *OtherNameIndividual) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *OtherNameIndividual) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *OtherNameIndividual) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *OtherNameIndividual) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetLegalName

`func (o *OtherNameIndividual) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OtherNameIndividual) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OtherNameIndividual) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *OtherNameIndividual) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetMiddleName

`func (o *OtherNameIndividual) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *OtherNameIndividual) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *OtherNameIndividual) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *OtherNameIndividual) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetFullName

`func (o *OtherNameIndividual) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OtherNameIndividual) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OtherNameIndividual) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OtherNameIndividual) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetFormattedName

`func (o *OtherNameIndividual) GetFormattedName() string`

GetFormattedName returns the FormattedName field if non-nil, zero value otherwise.

### GetFormattedNameOk

`func (o *OtherNameIndividual) GetFormattedNameOk() (*string, bool)`

GetFormattedNameOk returns a tuple with the FormattedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedName

`func (o *OtherNameIndividual) SetFormattedName(v string)`

SetFormattedName sets FormattedName field to given value.

### HasFormattedName

`func (o *OtherNameIndividual) HasFormattedName() bool`

HasFormattedName returns a boolean if a field has been set.

### GetValidFor

`func (o *OtherNameIndividual) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *OtherNameIndividual) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *OtherNameIndividual) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *OtherNameIndividual) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


