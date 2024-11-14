# \OrganizationAPI

All URIs are relative to *https://serverRoot/partyManagement/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationAPI.md#CreateOrganization) | **Post** /organization | Creates a Organization
[**DeleteOrganization**](OrganizationAPI.md#DeleteOrganization) | **Delete** /organization/{id} | Deletes a Organization
[**ListOrganization**](OrganizationAPI.md#ListOrganization) | **Get** /organization | List or find Organization objects
[**PatchOrganization**](OrganizationAPI.md#PatchOrganization) | **Patch** /organization/{id} | Updates partially a Organization
[**RetrieveOrganization**](OrganizationAPI.md#RetrieveOrganization) | **Get** /organization/{id} | Retrieves a Organization by ID



## CreateOrganization

> Organization CreateOrganization(ctx).OrganizationFVO(organizationFVO).Fields(fields).Execute()

Creates a Organization



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
	organizationFVO := *openapiclient.NewOrganizationFVO("Name_example", "Type_example") // OrganizationFVO | The Organization to be created
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).OrganizationFVO(organizationFVO).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationFVO** | [**OrganizationFVO**](OrganizationFVO.md) | The Organization to be created | 
 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, id).Execute()

Deletes a Organization



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
	r, err := apiClient.OrganizationAPI.DeleteOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteOrganization``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


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


## ListOrganization

> []Organization ListOrganization(ctx).Fields(fields).Offset(offset).Limit(limit).Execute()

List or find Organization objects



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
	resp, r, err := apiClient.OrganizationAPI.ListOrganization(context.Background()).Fields(fields).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.ListOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganization`: []Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.ListOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | Comma-separated properties to be provided in response | 
 **offset** | **int32** | Requested index for start of resources to be provided in response | 
 **limit** | **int32** | Requested number of resources to be provided in response | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrganization

> Organization PatchOrganization(ctx, id).OrganizationMVO(organizationMVO).Fields(fields).Execute()

Updates partially a Organization



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
	organizationMVO := *openapiclient.NewOrganizationMVO("Type_example") // OrganizationMVO | The Organization to be patched
	fields := "fields_example" // string | Comma-separated properties to be provided in response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.PatchOrganization(context.Background(), id).OrganizationMVO(organizationMVO).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.PatchOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.PatchOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationMVO** | [**OrganizationMVO**](OrganizationMVO.md) | The Organization to be patched | 
 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/merge-patch+json, application/json-patch+json, application/json-patch-query+json
- **Accept**: application/json, application/merge-patch+json, application/json-patch+json, application/json-patch-query+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOrganization

> Organization RetrieveOrganization(ctx, id).Fields(fields).Execute()

Retrieves a Organization by ID



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
	resp, r, err := apiClient.OrganizationAPI.RetrieveOrganization(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RetrieveOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveOrganization`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RetrieveOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the Resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Comma-separated properties to be provided in response | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

