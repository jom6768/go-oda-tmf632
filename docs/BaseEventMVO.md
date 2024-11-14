# BaseEventMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
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

### NewBaseEventMVO

`func NewBaseEventMVO(type_ string, ) *BaseEventMVO`

NewBaseEventMVO instantiates a new BaseEventMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventMVOWithDefaults

`func NewBaseEventMVOWithDefaults() *BaseEventMVO`

NewBaseEventMVOWithDefaults instantiates a new BaseEventMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseEventMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEventMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEventMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *BaseEventMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *BaseEventMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *BaseEventMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *BaseEventMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *BaseEventMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *BaseEventMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *BaseEventMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *BaseEventMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetEvent

`func (o *BaseEventMVO) GetEvent() map[string]interface{}`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *BaseEventMVO) GetEventOk() (*map[string]interface{}, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *BaseEventMVO) SetEvent(v map[string]interface{})`

SetEvent sets Event field to given value.

### HasEvent

`func (o *BaseEventMVO) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventId

`func (o *BaseEventMVO) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BaseEventMVO) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BaseEventMVO) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *BaseEventMVO) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *BaseEventMVO) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *BaseEventMVO) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *BaseEventMVO) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *BaseEventMVO) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventType

`func (o *BaseEventMVO) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseEventMVO) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseEventMVO) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseEventMVO) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetCorrelationId

`func (o *BaseEventMVO) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BaseEventMVO) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BaseEventMVO) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *BaseEventMVO) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetDomain

`func (o *BaseEventMVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BaseEventMVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BaseEventMVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BaseEventMVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTitle

`func (o *BaseEventMVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BaseEventMVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BaseEventMVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BaseEventMVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *BaseEventMVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseEventMVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseEventMVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseEventMVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *BaseEventMVO) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BaseEventMVO) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BaseEventMVO) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BaseEventMVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTimeOcurred

`func (o *BaseEventMVO) GetTimeOcurred() time.Time`

GetTimeOcurred returns the TimeOcurred field if non-nil, zero value otherwise.

### GetTimeOcurredOk

`func (o *BaseEventMVO) GetTimeOcurredOk() (*time.Time, bool)`

GetTimeOcurredOk returns a tuple with the TimeOcurred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOcurred

`func (o *BaseEventMVO) SetTimeOcurred(v time.Time)`

SetTimeOcurred sets TimeOcurred field to given value.

### HasTimeOcurred

`func (o *BaseEventMVO) HasTimeOcurred() bool`

HasTimeOcurred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


