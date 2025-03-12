# RepresentantRecherche

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualite** | Pointer to **string** | Qualité du représentant. | [optional] 
**PersonneMorale** | Pointer to **bool** | Vrai si le représentant est une personne morale, faux si personne physique. | [optional] 
**DatePriseDePoste** | Pointer to **string** | Date de prise de poste du représentant. | [optional] 
**Denomination** | Pointer to **string** | Dénomination du représentant si personne morale. | [optional] 
**Siren** | Pointer to **string** | Siren du représentant si personne morale. | [optional] 
**FormeJuridique** | Pointer to **string** | Forme juridique du représentant dans le cas d&#39;une personne morale. | [optional] 
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
**Actuel** | Pointer to **bool** | Vaut vrai si le représentant est toujours à son poste. | [optional] 
**DateDepartDePoste** | Pointer to **string** | Date de départ de poste dans le cas où le représentant n&#39;est plus un représentant actuel, au format AAAA-MM-JJ. | [optional] 

## Methods

### NewRepresentantRecherche

`func NewRepresentantRecherche() *RepresentantRecherche`

NewRepresentantRecherche instantiates a new RepresentantRecherche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepresentantRechercheWithDefaults

`func NewRepresentantRechercheWithDefaults() *RepresentantRecherche`

NewRepresentantRechercheWithDefaults instantiates a new RepresentantRecherche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualite

`func (o *RepresentantRecherche) GetQualite() string`

GetQualite returns the Qualite field if non-nil, zero value otherwise.

### GetQualiteOk

`func (o *RepresentantRecherche) GetQualiteOk() (*string, bool)`

GetQualiteOk returns a tuple with the Qualite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualite

`func (o *RepresentantRecherche) SetQualite(v string)`

SetQualite sets Qualite field to given value.

### HasQualite

`func (o *RepresentantRecherche) HasQualite() bool`

HasQualite returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *RepresentantRecherche) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *RepresentantRecherche) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *RepresentantRecherche) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *RepresentantRecherche) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDatePriseDePoste

`func (o *RepresentantRecherche) GetDatePriseDePoste() string`

GetDatePriseDePoste returns the DatePriseDePoste field if non-nil, zero value otherwise.

### GetDatePriseDePosteOk

`func (o *RepresentantRecherche) GetDatePriseDePosteOk() (*string, bool)`

GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePriseDePoste

`func (o *RepresentantRecherche) SetDatePriseDePoste(v string)`

SetDatePriseDePoste sets DatePriseDePoste field to given value.

### HasDatePriseDePoste

`func (o *RepresentantRecherche) HasDatePriseDePoste() bool`

HasDatePriseDePoste returns a boolean if a field has been set.

### GetDenomination

`func (o *RepresentantRecherche) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *RepresentantRecherche) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *RepresentantRecherche) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *RepresentantRecherche) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiren

`func (o *RepresentantRecherche) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *RepresentantRecherche) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *RepresentantRecherche) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *RepresentantRecherche) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *RepresentantRecherche) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *RepresentantRecherche) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *RepresentantRecherche) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *RepresentantRecherche) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetSexe

`func (o *RepresentantRecherche) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *RepresentantRecherche) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *RepresentantRecherche) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *RepresentantRecherche) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetNom

`func (o *RepresentantRecherche) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *RepresentantRecherche) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *RepresentantRecherche) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *RepresentantRecherche) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *RepresentantRecherche) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *RepresentantRecherche) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *RepresentantRecherche) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *RepresentantRecherche) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *RepresentantRecherche) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *RepresentantRecherche) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *RepresentantRecherche) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *RepresentantRecherche) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetNomComplet

`func (o *RepresentantRecherche) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *RepresentantRecherche) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *RepresentantRecherche) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *RepresentantRecherche) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *RepresentantRecherche) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *RepresentantRecherche) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *RepresentantRecherche) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *RepresentantRecherche) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *RepresentantRecherche) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *RepresentantRecherche) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *RepresentantRecherche) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *RepresentantRecherche) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetAge

`func (o *RepresentantRecherche) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RepresentantRecherche) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RepresentantRecherche) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *RepresentantRecherche) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetNationalite

`func (o *RepresentantRecherche) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *RepresentantRecherche) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *RepresentantRecherche) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *RepresentantRecherche) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *RepresentantRecherche) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *RepresentantRecherche) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *RepresentantRecherche) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *RepresentantRecherche) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *RepresentantRecherche) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *RepresentantRecherche) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *RepresentantRecherche) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *RepresentantRecherche) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *RepresentantRecherche) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *RepresentantRecherche) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *RepresentantRecherche) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *RepresentantRecherche) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *RepresentantRecherche) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *RepresentantRecherche) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *RepresentantRecherche) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *RepresentantRecherche) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *RepresentantRecherche) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *RepresentantRecherche) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *RepresentantRecherche) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *RepresentantRecherche) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *RepresentantRecherche) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *RepresentantRecherche) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *RepresentantRecherche) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *RepresentantRecherche) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *RepresentantRecherche) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *RepresentantRecherche) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *RepresentantRecherche) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *RepresentantRecherche) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *RepresentantRecherche) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *RepresentantRecherche) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *RepresentantRecherche) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *RepresentantRecherche) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *RepresentantRecherche) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *RepresentantRecherche) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *RepresentantRecherche) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *RepresentantRecherche) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *RepresentantRecherche) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *RepresentantRecherche) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *RepresentantRecherche) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *RepresentantRecherche) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *RepresentantRecherche) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *RepresentantRecherche) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *RepresentantRecherche) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *RepresentantRecherche) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetActuel

`func (o *RepresentantRecherche) GetActuel() bool`

GetActuel returns the Actuel field if non-nil, zero value otherwise.

### GetActuelOk

`func (o *RepresentantRecherche) GetActuelOk() (*bool, bool)`

GetActuelOk returns a tuple with the Actuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActuel

`func (o *RepresentantRecherche) SetActuel(v bool)`

SetActuel sets Actuel field to given value.

### HasActuel

`func (o *RepresentantRecherche) HasActuel() bool`

HasActuel returns a boolean if a field has been set.

### GetDateDepartDePoste

`func (o *RepresentantRecherche) GetDateDepartDePoste() string`

GetDateDepartDePoste returns the DateDepartDePoste field if non-nil, zero value otherwise.

### GetDateDepartDePosteOk

`func (o *RepresentantRecherche) GetDateDepartDePosteOk() (*string, bool)`

GetDateDepartDePosteOk returns a tuple with the DateDepartDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepartDePoste

`func (o *RepresentantRecherche) SetDateDepartDePoste(v string)`

SetDateDepartDePoste sets DateDepartDePoste field to given value.

### HasDateDepartDePoste

`func (o *RepresentantRecherche) HasDateDepartDePoste() bool`

HasDateDepartDePoste returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


