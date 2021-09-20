# \DataSourcesApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatasourceCallsDatasourceCallIdGet**](DataSourcesApi.md#DatasourceCallsDatasourceCallIdGet) | **Get** /datasource_calls/{datasourceCallId} | Get Data Source Call Details
[**DatasourceDatasourceIdCallsGet**](DataSourcesApi.md#DatasourceDatasourceIdCallsGet) | **Get** /datasource/{datasourceId}/calls | List Data Source Calls
[**DatasourceDatasourceIdGet**](DataSourcesApi.md#DatasourceDatasourceIdGet) | **Get** /datasource/{datasourceId} | Get Data Source Details
[**DatasourceGet**](DataSourcesApi.md#DatasourceGet) | **Get** /datasource | List Data Sources



## DatasourceCallsDatasourceCallIdGet

> DataSourceCall DatasourceCallsDatasourceCallIdGet(ctx, datasourceCallId).Execute()

Get Data Source Call Details



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
    datasourceCallId := int32(56) // int32 | Data source call id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DatasourceCallsDatasourceCallIdGet(context.Background(), datasourceCallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DatasourceCallsDatasourceCallIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasourceCallsDatasourceCallIdGet`: DataSourceCall
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DatasourceCallsDatasourceCallIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceCallId** | **int32** | Data source call id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceCallsDatasourceCallIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSourceCall**](DataSourceCall.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceDatasourceIdCallsGet

> ListResponse DatasourceDatasourceIdCallsGet(ctx, datasourceId).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Data Source Calls

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
    datasourceId := int32(56) // int32 | Specified workflow to run, to evaluate the target user.
    pageSize := int32(56) // int32 | Number of elements to return (default is 10) (optional)
    pageNumber := int32(56) // int32 | Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DatasourceDatasourceIdCallsGet(context.Background(), datasourceId).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DatasourceDatasourceIdCallsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasourceDatasourceIdCallsGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DatasourceDatasourceIdCallsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **int32** | Specified workflow to run, to evaluate the target user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceDatasourceIdCallsGetRequest struct via the builder pattern


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


## DatasourceDatasourceIdGet

> DataSource DatasourceDatasourceIdGet(ctx, datasourceId).Execute()

Get Data Source Details



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
    datasourceId := int32(56) // int32 | Data Source id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DatasourceDatasourceIdGet(context.Background(), datasourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DatasourceDatasourceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasourceDatasourceIdGet`: DataSource
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DatasourceDatasourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datasourceId** | **int32** | Data Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceDatasourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSource**](DataSource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceGet

> ListResponse DatasourceGet(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Data Sources

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
    resp, r, err := api_client.DataSourcesApi.DatasourceGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DatasourceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatasourceGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DatasourceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceGetRequest struct via the builder pattern


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

