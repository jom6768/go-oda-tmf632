# OrganizationAttributeValueChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**CorrelationId** | Pointer to **string** | The correlation id for this event. | [optional] 
**Domain** | Pointer to **string** | The domain of the event. | [optional] 
**Title** | Pointer to **string** | The title of the event. | [optional] 
**Description** | Pointer to **string** | An explnatory of the event. | [optional] 
**Priority** | Pointer to **string** | A priority. | [optional] 
**TimeOccurred** | Pointer to **time.Time** | The time the event occurred. | [optional] 
**Source** | Pointer to [**EntityRef**](EntityRef.md) |  | [optional] 
**ReportingSystem** | Pointer to [**EntityRef**](EntityRef.md) |  | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyRefOrPartyRoleRef**](RelatedPartyRefOrPartyRoleRef.md) |  | [optional] 
**AnalyticCharacteristic** | Pointer to [**[]Characteristic**](Characteristic.md) |  | [optional] 
**EventId** | Pointer to **string** | The identifier of the notification. | [optional] 
**EventTime** | Pointer to **time.Time** | Time of the event occurrence. | [optional] 
**EventType** | Pointer to **string** | The type of the notification. | [optional] 
**Event** | Pointer to [**OrganizationAttributeValueChangeEventPayload**](OrganizationAttributeValueChangeEventPayload.md) |  | [optional] 

## Methods

### NewOrganizationAttributeValueChangeEvent

`func NewOrganizationAttributeValueChangeEvent(type_ string, ) *OrganizationAttributeValueChangeEvent`

NewOrganizationAttributeValueChangeEvent instantiates a new OrganizationAttributeValueChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAttributeValueChangeEventWithDefaults

`func NewOrganizationAttributeValueChangeEventWithDefaults() *OrganizationAttributeValueChangeEvent`

NewOrganizationAttributeValueChangeEventWithDefaults instantiates a new OrganizationAttributeValueChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationAttributeValueChangeEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationAttributeValueChangeEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationAttributeValueChangeEvent) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OrganizationAttributeValueChangeEvent) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OrganizationAttributeValueChangeEvent) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OrganizationAttributeValueChangeEvent) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OrganizationAttributeValueChangeEvent) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OrganizationAttributeValueChangeEvent) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OrganizationAttributeValueChangeEvent) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OrganizationAttributeValueChangeEvent) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OrganizationAttributeValueChangeEvent) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *OrganizationAttributeValueChangeEvent) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OrganizationAttributeValueChangeEvent) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OrganizationAttributeValueChangeEvent) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *OrganizationAttributeValueChangeEvent) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *OrganizationAttributeValueChangeEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationAttributeValueChangeEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationAttributeValueChangeEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationAttributeValueChangeEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCorrelationId

`func (o *OrganizationAttributeValueChangeEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *OrganizationAttributeValueChangeEvent) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *OrganizationAttributeValueChangeEvent) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *OrganizationAttributeValueChangeEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDomain

`func (o *OrganizationAttributeValueChangeEvent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationAttributeValueChangeEvent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationAttributeValueChangeEvent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OrganizationAttributeValueChangeEvent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTitle

`func (o *OrganizationAttributeValueChangeEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationAttributeValueChangeEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationAttributeValueChangeEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OrganizationAttributeValueChangeEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationAttributeValueChangeEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationAttributeValueChangeEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationAttributeValueChangeEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationAttributeValueChangeEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *OrganizationAttributeValueChangeEvent) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OrganizationAttributeValueChangeEvent) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OrganizationAttributeValueChangeEvent) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OrganizationAttributeValueChangeEvent) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeOccurred

`func (o *OrganizationAttributeValueChangeEvent) GetTimeOccurred() time.Time`

GetTimeOccurred returns the TimeOccurred field if non-nil, zero value otherwise.

### GetTimeOccurredOk

`func (o *OrganizationAttributeValueChangeEvent) GetTimeOccurredOk() (*time.Time, bool)`

GetTimeOccurredOk returns a tuple with the TimeOccurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOccurred

`func (o *OrganizationAttributeValueChangeEvent) SetTimeOccurred(v time.Time)`

SetTimeOccurred sets TimeOccurred field to given value.

### HasTimeOccurred

`func (o *OrganizationAttributeValueChangeEvent) HasTimeOccurred() bool`

HasTimeOccurred returns a boolean if a field has been set.

### GetSource

`func (o *OrganizationAttributeValueChangeEvent) GetSource() EntityRef`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OrganizationAttributeValueChangeEvent) GetSourceOk() (*EntityRef, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OrganizationAttributeValueChangeEvent) SetSource(v EntityRef)`

SetSource sets Source field to given value.

### HasSource

`func (o *OrganizationAttributeValueChangeEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReportingSystem

`func (o *OrganizationAttributeValueChangeEvent) GetReportingSystem() EntityRef`

GetReportingSystem returns the ReportingSystem field if non-nil, zero value otherwise.

### GetReportingSystemOk

`func (o *OrganizationAttributeValueChangeEvent) GetReportingSystemOk() (*EntityRef, bool)`

GetReportingSystemOk returns a tuple with the ReportingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingSystem

`func (o *OrganizationAttributeValueChangeEvent) SetReportingSystem(v EntityRef)`

SetReportingSystem sets ReportingSystem field to given value.

### HasReportingSystem

`func (o *OrganizationAttributeValueChangeEvent) HasReportingSystem() bool`

HasReportingSystem returns a boolean if a field has been set.

### GetRelatedParty

`func (o *OrganizationAttributeValueChangeEvent) GetRelatedParty() []RelatedPartyRefOrPartyRoleRef`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *OrganizationAttributeValueChangeEvent) GetRelatedPartyOk() (*[]RelatedPartyRefOrPartyRoleRef, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *OrganizationAttributeValueChangeEvent) SetRelatedParty(v []RelatedPartyRefOrPartyRoleRef)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *OrganizationAttributeValueChangeEvent) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetAnalyticCharacteristic

`func (o *OrganizationAttributeValueChangeEvent) GetAnalyticCharacteristic() []Characteristic`

GetAnalyticCharacteristic returns the AnalyticCharacteristic field if non-nil, zero value otherwise.

### GetAnalyticCharacteristicOk

`func (o *OrganizationAttributeValueChangeEvent) GetAnalyticCharacteristicOk() (*[]Characteristic, bool)`

GetAnalyticCharacteristicOk returns a tuple with the AnalyticCharacteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticCharacteristic

`func (o *OrganizationAttributeValueChangeEvent) SetAnalyticCharacteristic(v []Characteristic)`

SetAnalyticCharacteristic sets AnalyticCharacteristic field to given value.

### HasAnalyticCharacteristic

`func (o *OrganizationAttributeValueChangeEvent) HasAnalyticCharacteristic() bool`

HasAnalyticCharacteristic returns a boolean if a field has been set.

### GetEventId

`func (o *OrganizationAttributeValueChangeEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *OrganizationAttributeValueChangeEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *OrganizationAttributeValueChangeEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *OrganizationAttributeValueChangeEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *OrganizationAttributeValueChangeEvent) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *OrganizationAttributeValueChangeEvent) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *OrganizationAttributeValueChangeEvent) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *OrganizationAttributeValueChangeEvent) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *OrganizationAttributeValueChangeEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OrganizationAttributeValueChangeEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OrganizationAttributeValueChangeEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OrganizationAttributeValueChangeEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEvent

`func (o *OrganizationAttributeValueChangeEvent) GetEvent() OrganizationAttributeValueChangeEventPayload`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OrganizationAttributeValueChangeEvent) GetEventOk() (*OrganizationAttributeValueChangeEventPayload, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OrganizationAttributeValueChangeEvent) SetEvent(v OrganizationAttributeValueChangeEventPayload)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OrganizationAttributeValueChangeEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


