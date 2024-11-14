# TaxDefinitionMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | Unique identifier of the tax. | [optional] 
**Name** | Pointer to **string** | Tax name. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**JurisdictionName** | Pointer to **string** | Name of the jurisdiction that levies the tax | [optional] 
**JurisdictionLevel** | Pointer to **string** | Level of the jurisdiction that levies the tax | [optional] 
**TaxType** | Pointer to **string** | Type of the tax. | [optional] 

## Methods

### NewTaxDefinitionMVO

`func NewTaxDefinitionMVO(type_ string, ) *TaxDefinitionMVO`

NewTaxDefinitionMVO instantiates a new TaxDefinitionMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDefinitionMVOWithDefaults

`func NewTaxDefinitionMVOWithDefaults() *TaxDefinitionMVO`

NewTaxDefinitionMVOWithDefaults instantiates a new TaxDefinitionMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxDefinitionMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxDefinitionMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxDefinitionMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *TaxDefinitionMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TaxDefinitionMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TaxDefinitionMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TaxDefinitionMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *TaxDefinitionMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *TaxDefinitionMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *TaxDefinitionMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *TaxDefinitionMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *TaxDefinitionMVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxDefinitionMVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxDefinitionMVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxDefinitionMVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxDefinitionMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxDefinitionMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxDefinitionMVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxDefinitionMVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidFor

`func (o *TaxDefinitionMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *TaxDefinitionMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *TaxDefinitionMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *TaxDefinitionMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetJurisdictionName

`func (o *TaxDefinitionMVO) GetJurisdictionName() string`

GetJurisdictionName returns the JurisdictionName field if non-nil, zero value otherwise.

### GetJurisdictionNameOk

`func (o *TaxDefinitionMVO) GetJurisdictionNameOk() (*string, bool)`

GetJurisdictionNameOk returns a tuple with the JurisdictionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionName

`func (o *TaxDefinitionMVO) SetJurisdictionName(v string)`

SetJurisdictionName sets JurisdictionName field to given value.

### HasJurisdictionName

`func (o *TaxDefinitionMVO) HasJurisdictionName() bool`

HasJurisdictionName returns a boolean if a field has been set.

### GetJurisdictionLevel

`func (o *TaxDefinitionMVO) GetJurisdictionLevel() string`

GetJurisdictionLevel returns the JurisdictionLevel field if non-nil, zero value otherwise.

### GetJurisdictionLevelOk

`func (o *TaxDefinitionMVO) GetJurisdictionLevelOk() (*string, bool)`

GetJurisdictionLevelOk returns a tuple with the JurisdictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionLevel

`func (o *TaxDefinitionMVO) SetJurisdictionLevel(v string)`

SetJurisdictionLevel sets JurisdictionLevel field to given value.

### HasJurisdictionLevel

`func (o *TaxDefinitionMVO) HasJurisdictionLevel() bool`

HasJurisdictionLevel returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxDefinitionMVO) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxDefinitionMVO) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxDefinitionMVO) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxDefinitionMVO) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


