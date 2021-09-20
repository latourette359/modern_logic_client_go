# DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DataSourceType**](DataSourceType.md) |  | [optional] 

## Methods

### NewDataSource

`func NewDataSource() *DataSource`

NewDataSource instantiates a new DataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceWithDefaults

`func NewDataSourceWithDefaults() *DataSource`

NewDataSourceWithDefaults instantiates a new DataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DataSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DataSource) GetType() DataSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataSource) GetTypeOk() (*DataSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataSource) SetType(v DataSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *DataSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


