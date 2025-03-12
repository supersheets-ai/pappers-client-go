# Suggestions200ResponseResultatsSirenInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | Pointer to **string** | Le numéro SIREN de l&#39;entreprise au format xxxxxxxxx. | [optional] 
**SirenFormate** | Pointer to **string** | Le numéro SIREN de l&#39;entreprise au format xxx xxx xxx. | [optional] 
**OppositionUtilisationCommerciale** | Pointer to **bool** | Indique si l&#39;entreprise s&#39;oppose à l&#39;utilisation commerciale de ses données. | [optional] 
**NomEntreprise** | Pointer to **string** | Le nom de l&#39;entreprise. Il est égal à sigle + dénomination en cas de personne morale, ou à nom + prénom en cas de personne physique. Nullable si le paramètre &#x60;integrer_diffusions_partielles&#x60; est à vrai. | [optional] 
**PersonneMorale** | Pointer to **bool** | Vrai en cas de personne morale, faux en cas de personne physique. | [optional] 
**Denomination** | Pointer to **string** | Dénomination de l&#39;entreprise si personne morale. | [optional] 
**Nom** | Pointer to **string** | Nom si personne physique. | [optional] 
**Prenom** | Pointer to **string** | Prénom si personne physique. | [optional] 
**Sexe** | Pointer to **string** | Sexe si personne physique. F pour féminin, M pour masculin. | [optional] 
**CodeNaf** | Pointer to **string** | Code NAF de l&#39;entreprise. | [optional] 
**LibelleCodeNaf** | Pointer to **string** | Libellé du code NAF de l&#39;entreprise. | [optional] 
**DomaineActivite** | Pointer to **string** | Domaine d&#39;activité de l&#39;entreprise. | [optional] 
**ConventionsCollectives** | Pointer to [**[]EntrepriseBaseConventionsCollectivesInner**](EntrepriseBaseConventionsCollectivesInner.md) | Liste des conventions collectives de l&#39;entreprise. | [optional] 
**DateCreation** | Pointer to **string** | Date de création de l&#39;entreprise au format AAAA-MM-JJ. | [optional] 
**DateCreationFormate** | Pointer to **string** | Date de création de l&#39;entreprise au format JJ/MM/AAAA. | [optional] 
**EntrepriseCessee** | Pointer to **bool** | Si vrai, l&#39;entreprise n&#39;est plus en activité. Sinon, elle est toujours en activité. | [optional] 
**DateCessation** | Pointer to **string** | Date de cessation de l&#39;entreprise au format AAAA-MM-JJ. | [optional] 
**EntrepriseEmployeuse** | Pointer to **bool** | Si vrai, l&#39;entreprise a au moins un employé. | [optional] 
**SocieteAMission** | Pointer to **bool** | Si vrai, l&#39;entreprise est société à mission. | [optional] 
**CategorieJuridique** | Pointer to **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | [optional] 
**FormeJuridique** | Pointer to **string** | Forme juridique de l&#39;entreprise. | [optional] 
**MicroEntreprise** | Pointer to **bool** | Si vrai, l&#39;entreprise est une micro-entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**FormeExercice** | Pointer to **string** | Forme d&#39;exercice de l&#39;activité principale. | [optional] 
**Effectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;entreprise. | [optional] 
**EffectifMin** | Pointer to **int32** | Effectif minimal de l&#39;entreprise. | [optional] 
**EffectifMax** | Pointer to **int32** | Effectif maximal de l&#39;entreprise. | [optional] 
**TrancheEffectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73). | [optional] 
**AnneeEffectif** | Pointer to **int32** | Année de validité des variables effectif, effectif_min et effectif_max. | [optional] 
**Capital** | Pointer to **float32** | Capital de l&#39;entreprise. | [optional] 
**StatutRcs** | Pointer to **string** | Statut de l&#39;entreprise au RCS | [optional] 
**Siege** | Pointer to [**EtablissementRecherche**](EtablissementRecherche.md) |  | [optional] 
**Villes** | Pointer to **[]string** | Liste des villes où l&#39;entreprise a au moins un établissement. | [optional] 
**ChiffreAffaires** | Pointer to **int32** | Chiffre d&#39;affaires de l&#39;entreprise. | [optional] 
**Resultat** | Pointer to **int32** | Résultat de l&#39;entreprise. | [optional] 
**EffectifsFinances** | Pointer to **int32** | Effectif de l&#39;entreprise. | [optional] 
**AnneeFinances** | Pointer to **string** | Année de correspondance des variables financières (chiffre_affaires, resultat, effectifs_finances). | [optional] 
**Mention** | Pointer to **string** | SIREN de l&#39;entreprise, avec le texte recherché encerclé d&#39;une balise HTML &#x60;&lt;em&gt;&#x60;. | [optional] 

