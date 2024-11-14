/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// OrganizationStateType Valid values for the lifecycle state of the organization
type OrganizationStateType string

// List of OrganizationStateType
const (
	INITIALIZED OrganizationStateType = "initialized"
	VALIDATED OrganizationStateType = "validated"
	CLOSED OrganizationStateType = "closed"
)

// All allowed values of OrganizationStateType enum
var AllowedOrganizationStateTypeEnumValues = []OrganizationStateType{
	"initialized",
	"validated",
	"closed",
}

func (v *OrganizationStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationStateType(value)
	for _, existing := range AllowedOrganizationStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationStateType", value)
}

// NewOrganizationStateTypeFromValue returns a pointer to a valid OrganizationStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationStateTypeFromValue(v string) (*OrganizationStateType, error) {
	ev := OrganizationStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationStateType: valid values are %v", v, AllowedOrganizationStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationStateType) IsValid() bool {
	for _, existing := range AllowedOrganizationStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationStateType value
func (v OrganizationStateType) Ptr() *OrganizationStateType {
	return &v
}

type NullableOrganizationStateType struct {
	value *OrganizationStateType
	isSet bool
}

func (v NullableOrganizationStateType) Get() *OrganizationStateType {
	return v.value
}

func (v *NullableOrganizationStateType) Set(val *OrganizationStateType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationStateType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationStateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationStateType(val *OrganizationStateType) *NullableOrganizationStateType {
	return &NullableOrganizationStateType{value: val, isSet: true}
}

func (v NullableOrganizationStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationStateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

