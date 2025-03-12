# BodaccImmatriculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NomEntreprise** | Pointer to **string** | Nom de l&#39;entreprise concernée par la publication. Correspond à la dénomination en cas de personne morale et à nom + prenom en cas de personne physique. | [optional] 
**PersonneMorale** | Pointer to **bool** | Vrai si l&#39;entreprise concernée par la publication est une personne morale, faux si c&#39;est une personne physique. | [optional] 
**Denomination** | Pointer to **string** | Dénomination de l&#39;entreprise concernée par la publication (uniquement en cas de personne morale). | [optional] 
**Nom** | Pointer to **string** | Nom de la personne physique concernée par la publication (uniquement en cas de personne physique). | [optional] 
**Prenom** | Pointer to **string** | Prénom de la personne physique concernée par la publication (uniquement en cas de personne physique). | [optional] 
**Administration** | Pointer to **string** | Administration (dans le cas d&#39;une personne morale). | [optional] 
**Adresse** | Pointer to **string** | Adresse de l&#39;entreprise concernée par la publication. | [optional] 

## Methods

### NewBodaccImmatriculation

`func NewBodaccImmatriculation() *BodaccImmatriculation`

NewBodaccImmatriculation instantiates a new BodaccImmatriculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccImmatriculationWithDefaults

`func NewBodaccImmatriculationWithDefaults() *BodaccImmatriculation`

NewBodaccImmatriculationWithDefaults instantiates a new BodaccImmatriculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomEntreprise

`func (o *BodaccImmatriculation) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *BodaccImmatriculation) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *BodaccImmatriculation) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *BodaccImmatriculation) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *BodaccImmatriculation) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *BodaccImmatriculation) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *BodaccImmatriculation) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *BodaccImmatriculation) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *BodaccImmatriculation) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *BodaccImmatriculation) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *BodaccImmatriculation) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *BodaccImmatriculation) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *BodaccImmatriculation) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *BodaccImmatriculation) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *BodaccImmatriculation) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *BodaccImmatriculation) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *BodaccImmatriculation) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *BodaccImmatriculation) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *BodaccImmatriculation) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *BodaccImmatriculation) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetAdministration

`func (o *BodaccImmatriculation) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *BodaccImmatriculation) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *BodaccImmatriculation) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *BodaccImmatriculation) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetAdresse

`func (o *BodaccImmatriculation) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *BodaccImmatriculation) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *BodaccImmatriculation) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *BodaccImmatriculation) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


