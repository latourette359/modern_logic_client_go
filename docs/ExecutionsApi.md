# \ExecutionsApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerCustomerIdExecutionsGet**](ExecutionsApi.md#CustomerCustomerIdExecutionsGet) | **Get** /customer/{customerId}/executions | List Customer Executions
[**ExecutionExecutionIdGet**](ExecutionsApi.md#ExecutionExecutionIdGet) | **Get** /execution/{executionId} | Get Execution Details
[**ExecutionExecutionIdResumePost**](ExecutionsApi.md#ExecutionExecutionIdResumePost) | **Post** /execution/{executionId}/resume | Resume Execution
[**ExecutionGet**](ExecutionsApi.md#ExecutionGet) | **Get** /execution | List executions



## CustomerCustomerIdExecutionsGet

> ListResponse CustomerCustomerIdExecutionsGet(ctx, customerId).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Customer Executions

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
    customerId := "customerId_example" // string | Customer id that the client supplied
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExecutionsApi.CustomerCustomerIdExecutionsGet(context.Background(), customerId).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsApi.CustomerCustomerIdExecutionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCustomerIdExecutionsGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExecutionsApi.CustomerCustomerIdExecutionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Customer id that the client supplied | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomerIdExecutionsGetRequest struct via the builder pattern


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


## ExecutionExecutionIdGet

> Execution ExecutionExecutionIdGet(ctx, executionId).Execute()

Get Execution Details

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
    executionId := int32(56) // int32 | Execution id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExecutionsApi.ExecutionExecutionIdGet(context.Background(), executionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsApi.ExecutionExecutionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutionExecutionIdGet`: Execution
    fmt.Fprintf(os.Stdout, "Response from `ExecutionsApi.ExecutionExecutionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int32** | Execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecutionExecutionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Execution**](Execution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecutionExecutionIdResumePost

> WorkflowExecutionResult ExecutionExecutionIdResumePost(ctx, executionId).RequestBody(requestBody).Execute()

Resume Execution

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
    executionId := int32(56) // int32 | execution id
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | Execution Information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExecutionsApi.ExecutionExecutionIdResumePost(context.Background(), executionId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsApi.ExecutionExecutionIdResumePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutionExecutionIdResumePost`: WorkflowExecutionResult
    fmt.Fprintf(os.Stdout, "Response from `ExecutionsApi.ExecutionExecutionIdResumePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int32** | execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecutionExecutionIdResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | Execution Information | 

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


## ExecutionGet

> ListResponse ExecutionGet(ctx).PageSize(pageSize).PageNumber(pageNumber).AlertType(alertType).Before(before).After(after).Execute()

List executions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)
    alertType := "alertType_example" // string | The alert status of this execution (optional)
    before := time.Now() // string | Filter executions to those that occurred before the given date. (optional)
    after := time.Now() // string | Filter executions to those that occurred after the given date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExecutionsApi.ExecutionGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).AlertType(alertType).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsApi.ExecutionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecutionGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExecutionsApi.ExecutionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecutionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 
 **alertType** | **string** | The alert status of this execution | 
 **before** | **string** | Filter executions to those that occurred before the given date. | 
 **after** | **string** | Filter executions to those that occurred after the given date. | 

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

