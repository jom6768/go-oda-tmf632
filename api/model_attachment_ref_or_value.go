/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// AttachmentRefOrValue - The polymorphic attributes @type, @schemaLocation & @referredType are related to the Attachment entity and not the AttachmentRefOrValue class itself
type AttachmentRefOrValue struct {
	Attachment *Attachment
	AttachmentRef *AttachmentRef
}

// AttachmentAsAttachmentRefOrValue is a convenience function that returns Attachment wrapped in AttachmentRefOrValue
func AttachmentAsAttachmentRefOrValue(v *Attachment) AttachmentRefOrValue {
	return AttachmentRefOrValue{
		Attachment: v,
	}
}

// AttachmentRefAsAttachmentRefOrValue is a convenience function that returns AttachmentRef wrapped in AttachmentRefOrValue
func AttachmentRefAsAttachmentRefOrValue(v *AttachmentRef) AttachmentRefOrValue {
	return AttachmentRefOrValue{
		AttachmentRef: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AttachmentRefOrValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Attachment
	err = newStrictDecoder(data).Decode(&dst.Attachment)
	if err == nil {
		jsonAttachment, _ := json.Marshal(dst.Attachment)
		if string(jsonAttachment) == "{}" { // empty struct
			dst.Attachment = nil
		} else {
			if err = validator.Validate(dst.Attachment); err != nil {
				dst.Attachment = nil
			} else {
				match++
			}
		}
	} else {
		dst.Attachment = nil
	}

	// try to unmarshal data into AttachmentRef
	err = newStrictDecoder(data).Decode(&dst.AttachmentRef)
	if err == nil {
		jsonAttachmentRef, _ := json.Marshal(dst.AttachmentRef)
		if string(jsonAttachmentRef) == "{}" { // empty struct
			dst.AttachmentRef = nil
		} else {
			if err = validator.Validate(dst.AttachmentRef); err != nil {
				dst.AttachmentRef = nil
			} else {
				match++
			}
		}
	} else {
		dst.AttachmentRef = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Attachment = nil
		dst.AttachmentRef = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AttachmentRefOrValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AttachmentRefOrValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AttachmentRefOrValue) MarshalJSON() ([]byte, error) {
	if src.Attachment != nil {
		return json.Marshal(&src.Attachment)
	}

	if src.AttachmentRef != nil {
		return json.Marshal(&src.AttachmentRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AttachmentRefOrValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Attachment != nil {
		return obj.Attachment
	}

	if obj.AttachmentRef != nil {
		return obj.AttachmentRef
	}

	// all schemas are nil
	return nil
}

type NullableAttachmentRefOrValue struct {
	value *AttachmentRefOrValue
	isSet bool
}

func (v NullableAttachmentRefOrValue) Get() *AttachmentRefOrValue {
	return v.value
}

func (v *NullableAttachmentRefOrValue) Set(val *AttachmentRefOrValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentRefOrValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentRefOrValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentRefOrValue(val *AttachmentRefOrValue) *NullableAttachmentRefOrValue {
	return &NullableAttachmentRefOrValue{value: val, isSet: true}
}

func (v NullableAttachmentRefOrValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentRefOrValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


