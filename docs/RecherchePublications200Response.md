# RecherchePublications200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resultats** | Pointer to [**[]RecherchePublications200ResponseResultatsInner**](RecherchePublications200ResponseResultatsInner.md) | Liste des publications qui correspondent à la recherche. | [optional] 
**Total** | Pointer to **int32** | Nombre de publications qui correspondent à la recherche. | [optional] 
**Page** | Pointer to **int32** | Page actuelle. | [optional] 

## Methods

### NewRecherchePublications200Response

`func NewRecherchePublications200Response() *RecherchePublications200Response`

NewRecherchePublications200Response instantiates a new RecherchePublications200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecherchePublications200ResponseWithDefaults

`func NewRecherchePublications200ResponseWithDefaults() *RecherchePublications200Response`

NewRecherchePublications200ResponseWithDefaults instantiates a new RecherchePublications200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultats

`func (o *RecherchePublications200Response) GetResultats() []RecherchePublications200ResponseResultatsInner`

GetResultats returns the Resultats field if non-nil, zero value otherwise.

### GetResultatsOk

`func (o *RecherchePublications200Response) GetResultatsOk() (*[]RecherchePublications200ResponseResultatsInner, bool)`

GetResultatsOk returns a tuple with the Resultats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultats

`func (o *RecherchePublications200Response) SetResultats(v []RecherchePublications200ResponseResultatsInner)`

SetResultats sets Resultats field to given value.

### HasResultats

`func (o *RecherchePublications200Response) HasResultats() bool`

HasResultats returns a boolean if a field has been set.

### GetTotal

`func (o *RecherchePublications200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RecherchePublications200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RecherchePublications200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RecherchePublications200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *RecherchePublications200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RecherchePublications200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RecherchePublications200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *RecherchePublications200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


