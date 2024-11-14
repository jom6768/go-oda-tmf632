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

// checks if the EventMVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMVO{}

// EventMVO struct for EventMVO
type EventMVO struct {
	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type"`
	// When sub-classing, this defines the super-class
	BaseType *string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation *string `json:"@schemaLocation,omitempty"`
	// Hyperlink reference
	Href *string `json:"href,omitempty"`
	// unique identifier
	Id *string `json:"id,omitempty"`
	// The correlation id for this event.
	CorrelationId *string `json:"correlationId,omitempty"`
	// The domain of the event.
	Domain *string `json:"domain,omitempty"`
	// The title of the event.
	Title *string `json:"title,omitempty"`
	// An explnatory of the event.
	Description *string `json:"description,omitempty"`
	// A priority.
	Priority *string `json:"priority,omitempty"`
	// The time the event occurred.
	TimeOccurred *time.Time `json:"timeOccurred,omitempty"`
	Source *EntityRef `json:"source,omitempty"`
	ReportingSystem *EntityRef `json:"reportingSystem,omitempty"`
	RelatedParty []RelatedPartyRefOrPartyRoleRefMVO `json:"relatedParty,omitempty"`
	AnalyticCharacteristic []CharacteristicMVO `json:"analyticCharacteristic,omitempty"`
	// The identifier of the notification.
	EventId string `json:"eventId"`
	// Time of the event occurrence.
	EventTime time.Time `json:"eventTime"`
	// The type of the notification.
	EventType string `json:"eventType"`
	// The event linked to the involved resource object
	Event map[string]interface{} `json:"event"`
}

type _EventMVO EventMVO

// NewEventMVO instantiates a new EventMVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMVO(type_ string, eventId string, eventTime time.Time, eventType string, event map[string]interface{}) *EventMVO {
	this := EventMVO{}
	this.Type = type_
	this.EventId = eventId
	this.EventTime = eventTime
	this.EventType = eventType
	this.Event = event
	return &this
}

// NewEventMVOWithDefaults instantiates a new EventMVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMVOWithDefaults() *EventMVO {
	this := EventMVO{}
	return &this
}

// GetType returns the Type field value
func (o *EventMVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventMVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventMVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *EventMVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *EventMVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *EventMVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *EventMVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *EventMVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *EventMVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *EventMVO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *EventMVO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *EventMVO) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventMVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventMVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventMVO) SetId(v string) {
	o.Id = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *EventMVO) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *EventMVO) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *EventMVO) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *EventMVO) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *EventMVO) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *EventMVO) SetDomain(v string) {
	o.Domain = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EventMVO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EventMVO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EventMVO) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventMVO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventMVO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventMVO) SetDescription(v string) {
	o.Description = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *EventMVO) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *EventMVO) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *EventMVO) SetPriority(v string) {
	o.Priority = &v
}

// GetTimeOccurred returns the TimeOccurred field value if set, zero value otherwise.
func (o *EventMVO) GetTimeOccurred() time.Time {
	if o == nil || IsNil(o.TimeOccurred) {
		var ret time.Time
		return ret
	}
	return *o.TimeOccurred
}

// GetTimeOccurredOk returns a tuple with the TimeOccurred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetTimeOccurredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeOccurred) {
		return nil, false
	}
	return o.TimeOccurred, true
}

// HasTimeOccurred returns a boolean if a field has been set.
func (o *EventMVO) HasTimeOccurred() bool {
	if o != nil && !IsNil(o.TimeOccurred) {
		return true
	}

	return false
}

// SetTimeOccurred gets a reference to the given time.Time and assigns it to the TimeOccurred field.
func (o *EventMVO) SetTimeOccurred(v time.Time) {
	o.TimeOccurred = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EventMVO) GetSource() EntityRef {
	if o == nil || IsNil(o.Source) {
		var ret EntityRef
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetSourceOk() (*EntityRef, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EventMVO) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given EntityRef and assigns it to the Source field.
func (o *EventMVO) SetSource(v EntityRef) {
	o.Source = &v
}

// GetReportingSystem returns the ReportingSystem field value if set, zero value otherwise.
func (o *EventMVO) GetReportingSystem() EntityRef {
	if o == nil || IsNil(o.ReportingSystem) {
		var ret EntityRef
		return ret
	}
	return *o.ReportingSystem
}

// GetReportingSystemOk returns a tuple with the ReportingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetReportingSystemOk() (*EntityRef, bool) {
	if o == nil || IsNil(o.ReportingSystem) {
		return nil, false
	}
	return o.ReportingSystem, true
}

// HasReportingSystem returns a boolean if a field has been set.
func (o *EventMVO) HasReportingSystem() bool {
	if o != nil && !IsNil(o.ReportingSystem) {
		return true
	}

	return false
}

// SetReportingSystem gets a reference to the given EntityRef and assigns it to the ReportingSystem field.
func (o *EventMVO) SetReportingSystem(v EntityRef) {
	o.ReportingSystem = &v
}

// GetRelatedParty returns the RelatedParty field value if set, zero value otherwise.
func (o *EventMVO) GetRelatedParty() []RelatedPartyRefOrPartyRoleRefMVO {
	if o == nil || IsNil(o.RelatedParty) {
		var ret []RelatedPartyRefOrPartyRoleRefMVO
		return ret
	}
	return o.RelatedParty
}

// GetRelatedPartyOk returns a tuple with the RelatedParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetRelatedPartyOk() ([]RelatedPartyRefOrPartyRoleRefMVO, bool) {
	if o == nil || IsNil(o.RelatedParty) {
		return nil, false
	}
	return o.RelatedParty, true
}

