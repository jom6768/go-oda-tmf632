# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLegalEntity** | Pointer to **bool** | If value is true, the organization is a legal entity known by a national referential. | [optional] 
**IsHeadOffice** | Pointer to **bool** | If value is true, the organization is the head office | [optional] 
**OrganizationType** | Pointer to **string** | Type of Organization (company, department...) | [optional] 
**ExistsDuring** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Name** | Pointer to **string** | Organization name (department name for example) | [optional] 
**NameType** | Pointer to **string** | Type of the name : Co, Inc, Ltd, etc. | [optional] 
**Status** | Pointer to [**OrganizationStateType**](OrganizationStateType.md) |  | [optional] 
**OtherName** | Pointer to [**[]OtherNameOrganization**](OtherNameOrganization.md) | List of additional names by which the organization is known | [optional] 
**OrganizationIdentification** | Pointer to [**[]OrganizationIdentification**](OrganizationIdentification.md) | List of official identifiers given to the organization, for example company number in the registry of companies | [optional] 
**OrganizationChildRelationship** | Pointer to [**[]OrganizationChildRelationship**](OrganizationChildRelationship.md) | List of organizations that are contained within this organization. For example if this organization is the Legal Department, the child organizations might include Claims, Courts, Contracts | [optional] 
**OrganizationParentRelationship** | Pointer to [**OrganizationParentRelationship**](OrganizationParentRelationship.md) |  | [optional] 
**TradingName** | Pointer to **string** | Name that the organization (unit) trades under | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLegalEntity

`func (o *Organization) GetIsLegalEntity() bool`

GetIsLegalEntity returns the IsLegalEntity field if non-nil, zero value otherwise.

### GetIsLegalEntityOk

`func (o *Organization) GetIsLegalEntityOk() (*bool, bool)`

GetIsLegalEntityOk returns a tuple with the IsLegalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegalEntity

`func (o *Organization) SetIsLegalEntity(v bool)`

SetIsLegalEntity sets IsLegalEntity field to given value.

### HasIsLegalEntity

`func (o *Organization) HasIsLegalEntity() bool`

HasIsLegalEntity returns a boolean if a field has been set.

### GetIsHeadOffice

`func (o *Organization) GetIsHeadOffice() bool`

GetIsHeadOffice returns the IsHeadOffice field if non-nil, zero value otherwise.

### GetIsHeadOfficeOk

`func (o *Organization) GetIsHeadOfficeOk() (*bool, bool)`

GetIsHeadOfficeOk returns a tuple with the IsHeadOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadOffice

`func (o *Organization) SetIsHeadOffice(v bool)`

SetIsHeadOffice sets IsHeadOffice field to given value.

### HasIsHeadOffice

`func (o *Organization) HasIsHeadOffice() bool`

HasIsHeadOffice returns a boolean if a field has been set.

### GetOrganizationType

`func (o *Organization) GetOrganizationType() string`

GetOrganizationType returns the OrganizationType field if non-nil, zero value otherwise.

### GetOrganizationTypeOk

`func (o *Organization) GetOrganizationTypeOk() (*string, bool)`

GetOrganizationTypeOk returns a tuple with the OrganizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationType

`func (o *Organization) SetOrganizationType(v string)`

SetOrganizationType sets OrganizationType field to given value.

### HasOrganizationType

`func (o *Organization) HasOrganizationType() bool`

HasOrganizationType returns a boolean if a field has been set.

### GetExistsDuring

`func (o *Organization) GetExistsDuring() TimePeriod`

GetExistsDuring returns the ExistsDuring field if non-nil, zero value otherwise.

### GetExistsDuringOk

`func (o *Organization) GetExistsDuringOk() (*TimePeriod, bool)`

GetExistsDuringOk returns a tuple with the ExistsDuring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistsDuring

`func (o *Organization) SetExistsDuring(v TimePeriod)`

SetExistsDuring sets ExistsDuring field to given value.

