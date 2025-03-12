# SurveillanceListeInformationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to **[]string** | Tableau d&#39;identifiant de notifications | [optional] 
**Informations** | Pointer to **string** | Information à rajouter sur les notifications | [optional] 

## Methods

### NewSurveillanceListeInformationsRequest

`func NewSurveillanceListeInformationsRequest() *SurveillanceListeInformationsRequest`

NewSurveillanceListeInformationsRequest instantiates a new SurveillanceListeInformationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveillanceListeInformationsRequestWithDefaults

`func NewSurveillanceListeInformationsRequestWithDefaults() *SurveillanceListeInformationsRequest`

NewSurveillanceListeInformationsRequestWithDefaults instantiates a new SurveillanceListeInformationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *SurveillanceListeInformationsRequest) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SurveillanceListeInformationsRequest) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SurveillanceListeInformationsRequest) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SurveillanceListeInformationsRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetInformations

`func (o *SurveillanceListeInformationsRequest) GetInformations() string`

GetInformations returns the Informations field if non-nil, zero value otherwise.

### GetInformationsOk

`func (o *SurveillanceListeInformationsRequest) GetInformationsOk() (*string, bool)`

GetInformationsOk returns a tuple with the Informations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformations

`func (o *SurveillanceListeInformationsRequest) SetInformations(v string)`

SetInformations sets Informations field to given value.

### HasInformations

`func (o *SurveillanceListeInformationsRequest) HasInformations() bool`

HasInformations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


