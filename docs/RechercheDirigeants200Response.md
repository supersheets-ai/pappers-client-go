# RechercheDirigeants200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resultats** | Pointer to [**[]RechercheDirigeants200ResponseResultatsInner**](RechercheDirigeants200ResponseResultatsInner.md) | Liste des dirigeants qui correspondent à la recherche. | [optional] 
**Total** | Pointer to **int32** | Nombre de dirigeants qui correspondent à la recherche. | [optional] 
**Page** | Pointer to **int32** | Page actuelle. | [optional] 

## Methods

### NewRechercheDirigeants200Response

`func NewRechercheDirigeants200Response() *RechercheDirigeants200Response`

NewRechercheDirigeants200Response instantiates a new RechercheDirigeants200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercheDirigeants200ResponseWithDefaults

`func NewRechercheDirigeants200ResponseWithDefaults() *RechercheDirigeants200Response`

NewRechercheDirigeants200ResponseWithDefaults instantiates a new RechercheDirigeants200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultats

`func (o *RechercheDirigeants200Response) GetResultats() []RechercheDirigeants200ResponseResultatsInner`

GetResultats returns the Resultats field if non-nil, zero value otherwise.

### GetResultatsOk

`func (o *RechercheDirigeants200Response) GetResultatsOk() (*[]RechercheDirigeants200ResponseResultatsInner, bool)`

GetResultatsOk returns a tuple with the Resultats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultats

`func (o *RechercheDirigeants200Response) SetResultats(v []RechercheDirigeants200ResponseResultatsInner)`

SetResultats sets Resultats field to given value.

### HasResultats

`func (o *RechercheDirigeants200Response) HasResultats() bool`

HasResultats returns a boolean if a field has been set.

### GetTotal

`func (o *RechercheDirigeants200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RechercheDirigeants200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RechercheDirigeants200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RechercheDirigeants200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *RechercheDirigeants200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RechercheDirigeants200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RechercheDirigeants200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *RechercheDirigeants200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


