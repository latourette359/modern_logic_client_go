# ExecutionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DateTimeExecuted** | Pointer to **time.Time** | When the execution happened | [optional] 
**WorkflowId** | Pointer to **int32** |  | [optional] 

## Methods

### NewExecutionItem

`func NewExecutionItem() *ExecutionItem`

NewExecutionItem instantiates a new ExecutionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionItemWithDefaults

`func NewExecutionItemWithDefaults() *ExecutionItem`

NewExecutionItemWithDefaults instantiates a new ExecutionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateTimeExecuted

`func (o *ExecutionItem) GetDateTimeExecuted() time.Time`

GetDateTimeExecuted returns the DateTimeExecuted field if non-nil, zero value otherwise.

### GetDateTimeExecutedOk

`func (o *ExecutionItem) GetDateTimeExecutedOk() (*time.Time, bool)`

GetDateTimeExecutedOk returns a tuple with the DateTimeExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeExecuted

`func (o *ExecutionItem) SetDateTimeExecuted(v time.Time)`

SetDateTimeExecuted sets DateTimeExecuted field to given value.

### HasDateTimeExecuted

`func (o *ExecutionItem) HasDateTimeExecuted() bool`

HasDateTimeExecuted returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ExecutionItem) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ExecutionItem) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ExecutionItem) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ExecutionItem) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