## Methods

### NewSuggestions200ResponseResultatsSirenInner

`func NewSuggestions200ResponseResultatsSirenInner() *Suggestions200ResponseResultatsSirenInner`

NewSuggestions200ResponseResultatsSirenInner instantiates a new Suggestions200ResponseResultatsSirenInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestions200ResponseResultatsSirenInnerWithDefaults

`func NewSuggestions200ResponseResultatsSirenInnerWithDefaults() *Suggestions200ResponseResultatsSirenInner`

NewSuggestions200ResponseResultatsSirenInnerWithDefaults instantiates a new Suggestions200ResponseResultatsSirenInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *Suggestions200ResponseResultatsSirenInner) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Suggestions200ResponseResultatsSirenInner) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Suggestions200ResponseResultatsSirenInner) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSirenFormate

`func (o *Suggestions200ResponseResultatsSirenInner) GetSirenFormate() string`

GetSirenFormate returns the SirenFormate field if non-nil, zero value otherwise.

### GetSirenFormateOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetSirenFormateOk() (*string, bool)`

GetSirenFormateOk returns a tuple with the SirenFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenFormate

`func (o *Suggestions200ResponseResultatsSirenInner) SetSirenFormate(v string)`

SetSirenFormate sets SirenFormate field to given value.

### HasSirenFormate

`func (o *Suggestions200ResponseResultatsSirenInner) HasSirenFormate() bool`

HasSirenFormate returns a boolean if a field has been set.

### GetOppositionUtilisationCommerciale

`func (o *Suggestions200ResponseResultatsSirenInner) GetOppositionUtilisationCommerciale() bool`

GetOppositionUtilisationCommerciale returns the OppositionUtilisationCommerciale field if non-nil, zero value otherwise.

### GetOppositionUtilisationCommercialeOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetOppositionUtilisationCommercialeOk() (*bool, bool)`

GetOppositionUtilisationCommercialeOk returns a tuple with the OppositionUtilisationCommerciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositionUtilisationCommerciale

`func (o *Suggestions200ResponseResultatsSirenInner) SetOppositionUtilisationCommerciale(v bool)`

SetOppositionUtilisationCommerciale sets OppositionUtilisationCommerciale field to given value.

### HasOppositionUtilisationCommerciale

`func (o *Suggestions200ResponseResultatsSirenInner) HasOppositionUtilisationCommerciale() bool`

HasOppositionUtilisationCommerciale returns a boolean if a field has been set.

### GetNomEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *Suggestions200ResponseResultatsSirenInner) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *Suggestions200ResponseResultatsSirenInner) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *Suggestions200ResponseResultatsSirenInner) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *Suggestions200ResponseResultatsSirenInner) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *Suggestions200ResponseResultatsSirenInner) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *Suggestions200ResponseResultatsSirenInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *Suggestions200ResponseResultatsSirenInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Suggestions200ResponseResultatsSirenInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Suggestions200ResponseResultatsSirenInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *Suggestions200ResponseResultatsSirenInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *Suggestions200ResponseResultatsSirenInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *Suggestions200ResponseResultatsSirenInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetSexe

`func (o *Suggestions200ResponseResultatsSirenInner) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *Suggestions200ResponseResultatsSirenInner) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *Suggestions200ResponseResultatsSirenInner) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *Suggestions200ResponseResultatsSirenInner) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetDomaineActivite

`func (o *Suggestions200ResponseResultatsSirenInner) GetDomaineActivite() string`

GetDomaineActivite returns the DomaineActivite field if non-nil, zero value otherwise.

### GetDomaineActiviteOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetDomaineActiviteOk() (*string, bool)`

