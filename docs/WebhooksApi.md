# \WebhooksApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookGet**](WebhooksApi.md#WebhookGet) | **Get** /webhook | List Webhooks
[**WebhookWebhookIdGet**](WebhooksApi.md#WebhookWebhookIdGet) | **Get** /webhook/{webhookId} | Get Webhook Details



## WebhookGet

> ListResponse WebhookGet(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Webhooks

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
    resp, r, err := api_client.WebhooksApi.WebhookGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


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


## WebhookWebhookIdGet

> Webhook WebhookWebhookIdGet(ctx, webhookId).Execute()

Get Webhook Details

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
    webhookId := int32(56) // int32 | Webhook Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.WebhookWebhookIdGet(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookWebhookIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookWebhookIdGet`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookWebhookIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | Webhook Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookWebhookIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

