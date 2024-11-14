# \NotificationListenerAPI

All URIs are relative to *https://serverRoot/partyManagement/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndividualAttributeValueChangeEvent**](NotificationListenerAPI.md#IndividualAttributeValueChangeEvent) | **Post** /listener/individualAttributeValueChangeEvent | Client listener for entity IndividualAttributeValueChangeEvent
[**IndividualCreateEvent**](NotificationListenerAPI.md#IndividualCreateEvent) | **Post** /listener/individualCreateEvent | Client listener for entity IndividualCreateEvent
[**IndividualDeleteEvent**](NotificationListenerAPI.md#IndividualDeleteEvent) | **Post** /listener/individualDeleteEvent | Client listener for entity IndividualDeleteEvent
[**IndividualStateChangeEvent**](NotificationListenerAPI.md#IndividualStateChangeEvent) | **Post** /listener/individualStateChangeEvent | Client listener for entity IndividualStateChangeEvent
[**OrganizationAttributeValueChangeEvent**](NotificationListenerAPI.md#OrganizationAttributeValueChangeEvent) | **Post** /listener/organizationAttributeValueChangeEvent | Client listener for entity OrganizationAttributeValueChangeEvent
[**OrganizationCreateEvent**](NotificationListenerAPI.md#OrganizationCreateEvent) | **Post** /listener/organizationCreateEvent | Client listener for entity OrganizationCreateEvent
[**OrganizationDeleteEvent**](NotificationListenerAPI.md#OrganizationDeleteEvent) | **Post** /listener/organizationDeleteEvent | Client listener for entity OrganizationDeleteEvent
[**OrganizationStateChangeEvent**](NotificationListenerAPI.md#OrganizationStateChangeEvent) | **Post** /listener/organizationStateChangeEvent | Client listener for entity OrganizationStateChangeEvent



## IndividualAttributeValueChangeEvent

> IndividualAttributeValueChangeEvent(ctx).IndividualAttributeValueChangeEvent(individualAttributeValueChangeEvent).Execute()

Client listener for entity IndividualAttributeValueChangeEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	individualAttributeValueChangeEvent := *openapiclient.NewIndividualAttributeValueChangeEvent("Type_example") // IndividualAttributeValueChangeEvent | Individual attributeValueChange Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.IndividualAttributeValueChangeEvent(context.Background()).IndividualAttributeValueChangeEvent(individualAttributeValueChangeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.IndividualAttributeValueChangeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndividualAttributeValueChangeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **individualAttributeValueChangeEvent** | [**IndividualAttributeValueChangeEvent**](IndividualAttributeValueChangeEvent.md) | Individual attributeValueChange Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndividualCreateEvent

> IndividualCreateEvent(ctx).IndividualCreateEvent(individualCreateEvent).Execute()

Client listener for entity IndividualCreateEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	individualCreateEvent := *openapiclient.NewIndividualCreateEvent("Type_example") // IndividualCreateEvent | Individual create Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.IndividualCreateEvent(context.Background()).IndividualCreateEvent(individualCreateEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.IndividualCreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndividualCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **individualCreateEvent** | [**IndividualCreateEvent**](IndividualCreateEvent.md) | Individual create Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndividualDeleteEvent

> IndividualDeleteEvent(ctx).IndividualDeleteEvent(individualDeleteEvent).Execute()

Client listener for entity IndividualDeleteEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	individualDeleteEvent := *openapiclient.NewIndividualDeleteEvent("Type_example") // IndividualDeleteEvent | Individual delete Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.IndividualDeleteEvent(context.Background()).IndividualDeleteEvent(individualDeleteEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.IndividualDeleteEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndividualDeleteEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **individualDeleteEvent** | [**IndividualDeleteEvent**](IndividualDeleteEvent.md) | Individual delete Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndividualStateChangeEvent

> IndividualStateChangeEvent(ctx).IndividualStateChangeEvent(individualStateChangeEvent).Execute()

Client listener for entity IndividualStateChangeEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	individualStateChangeEvent := *openapiclient.NewIndividualStateChangeEvent("Type_example") // IndividualStateChangeEvent | Individual stateChange Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.IndividualStateChangeEvent(context.Background()).IndividualStateChangeEvent(individualStateChangeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.IndividualStateChangeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndividualStateChangeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **individualStateChangeEvent** | [**IndividualStateChangeEvent**](IndividualStateChangeEvent.md) | Individual stateChange Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationAttributeValueChangeEvent

> OrganizationAttributeValueChangeEvent(ctx).OrganizationAttributeValueChangeEvent(organizationAttributeValueChangeEvent).Execute()

Client listener for entity OrganizationAttributeValueChangeEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationAttributeValueChangeEvent := *openapiclient.NewOrganizationAttributeValueChangeEvent("Type_example") // OrganizationAttributeValueChangeEvent | Organization attributeValueChange Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.OrganizationAttributeValueChangeEvent(context.Background()).OrganizationAttributeValueChangeEvent(organizationAttributeValueChangeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.OrganizationAttributeValueChangeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationAttributeValueChangeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationAttributeValueChangeEvent** | [**OrganizationAttributeValueChangeEvent**](OrganizationAttributeValueChangeEvent.md) | Organization attributeValueChange Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationCreateEvent

> OrganizationCreateEvent(ctx).OrganizationCreateEvent(organizationCreateEvent).Execute()

Client listener for entity OrganizationCreateEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationCreateEvent := *openapiclient.NewOrganizationCreateEvent("Type_example") // OrganizationCreateEvent | Organization create Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.OrganizationCreateEvent(context.Background()).OrganizationCreateEvent(organizationCreateEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.OrganizationCreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationCreateEvent** | [**OrganizationCreateEvent**](OrganizationCreateEvent.md) | Organization create Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationDeleteEvent

> OrganizationDeleteEvent(ctx).OrganizationDeleteEvent(organizationDeleteEvent).Execute()

Client listener for entity OrganizationDeleteEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationDeleteEvent := *openapiclient.NewOrganizationDeleteEvent("Type_example") // OrganizationDeleteEvent | Organization delete Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.OrganizationDeleteEvent(context.Background()).OrganizationDeleteEvent(organizationDeleteEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.OrganizationDeleteEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationDeleteEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationDeleteEvent** | [**OrganizationDeleteEvent**](OrganizationDeleteEvent.md) | Organization delete Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationStateChangeEvent

> OrganizationStateChangeEvent(ctx).OrganizationStateChangeEvent(organizationStateChangeEvent).Execute()

Client listener for entity OrganizationStateChangeEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationStateChangeEvent := *openapiclient.NewOrganizationStateChangeEvent("Type_example") // OrganizationStateChangeEvent | Organization stateChange Event payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationListenerAPI.OrganizationStateChangeEvent(context.Background()).OrganizationStateChangeEvent(organizationStateChangeEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationListenerAPI.OrganizationStateChangeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationStateChangeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationStateChangeEvent** | [**OrganizationStateChangeEvent**](OrganizationStateChangeEvent.md) | Organization stateChange Event payload | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

