# RechercheBeneficiaires200ResponseResultatsInner

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
**Entreprises** | Pointer to [**[]EntrepriseRecherche**](EntrepriseRecherche.md) | Liste des entreprises dont le bénéficiaire effectif est bénéficiaire effectif, dans la limite de 100 entreprises. | [optional] 
**NbEntreprisesTotal** | Pointer to **int32** | Nombre d&#39;entreprises du bénéficiaire effectif au total. | [optional] 
**EntreprisesDirigeant** | Pointer to [**[]EntrepriseRecherche**](EntrepriseRecherche.md) | Liste des entreprises dont le bénéficiaire effectif est dirigeant (sans forcément en être bénéficiaire effectif), dans la limite de 100 entreprises. | [optional] 
**NbEntreprisesDirigeantTotal** | Pointer to **int32** | Nombre d&#39;entreprises dont le bénéficiaire effectif est dirigeant au total. | [optional] 

## Methods

### NewRechercheBeneficiaires200ResponseResultatsInner

`func NewRechercheBeneficiaires200ResponseResultatsInner() *RechercheBeneficiaires200ResponseResultatsInner`

NewRechercheBeneficiaires200ResponseResultatsInner instantiates a new RechercheBeneficiaires200ResponseResultatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechercheBeneficiaires200ResponseResultatsInnerWithDefaults

`func NewRechercheBeneficiaires200ResponseResultatsInnerWithDefaults() *RechercheBeneficiaires200ResponseResultatsInner`

NewRechercheBeneficiaires200ResponseResultatsInnerWithDefaults instantiates a new RechercheBeneficiaires200ResponseResultatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetNomUsage

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNomUsage() string`

GetNomUsage returns the NomUsage field if non-nil, zero value otherwise.

### GetNomUsageOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNomUsageOk() (*string, bool)`

GetNomUsageOk returns a tuple with the NomUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomUsage

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNomUsage(v string)`

SetNomUsage sets NomUsage field to given value.

### HasNomUsage

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNomUsage() bool`

HasNomUsage returns a boolean if a field has been set.

### GetPrenom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetPseudonyme

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPseudonyme() string`

GetPseudonyme returns the Pseudonyme field if non-nil, zero value otherwise.

### GetPseudonymeOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPseudonymeOk() (*string, bool)`

GetPseudonymeOk returns a tuple with the Pseudonyme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonyme

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPseudonyme(v string)`

SetPseudonyme sets Pseudonyme field to given value.

### HasPseudonyme

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPseudonyme() bool`

HasPseudonyme returns a boolean if a field has been set.

### GetNomComplet

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNomComplet() string`

GetNomComplet returns the NomComplet field if non-nil, zero value otherwise.

### GetNomCompletOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNomCompletOk() (*string, bool)`

GetNomCompletOk returns a tuple with the NomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomComplet

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNomComplet(v string)`

SetNomComplet sets NomComplet field to given value.

### HasNomComplet

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNomComplet() bool`

HasNomComplet returns a boolean if a field has been set.

### GetDateDeNaissanceFormate

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceFormate() string`

GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceFormateOk() (*string, bool)`

GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormate

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDateDeNaissanceFormate(v string)`

SetDateDeNaissanceFormate sets DateDeNaissanceFormate field to given value.

### HasDateDeNaissanceFormate

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDateDeNaissanceFormate() bool`

HasDateDeNaissanceFormate returns a boolean if a field has been set.

### GetNationalite

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNationalite() string`

GetNationalite returns the Nationalite field if non-nil, zero value otherwise.

### GetNationaliteOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNationaliteOk() (*string, bool)`

GetNationaliteOk returns a tuple with the Nationalite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalite

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNationalite(v string)`

SetNationalite sets Nationalite field to given value.

### HasNationalite

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNationalite() bool`

HasNationalite returns a boolean if a field has been set.

### GetPourcentageParts

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageParts() float32`

GetPourcentageParts returns the PourcentageParts field if non-nil, zero value otherwise.

### GetPourcentagePartsOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsOk() (*float32, bool)`

GetPourcentagePartsOk returns a tuple with the PourcentageParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageParts

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentageParts(v float32)`

SetPourcentageParts sets PourcentageParts field to given value.

### HasPourcentageParts

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentageParts() bool`

HasPourcentageParts returns a boolean if a field has been set.

### GetPourcentageVotes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotes() float32`

GetPourcentageVotes returns the PourcentageVotes field if non-nil, zero value otherwise.

### GetPourcentageVotesOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotesOk() (*float32, bool)`

GetPourcentageVotesOk returns a tuple with the PourcentageVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentageVotes(v float32)`

SetPourcentageVotes sets PourcentageVotes field to given value.

### HasPourcentageVotes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentageVotes() bool`

