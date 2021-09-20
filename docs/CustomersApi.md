# \CustomersApi

All URIs are relative to *https://app.usemodernlogic.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerCustomerIdDelete**](CustomersApi.md#CustomerCustomerIdDelete) | **Delete** /customer/{customerId} | Delete Customer
[**CustomerCustomerIdGet**](CustomersApi.md#CustomerCustomerIdGet) | **Get** /customer/{customerId} | Get Customer Details
[**CustomerCustomerIdPut**](CustomersApi.md#CustomerCustomerIdPut) | **Put** /customer/{customerId} | Update Customer Details
[**CustomerGet**](CustomersApi.md#CustomerGet) | **Get** /customer | List Customers
[**CustomerPost**](CustomersApi.md#CustomerPost) | **Post** /customer | Create Customer



## CustomerCustomerIdDelete

> Customer CustomerCustomerIdDelete(ctx, customerId).Execute()

Delete Customer

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CustomerCustomerIdDelete(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CustomerCustomerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCustomerIdDelete`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CustomerCustomerIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Customer id that the client supplied | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCustomerIdGet

> Customer CustomerCustomerIdGet(ctx, customerId).Execute()

Get Customer Details

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CustomerCustomerIdGet(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CustomerCustomerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCustomerIdGet`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CustomerCustomerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Customer id that the client supplied | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCustomerIdPut

> Customer CustomerCustomerIdPut(ctx, customerId).Customer(customer).Execute()

Update Customer Details

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
    customer := *openapiclient.NewCustomer() // Customer | Customer Information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CustomerCustomerIdPut(context.Background(), customerId).Customer(customer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CustomerCustomerIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCustomerIdPut`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CustomerCustomerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Customer id that the client supplied | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCustomerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customer** | [**Customer**](Customer.md) | Customer Information | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGet

> ListResponse CustomerGet(ctx).PageSize(pageSize).PageNumber(pageNumber).Execute()

List Customers

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
    resp, r, err := api_client.CustomersApi.CustomerGet(context.Background()).PageSize(pageSize).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CustomerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerGet`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CustomerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetRequest struct via the builder pattern


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


## CustomerPost

> Customer CustomerPost(ctx).Customer(customer).Execute()

Create Customer

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
    customer := *openapiclient.NewCustomer() // Customer | Customer Information

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CustomerPost(context.Background()).Customer(customer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CustomerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerPost`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CustomerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | [**Customer**](Customer.md) | Customer Information | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

