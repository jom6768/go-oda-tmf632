# Quantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Numeric value in a given unit | [optional] [default to 1]
**Units** | Pointer to **string** | Unit | [optional] 

## Methods

### NewQuantity

`func NewQuantity() *Quantity`

NewQuantity instantiates a new Quantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityWithDefaults

`func NewQuantityWithDefaults() *Quantity`

NewQuantityWithDefaults instantiates a new Quantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Quantity) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Quantity) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Quantity) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Quantity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUnits

`func (o *Quantity) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Quantity) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Quantity) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Quantity) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


