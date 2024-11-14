# BaseEventFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**Event** | Pointer to **map[string]interface{}** |  | [optional] 
**EventId** | Pointer to **string** | The identifier of the notification. | [optional] 
**EventTime** | Pointer to **time.Time** | Time of the event occurrence. | [optional] 
**EventType** | Pointer to **string** | The type of the notification. | [optional] 
**CorrelationId** | Pointer to **string** | The correlation id for this event. | [optional] 
**Domain** | Pointer to **string** | The domain of the event. | [optional] 
**Title** | Pointer to **string** | The title of the event. | [optional] 
**Description** | Pointer to **string** | An explanatory of the event. | [optional] 
**Priority** | Pointer to **string** | A priority. | [optional] 
**TimeOcurred** | Pointer to **time.Time** | The time the event occured. | [optional] 

## Methods

### NewBaseEventFVO

`func NewBaseEventFVO(type_ string, ) *BaseEventFVO`

NewBaseEventFVO instantiates a new BaseEventFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventFVOWithDefaults

`func NewBaseEventFVOWithDefaults() *BaseEventFVO`

NewBaseEventFVOWithDefaults instantiates a new BaseEventFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseEventFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEventFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEventFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *BaseEventFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *BaseEventFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *BaseEventFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *BaseEventFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *BaseEventFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *BaseEventFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *BaseEventFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *BaseEventFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *BaseEventFVO) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BaseEventFVO) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BaseEventFVO) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BaseEventFVO) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BaseEventFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEventFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEventFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEventFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *BaseEventFVO) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BaseEventFVO) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BaseEventFVO) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BaseEventFVO) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventId

`func (o *BaseEventFVO) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BaseEventFVO) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BaseEventFVO) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *BaseEventFVO) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *BaseEventFVO) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *BaseEventFVO) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *BaseEventFVO) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *BaseEventFVO) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *BaseEventFVO) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEventFVO) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEventFVO) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseEventFVO) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetCorrelationId

`func (o *BaseEventFVO) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BaseEventFVO) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BaseEventFVO) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *BaseEventFVO) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDomain

`func (o *BaseEventFVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BaseEventFVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BaseEventFVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BaseEventFVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTitle

`func (o *BaseEventFVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BaseEventFVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BaseEventFVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BaseEventFVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *BaseEventFVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseEventFVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseEventFVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseEventFVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *BaseEventFVO) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BaseEventFVO) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BaseEventFVO) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BaseEventFVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeOcurred

`func (o *BaseEventFVO) GetTimeOcurred() time.Time`

GetTimeOcurred returns the TimeOcurred field if non-nil, zero value otherwise.

### GetTimeOcurredOk

`func (o *BaseEventFVO) GetTimeOcurredOk() (*time.Time, bool)`

GetTimeOcurredOk returns a tuple with the TimeOcurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOcurred

`func (o *BaseEventFVO) SetTimeOcurred(v time.Time)`

SetTimeOcurred sets TimeOcurred field to given value.

### HasTimeOcurred

`func (o *BaseEventFVO) HasTimeOcurred() bool`

HasTimeOcurred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


