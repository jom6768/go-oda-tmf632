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

// checks if the OrganizationChildRelationshipMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationChildRelationshipMVO{}

// OrganizationChildRelationshipMVO struct for OrganizationChildRelationshipMVO
type OrganizationChildRelationshipMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Type of the relationship. Could be juridical, hierarchical, geographical, functional for example.
	RelationshipType *string `json:"relationshipType,omitempty"`
	Organization *OrganizationRefMVO `json:"organization,omitempty"`
}

type _OrganizationChildRelationshipMVO OrganizationChildRelationshipMVO

// NewOrganizationChildRelationshipMVO instantiates a new OrganizationChildRelationshipMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationChildRelationshipMVO(type_ string) *OrganizationChildRelationshipMVO {
	this := OrganizationChildRelationshipMVO{}
	this.Type = type_
	return &this
}

// NewOrganizationChildRelationshipMVOWithDefaults instantiates a new OrganizationChildRelationshipMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationChildRelationshipMVOWithDefaults() *OrganizationChildRelationshipMVO {
	this := OrganizationChildRelationshipMVO{}
	return &this
}

// GetType returns the Type field value
func (o *OrganizationChildRelationshipMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrganizationChildRelationshipMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrganizationChildRelationshipMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *OrganizationChildRelationshipMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationChildRelationshipMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *OrganizationChildRelationshipMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *OrganizationChildRelationshipMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *OrganizationChildRelationshipMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationChildRelationshipMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *OrganizationChildRelationshipMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *OrganizationChildRelationshipMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetRelationshipType returns the RelationshipType field value if set, zero value otherwise.
func (o *OrganizationChildRelationshipMVO) GetRelationshipType() string {
	if o == nil || IsNil(o.RelationshipType) {
		var ret string
		return ret
	}
	return *o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationChildRelationshipMVO) GetRelationshipTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RelationshipType) {
		return nil, false
	}
	return o.RelationshipType, true
}

// HasRelationshipType returns a boolean if a field has been set.
func (o *OrganizationChildRelationshipMVO) HasRelationshipType() bool {
	if o != nil && !IsNil(o.RelationshipType) {
		return true
	}

	return false
}

// SetRelationshipType gets a reference to the given string and assigns it to the RelationshipType field.
func (o *OrganizationChildRelationshipMVO) SetRelationshipType(v string) {
	o.RelationshipType = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationChildRelationshipMVO) GetOrganization() OrganizationRefMVO {
	if o == nil || IsNil(o.Organization) {
		var ret OrganizationRefMVO
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationChildRelationshipMVO) GetOrganizationOk() (*OrganizationRefMVO, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationChildRelationshipMVO) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationRefMVO and assigns it to the Organization field.
func (o *OrganizationChildRelationshipMVO) SetOrganization(v OrganizationRefMVO) {
	o.Organization = &v
}

func (o OrganizationChildRelationshipMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationChildRelationshipMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	if !IsNil(o.RelationshipType) {
		toSerialize["relationshipType"] = o.RelationshipType
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

func (o *OrganizationChildRelationshipMVO) UnmarshalJSON(data []byte) (err error) {
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

	varOrganizationChildRelationshipMVO := _OrganizationChildRelationshipMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationChildRelationshipMVO)

	if err != nil {
		return err
	}

	*o = OrganizationChildRelationshipMVO(varOrganizationChildRelationshipMVO)

	return err
}

type NullableOrganizationChildRelationshipMVO struct {
	value *OrganizationChildRelationshipMVO
	isSet bool
}

func (v NullableOrganizationChildRelationshipMVO) Get() *OrganizationChildRelationshipMVO {
	return v.value
}

func (v *NullableOrganizationChildRelationshipMVO) Set(val *OrganizationChildRelationshipMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationChildRelationshipMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationChildRelationshipMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationChildRelationshipMVO(val *OrganizationChildRelationshipMVO) *NullableOrganizationChildRelationshipMVO {
	return &NullableOrganizationChildRelationshipMVO{value: val, isSet: true}
}

func (v NullableOrganizationChildRelationshipMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationChildRelationshipMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


