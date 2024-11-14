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

// checks if the BaseEventFVO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEventFVO{}

// BaseEventFVO struct for BaseEventFVO
type BaseEventFVO struct {
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
	Event map[string]interface{} `json:"event,omitempty"`
	// The identifier of the notification.
	EventId *string `json:"eventId,omitempty"`
	// Time of the event occurrence.
	EventTime *time.Time `json:"eventTime,omitempty"`
	// The type of the notification.
	EventType *string `json:"eventType,omitempty"`
	// The correlation id for this event.
	CorrelationId *string `json:"correlationId,omitempty"`
	// The domain of the event.
	Domain *string `json:"domain,omitempty"`
	// The title of the event.
	Title *string `json:"title,omitempty"`
	// An explanatory of the event.
	Description *string `json:"description,omitempty"`
	// A priority.
	Priority *string `json:"priority,omitempty"`
	// The time the event occured.
	TimeOcurred *time.Time `json:"timeOcurred,omitempty"`
}

type _BaseEventFVO BaseEventFVO

// NewBaseEventFVO instantiates a new BaseEventFVO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventFVO(type_ string) *BaseEventFVO {
	this := BaseEventFVO{}
	this.Type = type_
	return &this
}

// NewBaseEventFVOWithDefaults instantiates a new BaseEventFVO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventFVOWithDefaults() *BaseEventFVO {
	this := BaseEventFVO{}
	return &this
}

// GetType returns the Type field value
func (o *BaseEventFVO) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseEventFVO) SetType(v string) {
	o.Type = v
}

// GetBaseType returns the BaseType field value if set, zero value otherwise.
func (o *BaseEventFVO) GetBaseType() string {
	if o == nil || IsNil(o.BaseType) {
		var ret string
		return ret
	}
	return *o.BaseType
}

// GetBaseTypeOk returns a tuple with the BaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetBaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BaseType) {
		return nil, false
	}
	return o.BaseType, true
}

// HasBaseType returns a boolean if a field has been set.
func (o *BaseEventFVO) HasBaseType() bool {
	if o != nil && !IsNil(o.BaseType) {
		return true
	}

	return false
}

// SetBaseType gets a reference to the given string and assigns it to the BaseType field.
func (o *BaseEventFVO) SetBaseType(v string) {
	o.BaseType = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *BaseEventFVO) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *BaseEventFVO) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *BaseEventFVO) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BaseEventFVO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BaseEventFVO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BaseEventFVO) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseEventFVO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseEventFVO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseEventFVO) SetId(v string) {
	o.Id = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BaseEventFVO) GetEvent() map[string]interface{} {
	if o == nil || IsNil(o.Event) {
		var ret map[string]interface{}
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetEventOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Event) {
		return map[string]interface{}{}, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BaseEventFVO) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given map[string]interface{} and assigns it to the Event field.
func (o *BaseEventFVO) SetEvent(v map[string]interface{}) {
	o.Event = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *BaseEventFVO) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *BaseEventFVO) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *BaseEventFVO) SetEventId(v string) {
	o.EventId = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *BaseEventFVO) GetEventTime() time.Time {
	if o == nil || IsNil(o.EventTime) {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *BaseEventFVO) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *BaseEventFVO) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseEventFVO) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseEventFVO) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseEventFVO) SetEventType(v string) {
	o.EventType = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *BaseEventFVO) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *BaseEventFVO) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *BaseEventFVO) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *BaseEventFVO) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *BaseEventFVO) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *BaseEventFVO) SetDomain(v string) {
	o.Domain = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BaseEventFVO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BaseEventFVO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BaseEventFVO) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseEventFVO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseEventFVO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseEventFVO) SetDescription(v string) {
	o.Description = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *BaseEventFVO) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *BaseEventFVO) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *BaseEventFVO) SetPriority(v string) {
	o.Priority = &v
}

// GetTimeOcurred returns the TimeOcurred field value if set, zero value otherwise.
func (o *BaseEventFVO) GetTimeOcurred() time.Time {
	if o == nil || IsNil(o.TimeOcurred) {
		var ret time.Time
		return ret
	}
	return *o.TimeOcurred
}

// GetTimeOcurredOk returns a tuple with the TimeOcurred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventFVO) GetTimeOcurredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeOcurred) {
		return nil, false
	}
	return o.TimeOcurred, true
}

// HasTimeOcurred returns a boolean if a field has been set.
func (o *BaseEventFVO) HasTimeOcurred() bool {
	if o != nil && !IsNil(o.TimeOcurred) {
		return true
	}

	return false
}

// SetTimeOcurred gets a reference to the given time.Time and assigns it to the TimeOcurred field.
func (o *BaseEventFVO) SetTimeOcurred(v time.Time) {
	o.TimeOcurred = &v
}

func (o BaseEventFVO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEventFVO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
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
	if !IsNil(o.TimeOcurred) {
		toSerialize["timeOcurred"] = o.TimeOcurred
	}
	return toSerialize, nil
}

func (o *BaseEventFVO) UnmarshalJSON(data []byte) (err error) {
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

	varBaseEventFVO := _BaseEventFVO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseEventFVO)

	if err != nil {
		return err
	}

	*o = BaseEventFVO(varBaseEventFVO)

	return err
}

type NullableBaseEventFVO struct {
	value *BaseEventFVO
	isSet bool
}

func (v NullableBaseEventFVO) Get() *BaseEventFVO {
	return v.value
}

func (v *NullableBaseEventFVO) Set(val *BaseEventFVO) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventFVO) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventFVO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventFVO(val *BaseEventFVO) *NullableBaseEventFVO {
	return &NullableBaseEventFVO{value: val, isSet: true}
}

func (v NullableBaseEventFVO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventFVO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


