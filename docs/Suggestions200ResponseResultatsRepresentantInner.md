# Suggestions200ResponseResultatsRepresentantInner

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
**Mention** | Pointer to **string** | Nom du dirigeant, avec le texte recherché encerclé d&#39;une balise HTML &#x60;&lt;em&gt;&#x60;. | [optional] 

## Methods

### NewSuggestions200ResponseResultatsRepresentantInner

`func NewSuggestions200ResponseResultatsRepresentantInner() *Suggestions200ResponseResultatsRepresentantInner`

NewSuggestions200ResponseResultatsRepresentantInner instantiates a new Suggestions200ResponseResultatsRepresentantInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestions200ResponseResultatsRepresentantInnerWithDefaults

`func NewSuggestions200ResponseResultatsRepresentantInnerWithDefaults() *Suggestions200ResponseResultatsRepresentantInner`

NewSuggestions200ResponseResultatsRepresentantInnerWithDefaults instantiates a new Suggestions200ResponseResultatsRepresentantInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualite

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetQualite() string`

GetQualite returns the Qualite field if non-nil, zero value otherwise.

### GetQualiteOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetQualiteOk() (*string, bool)`

GetQualiteOk returns a tuple with the Qualite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualite

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetQualite(v string)`

SetQualite sets Qualite field to given value.

### HasQualite

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasQualite() bool`

HasQualite returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDatePriseDePoste

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDatePriseDePoste() string`

GetDatePriseDePoste returns the DatePriseDePoste field if non-nil, zero value otherwise.

### GetDatePriseDePosteOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDatePriseDePosteOk() (*string, bool)`

GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePriseDePoste

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetDatePriseDePoste(v string)`

SetDatePriseDePoste sets DatePriseDePoste field to given value.

### HasDatePriseDePoste

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasDatePriseDePoste() bool`

HasDatePriseDePoste returns a boolean if a field has been set.

### GetDenomination

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiren

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetSexe

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetNom

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetNomComplet

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetAge

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetMention

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetMention() string`

GetMention returns the Mention field if non-nil, zero value otherwise.

### GetMentionOk

`func (o *Suggestions200ResponseResultatsRepresentantInner) GetMentionOk() (*string, bool)`

GetMentionOk returns a tuple with the Mention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMention

`func (o *Suggestions200ResponseResultatsRepresentantInner) SetMention(v string)`

SetMention sets Mention field to given value.

### HasMention

`func (o *Suggestions200ResponseResultatsRepresentantInner) HasMention() bool`

HasMention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


