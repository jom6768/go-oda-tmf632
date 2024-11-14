# LanguageAbility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanguageCode** | Pointer to **string** | Language code (RFC 5646) | [optional] 
**LanguageName** | Pointer to **string** | Language name | [optional] 
**IsFavouriteLanguage** | Pointer to **bool** | A “true” value specifies whether the language is considered by the individual as his favourite one | [optional] 
**WritingProficiency** | Pointer to **string** | Writing proficiency evaluated for this language | [optional] 
**ReadingProficiency** | Pointer to **string** | Reading proficiency evaluated for this language | [optional] 
**SpeakingProficiency** | Pointer to **string** | Speaking proficiency evaluated for this language | [optional] 
**ListeningProficiency** | Pointer to **string** | Listening proficiency evaluated for this language | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewLanguageAbility

`func NewLanguageAbility() *LanguageAbility`

NewLanguageAbility instantiates a new LanguageAbility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageAbilityWithDefaults

`func NewLanguageAbilityWithDefaults() *LanguageAbility`

NewLanguageAbilityWithDefaults instantiates a new LanguageAbility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguageCode

`func (o *LanguageAbility) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *LanguageAbility) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *LanguageAbility) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *LanguageAbility) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetLanguageName

`func (o *LanguageAbility) GetLanguageName() string`

GetLanguageName returns the LanguageName field if non-nil, zero value otherwise.

### GetLanguageNameOk

`func (o *LanguageAbility) GetLanguageNameOk() (*string, bool)`

GetLanguageNameOk returns a tuple with the LanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageName

`func (o *LanguageAbility) SetLanguageName(v string)`

SetLanguageName sets LanguageName field to given value.

### HasLanguageName

`func (o *LanguageAbility) HasLanguageName() bool`

HasLanguageName returns a boolean if a field has been set.

### GetIsFavouriteLanguage

`func (o *LanguageAbility) GetIsFavouriteLanguage() bool`

GetIsFavouriteLanguage returns the IsFavouriteLanguage field if non-nil, zero value otherwise.

### GetIsFavouriteLanguageOk

`func (o *LanguageAbility) GetIsFavouriteLanguageOk() (*bool, bool)`

GetIsFavouriteLanguageOk returns a tuple with the IsFavouriteLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavouriteLanguage

`func (o *LanguageAbility) SetIsFavouriteLanguage(v bool)`

SetIsFavouriteLanguage sets IsFavouriteLanguage field to given value.

### HasIsFavouriteLanguage

`func (o *LanguageAbility) HasIsFavouriteLanguage() bool`

HasIsFavouriteLanguage returns a boolean if a field has been set.

### GetWritingProficiency

`func (o *LanguageAbility) GetWritingProficiency() string`

GetWritingProficiency returns the WritingProficiency field if non-nil, zero value otherwise.

### GetWritingProficiencyOk

`func (o *LanguageAbility) GetWritingProficiencyOk() (*string, bool)`

GetWritingProficiencyOk returns a tuple with the WritingProficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritingProficiency

`func (o *LanguageAbility) SetWritingProficiency(v string)`

SetWritingProficiency sets WritingProficiency field to given value.

### HasWritingProficiency

`func (o *LanguageAbility) HasWritingProficiency() bool`

HasWritingProficiency returns a boolean if a field has been set.

### GetReadingProficiency

`func (o *LanguageAbility) GetReadingProficiency() string`

GetReadingProficiency returns the ReadingProficiency field if non-nil, zero value otherwise.

### GetReadingProficiencyOk

`func (o *LanguageAbility) GetReadingProficiencyOk() (*string, bool)`

GetReadingProficiencyOk returns a tuple with the ReadingProficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingProficiency

`func (o *LanguageAbility) SetReadingProficiency(v string)`

SetReadingProficiency sets ReadingProficiency field to given value.

### HasReadingProficiency

`func (o *LanguageAbility) HasReadingProficiency() bool`

HasReadingProficiency returns a boolean if a field has been set.

### GetSpeakingProficiency

`func (o *LanguageAbility) GetSpeakingProficiency() string`

GetSpeakingProficiency returns the SpeakingProficiency field if non-nil, zero value otherwise.

### GetSpeakingProficiencyOk

`func (o *LanguageAbility) GetSpeakingProficiencyOk() (*string, bool)`

GetSpeakingProficiencyOk returns a tuple with the SpeakingProficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakingProficiency

`func (o *LanguageAbility) SetSpeakingProficiency(v string)`

SetSpeakingProficiency sets SpeakingProficiency field to given value.

### HasSpeakingProficiency

`func (o *LanguageAbility) HasSpeakingProficiency() bool`

HasSpeakingProficiency returns a boolean if a field has been set.

### GetListeningProficiency

`func (o *LanguageAbility) GetListeningProficiency() string`

GetListeningProficiency returns the ListeningProficiency field if non-nil, zero value otherwise.

### GetListeningProficiencyOk

`func (o *LanguageAbility) GetListeningProficiencyOk() (*string, bool)`

GetListeningProficiencyOk returns a tuple with the ListeningProficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeningProficiency

`func (o *LanguageAbility) SetListeningProficiency(v string)`

SetListeningProficiency sets ListeningProficiency field to given value.

### HasListeningProficiency

`func (o *LanguageAbility) HasListeningProficiency() bool`

HasListeningProficiency returns a boolean if a field has been set.

### GetValidFor

`func (o *LanguageAbility) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *LanguageAbility) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *LanguageAbility) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *LanguageAbility) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


