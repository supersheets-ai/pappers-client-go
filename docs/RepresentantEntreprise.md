# RepresentantEntreprise

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
**PersonnePolitiquementExposee** | Pointer to [**PersonnePolitiquementExposee**](PersonnePolitiquementExposee.md) |  | [optional] 
**SanctionsEnCours** | Pointer to **bool** | Vaut vrai si le représentant est actuellement sous sanction. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Sanctions** | Pointer to [**[]Sanction**](Sanction.md) | Liste des sanctions du représentant. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Decede** | Pointer to **bool** | Indique si une personne avec le même nom et la même date de naissance est présente dans le fichier des personnes décédées. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**DateDeDeces** | Pointer to **string** | Indique, si le champ decede est vrai, la date de décès précisée dans le fichier des personnes décédées. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 

## Methods

### NewRepresentantEntreprise

`func NewRepresentantEntreprise() *RepresentantEntreprise`

NewRepresentantEntreprise instantiates a new RepresentantEntreprise object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepresentantEntrepriseWithDefaults

`func NewRepresentantEntrepriseWithDefaults() *RepresentantEntreprise`

NewRepresentantEntrepriseWithDefaults instantiates a new RepresentantEntreprise object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualite

`func (o *RepresentantEntreprise) GetQualite() string`

GetQualite returns the Qualite field if non-nil, zero value otherwise.

### GetQualiteOk

`func (o *RepresentantEntreprise) GetQualiteOk() (*string, bool)`

GetQualiteOk returns a tuple with the Qualite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualite

`func (o *RepresentantEntreprise) SetQualite(v string)`

SetQualite sets Qualite field to given value.

### HasQualite

`func (o *RepresentantEntreprise) HasQualite() bool`

HasQualite returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *RepresentantEntreprise) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *RepresentantEntreprise) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *RepresentantEntreprise) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *RepresentantEntreprise) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDatePriseDePoste

`func (o *RepresentantEntreprise) GetDatePriseDePoste() string`

GetDatePriseDePoste returns the DatePriseDePoste field if non-nil, zero value otherwise.

### GetDatePriseDePosteOk

`func (o *RepresentantEntreprise) GetDatePriseDePosteOk() (*string, bool)`

GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePriseDePoste

`func (o *RepresentantEntreprise) SetDatePriseDePoste(v string)`

SetDatePriseDePoste sets DatePriseDePoste field to given value.

### HasDatePriseDePoste

`func (o *RepresentantEntreprise) HasDatePriseDePoste() bool`

HasDatePriseDePoste returns a boolean if a field has been set.

### GetDenomination

`func (o *RepresentantEntreprise) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *RepresentantEntreprise) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *RepresentantEntreprise) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *RepresentantEntreprise) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiren

`func (o *RepresentantEntreprise) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *RepresentantEntreprise) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *RepresentantEntreprise) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *RepresentantEntreprise) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *RepresentantEntreprise) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *RepresentantEntreprise) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *RepresentantEntreprise) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *RepresentantEntreprise) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetSexe

`func (o *RepresentantEntreprise) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *RepresentantEntreprise) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *RepresentantEntreprise) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *RepresentantEntreprise) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetNom

`func (o *RepresentantEntreprise) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *RepresentantEntreprise) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *RepresentantEntreprise) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *RepresentantEntreprise) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *RepresentantEntreprise) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *RepresentantEntreprise) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *RepresentantEntreprise) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *RepresentantEntreprise) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *RepresentantEntreprise) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *RepresentantEntreprise) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *RepresentantEntreprise) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *RepresentantEntreprise) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetNomComplet

`func (o *RepresentantEntreprise) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *RepresentantEntreprise) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *RepresentantEntreprise) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *RepresentantEntreprise) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *RepresentantEntreprise) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *RepresentantEntreprise) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *RepresentantEntreprise) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *RepresentantEntreprise) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *RepresentantEntreprise) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *RepresentantEntreprise) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *RepresentantEntreprise) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *RepresentantEntreprise) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetAge

`func (o *RepresentantEntreprise) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RepresentantEntreprise) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RepresentantEntreprise) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *RepresentantEntreprise) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetNationalite

