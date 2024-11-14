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

// checks if the TaxDefinitionMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDefinitionMVO{}

// TaxDefinitionMVO struct for TaxDefinitionMVO
type TaxDefinitionMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Unique identifier of the tax.
	Id *string `json:"id,omitempty"`
	// Tax name.
	Name *string `json:"name,omitempty"`
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// Name of the jurisdiction that levies the tax
	JurisdictionName *string `json:"jurisdictionName,omitempty"`
	// Level of the jurisdiction that levies the tax
	JurisdictionLevel *string `json:"jurisdictionLevel,omitempty"`
	// Type of the tax.
	TaxType *string `json:"taxType,omitempty"`
}

type _TaxDefinitionMVO TaxDefinitionMVO

// NewTaxDefinitionMVO instantiates a new TaxDefinitionMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDefinitionMVO(type_ string) *TaxDefinitionMVO {
	this := TaxDefinitionMVO{}
	this.Type = type_
	return &this
}

// NewTaxDefinitionMVOWithDefaults instantiates a new TaxDefinitionMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDefinitionMVOWithDefaults() *TaxDefinitionMVO {
	this := TaxDefinitionMVO{}
	return &this
}

// GetType returns the Type field value
func (o *TaxDefinitionMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxDefinitionMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *TaxDefinitionMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *TaxDefinitionMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaxDefinitionMVO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaxDefinitionMVO) SetName(v string) {
	o.Name = &v
}

// GetValidFor returns the ValidFor field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetValidFor() TimePeriod {
	if o == nil || IsNil(o.ValidFor) {
		var ret TimePeriod
		return ret
	}
	return *o.ValidFor
}

// GetValidForOk returns a tuple with the ValidFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetValidForOk() (*TimePeriod, bool) {
	if o == nil || IsNil(o.ValidFor) {
		return nil, false
	}
	return o.ValidFor, true
}

// HasValidFor returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasValidFor() bool {
	if o != nil && !IsNil(o.ValidFor) {
		return true
	}

	return false
}

// SetValidFor gets a reference to the given TimePeriod and assigns it to the ValidFor field.
func (o *TaxDefinitionMVO) SetValidFor(v TimePeriod) {
	o.ValidFor = &v
}

// GetJurisdictionName returns the JurisdictionName field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetJurisdictionName() string {
	if o == nil || IsNil(o.JurisdictionName) {
		var ret string
		return ret
	}
	return *o.JurisdictionName
}

// GetJurisdictionNameOk returns a tuple with the JurisdictionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetJurisdictionNameOk() (*string, bool) {
	if o == nil || IsNil(o.JurisdictionName) {
		return nil, false
	}
	return o.JurisdictionName, true
}

// HasJurisdictionName returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasJurisdictionName() bool {
	if o != nil && !IsNil(o.JurisdictionName) {
		return true
	}

	return false
}

// SetJurisdictionName gets a reference to the given string and assigns it to the JurisdictionName field.
func (o *TaxDefinitionMVO) SetJurisdictionName(v string) {
	o.JurisdictionName = &v
}

// GetJurisdictionLevel returns the JurisdictionLevel field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetJurisdictionLevel() string {
	if o == nil || IsNil(o.JurisdictionLevel) {
		var ret string
		return ret
	}
	return *o.JurisdictionLevel
}

// GetJurisdictionLevelOk returns a tuple with the JurisdictionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetJurisdictionLevelOk() (*string, bool) {
	if o == nil || IsNil(o.JurisdictionLevel) {
		return nil, false
	}
	return o.JurisdictionLevel, true
}

// HasJurisdictionLevel returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasJurisdictionLevel() bool {
	if o != nil && !IsNil(o.JurisdictionLevel) {
		return true
	}

	return false
}

// SetJurisdictionLevel gets a reference to the given string and assigns it to the JurisdictionLevel field.
func (o *TaxDefinitionMVO) SetJurisdictionLevel(v string) {
	o.JurisdictionLevel = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *TaxDefinitionMVO) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDefinitionMVO) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *TaxDefinitionMVO) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *TaxDefinitionMVO) SetTaxType(v string) {
	o.TaxType = &v
}

func (o TaxDefinitionMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxDefinitionMVO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ValidFor) {
		toSerialize["validFor"] = o.ValidFor
	}
	if !IsNil(o.JurisdictionName) {
		toSerialize["jurisdictionName"] = o.JurisdictionName
	}
	if !IsNil(o.JurisdictionLevel) {
		toSerialize["jurisdictionLevel"] = o.JurisdictionLevel
	}
	if !IsNil(o.TaxType) {
		toSerialize["taxType"] = o.TaxType
	}
	return toSerialize, nil
}

func (o *TaxDefinitionMVO) UnmarshalJSON(data []byte) (err error) {
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

	varTaxDefinitionMVO := _TaxDefinitionMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxDefinitionMVO)

	if err != nil {
		return err
	}

	*o = TaxDefinitionMVO(varTaxDefinitionMVO)

	return err
}

type NullableTaxDefinitionMVO struct {
	value *TaxDefinitionMVO
	isSet bool
}

func (v NullableTaxDefinitionMVO) Get() *TaxDefinitionMVO {
	return v.value
}

func (v *NullableTaxDefinitionMVO) Set(val *TaxDefinitionMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDefinitionMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDefinitionMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDefinitionMVO(val *TaxDefinitionMVO) *NullableTaxDefinitionMVO {
	return &NullableTaxDefinitionMVO{value: val, isSet: true}
}

func (v NullableTaxDefinitionMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxDefinitionMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

