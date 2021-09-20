# \WorkflowsApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowGet**](WorkflowsApi.md#WorkflowGet) | **Get** /workflow | List Available Workflows
[**WorkflowWorkflowIDExecutePost**](WorkflowsApi.md#WorkflowWorkflowIDExecutePost) | **Post** /workflow/{workflowID}/execute | Execute Workflow
[**WorkflowWorkflowIDExecutionsGet**](WorkflowsApi.md#WorkflowWorkflowIDExecutionsGet) | **Get** /workflow/{workflowID}/executions | List Workflow Executions
[**WorkflowWorkflowIDGet**](WorkflowsApi.md#WorkflowWorkflowIDGet) | **Get** /workflow/{workflowID} | Get Workflow Details
[**WorkflowWorkflowIDVersionsGet**](WorkflowsApi.md#WorkflowWorkflowIDVersionsGet) | **Get** /workflow/{workflowID}/versions | List Workflow Versions
[**WorkflowWorkflowIDWebhooksGet**](WorkflowsApi.md#WorkflowWorkflowIDWebhooksGet) | **Get** /workflow/{workflowID}/webhooks | List Active Callbacks



## WorkflowGet

> ListResponse WorkflowGet(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Available Workflows

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowWorkflowIDExecutePost

> WorkflowExecutionResult WorkflowWorkflowIDExecutePost(ctx, workflowID).RequestBody(requestBody).V(v).Execute()

Execute Workflow

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowID := int32(56) // int32 | Workflow id
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | Execution Information
    v := int32(56) // int32 | If you want to run a version other then the current production version, you can supply the version number to run. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowWorkflowIDExecutePost(context.Background(), workflowID).RequestBody(requestBody).V(v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowWorkflowIDExecutePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowWorkflowIDExecutePost`: WorkflowExecutionResult
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowWorkflowIDExecutePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowID** | **int32** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowWorkflowIDExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | Execution Information | 
 **v** | **int32** | If you want to run a version other then the current production version, you can supply the version number to run. | 

### Return type

[**WorkflowExecutionResult**](WorkflowExecutionResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowWorkflowIDExecutionsGet

> ListResponse WorkflowWorkflowIDExecutionsGet(ctx, workflowID).V(v).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Workflow Executions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowID := int32(56) // int32 | Workflow Id
    v := int32(56) // int32 | If you want to return the executions of a particular workflow version instead of all workflows (optional)
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowWorkflowIDExecutionsGet(context.Background(), workflowID).V(v).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowWorkflowIDExecutionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowWorkflowIDExecutionsGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowWorkflowIDExecutionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowID** | **int32** | Workflow Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowWorkflowIDExecutionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v** | **int32** | If you want to return the executions of a particular workflow version instead of all workflows | 
 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowWorkflowIDGet

> Workflow WorkflowWorkflowIDGet(ctx, workflowID).Execute()

Get Workflow Details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowID := int32(56) // int32 | Workflow id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowWorkflowIDGet(context.Background(), workflowID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowWorkflowIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowWorkflowIDGet`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowWorkflowIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowID** | **int32** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowWorkflowIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workflow**](Workflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowWorkflowIDVersionsGet

> ListResponse WorkflowWorkflowIDVersionsGet(ctx, workflowID).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Workflow Versions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowID := int32(56) // int32 | Workflow id
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowWorkflowIDVersionsGet(context.Background(), workflowID).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowWorkflowIDVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowWorkflowIDVersionsGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowWorkflowIDVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowID** | **int32** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowWorkflowIDVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowWorkflowIDWebhooksGet

> ListResponse WorkflowWorkflowIDWebhooksGet(ctx, workflowID).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Active Callbacks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowID := int32(56) // int32 | Workflow ID
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowsApi.WorkflowWorkflowIDWebhooksGet(context.Background(), workflowID).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.WorkflowWorkflowIDWebhooksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowWorkflowIDWebhooksGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.WorkflowWorkflowIDWebhooksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowID** | **int32** | Workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowWorkflowIDWebhooksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

