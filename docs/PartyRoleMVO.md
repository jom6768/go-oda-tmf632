# PartyRoleMVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Name** | **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | [**PartyRefMVO**](PartyRefMVO.md) |  | 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRefMVO**](PartyRoleSpecificationRefMVO.md) |  | [optional] 
**Characteristic** | Pointer to [**[]CharacteristicMVO**](CharacteristicMVO.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRefMVO**](AccountRefMVO.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRefMVO**](AgreementRefMVO.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumMVO**](ContactMediumMVO.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRefMVO**](PaymentMethodRefMVO.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfileMVO**](CreditProfileMVO.md) |  | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleMVO**](RelatedPartyOrPartyRoleMVO.md) |  | [optional] 
**Status** | Pointer to **string** | Used to track the lifecycle status of the party role. | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyRoleMVO

`func NewPartyRoleMVO(type_ string, name string, engagedParty PartyRefMVO, ) *PartyRoleMVO`

NewPartyRoleMVO instantiates a new PartyRoleMVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleMVOWithDefaults

`func NewPartyRoleMVOWithDefaults() *PartyRoleMVO`

NewPartyRoleMVOWithDefaults instantiates a new PartyRoleMVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRoleMVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRoleMVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRoleMVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRoleMVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRoleMVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRoleMVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRoleMVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRoleMVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRoleMVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRoleMVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRoleMVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetName

`func (o *PartyRoleMVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRoleMVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRoleMVO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PartyRoleMVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyRoleMVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyRoleMVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyRoleMVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyRoleMVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyRoleMVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyRoleMVO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyRoleMVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyRoleMVO) GetEngagedParty() PartyRefMVO`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyRoleMVO) GetEngagedPartyOk() (*PartyRefMVO, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyRoleMVO) SetEngagedParty(v PartyRefMVO)`

SetEngagedParty sets EngagedParty field to given value.


### GetPartyRoleSpecification

`func (o *PartyRoleMVO) GetPartyRoleSpecification() PartyRoleSpecificationRefMVO`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyRoleMVO) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRefMVO, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyRoleMVO) SetPartyRoleSpecification(v PartyRoleSpecificationRefMVO)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyRoleMVO) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyRoleMVO) GetCharacteristic() []CharacteristicMVO`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyRoleMVO) GetCharacteristicOk() (*[]CharacteristicMVO, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyRoleMVO) SetCharacteristic(v []CharacteristicMVO)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyRoleMVO) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyRoleMVO) GetAccount() []AccountRefMVO`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyRoleMVO) GetAccountOk() (*[]AccountRefMVO, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyRoleMVO) SetAccount(v []AccountRefMVO)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyRoleMVO) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyRoleMVO) GetAgreement() []AgreementRefMVO`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyRoleMVO) GetAgreementOk() (*[]AgreementRefMVO, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyRoleMVO) SetAgreement(v []AgreementRefMVO)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyRoleMVO) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyRoleMVO) GetContactMedium() []ContactMediumMVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyRoleMVO) GetContactMediumOk() (*[]ContactMediumMVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyRoleMVO) SetContactMedium(v []ContactMediumMVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyRoleMVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyRoleMVO) GetPaymentMethod() []PaymentMethodRefMVO`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyRoleMVO) GetPaymentMethodOk() (*[]PaymentMethodRefMVO, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyRoleMVO) SetPaymentMethod(v []PaymentMethodRefMVO)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyRoleMVO) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyRoleMVO) GetCreditProfile() []CreditProfileMVO`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyRoleMVO) GetCreditProfileOk() (*[]CreditProfileMVO, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyRoleMVO) SetCreditProfile(v []CreditProfileMVO)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyRoleMVO) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyRoleMVO) GetRelatedParty() []RelatedPartyOrPartyRoleMVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyRoleMVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleMVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyRoleMVO) SetRelatedParty(v []RelatedPartyOrPartyRoleMVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyRoleMVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetStatus

`func (o *PartyRoleMVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyRoleMVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyRoleMVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyRoleMVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyRoleMVO) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyRoleMVO) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyRoleMVO) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyRoleMVO) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyRoleMVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyRoleMVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyRoleMVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyRoleMVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


