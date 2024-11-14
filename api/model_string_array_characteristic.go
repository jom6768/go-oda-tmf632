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

// checks if the StringArrayCharacteristic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringArrayCharacteristic{}

// StringArrayCharacteristic struct for StringArrayCharacteristic
type StringArrayCharacteristic struct {
	Characteristic
	Value []string `json:"value,omitempty"`
}

type _StringArrayCharacteristic StringArrayCharacteristic

// NewStringArrayCharacteristic instantiates a new StringArrayCharacteristic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringArrayCharacteristic(type_ string) *StringArrayCharacteristic {
	this := StringArrayCharacteristic{}
	this.Type = type_
	return &this
}

// NewStringArrayCharacteristicWithDefaults instantiates a new StringArrayCharacteristic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringArrayCharacteristicWithDefaults() *StringArrayCharacteristic {
	this := StringArrayCharacteristic{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StringArrayCharacteristic) GetValue() []string {
	if o == nil || IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringArrayCharacteristic) GetValueOk() ([]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StringArrayCharacteristic) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *StringArrayCharacteristic) SetValue(v []string) {
	o.Value = v
}

func (o StringArrayCharacteristic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringArrayCharacteristic) ToMap() (map[string]interface{}, error) {
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

func (o *StringArrayCharacteristic) UnmarshalJSON(data []byte) (err error) {
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

	varStringArrayCharacteristic := _StringArrayCharacteristic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStringArrayCharacteristic)

	if err != nil {
		return err
	}

	*o = StringArrayCharacteristic(varStringArrayCharacteristic)

	return err
}

type NullableStringArrayCharacteristic struct {
	value *StringArrayCharacteristic
	isSet bool
}

func (v NullableStringArrayCharacteristic) Get() *StringArrayCharacteristic {
	return v.value
}

func (v *NullableStringArrayCharacteristic) Set(val *StringArrayCharacteristic) {
	v.value = val
	v.isSet = true
}

func (v NullableStringArrayCharacteristic) IsSet() bool {
	return v.isSet
}

func (v *NullableStringArrayCharacteristic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringArrayCharacteristic(val *StringArrayCharacteristic) *NullableStringArrayCharacteristic {
	return &NullableStringArrayCharacteristic{value: val, isSet: true}
}

func (v NullableStringArrayCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringArrayCharacteristic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

