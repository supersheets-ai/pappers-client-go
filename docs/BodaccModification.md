# BodaccModification

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
**Capital** | Pointer to **float32** | Capital de l&#39;entreprise concernée par la publication. | [optional] 
**Activite** | Pointer to **string** | Activité de l&#39;entreprise concernée par la publication. | [optional] 
**Description** | Pointer to **string** | Description de la modification. | [optional] 

## Methods

### NewBodaccModification

`func NewBodaccModification() *BodaccModification`

NewBodaccModification instantiates a new BodaccModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccModificationWithDefaults

`func NewBodaccModificationWithDefaults() *BodaccModification`

NewBodaccModificationWithDefaults instantiates a new BodaccModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomEntreprise

`func (o *BodaccModification) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *BodaccModification) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *BodaccModification) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *BodaccModification) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *BodaccModification) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *BodaccModification) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *BodaccModification) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *BodaccModification) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *BodaccModification) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *BodaccModification) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *BodaccModification) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *BodaccModification) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *BodaccModification) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *BodaccModification) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *BodaccModification) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *BodaccModification) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *BodaccModification) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *BodaccModification) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *BodaccModification) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *BodaccModification) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetAdministration

`func (o *BodaccModification) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *BodaccModification) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *BodaccModification) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *BodaccModification) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetAdresse

`func (o *BodaccModification) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *BodaccModification) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *BodaccModification) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *BodaccModification) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetCapital

`func (o *BodaccModification) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *BodaccModification) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *BodaccModification) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *BodaccModification) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetActivite

`func (o *BodaccModification) GetActivite() string`

GetActivite returns the Activite field if non-nil, zero value otherwise.

### GetActiviteOk

`func (o *BodaccModification) GetActiviteOk() (*string, bool)`

GetActiviteOk returns a tuple with the Activite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivite

`func (o *BodaccModification) SetActivite(v string)`

SetActivite sets Activite field to given value.

### HasActivite

`func (o *BodaccModification) HasActivite() bool`

HasActivite returns a boolean if a field has been set.

### GetDescription

`func (o *BodaccModification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BodaccModification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BodaccModification) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BodaccModification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


