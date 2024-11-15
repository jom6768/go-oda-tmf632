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

// checks if the ExternalIdentifierMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalIdentifierMVO{}

// ExternalIdentifierMVO struct for ExternalIdentifierMVO
type ExternalIdentifierMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Name of the external system that owns the entity.
	Owner *string `json:"owner,omitempty"`
	// Type of the identification, typically would be the type of the entity within the external system
	ExternalIdentifierType *string `json:"externalIdentifierType,omitempty"`
	// identification of the entity within the external system.
	Id *string `json:"id,omitempty"`
}

type _ExternalIdentifierMVO ExternalIdentifierMVO

// NewExternalIdentifierMVO instantiates a new ExternalIdentifierMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalIdentifierMVO(type_ string) *ExternalIdentifierMVO {
	this := ExternalIdentifierMVO{}
	this.Type = type_
	return &this
}

// NewExternalIdentifierMVOWithDefaults instantiates a new ExternalIdentifierMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalIdentifierMVOWithDefaults() *ExternalIdentifierMVO {
	this := ExternalIdentifierMVO{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalIdentifierMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalIdentifierMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *ExternalIdentifierMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *ExternalIdentifierMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *ExternalIdentifierMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *ExternalIdentifierMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *ExternalIdentifierMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *ExternalIdentifierMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ExternalIdentifierMVO) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ExternalIdentifierMVO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ExternalIdentifierMVO) SetOwner(v string) {
	o.Owner = &v
}

// GetExternalIdentifierType returns the ExternalIdentifierType field value if set, zero value otherwise.
func (o *ExternalIdentifierMVO) GetExternalIdentifierType() string {
	if o == nil || IsNil(o.ExternalIdentifierType) {
		var ret string
		return ret
	}
	return *o.ExternalIdentifierType
}

// GetExternalIdentifierTypeOk returns a tuple with the ExternalIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetExternalIdentifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalIdentifierType) {
		return nil, false
	}
	return o.ExternalIdentifierType, true
}

// HasExternalIdentifierType returns a boolean if a field has been set.
func (o *ExternalIdentifierMVO) HasExternalIdentifierType() bool {
	if o != nil && !IsNil(o.ExternalIdentifierType) {
		return true
	}

	return false
}

// SetExternalIdentifierType gets a reference to the given string and assigns it to the ExternalIdentifierType field.
func (o *ExternalIdentifierMVO) SetExternalIdentifierType(v string) {
	o.ExternalIdentifierType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalIdentifierMVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdentifierMVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalIdentifierMVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalIdentifierMVO) SetId(v string) {
	o.Id = &v
}

func (o ExternalIdentifierMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalIdentifierMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.ExternalIdentifierType) {
		toSerialize["externalIdentifierType"] = o.ExternalIdentifierType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

func (o *ExternalIdentifierMVO) UnmarshalJSON(data []byte) (err error) {
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

	varExternalIdentifierMVO := _ExternalIdentifierMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalIdentifierMVO)

	if err != nil {
		return err
	}

	*o = ExternalIdentifierMVO(varExternalIdentifierMVO)

	return err
}

type NullableExternalIdentifierMVO struct {
	value *ExternalIdentifierMVO
	isSet bool
}

func (v NullableExternalIdentifierMVO) Get() *ExternalIdentifierMVO {
	return v.value
}

func (v *NullableExternalIdentifierMVO) Set(val *ExternalIdentifierMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalIdentifierMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalIdentifierMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalIdentifierMVO(val *ExternalIdentifierMVO) *NullableExternalIdentifierMVO {
	return &NullableExternalIdentifierMVO{value: val, isSet: true}
}

func (v NullableExternalIdentifierMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalIdentifierMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


