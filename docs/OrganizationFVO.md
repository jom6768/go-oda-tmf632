# OrganizationFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLegalEntity** | Pointer to **bool** | If value is true, the organization is a legal entity known by a national referential. | [optional] 
**IsHeadOffice** | Pointer to **bool** | If value is true, the organization is the head office | [optional] 
**OrganizationType** | Pointer to **string** | Type of Organization (company, department...) | [optional] 
**ExistsDuring** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Name** | **string** | Organization name (department name for example) | 
**NameType** | Pointer to **string** | Type of the name : Co, Inc, Ltd, etc. | [optional] 
**Status** | Pointer to [**OrganizationStateType**](OrganizationStateType.md) |  | [optional] 
**OtherName** | Pointer to [**[]OtherNameOrganizationFVO**](OtherNameOrganizationFVO.md) | List of additional names by which the organization is known | [optional] 
**OrganizationIdentification** | Pointer to [**[]OrganizationIdentificationFVO**](OrganizationIdentificationFVO.md) | List of official identifiers given to the organization, for example company number in the registry of companies | [optional] 
**OrganizationChildRelationship** | Pointer to [**[]OrganizationChildRelationshipFVO**](OrganizationChildRelationshipFVO.md) | List of organizations that are contained within this organization. For example if this organization is the Legal Department, the child organizations might include Claims, Courts, Contracts | [optional] 
**OrganizationParentRelationship** | Pointer to [**OrganizationParentRelationshipFVO**](OrganizationParentRelationshipFVO.md) |  | [optional] 
**TradingName** | Pointer to **string** | Name that the organization (unit) trades under | [optional] 

## Methods

### NewOrganizationFVO

`func NewOrganizationFVO(name string, ) *OrganizationFVO`

NewOrganizationFVO instantiates a new OrganizationFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationFVOWithDefaults

`func NewOrganizationFVOWithDefaults() *OrganizationFVO`

NewOrganizationFVOWithDefaults instantiates a new OrganizationFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLegalEntity

`func (o *OrganizationFVO) GetIsLegalEntity() bool`

GetIsLegalEntity returns the IsLegalEntity field if non-nil, zero value otherwise.

### GetIsLegalEntityOk

`func (o *OrganizationFVO) GetIsLegalEntityOk() (*bool, bool)`

GetIsLegalEntityOk returns a tuple with the IsLegalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalEntity

`func (o *OrganizationFVO) SetIsLegalEntity(v bool)`

SetIsLegalEntity sets IsLegalEntity field to given value.

### HasIsLegalEntity

`func (o *OrganizationFVO) HasIsLegalEntity() bool`

HasIsLegalEntity returns a boolean if a field has been set.

### GetIsHeadOffice

`func (o *OrganizationFVO) GetIsHeadOffice() bool`

GetIsHeadOffice returns the IsHeadOffice field if non-nil, zero value otherwise.

### GetIsHeadOfficeOk

`func (o *OrganizationFVO) GetIsHeadOfficeOk() (*bool, bool)`

GetIsHeadOfficeOk returns a tuple with the IsHeadOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadOffice

`func (o *OrganizationFVO) SetIsHeadOffice(v bool)`

SetIsHeadOffice sets IsHeadOffice field to given value.

### HasIsHeadOffice

`func (o *OrganizationFVO) HasIsHeadOffice() bool`

HasIsHeadOffice returns a boolean if a field has been set.

### GetOrganizationType

`func (o *OrganizationFVO) GetOrganizationType() string`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *OrganizationFVO) GetOrganizationTypeOk() (*string, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *OrganizationFVO) SetOrganizationType(v string)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *OrganizationFVO) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetExistsDuring

`func (o *OrganizationFVO) GetExistsDuring() TimePeriod`

GetExistsDuring returns the ExistsDuring field if non-nil, zero value otherwise.

### GetExistsDuringOk

`func (o *OrganizationFVO) GetExistsDuringOk() (*TimePeriod, bool)`