GetDomaineActiviteOk returns a tuple with the DomaineActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaineActivite

`func (o *Suggestions200ResponseResultatsSirenInner) SetDomaineActivite(v string)`

SetDomaineActivite sets DomaineActivite field to given value.

### HasDomaineActivite

`func (o *Suggestions200ResponseResultatsSirenInner) HasDomaineActivite() bool`

HasDomaineActivite returns a boolean if a field has been set.

### GetConventionsCollectives

`func (o *Suggestions200ResponseResultatsSirenInner) GetConventionsCollectives() []EntrepriseBaseConventionsCollectivesInner`

GetConventionsCollectives returns the ConventionsCollectives field if non-nil, zero value otherwise.

### GetConventionsCollectivesOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetConventionsCollectivesOk() (*[]EntrepriseBaseConventionsCollectivesInner, bool)`

GetConventionsCollectivesOk returns a tuple with the ConventionsCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionsCollectives

`func (o *Suggestions200ResponseResultatsSirenInner) SetConventionsCollectives(v []EntrepriseBaseConventionsCollectivesInner)`

SetConventionsCollectives sets ConventionsCollectives field to given value.

### HasConventionsCollectives

`func (o *Suggestions200ResponseResultatsSirenInner) HasConventionsCollectives() bool`

HasConventionsCollectives returns a boolean if a field has been set.

### GetDateCreation

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *Suggestions200ResponseResultatsSirenInner) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *Suggestions200ResponseResultatsSirenInner) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateCreationFormate

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCreationFormate() string`

GetDateCreationFormate returns the DateCreationFormate field if non-nil, zero value otherwise.

### GetDateCreationFormateOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCreationFormateOk() (*string, bool)`

GetDateCreationFormateOk returns a tuple with the DateCreationFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreationFormate

`func (o *Suggestions200ResponseResultatsSirenInner) SetDateCreationFormate(v string)`

SetDateCreationFormate sets DateCreationFormate field to given value.

### HasDateCreationFormate

`func (o *Suggestions200ResponseResultatsSirenInner) HasDateCreationFormate() bool`

HasDateCreationFormate returns a boolean if a field has been set.

### GetEntrepriseCessee

`func (o *Suggestions200ResponseResultatsSirenInner) GetEntrepriseCessee() bool`

GetEntrepriseCessee returns the EntrepriseCessee field if non-nil, zero value otherwise.

### GetEntrepriseCesseeOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEntrepriseCesseeOk() (*bool, bool)`

GetEntrepriseCesseeOk returns a tuple with the EntrepriseCessee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseCessee

`func (o *Suggestions200ResponseResultatsSirenInner) SetEntrepriseCessee(v bool)`

SetEntrepriseCessee sets EntrepriseCessee field to given value.

### HasEntrepriseCessee

`func (o *Suggestions200ResponseResultatsSirenInner) HasEntrepriseCessee() bool`

HasEntrepriseCessee returns a boolean if a field has been set.

### GetDateCessation

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *Suggestions200ResponseResultatsSirenInner) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *Suggestions200ResponseResultatsSirenInner) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEntrepriseEmployeuse

`func (o *Suggestions200ResponseResultatsSirenInner) GetEntrepriseEmployeuse() bool`

GetEntrepriseEmployeuse returns the EntrepriseEmployeuse field if non-nil, zero value otherwise.

### GetEntrepriseEmployeuseOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEntrepriseEmployeuseOk() (*bool, bool)`

GetEntrepriseEmployeuseOk returns a tuple with the EntrepriseEmployeuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseEmployeuse

`func (o *Suggestions200ResponseResultatsSirenInner) SetEntrepriseEmployeuse(v bool)`

SetEntrepriseEmployeuse sets EntrepriseEmployeuse field to given value.

### HasEntrepriseEmployeuse

