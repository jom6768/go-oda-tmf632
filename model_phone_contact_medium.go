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

// checks if the PhoneContactMedium type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneContactMedium{}

// PhoneContactMedium struct for PhoneContactMedium
type PhoneContactMedium struct {
	ContactMedium
	// The phone number of the contact
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type _PhoneContactMedium PhoneContactMedium

// NewPhoneContactMedium instantiates a new PhoneContactMedium object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneContactMedium(type_ string) *PhoneContactMedium {
	this := PhoneContactMedium{}
	this.Type = type_
	return &this
}

// NewPhoneContactMediumWithDefaults instantiates a new PhoneContactMedium object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneContactMediumWithDefaults() *PhoneContactMedium {
	this := PhoneContactMedium{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PhoneContactMedium) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneContactMedium) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PhoneContactMedium) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PhoneContactMedium) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o PhoneContactMedium) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneContactMedium) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContactMedium, errContactMedium := json.Marshal(o.ContactMedium)
	if errContactMedium != nil {
		return map[string]interface{}{}, errContactMedium
	}
	errContactMedium = json.Unmarshal([]byte(serializedContactMedium), &toSerialize)
	if errContactMedium != nil {
		return map[string]interface{}{}, errContactMedium
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

func (o *PhoneContactMedium) UnmarshalJSON(data []byte) (err error) {
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

	varPhoneContactMedium := _PhoneContactMedium{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneContactMedium)

	if err != nil {
		return err
	}

	*o = PhoneContactMedium(varPhoneContactMedium)

	return err
}

type NullablePhoneContactMedium struct {
	value *PhoneContactMedium
	isSet bool
}

func (v NullablePhoneContactMedium) Get() *PhoneContactMedium {
	return v.value
}

func (v *NullablePhoneContactMedium) Set(val *PhoneContactMedium) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneContactMedium) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneContactMedium) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneContactMedium(val *PhoneContactMedium) *NullablePhoneContactMedium {
	return &NullablePhoneContactMedium{value: val, isSet: true}
}

func (v NullablePhoneContactMedium) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneContactMedium) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


