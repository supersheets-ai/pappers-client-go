# BodaccVente

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
**DenominationNouveauProprietaire** | Pointer to **string** | Dénomination du nouveau propriétaire de l&#39;établisement. | [optional] 
**SirenNouveauProprietaire** | Pointer to **string** | Siren du nouveau propriétaire de l&#39;établisement. | [optional] 
**DenominationNouvelExploitant** | Pointer to **string** | Dénomination du nouvel exploitant de l&#39;établisement. | [optional] 
**SirenNouvelExploitant** | Pointer to **string** | Siren du nouvel exploitant de l&#39;établisement. | [optional] 

## Methods

### NewBodaccVente

`func NewBodaccVente() *BodaccVente`

NewBodaccVente instantiates a new BodaccVente object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccVenteWithDefaults

`func NewBodaccVenteWithDefaults() *BodaccVente`

NewBodaccVenteWithDefaults instantiates a new BodaccVente object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomEntreprise

`func (o *BodaccVente) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *BodaccVente) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *BodaccVente) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *BodaccVente) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *BodaccVente) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *BodaccVente) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *BodaccVente) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *BodaccVente) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *BodaccVente) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *BodaccVente) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *BodaccVente) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *BodaccVente) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *BodaccVente) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *BodaccVente) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *BodaccVente) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *BodaccVente) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *BodaccVente) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *BodaccVente) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *BodaccVente) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *BodaccVente) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetAdministration

`func (o *BodaccVente) GetAdministration() string`

GetAdministration returns the Administration field if non-nil, zero value otherwise.

### GetAdministrationOk

`func (o *BodaccVente) GetAdministrationOk() (*string, bool)`

GetAdministrationOk returns a tuple with the Administration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministration

`func (o *BodaccVente) SetAdministration(v string)`

SetAdministration sets Administration field to given value.

### HasAdministration

`func (o *BodaccVente) HasAdministration() bool`

HasAdministration returns a boolean if a field has been set.

### GetAdresse

`func (o *BodaccVente) GetAdresse() string`

GetAdresse returns the Adresse field if non-nil, zero value otherwise.

### GetAdresseOk

`func (o *BodaccVente) GetAdresseOk() (*string, bool)`

GetAdresseOk returns a tuple with the Adresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresse

`func (o *BodaccVente) SetAdresse(v string)`

SetAdresse sets Adresse field to given value.

### HasAdresse

`func (o *BodaccVente) HasAdresse() bool`

HasAdresse returns a boolean if a field has been set.

### GetCommentaires

`func (o *BodaccVente) GetCommentaires() string`

GetCommentaires returns the Commentaires field if non-nil, zero value otherwise.

### GetCommentairesOk

`func (o *BodaccVente) GetCommentairesOk() (*string, bool)`

GetCommentairesOk returns a tuple with the Commentaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentaires

`func (o *BodaccVente) SetCommentaires(v string)`

SetCommentaires sets Commentaires field to given value.

### HasCommentaires

`func (o *BodaccVente) HasCommentaires() bool`

HasCommentaires returns a boolean if a field has been set.

### GetOppositions

`func (o *BodaccVente) GetOppositions() string`

GetOppositions returns the Oppositions field if non-nil, zero value otherwise.

### GetOppositionsOk

`func (o *BodaccVente) GetOppositionsOk() (*string, bool)`

GetOppositionsOk returns a tuple with the Oppositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositions

`func (o *BodaccVente) SetOppositions(v string)`

SetOppositions sets Oppositions field to given value.

### HasOppositions

`func (o *BodaccVente) HasOppositions() bool`

HasOppositions returns a boolean if a field has been set.

### GetDeclarationCreance

`func (o *BodaccVente) GetDeclarationCreance() string`

GetDeclarationCreance returns the DeclarationCreance field if non-nil, zero value otherwise.

### GetDeclarationCreanceOk

`func (o *BodaccVente) GetDeclarationCreanceOk() (*string, bool)`

GetDeclarationCreanceOk returns a tuple with the DeclarationCreance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationCreance

`func (o *BodaccVente) SetDeclarationCreance(v string)`

SetDeclarationCreance sets DeclarationCreance field to given value.

### HasDeclarationCreance

`func (o *BodaccVente) HasDeclarationCreance() bool`

HasDeclarationCreance returns a boolean if a field has been set.

### GetPublicationLegale

`func (o *BodaccVente) GetPublicationLegale() string`

GetPublicationLegale returns the PublicationLegale field if non-nil, zero value otherwise.

### GetPublicationLegaleOk

`func (o *BodaccVente) GetPublicationLegaleOk() (*string, bool)`

GetPublicationLegaleOk returns a tuple with the PublicationLegale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationLegale

`func (o *BodaccVente) SetPublicationLegale(v string)`

SetPublicationLegale sets PublicationLegale field to given value.

### HasPublicationLegale

`func (o *BodaccVente) HasPublicationLegale() bool`

HasPublicationLegale returns a boolean if a field has been set.

### GetDenominationNouveauProprietaire

`func (o *BodaccVente) GetDenominationNouveauProprietaire() string`

GetDenominationNouveauProprietaire returns the DenominationNouveauProprietaire field if non-nil, zero value otherwise.

### GetDenominationNouveauProprietaireOk

`func (o *BodaccVente) GetDenominationNouveauProprietaireOk() (*string, bool)`

GetDenominationNouveauProprietaireOk returns a tuple with the DenominationNouveauProprietaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationNouveauProprietaire

`func (o *BodaccVente) SetDenominationNouveauProprietaire(v string)`

SetDenominationNouveauProprietaire sets DenominationNouveauProprietaire field to given value.

### HasDenominationNouveauProprietaire

`func (o *BodaccVente) HasDenominationNouveauProprietaire() bool`

HasDenominationNouveauProprietaire returns a boolean if a field has been set.

### GetSirenNouveauProprietaire

`func (o *BodaccVente) GetSirenNouveauProprietaire() string`

GetSirenNouveauProprietaire returns the SirenNouveauProprietaire field if non-nil, zero value otherwise.

### GetSirenNouveauProprietaireOk

`func (o *BodaccVente) GetSirenNouveauProprietaireOk() (*string, bool)`

GetSirenNouveauProprietaireOk returns a tuple with the SirenNouveauProprietaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenNouveauProprietaire

`func (o *BodaccVente) SetSirenNouveauProprietaire(v string)`

SetSirenNouveauProprietaire sets SirenNouveauProprietaire field to given value.

### HasSirenNouveauProprietaire

`func (o *BodaccVente) HasSirenNouveauProprietaire() bool`

HasSirenNouveauProprietaire returns a boolean if a field has been set.

### GetDenominationNouvelExploitant

`func (o *BodaccVente) GetDenominationNouvelExploitant() string`

GetDenominationNouvelExploitant returns the DenominationNouvelExploitant field if non-nil, zero value otherwise.

### GetDenominationNouvelExploitantOk

`func (o *BodaccVente) GetDenominationNouvelExploitantOk() (*string, bool)`

GetDenominationNouvelExploitantOk returns a tuple with the DenominationNouvelExploitant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationNouvelExploitant

`func (o *BodaccVente) SetDenominationNouvelExploitant(v string)`

SetDenominationNouvelExploitant sets DenominationNouvelExploitant field to given value.

### HasDenominationNouvelExploitant

`func (o *BodaccVente) HasDenominationNouvelExploitant() bool`

HasDenominationNouvelExploitant returns a boolean if a field has been set.

### GetSirenNouvelExploitant

`func (o *BodaccVente) GetSirenNouvelExploitant() string`

GetSirenNouvelExploitant returns the SirenNouvelExploitant field if non-nil, zero value otherwise.

### GetSirenNouvelExploitantOk

`func (o *BodaccVente) GetSirenNouvelExploitantOk() (*string, bool)`

GetSirenNouvelExploitantOk returns a tuple with the SirenNouvelExploitant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenNouvelExploitant

`func (o *BodaccVente) SetSirenNouvelExploitant(v string)`

SetSirenNouvelExploitant sets SirenNouvelExploitant field to given value.

### HasSirenNouvelExploitant

`func (o *BodaccVente) HasSirenNouvelExploitant() bool`

HasSirenNouvelExploitant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


