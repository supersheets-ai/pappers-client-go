# EntrepriseFicheAllOfBeneficiairesEffectifs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateGreffe** | Pointer to **string** | Date de génération des bénéficiaires effectifs, au format AAAA-MM-JJ. | [optional] 
**Type** | Pointer to **string** | Type du bénéficiaire effectif | [optional] 
**Nom** | Pointer to **string** | Nom du bénéficiaire effectif. | [optional] 
**NomUsage** | Pointer to **string** | Nom d&#39;usage du bénéficiaire effectif. | [optional] 
**Prenom** | Pointer to **string** | Prénoms du bénéficiaire effectif. | [optional] 
**PrenomUsuel** | Pointer to **string** | Prénom usuel du bénéficiaire effectif. | [optional] 
**Pseudonyme** | Pointer to **string** |  | [optional] 
**Sexe** | Pointer to **string** | Sexe du bénéficiaire si personne physique. F pour féminin, M pour masculin. | [optional] 
**DateDeNaissanceFormatee** | Pointer to **string** | Mois et année de naissance du bénéficiaire effectif, au format MM/AAAA. | [optional] 
**DateDeNaissanceCompleteFormatee** | Pointer to **string** | Date de naissance complète du bénéficiaire effectif, au format JJ/MM/AAAA. | [optional] 
**Nationalite** | Pointer to **string** | Nationalité du bénéficiaire effectif. | [optional] 
**CodeNationalite** | Pointer to **string** | Code de la nationalité du bénéficiaire effectif. | [optional] 
**VilleDeNaissance** | Pointer to **string** | Ville de naissance du bénéficiaire effectif. | [optional] 
**PaysDeNaissance** | Pointer to **string** | Pays de naissance du bénéficiaire effectif. | [optional] 
**CodePaysDeNaissance** | Pointer to **string** | Code du pays de naissance du bénéficiaire effectif. | [optional] 
**AdresseLigne1** | Pointer to **string** | Première ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**AdresseLigne2** | Pointer to **string** | Deuxième ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**AdresseLigne3** | Pointer to **string** | Troisième ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**CodePostal** | Pointer to **string** | Code postal du bénéficiaire effectif. | [optional] 
**Ville** | Pointer to **string** | Ville du bénéficiaire effectif. | [optional] 
**Pays** | Pointer to **string** | Pays du bénéficiaire effectif. | [optional] 
**CodePays** | Pointer to **string** | Code du pays du bénéficiaire effectif. | [optional] 
**PourcentageParts** | Pointer to **float32** | Parts détenues par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PourcentagePartsDirectes** | Pointer to **float32** | Parts détenues de façon directe par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PourcentagePartsIndirectes** | Pointer to **float32** | Parts détenues de façon indirecte par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PourcentagePartsVocationTitulaire** | Pointer to **float32** | Parts dont le bénéficiaire effectif a vocation à devenir titulaire par l&#39;effet d&#39;un acte juridique, en pourcentage des parts totales. | [optional] 
**DetailsPartsDirectes** | Pointer to [**EntrepriseFicheAllOfDetailsPartsDirectes**](EntrepriseFicheAllOfDetailsPartsDirectes.md) |  | [optional] 
**DetailsPartsIndirectes** | Pointer to [**EntrepriseFicheAllOfDetailsPartsIndirectes**](EntrepriseFicheAllOfDetailsPartsIndirectes.md) |  | [optional] 
**DetailsPartsVocationTitulaire** | Pointer to [**EntrepriseFicheAllOfDetailsPartsVocationTitulaire**](EntrepriseFicheAllOfDetailsPartsVocationTitulaire.md) |  | [optional] 
**PourcentageVotes** | Pointer to **float32** | Droits de vote détenus par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**PourcentageVotesDirects** | Pointer to **float32** | Droits de vote détenus de façon directe par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**PourcentageVotesIndirect** | Pointer to **float32** | Droits de vote détenus de façon indirecte par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**DetailsVotesDirects** | Pointer to [**EntrepriseFicheAllOfDetailsVotesDirects**](EntrepriseFicheAllOfDetailsVotesDirects.md) |  | [optional] 
**DetailsVotesIndirects** | Pointer to [**EntrepriseFicheAllOfDetailsVotesIndirects**](EntrepriseFicheAllOfDetailsVotesIndirects.md) |  | [optional] 
**DetailsSocieteDeGestion** | Pointer to [**EntrepriseFicheAllOfDetailsSocieteDeGestion**](EntrepriseFicheAllOfDetailsSocieteDeGestion.md) |  | [optional] 
**DetentionPouvoirDecisionAg** | Pointer to **bool** | Vaut vrai pour les moyens de contrôle sur les organes d&#39;administration, de gestion, de direction ou de surveillance de la personne morale autre que le pouvoir de nommer ou de révoquer la majorité des membres. | [optional] 
**DetentionPouvoirNomMembreConseilAdministration** | Pointer to **bool** | Vaut vrai si le moyen de contrôle est le pouvoir de nommer ou de révoquer la majorité des membres des organes d&#39;administration, de gestion, de direction ou de surveillance de la personne morale. | [optional] 
**DetentionAutresMoyensControle** | Pointer to **bool** | Vaut vrai s&#39;il existe d&#39;autres moyens de contrôle. | [optional] 
**BeneficiaireRepresentantLegal** | Pointer to **bool** | Vaut vrai dans le cas où le bénéficiaire effectif a été défini comme le représentant légal par défaut. | [optional] 
**RepresentantLegalPlacementSansGestionDelegation** | Pointer to **bool** | Vaut vrai dans le cas où le bénéficiaire effectif est le représentant légal du placement collectif (cas où le placement collectif n&#39;a pas délégué sa gestion à une société de gestion). | [optional] 
**PersonnePolitiquementExposee** | Pointer to [**PersonnePolitiquementExposee**](PersonnePolitiquementExposee.md) |  | [optional] 
**SanctionsEnCours** | Pointer to **bool** | Vaut vrai si le bénéficiaire effectif est actuellement sous sanction. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Sanctions** | Pointer to [**[]Sanction**](Sanction.md) | Liste des sanctions du bénéficiaire effectif. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 

## Methods

### NewEntrepriseFicheAllOfBeneficiairesEffectifs

`func NewEntrepriseFicheAllOfBeneficiairesEffectifs() *EntrepriseFicheAllOfBeneficiairesEffectifs`

NewEntrepriseFicheAllOfBeneficiairesEffectifs instantiates a new EntrepriseFicheAllOfBeneficiairesEffectifs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfBeneficiairesEffectifsWithDefaults

`func NewEntrepriseFicheAllOfBeneficiairesEffectifsWithDefaults() *EntrepriseFicheAllOfBeneficiairesEffectifs`

NewEntrepriseFicheAllOfBeneficiairesEffectifsWithDefaults instantiates a new EntrepriseFicheAllOfBeneficiairesEffectifs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateGreffe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateGreffe() string`

GetDateGreffe returns the DateGreffe field if non-nil, zero value otherwise.

### GetDateGreffeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateGreffeOk() (*string, bool)`

GetDateGreffeOk returns a tuple with the DateGreffe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateGreffe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDateGreffe(v string)`

SetDateGreffe sets DateGreffe field to given value.

### HasDateGreffe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDateGreffe() bool`

HasDateGreffe returns a boolean if a field has been set.

### GetType

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetNomUsage

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNomUsage() string`

GetNomUsage returns the NomUsage field if non-nil, zero value otherwise.

### GetNomUsageOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNomUsageOk() (*string, bool)`

GetNomUsageOk returns a tuple with the NomUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomUsage

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetNomUsage(v string)`

SetNomUsage sets NomUsage field to given value.

### HasNomUsage

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasNomUsage() bool`

HasNomUsage returns a boolean if a field has been set.

### GetPrenom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPrenomUsuel

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPrenomUsuel() string`

GetPrenomUsuel returns the PrenomUsuel field if non-nil, zero value otherwise.

### GetPrenomUsuelOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPrenomUsuelOk() (*string, bool)`

GetPrenomUsuelOk returns a tuple with the PrenomUsuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenomUsuel

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPrenomUsuel(v string)`

SetPrenomUsuel sets PrenomUsuel field to given value.

### HasPrenomUsuel

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPrenomUsuel() bool`

HasPrenomUsuel returns a boolean if a field has been set.

### GetPseudonyme

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPseudonyme() string`

GetPseudonyme returns the Pseudonyme field if non-nil, zero value otherwise.

### GetPseudonymeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPseudonymeOk() (*string, bool)`

GetPseudonymeOk returns a tuple with the Pseudonyme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonyme

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPseudonyme(v string)`

SetPseudonyme sets Pseudonyme field to given value.

### HasPseudonyme

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPseudonyme() bool`

HasPseudonyme returns a boolean if a field has been set.

### GetSexe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetDateDeNaissanceFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateDeNaissanceFormatee() string`

GetDateDeNaissanceFormatee returns the DateDeNaissanceFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateDeNaissanceFormateeOk() (*string, bool)`

GetDateDeNaissanceFormateeOk returns a tuple with the DateDeNaissanceFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDateDeNaissanceFormatee(v string)`

SetDateDeNaissanceFormatee sets DateDeNaissanceFormatee field to given value.

### HasDateDeNaissanceFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDateDeNaissanceFormatee() bool`

HasDateDeNaissanceFormatee returns a boolean if a field has been set.

### GetDateDeNaissanceCompleteFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateDeNaissanceCompleteFormatee() string`

GetDateDeNaissanceCompleteFormatee returns the DateDeNaissanceCompleteFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceCompleteFormateeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDateDeNaissanceCompleteFormateeOk() (*string, bool)`

GetDateDeNaissanceCompleteFormateeOk returns a tuple with the DateDeNaissanceCompleteFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceCompleteFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDateDeNaissanceCompleteFormatee(v string)`

SetDateDeNaissanceCompleteFormatee sets DateDeNaissanceCompleteFormatee field to given value.

### HasDateDeNaissanceCompleteFormatee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDateDeNaissanceCompleteFormatee() bool`

HasDateDeNaissanceCompleteFormatee returns a boolean if a field has been set.

### GetNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetCodeNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodeNationalite() string`

GetCodeNationalite returns the CodeNationalite field if non-nil, zero value otherwise.

### GetCodeNationaliteOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodeNationaliteOk() (*string, bool)`

GetCodeNationaliteOk returns a tuple with the CodeNationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetCodeNationalite(v string)`

SetCodeNationalite sets CodeNationalite field to given value.

### HasCodeNationalite

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasCodeNationalite() bool`

HasCodeNationalite returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetCodePostal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetPourcentageParts

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageParts() float32`

GetPourcentageParts returns the PourcentageParts field if non-nil, zero value otherwise.

### GetPourcentagePartsOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsOk() (*float32, bool)`

GetPourcentagePartsOk returns a tuple with the PourcentageParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageParts

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentageParts(v float32)`

SetPourcentageParts sets PourcentageParts field to given value.

### HasPourcentageParts

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentageParts() bool`

HasPourcentageParts returns a boolean if a field has been set.

### GetPourcentagePartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsDirectes() float32`

GetPourcentagePartsDirectes returns the PourcentagePartsDirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsDirectesOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsDirectesOk() (*float32, bool)`

GetPourcentagePartsDirectesOk returns a tuple with the PourcentagePartsDirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentagePartsDirectes(v float32)`

SetPourcentagePartsDirectes sets PourcentagePartsDirectes field to given value.

### HasPourcentagePartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentagePartsDirectes() bool`

HasPourcentagePartsDirectes returns a boolean if a field has been set.

### GetPourcentagePartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsIndirectes() float32`

GetPourcentagePartsIndirectes returns the PourcentagePartsIndirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsIndirectesOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsIndirectesOk() (*float32, bool)`

GetPourcentagePartsIndirectesOk returns a tuple with the PourcentagePartsIndirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentagePartsIndirectes(v float32)`

SetPourcentagePartsIndirectes sets PourcentagePartsIndirectes field to given value.

### HasPourcentagePartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentagePartsIndirectes() bool`

HasPourcentagePartsIndirectes returns a boolean if a field has been set.

### GetPourcentagePartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsVocationTitulaire() float32`

GetPourcentagePartsVocationTitulaire returns the PourcentagePartsVocationTitulaire field if non-nil, zero value otherwise.

### GetPourcentagePartsVocationTitulaireOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentagePartsVocationTitulaireOk() (*float32, bool)`

GetPourcentagePartsVocationTitulaireOk returns a tuple with the PourcentagePartsVocationTitulaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentagePartsVocationTitulaire(v float32)`

SetPourcentagePartsVocationTitulaire sets PourcentagePartsVocationTitulaire field to given value.

### HasPourcentagePartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentagePartsVocationTitulaire() bool`

HasPourcentagePartsVocationTitulaire returns a boolean if a field has been set.

### GetDetailsPartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsDirectes() EntrepriseFicheAllOfDetailsPartsDirectes`

GetDetailsPartsDirectes returns the DetailsPartsDirectes field if non-nil, zero value otherwise.

### GetDetailsPartsDirectesOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsDirectesOk() (*EntrepriseFicheAllOfDetailsPartsDirectes, bool)`

GetDetailsPartsDirectesOk returns a tuple with the DetailsPartsDirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsPartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsPartsDirectes(v EntrepriseFicheAllOfDetailsPartsDirectes)`

SetDetailsPartsDirectes sets DetailsPartsDirectes field to given value.

### HasDetailsPartsDirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsPartsDirectes() bool`

HasDetailsPartsDirectes returns a boolean if a field has been set.

### GetDetailsPartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsIndirectes() EntrepriseFicheAllOfDetailsPartsIndirectes`

GetDetailsPartsIndirectes returns the DetailsPartsIndirectes field if non-nil, zero value otherwise.

### GetDetailsPartsIndirectesOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsIndirectesOk() (*EntrepriseFicheAllOfDetailsPartsIndirectes, bool)`

GetDetailsPartsIndirectesOk returns a tuple with the DetailsPartsIndirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsPartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsPartsIndirectes(v EntrepriseFicheAllOfDetailsPartsIndirectes)`

SetDetailsPartsIndirectes sets DetailsPartsIndirectes field to given value.

### HasDetailsPartsIndirectes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsPartsIndirectes() bool`

HasDetailsPartsIndirectes returns a boolean if a field has been set.

### GetDetailsPartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsVocationTitulaire() EntrepriseFicheAllOfDetailsPartsVocationTitulaire`

GetDetailsPartsVocationTitulaire returns the DetailsPartsVocationTitulaire field if non-nil, zero value otherwise.

### GetDetailsPartsVocationTitulaireOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsPartsVocationTitulaireOk() (*EntrepriseFicheAllOfDetailsPartsVocationTitulaire, bool)`

GetDetailsPartsVocationTitulaireOk returns a tuple with the DetailsPartsVocationTitulaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsPartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsPartsVocationTitulaire(v EntrepriseFicheAllOfDetailsPartsVocationTitulaire)`

SetDetailsPartsVocationTitulaire sets DetailsPartsVocationTitulaire field to given value.

### HasDetailsPartsVocationTitulaire

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsPartsVocationTitulaire() bool`

HasDetailsPartsVocationTitulaire returns a boolean if a field has been set.

### GetPourcentageVotes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotes() float32`

GetPourcentageVotes returns the PourcentageVotes field if non-nil, zero value otherwise.

### GetPourcentageVotesOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotesOk() (*float32, bool)`

GetPourcentageVotesOk returns a tuple with the PourcentageVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentageVotes(v float32)`

SetPourcentageVotes sets PourcentageVotes field to given value.

### HasPourcentageVotes

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentageVotes() bool`

HasPourcentageVotes returns a boolean if a field has been set.

### GetPourcentageVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotesDirects() float32`

GetPourcentageVotesDirects returns the PourcentageVotesDirects field if non-nil, zero value otherwise.

### GetPourcentageVotesDirectsOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotesDirectsOk() (*float32, bool)`

GetPourcentageVotesDirectsOk returns a tuple with the PourcentageVotesDirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentageVotesDirects(v float32)`

SetPourcentageVotesDirects sets PourcentageVotesDirects field to given value.

### HasPourcentageVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentageVotesDirects() bool`

HasPourcentageVotesDirects returns a boolean if a field has been set.

### GetPourcentageVotesIndirect

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotesIndirect() float32`

GetPourcentageVotesIndirect returns the PourcentageVotesIndirect field if non-nil, zero value otherwise.

### GetPourcentageVotesIndirectOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPourcentageVotesIndirectOk() (*float32, bool)`

GetPourcentageVotesIndirectOk returns a tuple with the PourcentageVotesIndirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesIndirect

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPourcentageVotesIndirect(v float32)`

SetPourcentageVotesIndirect sets PourcentageVotesIndirect field to given value.

### HasPourcentageVotesIndirect

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPourcentageVotesIndirect() bool`

HasPourcentageVotesIndirect returns a boolean if a field has been set.

### GetDetailsVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsVotesDirects() EntrepriseFicheAllOfDetailsVotesDirects`

GetDetailsVotesDirects returns the DetailsVotesDirects field if non-nil, zero value otherwise.

### GetDetailsVotesDirectsOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsVotesDirectsOk() (*EntrepriseFicheAllOfDetailsVotesDirects, bool)`

GetDetailsVotesDirectsOk returns a tuple with the DetailsVotesDirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsVotesDirects(v EntrepriseFicheAllOfDetailsVotesDirects)`

SetDetailsVotesDirects sets DetailsVotesDirects field to given value.

### HasDetailsVotesDirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsVotesDirects() bool`

HasDetailsVotesDirects returns a boolean if a field has been set.

### GetDetailsVotesIndirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsVotesIndirects() EntrepriseFicheAllOfDetailsVotesIndirects`

GetDetailsVotesIndirects returns the DetailsVotesIndirects field if non-nil, zero value otherwise.

### GetDetailsVotesIndirectsOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsVotesIndirectsOk() (*EntrepriseFicheAllOfDetailsVotesIndirects, bool)`

GetDetailsVotesIndirectsOk returns a tuple with the DetailsVotesIndirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsVotesIndirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsVotesIndirects(v EntrepriseFicheAllOfDetailsVotesIndirects)`

SetDetailsVotesIndirects sets DetailsVotesIndirects field to given value.

### HasDetailsVotesIndirects

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsVotesIndirects() bool`

HasDetailsVotesIndirects returns a boolean if a field has been set.

### GetDetailsSocieteDeGestion

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsSocieteDeGestion() EntrepriseFicheAllOfDetailsSocieteDeGestion`

GetDetailsSocieteDeGestion returns the DetailsSocieteDeGestion field if non-nil, zero value otherwise.

### GetDetailsSocieteDeGestionOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetailsSocieteDeGestionOk() (*EntrepriseFicheAllOfDetailsSocieteDeGestion, bool)`

GetDetailsSocieteDeGestionOk returns a tuple with the DetailsSocieteDeGestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsSocieteDeGestion

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetailsSocieteDeGestion(v EntrepriseFicheAllOfDetailsSocieteDeGestion)`

SetDetailsSocieteDeGestion sets DetailsSocieteDeGestion field to given value.

### HasDetailsSocieteDeGestion

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetailsSocieteDeGestion() bool`

HasDetailsSocieteDeGestion returns a boolean if a field has been set.

### GetDetentionPouvoirDecisionAg

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionPouvoirDecisionAg() bool`

GetDetentionPouvoirDecisionAg returns the DetentionPouvoirDecisionAg field if non-nil, zero value otherwise.

### GetDetentionPouvoirDecisionAgOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionPouvoirDecisionAgOk() (*bool, bool)`

GetDetentionPouvoirDecisionAgOk returns a tuple with the DetentionPouvoirDecisionAg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirDecisionAg

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetentionPouvoirDecisionAg(v bool)`

SetDetentionPouvoirDecisionAg sets DetentionPouvoirDecisionAg field to given value.

### HasDetentionPouvoirDecisionAg

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetentionPouvoirDecisionAg() bool`

HasDetentionPouvoirDecisionAg returns a boolean if a field has been set.

### GetDetentionPouvoirNomMembreConseilAdministration

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionPouvoirNomMembreConseilAdministration() bool`

GetDetentionPouvoirNomMembreConseilAdministration returns the DetentionPouvoirNomMembreConseilAdministration field if non-nil, zero value otherwise.

### GetDetentionPouvoirNomMembreConseilAdministrationOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionPouvoirNomMembreConseilAdministrationOk() (*bool, bool)`

GetDetentionPouvoirNomMembreConseilAdministrationOk returns a tuple with the DetentionPouvoirNomMembreConseilAdministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirNomMembreConseilAdministration

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetentionPouvoirNomMembreConseilAdministration(v bool)`

SetDetentionPouvoirNomMembreConseilAdministration sets DetentionPouvoirNomMembreConseilAdministration field to given value.

### HasDetentionPouvoirNomMembreConseilAdministration

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetentionPouvoirNomMembreConseilAdministration() bool`

HasDetentionPouvoirNomMembreConseilAdministration returns a boolean if a field has been set.

### GetDetentionAutresMoyensControle

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionAutresMoyensControle() bool`

GetDetentionAutresMoyensControle returns the DetentionAutresMoyensControle field if non-nil, zero value otherwise.

### GetDetentionAutresMoyensControleOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetDetentionAutresMoyensControleOk() (*bool, bool)`

GetDetentionAutresMoyensControleOk returns a tuple with the DetentionAutresMoyensControle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionAutresMoyensControle

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetDetentionAutresMoyensControle(v bool)`

SetDetentionAutresMoyensControle sets DetentionAutresMoyensControle field to given value.

### HasDetentionAutresMoyensControle

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasDetentionAutresMoyensControle() bool`

HasDetentionAutresMoyensControle returns a boolean if a field has been set.

### GetBeneficiaireRepresentantLegal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetBeneficiaireRepresentantLegal() bool`

GetBeneficiaireRepresentantLegal returns the BeneficiaireRepresentantLegal field if non-nil, zero value otherwise.

### GetBeneficiaireRepresentantLegalOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetBeneficiaireRepresentantLegalOk() (*bool, bool)`

GetBeneficiaireRepresentantLegalOk returns a tuple with the BeneficiaireRepresentantLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaireRepresentantLegal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetBeneficiaireRepresentantLegal(v bool)`

SetBeneficiaireRepresentantLegal sets BeneficiaireRepresentantLegal field to given value.

### HasBeneficiaireRepresentantLegal

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasBeneficiaireRepresentantLegal() bool`

HasBeneficiaireRepresentantLegal returns a boolean if a field has been set.

### GetRepresentantLegalPlacementSansGestionDelegation

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetRepresentantLegalPlacementSansGestionDelegation() bool`

GetRepresentantLegalPlacementSansGestionDelegation returns the RepresentantLegalPlacementSansGestionDelegation field if non-nil, zero value otherwise.

### GetRepresentantLegalPlacementSansGestionDelegationOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetRepresentantLegalPlacementSansGestionDelegationOk() (*bool, bool)`

GetRepresentantLegalPlacementSansGestionDelegationOk returns a tuple with the RepresentantLegalPlacementSansGestionDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentantLegalPlacementSansGestionDelegation

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetRepresentantLegalPlacementSansGestionDelegation(v bool)`

SetRepresentantLegalPlacementSansGestionDelegation sets RepresentantLegalPlacementSansGestionDelegation field to given value.

### HasRepresentantLegalPlacementSansGestionDelegation

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasRepresentantLegalPlacementSansGestionDelegation() bool`

HasRepresentantLegalPlacementSansGestionDelegation returns a boolean if a field has been set.

### GetPersonnePolitiquementExposee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPersonnePolitiquementExposee() PersonnePolitiquementExposee`

GetPersonnePolitiquementExposee returns the PersonnePolitiquementExposee field if non-nil, zero value otherwise.

### GetPersonnePolitiquementExposeeOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetPersonnePolitiquementExposeeOk() (*PersonnePolitiquementExposee, bool)`

GetPersonnePolitiquementExposeeOk returns a tuple with the PersonnePolitiquementExposee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonnePolitiquementExposee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetPersonnePolitiquementExposee(v PersonnePolitiquementExposee)`

SetPersonnePolitiquementExposee sets PersonnePolitiquementExposee field to given value.

### HasPersonnePolitiquementExposee

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasPersonnePolitiquementExposee() bool`

HasPersonnePolitiquementExposee returns a boolean if a field has been set.

### GetSanctionsEnCours

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSanctionsEnCours() bool`

GetSanctionsEnCours returns the SanctionsEnCours field if non-nil, zero value otherwise.

### GetSanctionsEnCoursOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSanctionsEnCoursOk() (*bool, bool)`

GetSanctionsEnCoursOk returns a tuple with the SanctionsEnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsEnCours

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetSanctionsEnCours(v bool)`

SetSanctionsEnCours sets SanctionsEnCours field to given value.

### HasSanctionsEnCours

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasSanctionsEnCours() bool`

HasSanctionsEnCours returns a boolean if a field has been set.

### GetSanctions

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSanctions() []Sanction`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) GetSanctionsOk() (*[]Sanction, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) SetSanctions(v []Sanction)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *EntrepriseFicheAllOfBeneficiairesEffectifs) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


