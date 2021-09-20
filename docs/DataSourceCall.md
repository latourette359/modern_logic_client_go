# DataSourceCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ExecutionId** | Pointer to **int32** |  | [optional] 
**DataSourceId** | Pointer to **int32** |  | [optional] 
**Request** | Pointer to **map[string]interface{}** |  | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**DateTimeOfCall** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDataSourceCall

`func NewDataSourceCall() *DataSourceCall`

NewDataSourceCall instantiates a new DataSourceCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceCallWithDefaults

`func NewDataSourceCallWithDefaults() *DataSourceCall`

NewDataSourceCallWithDefaults instantiates a new DataSourceCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceCall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceCall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceCall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DataSourceCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExecutionId

`func (o *DataSourceCall) GetExecutionId() int32`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DataSourceCall) GetExecutionIdOk() (*int32, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DataSourceCall) SetExecutionId(v int32)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DataSourceCall) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetDataSourceId

`func (o *DataSourceCall) GetDataSourceId() int32`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *DataSourceCall) GetDataSourceIdOk() (*int32, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *DataSourceCall) SetDataSourceId(v int32)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *DataSourceCall) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### GetRequest

`func (o *DataSourceCall) GetRequest() map[string]interface{}`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DataSourceCall) GetRequestOk() (*map[string]interface{}, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DataSourceCall) SetRequest(v map[string]interface{})`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DataSourceCall) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *DataSourceCall) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DataSourceCall) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DataSourceCall) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DataSourceCall) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetDateTimeOfCall

`func (o *DataSourceCall) GetDateTimeOfCall() time.Time`

GetDateTimeOfCall returns the DateTimeOfCall field if non-nil, zero value otherwise.

### GetDateTimeOfCallOk

`func (o *DataSourceCall) GetDateTimeOfCallOk() (*time.Time, bool)`

GetDateTimeOfCallOk returns a tuple with the DateTimeOfCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeOfCall

`func (o *DataSourceCall) SetDateTimeOfCall(v time.Time)`

SetDateTimeOfCall sets DateTimeOfCall field to given value.

### HasDateTimeOfCall

`func (o *DataSourceCall) HasDateTimeOfCall() bool`

HasDateTimeOfCall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


