# PartyRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When sub-classing, this defines the sub-class Extensible name | 
**BaseType** | Pointer to **string** | When sub-classing, this defines the super-class | [optional] 
**SchemaLocation** | Pointer to **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] 
**Href** | Pointer to **string** | Hyperlink reference | [optional] 
**Id** | Pointer to **string** | unique identifier | [optional] 
**Name** | Pointer to **string** | A word, term, or phrase by which the PartyRole is known and distinguished from other PartyRoles. It&#39;s the name of the PartyRole unique entity. | [optional] 
**Description** | Pointer to **string** | A description of the PartyRole. | [optional] 
**Role** | Pointer to **string** | Role played by the engagedParty in this context. As role is defined by partyRoleSpecification, this role attribute can be used to precise the role defined by partyRoleSpecification, or it can be used to define the role in case there is no partyRoleSpecification. | [optional] 
**EngagedParty** | Pointer to [**PartyRef**](PartyRef.md) |  | [optional] 
**PartyRoleSpecification** | Pointer to [**PartyRoleSpecificationRef**](PartyRoleSpecificationRef.md) |  | [optional] 
**Characteristic** | Pointer to [**[]Characteristic**](Characteristic.md) | Describes the characteristic of a party role. | [optional] 
**Account** | Pointer to [**[]AccountRef**](AccountRef.md) |  | [optional] 
**Agreement** | Pointer to [**[]AgreementRef**](AgreementRef.md) |  | [optional] 
**ContactMedium** | Pointer to [**[]ContactMedium**](ContactMedium.md) |  | [optional] 
**PaymentMethod** | Pointer to [**[]PaymentMethodRef**](PaymentMethodRef.md) |  | [optional] 
**CreditProfile** | Pointer to [**[]CreditProfile**](CreditProfile.md) |  | [optional] 
**RelatedParty** | Pointer to [**[]RelatedPartyOrPartyRole**](RelatedPartyOrPartyRole.md) |  | [optional] 
**Status** | Pointer to **string** | Used to track the lifecycle status of the party role. | [optional] 
**StatusReason** | Pointer to **string** | A string providing an explanation on the value of the status lifecycle. For instance if the status is Rejected, statusReason will provide the reason for rejection. | [optional] 
**ValidFor** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 

## Methods

### NewPartyRole

`func NewPartyRole(type_ string, ) *PartyRole`

NewPartyRole instantiates a new PartyRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyRoleWithDefaults

`func NewPartyRoleWithDefaults() *PartyRole`

NewPartyRoleWithDefaults instantiates a new PartyRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PartyRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartyRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartyRole) SetType(v string)`

SetType sets Type field to given value.


### GetBaseType

`func (o *PartyRole) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *PartyRole) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *PartyRole) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *PartyRole) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetSchemaLocation

`func (o *PartyRole) GetSchemaLocation() string`

GetSchemaLocation returns the SchemaLocation field if non-nil, zero value otherwise.

### GetSchemaLocationOk

`func (o *PartyRole) GetSchemaLocationOk() (*string, bool)`

GetSchemaLocationOk returns a tuple with the SchemaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaLocation

`func (o *PartyRole) SetSchemaLocation(v string)`

SetSchemaLocation sets SchemaLocation field to given value.

### HasSchemaLocation

`func (o *PartyRole) HasSchemaLocation() bool`

HasSchemaLocation returns a boolean if a field has been set.

### GetHref

