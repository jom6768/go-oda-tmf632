# TaxDefinitionFVO

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

### NewTaxDefinitionFVO

`func NewTaxDefinitionFVO(type_ string, ) *TaxDefinitionFVO`

NewTaxDefinitionFVO instantiates a new TaxDefinitionFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDefinitionFVOWithDefaults

`func NewTaxDefinitionFVOWithDefaults() *TaxDefinitionFVO`

NewTaxDefinitionFVOWithDefaults instantiates a new TaxDefinitionFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxDefinitionFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxDefinitionFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxDefinitionFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *TaxDefinitionFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *TaxDefinitionFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *TaxDefinitionFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *TaxDefinitionFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *TaxDefinitionFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *TaxDefinitionFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *TaxDefinitionFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *TaxDefinitionFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *TaxDefinitionFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxDefinitionFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxDefinitionFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxDefinitionFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxDefinitionFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxDefinitionFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxDefinitionFVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxDefinitionFVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidFor

`func (o *TaxDefinitionFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *TaxDefinitionFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *TaxDefinitionFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *TaxDefinitionFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.

### GetJurisdictionName

`func (o *TaxDefinitionFVO) GetJurisdictionName() string`

GetJurisdictionName returns the JurisdictionName field if non-nil, zero value otherwise.

### GetJurisdictionNameOk

`func (o *TaxDefinitionFVO) GetJurisdictionNameOk() (*string, bool)`

GetJurisdictionNameOk returns a tuple with the JurisdictionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionName

`func (o *TaxDefinitionFVO) SetJurisdictionName(v string)`

SetJurisdictionName sets JurisdictionName field to given value.

### HasJurisdictionName

`func (o *TaxDefinitionFVO) HasJurisdictionName() bool`

HasJurisdictionName returns a boolean if a field has been set.

### GetJurisdictionLevel

`func (o *TaxDefinitionFVO) GetJurisdictionLevel() string`

GetJurisdictionLevel returns the JurisdictionLevel field if non-nil, zero value otherwise.

### GetJurisdictionLevelOk

`func (o *TaxDefinitionFVO) GetJurisdictionLevelOk() (*string, bool)`

GetJurisdictionLevelOk returns a tuple with the JurisdictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdictionLevel

`func (o *TaxDefinitionFVO) SetJurisdictionLevel(v string)`

SetJurisdictionLevel sets JurisdictionLevel field to given value.

### HasJurisdictionLevel

`func (o *TaxDefinitionFVO) HasJurisdictionLevel() bool`

HasJurisdictionLevel returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxDefinitionFVO) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxDefinitionFVO) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxDefinitionFVO) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxDefinitionFVO) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


