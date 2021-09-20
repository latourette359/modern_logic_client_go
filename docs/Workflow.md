# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CurrentProductionVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow() *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workflow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Workflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrentProductionVersion

`func (o *Workflow) GetCurrentProductionVersion() int32`

GetCurrentProductionVersion returns the CurrentProductionVersion field if non-nil, zero value otherwise.

### GetCurrentProductionVersionOk

`func (o *Workflow) GetCurrentProductionVersionOk() (*int32, bool)`

GetCurrentProductionVersionOk returns a tuple with the CurrentProductionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProductionVersion

`func (o *Workflow) SetCurrentProductionVersion(v int32)`

SetCurrentProductionVersion sets CurrentProductionVersion field to given value.

### HasCurrentProductionVersion

`func (o *Workflow) HasCurrentProductionVersion() bool`

HasCurrentProductionVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