// HasRelatedParty returns a boolean if a field has been set.
func (o *EventMVO) HasRelatedParty() bool {
	if o != nil && !IsNil(o.RelatedParty) {
		return true
	}

	return false
}

// SetRelatedParty gets a reference to the given []RelatedPartyRefOrPartyRoleRefMVO and assigns it to the RelatedParty field.
func (o *EventMVO) SetRelatedParty(v []RelatedPartyRefOrPartyRoleRefMVO) {
	o.RelatedParty = v
}

// GetAnalyticCharacteristic returns the AnalyticCharacteristic field value if set, zero value otherwise.
func (o *EventMVO) GetAnalyticCharacteristic() []CharacteristicMVO {
	if o == nil || IsNil(o.AnalyticCharacteristic) {
		var ret []CharacteristicMVO
		return ret
	}
	return o.AnalyticCharacteristic
}

// GetAnalyticCharacteristicOk returns a tuple with the AnalyticCharacteristic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventMVO) GetAnalyticCharacteristicOk() ([]CharacteristicMVO, bool) {
	if o == nil || IsNil(o.AnalyticCharacteristic) {
		return nil, false
	}
	return o.AnalyticCharacteristic, true
}

// HasAnalyticCharacteristic returns a boolean if a field has been set.
func (o *EventMVO) HasAnalyticCharacteristic() bool {
	if o != nil && !IsNil(o.AnalyticCharacteristic) {
		return true
	}

	return false
}

// SetAnalyticCharacteristic gets a reference to the given []CharacteristicMVO and assigns it to the AnalyticCharacteristic field.
func (o *EventMVO) SetAnalyticCharacteristic(v []CharacteristicMVO) {
	o.AnalyticCharacteristic = v
}

// GetEventId returns the EventId field value
func (o *EventMVO) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *EventMVO) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *EventMVO) SetEventId(v string) {
	o.EventId = v
}

// GetEventTime returns the EventTime field value
func (o *EventMVO) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *EventMVO) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *EventMVO) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetEventType returns the EventType field value
func (o *EventMVO) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EventMVO) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EventMVO) SetEventType(v string) {
	o.EventType = v
}

// GetEvent returns the Event field value
func (o *EventMVO) GetEvent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventMVO) GetEventOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Event, true
}

// SetEvent sets field value
func (o *EventMVO) SetEvent(v map[string]interface{}) {
	o.Event = v
}

func (o EventMVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMVO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.TimeOccurred) {
		toSerialize["timeOccurred"] = o.TimeOccurred
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.ReportingSystem) {
		toSerialize["reportingSystem"] = o.ReportingSystem
	}
	if !IsNil(o.RelatedParty) {
		toSerialize["relatedParty"] = o.RelatedParty
	}
	if !IsNil(o.AnalyticCharacteristic) {
		toSerialize["analyticCharacteristic"] = o.AnalyticCharacteristic
	}
	toSerialize["eventId"] = o.EventId
	toSerialize["eventTime"] = o.EventTime
	toSerialize["eventType"] = o.EventType
	toSerialize["event"] = o.Event
	return toSerialize, nil
}

func (o *EventMVO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"@type",
		"eventId",
		"eventTime",
		"eventType",
		"event",
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

	varEventMVO := _EventMVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventMVO)

	if err != nil {
		return err
	}

	*o = EventMVO(varEventMVO)

	return err
}

type NullableEventMVO struct {
	value *EventMVO
	isSet bool
}

func (v NullableEventMVO) Get() *EventMVO {
	return v.value
}

func (v *NullableEventMVO) Set(val *EventMVO) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMVO) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMVO(val *EventMVO) *NullableEventMVO {
	return &NullableEventMVO{value: val, isSet: true}
}

func (v NullableEventMVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