`func (o *RepresentantEntreprise) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *RepresentantEntreprise) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *RepresentantEntreprise) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *RepresentantEntreprise) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *RepresentantEntreprise) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *RepresentantEntreprise) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *RepresentantEntreprise) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *RepresentantEntreprise) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *RepresentantEntreprise) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *RepresentantEntreprise) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *RepresentantEntreprise) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *RepresentantEntreprise) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *RepresentantEntreprise) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *RepresentantEntreprise) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *RepresentantEntreprise) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *RepresentantEntreprise) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *RepresentantEntreprise) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *RepresentantEntreprise) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *RepresentantEntreprise) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *RepresentantEntreprise) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *RepresentantEntreprise) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *RepresentantEntreprise) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *RepresentantEntreprise) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *RepresentantEntreprise) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *RepresentantEntreprise) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *RepresentantEntreprise) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *RepresentantEntreprise) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *RepresentantEntreprise) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *RepresentantEntreprise) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *RepresentantEntreprise) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *RepresentantEntreprise) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *RepresentantEntreprise) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *RepresentantEntreprise) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *RepresentantEntreprise) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *RepresentantEntreprise) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *RepresentantEntreprise) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *RepresentantEntreprise) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *RepresentantEntreprise) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *RepresentantEntreprise) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *RepresentantEntreprise) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *RepresentantEntreprise) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *RepresentantEntreprise) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *RepresentantEntreprise) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *RepresentantEntreprise) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *RepresentantEntreprise) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *RepresentantEntreprise) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *RepresentantEntreprise) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *RepresentantEntreprise) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetPersonnePolitiquementExposee

`func (o *RepresentantEntreprise) GetPersonnePolitiquementExposee() PersonnePolitiquementExposee`

GetPersonnePolitiquementExposee returns the PersonnePolitiquementExposee field if non-nil, zero value otherwise.

### GetPersonnePolitiquementExposeeOk

`func (o *RepresentantEntreprise) GetPersonnePolitiquementExposeeOk() (*PersonnePolitiquementExposee, bool)`

GetPersonnePolitiquementExposeeOk returns a tuple with the PersonnePolitiquementExposee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonnePolitiquementExposee

`func (o *RepresentantEntreprise) SetPersonnePolitiquementExposee(v PersonnePolitiquementExposee)`

SetPersonnePolitiquementExposee sets PersonnePolitiquementExposee field to given value.

### HasPersonnePolitiquementExposee

`func (o *RepresentantEntreprise) HasPersonnePolitiquementExposee() bool`

HasPersonnePolitiquementExposee returns a boolean if a field has been set.

### GetSanctionsEnCours

`func (o *RepresentantEntreprise) GetSanctionsEnCours() bool`

GetSanctionsEnCours returns the SanctionsEnCours field if non-nil, zero value otherwise.

### GetSanctionsEnCoursOk

`func (o *RepresentantEntreprise) GetSanctionsEnCoursOk() (*bool, bool)`

GetSanctionsEnCoursOk returns a tuple with the SanctionsEnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsEnCours

`func (o *RepresentantEntreprise) SetSanctionsEnCours(v bool)`

SetSanctionsEnCours sets SanctionsEnCours field to given value.

### HasSanctionsEnCours

`func (o *RepresentantEntreprise) HasSanctionsEnCours() bool`

HasSanctionsEnCours returns a boolean if a field has been set.

### GetSanctions

`func (o *RepresentantEntreprise) GetSanctions() []Sanction`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *RepresentantEntreprise) GetSanctionsOk() (*[]Sanction, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *RepresentantEntreprise) SetSanctions(v []Sanction)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *RepresentantEntreprise) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.

### GetDecede

`func (o *RepresentantEntreprise) GetDecede() bool`

GetDecede returns the Decede field if non-nil, zero value otherwise.

### GetDecedeOk

`func (o *RepresentantEntreprise) GetDecedeOk() (*bool, bool)`

GetDecedeOk returns a tuple with the Decede field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecede

`func (o *RepresentantEntreprise) SetDecede(v bool)`

SetDecede sets Decede field to given value.

### HasDecede

`func (o *RepresentantEntreprise) HasDecede() bool`

HasDecede returns a boolean if a field has been set.

### GetDateDeDeces

`func (o *RepresentantEntreprise) GetDateDeDeces() string`

GetDateDeDeces returns the DateDeDeces field if non-nil, zero value otherwise.

### GetDateDeDecesOk

`func (o *RepresentantEntreprise) GetDateDeDecesOk() (*string, bool)`

GetDateDeDecesOk returns a tuple with the DateDeDeces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeDeces

`func (o *RepresentantEntreprise) SetDateDeDeces(v string)`

SetDateDeDeces sets DateDeDeces field to given value.

### HasDateDeDeces

`func (o *RepresentantEntreprise) HasDateDeDeces() bool`

HasDateDeDeces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


