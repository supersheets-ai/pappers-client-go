# Beneficiaire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | Pointer to **string** | Nom du bénéficiaire effectif. | [optional] 
**NomUsage** | Pointer to **string** | Nom d&#39;usage du bénéficiaire effectif. | [optional] 
**Prenom** | Pointer to **string** | Prénom du bénéficiaire effectif. | [optional] 
**Pseudonyme** | Pointer to **string** | Pseudonyme du bénéficiaire effectif. | [optional] 
**NomComplet** | Pointer to **string** | Nom complet du bénéficiaire effectif. | [optional] 
**DateDeNaissanceFormate** | Pointer to **string** | Mois et année de naissance du bénéficiaire effectif, au format MM/AAAA. | [optional] 
**Nationalite** | Pointer to **string** | Nationalité du bénéficiaire effectif. | [optional] 
**PourcentageParts** | Pointer to **float32** | Parts détenues par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PourcentageVotes** | Pointer to **float32** | Droits de vote détenus par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**PourcentageVotesIndirect** | Pointer to **float32** | Droits de vote détenus de façon indirecte par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**PourcentageVotesDirects** | Pointer to **float32** | Droits de vote détenus de façon directe par le bénéficiaire effectif, en pourcentage des droits de vote totaux. | [optional] 
**DetentionAutresMoyensControle** | Pointer to **bool** | Vaut vrai s&#39;il existe d&#39;autres moyens de contrôle. | [optional] 
**BeneficiaireRepresentantLegal** | Pointer to **bool** | Vaut vrai dans le cas où le bénéficiaire effectif a été défini comme le représentant légal par défaut. | [optional] 
**AdresseLigne1** | Pointer to **string** | Première ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**AdresseLigne2** | Pointer to **string** | Deuxième ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**AdresseLigne3** | Pointer to **string** | Troisième ligne de l&#39;adresse du bénéficiaire effectif. | [optional] 
**PourcentagePartsVocationTitulaire** | Pointer to **float32** | Parts dont le bénéficiaire effectif a vocation à devenir titulaire par l&#39;effet d&#39;un acte juridique, en pourcentage des parts totales. | [optional] 
**RepresentantLegalPlacementSansGestionDelegation** | Pointer to **bool** | Vaut vrai dans le cas où le bénéficiaire effectif est le représentant légal du placement collectif (cas où le placement collectif n&#39;a pas délégué sa gestion à une société de gestion). | [optional] 
**CodePostal** | Pointer to **string** | Code postal du bénéficiaire effectif. | [optional] 
**DetentionPouvoirNomMembreConseilAdministration** | Pointer to **bool** | Vaut vrai si le moyen de contrôle est le pouvoir de nommer ou de révoquer la majorité des membres des organes d&#39;administration, de gestion, de direction ou de surveillance de la personne morale. | [optional] 
**Ville** | Pointer to **string** | Ville du bénéficiaire effectif. | [optional] 
**DateDeNaissanceCompleteFormatee** | Pointer to **string** | Date de naissance complète du bénéficiaire effectif, au format JJ/MM/AAAA. | [optional] 
**PourcentagePartsDirectes** | Pointer to **float32** | Parts détenues de façon directe par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PourcentagePartsIndirectes** | Pointer to **float32** | Parts détenues de façon indirecte par le bénéficiaire effectif, en pourcentage des parts totales. | [optional] 
**PaysDeNaissance** | Pointer to **string** | Pays de naissance du bénéficiaire effectif. | [optional] 
**CodePaysDeNaissance** | Pointer to **string** | Code du pays de naissance du bénéficiaire effectif. | [optional] 
**VilleDeNaissance** | Pointer to **string** | Ville de naissance du bénéficiaire effectif. | [optional] 
**DetentionPouvoirDecisionAg** | Pointer to **bool** | Vaut vrai pour les moyens de contrôle sur les organes d&#39;administration, de gestion, de direction ou de surveillance de la personne morale autre que le pouvoir de nommer ou de révoquer la majorité des membres. | [optional] 
**Pays** | Pointer to **string** | Pays du bénéficiaire effectif. | [optional] 
**DateDeNaissanceFormatee** | Pointer to **string** | Mois et année de naissance du bénéficiaire effectif, au format MM/AAAA. | [optional] 
**CodePays** | Pointer to **string** | Code pays du bénéficiaire effectif. | [optional] 

## Methods

### NewBeneficiaire

`func NewBeneficiaire() *Beneficiaire`

NewBeneficiaire instantiates a new Beneficiaire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficiaireWithDefaults

`func NewBeneficiaireWithDefaults() *Beneficiaire`

NewBeneficiaireWithDefaults instantiates a new Beneficiaire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *Beneficiaire) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Beneficiaire) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Beneficiaire) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Beneficiaire) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetNomUsage

