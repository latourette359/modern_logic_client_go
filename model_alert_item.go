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

// AlertItem struct for AlertItem
type AlertItem struct {
	Id *int32 `json:"id,omitempty"`
	ExecutionId *int32 `json:"executionId,omitempty"`
}

// NewAlertItem instantiates a new AlertItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertItem() *AlertItem {
	this := AlertItem{}
	return &this
}

// NewAlertItemWithDefaults instantiates a new AlertItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertItemWithDefaults() *AlertItem {
	this := AlertItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertItem) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertItem) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertItem) SetId(v int32) {
	o.Id = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *AlertItem) GetExecutionId() int32 {
	if o == nil || o.ExecutionId == nil {
		var ret int32
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertItem) GetExecutionIdOk() (*int32, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AlertItem) HasExecutionId() bool {
	if o != nil && o.ExecutionId != nil {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given int32 and assigns it to the ExecutionId field.
func (o *AlertItem) SetExecutionId(v int32) {
	o.ExecutionId = &v
}

func (o AlertItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExecutionId != nil {
		toSerialize["executionId"] = o.ExecutionId
	}
	return json.Marshal(toSerialize)
}

type NullableAlertItem struct {
	value *AlertItem
	isSet bool
}

func (v NullableAlertItem) Get() *AlertItem {
	return v.value
}

func (v *NullableAlertItem) Set(val *AlertItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertItem(val *AlertItem) *NullableAlertItem {
	return &NullableAlertItem{value: val, isSet: true}
}

func (v NullableAlertItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


