/*
Modern Logic Api

Manage and version your customer decision logic outside of your codebase

API version: 1.0.0
Contact: info@usemodernlogic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modern_logic_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ExecutionsApiService ExecutionsApi service
type ExecutionsApiService service

type ApiCustomerCustomerIdExecutionsGetRequest struct {
	ctx _context.Context
	ApiService *ExecutionsApiService
	customerId string
	pageSize *int32
	pageNumber *int32
}

// Number of elements to return (default is 10)
func (r ApiCustomerCustomerIdExecutionsGetRequest) PageSize(pageSize int32) ApiCustomerCustomerIdExecutionsGetRequest {
	r.pageSize = &pageSize
	return r
}
// Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero
func (r ApiCustomerCustomerIdExecutionsGetRequest) PageNumber(pageNumber int32) ApiCustomerCustomerIdExecutionsGetRequest {
	r.pageNumber = &pageNumber
	return r
}

func (r ApiCustomerCustomerIdExecutionsGetRequest) Execute() (ListResponse, *_nethttp.Response, error) {
	return r.ApiService.CustomerCustomerIdExecutionsGetExecute(r)
}

/*
CustomerCustomerIdExecutionsGet List Customer Executions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId Customer id that the client supplied
 @return ApiCustomerCustomerIdExecutionsGetRequest
*/
func (a *ExecutionsApiService) CustomerCustomerIdExecutionsGet(ctx _context.Context, customerId string) ApiCustomerCustomerIdExecutionsGetRequest {
	return ApiCustomerCustomerIdExecutionsGetRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return ListResponse
func (a *ExecutionsApiService) CustomerCustomerIdExecutionsGetExecute(r ApiCustomerCustomerIdExecutionsGetRequest) (ListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsApiService.CustomerCustomerIdExecutionsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer/{customerId}/executions"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", _neturl.PathEscape(parameterToString(r.customerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiExecutionExecutionIdGetRequest struct {
	ctx _context.Context
	ApiService *ExecutionsApiService
	executionId int32
}


func (r ApiExecutionExecutionIdGetRequest) Execute() (Execution, *_nethttp.Response, error) {
	return r.ApiService.ExecutionExecutionIdGetExecute(r)
}

/*
ExecutionExecutionIdGet Get Execution Details

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param executionId Execution id
 @return ApiExecutionExecutionIdGetRequest
*/
func (a *ExecutionsApiService) ExecutionExecutionIdGet(ctx _context.Context, executionId int32) ApiExecutionExecutionIdGetRequest {
	return ApiExecutionExecutionIdGetRequest{
		ApiService: a,
		ctx: ctx,
		executionId: executionId,
	}
}

// Execute executes the request
//  @return Execution
func (a *ExecutionsApiService) ExecutionExecutionIdGetExecute(r ApiExecutionExecutionIdGetRequest) (Execution, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Execution
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsApiService.ExecutionExecutionIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/execution/{executionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"executionId"+"}", _neturl.PathEscape(parameterToString(r.executionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiExecutionExecutionIdResumePostRequest struct {
	ctx _context.Context
	ApiService *ExecutionsApiService
	executionId int32
	requestBody *map[string]interface{}
}

// Execution Information
func (r ApiExecutionExecutionIdResumePostRequest) RequestBody(requestBody map[string]interface{}) ApiExecutionExecutionIdResumePostRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiExecutionExecutionIdResumePostRequest) Execute() (WorkflowExecutionResult, *_nethttp.Response, error) {
	return r.ApiService.ExecutionExecutionIdResumePostExecute(r)
}

/*
ExecutionExecutionIdResumePost Resume Execution

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param executionId execution id
 @return ApiExecutionExecutionIdResumePostRequest
*/
func (a *ExecutionsApiService) ExecutionExecutionIdResumePost(ctx _context.Context, executionId int32) ApiExecutionExecutionIdResumePostRequest {
	return ApiExecutionExecutionIdResumePostRequest{
		ApiService: a,
		ctx: ctx,
		executionId: executionId,
	}
}

// Execute executes the request
//  @return WorkflowExecutionResult
func (a *ExecutionsApiService) ExecutionExecutionIdResumePostExecute(r ApiExecutionExecutionIdResumePostRequest) (WorkflowExecutionResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  WorkflowExecutionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsApiService.ExecutionExecutionIdResumePost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/execution/{executionId}/resume"
	localVarPath = strings.Replace(localVarPath, "{"+"executionId"+"}", _neturl.PathEscape(parameterToString(r.executionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiExecutionGetRequest struct {
	ctx _context.Context
	ApiService *ExecutionsApiService
	pageSize *int32
	pageNumber *int32
	alertType *string
	before *string
	after *string
}

// Number of elements to return (default is 10)
func (r ApiExecutionGetRequest) PageSize(pageSize int32) ApiExecutionGetRequest {
	r.pageSize = &pageSize
	return r
}
// Lists are ordered by creation date ascending. To return the first page, set pageNumber to zero
func (r ApiExecutionGetRequest) PageNumber(pageNumber int32) ApiExecutionGetRequest {
	r.pageNumber = &pageNumber
	return r
}
// The alert status of this execution
func (r ApiExecutionGetRequest) AlertType(alertType string) ApiExecutionGetRequest {
	r.alertType = &alertType
	return r
}
// Filter executions to those that occurred before the given date.
func (r ApiExecutionGetRequest) Before(before string) ApiExecutionGetRequest {
	r.before = &before
	return r
}
// Filter executions to those that occurred after the given date.
func (r ApiExecutionGetRequest) After(after string) ApiExecutionGetRequest {
	r.after = &after
	return r
}

func (r ApiExecutionGetRequest) Execute() (ListResponse, *_nethttp.Response, error) {
	return r.ApiService.ExecutionGetExecute(r)
}

/*
ExecutionGet List executions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExecutionGetRequest
*/
func (a *ExecutionsApiService) ExecutionGet(ctx _context.Context) ApiExecutionGetRequest {
	return ApiExecutionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListResponse
func (a *ExecutionsApiService) ExecutionGetExecute(r ApiExecutionGetRequest) (ListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsApiService.ExecutionGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/execution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.alertType != nil {
		localVarQueryParams.Add("alertType", parameterToString(*r.alertType, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