HasPourcentageVotes returns a boolean if a field has been set.

### GetPourcentageVotesIndirect

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotesIndirect() float32`

GetPourcentageVotesIndirect returns the PourcentageVotesIndirect field if non-nil, zero value otherwise.

### GetPourcentageVotesIndirectOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotesIndirectOk() (*float32, bool)`

GetPourcentageVotesIndirectOk returns a tuple with the PourcentageVotesIndirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesIndirect

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentageVotesIndirect(v float32)`

SetPourcentageVotesIndirect sets PourcentageVotesIndirect field to given value.

### HasPourcentageVotesIndirect

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentageVotesIndirect() bool`

HasPourcentageVotesIndirect returns a boolean if a field has been set.

### GetPourcentageVotesDirects

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotesDirects() float32`

GetPourcentageVotesDirects returns the PourcentageVotesDirects field if non-nil, zero value otherwise.

### GetPourcentageVotesDirectsOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentageVotesDirectsOk() (*float32, bool)`

GetPourcentageVotesDirectsOk returns a tuple with the PourcentageVotesDirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentageVotesDirects

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentageVotesDirects(v float32)`

SetPourcentageVotesDirects sets PourcentageVotesDirects field to given value.

### HasPourcentageVotesDirects

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentageVotesDirects() bool`

HasPourcentageVotesDirects returns a boolean if a field has been set.

### GetDetentionAutresMoyensControle

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionAutresMoyensControle() bool`

GetDetentionAutresMoyensControle returns the DetentionAutresMoyensControle field if non-nil, zero value otherwise.

### GetDetentionAutresMoyensControleOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionAutresMoyensControleOk() (*bool, bool)`

GetDetentionAutresMoyensControleOk returns a tuple with the DetentionAutresMoyensControle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionAutresMoyensControle

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDetentionAutresMoyensControle(v bool)`

SetDetentionAutresMoyensControle sets DetentionAutresMoyensControle field to given value.

### HasDetentionAutresMoyensControle

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDetentionAutresMoyensControle() bool`

HasDetentionAutresMoyensControle returns a boolean if a field has been set.

### GetBeneficiaireRepresentantLegal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetBeneficiaireRepresentantLegal() bool`

GetBeneficiaireRepresentantLegal returns the BeneficiaireRepresentantLegal field if non-nil, zero value otherwise.

### GetBeneficiaireRepresentantLegalOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetBeneficiaireRepresentantLegalOk() (*bool, bool)`

GetBeneficiaireRepresentantLegalOk returns a tuple with the BeneficiaireRepresentantLegal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaireRepresentantLegal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetBeneficiaireRepresentantLegal(v bool)`

SetBeneficiaireRepresentantLegal sets BeneficiaireRepresentantLegal field to given value.

### HasBeneficiaireRepresentantLegal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasBeneficiaireRepresentantLegal() bool`

HasBeneficiaireRepresentantLegal returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetAdresseLigne3

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne3() string`

GetAdresseLigne3 returns the AdresseLigne3 field if non-nil, zero value otherwise.

### GetAdresseLigne3Ok

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetAdresseLigne3Ok() (*string, bool)`

GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne3

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetAdresseLigne3(v string)`

SetAdresseLigne3 sets AdresseLigne3 field to given value.

### HasAdresseLigne3

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasAdresseLigne3() bool`

HasAdresseLigne3 returns a boolean if a field has been set.

### GetPourcentagePartsVocationTitulaire

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsVocationTitulaire() float32`

GetPourcentagePartsVocationTitulaire returns the PourcentagePartsVocationTitulaire field if non-nil, zero value otherwise.

### GetPourcentagePartsVocationTitulaireOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsVocationTitulaireOk() (*float32, bool)`

GetPourcentagePartsVocationTitulaireOk returns a tuple with the PourcentagePartsVocationTitulaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsVocationTitulaire

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentagePartsVocationTitulaire(v float32)`

SetPourcentagePartsVocationTitulaire sets PourcentagePartsVocationTitulaire field to given value.

### HasPourcentagePartsVocationTitulaire

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentagePartsVocationTitulaire() bool`

HasPourcentagePartsVocationTitulaire returns a boolean if a field has been set.

### GetRepresentantLegalPlacementSansGestionDelegation

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetRepresentantLegalPlacementSansGestionDelegation() bool`

GetRepresentantLegalPlacementSansGestionDelegation returns the RepresentantLegalPlacementSansGestionDelegation field if non-nil, zero value otherwise.

### GetRepresentantLegalPlacementSansGestionDelegationOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetRepresentantLegalPlacementSansGestionDelegationOk() (*bool, bool)`

