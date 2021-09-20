# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**DateOpened** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExecutionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alert) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignedTo

`func (o *Alert) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Alert) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Alert) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *Alert) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetDateOpened

`func (o *Alert) GetDateOpened() string`

GetDateOpened returns the DateOpened field if non-nil, zero value otherwise.

### GetDateOpenedOk

`func (o *Alert) GetDateOpenedOk() (*string, bool)`

GetDateOpenedOk returns a tuple with the DateOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOpened

`func (o *Alert) SetDateOpened(v string)`

SetDateOpened sets DateOpened field to given value.

### HasDateOpened

`func (o *Alert) HasDateOpened() bool`

HasDateOpened returns a boolean if a field has been set.

### GetDescription

`func (o *Alert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionId

`func (o *Alert) GetExecutionId() int32`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Alert) GetExecutionIdOk() (*int32, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Alert) SetExecutionId(v int32)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Alert) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


