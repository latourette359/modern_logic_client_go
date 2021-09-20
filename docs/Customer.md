# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerId** | Pointer to **string** | Way that you uniquely identify customers | [optional] 
**FirstName** | Pointer to **string** | Legal first name of the user being evaluated. | [optional] 
**LastName** | Pointer to **string** | Legal last name (surname) of the user being evaluated. | [optional] 
**Email** | Pointer to **string** | Email address provided by user being evaluated. | [optional] 
**Phone** | Pointer to **string** | Phone number of user being evaluated. | [optional] 
**AddressStreet** | Pointer to **string** | Home address of user being evaluated | [optional] 
**AddressStreet2** | Pointer to **string** |  | [optional] 
**AddressCity** | Pointer to **string** |  | [optional] 
**AddressStateCode** | Pointer to **string** |  | [optional] 
**AddressZip** | Pointer to **string** |  | [optional] 
**AddressCountryCode** | Pointer to **string** |  | [optional] 
**Dob** | Pointer to **string** | Date of birth for user being evaluated | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraProperties

`func (o *Customer) GetExtraProperties() map[string]interface{}`

GetExtraProperties returns the ExtraProperties field if non-nil, zero value otherwise.

### GetExtraPropertiesOk

`func (o *Customer) GetExtraPropertiesOk() (*map[string]interface{}, bool)`

GetExtraPropertiesOk returns a tuple with the ExtraProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraProperties

`func (o *Customer) SetExtraProperties(v map[string]interface{})`

SetExtraProperties sets ExtraProperties field to given value.

### HasExtraProperties

`func (o *Customer) HasExtraProperties() bool`

HasExtraProperties returns a boolean if a field has been set.

### GetCustomerId

`func (o *Customer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Customer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Customer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Customer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFirstName

`func (o *Customer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Customer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Customer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Customer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Customer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Customer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Customer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Customer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Customer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddressStreet

`func (o *Customer) GetAddressStreet() string`

GetAddressStreet returns the AddressStreet field if non-nil, zero value otherwise.

### GetAddressStreetOk

`func (o *Customer) GetAddressStreetOk() (*string, bool)`

GetAddressStreetOk returns a tuple with the AddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet

`func (o *Customer) SetAddressStreet(v string)`

SetAddressStreet sets AddressStreet field to given value.

### HasAddressStreet

`func (o *Customer) HasAddressStreet() bool`

HasAddressStreet returns a boolean if a field has been set.

### GetAddressStreet2

`func (o *Customer) GetAddressStreet2() string`

GetAddressStreet2 returns the AddressStreet2 field if non-nil, zero value otherwise.

### GetAddressStreet2Ok

`func (o *Customer) GetAddressStreet2Ok() (*string, bool)`

GetAddressStreet2Ok returns a tuple with the AddressStreet2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStreet2

`func (o *Customer) SetAddressStreet2(v string)`

SetAddressStreet2 sets AddressStreet2 field to given value.

### HasAddressStreet2

`func (o *Customer) HasAddressStreet2() bool`

HasAddressStreet2 returns a boolean if a field has been set.

### GetAddressCity

`func (o *Customer) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *Customer) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *Customer) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *Customer) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressStateCode

`func (o *Customer) GetAddressStateCode() string`

GetAddressStateCode returns the AddressStateCode field if non-nil, zero value otherwise.

### GetAddressStateCodeOk

`func (o *Customer) GetAddressStateCodeOk() (*string, bool)`

GetAddressStateCodeOk returns a tuple with the AddressStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressStateCode

`func (o *Customer) SetAddressStateCode(v string)`

SetAddressStateCode sets AddressStateCode field to given value.

### HasAddressStateCode

`func (o *Customer) HasAddressStateCode() bool`

HasAddressStateCode returns a boolean if a field has been set.

### GetAddressZip

`func (o *Customer) GetAddressZip() string`

GetAddressZip returns the AddressZip field if non-nil, zero value otherwise.

### GetAddressZipOk

`func (o *Customer) GetAddressZipOk() (*string, bool)`

GetAddressZipOk returns a tuple with the AddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressZip

`func (o *Customer) SetAddressZip(v string)`

SetAddressZip sets AddressZip field to given value.

### HasAddressZip

`func (o *Customer) HasAddressZip() bool`

HasAddressZip returns a boolean if a field has been set.

### GetAddressCountryCode

`func (o *Customer) GetAddressCountryCode() string`

GetAddressCountryCode returns the AddressCountryCode field if non-nil, zero value otherwise.

### GetAddressCountryCodeOk

`func (o *Customer) GetAddressCountryCodeOk() (*string, bool)`

GetAddressCountryCodeOk returns a tuple with the AddressCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountryCode

`func (o *Customer) SetAddressCountryCode(v string)`

SetAddressCountryCode sets AddressCountryCode field to given value.

### HasAddressCountryCode

`func (o *Customer) HasAddressCountryCode() bool`

HasAddressCountryCode returns a boolean if a field has been set.

### GetDob

`func (o *Customer) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Customer) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Customer) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Customer) HasDob() bool`

HasDob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


