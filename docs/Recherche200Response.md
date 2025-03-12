# Recherche200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resultats** | Pointer to [**[]Recherche200ResponseResultatsInner**](Recherche200ResponseResultatsInner.md) | Liste des entreprises qui correspondent à la recherche. | [optional] 
**Total** | Pointer to **int32** | Nombre d&#39;entreprises qui correspondent à la recherche. | [optional] 
**Page** | Pointer to **int32** | Page actuelle. | [optional] 
**CurseurSuivant** | Pointer to **string** | Présent uniquement en cas d&#39;utilisation du paramètre &#x60;curseur&#x60;. Contient le curseur courant à envoyer en paramètre &#x60;curseur&#x60; de la requête suivantes. | [optional] 

## Methods

### NewRecherche200Response

`func NewRecherche200Response() *Recherche200Response`

NewRecherche200Response instantiates a new Recherche200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecherche200ResponseWithDefaults

`func NewRecherche200ResponseWithDefaults() *Recherche200Response`

NewRecherche200ResponseWithDefaults instantiates a new Recherche200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultats

`func (o *Recherche200Response) GetResultats() []Recherche200ResponseResultatsInner`

GetResultats returns the Resultats field if non-nil, zero value otherwise.

### GetResultatsOk

`func (o *Recherche200Response) GetResultatsOk() (*[]Recherche200ResponseResultatsInner, bool)`

GetResultatsOk returns a tuple with the Resultats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultats

`func (o *Recherche200Response) SetResultats(v []Recherche200ResponseResultatsInner)`

SetResultats sets Resultats field to given value.

### HasResultats

`func (o *Recherche200Response) HasResultats() bool`

HasResultats returns a boolean if a field has been set.

### GetTotal

`func (o *Recherche200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Recherche200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Recherche200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Recherche200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *Recherche200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Recherche200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Recherche200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *Recherche200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetCurseurSuivant

`func (o *Recherche200Response) GetCurseurSuivant() string`

GetCurseurSuivant returns the CurseurSuivant field if non-nil, zero value otherwise.

### GetCurseurSuivantOk

`func (o *Recherche200Response) GetCurseurSuivantOk() (*string, bool)`

GetCurseurSuivantOk returns a tuple with the CurseurSuivant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurseurSuivant

`func (o *Recherche200Response) SetCurseurSuivant(v string)`

SetCurseurSuivant sets CurseurSuivant field to given value.

### HasCurseurSuivant

`func (o *Recherche200Response) HasCurseurSuivant() bool`

HasCurseurSuivant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