`func (o *Beneficiaire) GetNomUsage() string`

GetNomUsage returns the NomUsage field if non-nil, zero value otherwise.

### GetNomUsageOk

`func (o *Beneficiaire) GetNomUsageOk() (*string, bool)`

GetNomUsageOk returns a tuple with the NomUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomUsage

`func (o *Beneficiaire) SetNomUsage(v string)`

SetNomUsage sets NomUsage field to given value.

### HasNomUsage

`func (o *Beneficiaire) HasNomUsage() bool`

HasNomUsage returns a boolean if a field has been set.

### GetPrenom

`func (o *Beneficiaire) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *Beneficiaire) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *Beneficiaire) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *Beneficiaire) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPseudonyme

`func (o *Beneficiaire) GetPseudonyme() string`

GetPseudonyme returns the Pseudonyme field if non-nil, zero value otherwise.

### GetPseudonymeOk

`func (o *Beneficiaire) GetPseudonymeOk() (*string, bool)`

GetPseudonymeOk returns a tuple with the Pseudonyme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonyme

`func (o *Beneficiaire) SetPseudonyme(v string)`

SetPseudonyme sets Pseudonyme field to given value.

### HasPseudonyme

`func (o *Beneficiaire) HasPseudonyme() bool`

HasPseudonyme returns a boolean if a field has been set.

### GetNomComplet

`func (o *Beneficiaire) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *Beneficiaire) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *Beneficiaire) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *Beneficiaire) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *Beneficiaire) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *Beneficiaire) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *Beneficiaire) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *Beneficiaire) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetNationalite

