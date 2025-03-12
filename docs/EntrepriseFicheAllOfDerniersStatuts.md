# EntrepriseFicheAllOfDerniersStatuts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateDepot** | Pointer to **string** | Date de dépôt des statuts, au format AAAA-MM-JJ. | [optional] 
**DateDepotFormate** | Pointer to **string** | Date de dépôt formaté des statuts, au format JJ/MM/AAAA. | [optional] 
**Disponible** | Pointer to **bool** | Disponibilité des statuts. | [optional] 
**NomFichierPdf** | Pointer to **string** | Nom du fichier pdf des statuts. | [optional] 
**Token** | Pointer to **string** | Token des statuts. | [optional] 
**Type** | Pointer to **string** | Champ \&quot;type\&quot; du document contenant les statuts. | [optional] 
**Decision** | Pointer to **string** | Champ \&quot;decision\&quot; du document contenant les statuts. | [optional] 
**DateActe** | Pointer to **string** | Date de publication des statuts, au format AAAA-MM-JJ. | [optional] 
**DateActeFormate** | Pointer to **string** | Date de publication des statuts, au format JJ/MM/AAAA. | [optional] 

## Methods

### NewEntrepriseFicheAllOfDerniersStatuts

`func NewEntrepriseFicheAllOfDerniersStatuts() *EntrepriseFicheAllOfDerniersStatuts`

NewEntrepriseFicheAllOfDerniersStatuts instantiates a new EntrepriseFicheAllOfDerniersStatuts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfDerniersStatutsWithDefaults

`func NewEntrepriseFicheAllOfDerniersStatutsWithDefaults() *EntrepriseFicheAllOfDerniersStatuts`

NewEntrepriseFicheAllOfDerniersStatutsWithDefaults instantiates a new EntrepriseFicheAllOfDerniersStatuts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateDepot

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateDepot() string`

GetDateDepot returns the DateDepot field if non-nil, zero value otherwise.

### GetDateDepotOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateDepotOk() (*string, bool)`

GetDateDepotOk returns a tuple with the DateDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepot

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDateDepot(v string)`

SetDateDepot sets DateDepot field to given value.

### HasDateDepot

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDateDepot() bool`

HasDateDepot returns a boolean if a field has been set.

### GetDateDepotFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateDepotFormate() string`

GetDateDepotFormate returns the DateDepotFormate field if non-nil, zero value otherwise.

### GetDateDepotFormateOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateDepotFormateOk() (*string, bool)`

GetDateDepotFormateOk returns a tuple with the DateDepotFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepotFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDateDepotFormate(v string)`

SetDateDepotFormate sets DateDepotFormate field to given value.

### HasDateDepotFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDateDepotFormate() bool`

HasDateDepotFormate returns a boolean if a field has been set.

### GetDisponible

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDisponible() bool`

GetDisponible returns the Disponible field if non-nil, zero value otherwise.

### GetDisponibleOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDisponibleOk() (*bool, bool)`

GetDisponibleOk returns a tuple with the Disponible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisponible

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDisponible(v bool)`

SetDisponible sets Disponible field to given value.

### HasDisponible

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDisponible() bool`

HasDisponible returns a boolean if a field has been set.

### GetNomFichierPdf

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetNomFichierPdf() string`

GetNomFichierPdf returns the NomFichierPdf field if non-nil, zero value otherwise.

### GetNomFichierPdfOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetNomFichierPdfOk() (*string, bool)`

GetNomFichierPdfOk returns a tuple with the NomFichierPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomFichierPdf

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetNomFichierPdf(v string)`

SetNomFichierPdf sets NomFichierPdf field to given value.

### HasNomFichierPdf

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasNomFichierPdf() bool`

HasNomFichierPdf returns a boolean if a field has been set.

### GetToken

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDecision

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetDateActe

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateActe() string`

GetDateActe returns the DateActe field if non-nil, zero value otherwise.

### GetDateActeOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateActeOk() (*string, bool)`

GetDateActeOk returns a tuple with the DateActe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateActe

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDateActe(v string)`

SetDateActe sets DateActe field to given value.

### HasDateActe

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDateActe() bool`

HasDateActe returns a boolean if a field has been set.

### GetDateActeFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateActeFormate() string`

GetDateActeFormate returns the DateActeFormate field if non-nil, zero value otherwise.

### GetDateActeFormateOk

`func (o *EntrepriseFicheAllOfDerniersStatuts) GetDateActeFormateOk() (*string, bool)`

GetDateActeFormateOk returns a tuple with the DateActeFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateActeFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) SetDateActeFormate(v string)`

SetDateActeFormate sets DateActeFormate field to given value.

### HasDateActeFormate

`func (o *EntrepriseFicheAllOfDerniersStatuts) HasDateActeFormate() bool`

HasDateActeFormate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


