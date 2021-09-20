# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **int32** |  | [optional] 
**WorkflowId** | Pointer to **int32** |  | [optional] 
**WorkflowVersion** | Pointer to **int32** |  | [optional] 
**ExecutionId** | Pointer to **int32** |  | [optional] 
**CustomerId** | Pointer to **int32** |  | [optional] 
**CustomerUrl** | Pointer to **string** |  | [optional] 
**AlertId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**EventMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse() *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookResponse) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookResponse) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookResponse) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookResponse) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WebhookResponse) GetWorkflowId() int32`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WebhookResponse) GetWorkflowIdOk() (*int32, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WebhookResponse) SetWorkflowId(v int32)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WebhookResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *WebhookResponse) GetWorkflowVersion() int32`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *WebhookResponse) GetWorkflowVersionOk() (*int32, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *WebhookResponse) SetWorkflowVersion(v int32)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *WebhookResponse) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.

### GetExecutionId

`func (o *WebhookResponse) GetExecutionId() int32`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WebhookResponse) GetExecutionIdOk() (*int32, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WebhookResponse) SetExecutionId(v int32)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *WebhookResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetCustomerId

`func (o *WebhookResponse) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WebhookResponse) GetCustomerIdOk() (*int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WebhookResponse) SetCustomerId(v int32)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *WebhookResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerUrl

`func (o *WebhookResponse) GetCustomerUrl() string`

GetCustomerUrl returns the CustomerUrl field if non-nil, zero value otherwise.

### GetCustomerUrlOk

`func (o *WebhookResponse) GetCustomerUrlOk() (*string, bool)`

GetCustomerUrlOk returns a tuple with the CustomerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUrl

`func (o *WebhookResponse) SetCustomerUrl(v string)`

SetCustomerUrl sets CustomerUrl field to given value.

### HasCustomerUrl

`func (o *WebhookResponse) HasCustomerUrl() bool`

HasCustomerUrl returns a boolean if a field has been set.

### GetAlertId

`func (o *WebhookResponse) GetAlertId() int32`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *WebhookResponse) GetAlertIdOk() (*int32, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *WebhookResponse) SetAlertId(v int32)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *WebhookResponse) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetDescription

`func (o *WebhookResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvent

`func (o *WebhookResponse) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookResponse) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookResponse) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhookResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventMetadata

`func (o *WebhookResponse) GetEventMetadata() map[string]interface{}`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *WebhookResponse) GetEventMetadataOk() (*map[string]interface{}, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *WebhookResponse) SetEventMetadata(v map[string]interface{})`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *WebhookResponse) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


