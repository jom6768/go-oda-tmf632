# OrganizationParentRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**RelationshipType** | Pointer to **string** | Type of the relationship. Could be juridical, hierarchical, geographical, functional for example. | [optional] 
**Organization** | Pointer to [**OrganizationRef**](OrganizationRef.md) |  | [optional] 

## Methods

### NewOrganizationParentRelationship

`func NewOrganizationParentRelationship(type_ string, ) *OrganizationParentRelationship`

NewOrganizationParentRelationship instantiates a new OrganizationParentRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationParentRelationshipWithDefaults

`func NewOrganizationParentRelationshipWithDefaults() *OrganizationParentRelationship`

NewOrganizationParentRelationshipWithDefaults instantiates a new OrganizationParentRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationParentRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationParentRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationParentRelationship) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OrganizationParentRelationship) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OrganizationParentRelationship) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OrganizationParentRelationship) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OrganizationParentRelationship) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OrganizationParentRelationship) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OrganizationParentRelationship) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OrganizationParentRelationship) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OrganizationParentRelationship) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetRelationshipType

`func (o *OrganizationParentRelationship) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OrganizationParentRelationship) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OrganizationParentRelationship) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *OrganizationParentRelationship) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationParentRelationship) GetOrganization() OrganizationRef`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationParentRelationship) GetOrganizationOk() (*OrganizationRef, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationParentRelationship) SetOrganization(v OrganizationRef)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationParentRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


