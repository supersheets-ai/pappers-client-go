# RecherchePublications200ResponseResultatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type de publication | [optional] 
**Date** | Pointer to **string** | Date de publication, au format AAAA-MM-JJ. | [optional] 
**Contenu** | Pointer to **string** | Contenu de la publication, avec les mentions correspondant à la recherche encerclée par une balise HTML &lt;em&gt;. | [optional] 
**Entreprise** | Pointer to [**EntrepriseRecherche**](EntrepriseRecherche.md) |  | [optional] 

## Methods

### NewRecherchePublications200ResponseResultatsInner

`func NewRecherchePublications200ResponseResultatsInner() *RecherchePublications200ResponseResultatsInner`

NewRecherchePublications200ResponseResultatsInner instantiates a new RecherchePublications200ResponseResultatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecherchePublications200ResponseResultatsInnerWithDefaults

`func NewRecherchePublications200ResponseResultatsInnerWithDefaults() *RecherchePublications200ResponseResultatsInner`

NewRecherchePublications200ResponseResultatsInnerWithDefaults instantiates a new RecherchePublications200ResponseResultatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecherchePublications200ResponseResultatsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecherchePublications200ResponseResultatsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecherchePublications200ResponseResultatsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecherchePublications200ResponseResultatsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDate

`func (o *RecherchePublications200ResponseResultatsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RecherchePublications200ResponseResultatsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RecherchePublications200ResponseResultatsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RecherchePublications200ResponseResultatsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetContenu

`func (o *RecherchePublications200ResponseResultatsInner) GetContenu() string`

GetContenu returns the Contenu field if non-nil, zero value otherwise.

### GetContenuOk

`func (o *RecherchePublications200ResponseResultatsInner) GetContenuOk() (*string, bool)`

GetContenuOk returns a tuple with the Contenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContenu

`func (o *RecherchePublications200ResponseResultatsInner) SetContenu(v string)`

SetContenu sets Contenu field to given value.

### HasContenu

`func (o *RecherchePublications200ResponseResultatsInner) HasContenu() bool`

HasContenu returns a boolean if a field has been set.

### GetEntreprise

`func (o *RecherchePublications200ResponseResultatsInner) GetEntreprise() EntrepriseRecherche`

GetEntreprise returns the Entreprise field if non-nil, zero value otherwise.

### GetEntrepriseOk

`func (o *RecherchePublications200ResponseResultatsInner) GetEntrepriseOk() (*EntrepriseRecherche, bool)`

GetEntrepriseOk returns a tuple with the Entreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntreprise

`func (o *RecherchePublications200ResponseResultatsInner) SetEntreprise(v EntrepriseRecherche)`

SetEntreprise sets Entreprise field to given value.

### HasEntreprise

`func (o *RecherchePublications200ResponseResultatsInner) HasEntreprise() bool`

HasEntreprise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


