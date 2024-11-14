# Go API client for openapi

TMF API Reference : TMF 632 - Party  Release: 22.5
The party API provides standardized mechanism for party management such as creation, update, retrieval, deletion, and notification of events. Party can be an individual or an organization that has any kind of relation with the enterprise. Party is created to record individual or organization information before the assignment of any role.
For example, within the context of a split billing mechanism, Party API allows creation of the individual or organization that will play the role of 3rd payer for a given offer and, then, allows consultation or update of his information.
Resources - Party (abstract base class with concrete subclasses Individual and Organization) Party API performs the following operations: - Retrieve an organization or an individual - Retrieve a collection of organizations or individuals according to given criteria - Create a new organization or a new individual - Update an existing organization or an existing individual - Delete an existing organization or an existing individual - Notify events on organization or individual

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 5.0.0
- Package version: 1.23.1
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://serverRoot/partyManagement/v5*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EventsSubscriptionAPI* | [**CreateHub**](docs/EventsSubscriptionAPI.md#createhub) | **Post** /hub | Create a subscription (hub) to receive Events
*EventsSubscriptionAPI* | [**HubDelete**](docs/EventsSubscriptionAPI.md#hubdelete) | **Delete** /hub/{id} | Remove a subscription (hub) to receive Events
*IndividualAPI* | [**CreateIndividual**](docs/IndividualAPI.md#createindividual) | **Post** /individual | Creates a Individual
*IndividualAPI* | [**DeleteIndividual**](docs/IndividualAPI.md#deleteindividual) | **Delete** /individual/{id} | Deletes a Individual
*IndividualAPI* | [**ListIndividual**](docs/IndividualAPI.md#listindividual) | **Get** /individual | List or find Individual objects
*IndividualAPI* | [**PatchIndividual**](docs/IndividualAPI.md#patchindividual) | **Patch** /individual/{id} | Updates partially a Individual
*IndividualAPI* | [**RetrieveIndividual**](docs/IndividualAPI.md#retrieveindividual) | **Get** /individual/{id} | Retrieves a Individual by ID
*NotificationListenerAPI* | [**IndividualAttributeValueChangeEvent**](docs/NotificationListenerAPI.md#individualattributevaluechangeevent) | **Post** /listener/individualAttributeValueChangeEvent | Client listener for entity IndividualAttributeValueChangeEvent
*NotificationListenerAPI* | [**IndividualCreateEvent**](docs/NotificationListenerAPI.md#individualcreateevent) | **Post** /listener/individualCreateEvent | Client listener for entity IndividualCreateEvent
*NotificationListenerAPI* | [**IndividualDeleteEvent**](docs/NotificationListenerAPI.md#individualdeleteevent) | **Post** /listener/individualDeleteEvent | Client listener for entity IndividualDeleteEvent
*NotificationListenerAPI* | [**IndividualStateChangeEvent**](docs/NotificationListenerAPI.md#individualstatechangeevent) | **Post** /listener/individualStateChangeEvent | Client listener for entity IndividualStateChangeEvent
*NotificationListenerAPI* | [**OrganizationAttributeValueChangeEvent**](docs/NotificationListenerAPI.md#organizationattributevaluechangeevent) | **Post** /listener/organizationAttributeValueChangeEvent | Client listener for entity OrganizationAttributeValueChangeEvent
*NotificationListenerAPI* | [**OrganizationCreateEvent**](docs/NotificationListenerAPI.md#organizationcreateevent) | **Post** /listener/organizationCreateEvent | Client listener for entity OrganizationCreateEvent
*NotificationListenerAPI* | [**OrganizationDeleteEvent**](docs/NotificationListenerAPI.md#organizationdeleteevent) | **Post** /listener/organizationDeleteEvent | Client listener for entity OrganizationDeleteEvent
*NotificationListenerAPI* | [**OrganizationStateChangeEvent**](docs/NotificationListenerAPI.md#organizationstatechangeevent) | **Post** /listener/organizationStateChangeEvent | Client listener for entity OrganizationStateChangeEvent
*OrganizationAPI* | [**CreateOrganization**](docs/OrganizationAPI.md#createorganization) | **Post** /organization | Creates a Organization
*OrganizationAPI* | [**DeleteOrganization**](docs/OrganizationAPI.md#deleteorganization) | **Delete** /organization/{id} | Deletes a Organization
*OrganizationAPI* | [**ListOrganization**](docs/OrganizationAPI.md#listorganization) | **Get** /organization | List or find Organization objects
*OrganizationAPI* | [**PatchOrganization**](docs/OrganizationAPI.md#patchorganization) | **Patch** /organization/{id} | Updates partially a Organization
*OrganizationAPI* | [**RetrieveOrganization**](docs/OrganizationAPI.md#retrieveorganization) | **Get** /organization/{id} | Retrieves a Organization by ID


## Documentation For Models

 - [AccountRef](docs/AccountRef.md)
 - [AccountRefFVO](docs/AccountRefFVO.md)
 - [AccountRefMVO](docs/AccountRefMVO.md)
 - [Addressable](docs/Addressable.md)
 - [AddressableFVO](docs/AddressableFVO.md)
 - [AgreementRef](docs/AgreementRef.md)
 - [AgreementRefFVO](docs/AgreementRefFVO.md)
 - [AgreementRefMVO](docs/AgreementRefMVO.md)
 - [Attachment](docs/Attachment.md)
 - [AttachmentFVO](docs/AttachmentFVO.md)
 - [AttachmentMVO](docs/AttachmentMVO.md)
 - [AttachmentRef](docs/AttachmentRef.md)
 - [AttachmentRefFVO](docs/AttachmentRefFVO.md)
 - [AttachmentRefMVO](docs/AttachmentRefMVO.md)
 - [AttachmentRefOrValue](docs/AttachmentRefOrValue.md)
 - [AttachmentRefOrValueFVO](docs/AttachmentRefOrValueFVO.md)
 - [AttachmentRefOrValueMVO](docs/AttachmentRefOrValueMVO.md)
 - [BaseEvent](docs/BaseEvent.md)
 - [BaseEventFVO](docs/BaseEventFVO.md)
 - [BaseEventMVO](docs/BaseEventMVO.md)
 - [BooleanArrayCharacteristic](docs/BooleanArrayCharacteristic.md)
 - [BooleanArrayCharacteristicFVO](docs/BooleanArrayCharacteristicFVO.md)
 - [BooleanArrayCharacteristicMVO](docs/BooleanArrayCharacteristicMVO.md)
 - [BooleanCharacteristic](docs/BooleanCharacteristic.md)
 - [BooleanCharacteristicFVO](docs/BooleanCharacteristicFVO.md)
 - [BooleanCharacteristicMVO](docs/BooleanCharacteristicMVO.md)
 - [BusinessPartner](docs/BusinessPartner.md)
 - [BusinessPartnerFVO](docs/BusinessPartnerFVO.md)
 - [BusinessPartnerMVO](docs/BusinessPartnerMVO.md)
 - [Characteristic](docs/Characteristic.md)
 - [CharacteristicFVO](docs/CharacteristicFVO.md)
 - [CharacteristicMVO](docs/CharacteristicMVO.md)
 - [CharacteristicRelationship](docs/CharacteristicRelationship.md)
 - [CharacteristicRelationshipFVO](docs/CharacteristicRelationshipFVO.md)
 - [CharacteristicRelationshipMVO](docs/CharacteristicRelationshipMVO.md)
 - [Consumer](docs/Consumer.md)
 - [ConsumerFVO](docs/ConsumerFVO.md)
 - [ConsumerMVO](docs/ConsumerMVO.md)
 - [ContactMedium](docs/ContactMedium.md)
 - [ContactMediumFVO](docs/ContactMediumFVO.md)
 - [ContactMediumMVO](docs/ContactMediumMVO.md)
 - [CreditProfile](docs/CreditProfile.md)
 - [CreditProfileFVO](docs/CreditProfileFVO.md)
 - [CreditProfileMVO](docs/CreditProfileMVO.md)
 - [Disability](docs/Disability.md)
 - [EmailContactMedium](docs/EmailContactMedium.md)
 - [EmailContactMediumFVO](docs/EmailContactMediumFVO.md)
 - [EmailContactMediumMVO](docs/EmailContactMediumMVO.md)
 - [Entity](docs/Entity.md)
 - [EntityFVO](docs/EntityFVO.md)
 - [EntityMVO](docs/EntityMVO.md)
 - [EntityRef](docs/EntityRef.md)
 - [EntityRefFVO](docs/EntityRefFVO.md)
 - [Error](docs/Error.md)
 - [Event](docs/Event.md)
 - [EventFVO](docs/EventFVO.md)
 - [EventMVO](docs/EventMVO.md)
 - [Extensible](docs/Extensible.md)
 - [ExtensibleFVO](docs/ExtensibleFVO.md)
 - [ExternalIdentifier](docs/ExternalIdentifier.md)
 - [ExternalIdentifierFVO](docs/ExternalIdentifierFVO.md)
 - [ExternalIdentifierMVO](docs/ExternalIdentifierMVO.md)
 - [FaxContactMedium](docs/FaxContactMedium.md)
 - [FaxContactMediumFVO](docs/FaxContactMediumFVO.md)
 - [FaxContactMediumMVO](docs/FaxContactMediumMVO.md)
 - [FloatCharacteristic](docs/FloatCharacteristic.md)
 - [FloatCharacteristicFVO](docs/FloatCharacteristicFVO.md)
 - [FloatCharacteristicMVO](docs/FloatCharacteristicMVO.md)
 - [GeographicAddressContactMedium](docs/GeographicAddressContactMedium.md)
 - [GeographicAddressContactMediumFVO](docs/GeographicAddressContactMediumFVO.md)
 - [GeographicAddressContactMediumMVO](docs/GeographicAddressContactMediumMVO.md)
 - [GeographicAddressRef](docs/GeographicAddressRef.md)
 - [GeographicAddressRefFVO](docs/GeographicAddressRefFVO.md)
 - [GeographicAddressRefMVO](docs/GeographicAddressRefMVO.md)
 - [Hub](docs/Hub.md)
 - [HubFVO](docs/HubFVO.md)
 - [Individual](docs/Individual.md)
 - [IndividualAttributeValueChangeEvent](docs/IndividualAttributeValueChangeEvent.md)
 - [IndividualAttributeValueChangeEventPayload](docs/IndividualAttributeValueChangeEventPayload.md)
 - [IndividualCreateEvent](docs/IndividualCreateEvent.md)
 - [IndividualCreateEventPayload](docs/IndividualCreateEventPayload.md)
 - [IndividualDeleteEvent](docs/IndividualDeleteEvent.md)
 - [IndividualDeleteEventPayload](docs/IndividualDeleteEventPayload.md)
 - [IndividualFVO](docs/IndividualFVO.md)
 - [IndividualIdentification](docs/IndividualIdentification.md)
 - [IndividualIdentificationFVO](docs/IndividualIdentificationFVO.md)
 - [IndividualIdentificationMVO](docs/IndividualIdentificationMVO.md)
 - [IndividualMVO](docs/IndividualMVO.md)
 - [IndividualStateChangeEvent](docs/IndividualStateChangeEvent.md)
 - [IndividualStateChangeEventPayload](docs/IndividualStateChangeEventPayload.md)
 - [IndividualStateType](docs/IndividualStateType.md)
 - [IntegerArrayCharacteristic](docs/IntegerArrayCharacteristic.md)
 - [IntegerArrayCharacteristicFVO](docs/IntegerArrayCharacteristicFVO.md)
 - [IntegerArrayCharacteristicMVO](docs/IntegerArrayCharacteristicMVO.md)
 - [IntegerCharacteristic](docs/IntegerCharacteristic.md)
 - [IntegerCharacteristicFVO](docs/IntegerCharacteristicFVO.md)
 - [IntegerCharacteristicMVO](docs/IntegerCharacteristicMVO.md)
 - [JsonPatch](docs/JsonPatch.md)
 - [LanguageAbility](docs/LanguageAbility.md)
 - [NumberArrayCharacteristic](docs/NumberArrayCharacteristic.md)
 - [NumberArrayCharacteristicFVO](docs/NumberArrayCharacteristicFVO.md)
 - [NumberArrayCharacteristicMVO](docs/NumberArrayCharacteristicMVO.md)
 - [NumberCharacteristic](docs/NumberCharacteristic.md)
 - [NumberCharacteristicFVO](docs/NumberCharacteristicFVO.md)
 - [NumberCharacteristicMVO](docs/NumberCharacteristicMVO.md)
 - [ObjectArrayCharacteristic](docs/ObjectArrayCharacteristic.md)
 - [ObjectArrayCharacteristicFVO](docs/ObjectArrayCharacteristicFVO.md)
 - [ObjectArrayCharacteristicMVO](docs/ObjectArrayCharacteristicMVO.md)
 - [ObjectCharacteristic](docs/ObjectCharacteristic.md)
 - [ObjectCharacteristicFVO](docs/ObjectCharacteristicFVO.md)
 - [ObjectCharacteristicMVO](docs/ObjectCharacteristicMVO.md)
 - [Organization](docs/Organization.md)
 - [OrganizationAttributeValueChangeEvent](docs/OrganizationAttributeValueChangeEvent.md)
 - [OrganizationAttributeValueChangeEventPayload](docs/OrganizationAttributeValueChangeEventPayload.md)
 - [OrganizationChildRelationship](docs/OrganizationChildRelationship.md)
 - [OrganizationChildRelationshipFVO](docs/OrganizationChildRelationshipFVO.md)
 - [OrganizationChildRelationshipMVO](docs/OrganizationChildRelationshipMVO.md)
 - [OrganizationCreateEvent](docs/OrganizationCreateEvent.md)
 - [OrganizationCreateEventPayload](docs/OrganizationCreateEventPayload.md)
 - [OrganizationDeleteEvent](docs/OrganizationDeleteEvent.md)
 - [OrganizationDeleteEventPayload](docs/OrganizationDeleteEventPayload.md)
 - [OrganizationFVO](docs/OrganizationFVO.md)
 - [OrganizationIdentification](docs/OrganizationIdentification.md)
 - [OrganizationIdentificationFVO](docs/OrganizationIdentificationFVO.md)
 - [OrganizationIdentificationMVO](docs/OrganizationIdentificationMVO.md)
 - [OrganizationMVO](docs/OrganizationMVO.md)
 - [OrganizationParentRelationship](docs/OrganizationParentRelationship.md)
 - [OrganizationParentRelationshipFVO](docs/OrganizationParentRelationshipFVO.md)
 - [OrganizationParentRelationshipMVO](docs/OrganizationParentRelationshipMVO.md)
 - [OrganizationRef](docs/OrganizationRef.md)
 - [OrganizationRefFVO](docs/OrganizationRefFVO.md)
 - [OrganizationRefMVO](docs/OrganizationRefMVO.md)
 - [OrganizationStateChangeEvent](docs/OrganizationStateChangeEvent.md)
 - [OrganizationStateChangeEventPayload](docs/OrganizationStateChangeEventPayload.md)
 - [OrganizationStateType](docs/OrganizationStateType.md)
 - [OtherNameIndividual](docs/OtherNameIndividual.md)
 - [OtherNameOrganization](docs/OtherNameOrganization.md)
 - [OtherNameOrganizationFVO](docs/OtherNameOrganizationFVO.md)
 - [OtherNameOrganizationMVO](docs/OtherNameOrganizationMVO.md)
 - [Party](docs/Party.md)
 - [PartyCreditProfile](docs/PartyCreditProfile.md)
 - [PartyCreditProfileFVO](docs/PartyCreditProfileFVO.md)
 - [PartyCreditProfileMVO](docs/PartyCreditProfileMVO.md)
 - [PartyFVO](docs/PartyFVO.md)
 - [PartyMVO](docs/PartyMVO.md)
 - [PartyOrPartyRole](docs/PartyOrPartyRole.md)
 - [PartyOrPartyRoleFVO](docs/PartyOrPartyRoleFVO.md)
 - [PartyOrPartyRoleMVO](docs/PartyOrPartyRoleMVO.md)
 - [PartyRef](docs/PartyRef.md)
 - [PartyRefFVO](docs/PartyRefFVO.md)
 - [PartyRefMVO](docs/PartyRefMVO.md)
 - [PartyRefOrPartyRoleRef](docs/PartyRefOrPartyRoleRef.md)
 - [PartyRole](docs/PartyRole.md)
 - [PartyRoleFVO](docs/PartyRoleFVO.md)
 - [PartyRoleMVO](docs/PartyRoleMVO.md)
 - [PartyRoleRef](docs/PartyRoleRef.md)
 - [PartyRoleRefFVO](docs/PartyRoleRefFVO.md)
 - [PartyRoleRefMVO](docs/PartyRoleRefMVO.md)
 - [PartyRoleSpecificationRef](docs/PartyRoleSpecificationRef.md)
 - [PartyRoleSpecificationRefFVO](docs/PartyRoleSpecificationRefFVO.md)
 - [PartyRoleSpecificationRefMVO](docs/PartyRoleSpecificationRefMVO.md)
 - [PatchIndividual200Response](docs/PatchIndividual200Response.md)
 - [PatchOrganization200Response](docs/PatchOrganization200Response.md)
 - [PaymentMethodRef](docs/PaymentMethodRef.md)
 - [PaymentMethodRefFVO](docs/PaymentMethodRefFVO.md)
 - [PaymentMethodRefMVO](docs/PaymentMethodRefMVO.md)
 - [PhoneContactMedium](docs/PhoneContactMedium.md)
 - [PhoneContactMediumFVO](docs/PhoneContactMediumFVO.md)
 - [PhoneContactMediumMVO](docs/PhoneContactMediumMVO.md)
 - [Producer](docs/Producer.md)
 - [ProducerFVO](docs/ProducerFVO.md)
 - [ProducerMVO](docs/ProducerMVO.md)
 - [ProductRelationshipType](docs/ProductRelationshipType.md)
 - [Quantity](docs/Quantity.md)
 - [RelatedPartyOrPartyRole](docs/RelatedPartyOrPartyRole.md)
 - [RelatedPartyOrPartyRoleFVO](docs/RelatedPartyOrPartyRoleFVO.md)
 - [RelatedPartyOrPartyRoleMVO](docs/RelatedPartyOrPartyRoleMVO.md)
 - [RelatedPartyRefOrPartyRoleRef](docs/RelatedPartyRefOrPartyRoleRef.md)
 - [RelatedPartyRefOrPartyRoleRefFVO](docs/RelatedPartyRefOrPartyRoleRefFVO.md)
 - [RelatedPartyRefOrPartyRoleRefMVO](docs/RelatedPartyRefOrPartyRoleRefMVO.md)
 - [Skill](docs/Skill.md)
 - [SocialContactMedium](docs/SocialContactMedium.md)
 - [SocialContactMediumFVO](docs/SocialContactMediumFVO.md)
 - [SocialContactMediumMVO](docs/SocialContactMediumMVO.md)
 - [StringArrayCharacteristic](docs/StringArrayCharacteristic.md)
 - [StringArrayCharacteristicFVO](docs/StringArrayCharacteristicFVO.md)
 - [StringArrayCharacteristicMVO](docs/StringArrayCharacteristicMVO.md)
 - [StringCharacteristic](docs/StringCharacteristic.md)
 - [StringCharacteristicFVO](docs/StringCharacteristicFVO.md)
 - [StringCharacteristicMVO](docs/StringCharacteristicMVO.md)
 - [Supplier](docs/Supplier.md)
 - [SupplierFVO](docs/SupplierFVO.md)
 - [SupplierMVO](docs/SupplierMVO.md)
 - [TaxDefinition](docs/TaxDefinition.md)
 - [TaxDefinitionFVO](docs/TaxDefinitionFVO.md)
 - [TaxDefinitionMVO](docs/TaxDefinitionMVO.md)
 - [TaxExemptionCertificate](docs/TaxExemptionCertificate.md)
 - [TaxExemptionCertificateFVO](docs/TaxExemptionCertificateFVO.md)
 - [TaxExemptionCertificateMVO](docs/TaxExemptionCertificateMVO.md)
 - [TimePeriod](docs/TimePeriod.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


