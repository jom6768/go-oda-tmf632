# \IndividualAPI

All URIs are relative to *https://serverRoot/partyManagement/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividual**](IndividualAPI.md#CreateIndividual) | **Post** /individual | Creates a Individual
[**DeleteIndividual**](IndividualAPI.md#DeleteIndividual) | **Delete** /individual/{id} | Deletes a Individual
[**ListIndividual**](IndividualAPI.md#ListIndividual) | **Get** /individual | List or find Individual objects
[**PatchIndividual**](IndividualAPI.md#PatchIndividual) | **Patch** /individual/{id} | Updates partially a Individual
[**RetrieveIndividual**](IndividualAPI.md#RetrieveIndividual) | **Get** /individual/{id} | Retrieves a Individual by ID



## CreateIndividual

> Individual CreateIndividual(ctx).IndividualFVO(individualFVO).Fields(fields).Execute()

Creates a Individual



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
	individualFVO := *openapiclient.NewIndividualFVO("FamilyName_example", "GivenName_example", "Type_example") // IndividualFVO | The Individual to be created
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndividualAPI.CreateIndividual(context.Background()).IndividualFVO(individualFVO).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPI.CreateIndividual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIndividual`: Individual
	fmt.Fprintf(os.Stdout, "Response from `IndividualAPI.CreateIndividual`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **individualFVO** | [**IndividualFVO**](IndividualFVO.md) | The Individual to be created | 
 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Individual**](Individual.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividual

> DeleteIndividual(ctx, id).Execute()

Deletes a Individual



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
	r, err := apiClient.IndividualAPI.DeleteIndividual(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPI.DeleteIndividual``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndividualRequest struct via the builder pattern


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


## ListIndividual

> []Individual ListIndividual(ctx).Fields(fields).Offset(offset).Limit(limit).Execute()

List or find Individual objects



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
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)
	offset := int32(56) // int32 | Requested index for start of resources to be provided in response (optional)
	limit := int32(56) // int32 | Requested number of resources to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndividualAPI.ListIndividual(context.Background()).Fields(fields).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPI.ListIndividual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIndividual`: []Individual
	fmt.Fprintf(os.Stdout, "Response from `IndividualAPI.ListIndividual`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated properties to be provided in response | 
 **offset** | **int32** | Requested index for start of resources to be provided in response | 
 **limit** | **int32** | Requested number of resources to be provided in response | 

### Return type

[**[]Individual**](Individual.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIndividual

> Individual PatchIndividual(ctx, id).IndividualMVO(individualMVO).Fields(fields).Execute()

Updates partially a Individual



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
	individualMVO := *openapiclient.NewIndividualMVO("Type_example") // IndividualMVO | The Individual to be patched
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndividualAPI.PatchIndividual(context.Background(), id).IndividualMVO(individualMVO).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPI.PatchIndividual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIndividual`: Individual
	fmt.Fprintf(os.Stdout, "Response from `IndividualAPI.PatchIndividual`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **individualMVO** | [**IndividualMVO**](IndividualMVO.md) | The Individual to be patched | 
 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Individual**](Individual.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json, application/json-patch-query+json
- **Accept**: application/json, application/merge-patch+json, application/json-patch+json, application/json-patch-query+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndividual

> Individual RetrieveIndividual(ctx, id).Fields(fields).Execute()

Retrieves a Individual by ID



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
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IndividualAPI.RetrieveIndividual(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPI.RetrieveIndividual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveIndividual`: Individual
	fmt.Fprintf(os.Stdout, "Response from `IndividualAPI.RetrieveIndividual`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndividualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Individual**](Individual.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

