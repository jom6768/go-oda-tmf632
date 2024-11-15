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

// checks if the CharacteristicRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacteristicRelationship{}

// CharacteristicRelationship struct for CharacteristicRelationship
type CharacteristicRelationship struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Unique identifier of the characteristic
	Id *string `json:"id,omitempty"`
	// The type of relationship
	RelationshipType *string `json:"relationshipType,omitempty"`
}

type _CharacteristicRelationship CharacteristicRelationship

// NewCharacteristicRelationship instantiates a new CharacteristicRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacteristicRelationship(type_ string) *CharacteristicRelationship {
	this := CharacteristicRelationship{}
	this.Type = type_
	return &this
}

// NewCharacteristicRelationshipWithDefaults instantiates a new CharacteristicRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacteristicRelationshipWithDefaults() *CharacteristicRelationship {
	this := CharacteristicRelationship{}
	return &this
}

// GetType returns the Type field value
func (o *CharacteristicRelationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CharacteristicRelationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CharacteristicRelationship) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *CharacteristicRelationship) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicRelationship) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *CharacteristicRelationship) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *CharacteristicRelationship) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *CharacteristicRelationship) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicRelationship) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *CharacteristicRelationship) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *CharacteristicRelationship) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CharacteristicRelationship) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicRelationship) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CharacteristicRelationship) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CharacteristicRelationship) SetId(v string) {
	o.Id = &v
}

// GetRelationshipType returns the RelationshipType field value if set, zero value otherwise.
func (o *CharacteristicRelationship) GetRelationshipType() string {
	if o == nil || IsNil(o.RelationshipType) {
		var ret string
		return ret
	}
	return *o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CharacteristicRelationship) GetRelationshipTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RelationshipType) {
		return nil, false
	}
	return o.RelationshipType, true
}

// HasRelationshipType returns a boolean if a field has been set.
func (o *CharacteristicRelationship) HasRelationshipType() bool {
	if o != nil && !IsNil(o.RelationshipType) {
		return true
	}

	return false
}

// SetRelationshipType gets a reference to the given string and assigns it to the RelationshipType field.
func (o *CharacteristicRelationship) SetRelationshipType(v string) {
	o.RelationshipType = &v
}

func (o CharacteristicRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacteristicRelationship) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RelationshipType) {
		toSerialize["relationshipType"] = o.RelationshipType
	}
	return toSerialize, nil
}

func (o *CharacteristicRelationship) UnmarshalJSON(data []byte) (err error) {
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

	varCharacteristicRelationship := _CharacteristicRelationship{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCharacteristicRelationship)

	if err != nil {
		return err
	}

	*o = CharacteristicRelationship(varCharacteristicRelationship)

	return err
}

type NullableCharacteristicRelationship struct {
	value *CharacteristicRelationship
	isSet bool
}

func (v NullableCharacteristicRelationship) Get() *CharacteristicRelationship {
	return v.value
}

func (v *NullableCharacteristicRelationship) Set(val *CharacteristicRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacteristicRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacteristicRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacteristicRelationship(val *CharacteristicRelationship) *NullableCharacteristicRelationship {
	return &NullableCharacteristicRelationship{value: val, isSet: true}
}

func (v NullableCharacteristicRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacteristicRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


