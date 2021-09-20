# WorkflowExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to [**[]RuleSetWarning**](RuleSetWarning.md) |  | [optional] 
**ExecutionId** | Pointer to **int32** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **[]string** |  | [optional] 
**DateEvaluated** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **int32** |  | [optional] 
**WorkflowVersion** | Pointer to **int32** |  | [optional] 
**DefinedVariables** | Pointer to **map[string]map[string]interface{}** | this field contains the variables that were defined during your execution | [optional] 
**FollowUpRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkflowExecutionResult

`func NewWorkflowExecutionResult() *WorkflowExecutionResult`

NewWorkflowExecutionResult instantiates a new WorkflowExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionResultWithDefaults

`func NewWorkflowExecutionResultWithDefaults() *WorkflowExecutionResult`

NewWorkflowExecutionResultWithDefaults instantiates a new WorkflowExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *WorkflowExecutionResult) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *WorkflowExecutionResult) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *WorkflowExecutionResult) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *WorkflowExecutionResult) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetWarnings

`func (o *WorkflowExecutionResult) GetWarnings() []RuleSetWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *WorkflowExecutionResult) GetWarningsOk() (*[]RuleSetWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *WorkflowExecutionResult) SetWarnings(v []RuleSetWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *WorkflowExecutionResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetExecutionId

`func (o *WorkflowExecutionResult) GetExecutionId() int32`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionResult) GetExecutionIdOk() (*int32, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionResult) SetExecutionId(v int32)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *WorkflowExecutionResult) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetResult

`func (o *WorkflowExecutionResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *WorkflowExecutionResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScore

`func (o *WorkflowExecutionResult) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *WorkflowExecutionResult) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *WorkflowExecutionResult) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *WorkflowExecutionResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetReason

`func (o *WorkflowExecutionResult) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowExecutionResult) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowExecutionResult) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkflowExecutionResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDateEvaluated

`func (o *WorkflowExecutionResult) GetDateEvaluated() string`

GetDateEvaluated returns the DateEvaluated field if non-nil, zero value otherwise.

### GetDateEvaluatedOk

`func (o *WorkflowExecutionResult) GetDateEvaluatedOk() (*string, bool)`

GetDateEvaluatedOk returns a tuple with the DateEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEvaluated

`func (o *WorkflowExecutionResult) SetDateEvaluated(v string)`

SetDateEvaluated sets DateEvaluated field to given value.

### HasDateEvaluated

`func (o *WorkflowExecutionResult) HasDateEvaluated() bool`

HasDateEvaluated returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowExecutionResult) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowExecutionResult) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowExecutionResult) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowExecutionResult) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *WorkflowExecutionResult) GetWorkflowVersion() int32`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *WorkflowExecutionResult) GetWorkflowVersionOk() (*int32, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *WorkflowExecutionResult) SetWorkflowVersion(v int32)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *WorkflowExecutionResult) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.

### GetDefinedVariables

`func (o *WorkflowExecutionResult) GetDefinedVariables() map[string]map[string]interface{}`

GetDefinedVariables returns the DefinedVariables field if non-nil, zero value otherwise.

### GetDefinedVariablesOk

`func (o *WorkflowExecutionResult) GetDefinedVariablesOk() (*map[string]map[string]interface{}, bool)`

GetDefinedVariablesOk returns a tuple with the DefinedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinedVariables

`func (o *WorkflowExecutionResult) SetDefinedVariables(v map[string]map[string]interface{})`

SetDefinedVariables sets DefinedVariables field to given value.

### HasDefinedVariables

`func (o *WorkflowExecutionResult) HasDefinedVariables() bool`

HasDefinedVariables returns a boolean if a field has been set.

### GetFollowUpRequired

`func (o *WorkflowExecutionResult) GetFollowUpRequired() bool`

GetFollowUpRequired returns the FollowUpRequired field if non-nil, zero value otherwise.

### GetFollowUpRequiredOk

`func (o *WorkflowExecutionResult) GetFollowUpRequiredOk() (*bool, bool)`

GetFollowUpRequiredOk returns a tuple with the FollowUpRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpRequired

`func (o *WorkflowExecutionResult) SetFollowUpRequired(v bool)`

SetFollowUpRequired sets FollowUpRequired field to given value.

### HasFollowUpRequired

`func (o *WorkflowExecutionResult) HasFollowUpRequired() bool`

HasFollowUpRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


