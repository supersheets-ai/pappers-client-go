# RechercheDocuments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resultats** | Pointer to [**[]RechercheDocuments200ResponseResultatsInner**](RechercheDocuments200ResponseResultatsInner.md) | Liste des documents qui correspondent à la recherche. | [optional] 
**Total** | Pointer to **int32** | Nombre de documents qui correspondent à la recherche. | [optional] 
**Page** | Pointer to **int32** | Page actuelle. | [optional] 

## Methods

### NewRechercheDocuments200Response

`func NewRechercheDocuments200Response() *RechercheDocuments200Response`

NewRechercheDocuments200Response instantiates a new RechercheDocuments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercheDocuments200ResponseWithDefaults

`func NewRechercheDocuments200ResponseWithDefaults() *RechercheDocuments200Response`

NewRechercheDocuments200ResponseWithDefaults instantiates a new RechercheDocuments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultats

`func (o *RechercheDocuments200Response) GetResultats() []RechercheDocuments200ResponseResultatsInner`

GetResultats returns the Resultats field if non-nil, zero value otherwise.

### GetResultatsOk

`func (o *RechercheDocuments200Response) GetResultatsOk() (*[]RechercheDocuments200ResponseResultatsInner, bool)`

GetResultatsOk returns a tuple with the Resultats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultats

`func (o *RechercheDocuments200Response) SetResultats(v []RechercheDocuments200ResponseResultatsInner)`

SetResultats sets Resultats field to given value.

### HasResultats

`func (o *RechercheDocuments200Response) HasResultats() bool`

HasResultats returns a boolean if a field has been set.

### GetTotal

`func (o *RechercheDocuments200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RechercheDocuments200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RechercheDocuments200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RechercheDocuments200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *RechercheDocuments200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RechercheDocuments200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RechercheDocuments200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *RechercheDocuments200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