`func (o *Suggestions200ResponseResultatsSirenInner) HasEntrepriseEmployeuse() bool`

HasEntrepriseEmployeuse returns a boolean if a field has been set.

### GetSocieteAMission

`func (o *Suggestions200ResponseResultatsSirenInner) GetSocieteAMission() bool`

GetSocieteAMission returns the SocieteAMission field if non-nil, zero value otherwise.

### GetSocieteAMissionOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetSocieteAMissionOk() (*bool, bool)`

GetSocieteAMissionOk returns a tuple with the SocieteAMission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieteAMission

`func (o *Suggestions200ResponseResultatsSirenInner) SetSocieteAMission(v bool)`

SetSocieteAMission sets SocieteAMission field to given value.

### HasSocieteAMission

`func (o *Suggestions200ResponseResultatsSirenInner) HasSocieteAMission() bool`

HasSocieteAMission returns a boolean if a field has been set.

### GetCategorieJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) GetCategorieJuridique() string`

GetCategorieJuridique returns the CategorieJuridique field if non-nil, zero value otherwise.

### GetCategorieJuridiqueOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetCategorieJuridiqueOk() (*string, bool)`

GetCategorieJuridiqueOk returns a tuple with the CategorieJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) SetCategorieJuridique(v string)`

SetCategorieJuridique sets CategorieJuridique field to given value.

### HasCategorieJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) HasCategorieJuridique() bool`

HasCategorieJuridique returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *Suggestions200ResponseResultatsSirenInner) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetMicroEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) GetMicroEntreprise() bool`

GetMicroEntreprise returns the MicroEntreprise field if non-nil, zero value otherwise.

### GetMicroEntrepriseOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetMicroEntrepriseOk() (*bool, bool)`

GetMicroEntrepriseOk returns a tuple with the MicroEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) SetMicroEntreprise(v bool)`

SetMicroEntreprise sets MicroEntreprise field to given value.

### HasMicroEntreprise

`func (o *Suggestions200ResponseResultatsSirenInner) HasMicroEntreprise() bool`

HasMicroEntreprise returns a boolean if a field has been set.

### GetFormeExercice

`func (o *Suggestions200ResponseResultatsSirenInner) GetFormeExercice() string`

GetFormeExercice returns the FormeExercice field if non-nil, zero value otherwise.

### GetFormeExerciceOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetFormeExerciceOk() (*string, bool)`

GetFormeExerciceOk returns a tuple with the FormeExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeExercice

`func (o *Suggestions200ResponseResultatsSirenInner) SetFormeExercice(v string)`

SetFormeExercice sets FormeExercice field to given value.

### HasFormeExercice

`func (o *Suggestions200ResponseResultatsSirenInner) HasFormeExercice() bool`

HasFormeExercice returns a boolean if a field has been set.

### GetEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *Suggestions200ResponseResultatsSirenInner) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *Suggestions200ResponseResultatsSirenInner) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *Suggestions200ResponseResultatsSirenInner) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *Suggestions200ResponseResultatsSirenInner) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *Suggestions200ResponseResultatsSirenInner) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCapital

`func (o *Suggestions200ResponseResultatsSirenInner) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *Suggestions200ResponseResultatsSirenInner) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *Suggestions200ResponseResultatsSirenInner) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetStatutRcs

`func (o *Suggestions200ResponseResultatsSirenInner) GetStatutRcs() string`

GetStatutRcs returns the StatutRcs field if non-nil, zero value otherwise.

### GetStatutRcsOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetStatutRcsOk() (*string, bool)`

GetStatutRcsOk returns a tuple with the StatutRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRcs

`func (o *Suggestions200ResponseResultatsSirenInner) SetStatutRcs(v string)`

SetStatutRcs sets StatutRcs field to given value.

### HasStatutRcs

`func (o *Suggestions200ResponseResultatsSirenInner) HasStatutRcs() bool`

HasStatutRcs returns a boolean if a field has been set.

### GetSiege

`func (o *Suggestions200ResponseResultatsSirenInner) GetSiege() EtablissementRecherche`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetSiegeOk() (*EtablissementRecherche, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *Suggestions200ResponseResultatsSirenInner) SetSiege(v EtablissementRecherche)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *Suggestions200ResponseResultatsSirenInner) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetVilles

