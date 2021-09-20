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
	"fmt"
)

// DataSourceType the model 'DataSourceType'
type DataSourceType string

// List of DataSourceType
const (
	PG_DB DataSourceType = "PG_DB"
	API_POST DataSourceType = "API_POST"
	MANAGED DataSourceType = "MANAGED"
)

var allowedDataSourceTypeEnumValues = []DataSourceType{
	"PG_DB",
	"API_POST",
	"MANAGED",
}

func (v *DataSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSourceType(value)
	for _, existing := range allowedDataSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSourceType", value)
}

// NewDataSourceTypeFromValue returns a pointer to a valid DataSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourceTypeFromValue(v string) (*DataSourceType, error) {
	ev := DataSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSourceType: valid values are %v", v, allowedDataSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSourceType) IsValid() bool {
	for _, existing := range allowedDataSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSourceType value
func (v DataSourceType) Ptr() *DataSourceType {
	return &v
}

type NullableDataSourceType struct {
	value *DataSourceType
	isSet bool
}

func (v NullableDataSourceType) Get() *DataSourceType {
	return v.value
}

func (v *NullableDataSourceType) Set(val *DataSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceType(val *DataSourceType) *NullableDataSourceType {
	return &NullableDataSourceType{value: val, isSet: true}
}

func (v NullableDataSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

