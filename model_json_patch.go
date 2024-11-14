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

// checks if the JsonPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonPatch{}

// JsonPatch A JSONPatch document as defined by RFC 6902
type JsonPatch struct {
	// The operation to be performed
	Op string `json:"op"`
	// A JSON-Pointer
	Path string `json:"path"`
	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
	// A string containing a JSON Pointer value.
	From *string `json:"from,omitempty"`
}

type _JsonPatch JsonPatch

// NewJsonPatch instantiates a new JsonPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatch(op string, path string) *JsonPatch {
	this := JsonPatch{}
	this.Op = op
	this.Path = path
	return &this
}

// NewJsonPatchWithDefaults instantiates a new JsonPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchWithDefaults() *JsonPatch {
	this := JsonPatch{}
	return &this
}

// GetOp returns the Op field value
func (o *JsonPatch) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatch) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatch) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatch) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JsonPatch) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JsonPatch) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JsonPatch) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *JsonPatch) SetValue(v interface{}) {
	o.Value = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *JsonPatch) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *JsonPatch) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *JsonPatch) SetFrom(v string) {
	o.From = &v
}

func (o JsonPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	return toSerialize, nil
}

func (o *JsonPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
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

	varJsonPatch := _JsonPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJsonPatch)

	if err != nil {
		return err
	}

	*o = JsonPatch(varJsonPatch)

	return err
}

type NullableJsonPatch struct {
	value *JsonPatch
	isSet bool
}

func (v NullableJsonPatch) Get() *JsonPatch {
	return v.value
}

func (v *NullableJsonPatch) Set(val *JsonPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatch(val *JsonPatch) *NullableJsonPatch {
	return &NullableJsonPatch{value: val, isSet: true}
}

func (v NullableJsonPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