`func (o *Suggestions200ResponseResultatsSirenInner) GetVilles() []string`

GetVilles returns the Villes field if non-nil, zero value otherwise.

### GetVillesOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetVillesOk() (*[]string, bool)`

GetVillesOk returns a tuple with the Villes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilles

`func (o *Suggestions200ResponseResultatsSirenInner) SetVilles(v []string)`

SetVilles sets Villes field to given value.

### HasVilles

`func (o *Suggestions200ResponseResultatsSirenInner) HasVilles() bool`

HasVilles returns a boolean if a field has been set.

### GetChiffreAffaires

`func (o *Suggestions200ResponseResultatsSirenInner) GetChiffreAffaires() int32`

GetChiffreAffaires returns the ChiffreAffaires field if non-nil, zero value otherwise.

### GetChiffreAffairesOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetChiffreAffairesOk() (*int32, bool)`

GetChiffreAffairesOk returns a tuple with the ChiffreAffaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChiffreAffaires

`func (o *Suggestions200ResponseResultatsSirenInner) SetChiffreAffaires(v int32)`

SetChiffreAffaires sets ChiffreAffaires field to given value.

### HasChiffreAffaires

`func (o *Suggestions200ResponseResultatsSirenInner) HasChiffreAffaires() bool`

HasChiffreAffaires returns a boolean if a field has been set.

### GetResultat

`func (o *Suggestions200ResponseResultatsSirenInner) GetResultat() int32`

GetResultat returns the Resultat field if non-nil, zero value otherwise.

### GetResultatOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetResultatOk() (*int32, bool)`

GetResultatOk returns a tuple with the Resultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultat

`func (o *Suggestions200ResponseResultatsSirenInner) SetResultat(v int32)`

SetResultat sets Resultat field to given value.

### HasResultat

`func (o *Suggestions200ResponseResultatsSirenInner) HasResultat() bool`

HasResultat returns a boolean if a field has been set.

### GetEffectifsFinances

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifsFinances() int32`

GetEffectifsFinances returns the EffectifsFinances field if non-nil, zero value otherwise.

### GetEffectifsFinancesOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetEffectifsFinancesOk() (*int32, bool)`

GetEffectifsFinancesOk returns a tuple with the EffectifsFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifsFinances

`func (o *Suggestions200ResponseResultatsSirenInner) SetEffectifsFinances(v int32)`

SetEffectifsFinances sets EffectifsFinances field to given value.

### HasEffectifsFinances

`func (o *Suggestions200ResponseResultatsSirenInner) HasEffectifsFinances() bool`

HasEffectifsFinances returns a boolean if a field has been set.

### GetAnneeFinances

`func (o *Suggestions200ResponseResultatsSirenInner) GetAnneeFinances() string`

GetAnneeFinances returns the AnneeFinances field if non-nil, zero value otherwise.

### GetAnneeFinancesOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetAnneeFinancesOk() (*string, bool)`

GetAnneeFinancesOk returns a tuple with the AnneeFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeFinances

`func (o *Suggestions200ResponseResultatsSirenInner) SetAnneeFinances(v string)`

SetAnneeFinances sets AnneeFinances field to given value.

### HasAnneeFinances

`func (o *Suggestions200ResponseResultatsSirenInner) HasAnneeFinances() bool`

HasAnneeFinances returns a boolean if a field has been set.

### GetMention

`func (o *Suggestions200ResponseResultatsSirenInner) GetMention() string`

GetMention returns the Mention field if non-nil, zero value otherwise.

### GetMentionOk

`func (o *Suggestions200ResponseResultatsSirenInner) GetMentionOk() (*string, bool)`

GetMentionOk returns a tuple with the Mention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMention

`func (o *Suggestions200ResponseResultatsSirenInner) SetMention(v string)`

SetMention sets Mention field to given value.

### HasMention

`func (o *Suggestions200ResponseResultatsSirenInner) HasMention() bool`

HasMention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


