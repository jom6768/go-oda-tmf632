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

// checks if the GeographicAddressContactMediumFVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeographicAddressContactMediumFVO{}

// GeographicAddressContactMediumFVO struct for GeographicAddressContactMediumFVO
type GeographicAddressContactMediumFVO struct {
	ContactMediumFVO
	// The city
	City *string `json:"city,omitempty"`
	// The country
	Country *string `json:"country,omitempty"`
	// Postcode
	PostCode *string `json:"postCode,omitempty"`
	// State or province
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// Describes the street
	Street1 *string `json:"street1,omitempty"`
	// Complementary street description
	Street2 *string `json:"street2,omitempty"`
	GeographicAddress *GeographicAddressRefFVO `json:"geographicAddress,omitempty"`
}

type _GeographicAddressContactMediumFVO GeographicAddressContactMediumFVO

// NewGeographicAddressContactMediumFVO instantiates a new GeographicAddressContactMediumFVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicAddressContactMediumFVO(type_ string) *GeographicAddressContactMediumFVO {
	this := GeographicAddressContactMediumFVO{}
	this.Type = type_
	return &this
}

// NewGeographicAddressContactMediumFVOWithDefaults instantiates a new GeographicAddressContactMediumFVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicAddressContactMediumFVOWithDefaults() *GeographicAddressContactMediumFVO {
	this := GeographicAddressContactMediumFVO{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GeographicAddressContactMediumFVO) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GeographicAddressContactMediumFVO) SetCountry(v string) {
	o.Country = &v
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetPostCode() string {
	if o == nil || IsNil(o.PostCode) {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetPostCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasPostCode() bool {
	if o != nil && !IsNil(o.PostCode) {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *GeographicAddressContactMediumFVO) SetPostCode(v string) {
	o.PostCode = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetStateOrProvince() string {
	if o == nil || IsNil(o.StateOrProvince) {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.StateOrProvince) {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasStateOrProvince() bool {
	if o != nil && !IsNil(o.StateOrProvince) {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *GeographicAddressContactMediumFVO) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetStreet1() string {
	if o == nil || IsNil(o.Street1) {
		var ret string
		return ret
	}
	return *o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetStreet1Ok() (*string, bool) {
	if o == nil || IsNil(o.Street1) {
		return nil, false
	}
	return o.Street1, true
}

// HasStreet1 returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasStreet1() bool {
	if o != nil && !IsNil(o.Street1) {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given string and assigns it to the Street1 field.
func (o *GeographicAddressContactMediumFVO) SetStreet1(v string) {
	o.Street1 = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *GeographicAddressContactMediumFVO) SetStreet2(v string) {
	o.Street2 = &v
}

// GetGeographicAddress returns the GeographicAddress field value if set, zero value otherwise.
func (o *GeographicAddressContactMediumFVO) GetGeographicAddress() GeographicAddressRefFVO {
	if o == nil || IsNil(o.GeographicAddress) {
		var ret GeographicAddressRefFVO
		return ret
	}
	return *o.GeographicAddress
}

// GetGeographicAddressOk returns a tuple with the GeographicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicAddressContactMediumFVO) GetGeographicAddressOk() (*GeographicAddressRefFVO, bool) {
	if o == nil || IsNil(o.GeographicAddress) {
		return nil, false
	}
	return o.GeographicAddress, true
}

// HasGeographicAddress returns a boolean if a field has been set.
func (o *GeographicAddressContactMediumFVO) HasGeographicAddress() bool {
	if o != nil && !IsNil(o.GeographicAddress) {
		return true
	}

	return false
}

// SetGeographicAddress gets a reference to the given GeographicAddressRefFVO and assigns it to the GeographicAddress field.
func (o *GeographicAddressContactMediumFVO) SetGeographicAddress(v GeographicAddressRefFVO) {
	o.GeographicAddress = &v
}

func (o GeographicAddressContactMediumFVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeographicAddressContactMediumFVO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContactMediumFVO, errContactMediumFVO := json.Marshal(o.ContactMediumFVO)
	if errContactMediumFVO != nil {
		return map[string]interface{}{}, errContactMediumFVO
	}
	errContactMediumFVO = json.Unmarshal([]byte(serializedContactMediumFVO), &toSerialize)
	if errContactMediumFVO != nil {
		return map[string]interface{}{}, errContactMediumFVO
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostCode) {
		toSerialize["postCode"] = o.PostCode
	}
	if !IsNil(o.StateOrProvince) {
		toSerialize["stateOrProvince"] = o.StateOrProvince
	}
	if !IsNil(o.Street1) {
		toSerialize["street1"] = o.Street1
	}
	if !IsNil(o.Street2) {
		toSerialize["street2"] = o.Street2
	}
	if !IsNil(o.GeographicAddress) {
		toSerialize["geographicAddress"] = o.GeographicAddress
	}
	return toSerialize, nil
}

func (o *GeographicAddressContactMediumFVO) UnmarshalJSON(data []byte) (err error) {
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

	varGeographicAddressContactMediumFVO := _GeographicAddressContactMediumFVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGeographicAddressContactMediumFVO)

	if err != nil {
		return err
	}

	*o = GeographicAddressContactMediumFVO(varGeographicAddressContactMediumFVO)

	return err
}

type NullableGeographicAddressContactMediumFVO struct {
	value *GeographicAddressContactMediumFVO
	isSet bool
}

func (v NullableGeographicAddressContactMediumFVO) Get() *GeographicAddressContactMediumFVO {
	return v.value
}

func (v *NullableGeographicAddressContactMediumFVO) Set(val *GeographicAddressContactMediumFVO) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicAddressContactMediumFVO) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicAddressContactMediumFVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicAddressContactMediumFVO(val *GeographicAddressContactMediumFVO) *NullableGeographicAddressContactMediumFVO {
	return &NullableGeographicAddressContactMediumFVO{value: val, isSet: true}
}

func (v NullableGeographicAddressContactMediumFVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicAddressContactMediumFVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

