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

// checks if the IndividualAttributeValueChangeEventPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndividualAttributeValueChangeEventPayload{}

// IndividualAttributeValueChangeEventPayload IndividualAttributeValueChangeEventPayload generic structure
type IndividualAttributeValueChangeEventPayload struct {
	Individual *Individual `json:"individual,omitempty"`
}

// NewIndividualAttributeValueChangeEventPayload instantiates a new IndividualAttributeValueChangeEventPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualAttributeValueChangeEventPayload() *IndividualAttributeValueChangeEventPayload {
	this := IndividualAttributeValueChangeEventPayload{}
	return &this
}

// NewIndividualAttributeValueChangeEventPayloadWithDefaults instantiates a new IndividualAttributeValueChangeEventPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualAttributeValueChangeEventPayloadWithDefaults() *IndividualAttributeValueChangeEventPayload {
	this := IndividualAttributeValueChangeEventPayload{}
	return &this
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *IndividualAttributeValueChangeEventPayload) GetIndividual() Individual {
	if o == nil || IsNil(o.Individual) {
		var ret Individual
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualAttributeValueChangeEventPayload) GetIndividualOk() (*Individual, bool) {
	if o == nil || IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *IndividualAttributeValueChangeEventPayload) HasIndividual() bool {
	if o != nil && !IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given Individual and assigns it to the Individual field.
func (o *IndividualAttributeValueChangeEventPayload) SetIndividual(v Individual) {
	o.Individual = &v
}

func (o IndividualAttributeValueChangeEventPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualAttributeValueChangeEventPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	return toSerialize, nil
}

type NullableIndividualAttributeValueChangeEventPayload struct {
	value *IndividualAttributeValueChangeEventPayload
	isSet bool
}

func (v NullableIndividualAttributeValueChangeEventPayload) Get() *IndividualAttributeValueChangeEventPayload {
	return v.value
}

func (v *NullableIndividualAttributeValueChangeEventPayload) Set(val *IndividualAttributeValueChangeEventPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualAttributeValueChangeEventPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualAttributeValueChangeEventPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualAttributeValueChangeEventPayload(val *IndividualAttributeValueChangeEventPayload) *NullableIndividualAttributeValueChangeEventPayload {
	return &NullableIndividualAttributeValueChangeEventPayload{value: val, isSet: true}
}

func (v NullableIndividualAttributeValueChangeEventPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualAttributeValueChangeEventPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

