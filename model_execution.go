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
	"time"
)

// Execution struct for Execution
type Execution struct {
	Id *int32 `json:"id,omitempty"`
	AlertId *int32 `json:"alertId,omitempty"`
	// Time it took to execute the workflow in milliseconds
	Runtime *int32 `json:"runtime,omitempty"`
	// When the execution happened
	DateTimeExecuted *time.Time `json:"dateTimeExecuted,omitempty"`
	WorkflowId *int32 `json:"workflowId,omitempty"`
	Request *ExecutionInput `json:"request,omitempty"`
	Response *WorkflowExecutionResult `json:"response,omitempty"`
}

// NewExecution instantiates a new Execution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecution() *Execution {
	this := Execution{}
	return &this
}

// NewExecutionWithDefaults instantiates a new Execution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionWithDefaults() *Execution {
	this := Execution{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Execution) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Execution) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Execution) SetId(v int32) {
	o.Id = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *Execution) GetAlertId() int32 {
	if o == nil || o.AlertId == nil {
		var ret int32
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetAlertIdOk() (*int32, bool) {
	if o == nil || o.AlertId == nil {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *Execution) HasAlertId() bool {
	if o != nil && o.AlertId != nil {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given int32 and assigns it to the AlertId field.
func (o *Execution) SetAlertId(v int32) {
	o.AlertId = &v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *Execution) GetRuntime() int32 {
	if o == nil || o.Runtime == nil {
		var ret int32
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetRuntimeOk() (*int32, bool) {
	if o == nil || o.Runtime == nil {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *Execution) HasRuntime() bool {
	if o != nil && o.Runtime != nil {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given int32 and assigns it to the Runtime field.
func (o *Execution) SetRuntime(v int32) {
	o.Runtime = &v
}

// GetDateTimeExecuted returns the DateTimeExecuted field value if set, zero value otherwise.
func (o *Execution) GetDateTimeExecuted() time.Time {
	if o == nil || o.DateTimeExecuted == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTimeExecuted
}

// GetDateTimeExecutedOk returns a tuple with the DateTimeExecuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetDateTimeExecutedOk() (*time.Time, bool) {
	if o == nil || o.DateTimeExecuted == nil {
		return nil, false
	}
	return o.DateTimeExecuted, true
}

// HasDateTimeExecuted returns a boolean if a field has been set.
func (o *Execution) HasDateTimeExecuted() bool {
	if o != nil && o.DateTimeExecuted != nil {
		return true
	}

	return false
}

// SetDateTimeExecuted gets a reference to the given time.Time and assigns it to the DateTimeExecuted field.
func (o *Execution) SetDateTimeExecuted(v time.Time) {
	o.DateTimeExecuted = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *Execution) GetWorkflowId() int32 {
	if o == nil || o.WorkflowId == nil {
		var ret int32
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetWorkflowIdOk() (*int32, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *Execution) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given int32 and assigns it to the WorkflowId field.
func (o *Execution) SetWorkflowId(v int32) {
	o.WorkflowId = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *Execution) GetRequest() ExecutionInput {
	if o == nil || o.Request == nil {
		var ret ExecutionInput
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetRequestOk() (*ExecutionInput, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *Execution) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given ExecutionInput and assigns it to the Request field.
func (o *Execution) SetRequest(v ExecutionInput) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *Execution) GetResponse() WorkflowExecutionResult {
	if o == nil || o.Response == nil {
		var ret WorkflowExecutionResult
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Execution) GetResponseOk() (*WorkflowExecutionResult, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *Execution) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given WorkflowExecutionResult and assigns it to the Response field.
func (o *Execution) SetResponse(v WorkflowExecutionResult) {
	o.Response = &v
}

func (o Execution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AlertId != nil {
		toSerialize["alertId"] = o.AlertId
	}
	if o.Runtime != nil {
		toSerialize["runtime"] = o.Runtime
	}
	if o.DateTimeExecuted != nil {
		toSerialize["dateTimeExecuted"] = o.DateTimeExecuted
	}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	return json.Marshal(toSerialize)
}

type NullableExecution struct {
	value *Execution
	isSet bool
}

func (v NullableExecution) Get() *Execution {
	return v.value
}

func (v *NullableExecution) Set(val *Execution) {
	v.value = val
	v.isSet = true
}

func (v NullableExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecution(val *Execution) *NullableExecution {
	return &NullableExecution{value: val, isSet: true}
}

func (v NullableExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