`func (o *PartyRole) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PartyRole) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PartyRole) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PartyRole) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PartyRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartyRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PartyRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PartyRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartyRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartyRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartyRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *PartyRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PartyRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PartyRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PartyRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEngagedParty

`func (o *PartyRole) GetEngagedParty() PartyRef`

GetEngagedParty returns the EngagedParty field if non-nil, zero value otherwise.

### GetEngagedPartyOk

`func (o *PartyRole) GetEngagedPartyOk() (*PartyRef, bool)`

GetEngagedPartyOk returns a tuple with the EngagedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagedParty

`func (o *PartyRole) SetEngagedParty(v PartyRef)`

SetEngagedParty sets EngagedParty field to given value.

### HasEngagedParty

`func (o *PartyRole) HasEngagedParty() bool`

HasEngagedParty returns a boolean if a field has been set.

### GetPartyRoleSpecification

`func (o *PartyRole) GetPartyRoleSpecification() PartyRoleSpecificationRef`

GetPartyRoleSpecification returns the PartyRoleSpecification field if non-nil, zero value otherwise.

### GetPartyRoleSpecificationOk

`func (o *PartyRole) GetPartyRoleSpecificationOk() (*PartyRoleSpecificationRef, bool)`

GetPartyRoleSpecificationOk returns a tuple with the PartyRoleSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyRoleSpecification

`func (o *PartyRole) SetPartyRoleSpecification(v PartyRoleSpecificationRef)`

SetPartyRoleSpecification sets PartyRoleSpecification field to given value.

### HasPartyRoleSpecification

`func (o *PartyRole) HasPartyRoleSpecification() bool`

HasPartyRoleSpecification returns a boolean if a field has been set.

### GetCharacteristic

`func (o *PartyRole) GetCharacteristic() []Characteristic`

GetCharacteristic returns the Characteristic field if non-nil, zero value otherwise.

### GetCharacteristicOk

`func (o *PartyRole) GetCharacteristicOk() (*[]Characteristic, bool)`

GetCharacteristicOk returns a tuple with the Characteristic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristic

`func (o *PartyRole) SetCharacteristic(v []Characteristic)`

SetCharacteristic sets Characteristic field to given value.

### HasCharacteristic

`func (o *PartyRole) HasCharacteristic() bool`

HasCharacteristic returns a boolean if a field has been set.

### GetAccount

`func (o *PartyRole) GetAccount() []AccountRef`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartyRole) GetAccountOk() (*[]AccountRef, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartyRole) SetAccount(v []AccountRef)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PartyRole) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgreement

`func (o *PartyRole) GetAgreement() []AgreementRef`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *PartyRole) GetAgreementOk() (*[]AgreementRef, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *PartyRole) SetAgreement(v []AgreementRef)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *PartyRole) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetContactMedium

`func (o *PartyRole) GetContactMedium() []ContactMedium`

GetContactMedium returns the ContactMedium field if non-nil, zero value otherwise.

### GetContactMediumOk

`func (o *PartyRole) GetContactMediumOk() (*[]ContactMedium, bool)`

GetContactMediumOk returns a tuple with the ContactMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMedium

`func (o *PartyRole) SetContactMedium(v []ContactMedium)`

SetContactMedium sets ContactMedium field to given value.

### HasContactMedium

`func (o *PartyRole) HasContactMedium() bool`

HasContactMedium returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PartyRole) GetPaymentMethod() []PaymentMethodRef`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PartyRole) GetPaymentMethodOk() (*[]PaymentMethodRef, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PartyRole) SetPaymentMethod(v []PaymentMethodRef)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PartyRole) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCreditProfile

`func (o *PartyRole) GetCreditProfile() []CreditProfile`

GetCreditProfile returns the CreditProfile field if non-nil, zero value otherwise.

### GetCreditProfileOk

`func (o *PartyRole) GetCreditProfileOk() (*[]CreditProfile, bool)`

GetCreditProfileOk returns a tuple with the CreditProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProfile

`func (o *PartyRole) SetCreditProfile(v []CreditProfile)`

SetCreditProfile sets CreditProfile field to given value.

### HasCreditProfile

`func (o *PartyRole) HasCreditProfile() bool`

HasCreditProfile returns a boolean if a field has been set.

### GetRelatedParty

`func (o *PartyRole) GetRelatedParty() []RelatedPartyOrPartyRole`

GetRelatedParty returns the RelatedParty field if non-nil, zero value otherwise.

### GetRelatedPartyOk

`func (o *PartyRole) GetRelatedPartyOk() (*[]RelatedPartyOrPartyRole, bool)`

GetRelatedPartyOk returns a tuple with the RelatedParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedParty

`func (o *PartyRole) SetRelatedParty(v []RelatedPartyOrPartyRole)`

SetRelatedParty sets RelatedParty field to given value.

### HasRelatedParty

`func (o *PartyRole) HasRelatedParty() bool`

HasRelatedParty returns a boolean if a field has been set.

### GetStatus

`func (o *PartyRole) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartyRole) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartyRole) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartyRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *PartyRole) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PartyRole) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PartyRole) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *PartyRole) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetValidFor

`func (o *PartyRole) GetValidFor() TimePeriod`

GetValidFor returns the ValidFor field if non-nil, zero value otherwise.

### GetValidForOk

`func (o *PartyRole) GetValidForOk() (*TimePeriod, bool)`

GetValidForOk returns a tuple with the ValidFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFor

`func (o *PartyRole) SetValidFor(v TimePeriod)`

SetValidFor sets ValidFor field to given value.

### HasValidFor

`func (o *PartyRole) HasValidFor() bool`

HasValidFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


