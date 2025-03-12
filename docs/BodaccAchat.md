# BodaccAchat

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
**Commentaires** | Pointer to **string** | Commentaires sur la publication. | [optional] 
**Oppositions** | Pointer to **string** | Détails sur les oppositions. | [optional] 
**DeclarationCreance** | Pointer to **string** | Détails sur la déclaration de créance. | [optional] 
**PublicationLegale** | Pointer to **string** | Journal où a été publiée la publication légale. | [optional] 
**DenominationAncienProprietaire** | Pointer to **string** | Dénomination de l&#39;ancien propriétaire de l&#39;établisement. | [optional] 
**SirenAncienProprietaire** | Pointer to **string** | Siren de l&#39;ancien propriétaire de l&#39;établisement. | [optional] 
**DenominationAncienExploitant** | Pointer to **string** | Dénomination de l&#39;ancien exploitant de l&#39;établisement. | [optional] 
**SirenAncienExploitant** | Pointer to **string** | Siren de l&#39;ancien exploitant de l&#39;établisement. | [optional] 

## Methods

### NewBodaccAchat

`func NewBodaccAchat() *BodaccAchat`

NewBodaccAchat instantiates a new BodaccAchat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccAchatWithDefaults

`func NewBodaccAchatWithDefaults() *BodaccAchat`

NewBodaccAchatWithDefaults instantiates a new BodaccAchat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomEntreprise

`func (o *BodaccAchat) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *BodaccAchat) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *BodaccAchat) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *BodaccAchat) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *BodaccAchat) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *BodaccAchat) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *BodaccAchat) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *BodaccAchat) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *BodaccAchat) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *BodaccAchat) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *BodaccAchat) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *BodaccAchat) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *BodaccAchat) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *BodaccAchat) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *BodaccAchat) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *BodaccAchat) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *BodaccAchat) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *BodaccAchat) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *BodaccAchat) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *BodaccAchat) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetAdministration

`func (o *BodaccAchat) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *BodaccAchat) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *BodaccAchat) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *BodaccAchat) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetAdresse

