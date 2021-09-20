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

// WorkflowVersion struct for WorkflowVersion
type WorkflowVersion struct {
	VersionNumber *int32 `json:"versionNumber,omitempty"`
	Name *string `json:"name,omitempty"`
	PublishDate *string `json:"publishDate,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
}

// NewWorkflowVersion instantiates a new WorkflowVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowVersion() *WorkflowVersion {
	this := WorkflowVersion{}
	return &this
}

// NewWorkflowVersionWithDefaults instantiates a new WorkflowVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowVersionWithDefaults() *WorkflowVersion {
	this := WorkflowVersion{}
	return &this
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *WorkflowVersion) GetVersionNumber() int32 {
	if o == nil || o.VersionNumber == nil {
		var ret int32
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVersion) GetVersionNumberOk() (*int32, bool) {
	if o == nil || o.VersionNumber == nil {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *WorkflowVersion) HasVersionNumber() bool {
	if o != nil && o.VersionNumber != nil {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given int32 and assigns it to the VersionNumber field.
func (o *WorkflowVersion) SetVersionNumber(v int32) {
	o.VersionNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVersion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowVersion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowVersion) SetName(v string) {
	o.Name = &v
}

// GetPublishDate returns the PublishDate field value if set, zero value otherwise.
func (o *WorkflowVersion) GetPublishDate() string {
	if o == nil || o.PublishDate == nil {
		var ret string
		return ret
	}
	return *o.PublishDate
}

// GetPublishDateOk returns a tuple with the PublishDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVersion) GetPublishDateOk() (*string, bool) {
	if o == nil || o.PublishDate == nil {
		return nil, false
	}
	return o.PublishDate, true
}

// HasPublishDate returns a boolean if a field has been set.
func (o *WorkflowVersion) HasPublishDate() bool {
	if o != nil && o.PublishDate != nil {
		return true
	}

	return false
}

// SetPublishDate gets a reference to the given string and assigns it to the PublishDate field.
func (o *WorkflowVersion) SetPublishDate(v string) {
	o.PublishDate = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *WorkflowVersion) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowVersion) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *WorkflowVersion) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *WorkflowVersion) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o WorkflowVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VersionNumber != nil {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublishDate != nil {
		toSerialize["publishDate"] = o.PublishDate
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowVersion struct {
	value *WorkflowVersion
	isSet bool
}

func (v NullableWorkflowVersion) Get() *WorkflowVersion {
	return v.value
}

func (v *NullableWorkflowVersion) Set(val *WorkflowVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowVersion(val *WorkflowVersion) *NullableWorkflowVersion {
	return &NullableWorkflowVersion{value: val, isSet: true}
}

func (v NullableWorkflowVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

