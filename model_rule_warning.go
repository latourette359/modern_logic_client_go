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

// RuleWarning struct for RuleWarning
type RuleWarning struct {
	Name *string `json:"name,omitempty"`
	Warning *string `json:"warning,omitempty"`
}

// NewRuleWarning instantiates a new RuleWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleWarning() *RuleWarning {
	this := RuleWarning{}
	return &this
}

// NewRuleWarningWithDefaults instantiates a new RuleWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWarningWithDefaults() *RuleWarning {
	this := RuleWarning{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleWarning) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWarning) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleWarning) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleWarning) SetName(v string) {
	o.Name = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *RuleWarning) GetWarning() string {
	if o == nil || o.Warning == nil {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWarning) GetWarningOk() (*string, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *RuleWarning) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *RuleWarning) SetWarning(v string) {
	o.Warning = &v
}

func (o RuleWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Warning != nil {
		toSerialize["warning"] = o.Warning
	}
	return json.Marshal(toSerialize)
}

type NullableRuleWarning struct {
	value *RuleWarning
	isSet bool
}

func (v NullableRuleWarning) Get() *RuleWarning {
	return v.value
}

func (v *NullableRuleWarning) Set(val *RuleWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleWarning(val *RuleWarning) *NullableRuleWarning {
	return &NullableRuleWarning{value: val, isSet: true}
}

func (v NullableRuleWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


