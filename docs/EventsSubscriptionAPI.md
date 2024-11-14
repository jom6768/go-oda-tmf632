# \EventsSubscriptionAPI

All URIs are relative to *https://serverRoot/partyManagement/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHub**](EventsSubscriptionAPI.md#CreateHub) | **Post** /hub | Create a subscription (hub) to receive Events
[**HubDelete**](EventsSubscriptionAPI.md#HubDelete) | **Delete** /hub/{id} | Remove a subscription (hub) to receive Events



## CreateHub

> Hub CreateHub(ctx).HubFVO(hubFVO).Execute()

Create a subscription (hub) to receive Events



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
	hubFVO := *openapiclient.NewHubFVO("Type_example", "Callback_example") // HubFVO | Data containing the callback endpoint to deliver the information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsSubscriptionAPI.CreateHub(context.Background()).HubFVO(hubFVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsSubscriptionAPI.CreateHub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHub`: Hub
	fmt.Fprintf(os.Stdout, "Response from `EventsSubscriptionAPI.CreateHub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hubFVO** | [**HubFVO**](HubFVO.md) | Data containing the callback endpoint to deliver the information | 

### Return type

[**Hub**](Hub.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HubDelete

> HubDelete(ctx, id).Execute()

Remove a subscription (hub) to receive Events



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
	id := "id_example" // string | Identifier of the Resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsSubscriptionAPI.HubDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsSubscriptionAPI.HubDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiHubDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

