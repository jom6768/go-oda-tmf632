# PartyRoleFVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**Name** | **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | [**PartyRefFVO**](PartyRefFVO.md) |  | 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRefFVO**](PartyRoleSpecificationRefFVO.md) |  | [optional] 
**Characteristic** | Pointer to [**[]CharacteristicFVO**](CharacteristicFVO.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRefFVO**](AccountRefFVO.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRefFVO**](AgreementRefFVO.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMediumFVO**](ContactMediumFVO.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRefFVO**](PaymentMethodRefFVO.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfileFVO**](CreditProfileFVO.md) |  | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRoleFVO**](RelatedPartyOrPartyRoleFVO.md) |  | [optional] 
**Status** | Pointer to **string** | Used to track the lifecycle status of the party role. | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyRoleFVO

`func NewPartyRoleFVO(type_ string, name string, engagedParty PartyRefFVO, ) *PartyRoleFVO`

NewPartyRoleFVO instantiates a new PartyRoleFVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleFVOWithDefaults

`func NewPartyRoleFVOWithDefaults() *PartyRoleFVO`

NewPartyRoleFVOWithDefaults instantiates a new PartyRoleFVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRoleFVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRoleFVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRoleFVO) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRoleFVO) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRoleFVO) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRoleFVO) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRoleFVO) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRoleFVO) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRoleFVO) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRoleFVO) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRoleFVO) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetId

`func (o *PartyRoleFVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRoleFVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRoleFVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartyRoleFVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PartyRoleFVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRoleFVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRoleFVO) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PartyRoleFVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyRoleFVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyRoleFVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyRoleFVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyRoleFVO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyRoleFVO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyRoleFVO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyRoleFVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyRoleFVO) GetEngagedParty() PartyRefFVO`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyRoleFVO) GetEngagedPartyOk() (*PartyRefFVO, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyRoleFVO) SetEngagedParty(v PartyRefFVO)`

SetEngagedParty sets EngagedParty field to given value.


### GetPartyRoleSpecification

`func (o *PartyRoleFVO) GetPartyRoleSpecification() PartyRoleSpecificationRefFVO`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyRoleFVO) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRefFVO, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyRoleFVO) SetPartyRoleSpecification(v PartyRoleSpecificationRefFVO)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyRoleFVO) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyRoleFVO) GetCharacteristic() []CharacteristicFVO`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyRoleFVO) GetCharacteristicOk() (*[]CharacteristicFVO, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyRoleFVO) SetCharacteristic(v []CharacteristicFVO)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyRoleFVO) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyRoleFVO) GetAccount() []AccountRefFVO`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyRoleFVO) GetAccountOk() (*[]AccountRefFVO, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyRoleFVO) SetAccount(v []AccountRefFVO)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyRoleFVO) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyRoleFVO) GetAgreement() []AgreementRefFVO`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyRoleFVO) GetAgreementOk() (*[]AgreementRefFVO, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyRoleFVO) SetAgreement(v []AgreementRefFVO)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyRoleFVO) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyRoleFVO) GetContactMedium() []ContactMediumFVO`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyRoleFVO) GetContactMediumOk() (*[]ContactMediumFVO, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyRoleFVO) SetContactMedium(v []ContactMediumFVO)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyRoleFVO) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyRoleFVO) GetPaymentMethod() []PaymentMethodRefFVO`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyRoleFVO) GetPaymentMethodOk() (*[]PaymentMethodRefFVO, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyRoleFVO) SetPaymentMethod(v []PaymentMethodRefFVO)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyRoleFVO) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyRoleFVO) GetCreditProfile() []CreditProfileFVO`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyRoleFVO) GetCreditProfileOk() (*[]CreditProfileFVO, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyRoleFVO) SetCreditProfile(v []CreditProfileFVO)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyRoleFVO) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyRoleFVO) GetRelatedParty() []RelatedPartyOrPartyRoleFVO`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyRoleFVO) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRoleFVO, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyRoleFVO) SetRelatedParty(v []RelatedPartyOrPartyRoleFVO)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyRoleFVO) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetStatus

`func (o *PartyRoleFVO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyRoleFVO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyRoleFVO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyRoleFVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyRoleFVO) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyRoleFVO) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyRoleFVO) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyRoleFVO) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyRoleFVO) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyRoleFVO) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyRoleFVO) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyRoleFVO) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