`func (o *Beneficiaire) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *Beneficiaire) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *Beneficiaire) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *Beneficiaire) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetPourcentageParts

`func (o *Beneficiaire) GetPourcentageParts() float32`

GetPourcentageParts returns the PourcentageParts field if non-nil, zero value otherwise.

### GetPourcentagePartsOk

`func (o *Beneficiaire) GetPourcentagePartsOk() (*float32, bool)`

GetPourcentagePartsOk returns a tuple with the PourcentageParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageParts

`func (o *Beneficiaire) SetPourcentageParts(v float32)`

SetPourcentageParts sets PourcentageParts field to given value.

### HasPourcentageParts

`func (o *Beneficiaire) HasPourcentageParts() bool`

HasPourcentageParts returns a boolean if a field has been set.

### GetPourcentageVotes

`func (o *Beneficiaire) GetPourcentageVotes() float32`

GetPourcentageVotes returns the PourcentageVotes field if non-nil, zero value otherwise.

### GetPourcentageVotesOk

`func (o *Beneficiaire) GetPourcentageVotesOk() (*float32, bool)`

GetPourcentageVotesOk returns a tuple with the PourcentageVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotes

`func (o *Beneficiaire) SetPourcentageVotes(v float32)`

SetPourcentageVotes sets PourcentageVotes field to given value.

### HasPourcentageVotes

`func (o *Beneficiaire) HasPourcentageVotes() bool`

HasPourcentageVotes returns a boolean if a field has been set.

### GetPourcentageVotesIndirect

`func (o *Beneficiaire) GetPourcentageVotesIndirect() float32`

GetPourcentageVotesIndirect returns the PourcentageVotesIndirect field if non-nil, zero value otherwise.

### GetPourcentageVotesIndirectOk

`func (o *Beneficiaire) GetPourcentageVotesIndirectOk() (*float32, bool)`

GetPourcentageVotesIndirectOk returns a tuple with the PourcentageVotesIndirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesIndirect

`func (o *Beneficiaire) SetPourcentageVotesIndirect(v float32)`

SetPourcentageVotesIndirect sets PourcentageVotesIndirect field to given value.

### HasPourcentageVotesIndirect

`func (o *Beneficiaire) HasPourcentageVotesIndirect() bool`

HasPourcentageVotesIndirect returns a boolean if a field has been set.

### GetPourcentageVotesDirects

`func (o *Beneficiaire) GetPourcentageVotesDirects() float32`

GetPourcentageVotesDirects returns the PourcentageVotesDirects field if non-nil, zero value otherwise.

### GetPourcentageVotesDirectsOk

`func (o *Beneficiaire) GetPourcentageVotesDirectsOk() (*float32, bool)`

GetPourcentageVotesDirectsOk returns a tuple with the PourcentageVotesDirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesDirects

`func (o *Beneficiaire) SetPourcentageVotesDirects(v float32)`

SetPourcentageVotesDirects sets PourcentageVotesDirects field to given value.

### HasPourcentageVotesDirects

`func (o *Beneficiaire) HasPourcentageVotesDirects() bool`

HasPourcentageVotesDirects returns a boolean if a field has been set.

### GetDetentionAutresMoyensControle

`func (o *Beneficiaire) GetDetentionAutresMoyensControle() bool`

GetDetentionAutresMoyensControle returns the DetentionAutresMoyensControle field if non-nil, zero value otherwise.

### GetDetentionAutresMoyensControleOk

`func (o *Beneficiaire) GetDetentionAutresMoyensControleOk() (*bool, bool)`

GetDetentionAutresMoyensControleOk returns a tuple with the DetentionAutresMoyensControle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionAutresMoyensControle

`func (o *Beneficiaire) SetDetentionAutresMoyensControle(v bool)`

SetDetentionAutresMoyensControle sets DetentionAutresMoyensControle field to given value.

### HasDetentionAutresMoyensControle

`func (o *Beneficiaire) HasDetentionAutresMoyensControle() bool`

HasDetentionAutresMoyensControle returns a boolean if a field has been set.

### GetBeneficiaireRepresentantLegal

`func (o *Beneficiaire) GetBeneficiaireRepresentantLegal() bool`

GetBeneficiaireRepresentantLegal returns the BeneficiaireRepresentantLegal field if non-nil, zero value otherwise.

### GetBeneficiaireRepresentantLegalOk

`func (o *Beneficiaire) GetBeneficiaireRepresentantLegalOk() (*bool, bool)`

GetBeneficiaireRepresentantLegalOk returns a tuple with the BeneficiaireRepresentantLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaireRepresentantLegal

`func (o *Beneficiaire) SetBeneficiaireRepresentantLegal(v bool)`

SetBeneficiaireRepresentantLegal sets BeneficiaireRepresentantLegal field to given value.

### HasBeneficiaireRepresentantLegal

`func (o *Beneficiaire) HasBeneficiaireRepresentantLegal() bool`

HasBeneficiaireRepresentantLegal returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *Beneficiaire) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *Beneficiaire) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *Beneficiaire) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *Beneficiaire) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *Beneficiaire) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *Beneficiaire) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *Beneficiaire) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *Beneficiaire) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *Beneficiaire) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *Beneficiaire) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *Beneficiaire) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *Beneficiaire) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetPourcentagePartsVocationTitulaire

`func (o *Beneficiaire) GetPourcentagePartsVocationTitulaire() float32`

GetPourcentagePartsVocationTitulaire returns the PourcentagePartsVocationTitulaire field if non-nil, zero value otherwise.

### GetPourcentagePartsVocationTitulaireOk

`func (o *Beneficiaire) GetPourcentagePartsVocationTitulaireOk() (*float32, bool)`

GetPourcentagePartsVocationTitulaireOk returns a tuple with the PourcentagePartsVocationTitulaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsVocationTitulaire

`func (o *Beneficiaire) SetPourcentagePartsVocationTitulaire(v float32)`

SetPourcentagePartsVocationTitulaire sets PourcentagePartsVocationTitulaire field to given value.

### HasPourcentagePartsVocationTitulaire

`func (o *Beneficiaire) HasPourcentagePartsVocationTitulaire() bool`

HasPourcentagePartsVocationTitulaire returns a boolean if a field has been set.

### GetRepresentantLegalPlacementSansGestionDelegation

`func (o *Beneficiaire) GetRepresentantLegalPlacementSansGestionDelegation() bool`

GetRepresentantLegalPlacementSansGestionDelegation returns the RepresentantLegalPlacementSansGestionDelegation field if non-nil, zero value otherwise.

### GetRepresentantLegalPlacementSansGestionDelegationOk

`func (o *Beneficiaire) GetRepresentantLegalPlacementSansGestionDelegationOk() (*bool, bool)`

GetRepresentantLegalPlacementSansGestionDelegationOk returns a tuple with the RepresentantLegalPlacementSansGestionDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentantLegalPlacementSansGestionDelegation

`func (o *Beneficiaire) SetRepresentantLegalPlacementSansGestionDelegation(v bool)`

SetRepresentantLegalPlacementSansGestionDelegation sets RepresentantLegalPlacementSansGestionDelegation field to given value.

### HasRepresentantLegalPlacementSansGestionDelegation

`func (o *Beneficiaire) HasRepresentantLegalPlacementSansGestionDelegation() bool`

HasRepresentantLegalPlacementSansGestionDelegation returns a boolean if a field has been set.

### GetCodePostal

`func (o *Beneficiaire) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *Beneficiaire) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *Beneficiaire) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *Beneficiaire) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetDetentionPouvoirNomMembreConseilAdministration

