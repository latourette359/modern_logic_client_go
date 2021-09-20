# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AlertId** | Pointer to **int32** |  | [optional] 
**Runtime** | Pointer to **int32** | Time it took to execute the workflow in milliseconds | [optional] 
**DateTimeExecuted** | Pointer to **time.Time** | When the execution happened | [optional] 
**WorkflowId** | Pointer to **int32** |  | [optional] 
**Request** | Pointer to [**ExecutionInput**](ExecutionInput.md) |  | [optional] 
**Response** | Pointer to [**WorkflowExecutionResult**](WorkflowExecutionResult.md) |  | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Execution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Execution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Execution) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Execution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlertId

`func (o *Execution) GetAlertId() int32`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *Execution) GetAlertIdOk() (*int32, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *Execution) SetAlertId(v int32)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *Execution) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetRuntime

`func (o *Execution) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Execution) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Execution) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Execution) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetDateTimeExecuted

`func (o *Execution) GetDateTimeExecuted() time.Time`

GetDateTimeExecuted returns the DateTimeExecuted field if non-nil, zero value otherwise.

### GetDateTimeExecutedOk

`func (o *Execution) GetDateTimeExecutedOk() (*time.Time, bool)`

GetDateTimeExecutedOk returns a tuple with the DateTimeExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeExecuted

`func (o *Execution) SetDateTimeExecuted(v time.Time)`

SetDateTimeExecuted sets DateTimeExecuted field to given value.

### HasDateTimeExecuted

`func (o *Execution) HasDateTimeExecuted() bool`

HasDateTimeExecuted returns a boolean if a field has been set.

### GetWorkflowId

`func (o *Execution) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Execution) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Execution) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *Execution) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetRequest

`func (o *Execution) GetRequest() ExecutionInput`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Execution) GetRequestOk() (*ExecutionInput, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Execution) SetRequest(v ExecutionInput)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Execution) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *Execution) GetResponse() WorkflowExecutionResult`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Execution) GetResponseOk() (*WorkflowExecutionResult, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Execution) SetResponse(v WorkflowExecutionResult)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Execution) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


