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

// checks if the ProducerMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProducerMVO{}

// ProducerMVO struct for ProducerMVO
type ProducerMVO struct {
	PartyRoleMVO
}

type _ProducerMVO ProducerMVO

// NewProducerMVO instantiates a new ProducerMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProducerMVO(type_ string, name string, engagedParty PartyRefMVO) *ProducerMVO {
	this := ProducerMVO{}
	this.Type = type_
	this.Name = name
	this.EngagedParty = engagedParty
	return &this
}

// NewProducerMVOWithDefaults instantiates a new ProducerMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProducerMVOWithDefaults() *ProducerMVO {
	this := ProducerMVO{}
	return &this
}

func (o ProducerMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProducerMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *ProducerMVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
		"name",
		"engagedParty",
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

	varProducerMVO := _ProducerMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProducerMVO)

	if err != nil {
		return err
	}

	*o = ProducerMVO(varProducerMVO)

	return err
}

type NullableProducerMVO struct {
	value *ProducerMVO
	isSet bool
}

func (v NullableProducerMVO) Get() *ProducerMVO {
	return v.value
}

func (v *NullableProducerMVO) Set(val *ProducerMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableProducerMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableProducerMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProducerMVO(val *ProducerMVO) *NullableProducerMVO {
	return &NullableProducerMVO{value: val, isSet: true}
}

func (v NullableProducerMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProducerMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


