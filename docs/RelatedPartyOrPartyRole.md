# RelatedPartyOrPartyRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Role** | Pointer to **string** | Role played by the related party or party role in the context of the specific entity it is linked to. Such as &#39;initiator&#39;, &#39;customer&#39;,  &#39;salesAgent&#39;, &#39;user&#39; | [optional] 
**PartyOrPartyRole** | Pointer to [**PartyOrPartyRole**](PartyOrPartyRole.md) |  | [optional] 

## Methods

### NewRelatedPartyOrPartyRole

`func NewRelatedPartyOrPartyRole(type_ string, ) *RelatedPartyOrPartyRole`

NewRelatedPartyOrPartyRole instantiates a new RelatedPartyOrPartyRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedPartyOrPartyRoleWithDefaults

`func NewRelatedPartyOrPartyRoleWithDefaults() *RelatedPartyOrPartyRole`

NewRelatedPartyOrPartyRoleWithDefaults instantiates a new RelatedPartyOrPartyRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelatedPartyOrPartyRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelatedPartyOrPartyRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelatedPartyOrPartyRole) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *RelatedPartyOrPartyRole) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *RelatedPartyOrPartyRole) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *RelatedPartyOrPartyRole) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *RelatedPartyOrPartyRole) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *RelatedPartyOrPartyRole) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *RelatedPartyOrPartyRole) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *RelatedPartyOrPartyRole) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *RelatedPartyOrPartyRole) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetRole

`func (o *RelatedPartyOrPartyRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RelatedPartyOrPartyRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RelatedPartyOrPartyRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *RelatedPartyOrPartyRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetPartyOrPartyRole

`func (o *RelatedPartyOrPartyRole) GetPartyOrPartyRole() PartyOrPartyRole`

GetPartyOrPartyRole returns the PartyOrPartyRole field if non-nil, zero value otherwise.

### GetPartyOrPartyRoleOk

`func (o *RelatedPartyOrPartyRole) GetPartyOrPartyRoleOk() (*PartyOrPartyRole, bool)`

GetPartyOrPartyRoleOk returns a tuple with the PartyOrPartyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyOrPartyRole

`func (o *RelatedPartyOrPartyRole) SetPartyOrPartyRole(v PartyOrPartyRole)`

SetPartyOrPartyRole sets PartyOrPartyRole field to given value.

### HasPartyOrPartyRole

`func (o *RelatedPartyOrPartyRole) HasPartyOrPartyRole() bool`

HasPartyOrPartyRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


