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

// checks if the OrganizationStateChangeEventPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationStateChangeEventPayload{}

// OrganizationStateChangeEventPayload OrganizationStateChangeEventPayload generic structure
type OrganizationStateChangeEventPayload struct {
	Organization *Organization `json:"organization,omitempty"`
}

// NewOrganizationStateChangeEventPayload instantiates a new OrganizationStateChangeEventPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationStateChangeEventPayload() *OrganizationStateChangeEventPayload {
	this := OrganizationStateChangeEventPayload{}
	return &this
}

// NewOrganizationStateChangeEventPayloadWithDefaults instantiates a new OrganizationStateChangeEventPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationStateChangeEventPayloadWithDefaults() *OrganizationStateChangeEventPayload {
	this := OrganizationStateChangeEventPayload{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationStateChangeEventPayload) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationStateChangeEventPayload) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationStateChangeEventPayload) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *OrganizationStateChangeEventPayload) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o OrganizationStateChangeEventPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationStateChangeEventPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableOrganizationStateChangeEventPayload struct {
	value *OrganizationStateChangeEventPayload
	isSet bool
}

func (v NullableOrganizationStateChangeEventPayload) Get() *OrganizationStateChangeEventPayload {
	return v.value
}

func (v *NullableOrganizationStateChangeEventPayload) Set(val *OrganizationStateChangeEventPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationStateChangeEventPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationStateChangeEventPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationStateChangeEventPayload(val *OrganizationStateChangeEventPayload) *NullableOrganizationStateChangeEventPayload {
	return &NullableOrganizationStateChangeEventPayload{value: val, isSet: true}
}

func (v NullableOrganizationStateChangeEventPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationStateChangeEventPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

