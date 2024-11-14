/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Skill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Skill{}

// Skill Skills evaluated for an individual with a level and possibly with a limited validity when an obsolescence is defined (Ex: the first-aid certificate first level is limited to one year and an update training is required each year to keep the level).
type Skill struct {
	// Code of the skill
	SkillCode *string `json:"skillCode,omitempty"`
	// Name of the skill, such as Java language
	SkillName *string `json:"skillName,omitempty"`
	// Level of expertise in a skill evaluated for an individual
	EvaluatedLevel *string `json:"evaluatedLevel,omitempty"`
	// A free text comment linked to the evaluation done
	Comment *string `json:"comment,omitempty"`
	ValidFor *TimePeriod `json:"validFor,omitempty"`
}

// NewSkill instantiates a new Skill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkill() *Skill {
	this := Skill{}
	return &this
}

// NewSkillWithDefaults instantiates a new Skill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillWithDefaults() *Skill {
	this := Skill{}
	return &this
}

// GetSkillCode returns the SkillCode field value if set, zero value otherwise.
func (o *Skill) GetSkillCode() string {
	if o == nil || IsNil(o.SkillCode) {
		var ret string
		return ret
	}
	return *o.SkillCode
}

// GetSkillCodeOk returns a tuple with the SkillCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetSkillCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SkillCode) {
		return nil, false
	}
	return o.SkillCode, true
}

// HasSkillCode returns a boolean if a field has been set.
func (o *Skill) HasSkillCode() bool {
	if o != nil && !IsNil(o.SkillCode) {
		return true
	}

	return false
}

// SetSkillCode gets a reference to the given string and assigns it to the SkillCode field.
func (o *Skill) SetSkillCode(v string) {
	o.SkillCode = &v
}

// GetSkillName returns the SkillName field value if set, zero value otherwise.
func (o *Skill) GetSkillName() string {
	if o == nil || IsNil(o.SkillName) {
		var ret string
		return ret
	}
	return *o.SkillName
}

// GetSkillNameOk returns a tuple with the SkillName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetSkillNameOk() (*string, bool) {
	if o == nil || IsNil(o.SkillName) {
		return nil, false
	}
	return o.SkillName, true
}

// HasSkillName returns a boolean if a field has been set.
func (o *Skill) HasSkillName() bool {
	if o != nil && !IsNil(o.SkillName) {
		return true
	}

	return false
}

// SetSkillName gets a reference to the given string and assigns it to the SkillName field.
func (o *Skill) SetSkillName(v string) {
	o.SkillName = &v
}

// GetEvaluatedLevel returns the EvaluatedLevel field value if set, zero value otherwise.
func (o *Skill) GetEvaluatedLevel() string {
	if o == nil || IsNil(o.EvaluatedLevel) {
		var ret string
		return ret
	}
	return *o.EvaluatedLevel
}

// GetEvaluatedLevelOk returns a tuple with the EvaluatedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetEvaluatedLevelOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluatedLevel) {
		return nil, false
	}
	return o.EvaluatedLevel, true
}

// HasEvaluatedLevel returns a boolean if a field has been set.
func (o *Skill) HasEvaluatedLevel() bool {
	if o != nil && !IsNil(o.EvaluatedLevel) {
		return true
	}

	return false
}

// SetEvaluatedLevel gets a reference to the given string and assigns it to the EvaluatedLevel field.
func (o *Skill) SetEvaluatedLevel(v string) {
	o.EvaluatedLevel = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Skill) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Skill) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Skill) SetComment(v string) {
	o.Comment = &v
}

// GetValidFor returns the ValidFor field value if set, zero value otherwise.
func (o *Skill) GetValidFor() TimePeriod {
	if o == nil || IsNil(o.ValidFor) {
		var ret TimePeriod
		return ret
	}
	return *o.ValidFor
}

// GetValidForOk returns a tuple with the ValidFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetValidForOk() (*TimePeriod, bool) {
	if o == nil || IsNil(o.ValidFor) {
		return nil, false
	}
	return o.ValidFor, true
}

// HasValidFor returns a boolean if a field has been set.
func (o *Skill) HasValidFor() bool {
	if o != nil && !IsNil(o.ValidFor) {
		return true
	}

	return false
}

// SetValidFor gets a reference to the given TimePeriod and assigns it to the ValidFor field.
func (o *Skill) SetValidFor(v TimePeriod) {
	o.ValidFor = &v
}

func (o Skill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Skill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkillCode) {
		toSerialize["skillCode"] = o.SkillCode
	}
	if !IsNil(o.SkillName) {
		toSerialize["skillName"] = o.SkillName
	}
	if !IsNil(o.EvaluatedLevel) {
		toSerialize["evaluatedLevel"] = o.EvaluatedLevel
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ValidFor) {
		toSerialize["validFor"] = o.ValidFor
	}
	return toSerialize, nil
}

type NullableSkill struct {
	value *Skill
	isSet bool
}

func (v NullableSkill) Get() *Skill {
	return v.value
}

func (v *NullableSkill) Set(val *Skill) {
	v.value = val
	v.isSet = true
}

func (v NullableSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkill(val *Skill) *NullableSkill {
	return &NullableSkill{value: val, isSet: true}
}

func (v NullableSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

