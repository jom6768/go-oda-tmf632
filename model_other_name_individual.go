/*
Party Management

TMF API Reference : TMF 632 - Party  Release: 22.5 The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role. For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information. Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

API version: 5.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OtherNameIndividual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherNameIndividual{}

// OtherNameIndividual Keeps track of other names, for example the old name of a woman before marriage or an artist name.
type OtherNameIndividual struct {
	// Use for titles (aristrocatic, social, ...): Pr, Dr, Sir,....
	Title *string `json:"title,omitempty"`
	// e.g. Baron, Graf, Earl, etc.
	AristocraticTitle *string `json:"aristocraticTitle,omitempty"`
	// e.g. Sr, Jr, etc.
	Generation *string `json:"generation,omitempty"`
	// First name
	GivenName *string `json:"givenName,omitempty"`
	// Contains the chosen name by which the person prefers to be addressed. Note: This name may be a name other than a given name, such as a nickname
	PreferredGivenName *string `json:"preferredGivenName,omitempty"`
	// Family name prefix
	FamilyNamePrefix *string `json:"familyNamePrefix,omitempty"`
	// Contains the non-chosen or inherited name. Also known as last name in the Western context
	FamilyName *string `json:"familyName,omitempty"`
	// Legal name or birth name (name one has for official purposes)
	LegalName *string `json:"legalName,omitempty"`
	// Middle name or initial
	MiddleName *string `json:"middleName,omitempty"`
	// Full name flatten (first, middle, and last names)
	FullName *string `json:"fullName,omitempty"`
	// . A fully formatted name in one string with all of its pieces in their proper place and all of the necessary punctuation. Useful for specific contexts (Chinese, Japanese, Korean, etc.)
	FormattedName *string `json:"formattedName,omitempty"`
	ValidFor *TimePeriod `json:"validFor,omitempty"`
}

// NewOtherNameIndividual instantiates a new OtherNameIndividual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherNameIndividual() *OtherNameIndividual {
	this := OtherNameIndividual{}
	return &this
}

// NewOtherNameIndividualWithDefaults instantiates a new OtherNameIndividual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherNameIndividualWithDefaults() *OtherNameIndividual {
	this := OtherNameIndividual{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OtherNameIndividual) SetTitle(v string) {
	o.Title = &v
}

// GetAristocraticTitle returns the AristocraticTitle field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetAristocraticTitle() string {
	if o == nil || IsNil(o.AristocraticTitle) {
		var ret string
		return ret
	}
	return *o.AristocraticTitle
}

// GetAristocraticTitleOk returns a tuple with the AristocraticTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetAristocraticTitleOk() (*string, bool) {
	if o == nil || IsNil(o.AristocraticTitle) {
		return nil, false
	}
	return o.AristocraticTitle, true
}

// HasAristocraticTitle returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasAristocraticTitle() bool {
	if o != nil && !IsNil(o.AristocraticTitle) {
		return true
	}

	return false
}

// SetAristocraticTitle gets a reference to the given string and assigns it to the AristocraticTitle field.
func (o *OtherNameIndividual) SetAristocraticTitle(v string) {
	o.AristocraticTitle = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetGeneration() string {
	if o == nil || IsNil(o.Generation) {
		var ret string
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetGenerationOk() (*string, bool) {
	if o == nil || IsNil(o.Generation) {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasGeneration() bool {
	if o != nil && !IsNil(o.Generation) {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given string and assigns it to the Generation field.
func (o *OtherNameIndividual) SetGeneration(v string) {
	o.Generation = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *OtherNameIndividual) SetGivenName(v string) {
	o.GivenName = &v
}

// GetPreferredGivenName returns the PreferredGivenName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetPreferredGivenName() string {
	if o == nil || IsNil(o.PreferredGivenName) {
		var ret string
		return ret
	}
	return *o.PreferredGivenName
}

// GetPreferredGivenNameOk returns a tuple with the PreferredGivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetPreferredGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredGivenName) {
		return nil, false
	}
	return o.PreferredGivenName, true
}

// HasPreferredGivenName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasPreferredGivenName() bool {
	if o != nil && !IsNil(o.PreferredGivenName) {
		return true
	}

	return false
}

// SetPreferredGivenName gets a reference to the given string and assigns it to the PreferredGivenName field.
func (o *OtherNameIndividual) SetPreferredGivenName(v string) {
	o.PreferredGivenName = &v
}

// GetFamilyNamePrefix returns the FamilyNamePrefix field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetFamilyNamePrefix() string {
	if o == nil || IsNil(o.FamilyNamePrefix) {
		var ret string
		return ret
	}
	return *o.FamilyNamePrefix
}

// GetFamilyNamePrefixOk returns a tuple with the FamilyNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetFamilyNamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyNamePrefix) {
		return nil, false
	}
	return o.FamilyNamePrefix, true
}

// HasFamilyNamePrefix returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasFamilyNamePrefix() bool {
	if o != nil && !IsNil(o.FamilyNamePrefix) {
		return true
	}

	return false
}

// SetFamilyNamePrefix gets a reference to the given string and assigns it to the FamilyNamePrefix field.
func (o *OtherNameIndividual) SetFamilyNamePrefix(v string) {
	o.FamilyNamePrefix = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *OtherNameIndividual) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *OtherNameIndividual) SetLegalName(v string) {
	o.LegalName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *OtherNameIndividual) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *OtherNameIndividual) SetFullName(v string) {
	o.FullName = &v
}

// GetFormattedName returns the FormattedName field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetFormattedName() string {
	if o == nil || IsNil(o.FormattedName) {
		var ret string
		return ret
	}
	return *o.FormattedName
}

// GetFormattedNameOk returns a tuple with the FormattedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetFormattedNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedName) {
		return nil, false
	}
	return o.FormattedName, true
}

// HasFormattedName returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasFormattedName() bool {
	if o != nil && !IsNil(o.FormattedName) {
		return true
	}

	return false
}

// SetFormattedName gets a reference to the given string and assigns it to the FormattedName field.
func (o *OtherNameIndividual) SetFormattedName(v string) {
	o.FormattedName = &v
}

// GetValidFor returns the ValidFor field value if set, zero value otherwise.
func (o *OtherNameIndividual) GetValidFor() TimePeriod {
	if o == nil || IsNil(o.ValidFor) {
		var ret TimePeriod
		return ret
	}
	return *o.ValidFor
}

// GetValidForOk returns a tuple with the ValidFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtherNameIndividual) GetValidForOk() (*TimePeriod, bool) {
	if o == nil || IsNil(o.ValidFor) {
		return nil, false
	}
	return o.ValidFor, true
}

// HasValidFor returns a boolean if a field has been set.
func (o *OtherNameIndividual) HasValidFor() bool {
	if o != nil && !IsNil(o.ValidFor) {
		return true
	}

	return false
}

// SetValidFor gets a reference to the given TimePeriod and assigns it to the ValidFor field.
func (o *OtherNameIndividual) SetValidFor(v TimePeriod) {
	o.ValidFor = &v
}

func (o OtherNameIndividual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtherNameIndividual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.AristocraticTitle) {
		toSerialize["aristocraticTitle"] = o.AristocraticTitle
	}
	if !IsNil(o.Generation) {
		toSerialize["generation"] = o.Generation
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.PreferredGivenName) {
		toSerialize["preferredGivenName"] = o.PreferredGivenName
	}
	if !IsNil(o.FamilyNamePrefix) {
		toSerialize["familyNamePrefix"] = o.FamilyNamePrefix
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.LegalName) {
		toSerialize["legalName"] = o.LegalName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.FormattedName) {
		toSerialize["formattedName"] = o.FormattedName
	}
	if !IsNil(o.ValidFor) {
		toSerialize["validFor"] = o.ValidFor
	}
	return toSerialize, nil
}

type NullableOtherNameIndividual struct {
	value *OtherNameIndividual
	isSet bool
}

func (v NullableOtherNameIndividual) Get() *OtherNameIndividual {
	return v.value
}

func (v *NullableOtherNameIndividual) Set(val *OtherNameIndividual) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherNameIndividual) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherNameIndividual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherNameIndividual(val *OtherNameIndividual) *NullableOtherNameIndividual {
	return &NullableOtherNameIndividual{value: val, isSet: true}
}

func (v NullableOtherNameIndividual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherNameIndividual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


