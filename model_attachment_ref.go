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

// checks if the AttachmentRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentRef{}

// AttachmentRef struct for AttachmentRef
type AttachmentRef struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// The URI of the referred entity.
	Href *string `json:"href,omitempty"`
	// The identifier of the referred entity.
	Id string `json:"id"`
	// Name of the referred entity.
	Name *string `json:"name,omitempty"`
	// The actual type of the target instance when needed for disambiguation.
	ReferredType *string `json:"@referredType,omitempty"`
	// A narrative text describing the content of the attachment
	Description *string `json:"description,omitempty"`
	// Link to the attachment media/content
	Url *string `json:"url,omitempty"`
}

type _AttachmentRef AttachmentRef

// NewAttachmentRef instantiates a new AttachmentRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentRef(type_ string, id string) *AttachmentRef {
	this := AttachmentRef{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAttachmentRefWithDefaults instantiates a new AttachmentRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentRefWithDefaults() *AttachmentRef {
	this := AttachmentRef{}
	return &this
}

// GetType returns the Type field value
func (o *AttachmentRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentRef) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *AttachmentRef) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *AttachmentRef) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *AttachmentRef) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *AttachmentRef) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *AttachmentRef) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *AttachmentRef) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *AttachmentRef) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *AttachmentRef) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *AttachmentRef) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *AttachmentRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentRef) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttachmentRef) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttachmentRef) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttachmentRef) SetName(v string) {
	o.Name = &v
}

// GetReferredType returns the ReferredType field value if set, zero value otherwise.
func (o *AttachmentRef) GetReferredType() string {
	if o == nil || IsNil(o.ReferredType) {
		var ret string
		return ret
	}
	return *o.ReferredType
}

// GetReferredTypeOk returns a tuple with the ReferredType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetReferredTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferredType) {
		return nil, false
	}
	return o.ReferredType, true
}

// HasReferredType returns a boolean if a field has been set.
func (o *AttachmentRef) HasReferredType() bool {
	if o != nil && !IsNil(o.ReferredType) {
		return true
	}

	return false
}

// SetReferredType gets a reference to the given string and assigns it to the ReferredType field.
func (o *AttachmentRef) SetReferredType(v string) {
	o.ReferredType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttachmentRef) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttachmentRef) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttachmentRef) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AttachmentRef) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentRef) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AttachmentRef) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AttachmentRef) SetUrl(v string) {
	o.Url = &v
}

func (o AttachmentRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["@type"] = o.Type
	if !IsNil(o.BaseType) {
		toSerialize["@baseType"] = o.BaseType
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["@schemaLocation"] = o.SchemaLocation
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReferredType) {
		toSerialize["@referredType"] = o.ReferredType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *AttachmentRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
		"id",
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

	varAttachmentRef := _AttachmentRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachmentRef)

	if err != nil {
		return err
	}

	*o = AttachmentRef(varAttachmentRef)

	return err
}

type NullableAttachmentRef struct {
	value *AttachmentRef
	isSet bool
}

func (v NullableAttachmentRef) Get() *AttachmentRef {
	return v.value
}

func (v *NullableAttachmentRef) Set(val *AttachmentRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentRef(val *AttachmentRef) *NullableAttachmentRef {
	return &NullableAttachmentRef{value: val, isSet: true}
}

func (v NullableAttachmentRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


