# EntrepriseFicheAllOfComptes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateDepot** | Pointer to **string** | Date de dépôt des comptes. | [optional] 
**DateDepotFormate** | Pointer to **string** | Date de dépôt formatée des comptes. | [optional] 
**DateCloture** | Pointer to **string** | Date de clôture des comptes, au format AAAA-MM-JJ. | [optional] 
**AnneeCloture** | Pointer to **int32** | Année de clôture des comptes. | [optional] 
**Confidentialite** | Pointer to **bool** | Confidentialité totale des comptes. | [optional] 
**ConfidentialiteCompteDeResultat** | Pointer to **bool** | Confidentialité partielle des comptes (seul le compte de résultat est confidentiel, le reste des comptes sont disponibles). | [optional] 
**Disponible** | Pointer to **bool** | Disponibilité des comptes au format PDF. | [optional] 
**NomFichierPdf** | Pointer to **string** | Nom du fichier PDF des comptes. | [optional] 
**Token** | Pointer to **string** | Token des comptes. | [optional] 
**DisponibleXlsx** | Pointer to **bool** | Disponibilité des comptes au format XLSX. | [optional] 
**NomFichierXlsx** | Pointer to **string** | Nom du fichier XLSX des comptes. | [optional] 
**TokenXlsx** | Pointer to **string** | Token des comptes XLSX. | [optional] 

## Methods

### NewEntrepriseFicheAllOfComptes

`func NewEntrepriseFicheAllOfComptes() *EntrepriseFicheAllOfComptes`

NewEntrepriseFicheAllOfComptes instantiates a new EntrepriseFicheAllOfComptes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfComptesWithDefaults

`func NewEntrepriseFicheAllOfComptesWithDefaults() *EntrepriseFicheAllOfComptes`

NewEntrepriseFicheAllOfComptesWithDefaults instantiates a new EntrepriseFicheAllOfComptes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateDepot

`func (o *EntrepriseFicheAllOfComptes) GetDateDepot() string`

GetDateDepot returns the DateDepot field if non-nil, zero value otherwise.

### GetDateDepotOk

`func (o *EntrepriseFicheAllOfComptes) GetDateDepotOk() (*string, bool)`

GetDateDepotOk returns a tuple with the DateDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepot

`func (o *EntrepriseFicheAllOfComptes) SetDateDepot(v string)`

SetDateDepot sets DateDepot field to given value.

### HasDateDepot

`func (o *EntrepriseFicheAllOfComptes) HasDateDepot() bool`

HasDateDepot returns a boolean if a field has been set.

### GetDateDepotFormate

`func (o *EntrepriseFicheAllOfComptes) GetDateDepotFormate() string`

GetDateDepotFormate returns the DateDepotFormate field if non-nil, zero value otherwise.

### GetDateDepotFormateOk

`func (o *EntrepriseFicheAllOfComptes) GetDateDepotFormateOk() (*string, bool)`

GetDateDepotFormateOk returns a tuple with the DateDepotFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepotFormate

`func (o *EntrepriseFicheAllOfComptes) SetDateDepotFormate(v string)`

SetDateDepotFormate sets DateDepotFormate field to given value.

### HasDateDepotFormate

`func (o *EntrepriseFicheAllOfComptes) HasDateDepotFormate() bool`

HasDateDepotFormate returns a boolean if a field has been set.

### GetDateCloture

`func (o *EntrepriseFicheAllOfComptes) GetDateCloture() string`

GetDateCloture returns the DateCloture field if non-nil, zero value otherwise.

### GetDateClotureOk

`func (o *EntrepriseFicheAllOfComptes) GetDateClotureOk() (*string, bool)`

GetDateClotureOk returns a tuple with the DateCloture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCloture

`func (o *EntrepriseFicheAllOfComptes) SetDateCloture(v string)`

SetDateCloture sets DateCloture field to given value.

### HasDateCloture

`func (o *EntrepriseFicheAllOfComptes) HasDateCloture() bool`

HasDateCloture returns a boolean if a field has been set.

### GetAnneeCloture

`func (o *EntrepriseFicheAllOfComptes) GetAnneeCloture() int32`

GetAnneeCloture returns the AnneeCloture field if non-nil, zero value otherwise.

### GetAnneeClotureOk

`func (o *EntrepriseFicheAllOfComptes) GetAnneeClotureOk() (*int32, bool)`

GetAnneeClotureOk returns a tuple with the AnneeCloture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeCloture

`func (o *EntrepriseFicheAllOfComptes) SetAnneeCloture(v int32)`

SetAnneeCloture sets AnneeCloture field to given value.

### HasAnneeCloture

`func (o *EntrepriseFicheAllOfComptes) HasAnneeCloture() bool`

HasAnneeCloture returns a boolean if a field has been set.

### GetConfidentialite

`func (o *EntrepriseFicheAllOfComptes) GetConfidentialite() bool`

GetConfidentialite returns the Confidentialite field if non-nil, zero value otherwise.

### GetConfidentialiteOk

`func (o *EntrepriseFicheAllOfComptes) GetConfidentialiteOk() (*bool, bool)`

GetConfidentialiteOk returns a tuple with the Confidentialite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialite

`func (o *EntrepriseFicheAllOfComptes) SetConfidentialite(v bool)`

