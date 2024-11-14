# OrganizationParentRelationshipMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**RelationshipType** | Pointer to **string** | Type of the relationship. Could be juridical, hierarchical, geographical, functional for example. | [optional] 
**Organization** | Pointer to [**OrganizationRefMVO**](OrganizationRefMVO.md) |  | [optional] 

## Methods

### NewOrganizationParentRelationshipMVO

`func NewOrganizationParentRelationshipMVO(type_ string, ) *OrganizationParentRelationshipMVO`

NewOrganizationParentRelationshipMVO instantiates a new OrganizationParentRelationshipMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationParentRelationshipMVOWithDefaults

`func NewOrganizationParentRelationshipMVOWithDefaults() *OrganizationParentRelationshipMVO`

NewOrganizationParentRelationshipMVOWithDefaults instantiates a new OrganizationParentRelationshipMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrganizationParentRelationshipMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationParentRelationshipMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationParentRelationshipMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OrganizationParentRelationshipMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OrganizationParentRelationshipMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OrganizationParentRelationshipMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OrganizationParentRelationshipMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OrganizationParentRelationshipMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OrganizationParentRelationshipMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OrganizationParentRelationshipMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OrganizationParentRelationshipMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetRelationshipType

`func (o *OrganizationParentRelationshipMVO) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *OrganizationParentRelationshipMVO) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *OrganizationParentRelationshipMVO) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *OrganizationParentRelationshipMVO) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationParentRelationshipMVO) GetOrganization() OrganizationRefMVO`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationParentRelationshipMVO) GetOrganizationOk() (*OrganizationRefMVO, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationParentRelationshipMVO) SetOrganization(v OrganizationRefMVO)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationParentRelationshipMVO) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


