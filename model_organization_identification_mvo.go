/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the OrganizationIdentificationMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationIdentificationMVO{}

// OrganizationIdentificationMVO struct for OrganizationIdentificationMVO
type OrganizationIdentificationMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Identifier
	IdentificationId *string `json:"identificationId,omitempty"`
	// Authority which has issued the identifier (chamber of commerce...)
	IssuingAuthority *string `json:"issuingAuthority,omitempty"`
	// Date at which the identifier was issued
	IssuingDate *time.Time `json:"issuingDate,omitempty"`
	// Type of identification information used to identify the company in a country or internationally
	IdentificationType *string `json:"identificationType,omitempty"`
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	Attachment *AttachmentRefOrValueMVO `json:"attachment,omitempty"`
}

type _OrganizationIdentificationMVO OrganizationIdentificationMVO

// NewOrganizationIdentificationMVO instantiates a new OrganizationIdentificationMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationIdentificationMVO(type_ string) *OrganizationIdentificationMVO {
	this := OrganizationIdentificationMVO{}
	this.Type = type_
	return &this
}

// NewOrganizationIdentificationMVOWithDefaults instantiates a new OrganizationIdentificationMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationIdentificationMVOWithDefaults() *OrganizationIdentificationMVO {
	this := OrganizationIdentificationMVO{}
	return &this
}

// GetType returns the Type field value
func (o *OrganizationIdentificationMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrganizationIdentificationMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *OrganizationIdentificationMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *OrganizationIdentificationMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetIdentificationId returns the IdentificationId field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetIdentificationId() string {
	if o == nil || IsNil(o.IdentificationId) {
		var ret string
		return ret
	}
	return *o.IdentificationId
}

// GetIdentificationIdOk returns a tuple with the IdentificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetIdentificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentificationId) {
		return nil, false
	}
	return o.IdentificationId, true
}

// HasIdentificationId returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasIdentificationId() bool {
	if o != nil && !IsNil(o.IdentificationId) {
		return true
	}

	return false
}

// SetIdentificationId gets a reference to the given string and assigns it to the IdentificationId field.
func (o *OrganizationIdentificationMVO) SetIdentificationId(v string) {
	o.IdentificationId = &v
}

// GetIssuingAuthority returns the IssuingAuthority field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetIssuingAuthority() string {
	if o == nil || IsNil(o.IssuingAuthority) {
		var ret string
		return ret
	}
	return *o.IssuingAuthority
}

// GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetIssuingAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingAuthority) {
		return nil, false
	}
	return o.IssuingAuthority, true
}

// HasIssuingAuthority returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasIssuingAuthority() bool {
	if o != nil && !IsNil(o.IssuingAuthority) {
		return true
	}

	return false
}

// SetIssuingAuthority gets a reference to the given string and assigns it to the IssuingAuthority field.
func (o *OrganizationIdentificationMVO) SetIssuingAuthority(v string) {
	o.IssuingAuthority = &v
}

// GetIssuingDate returns the IssuingDate field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetIssuingDate() time.Time {
	if o == nil || IsNil(o.IssuingDate) {
		var ret time.Time
		return ret
	}
	return *o.IssuingDate
}

// GetIssuingDateOk returns a tuple with the IssuingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetIssuingDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.IssuingDate) {
		return nil, false
	}
	return o.IssuingDate, true
}

// HasIssuingDate returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasIssuingDate() bool {
	if o != nil && !IsNil(o.IssuingDate) {
		return true
	}

	return false
}

// SetIssuingDate gets a reference to the given time.Time and assigns it to the IssuingDate field.
func (o *OrganizationIdentificationMVO) SetIssuingDate(v time.Time) {
	o.IssuingDate = &v
}

// GetIdentificationType returns the IdentificationType field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetIdentificationType() string {
	if o == nil || IsNil(o.IdentificationType) {
		var ret string
		return ret
	}
	return *o.IdentificationType
}

// GetIdentificationTypeOk returns a tuple with the IdentificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetIdentificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentificationType) {
		return nil, false
	}
	return o.IdentificationType, true
}

// HasIdentificationType returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasIdentificationType() bool {
	if o != nil && !IsNil(o.IdentificationType) {
		return true
	}

	return false
}

// SetIdentificationType gets a reference to the given string and assigns it to the IdentificationType field.
func (o *OrganizationIdentificationMVO) SetIdentificationType(v string) {
	o.IdentificationType = &v
}

// GetValidFor returns the ValidFor field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetValidFor() TimePeriod {
	if o == nil || IsNil(o.ValidFor) {
		var ret TimePeriod
		return ret
	}
	return *o.ValidFor
}

// GetValidForOk returns a tuple with the ValidFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetValidForOk() (*TimePeriod, bool) {
	if o == nil || IsNil(o.ValidFor) {
		return nil, false
	}
	return o.ValidFor, true
}

// HasValidFor returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasValidFor() bool {
	if o != nil && !IsNil(o.ValidFor) {
		return true
	}

	return false
}

// SetValidFor gets a reference to the given TimePeriod and assigns it to the ValidFor field.
func (o *OrganizationIdentificationMVO) SetValidFor(v TimePeriod) {
	o.ValidFor = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *OrganizationIdentificationMVO) GetAttachment() AttachmentRefOrValueMVO {
	if o == nil || IsNil(o.Attachment) {
		var ret AttachmentRefOrValueMVO
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationIdentificationMVO) GetAttachmentOk() (*AttachmentRefOrValueMVO, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *OrganizationIdentificationMVO) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given AttachmentRefOrValueMVO and assigns it to the Attachment field.
func (o *OrganizationIdentificationMVO) SetAttachment(v AttachmentRefOrValueMVO) {
	o.Attachment = &v
}

func (o OrganizationIdentificationMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationIdentificationMVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	if !IsNil(o.IdentificationId) {
		toSerialize["identificationId"] = o.IdentificationId
	}
	if !IsNil(o.IssuingAuthority) {
		toSerialize["issuingAuthority"] = o.IssuingAuthority
	}
	if !IsNil(o.IssuingDate) {
		toSerialize["issuingDate"] = o.IssuingDate
	}
	if !IsNil(o.IdentificationType) {
		toSerialize["identificationType"] = o.IdentificationType
	}
	if !IsNil(o.ValidFor) {
		toSerialize["validFor"] = o.ValidFor
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	return toSerialize, nil
}

func (o *OrganizationIdentificationMVO) UnmarshalJSON(data []byte) (err error) {
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

	varOrganizationIdentificationMVO := _OrganizationIdentificationMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationIdentificationMVO)

	if err != nil {
		return err
	}

	*o = OrganizationIdentificationMVO(varOrganizationIdentificationMVO)

	return err
}

type NullableOrganizationIdentificationMVO struct {
	value *OrganizationIdentificationMVO
	isSet bool
}

func (v NullableOrganizationIdentificationMVO) Get() *OrganizationIdentificationMVO {
	return v.value
}

func (v *NullableOrganizationIdentificationMVO) Set(val *OrganizationIdentificationMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationIdentificationMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationIdentificationMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationIdentificationMVO(val *OrganizationIdentificationMVO) *NullableOrganizationIdentificationMVO {
	return &NullableOrganizationIdentificationMVO{value: val, isSet: true}
}

func (v NullableOrganizationIdentificationMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationIdentificationMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


