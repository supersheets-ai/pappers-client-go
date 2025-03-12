# BodaccCreation

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
**DateDebutActivite** | Pointer to **string** | Date de début d&#39;activité de l&#39;entreprise concernée par la publication, au format AAAA-MM-JJ. | [optional] 

## Methods

### NewBodaccCreation

`func NewBodaccCreation() *BodaccCreation`

NewBodaccCreation instantiates a new BodaccCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccCreationWithDefaults

`func NewBodaccCreationWithDefaults() *BodaccCreation`

NewBodaccCreationWithDefaults instantiates a new BodaccCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomEntreprise

`func (o *BodaccCreation) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *BodaccCreation) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *BodaccCreation) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *BodaccCreation) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *BodaccCreation) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *BodaccCreation) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *BodaccCreation) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *BodaccCreation) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *BodaccCreation) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *BodaccCreation) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *BodaccCreation) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *BodaccCreation) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *BodaccCreation) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *BodaccCreation) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *BodaccCreation) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *BodaccCreation) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *BodaccCreation) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *BodaccCreation) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *BodaccCreation) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *BodaccCreation) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetAdministration

`func (o *BodaccCreation) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *BodaccCreation) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *BodaccCreation) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *BodaccCreation) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetAdresse

`func (o *BodaccCreation) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *BodaccCreation) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *BodaccCreation) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *BodaccCreation) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetCapital

`func (o *BodaccCreation) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *BodaccCreation) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *BodaccCreation) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *BodaccCreation) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetActivite

`func (o *BodaccCreation) GetActivite() string`

GetActivite returns the Activite field if non-nil, zero value otherwise.

### GetActiviteOk

`func (o *BodaccCreation) GetActiviteOk() (*string, bool)`

GetActiviteOk returns a tuple with the Activite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivite

`func (o *BodaccCreation) SetActivite(v string)`

SetActivite sets Activite field to given value.

### HasActivite

`func (o *BodaccCreation) HasActivite() bool`

HasActivite returns a boolean if a field has been set.

### GetDateDebutActivite

`func (o *BodaccCreation) GetDateDebutActivite() string`

GetDateDebutActivite returns the DateDebutActivite field if non-nil, zero value otherwise.

### GetDateDebutActiviteOk

`func (o *BodaccCreation) GetDateDebutActiviteOk() (*string, bool)`

GetDateDebutActiviteOk returns a tuple with the DateDebutActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebutActivite

`func (o *BodaccCreation) SetDateDebutActivite(v string)`

SetDateDebutActivite sets DateDebutActivite field to given value.

### HasDateDebutActivite

`func (o *BodaccCreation) HasDateDebutActivite() bool`

HasDateDebutActivite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


