/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SocialContactMediumMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SocialContactMediumMVO{}

// SocialContactMediumMVO struct for SocialContactMediumMVO
type SocialContactMediumMVO struct {
	ContactMediumMVO
	// Identifier as a member of a social network
	SocialNetworkId *string `json:"socialNetworkId,omitempty"`
}

type _SocialContactMediumMVO SocialContactMediumMVO

// NewSocialContactMediumMVO instantiates a new SocialContactMediumMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialContactMediumMVO(type_ string) *SocialContactMediumMVO {
	this := SocialContactMediumMVO{}
	this.Type = type_
	return &this
}

// NewSocialContactMediumMVOWithDefaults instantiates a new SocialContactMediumMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialContactMediumMVOWithDefaults() *SocialContactMediumMVO {
	this := SocialContactMediumMVO{}
	return &this
}

// GetSocialNetworkId returns the SocialNetworkId field value if set, zero value otherwise.
func (o *SocialContactMediumMVO) GetSocialNetworkId() string {
	if o == nil || IsNil(o.SocialNetworkId) {
		var ret string
		return ret
	}
	return *o.SocialNetworkId
}

// GetSocialNetworkIdOk returns a tuple with the SocialNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialContactMediumMVO) GetSocialNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.SocialNetworkId) {
		return nil, false
	}
	return o.SocialNetworkId, true
}

// HasSocialNetworkId returns a boolean if a field has been set.
func (o *SocialContactMediumMVO) HasSocialNetworkId() bool {
	if o != nil && !IsNil(o.SocialNetworkId) {
		return true
	}

	return false
}

// SetSocialNetworkId gets a reference to the given string and assigns it to the SocialNetworkId field.
func (o *SocialContactMediumMVO) SetSocialNetworkId(v string) {
	o.SocialNetworkId = &v
}

func (o SocialContactMediumMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SocialContactMediumMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContactMediumMVO, errContactMediumMVO := json.Marshal(o.ContactMediumMVO)
	if errContactMediumMVO != nil {
		return map[string]interface{}{}, errContactMediumMVO
	}
	errContactMediumMVO = json.Unmarshal([]byte(serializedContactMediumMVO), &toSerialize)
	if errContactMediumMVO != nil {
		return map[string]interface{}{}, errContactMediumMVO
	}
	if !IsNil(o.SocialNetworkId) {
		toSerialize["socialNetworkId"] = o.SocialNetworkId
	}
	return toSerialize, nil
}

func (o *SocialContactMediumMVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSocialContactMediumMVO := _SocialContactMediumMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSocialContactMediumMVO)

	if err != nil {
		return err
	}

	*o = SocialContactMediumMVO(varSocialContactMediumMVO)

	return err
}

type NullableSocialContactMediumMVO struct {
	value *SocialContactMediumMVO
	isSet bool
}

func (v NullableSocialContactMediumMVO) Get() *SocialContactMediumMVO {
	return v.value
}

func (v *NullableSocialContactMediumMVO) Set(val *SocialContactMediumMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialContactMediumMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialContactMediumMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialContactMediumMVO(val *SocialContactMediumMVO) *NullableSocialContactMediumMVO {
	return &NullableSocialContactMediumMVO{value: val, isSet: true}
}

func (v NullableSocialContactMediumMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialContactMediumMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


