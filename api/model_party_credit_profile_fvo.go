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

// checks if the PartyCreditProfileFVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartyCreditProfileFVO{}

// PartyCreditProfileFVO struct for PartyCreditProfileFVO
type PartyCreditProfileFVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// unique identifier
	Id *string `json:"id,omitempty"`
	// Name of the credit agency giving the score
	CreditAgencyName *string `json:"creditAgencyName,omitempty"`
	// Type of the credit agency giving the score
	CreditAgencyType *string `json:"creditAgencyType,omitempty"`
	// Reference corresponding to the credit rating
	RatingReference *string `json:"ratingReference,omitempty"`
	// A measure of a party's creditworthiness calculated on the basis of a combination of factors such as their income and credit history
	RatingScore *int32 `json:"ratingScore,omitempty"`
	ValidFor *TimePeriod `json:"validFor,omitempty"`
}

type _PartyCreditProfileFVO PartyCreditProfileFVO

// NewPartyCreditProfileFVO instantiates a new PartyCreditProfileFVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyCreditProfileFVO(type_ string) *PartyCreditProfileFVO {
	this := PartyCreditProfileFVO{}
	this.Type = type_
	return &this
}

// NewPartyCreditProfileFVOWithDefaults instantiates a new PartyCreditProfileFVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyCreditProfileFVOWithDefaults() *PartyCreditProfileFVO {
	this := PartyCreditProfileFVO{}
	return &this
}

// GetType returns the Type field value
func (o *PartyCreditProfileFVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PartyCreditProfileFVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *PartyCreditProfileFVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *PartyCreditProfileFVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PartyCreditProfileFVO) SetId(v string) {
	o.Id = &v
}

// GetCreditAgencyName returns the CreditAgencyName field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetCreditAgencyName() string {
	if o == nil || IsNil(o.CreditAgencyName) {
		var ret string
		return ret
	}
	return *o.CreditAgencyName
}

// GetCreditAgencyNameOk returns a tuple with the CreditAgencyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetCreditAgencyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreditAgencyName) {
		return nil, false
	}
	return o.CreditAgencyName, true
}

// HasCreditAgencyName returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasCreditAgencyName() bool {
	if o != nil && !IsNil(o.CreditAgencyName) {
		return true
	}

	return false
}

// SetCreditAgencyName gets a reference to the given string and assigns it to the CreditAgencyName field.
func (o *PartyCreditProfileFVO) SetCreditAgencyName(v string) {
	o.CreditAgencyName = &v
}

// GetCreditAgencyType returns the CreditAgencyType field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetCreditAgencyType() string {
	if o == nil || IsNil(o.CreditAgencyType) {
		var ret string
		return ret
	}
	return *o.CreditAgencyType
}

// GetCreditAgencyTypeOk returns a tuple with the CreditAgencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetCreditAgencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreditAgencyType) {
		return nil, false
	}
	return o.CreditAgencyType, true
}

// HasCreditAgencyType returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasCreditAgencyType() bool {
	if o != nil && !IsNil(o.CreditAgencyType) {
		return true
	}

	return false
}

// SetCreditAgencyType gets a reference to the given string and assigns it to the CreditAgencyType field.
func (o *PartyCreditProfileFVO) SetCreditAgencyType(v string) {
	o.CreditAgencyType = &v
}

// GetRatingReference returns the RatingReference field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetRatingReference() string {
	if o == nil || IsNil(o.RatingReference) {
		var ret string
		return ret
	}
	return *o.RatingReference
}

// GetRatingReferenceOk returns a tuple with the RatingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetRatingReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.RatingReference) {
		return nil, false
	}
	return o.RatingReference, true
}

// HasRatingReference returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasRatingReference() bool {
	if o != nil && !IsNil(o.RatingReference) {
		return true
	}

	return false
}

// SetRatingReference gets a reference to the given string and assigns it to the RatingReference field.
func (o *PartyCreditProfileFVO) SetRatingReference(v string) {
	o.RatingReference = &v
}

// GetRatingScore returns the RatingScore field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetRatingScore() int32 {
	if o == nil || IsNil(o.RatingScore) {
		var ret int32
		return ret
	}
	return *o.RatingScore
}

// GetRatingScoreOk returns a tuple with the RatingScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetRatingScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingScore) {
		return nil, false
	}
	return o.RatingScore, true
}

// HasRatingScore returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasRatingScore() bool {
	if o != nil && !IsNil(o.RatingScore) {
		return true
	}

	return false
}

// SetRatingScore gets a reference to the given int32 and assigns it to the RatingScore field.
func (o *PartyCreditProfileFVO) SetRatingScore(v int32) {
	o.RatingScore = &v
}

// GetValidFor returns the ValidFor field value if set, zero value otherwise.
func (o *PartyCreditProfileFVO) GetValidFor() TimePeriod {
	if o == nil || IsNil(o.ValidFor) {
		var ret TimePeriod
		return ret
	}
	return *o.ValidFor
}

// GetValidForOk returns a tuple with the ValidFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyCreditProfileFVO) GetValidForOk() (*TimePeriod, bool) {
	if o == nil || IsNil(o.ValidFor) {
		return nil, false
	}
	return o.ValidFor, true
}

// HasValidFor returns a boolean if a field has been set.
func (o *PartyCreditProfileFVO) HasValidFor() bool {
	if o != nil && !IsNil(o.ValidFor) {
		return true
	}

	return false
}

// SetValidFor gets a reference to the given TimePeriod and assigns it to the ValidFor field.
func (o *PartyCreditProfileFVO) SetValidFor(v TimePeriod) {
	o.ValidFor = &v
}

func (o PartyCreditProfileFVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartyCreditProfileFVO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreditAgencyName) {
		toSerialize["creditAgencyName"] = o.CreditAgencyName
	}
	if !IsNil(o.CreditAgencyType) {
		toSerialize["creditAgencyType"] = o.CreditAgencyType
	}
	if !IsNil(o.RatingReference) {
		toSerialize["ratingReference"] = o.RatingReference
	}
	if !IsNil(o.RatingScore) {
		toSerialize["ratingScore"] = o.RatingScore
	}
	if !IsNil(o.ValidFor) {
		toSerialize["validFor"] = o.ValidFor
	}
	return toSerialize, nil
}

func (o *PartyCreditProfileFVO) UnmarshalJSON(data []byte) (err error) {
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

	varPartyCreditProfileFVO := _PartyCreditProfileFVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPartyCreditProfileFVO)

	if err != nil {
		return err
	}

	*o = PartyCreditProfileFVO(varPartyCreditProfileFVO)

	return err
}

type NullablePartyCreditProfileFVO struct {
	value *PartyCreditProfileFVO
	isSet bool
}

func (v NullablePartyCreditProfileFVO) Get() *PartyCreditProfileFVO {
	return v.value
}

func (v *NullablePartyCreditProfileFVO) Set(val *PartyCreditProfileFVO) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyCreditProfileFVO) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyCreditProfileFVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyCreditProfileFVO(val *PartyCreditProfileFVO) *NullablePartyCreditProfileFVO {
	return &NullablePartyCreditProfileFVO{value: val, isSet: true}
}

func (v NullablePartyCreditProfileFVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyCreditProfileFVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