### HasExistsDuring

`func (o *Organization) HasExistsDuring() bool`

HasExistsDuring returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameType

`func (o *Organization) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *Organization) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *Organization) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *Organization) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### GetStatus

`func (o *Organization) GetStatus() OrganizationStateType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Organization) GetStatusOk() (*OrganizationStateType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Organization) SetStatus(v OrganizationStateType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Organization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOtherName

`func (o *Organization) GetOtherName() []OtherNameOrganization`

GetOtherName returns the OtherName field if non-nil, zero value otherwise.

### GetOtherNameOk

`func (o *Organization) GetOtherNameOk() (*[]OtherNameOrganization, bool)`

GetOtherNameOk returns a tuple with the OtherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherName

`func (o *Organization) SetOtherName(v []OtherNameOrganization)`

SetOtherName sets OtherName field to given value.

### HasOtherName

`func (o *Organization) HasOtherName() bool`

HasOtherName returns a boolean if a field has been set.

### GetOrganizationIdentification

`func (o *Organization) GetOrganizationIdentification() []OrganizationIdentification`

GetOrganizationIdentification returns the OrganizationIdentification field if non-nil, zero value otherwise.

### GetOrganizationIdentificationOk

`func (o *Organization) GetOrganizationIdentificationOk() (*[]OrganizationIdentification, bool)`

GetOrganizationIdentificationOk returns a tuple with the OrganizationIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIdentification

`func (o *Organization) SetOrganizationIdentification(v []OrganizationIdentification)`

SetOrganizationIdentification sets OrganizationIdentification field to given value.

### HasOrganizationIdentification

`func (o *Organization) HasOrganizationIdentification() bool`

HasOrganizationIdentification returns a boolean if a field has been set.

### GetOrganizationChildRelationship

`func (o *Organization) GetOrganizationChildRelationship() []OrganizationChildRelationship`

GetOrganizationChildRelationship returns the OrganizationChildRelationship field if non-nil, zero value otherwise.

### GetOrganizationChildRelationshipOk

`func (o *Organization) GetOrganizationChildRelationshipOk() (*[]OrganizationChildRelationship, bool)`

GetOrganizationChildRelationshipOk returns a tuple with the OrganizationChildRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationChildRelationship

`func (o *Organization) SetOrganizationChildRelationship(v []OrganizationChildRelationship)`

SetOrganizationChildRelationship sets OrganizationChildRelationship field to given value.

### HasOrganizationChildRelationship

`func (o *Organization) HasOrganizationChildRelationship() bool`

HasOrganizationChildRelationship returns a boolean if a field has been set.

### GetOrganizationParentRelationship

`func (o *Organization) GetOrganizationParentRelationship() OrganizationParentRelationship`

GetOrganizationParentRelationship returns the OrganizationParentRelationship field if non-nil, zero value otherwise.

### GetOrganizationParentRelationshipOk

`func (o *Organization) GetOrganizationParentRelationshipOk() (*OrganizationParentRelationship, bool)`

GetOrganizationParentRelationshipOk returns a tuple with the OrganizationParentRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationParentRelationship

`func (o *Organization) SetOrganizationParentRelationship(v OrganizationParentRelationship)`

SetOrganizationParentRelationship sets OrganizationParentRelationship field to given value.

### HasOrganizationParentRelationship

`func (o *Organization) HasOrganizationParentRelationship() bool`

HasOrganizationParentRelationship returns a boolean if a field has been set.

### GetTradingName

`func (o *Organization) GetTradingName() string`

GetTradingName returns the TradingName field if non-nil, zero value otherwise.

### GetTradingNameOk

`func (o *Organization) GetTradingNameOk() (*string, bool)`

GetTradingNameOk returns a tuple with the TradingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingName

`func (o *Organization) SetTradingName(v string)`

SetTradingName sets TradingName field to given value.

### HasTradingName

`func (o *Organization) HasTradingName() bool`

HasTradingName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