`func (o *BodaccAchat) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *BodaccAchat) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *BodaccAchat) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *BodaccAchat) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetCommentaires

`func (o *BodaccAchat) GetCommentaires() string`

GetCommentaires returns the Commentaires field if non-nil, zero value otherwise.

### GetCommentairesOk

`func (o *BodaccAchat) GetCommentairesOk() (*string, bool)`

GetCommentairesOk returns a tuple with the Commentaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentaires

`func (o *BodaccAchat) SetCommentaires(v string)`

SetCommentaires sets Commentaires field to given value.

### HasCommentaires

`func (o *BodaccAchat) HasCommentaires() bool`

HasCommentaires returns a boolean if a field has been set.

### GetOppositions

`func (o *BodaccAchat) GetOppositions() string`

GetOppositions returns the Oppositions field if non-nil, zero value otherwise.

### GetOppositionsOk

`func (o *BodaccAchat) GetOppositionsOk() (*string, bool)`

GetOppositionsOk returns a tuple with the Oppositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositions

`func (o *BodaccAchat) SetOppositions(v string)`

SetOppositions sets Oppositions field to given value.

### HasOppositions

`func (o *BodaccAchat) HasOppositions() bool`

HasOppositions returns a boolean if a field has been set.

### GetDeclarationCreance

`func (o *BodaccAchat) GetDeclarationCreance() string`

GetDeclarationCreance returns the DeclarationCreance field if non-nil, zero value otherwise.

### GetDeclarationCreanceOk

`func (o *BodaccAchat) GetDeclarationCreanceOk() (*string, bool)`

GetDeclarationCreanceOk returns a tuple with the DeclarationCreance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationCreance

`func (o *BodaccAchat) SetDeclarationCreance(v string)`

SetDeclarationCreance sets DeclarationCreance field to given value.

### HasDeclarationCreance

`func (o *BodaccAchat) HasDeclarationCreance() bool`

HasDeclarationCreance returns a boolean if a field has been set.

### GetPublicationLegale

`func (o *BodaccAchat) GetPublicationLegale() string`

GetPublicationLegale returns the PublicationLegale field if non-nil, zero value otherwise.

### GetPublicationLegaleOk

`func (o *BodaccAchat) GetPublicationLegaleOk() (*string, bool)`

GetPublicationLegaleOk returns a tuple with the PublicationLegale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationLegale

`func (o *BodaccAchat) SetPublicationLegale(v string)`

SetPublicationLegale sets PublicationLegale field to given value.

### HasPublicationLegale

`func (o *BodaccAchat) HasPublicationLegale() bool`

HasPublicationLegale returns a boolean if a field has been set.

### GetDenominationAncienProprietaire

`func (o *BodaccAchat) GetDenominationAncienProprietaire() string`

GetDenominationAncienProprietaire returns the DenominationAncienProprietaire field if non-nil, zero value otherwise.

### GetDenominationAncienProprietaireOk

`func (o *BodaccAchat) GetDenominationAncienProprietaireOk() (*string, bool)`

GetDenominationAncienProprietaireOk returns a tuple with the DenominationAncienProprietaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationAncienProprietaire

`func (o *BodaccAchat) SetDenominationAncienProprietaire(v string)`

SetDenominationAncienProprietaire sets DenominationAncienProprietaire field to given value.

### HasDenominationAncienProprietaire

`func (o *BodaccAchat) HasDenominationAncienProprietaire() bool`

HasDenominationAncienProprietaire returns a boolean if a field has been set.

### GetSirenAncienProprietaire

`func (o *BodaccAchat) GetSirenAncienProprietaire() string`

GetSirenAncienProprietaire returns the SirenAncienProprietaire field if non-nil, zero value otherwise.

### GetSirenAncienProprietaireOk

`func (o *BodaccAchat) GetSirenAncienProprietaireOk() (*string, bool)`

GetSirenAncienProprietaireOk returns a tuple with the SirenAncienProprietaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenAncienProprietaire

`func (o *BodaccAchat) SetSirenAncienProprietaire(v string)`

SetSirenAncienProprietaire sets SirenAncienProprietaire field to given value.

### HasSirenAncienProprietaire

`func (o *BodaccAchat) HasSirenAncienProprietaire() bool`

HasSirenAncienProprietaire returns a boolean if a field has been set.

### GetDenominationAncienExploitant

`func (o *BodaccAchat) GetDenominationAncienExploitant() string`

GetDenominationAncienExploitant returns the DenominationAncienExploitant field if non-nil, zero value otherwise.

### GetDenominationAncienExploitantOk

`func (o *BodaccAchat) GetDenominationAncienExploitantOk() (*string, bool)`

GetDenominationAncienExploitantOk returns a tuple with the DenominationAncienExploitant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationAncienExploitant

`func (o *BodaccAchat) SetDenominationAncienExploitant(v string)`

SetDenominationAncienExploitant sets DenominationAncienExploitant field to given value.

### HasDenominationAncienExploitant

`func (o *BodaccAchat) HasDenominationAncienExploitant() bool`

HasDenominationAncienExploitant returns a boolean if a field has been set.

### GetSirenAncienExploitant

`func (o *BodaccAchat) GetSirenAncienExploitant() string`

GetSirenAncienExploitant returns the SirenAncienExploitant field if non-nil, zero value otherwise.

### GetSirenAncienExploitantOk

`func (o *BodaccAchat) GetSirenAncienExploitantOk() (*string, bool)`

GetSirenAncienExploitantOk returns a tuple with the SirenAncienExploitant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenAncienExploitant

`func (o *BodaccAchat) SetSirenAncienExploitant(v string)`

SetSirenAncienExploitant sets SirenAncienExploitant field to given value.

### HasSirenAncienExploitant

`func (o *BodaccAchat) HasSirenAncienExploitant() bool`

HasSirenAncienExploitant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


