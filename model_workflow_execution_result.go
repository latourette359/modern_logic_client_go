/*
Modern Logic Api

Manage and version your customer decision logic outside of your codebase

API version: 1.0.0
Contact: info@usemodernlogic.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modern_logic_client

import (
	"encoding/json"
)

// WorkflowExecutionResult Contains the results of the executed workflow
type WorkflowExecutionResult struct {
	Flags *[]string `json:"flags,omitempty"`
	Warnings *[]RuleSetWarning `json:"warnings,omitempty"`
	ExecutionId *int32 `json:"executionId,omitempty"`
	Result *string `json:"result,omitempty"`
	Score *string `json:"score,omitempty"`
	Reason *[]string `json:"reason,omitempty"`
	DateEvaluated *string `json:"dateEvaluated,omitempty"`
	WorkflowId *int32 `json:"workflowId,omitempty"`
	WorkflowVersion *int32 `json:"workflowVersion,omitempty"`
	// this field contains the variables that were defined during your execution
	DefinedVariables *map[string]interface{} `json:"definedVariables,omitempty"`
	FollowUpRequired *bool `json:"followUpRequired,omitempty"`
}

// NewWorkflowExecutionResult instantiates a new WorkflowExecutionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowExecutionResult() *WorkflowExecutionResult {
	this := WorkflowExecutionResult{}
	return &this
}

// NewWorkflowExecutionResultWithDefaults instantiates a new WorkflowExecutionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowExecutionResultWithDefaults() *WorkflowExecutionResult {
	this := WorkflowExecutionResult{}
	return &this
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetFlags() []string {
	if o == nil || o.Flags == nil {
		var ret []string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetFlagsOk() (*[]string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *WorkflowExecutionResult) SetFlags(v []string) {
	o.Flags = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetWarnings() []RuleSetWarning {
	if o == nil || o.Warnings == nil {
		var ret []RuleSetWarning
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetWarningsOk() (*[]RuleSetWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []RuleSetWarning and assigns it to the Warnings field.
func (o *WorkflowExecutionResult) SetWarnings(v []RuleSetWarning) {
	o.Warnings = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetExecutionId() int32 {
	if o == nil || o.ExecutionId == nil {
		var ret int32
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetExecutionIdOk() (*int32, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasExecutionId() bool {
	if o != nil && o.ExecutionId != nil {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given int32 and assigns it to the ExecutionId field.
func (o *WorkflowExecutionResult) SetExecutionId(v int32) {
	o.ExecutionId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *WorkflowExecutionResult) SetResult(v string) {
	o.Result = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetScore() string {
	if o == nil || o.Score == nil {
		var ret string
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetScoreOk() (*string, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given string and assigns it to the Score field.
func (o *WorkflowExecutionResult) SetScore(v string) {
	o.Score = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetReason() []string {
	if o == nil || o.Reason == nil {
		var ret []string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetReasonOk() (*[]string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given []string and assigns it to the Reason field.
func (o *WorkflowExecutionResult) SetReason(v []string) {
	o.Reason = &v
}

// GetDateEvaluated returns the DateEvaluated field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetDateEvaluated() string {
	if o == nil || o.DateEvaluated == nil {
		var ret string
		return ret
	}
	return *o.DateEvaluated
}

// GetDateEvaluatedOk returns a tuple with the DateEvaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetDateEvaluatedOk() (*string, bool) {
	if o == nil || o.DateEvaluated == nil {
		return nil, false
	}
	return o.DateEvaluated, true
}

// HasDateEvaluated returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasDateEvaluated() bool {
	if o != nil && o.DateEvaluated != nil {
		return true
	}

	return false
}

// SetDateEvaluated gets a reference to the given string and assigns it to the DateEvaluated field.
func (o *WorkflowExecutionResult) SetDateEvaluated(v string) {
	o.DateEvaluated = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetWorkflowId() int32 {
	if o == nil || o.WorkflowId == nil {
		var ret int32
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetWorkflowIdOk() (*int32, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given int32 and assigns it to the WorkflowId field.
func (o *WorkflowExecutionResult) SetWorkflowId(v int32) {
	o.WorkflowId = &v
}

// GetWorkflowVersion returns the WorkflowVersion field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetWorkflowVersion() int32 {
	if o == nil || o.WorkflowVersion == nil {
		var ret int32
		return ret
	}
	return *o.WorkflowVersion
}

// GetWorkflowVersionOk returns a tuple with the WorkflowVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetWorkflowVersionOk() (*int32, bool) {
	if o == nil || o.WorkflowVersion == nil {
		return nil, false
	}
	return o.WorkflowVersion, true
}

// HasWorkflowVersion returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasWorkflowVersion() bool {
	if o != nil && o.WorkflowVersion != nil {
		return true
	}

	return false
}

// SetWorkflowVersion gets a reference to the given int32 and assigns it to the WorkflowVersion field.
func (o *WorkflowExecutionResult) SetWorkflowVersion(v int32) {
	o.WorkflowVersion = &v
}

// GetDefinedVariables returns the DefinedVariables field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetDefinedVariables() map[string]interface{} {
	if o == nil || o.DefinedVariables == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.DefinedVariables
}

// GetDefinedVariablesOk returns a tuple with the DefinedVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetDefinedVariablesOk() (*map[string]interface{}, bool) {
	if o == nil || o.DefinedVariables == nil {
		return nil, false
	}
	return o.DefinedVariables, true
}

// HasDefinedVariables returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasDefinedVariables() bool {
	if o != nil && o.DefinedVariables != nil {
		return true
	}

	return false
}

// SetDefinedVariables gets a reference to the given map[string]interface{} and assigns it to the DefinedVariables field.
func (o *WorkflowExecutionResult) SetDefinedVariables(v map[string]interface{}) {
	o.DefinedVariables = &v
}

// GetFollowUpRequired returns the FollowUpRequired field value if set, zero value otherwise.
func (o *WorkflowExecutionResult) GetFollowUpRequired() bool {
	if o == nil || o.FollowUpRequired == nil {
		var ret bool
		return ret
	}
	return *o.FollowUpRequired
}

// GetFollowUpRequiredOk returns a tuple with the FollowUpRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionResult) GetFollowUpRequiredOk() (*bool, bool) {
	if o == nil || o.FollowUpRequired == nil {
		return nil, false
	}
	return o.FollowUpRequired, true
}

// HasFollowUpRequired returns a boolean if a field has been set.
func (o *WorkflowExecutionResult) HasFollowUpRequired() bool {
	if o != nil && o.FollowUpRequired != nil {
		return true
	}

	return false
}

// SetFollowUpRequired gets a reference to the given bool and assigns it to the FollowUpRequired field.
func (o *WorkflowExecutionResult) SetFollowUpRequired(v bool) {
	o.FollowUpRequired = &v
}

func (o WorkflowExecutionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.ExecutionId != nil {
		toSerialize["executionId"] = o.ExecutionId
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.DateEvaluated != nil {
		toSerialize["dateEvaluated"] = o.DateEvaluated
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.WorkflowVersion != nil {
		toSerialize["workflowVersion"] = o.WorkflowVersion
	}
	if o.DefinedVariables != nil {
		toSerialize["definedVariables"] = o.DefinedVariables
	}
	if o.FollowUpRequired != nil {
		toSerialize["followUpRequired"] = o.FollowUpRequired
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowExecutionResult struct {
	value *WorkflowExecutionResult
	isSet bool
}

func (v NullableWorkflowExecutionResult) Get() *WorkflowExecutionResult {
	return v.value
}

func (v *NullableWorkflowExecutionResult) Set(val *WorkflowExecutionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowExecutionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowExecutionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowExecutionResult(val *WorkflowExecutionResult) *NullableWorkflowExecutionResult {
	return &NullableWorkflowExecutionResult{value: val, isSet: true}
}

func (v NullableWorkflowExecutionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowExecutionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


