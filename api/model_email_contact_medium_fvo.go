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

// checks if the EmailContactMediumFVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailContactMediumFVO{}

// EmailContactMediumFVO struct for EmailContactMediumFVO
type EmailContactMediumFVO struct {
	ContactMediumFVO
	// Full email address in standard format
	EmailAddress *string `json:"emailAddress,omitempty"`
}

type _EmailContactMediumFVO EmailContactMediumFVO

// NewEmailContactMediumFVO instantiates a new EmailContactMediumFVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailContactMediumFVO(type_ string) *EmailContactMediumFVO {
	this := EmailContactMediumFVO{}
	this.Type = type_
	return &this
}

// NewEmailContactMediumFVOWithDefaults instantiates a new EmailContactMediumFVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailContactMediumFVOWithDefaults() *EmailContactMediumFVO {
	this := EmailContactMediumFVO{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *EmailContactMediumFVO) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailContactMediumFVO) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *EmailContactMediumFVO) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *EmailContactMediumFVO) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o EmailContactMediumFVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailContactMediumFVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContactMediumFVO, errContactMediumFVO := json.Marshal(o.ContactMediumFVO)
	if errContactMediumFVO != nil {
		return map[string]interface{}{}, errContactMediumFVO
	}
	errContactMediumFVO = json.Unmarshal([]byte(serializedContactMediumFVO), &toSerialize)
	if errContactMediumFVO != nil {
		return map[string]interface{}{}, errContactMediumFVO
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	return toSerialize, nil
}

func (o *EmailContactMediumFVO) UnmarshalJSON(data []byte) (err error) {
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

	varEmailContactMediumFVO := _EmailContactMediumFVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailContactMediumFVO)

	if err != nil {
		return err
	}

	*o = EmailContactMediumFVO(varEmailContactMediumFVO)

	return err
}

type NullableEmailContactMediumFVO struct {
	value *EmailContactMediumFVO
	isSet bool
}

func (v NullableEmailContactMediumFVO) Get() *EmailContactMediumFVO {
	return v.value
}

func (v *NullableEmailContactMediumFVO) Set(val *EmailContactMediumFVO) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailContactMediumFVO) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailContactMediumFVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailContactMediumFVO(val *EmailContactMediumFVO) *NullableEmailContactMediumFVO {
	return &NullableEmailContactMediumFVO{value: val, isSet: true}
}

func (v NullableEmailContactMediumFVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailContactMediumFVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


