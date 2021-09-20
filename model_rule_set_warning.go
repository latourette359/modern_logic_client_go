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

// RuleSetWarning struct for RuleSetWarning
type RuleSetWarning struct {
	Name *string `json:"name,omitempty"`
	Rules *[]RuleWarning `json:"rules,omitempty"`
}

// NewRuleSetWarning instantiates a new RuleSetWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSetWarning() *RuleSetWarning {
	this := RuleSetWarning{}
	return &this
}

// NewRuleSetWarningWithDefaults instantiates a new RuleSetWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetWarningWithDefaults() *RuleSetWarning {
	this := RuleSetWarning{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleSetWarning) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetWarning) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleSetWarning) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleSetWarning) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *RuleSetWarning) GetRules() []RuleWarning {
	if o == nil || o.Rules == nil {
		var ret []RuleWarning
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetWarning) GetRulesOk() (*[]RuleWarning, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *RuleSetWarning) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []RuleWarning and assigns it to the Rules field.
func (o *RuleSetWarning) SetRules(v []RuleWarning) {
	o.Rules = &v
}

func (o RuleSetWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableRuleSetWarning struct {
	value *RuleSetWarning
	isSet bool
}

func (v NullableRuleSetWarning) Get() *RuleSetWarning {
	return v.value
}

func (v *NullableRuleSetWarning) Set(val *RuleSetWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSetWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSetWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSetWarning(val *RuleSetWarning) *NullableRuleSetWarning {
	return &NullableRuleSetWarning{value: val, isSet: true}
}

func (v NullableRuleSetWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSetWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

