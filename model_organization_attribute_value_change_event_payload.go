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

// checks if the OrganizationAttributeValueChangeEventPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAttributeValueChangeEventPayload{}

// OrganizationAttributeValueChangeEventPayload OrganizationAttributeValueChangeEventPayload generic structure
type OrganizationAttributeValueChangeEventPayload struct {
	Organization *Organization `json:"organization,omitempty"`
}

// NewOrganizationAttributeValueChangeEventPayload instantiates a new OrganizationAttributeValueChangeEventPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAttributeValueChangeEventPayload() *OrganizationAttributeValueChangeEventPayload {
	this := OrganizationAttributeValueChangeEventPayload{}
	return &this
}

// NewOrganizationAttributeValueChangeEventPayloadWithDefaults instantiates a new OrganizationAttributeValueChangeEventPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAttributeValueChangeEventPayloadWithDefaults() *OrganizationAttributeValueChangeEventPayload {
	this := OrganizationAttributeValueChangeEventPayload{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationAttributeValueChangeEventPayload) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAttributeValueChangeEventPayload) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationAttributeValueChangeEventPayload) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *OrganizationAttributeValueChangeEventPayload) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o OrganizationAttributeValueChangeEventPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAttributeValueChangeEventPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableOrganizationAttributeValueChangeEventPayload struct {
	value *OrganizationAttributeValueChangeEventPayload
	isSet bool
}

func (v NullableOrganizationAttributeValueChangeEventPayload) Get() *OrganizationAttributeValueChangeEventPayload {
	return v.value
}

func (v *NullableOrganizationAttributeValueChangeEventPayload) Set(val *OrganizationAttributeValueChangeEventPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAttributeValueChangeEventPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAttributeValueChangeEventPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAttributeValueChangeEventPayload(val *OrganizationAttributeValueChangeEventPayload) *NullableOrganizationAttributeValueChangeEventPayload {
	return &NullableOrganizationAttributeValueChangeEventPayload{value: val, isSet: true}
}

func (v NullableOrganizationAttributeValueChangeEventPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAttributeValueChangeEventPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


