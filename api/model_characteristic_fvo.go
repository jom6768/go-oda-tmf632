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

// checks if the CharacteristicFVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacteristicFVO{}

// CharacteristicFVO struct for CharacteristicFVO
type CharacteristicFVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Unique identifier of the characteristic
	Id *string `json:"id,omitempty"`
	// Name of the characteristic
	Name string `json:"name"`
	// Data type of the value of the characteristic
	ValueType *string `json:"valueType,omitempty"`
	CharacteristicRelationship []CharacteristicRelationshipFVO `json:"characteristicRelationship,omitempty"`
}

type _CharacteristicFVO CharacteristicFVO

// NewCharacteristicFVO instantiates a new CharacteristicFVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacteristicFVO(type_ string, name string) *CharacteristicFVO {
	this := CharacteristicFVO{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewCharacteristicFVOWithDefaults instantiates a new CharacteristicFVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacteristicFVOWithDefaults() *CharacteristicFVO {
	this := CharacteristicFVO{}
	return &this
}

// GetType returns the Type field value
func (o *CharacteristicFVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharacteristicFVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *CharacteristicFVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *CharacteristicFVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *CharacteristicFVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *CharacteristicFVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *CharacteristicFVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *CharacteristicFVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CharacteristicFVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CharacteristicFVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CharacteristicFVO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CharacteristicFVO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CharacteristicFVO) SetName(v string) {
	o.Name = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *CharacteristicFVO) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *CharacteristicFVO) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *CharacteristicFVO) SetValueType(v string) {
	o.ValueType = &v
}

// GetCharacteristicRelationship returns the CharacteristicRelationship field value if set, zero value otherwise.
func (o *CharacteristicFVO) GetCharacteristicRelationship() []CharacteristicRelationshipFVO {
	if o == nil || IsNil(o.CharacteristicRelationship) {
		var ret []CharacteristicRelationshipFVO
		return ret
	}
	return o.CharacteristicRelationship
}

// GetCharacteristicRelationshipOk returns a tuple with the CharacteristicRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicFVO) GetCharacteristicRelationshipOk() ([]CharacteristicRelationshipFVO, bool) {
	if o == nil || IsNil(o.CharacteristicRelationship) {
		return nil, false
	}
	return o.CharacteristicRelationship, true
}

// HasCharacteristicRelationship returns a boolean if a field has been set.
func (o *CharacteristicFVO) HasCharacteristicRelationship() bool {
	if o != nil && !IsNil(o.CharacteristicRelationship) {
		return true
	}

	return false
}

// SetCharacteristicRelationship gets a reference to the given []CharacteristicRelationshipFVO and assigns it to the CharacteristicRelationship field.
func (o *CharacteristicFVO) SetCharacteristicRelationship(v []CharacteristicRelationshipFVO) {
	o.CharacteristicRelationship = v
}

func (o CharacteristicFVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacteristicFVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.CharacteristicRelationship) {
		toSerialize["characteristicRelationship"] = o.CharacteristicRelationship
	}
	return toSerialize, nil
}

func (o *CharacteristicFVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
		"name",
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

	varCharacteristicFVO := _CharacteristicFVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCharacteristicFVO)

	if err != nil {
		return err
	}

	*o = CharacteristicFVO(varCharacteristicFVO)

	return err
}

type NullableCharacteristicFVO struct {
	value *CharacteristicFVO
	isSet bool
}

func (v NullableCharacteristicFVO) Get() *CharacteristicFVO {
	return v.value
}

func (v *NullableCharacteristicFVO) Set(val *CharacteristicFVO) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacteristicFVO) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacteristicFVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacteristicFVO(val *CharacteristicFVO) *NullableCharacteristicFVO {
	return &NullableCharacteristicFVO{value: val, isSet: true}
}

func (v NullableCharacteristicFVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacteristicFVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