`func (o *Beneficiaire) GetDetentionPouvoirNomMembreConseilAdministration() bool`

GetDetentionPouvoirNomMembreConseilAdministration returns the DetentionPouvoirNomMembreConseilAdministration field if non-nil, zero value otherwise.

### GetDetentionPouvoirNomMembreConseilAdministrationOk

`func (o *Beneficiaire) GetDetentionPouvoirNomMembreConseilAdministrationOk() (*bool, bool)`

GetDetentionPouvoirNomMembreConseilAdministrationOk returns a tuple with the DetentionPouvoirNomMembreConseilAdministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirNomMembreConseilAdministration

`func (o *Beneficiaire) SetDetentionPouvoirNomMembreConseilAdministration(v bool)`

SetDetentionPouvoirNomMembreConseilAdministration sets DetentionPouvoirNomMembreConseilAdministration field to given value.

### HasDetentionPouvoirNomMembreConseilAdministration

`func (o *Beneficiaire) HasDetentionPouvoirNomMembreConseilAdministration() bool`

HasDetentionPouvoirNomMembreConseilAdministration returns a boolean if a field has been set.

### GetVille

`func (o *Beneficiaire) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *Beneficiaire) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *Beneficiaire) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *Beneficiaire) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetDateDeNaissanceCompleteFormatee

`func (o *Beneficiaire) GetDateDeNaissanceCompleteFormatee() string`

GetDateDeNaissanceCompleteFormatee returns the DateDeNaissanceCompleteFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceCompleteFormateeOk

`func (o *Beneficiaire) GetDateDeNaissanceCompleteFormateeOk() (*string, bool)`

GetDateDeNaissanceCompleteFormateeOk returns a tuple with the DateDeNaissanceCompleteFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceCompleteFormatee

`func (o *Beneficiaire) SetDateDeNaissanceCompleteFormatee(v string)`

SetDateDeNaissanceCompleteFormatee sets DateDeNaissanceCompleteFormatee field to given value.

### HasDateDeNaissanceCompleteFormatee

`func (o *Beneficiaire) HasDateDeNaissanceCompleteFormatee() bool`

HasDateDeNaissanceCompleteFormatee returns a boolean if a field has been set.

### GetPourcentagePartsDirectes

`func (o *Beneficiaire) GetPourcentagePartsDirectes() float32`

GetPourcentagePartsDirectes returns the PourcentagePartsDirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsDirectesOk

`func (o *Beneficiaire) GetPourcentagePartsDirectesOk() (*float32, bool)`

GetPourcentagePartsDirectesOk returns a tuple with the PourcentagePartsDirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsDirectes

`func (o *Beneficiaire) SetPourcentagePartsDirectes(v float32)`

SetPourcentagePartsDirectes sets PourcentagePartsDirectes field to given value.

### HasPourcentagePartsDirectes

`func (o *Beneficiaire) HasPourcentagePartsDirectes() bool`

HasPourcentagePartsDirectes returns a boolean if a field has been set.

### GetPourcentagePartsIndirectes

`func (o *Beneficiaire) GetPourcentagePartsIndirectes() float32`

GetPourcentagePartsIndirectes returns the PourcentagePartsIndirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsIndirectesOk

`func (o *Beneficiaire) GetPourcentagePartsIndirectesOk() (*float32, bool)`

GetPourcentagePartsIndirectesOk returns a tuple with the PourcentagePartsIndirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsIndirectes

`func (o *Beneficiaire) SetPourcentagePartsIndirectes(v float32)`

SetPourcentagePartsIndirectes sets PourcentagePartsIndirectes field to given value.

### HasPourcentagePartsIndirectes

`func (o *Beneficiaire) HasPourcentagePartsIndirectes() bool`

HasPourcentagePartsIndirectes returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *Beneficiaire) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *Beneficiaire) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *Beneficiaire) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *Beneficiaire) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *Beneficiaire) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *Beneficiaire) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *Beneficiaire) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *Beneficiaire) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *Beneficiaire) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *Beneficiaire) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *Beneficiaire) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *Beneficiaire) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetDetentionPouvoirDecisionAg

`func (o *Beneficiaire) GetDetentionPouvoirDecisionAg() bool`

GetDetentionPouvoirDecisionAg returns the DetentionPouvoirDecisionAg field if non-nil, zero value otherwise.

### GetDetentionPouvoirDecisionAgOk

`func (o *Beneficiaire) GetDetentionPouvoirDecisionAgOk() (*bool, bool)`

GetDetentionPouvoirDecisionAgOk returns a tuple with the DetentionPouvoirDecisionAg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirDecisionAg

`func (o *Beneficiaire) SetDetentionPouvoirDecisionAg(v bool)`

SetDetentionPouvoirDecisionAg sets DetentionPouvoirDecisionAg field to given value.

### HasDetentionPouvoirDecisionAg

`func (o *Beneficiaire) HasDetentionPouvoirDecisionAg() bool`

HasDetentionPouvoirDecisionAg returns a boolean if a field has been set.

### GetPays

`func (o *Beneficiaire) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *Beneficiaire) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *Beneficiaire) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *Beneficiaire) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetDateDeNaissanceFormatee

`func (o *Beneficiaire) GetDateDeNaissanceFormatee() string`

GetDateDeNaissanceFormatee returns the DateDeNaissanceFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateeOk

`func (o *Beneficiaire) GetDateDeNaissanceFormateeOk() (*string, bool)`

GetDateDeNaissanceFormateeOk returns a tuple with the DateDeNaissanceFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormatee

`func (o *Beneficiaire) SetDateDeNaissanceFormatee(v string)`

SetDateDeNaissanceFormatee sets DateDeNaissanceFormatee field to given value.

### HasDateDeNaissanceFormatee

`func (o *Beneficiaire) HasDateDeNaissanceFormatee() bool`

HasDateDeNaissanceFormatee returns a boolean if a field has been set.

### GetCodePays

`func (o *Beneficiaire) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *Beneficiaire) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *Beneficiaire) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *Beneficiaire) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


