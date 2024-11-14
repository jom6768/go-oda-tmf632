# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillCode** | Pointer to **string** | Code of the skill | [optional] 
**SkillName** | Pointer to **string** | Name of the skill, such as Java language | [optional] 
**EvaluatedLevel** | Pointer to **string** | Level of expertise in a skill evaluated for an individual | [optional] 
**Comment** | Pointer to **string** | A free text comment linked to the evaluation done | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewSkill

`func NewSkill() *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillCode

`func (o *Skill) GetSkillCode() string`

GetSkillCode returns the SkillCode field if non-nil, zero value otherwise.

### GetSkillCodeOk

`func (o *Skill) GetSkillCodeOk() (*string, bool)`

GetSkillCodeOk returns a tuple with the SkillCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillCode

`func (o *Skill) SetSkillCode(v string)`

SetSkillCode sets SkillCode field to given value.

### HasSkillCode

`func (o *Skill) HasSkillCode() bool`

HasSkillCode returns a boolean if a field has been set.

### GetSkillName

`func (o *Skill) GetSkillName() string`

GetSkillName returns the SkillName field if non-nil, zero value otherwise.

### GetSkillNameOk

`func (o *Skill) GetSkillNameOk() (*string, bool)`

GetSkillNameOk returns a tuple with the SkillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillName

`func (o *Skill) SetSkillName(v string)`

SetSkillName sets SkillName field to given value.

### HasSkillName

`func (o *Skill) HasSkillName() bool`

HasSkillName returns a boolean if a field has been set.

### GetEvaluatedLevel

`func (o *Skill) GetEvaluatedLevel() string`

GetEvaluatedLevel returns the EvaluatedLevel field if non-nil, zero value otherwise.

### GetEvaluatedLevelOk

`func (o *Skill) GetEvaluatedLevelOk() (*string, bool)`

GetEvaluatedLevelOk returns a tuple with the EvaluatedLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedLevel

`func (o *Skill) SetEvaluatedLevel(v string)`

SetEvaluatedLevel sets EvaluatedLevel field to given value.

### HasEvaluatedLevel

`func (o *Skill) HasEvaluatedLevel() bool`

HasEvaluatedLevel returns a boolean if a field has been set.

### GetComment

`func (o *Skill) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Skill) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Skill) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Skill) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetValidFor

`func (o *Skill) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *Skill) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *Skill) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *Skill) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