SetConfidentialite sets Confidentialite field to given value.

### HasConfidentialite

`func (o *EntrepriseFicheAllOfComptes) HasConfidentialite() bool`

HasConfidentialite returns a boolean if a field has been set.

### GetConfidentialiteCompteDeResultat

`func (o *EntrepriseFicheAllOfComptes) GetConfidentialiteCompteDeResultat() bool`

GetConfidentialiteCompteDeResultat returns the ConfidentialiteCompteDeResultat field if non-nil, zero value otherwise.

### GetConfidentialiteCompteDeResultatOk

`func (o *EntrepriseFicheAllOfComptes) GetConfidentialiteCompteDeResultatOk() (*bool, bool)`

GetConfidentialiteCompteDeResultatOk returns a tuple with the ConfidentialiteCompteDeResultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialiteCompteDeResultat

`func (o *EntrepriseFicheAllOfComptes) SetConfidentialiteCompteDeResultat(v bool)`

SetConfidentialiteCompteDeResultat sets ConfidentialiteCompteDeResultat field to given value.

### HasConfidentialiteCompteDeResultat

`func (o *EntrepriseFicheAllOfComptes) HasConfidentialiteCompteDeResultat() bool`

HasConfidentialiteCompteDeResultat returns a boolean if a field has been set.

### GetDisponible

`func (o *EntrepriseFicheAllOfComptes) GetDisponible() bool`

GetDisponible returns the Disponible field if non-nil, zero value otherwise.

### GetDisponibleOk

`func (o *EntrepriseFicheAllOfComptes) GetDisponibleOk() (*bool, bool)`

GetDisponibleOk returns a tuple with the Disponible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisponible

`func (o *EntrepriseFicheAllOfComptes) SetDisponible(v bool)`

SetDisponible sets Disponible field to given value.

### HasDisponible

`func (o *EntrepriseFicheAllOfComptes) HasDisponible() bool`

HasDisponible returns a boolean if a field has been set.

### GetNomFichierPdf

`func (o *EntrepriseFicheAllOfComptes) GetNomFichierPdf() string`

GetNomFichierPdf returns the NomFichierPdf field if non-nil, zero value otherwise.

### GetNomFichierPdfOk

`func (o *EntrepriseFicheAllOfComptes) GetNomFichierPdfOk() (*string, bool)`

GetNomFichierPdfOk returns a tuple with the NomFichierPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomFichierPdf

`func (o *EntrepriseFicheAllOfComptes) SetNomFichierPdf(v string)`

SetNomFichierPdf sets NomFichierPdf field to given value.

### HasNomFichierPdf

`func (o *EntrepriseFicheAllOfComptes) HasNomFichierPdf() bool`

HasNomFichierPdf returns a boolean if a field has been set.

### GetToken

`func (o *EntrepriseFicheAllOfComptes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntrepriseFicheAllOfComptes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntrepriseFicheAllOfComptes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EntrepriseFicheAllOfComptes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDisponibleXlsx

`func (o *EntrepriseFicheAllOfComptes) GetDisponibleXlsx() bool`

GetDisponibleXlsx returns the DisponibleXlsx field if non-nil, zero value otherwise.

### GetDisponibleXlsxOk

`func (o *EntrepriseFicheAllOfComptes) GetDisponibleXlsxOk() (*bool, bool)`

GetDisponibleXlsxOk returns a tuple with the DisponibleXlsx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisponibleXlsx

`func (o *EntrepriseFicheAllOfComptes) SetDisponibleXlsx(v bool)`

SetDisponibleXlsx sets DisponibleXlsx field to given value.

### HasDisponibleXlsx

`func (o *EntrepriseFicheAllOfComptes) HasDisponibleXlsx() bool`

HasDisponibleXlsx returns a boolean if a field has been set.

### GetNomFichierXlsx

`func (o *EntrepriseFicheAllOfComptes) GetNomFichierXlsx() string`

GetNomFichierXlsx returns the NomFichierXlsx field if non-nil, zero value otherwise.

### GetNomFichierXlsxOk

`func (o *EntrepriseFicheAllOfComptes) GetNomFichierXlsxOk() (*string, bool)`

GetNomFichierXlsxOk returns a tuple with the NomFichierXlsx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomFichierXlsx

`func (o *EntrepriseFicheAllOfComptes) SetNomFichierXlsx(v string)`

SetNomFichierXlsx sets NomFichierXlsx field to given value.

### HasNomFichierXlsx

`func (o *EntrepriseFicheAllOfComptes) HasNomFichierXlsx() bool`

HasNomFichierXlsx returns a boolean if a field has been set.

### GetTokenXlsx

`func (o *EntrepriseFicheAllOfComptes) GetTokenXlsx() string`

GetTokenXlsx returns the TokenXlsx field if non-nil, zero value otherwise.

### GetTokenXlsxOk

`func (o *EntrepriseFicheAllOfComptes) GetTokenXlsxOk() (*string, bool)`

GetTokenXlsxOk returns a tuple with the TokenXlsx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenXlsx

`func (o *EntrepriseFicheAllOfComptes) SetTokenXlsx(v string)`

SetTokenXlsx sets TokenXlsx field to given value.

### HasTokenXlsx

`func (o *EntrepriseFicheAllOfComptes) HasTokenXlsx() bool`

HasTokenXlsx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


