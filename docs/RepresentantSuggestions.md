# RepresentantSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualite** | Pointer to **string** | Qualité du représentant. | [optional] 
**PersonneMorale** | Pointer to **bool** | Vrai si le représentant est une personne morale, faux si personne physique. | [optional] 
**DatePriseDePoste** | Pointer to **string** | Date de prise de poste du représentant. | [optional] 
**Denomination** | Pointer to **string** | Dénomination du représentant si personne morale. | [optional] 
**Siren** | Pointer to **string** | Siren du représentant si personne morale. | [optional] 
**FormeJuridique** | Pointer to **string** | Forme juridique du représentant si personne morale. | [optional] 
**Sexe** | Pointer to **string** | Sexe supposé du représentant si personne physique. F pour féminin, M pour masculin. Ce champ est estimé à partir du prénom du représentant. | [optional] 
**Nom** | Pointer to **string** | Nom du représentant. | [optional] 
**Prenom** | Pointer to **string** | Prénoms du représentant. | [optional] 
**PrenomUsuel** | Pointer to **string** | Prénom usuel du représentant. | [optional] 
**NomComplet** | Pointer to **string** | Nom complet du représentant. | [optional] 
**DateDeNaissance** | Pointer to **string** | Date de naissance du représentant. | [optional] 
**DateDeNaissanceFormate** | Pointer to **string** | Date de naissance formatée du représentant. | [optional] 
**Age** | Pointer to **int32** | Age du représentant. | [optional] 
**Nationalite** | Pointer to **string** | Nationalité du représentant. | [optional] 
**CodeNationalite** | Pointer to **string** | Code nationalité du représentant | [optional] 
**VilleDeNaissance** | Pointer to **string** | Ville de naissance du représentant. | [optional] 
**PaysDeNaissance** | Pointer to **string** | Pays de naissance du représentant. | [optional] 
**CodePaysDeNaissance** | Pointer to **string** | Code du pays de naissance du représentant. | [optional] 
**AdresseLigne1** | Pointer to **string** | Première ligne de l&#39;adresse du représentant. | [optional] 
**AdresseLigne2** | Pointer to **string** | Deuxième ligne de l&#39;adresse du représentant. | [optional] 
**AdresseLigne3** | Pointer to **string** | Troisième ligne de l&#39;adresse du représentant. | [optional] 
**CodePostal** | Pointer to **string** | Code postal du représentant. | [optional] 
**Ville** | Pointer to **string** | Ville du représentant. | [optional] 
**Pays** | Pointer to **string** | Pays du représentant. | [optional] 
**CodePays** | Pointer to **string** | Code du pays du représentant | [optional] 

## Methods

### NewRepresentantSuggestions

`func NewRepresentantSuggestions() *RepresentantSuggestions`

NewRepresentantSuggestions instantiates a new RepresentantSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepresentantSuggestionsWithDefaults

`func NewRepresentantSuggestionsWithDefaults() *RepresentantSuggestions`

NewRepresentantSuggestionsWithDefaults instantiates a new RepresentantSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualite

`func (o *RepresentantSuggestions) GetQualite() string`

GetQualite returns the Qualite field if non-nil, zero value otherwise.

### GetQualiteOk

`func (o *RepresentantSuggestions) GetQualiteOk() (*string, bool)`

GetQualiteOk returns a tuple with the Qualite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualite

`func (o *RepresentantSuggestions) SetQualite(v string)`

SetQualite sets Qualite field to given value.

### HasQualite

`func (o *RepresentantSuggestions) HasQualite() bool`

HasQualite returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *RepresentantSuggestions) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *RepresentantSuggestions) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *RepresentantSuggestions) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *RepresentantSuggestions) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDatePriseDePoste

`func (o *RepresentantSuggestions) GetDatePriseDePoste() string`

GetDatePriseDePoste returns the DatePriseDePoste field if non-nil, zero value otherwise.

### GetDatePriseDePosteOk

`func (o *RepresentantSuggestions) GetDatePriseDePosteOk() (*string, bool)`

GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePriseDePoste

`func (o *RepresentantSuggestions) SetDatePriseDePoste(v string)`

SetDatePriseDePoste sets DatePriseDePoste field to given value.

### HasDatePriseDePoste

`func (o *RepresentantSuggestions) HasDatePriseDePoste() bool`

HasDatePriseDePoste returns a boolean if a field has been set.

### GetDenomination

`func (o *RepresentantSuggestions) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *RepresentantSuggestions) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *RepresentantSuggestions) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *RepresentantSuggestions) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiren

`func (o *RepresentantSuggestions) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *RepresentantSuggestions) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *RepresentantSuggestions) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *RepresentantSuggestions) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *RepresentantSuggestions) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *RepresentantSuggestions) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *RepresentantSuggestions) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *RepresentantSuggestions) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetSexe

`func (o *RepresentantSuggestions) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *RepresentantSuggestions) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *RepresentantSuggestions) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *RepresentantSuggestions) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetNom

`func (o *RepresentantSuggestions) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *RepresentantSuggestions) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *RepresentantSuggestions) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *RepresentantSuggestions) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *RepresentantSuggestions) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *RepresentantSuggestions) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *RepresentantSuggestions) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *RepresentantSuggestions) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *RepresentantSuggestions) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *RepresentantSuggestions) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *RepresentantSuggestions) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *RepresentantSuggestions) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetNomComplet

`func (o *RepresentantSuggestions) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *RepresentantSuggestions) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *RepresentantSuggestions) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *RepresentantSuggestions) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *RepresentantSuggestions) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *RepresentantSuggestions) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *RepresentantSuggestions) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *RepresentantSuggestions) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *RepresentantSuggestions) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *RepresentantSuggestions) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *RepresentantSuggestions) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *RepresentantSuggestions) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetAge

`func (o *RepresentantSuggestions) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RepresentantSuggestions) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RepresentantSuggestions) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *RepresentantSuggestions) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetNationalite

`func (o *RepresentantSuggestions) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *RepresentantSuggestions) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *RepresentantSuggestions) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *RepresentantSuggestions) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *RepresentantSuggestions) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *RepresentantSuggestions) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *RepresentantSuggestions) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *RepresentantSuggestions) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *RepresentantSuggestions) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *RepresentantSuggestions) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *RepresentantSuggestions) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *RepresentantSuggestions) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *RepresentantSuggestions) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *RepresentantSuggestions) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *RepresentantSuggestions) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *RepresentantSuggestions) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *RepresentantSuggestions) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *RepresentantSuggestions) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *RepresentantSuggestions) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *RepresentantSuggestions) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *RepresentantSuggestions) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *RepresentantSuggestions) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *RepresentantSuggestions) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *RepresentantSuggestions) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *RepresentantSuggestions) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *RepresentantSuggestions) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *RepresentantSuggestions) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *RepresentantSuggestions) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *RepresentantSuggestions) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *RepresentantSuggestions) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *RepresentantSuggestions) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *RepresentantSuggestions) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *RepresentantSuggestions) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *RepresentantSuggestions) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *RepresentantSuggestions) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *RepresentantSuggestions) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *RepresentantSuggestions) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *RepresentantSuggestions) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *RepresentantSuggestions) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *RepresentantSuggestions) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *RepresentantSuggestions) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *RepresentantSuggestions) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *RepresentantSuggestions) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *RepresentantSuggestions) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *RepresentantSuggestions) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *RepresentantSuggestions) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *RepresentantSuggestions) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *RepresentantSuggestions) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