GetRepresentantLegalPlacementSansGestionDelegationOk returns a tuple with the RepresentantLegalPlacementSansGestionDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentantLegalPlacementSansGestionDelegation

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetRepresentantLegalPlacementSansGestionDelegation(v bool)`

SetRepresentantLegalPlacementSansGestionDelegation sets RepresentantLegalPlacementSansGestionDelegation field to given value.

### HasRepresentantLegalPlacementSansGestionDelegation

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasRepresentantLegalPlacementSansGestionDelegation() bool`

HasRepresentantLegalPlacementSansGestionDelegation returns a boolean if a field has been set.

### GetCodePostal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetDetentionPouvoirNomMembreConseilAdministration

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionPouvoirNomMembreConseilAdministration() bool`

GetDetentionPouvoirNomMembreConseilAdministration returns the DetentionPouvoirNomMembreConseilAdministration field if non-nil, zero value otherwise.

### GetDetentionPouvoirNomMembreConseilAdministrationOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionPouvoirNomMembreConseilAdministrationOk() (*bool, bool)`

GetDetentionPouvoirNomMembreConseilAdministrationOk returns a tuple with the DetentionPouvoirNomMembreConseilAdministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirNomMembreConseilAdministration

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDetentionPouvoirNomMembreConseilAdministration(v bool)`

SetDetentionPouvoirNomMembreConseilAdministration sets DetentionPouvoirNomMembreConseilAdministration field to given value.

### HasDetentionPouvoirNomMembreConseilAdministration

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDetentionPouvoirNomMembreConseilAdministration() bool`

HasDetentionPouvoirNomMembreConseilAdministration returns a boolean if a field has been set.

### GetVille

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetDateDeNaissanceCompleteFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceCompleteFormatee() string`

GetDateDeNaissanceCompleteFormatee returns the DateDeNaissanceCompleteFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceCompleteFormateeOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceCompleteFormateeOk() (*string, bool)`

GetDateDeNaissanceCompleteFormateeOk returns a tuple with the DateDeNaissanceCompleteFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceCompleteFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDateDeNaissanceCompleteFormatee(v string)`

SetDateDeNaissanceCompleteFormatee sets DateDeNaissanceCompleteFormatee field to given value.

### HasDateDeNaissanceCompleteFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDateDeNaissanceCompleteFormatee() bool`

HasDateDeNaissanceCompleteFormatee returns a boolean if a field has been set.

### GetPourcentagePartsDirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsDirectes() float32`

GetPourcentagePartsDirectes returns the PourcentagePartsDirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsDirectesOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsDirectesOk() (*float32, bool)`

GetPourcentagePartsDirectesOk returns a tuple with the PourcentagePartsDirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsDirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentagePartsDirectes(v float32)`

SetPourcentagePartsDirectes sets PourcentagePartsDirectes field to given value.

### HasPourcentagePartsDirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentagePartsDirectes() bool`

HasPourcentagePartsDirectes returns a boolean if a field has been set.

### GetPourcentagePartsIndirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsIndirectes() float32`

GetPourcentagePartsIndirectes returns the PourcentagePartsIndirectes field if non-nil, zero value otherwise.

### GetPourcentagePartsIndirectesOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPourcentagePartsIndirectesOk() (*float32, bool)`

GetPourcentagePartsIndirectesOk returns a tuple with the PourcentagePartsIndirectes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentagePartsIndirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPourcentagePartsIndirectes(v float32)`

SetPourcentagePartsIndirectes sets PourcentagePartsIndirectes field to given value.

### HasPourcentagePartsIndirectes

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPourcentagePartsIndirectes() bool`

HasPourcentagePartsIndirectes returns a boolean if a field has been set.

### GetPaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPaysDeNaissance() string`

GetPaysDeNaissance returns the PaysDeNaissance field if non-nil, zero value otherwise.

### GetPaysDeNaissanceOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPaysDeNaissanceOk() (*string, bool)`

GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPaysDeNaissance(v string)`

SetPaysDeNaissance sets PaysDeNaissance field to given value.

### HasPaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPaysDeNaissance() bool`

HasPaysDeNaissance returns a boolean if a field has been set.

### GetCodePaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePaysDeNaissance() string`

GetCodePaysDeNaissance returns the CodePaysDeNaissance field if non-nil, zero value otherwise.

### GetCodePaysDeNaissanceOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePaysDeNaissanceOk() (*string, bool)`

GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetCodePaysDeNaissance(v string)`

SetCodePaysDeNaissance sets CodePaysDeNaissance field to given value.

### HasCodePaysDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasCodePaysDeNaissance() bool`

HasCodePaysDeNaissance returns a boolean if a field has been set.

### GetVilleDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetVilleDeNaissance() string`

GetVilleDeNaissance returns the VilleDeNaissance field if non-nil, zero value otherwise.

### GetVilleDeNaissanceOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetVilleDeNaissanceOk() (*string, bool)`

GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilleDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetVilleDeNaissance(v string)`

SetVilleDeNaissance sets VilleDeNaissance field to given value.

### HasVilleDeNaissance

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasVilleDeNaissance() bool`

HasVilleDeNaissance returns a boolean if a field has been set.

### GetDetentionPouvoirDecisionAg

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionPouvoirDecisionAg() bool`

GetDetentionPouvoirDecisionAg returns the DetentionPouvoirDecisionAg field if non-nil, zero value otherwise.

### GetDetentionPouvoirDecisionAgOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDetentionPouvoirDecisionAgOk() (*bool, bool)`

GetDetentionPouvoirDecisionAgOk returns a tuple with the DetentionPouvoirDecisionAg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetentionPouvoirDecisionAg

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDetentionPouvoirDecisionAg(v bool)`

SetDetentionPouvoirDecisionAg sets DetentionPouvoirDecisionAg field to given value.

### HasDetentionPouvoirDecisionAg

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDetentionPouvoirDecisionAg() bool`

HasDetentionPouvoirDecisionAg returns a boolean if a field has been set.

### GetPays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetDateDeNaissanceFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceFormatee() string`

GetDateDeNaissanceFormatee returns the DateDeNaissanceFormatee field if non-nil, zero value otherwise.

### GetDateDeNaissanceFormateeOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetDateDeNaissanceFormateeOk() (*string, bool)`

GetDateDeNaissanceFormateeOk returns a tuple with the DateDeNaissanceFormatee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissanceFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetDateDeNaissanceFormatee(v string)`

SetDateDeNaissanceFormatee sets DateDeNaissanceFormatee field to given value.

### HasDateDeNaissanceFormatee

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasDateDeNaissanceFormatee() bool`

HasDateDeNaissanceFormatee returns a boolean if a field has been set.

### GetCodePays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetEntreprises

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetEntreprises() []EntrepriseRecherche`

GetEntreprises returns the Entreprises field if non-nil, zero value otherwise.

### GetEntreprisesOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetEntreprisesOk() (*[]EntrepriseRecherche, bool)`

GetEntreprisesOk returns a tuple with the Entreprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntreprises

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetEntreprises(v []EntrepriseRecherche)`

SetEntreprises sets Entreprises field to given value.

### HasEntreprises

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasEntreprises() bool`

HasEntreprises returns a boolean if a field has been set.

### GetNbEntreprisesTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNbEntreprisesTotal() int32`

GetNbEntreprisesTotal returns the NbEntreprisesTotal field if non-nil, zero value otherwise.

### GetNbEntreprisesTotalOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNbEntreprisesTotalOk() (*int32, bool)`

GetNbEntreprisesTotalOk returns a tuple with the NbEntreprisesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbEntreprisesTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNbEntreprisesTotal(v int32)`

SetNbEntreprisesTotal sets NbEntreprisesTotal field to given value.

### HasNbEntreprisesTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNbEntreprisesTotal() bool`

HasNbEntreprisesTotal returns a boolean if a field has been set.

### GetEntreprisesDirigeant

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetEntreprisesDirigeant() []EntrepriseRecherche`

GetEntreprisesDirigeant returns the EntreprisesDirigeant field if non-nil, zero value otherwise.

### GetEntreprisesDirigeantOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetEntreprisesDirigeantOk() (*[]EntrepriseRecherche, bool)`

GetEntreprisesDirigeantOk returns a tuple with the EntreprisesDirigeant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntreprisesDirigeant

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetEntreprisesDirigeant(v []EntrepriseRecherche)`

SetEntreprisesDirigeant sets EntreprisesDirigeant field to given value.

### HasEntreprisesDirigeant

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasEntreprisesDirigeant() bool`

HasEntreprisesDirigeant returns a boolean if a field has been set.

### GetNbEntreprisesDirigeantTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNbEntreprisesDirigeantTotal() int32`

GetNbEntreprisesDirigeantTotal returns the NbEntreprisesDirigeantTotal field if non-nil, zero value otherwise.

### GetNbEntreprisesDirigeantTotalOk

`func (o *RechercheBeneficiaires200ResponseResultatsInner) GetNbEntreprisesDirigeantTotalOk() (*int32, bool)`

GetNbEntreprisesDirigeantTotalOk returns a tuple with the NbEntreprisesDirigeantTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbEntreprisesDirigeantTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) SetNbEntreprisesDirigeantTotal(v int32)`

SetNbEntreprisesDirigeantTotal sets NbEntreprisesDirigeantTotal field to given value.

### HasNbEntreprisesDirigeantTotal

`func (o *RechercheBeneficiaires200ResponseResultatsInner) HasNbEntreprisesDirigeantTotal() bool`

HasNbEntreprisesDirigeantTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


