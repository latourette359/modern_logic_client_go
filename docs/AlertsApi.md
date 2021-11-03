# \AlertsApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertAlertIdGet**](AlertsApi.md#AlertAlertIdGet) | **Get** /alert/{alertId} | Get Alert Details
[**AlertGet**](AlertsApi.md#AlertGet) | **Get** /alert | List Alerts



## AlertAlertIdGet

> Alert AlertAlertIdGet(ctx, alertId).Execute()

Get Alert Details

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
    alertId := int32(56) // int32 | Alert id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AlertAlertIdGet(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AlertAlertIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertAlertIdGet`: Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AlertAlertIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | Alert id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertAlertIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Alert**](Alert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertGet

> ListResponse AlertGet(ctx).PageSize(pageSize).PageNumber(pageNumber).AssignedTo(assignedTo).AssignedToTeam(assignedToTeam).Workflow(workflow).Before(before).After(after).Execute()

List Alerts

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
    assignedTo := "assignedTo_example" // string | The username of the user the alert is assigned to (optional)
    assignedToTeam := "assignedToTeam_example" // string | The name of the team the alert is assigned to (optional)
    workflow := "workflow_example" // string | The name of the workflow associated to the alert (optional)
    before := time.Now() // string | Filter alerts to those that occurred before the given date. (optional)
    after := time.Now() // string | Filter alerts to those that occurred after the given date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AlertGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).AssignedTo(assignedTo).AssignedToTeam(assignedToTeam).Workflow(workflow).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AlertGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AlertGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of elements to return (default is 10) | 
 **pageNumber** | **int32** | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero | 
 **assignedTo** | **string** | The username of the user the alert is assigned to | 
 **assignedToTeam** | **string** | The name of the team the alert is assigned to | 
 **workflow** | **string** | The name of the workflow associated to the alert | 
 **before** | **string** | Filter alerts to those that occurred before the given date. | 
 **after** | **string** | Filter alerts to those that occurred after the given date. | 

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

