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

// checks if the IndividualDeleteEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndividualDeleteEvent{}

// IndividualDeleteEvent struct for IndividualDeleteEvent
type IndividualDeleteEvent struct {
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
	RelatedParty []RelatedPartyRefOrPartyRoleRef `json:"relatedParty,omitempty"`
	AnalyticCharacteristic []Characteristic `json:"analyticCharacteristic,omitempty"`
	// The identifier of the notification.
	EventId *string `json:"eventId,omitempty"`
	// Time of the event occurrence.
	EventTime *time.Time `json:"eventTime,omitempty"`
	// The type of the notification.
	EventType *string `json:"eventType,omitempty"`
	Event *IndividualDeleteEventPayload `json:"event,omitempty"`
}

type _IndividualDeleteEvent IndividualDeleteEvent

// NewIndividualDeleteEvent instantiates a new IndividualDeleteEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualDeleteEvent(type_ string) *IndividualDeleteEvent {
	this := IndividualDeleteEvent{}
	this.Type = type_
	return &this
}

// NewIndividualDeleteEventWithDefaults instantiates a new IndividualDeleteEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualDeleteEventWithDefaults() *IndividualDeleteEvent {
	this := IndividualDeleteEvent{}
	return &this
}

// GetType returns the Type field value
func (o *IndividualDeleteEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IndividualDeleteEvent) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *IndividualDeleteEvent) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *IndividualDeleteEvent) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IndividualDeleteEvent) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IndividualDeleteEvent) SetId(v string) {
	o.Id = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *IndividualDeleteEvent) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *IndividualDeleteEvent) SetDomain(v string) {
	o.Domain = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IndividualDeleteEvent) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IndividualDeleteEvent) SetDescription(v string) {
	o.Description = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *IndividualDeleteEvent) SetPriority(v string) {
	o.Priority = &v
}

// GetTimeOccurred returns the TimeOccurred field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetTimeOccurred() time.Time {
	if o == nil || IsNil(o.TimeOccurred) {
		var ret time.Time
		return ret
	}
	return *o.TimeOccurred
}

// GetTimeOccurredOk returns a tuple with the TimeOccurred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetTimeOccurredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeOccurred) {
		return nil, false
	}
	return o.TimeOccurred, true
}

// HasTimeOccurred returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasTimeOccurred() bool {
	if o != nil && !IsNil(o.TimeOccurred) {
		return true
	}

	return false
}

// SetTimeOccurred gets a reference to the given time.Time and assigns it to the TimeOccurred field.
func (o *IndividualDeleteEvent) SetTimeOccurred(v time.Time) {
	o.TimeOccurred = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetSource() EntityRef {
	if o == nil || IsNil(o.Source) {
		var ret EntityRef
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetSourceOk() (*EntityRef, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given EntityRef and assigns it to the Source field.
func (o *IndividualDeleteEvent) SetSource(v EntityRef) {
	o.Source = &v
}

// GetReportingSystem returns the ReportingSystem field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetReportingSystem() EntityRef {
	if o == nil || IsNil(o.ReportingSystem) {
		var ret EntityRef
		return ret
	}
	return *o.ReportingSystem
}

// GetReportingSystemOk returns a tuple with the ReportingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetReportingSystemOk() (*EntityRef, bool) {
	if o == nil || IsNil(o.ReportingSystem) {
		return nil, false
	}
	return o.ReportingSystem, true
}

// HasReportingSystem returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasReportingSystem() bool {
	if o != nil && !IsNil(o.ReportingSystem) {
		return true
	}

	return false
}

// SetReportingSystem gets a reference to the given EntityRef and assigns it to the ReportingSystem field.
func (o *IndividualDeleteEvent) SetReportingSystem(v EntityRef) {
	o.ReportingSystem = &v
}

// GetRelatedParty returns the RelatedParty field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetRelatedParty() []RelatedPartyRefOrPartyRoleRef {
	if o == nil || IsNil(o.RelatedParty) {
		var ret []RelatedPartyRefOrPartyRoleRef
		return ret
	}
	return o.RelatedParty
}

// GetRelatedPartyOk returns a tuple with the RelatedParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetRelatedPartyOk() ([]RelatedPartyRefOrPartyRoleRef, bool) {
	if o == nil || IsNil(o.RelatedParty) {
		return nil, false
	}
	return o.RelatedParty, true
}

// HasRelatedParty returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasRelatedParty() bool {
	if o != nil && !IsNil(o.RelatedParty) {
		return true
	}

	return false
}

// SetRelatedParty gets a reference to the given []RelatedPartyRefOrPartyRoleRef and assigns it to the RelatedParty field.
func (o *IndividualDeleteEvent) SetRelatedParty(v []RelatedPartyRefOrPartyRoleRef) {
	o.RelatedParty = v
}

// GetAnalyticCharacteristic returns the AnalyticCharacteristic field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetAnalyticCharacteristic() []Characteristic {
	if o == nil || IsNil(o.AnalyticCharacteristic) {
		var ret []Characteristic
		return ret
	}
	return o.AnalyticCharacteristic
}

// GetAnalyticCharacteristicOk returns a tuple with the AnalyticCharacteristic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetAnalyticCharacteristicOk() ([]Characteristic, bool) {
	if o == nil || IsNil(o.AnalyticCharacteristic) {
		return nil, false
	}
	return o.AnalyticCharacteristic, true
}

// HasAnalyticCharacteristic returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasAnalyticCharacteristic() bool {
	if o != nil && !IsNil(o.AnalyticCharacteristic) {
		return true
	}

	return false
}

// SetAnalyticCharacteristic gets a reference to the given []Characteristic and assigns it to the AnalyticCharacteristic field.
func (o *IndividualDeleteEvent) SetAnalyticCharacteristic(v []Characteristic) {
	o.AnalyticCharacteristic = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *IndividualDeleteEvent) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *IndividualDeleteEvent) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *IndividualDeleteEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *IndividualDeleteEvent) GetEvent() IndividualDeleteEventPayload {
	if o == nil || IsNil(o.Event) {
		var ret IndividualDeleteEventPayload
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualDeleteEvent) GetEventOk() (*IndividualDeleteEventPayload, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *IndividualDeleteEvent) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given IndividualDeleteEventPayload and assigns it to the Event field.
func (o *IndividualDeleteEvent) SetEvent(v IndividualDeleteEventPayload) {
	o.Event = &v
}

func (o IndividualDeleteEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualDeleteEvent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	return toSerialize, nil
}

func (o *IndividualDeleteEvent) UnmarshalJSON(data []byte) (err error) {
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

	varIndividualDeleteEvent := _IndividualDeleteEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIndividualDeleteEvent)

	if err != nil {
		return err
	}

	*o = IndividualDeleteEvent(varIndividualDeleteEvent)

	return err
}

type NullableIndividualDeleteEvent struct {
	value *IndividualDeleteEvent
	isSet bool
}

func (v NullableIndividualDeleteEvent) Get() *IndividualDeleteEvent {
	return v.value
}

func (v *NullableIndividualDeleteEvent) Set(val *IndividualDeleteEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualDeleteEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualDeleteEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualDeleteEvent(val *IndividualDeleteEvent) *NullableIndividualDeleteEvent {
	return &NullableIndividualDeleteEvent{value: val, isSet: true}
}

func (v NullableIndividualDeleteEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualDeleteEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


