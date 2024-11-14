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

// checks if the RelatedPartyOrPartyRoleMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedPartyOrPartyRoleMVO{}

// RelatedPartyOrPartyRoleMVO struct for RelatedPartyOrPartyRoleMVO
type RelatedPartyOrPartyRoleMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Role played by the related party or party role in the context of the specific entity it is linked to. Such as 'initiator', 'customer',  'salesAgent', 'user'
	Role string `json:"role"`
	PartyOrPartyRole *PartyOrPartyRoleMVO `json:"partyOrPartyRole,omitempty"`
}

type _RelatedPartyOrPartyRoleMVO RelatedPartyOrPartyRoleMVO

// NewRelatedPartyOrPartyRoleMVO instantiates a new RelatedPartyOrPartyRoleMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedPartyOrPartyRoleMVO(type_ string, role string) *RelatedPartyOrPartyRoleMVO {
	this := RelatedPartyOrPartyRoleMVO{}
	this.Type = type_
	this.Role = role
	return &this
}

// NewRelatedPartyOrPartyRoleMVOWithDefaults instantiates a new RelatedPartyOrPartyRoleMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedPartyOrPartyRoleMVOWithDefaults() *RelatedPartyOrPartyRoleMVO {
	this := RelatedPartyOrPartyRoleMVO{}
	return &this
}

// GetType returns the Type field value
func (o *RelatedPartyOrPartyRoleMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelatedPartyOrPartyRoleMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RelatedPartyOrPartyRoleMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *RelatedPartyOrPartyRoleMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPartyOrPartyRoleMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *RelatedPartyOrPartyRoleMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *RelatedPartyOrPartyRoleMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *RelatedPartyOrPartyRoleMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPartyOrPartyRoleMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *RelatedPartyOrPartyRoleMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *RelatedPartyOrPartyRoleMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetRole returns the Role field value
func (o *RelatedPartyOrPartyRoleMVO) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RelatedPartyOrPartyRoleMVO) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RelatedPartyOrPartyRoleMVO) SetRole(v string) {
	o.Role = v
}

// GetPartyOrPartyRole returns the PartyOrPartyRole field value if set, zero value otherwise.
func (o *RelatedPartyOrPartyRoleMVO) GetPartyOrPartyRole() PartyOrPartyRoleMVO {
	if o == nil || IsNil(o.PartyOrPartyRole) {
		var ret PartyOrPartyRoleMVO
		return ret
	}
	return *o.PartyOrPartyRole
}

// GetPartyOrPartyRoleOk returns a tuple with the PartyOrPartyRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedPartyOrPartyRoleMVO) GetPartyOrPartyRoleOk() (*PartyOrPartyRoleMVO, bool) {
	if o == nil || IsNil(o.PartyOrPartyRole) {
		return nil, false
	}
	return o.PartyOrPartyRole, true
}

// HasPartyOrPartyRole returns a boolean if a field has been set.
func (o *RelatedPartyOrPartyRoleMVO) HasPartyOrPartyRole() bool {
	if o != nil && !IsNil(o.PartyOrPartyRole) {
		return true
	}

	return false
}

// SetPartyOrPartyRole gets a reference to the given PartyOrPartyRoleMVO and assigns it to the PartyOrPartyRole field.
func (o *RelatedPartyOrPartyRoleMVO) SetPartyOrPartyRole(v PartyOrPartyRoleMVO) {
	o.PartyOrPartyRole = &v
}

func (o RelatedPartyOrPartyRoleMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedPartyOrPartyRoleMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	toSerialize["role"] = o.Role
	if !IsNil(o.PartyOrPartyRole) {
		toSerialize["partyOrPartyRole"] = o.PartyOrPartyRole
	}
	return toSerialize, nil
}

func (o *RelatedPartyOrPartyRoleMVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
		"role",
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

	varRelatedPartyOrPartyRoleMVO := _RelatedPartyOrPartyRoleMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelatedPartyOrPartyRoleMVO)

	if err != nil {
		return err
	}

	*o = RelatedPartyOrPartyRoleMVO(varRelatedPartyOrPartyRoleMVO)

	return err
}

type NullableRelatedPartyOrPartyRoleMVO struct {
	value *RelatedPartyOrPartyRoleMVO
	isSet bool
}

func (v NullableRelatedPartyOrPartyRoleMVO) Get() *RelatedPartyOrPartyRoleMVO {
	return v.value
}

func (v *NullableRelatedPartyOrPartyRoleMVO) Set(val *RelatedPartyOrPartyRoleMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedPartyOrPartyRoleMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedPartyOrPartyRoleMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedPartyOrPartyRoleMVO(val *RelatedPartyOrPartyRoleMVO) *NullableRelatedPartyOrPartyRoleMVO {
	return &NullableRelatedPartyOrPartyRoleMVO{value: val, isSet: true}
}

func (v NullableRelatedPartyOrPartyRoleMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedPartyOrPartyRoleMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


