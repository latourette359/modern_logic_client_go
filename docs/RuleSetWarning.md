# RuleSetWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]RuleWarning**](RuleWarning.md) |  | [optional] 

## Methods

### NewRuleSetWarning

`func NewRuleSetWarning() *RuleSetWarning`

NewRuleSetWarning instantiates a new RuleSetWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetWarningWithDefaults

`func NewRuleSetWarningWithDefaults() *RuleSetWarning`

NewRuleSetWarningWithDefaults instantiates a new RuleSetWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleSetWarning) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSetWarning) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSetWarning) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleSetWarning) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *RuleSetWarning) GetRules() []RuleWarning`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RuleSetWarning) GetRulesOk() (*[]RuleWarning, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RuleSetWarning) SetRules(v []RuleWarning)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RuleSetWarning) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


