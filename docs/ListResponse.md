# ListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of entities | [optional] 
**TotalReturned** | Pointer to **int32** | Total number of entities returned | [optional] 
**Items** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListResponse

`func NewListResponse() *ListResponse`

NewListResponse instantiates a new ListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseWithDefaults

`func NewListResponseWithDefaults() *ListResponse`

NewListResponseWithDefaults instantiates a new ListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalReturned

`func (o *ListResponse) GetTotalReturned() int32`

GetTotalReturned returns the TotalReturned field if non-nil, zero value otherwise.

### GetTotalReturnedOk

`func (o *ListResponse) GetTotalReturnedOk() (*int32, bool)`

GetTotalReturnedOk returns a tuple with the TotalReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturned

`func (o *ListResponse) SetTotalReturned(v int32)`

SetTotalReturned sets TotalReturned field to given value.

### HasTotalReturned

`func (o *ListResponse) HasTotalReturned() bool`

HasTotalReturned returns a boolean if a field has been set.

### GetItems

`func (o *ListResponse) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListResponse) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListResponse) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *ListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


