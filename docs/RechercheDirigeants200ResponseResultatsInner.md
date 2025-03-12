# RechercheDirigeants200ResponseResultatsInner

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
**Entreprises** | Pointer to [**[]EntrepriseRecherche**](EntrepriseRecherche.md) | Liste des entreprises du dirigeant, dans la limite de 100 entreprises. | [optional] 
**NbEntreprisesTotal** | Pointer to **int32** | Nombre d&#39;entreprises du dirigeant au total | [optional] 

## Methods

### NewRechercheDirigeants200ResponseResultatsInner

`func NewRechercheDirigeants200ResponseResultatsInner() *RechercheDirigeants200ResponseResultatsInner`

NewRechercheDirigeants200ResponseResultatsInner instantiates a new RechercheDirigeants200ResponseResultatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercheDirigeants200ResponseResultatsInnerWithDefaults

`func NewRechercheDirigeants200ResponseResultatsInnerWithDefaults() *RechercheDirigeants200ResponseResultatsInner`

NewRechercheDirigeants200ResponseResultatsInnerWithDefaults instantiates a new RechercheDirigeants200ResponseResultatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualite

`func (o *RechercheDirigeants200ResponseResultatsInner) GetQualite() string`

GetQualite returns the Qualite field if non-nil, zero value otherwise.

### GetQualiteOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetQualiteOk() (*string, bool)`

GetQualiteOk returns a tuple with the Qualite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualite

`func (o *RechercheDirigeants200ResponseResultatsInner) SetQualite(v string)`

SetQualite sets Qualite field to given value.

### HasQualite

`func (o *RechercheDirigeants200ResponseResultatsInner) HasQualite() bool`

HasQualite returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *RechercheDirigeants200ResponseResultatsInner) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *RechercheDirigeants200ResponseResultatsInner) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDatePriseDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDatePriseDePoste() string`

GetDatePriseDePoste returns the DatePriseDePoste field if non-nil, zero value otherwise.

### GetDatePriseDePosteOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDatePriseDePosteOk() (*string, bool)`

GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePriseDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) SetDatePriseDePoste(v string)`

SetDatePriseDePoste sets DatePriseDePoste field to given value.

### HasDatePriseDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) HasDatePriseDePoste() bool`

HasDatePriseDePoste returns a boolean if a field has been set.

### GetDenomination

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *RechercheDirigeants200ResponseResultatsInner) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *RechercheDirigeants200ResponseResultatsInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiren

`func (o *RechercheDirigeants200ResponseResultatsInner) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *RechercheDirigeants200ResponseResultatsInner) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *RechercheDirigeants200ResponseResultatsInner) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *RechercheDirigeants200ResponseResultatsInner) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *RechercheDirigeants200ResponseResultatsInner) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *RechercheDirigeants200ResponseResultatsInner) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetSexe

`func (o *RechercheDirigeants200ResponseResultatsInner) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *RechercheDirigeants200ResponseResultatsInner) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *RechercheDirigeants200ResponseResultatsInner) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetNom

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *RechercheDirigeants200ResponseResultatsInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *RechercheDirigeants200ResponseResultatsInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *RechercheDirigeants200ResponseResultatsInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *RechercheDirigeants200ResponseResultatsInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *RechercheDirigeants200ResponseResultatsInner) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *RechercheDirigeants200ResponseResultatsInner) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetNomComplet

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *RechercheDirigeants200ResponseResultatsInner) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *RechercheDirigeants200ResponseResultatsInner) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *RechercheDirigeants200ResponseResultatsInner) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *RechercheDirigeants200ResponseResultatsInner) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetAge

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RechercheDirigeants200ResponseResultatsInner) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *RechercheDirigeants200ResponseResultatsInner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *RechercheDirigeants200ResponseResultatsInner) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *RechercheDirigeants200ResponseResultatsInner) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *RechercheDirigeants200ResponseResultatsInner) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *RechercheDirigeants200ResponseResultatsInner) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *RechercheDirigeants200ResponseResultatsInner) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *RechercheDirigeants200ResponseResultatsInner) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *RechercheDirigeants200ResponseResultatsInner) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *RechercheDirigeants200ResponseResultatsInner) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *RechercheDirigeants200ResponseResultatsInner) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *RechercheDirigeants200ResponseResultatsInner) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *RechercheDirigeants200ResponseResultatsInner) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *RechercheDirigeants200ResponseResultatsInner) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *RechercheDirigeants200ResponseResultatsInner) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *RechercheDirigeants200ResponseResultatsInner) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *RechercheDirigeants200ResponseResultatsInner) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *RechercheDirigeants200ResponseResultatsInner) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *RechercheDirigeants200ResponseResultatsInner) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *RechercheDirigeants200ResponseResultatsInner) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetActuel

`func (o *RechercheDirigeants200ResponseResultatsInner) GetActuel() bool`

GetActuel returns the Actuel field if non-nil, zero value otherwise.

### GetActuelOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetActuelOk() (*bool, bool)`

GetActuelOk returns a tuple with the Actuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActuel

`func (o *RechercheDirigeants200ResponseResultatsInner) SetActuel(v bool)`

SetActuel sets Actuel field to given value.

### HasActuel

`func (o *RechercheDirigeants200ResponseResultatsInner) HasActuel() bool`

HasActuel returns a boolean if a field has been set.

### GetDateDepartDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDepartDePoste() string`

GetDateDepartDePoste returns the DateDepartDePoste field if non-nil, zero value otherwise.

### GetDateDepartDePosteOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetDateDepartDePosteOk() (*string, bool)`

GetDateDepartDePosteOk returns a tuple with the DateDepartDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepartDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) SetDateDepartDePoste(v string)`

SetDateDepartDePoste sets DateDepartDePoste field to given value.

### HasDateDepartDePoste

`func (o *RechercheDirigeants200ResponseResultatsInner) HasDateDepartDePoste() bool`

HasDateDepartDePoste returns a boolean if a field has been set.

### GetEntreprises

`func (o *RechercheDirigeants200ResponseResultatsInner) GetEntreprises() []EntrepriseRecherche`

GetEntreprises returns the Entreprises field if non-nil, zero value otherwise.

### GetEntreprisesOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetEntreprisesOk() (*[]EntrepriseRecherche, bool)`

GetEntreprisesOk returns a tuple with the Entreprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntreprises

`func (o *RechercheDirigeants200ResponseResultatsInner) SetEntreprises(v []EntrepriseRecherche)`

SetEntreprises sets Entreprises field to given value.

### HasEntreprises

`func (o *RechercheDirigeants200ResponseResultatsInner) HasEntreprises() bool`

HasEntreprises returns a boolean if a field has been set.

### GetNbEntreprisesTotal

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNbEntreprisesTotal() int32`

GetNbEntreprisesTotal returns the NbEntreprisesTotal field if non-nil, zero value otherwise.

### GetNbEntreprisesTotalOk

`func (o *RechercheDirigeants200ResponseResultatsInner) GetNbEntreprisesTotalOk() (*int32, bool)`

GetNbEntreprisesTotalOk returns a tuple with the NbEntreprisesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbEntreprisesTotal

`func (o *RechercheDirigeants200ResponseResultatsInner) SetNbEntreprisesTotal(v int32)`

SetNbEntreprisesTotal sets NbEntreprisesTotal field to given value.

### HasNbEntreprisesTotal

`func (o *RechercheDirigeants200ResponseResultatsInner) HasNbEntreprisesTotal() bool`

HasNbEntreprisesTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


