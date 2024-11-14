# RelatedPartyOrPartyRoleMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Role** | **string** | Role played by the related party or party role in the context of the specific entity it is linked to. Such as &#39;initiator&#39;, &#39;customer&#39;,  &#39;salesAgent&#39;, &#39;user&#39; | 
**PartyOrPartyRole** | Pointer to [**PartyOrPartyRoleMVO**](PartyOrPartyRoleMVO.md) |  | [optional] 

## Methods

### NewRelatedPartyOrPartyRoleMVO

`func NewRelatedPartyOrPartyRoleMVO(type_ string, role string, ) *RelatedPartyOrPartyRoleMVO`

NewRelatedPartyOrPartyRoleMVO instantiates a new RelatedPartyOrPartyRoleMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedPartyOrPartyRoleMVOWithDefaults

`func NewRelatedPartyOrPartyRoleMVOWithDefaults() *RelatedPartyOrPartyRoleMVO`

NewRelatedPartyOrPartyRoleMVOWithDefaults instantiates a new RelatedPartyOrPartyRoleMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelatedPartyOrPartyRoleMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelatedPartyOrPartyRoleMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelatedPartyOrPartyRoleMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *RelatedPartyOrPartyRoleMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *RelatedPartyOrPartyRoleMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *RelatedPartyOrPartyRoleMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *RelatedPartyOrPartyRoleMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *RelatedPartyOrPartyRoleMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *RelatedPartyOrPartyRoleMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *RelatedPartyOrPartyRoleMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *RelatedPartyOrPartyRoleMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetRole

`func (o *RelatedPartyOrPartyRoleMVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RelatedPartyOrPartyRoleMVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RelatedPartyOrPartyRoleMVO) SetRole(v string)`

SetRole sets Role field to given value.


### GetPartyOrPartyRole

`func (o *RelatedPartyOrPartyRoleMVO) GetPartyOrPartyRole() PartyOrPartyRoleMVO`

GetPartyOrPartyRole returns the PartyOrPartyRole field if non-nil, zero value otherwise.

### GetPartyOrPartyRoleOk

`func (o *RelatedPartyOrPartyRoleMVO) GetPartyOrPartyRoleOk() (*PartyOrPartyRoleMVO, bool)`

GetPartyOrPartyRoleOk returns a tuple with the PartyOrPartyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyOrPartyRole

`func (o *RelatedPartyOrPartyRoleMVO) SetPartyOrPartyRole(v PartyOrPartyRoleMVO)`

SetPartyOrPartyRole sets PartyOrPartyRole field to given value.

### HasPartyOrPartyRole

`func (o *RelatedPartyOrPartyRoleMVO) HasPartyOrPartyRole() bool`

HasPartyOrPartyRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


