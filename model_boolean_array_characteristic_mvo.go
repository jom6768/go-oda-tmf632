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

// checks if the BooleanArrayCharacteristicMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanArrayCharacteristicMVO{}

// BooleanArrayCharacteristicMVO struct for BooleanArrayCharacteristicMVO
type BooleanArrayCharacteristicMVO struct {
	CharacteristicMVO
	Value []bool `json:"value"`
}

type _BooleanArrayCharacteristicMVO BooleanArrayCharacteristicMVO

// NewBooleanArrayCharacteristicMVO instantiates a new BooleanArrayCharacteristicMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanArrayCharacteristicMVO(value []bool, type_ string) *BooleanArrayCharacteristicMVO {
	this := BooleanArrayCharacteristicMVO{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewBooleanArrayCharacteristicMVOWithDefaults instantiates a new BooleanArrayCharacteristicMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanArrayCharacteristicMVOWithDefaults() *BooleanArrayCharacteristicMVO {
	this := BooleanArrayCharacteristicMVO{}
	return &this
}

// GetValue returns the Value field value
func (o *BooleanArrayCharacteristicMVO) GetValue() []bool {
	if o == nil {
		var ret []bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BooleanArrayCharacteristicMVO) GetValueOk() ([]bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *BooleanArrayCharacteristicMVO) SetValue(v []bool) {
	o.Value = v
}

func (o BooleanArrayCharacteristicMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanArrayCharacteristicMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCharacteristicMVO, errCharacteristicMVO := json.Marshal(o.CharacteristicMVO)
	if errCharacteristicMVO != nil {
		return map[string]interface{}{}, errCharacteristicMVO
	}
	errCharacteristicMVO = json.Unmarshal([]byte(serializedCharacteristicMVO), &toSerialize)
	if errCharacteristicMVO != nil {
		return map[string]interface{}{}, errCharacteristicMVO
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *BooleanArrayCharacteristicMVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varBooleanArrayCharacteristicMVO := _BooleanArrayCharacteristicMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBooleanArrayCharacteristicMVO)

	if err != nil {
		return err
	}

	*o = BooleanArrayCharacteristicMVO(varBooleanArrayCharacteristicMVO)

	return err
}

type NullableBooleanArrayCharacteristicMVO struct {
	value *BooleanArrayCharacteristicMVO
	isSet bool
}

func (v NullableBooleanArrayCharacteristicMVO) Get() *BooleanArrayCharacteristicMVO {
	return v.value
}

func (v *NullableBooleanArrayCharacteristicMVO) Set(val *BooleanArrayCharacteristicMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanArrayCharacteristicMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanArrayCharacteristicMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanArrayCharacteristicMVO(val *BooleanArrayCharacteristicMVO) *NullableBooleanArrayCharacteristicMVO {
	return &NullableBooleanArrayCharacteristicMVO{value: val, isSet: true}
}

func (v NullableBooleanArrayCharacteristicMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanArrayCharacteristicMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


