# EntrepriseFicheAllOfDepotsActes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateDepot** | Pointer to **string** | Date de dépôt de l&#39;acte, au format AAAA-MM-JJ. | [optional] 
**DateDepotFormate** | Pointer to **string** | Date de dépôt formatée de l&#39;acte, au format JJ/MM/AAAA. | [optional] 
**Disponible** | Pointer to **bool** | Disponibilité de l&#39;acte. Un acte peut être indisponible car il a été publié avant le 1er janvier 1993 ou bien car il est confidentiel. | [optional] 
**NomFichierPdf** | Pointer to **string** | Nom du fichier pdf de l&#39;acte. | [optional] 
**Token** | Pointer to **string** | Token de l&#39;acte. | [optional] 
**Actes** | Pointer to [**[]EntrepriseFicheAllOfActes**](EntrepriseFicheAllOfActes.md) | Détails de l&#39;acte. | [optional] 

## Methods

### NewEntrepriseFicheAllOfDepotsActes

`func NewEntrepriseFicheAllOfDepotsActes() *EntrepriseFicheAllOfDepotsActes`

NewEntrepriseFicheAllOfDepotsActes instantiates a new EntrepriseFicheAllOfDepotsActes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfDepotsActesWithDefaults

`func NewEntrepriseFicheAllOfDepotsActesWithDefaults() *EntrepriseFicheAllOfDepotsActes`

NewEntrepriseFicheAllOfDepotsActesWithDefaults instantiates a new EntrepriseFicheAllOfDepotsActes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateDepot

`func (o *EntrepriseFicheAllOfDepotsActes) GetDateDepot() string`

GetDateDepot returns the DateDepot field if non-nil, zero value otherwise.

### GetDateDepotOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetDateDepotOk() (*string, bool)`

GetDateDepotOk returns a tuple with the DateDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepot

`func (o *EntrepriseFicheAllOfDepotsActes) SetDateDepot(v string)`

SetDateDepot sets DateDepot field to given value.

### HasDateDepot

`func (o *EntrepriseFicheAllOfDepotsActes) HasDateDepot() bool`

HasDateDepot returns a boolean if a field has been set.

### GetDateDepotFormate

`func (o *EntrepriseFicheAllOfDepotsActes) GetDateDepotFormate() string`

GetDateDepotFormate returns the DateDepotFormate field if non-nil, zero value otherwise.

### GetDateDepotFormateOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetDateDepotFormateOk() (*string, bool)`

GetDateDepotFormateOk returns a tuple with the DateDepotFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepotFormate

`func (o *EntrepriseFicheAllOfDepotsActes) SetDateDepotFormate(v string)`

SetDateDepotFormate sets DateDepotFormate field to given value.

### HasDateDepotFormate

`func (o *EntrepriseFicheAllOfDepotsActes) HasDateDepotFormate() bool`

HasDateDepotFormate returns a boolean if a field has been set.

### GetDisponible

`func (o *EntrepriseFicheAllOfDepotsActes) GetDisponible() bool`

GetDisponible returns the Disponible field if non-nil, zero value otherwise.

### GetDisponibleOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetDisponibleOk() (*bool, bool)`

GetDisponibleOk returns a tuple with the Disponible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisponible

`func (o *EntrepriseFicheAllOfDepotsActes) SetDisponible(v bool)`

SetDisponible sets Disponible field to given value.

### HasDisponible

`func (o *EntrepriseFicheAllOfDepotsActes) HasDisponible() bool`

HasDisponible returns a boolean if a field has been set.

### GetNomFichierPdf

`func (o *EntrepriseFicheAllOfDepotsActes) GetNomFichierPdf() string`

GetNomFichierPdf returns the NomFichierPdf field if non-nil, zero value otherwise.

### GetNomFichierPdfOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetNomFichierPdfOk() (*string, bool)`

GetNomFichierPdfOk returns a tuple with the NomFichierPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomFichierPdf

`func (o *EntrepriseFicheAllOfDepotsActes) SetNomFichierPdf(v string)`

SetNomFichierPdf sets NomFichierPdf field to given value.

### HasNomFichierPdf

`func (o *EntrepriseFicheAllOfDepotsActes) HasNomFichierPdf() bool`

HasNomFichierPdf returns a boolean if a field has been set.

### GetToken

`func (o *EntrepriseFicheAllOfDepotsActes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntrepriseFicheAllOfDepotsActes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EntrepriseFicheAllOfDepotsActes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetActes

`func (o *EntrepriseFicheAllOfDepotsActes) GetActes() []EntrepriseFicheAllOfActes`

GetActes returns the Actes field if non-nil, zero value otherwise.

### GetActesOk

`func (o *EntrepriseFicheAllOfDepotsActes) GetActesOk() (*[]EntrepriseFicheAllOfActes, bool)`

GetActesOk returns a tuple with the Actes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActes

`func (o *EntrepriseFicheAllOfDepotsActes) SetActes(v []EntrepriseFicheAllOfActes)`

SetActes sets Actes field to given value.

### HasActes

`func (o *EntrepriseFicheAllOfDepotsActes) HasActes() bool`

HasActes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