GetExistsDuringOk returns a tuple with the ExistsDuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsDuring

`func (o *OrganizationFVO) SetExistsDuring(v TimePeriod)`

SetExistsDuring sets ExistsDuring field to given value.

### HasExistsDuring

`func (o *OrganizationFVO) HasExistsDuring() bool`

HasExistsDuring returns a boolean if a field has been set.

### GetName

`func (o *OrganizationFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationFVO) SetName(v string)`

SetName sets Name field to given value.


### GetNameType

`func (o *OrganizationFVO) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *OrganizationFVO) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *OrganizationFVO) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *OrganizationFVO) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationFVO) GetStatus() OrganizationStateType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationFVO) GetStatusOk() (*OrganizationStateType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationFVO) SetStatus(v OrganizationStateType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationFVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *OrganizationFVO) GetOtherName() []OtherNameOrganizationFVO`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *OrganizationFVO) GetOtherNameOk() (*[]OtherNameOrganizationFVO, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *OrganizationFVO) SetOtherName(v []OtherNameOrganizationFVO)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *OrganizationFVO) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetOrganizationIdentification

`func (o *OrganizationFVO) GetOrganizationIdentification() []OrganizationIdentificationFVO`

GetOrganizationIdentification returns the OrganizationIdentification field if non-nil, zero value otherwise.

### GetOrganizationIdentificationOk

`func (o *OrganizationFVO) GetOrganizationIdentificationOk() (*[]OrganizationIdentificationFVO, bool)`

GetOrganizationIdentificationOk returns a tuple with the OrganizationIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentification

`func (o *OrganizationFVO) SetOrganizationIdentification(v []OrganizationIdentificationFVO)`

SetOrganizationIdentification sets OrganizationIdentification field to given value.

### HasOrganizationIdentification

`func (o *OrganizationFVO) HasOrganizationIdentification() bool`

HasOrganizationIdentification returns a boolean if a field has been set.

### GetOrganizationChildRelationship

`func (o *OrganizationFVO) GetOrganizationChildRelationship() []OrganizationChildRelationshipFVO`

GetOrganizationChildRelationship returns the OrganizationChildRelationship field if non-nil, zero value otherwise.

### GetOrganizationChildRelationshipOk

`func (o *OrganizationFVO) GetOrganizationChildRelationshipOk() (*[]OrganizationChildRelationshipFVO, bool)`

GetOrganizationChildRelationshipOk returns a tuple with the OrganizationChildRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationChildRelationship

`func (o *OrganizationFVO) SetOrganizationChildRelationship(v []OrganizationChildRelationshipFVO)`

SetOrganizationChildRelationship sets OrganizationChildRelationship field to given value.

### HasOrganizationChildRelationship

`func (o *OrganizationFVO) HasOrganizationChildRelationship() bool`

HasOrganizationChildRelationship returns a boolean if a field has been set.

### GetOrganizationParentRelationship

`func (o *OrganizationFVO) GetOrganizationParentRelationship() OrganizationParentRelationshipFVO`

GetOrganizationParentRelationship returns the OrganizationParentRelationship field if non-nil, zero value otherwise.

### GetOrganizationParentRelationshipOk

`func (o *OrganizationFVO) GetOrganizationParentRelationshipOk() (*OrganizationParentRelationshipFVO, bool)`

GetOrganizationParentRelationshipOk returns a tuple with the OrganizationParentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationParentRelationship

`func (o *OrganizationFVO) SetOrganizationParentRelationship(v OrganizationParentRelationshipFVO)`

SetOrganizationParentRelationship sets OrganizationParentRelationship field to given value.

### HasOrganizationParentRelationship

`func (o *OrganizationFVO) HasOrganizationParentRelationship() bool`

HasOrganizationParentRelationship returns a boolean if a field has been set.

### GetTradingName

`func (o *OrganizationFVO) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *OrganizationFVO) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *OrganizationFVO) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *OrganizationFVO) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


