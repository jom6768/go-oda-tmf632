# OtherNameOrganization

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

### NewOtherNameOrganization

`func NewOtherNameOrganization(type_ string, ) *OtherNameOrganization`

NewOtherNameOrganization instantiates a new OtherNameOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherNameOrganizationWithDefaults

`func NewOtherNameOrganizationWithDefaults() *OtherNameOrganization`

NewOtherNameOrganizationWithDefaults instantiates a new OtherNameOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OtherNameOrganization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OtherNameOrganization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OtherNameOrganization) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *OtherNameOrganization) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *OtherNameOrganization) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *OtherNameOrganization) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *OtherNameOrganization) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *OtherNameOrganization) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *OtherNameOrganization) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *OtherNameOrganization) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *OtherNameOrganization) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetTradingName

`func (o *OtherNameOrganization) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *OtherNameOrganization) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *OtherNameOrganization) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *OtherNameOrganization) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.

### GetNameType

`func (o *OtherNameOrganization) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *OtherNameOrganization) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *OtherNameOrganization) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *OtherNameOrganization) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetName

`func (o *OtherNameOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtherNameOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtherNameOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OtherNameOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidFor

`func (o *OtherNameOrganization) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *OtherNameOrganization) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *OtherNameOrganization) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *OtherNameOrganization) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


