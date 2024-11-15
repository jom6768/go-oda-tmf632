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

// checks if the StringCharacteristic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringCharacteristic{}

// StringCharacteristic struct for StringCharacteristic
type StringCharacteristic struct {
	Characteristic
	// Value of the characteristic
	Value *string `json:"value,omitempty"`
}

type _StringCharacteristic StringCharacteristic

// NewStringCharacteristic instantiates a new StringCharacteristic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringCharacteristic(type_ string) *StringCharacteristic {
	this := StringCharacteristic{}
	this.Type = type_
	return &this
}

// NewStringCharacteristicWithDefaults instantiates a new StringCharacteristic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringCharacteristicWithDefaults() *StringCharacteristic {
	this := StringCharacteristic{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringCharacteristic) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringCharacteristic) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringCharacteristic) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StringCharacteristic) SetValue(v string) {
	o.Value = &v
}

func (o StringCharacteristic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringCharacteristic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCharacteristic, errCharacteristic := json.Marshal(o.Characteristic)
	if errCharacteristic != nil {
		return map[string]interface{}{}, errCharacteristic
	}
	errCharacteristic = json.Unmarshal([]byte(serializedCharacteristic), &toSerialize)
	if errCharacteristic != nil {
		return map[string]interface{}{}, errCharacteristic
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *StringCharacteristic) UnmarshalJSON(data []byte) (err error) {
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

	varStringCharacteristic := _StringCharacteristic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringCharacteristic)

	if err != nil {
		return err
	}

	*o = StringCharacteristic(varStringCharacteristic)

	return err
}

type NullableStringCharacteristic struct {
	value *StringCharacteristic
	isSet bool
}

func (v NullableStringCharacteristic) Get() *StringCharacteristic {
	return v.value
}

func (v *NullableStringCharacteristic) Set(val *StringCharacteristic) {
	v.value = val
	v.isSet = true
}

func (v NullableStringCharacteristic) IsSet() bool {
	return v.isSet
}

func (v *NullableStringCharacteristic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringCharacteristic(val *StringCharacteristic) *NullableStringCharacteristic {
	return &NullableStringCharacteristic{value: val, isSet: true}
}

func (v NullableStringCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringCharacteristic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


