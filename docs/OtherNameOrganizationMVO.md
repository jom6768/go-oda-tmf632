# OtherNameOrganizationMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**TradingName** | Pointer to **string** | The name that the organization trades under | [optional] 
**NameType** | Pointer to **string** | Co. , Inc. , Ltd. , Pty Ltd. , Plc; , Gmbh | [optional] 
**Name** | Pointer to **string** | Organization name (department name for example) | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewOtherNameOrganizationMVO

`func NewOtherNameOrganizationMVO(type_ string, ) *OtherNameOrganizationMVO`

NewOtherNameOrganizationMVO instantiates a new OtherNameOrganizationMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherNameOrganizationMVOWithDefaults

`func NewOtherNameOrganizationMVOWithDefaults() *OtherNameOrganizationMVO`

NewOtherNameOrganizationMVOWithDefaults instantiates a new OtherNameOrganizationMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OtherNameOrganizationMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OtherNameOrganizationMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OtherNameOrganizationMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OtherNameOrganizationMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OtherNameOrganizationMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OtherNameOrganizationMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OtherNameOrganizationMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OtherNameOrganizationMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OtherNameOrganizationMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OtherNameOrganizationMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OtherNameOrganizationMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetTradingName

`func (o *OtherNameOrganizationMVO) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *OtherNameOrganizationMVO) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *OtherNameOrganizationMVO) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *OtherNameOrganizationMVO) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetNameType

`func (o *OtherNameOrganizationMVO) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *OtherNameOrganizationMVO) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *OtherNameOrganizationMVO) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *OtherNameOrganizationMVO) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetName

`func (o *OtherNameOrganizationMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtherNameOrganizationMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtherNameOrganizationMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OtherNameOrganizationMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidFor

`func (o *OtherNameOrganizationMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *OtherNameOrganizationMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *OtherNameOrganizationMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *OtherNameOrganizationMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


