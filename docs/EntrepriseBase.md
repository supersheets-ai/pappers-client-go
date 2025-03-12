# EntrepriseBase

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

## Methods

### NewEntrepriseBase

`func NewEntrepriseBase() *EntrepriseBase`

NewEntrepriseBase instantiates a new EntrepriseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseBaseWithDefaults

`func NewEntrepriseBaseWithDefaults() *EntrepriseBase`

NewEntrepriseBaseWithDefaults instantiates a new EntrepriseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *EntrepriseBase) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *EntrepriseBase) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *EntrepriseBase) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *EntrepriseBase) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSirenFormate

`func (o *EntrepriseBase) GetSirenFormate() string`

GetSirenFormate returns the SirenFormate field if non-nil, zero value otherwise.

### GetSirenFormateOk

`func (o *EntrepriseBase) GetSirenFormateOk() (*string, bool)`

GetSirenFormateOk returns a tuple with the SirenFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenFormate

`func (o *EntrepriseBase) SetSirenFormate(v string)`

SetSirenFormate sets SirenFormate field to given value.

### HasSirenFormate

`func (o *EntrepriseBase) HasSirenFormate() bool`

HasSirenFormate returns a boolean if a field has been set.

### GetOppositionUtilisationCommerciale

`func (o *EntrepriseBase) GetOppositionUtilisationCommerciale() bool`

GetOppositionUtilisationCommerciale returns the OppositionUtilisationCommerciale field if non-nil, zero value otherwise.

### GetOppositionUtilisationCommercialeOk

`func (o *EntrepriseBase) GetOppositionUtilisationCommercialeOk() (*bool, bool)`

GetOppositionUtilisationCommercialeOk returns a tuple with the OppositionUtilisationCommerciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositionUtilisationCommerciale

`func (o *EntrepriseBase) SetOppositionUtilisationCommerciale(v bool)`

SetOppositionUtilisationCommerciale sets OppositionUtilisationCommerciale field to given value.

### HasOppositionUtilisationCommerciale

`func (o *EntrepriseBase) HasOppositionUtilisationCommerciale() bool`

HasOppositionUtilisationCommerciale returns a boolean if a field has been set.

### GetNomEntreprise

`func (o *EntrepriseBase) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *EntrepriseBase) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *EntrepriseBase) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *EntrepriseBase) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *EntrepriseBase) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *EntrepriseBase) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *EntrepriseBase) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *EntrepriseBase) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *EntrepriseBase) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *EntrepriseBase) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *EntrepriseBase) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *EntrepriseBase) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *EntrepriseBase) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *EntrepriseBase) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *EntrepriseBase) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *EntrepriseBase) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *EntrepriseBase) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *EntrepriseBase) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *EntrepriseBase) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *EntrepriseBase) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetSexe

`func (o *EntrepriseBase) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *EntrepriseBase) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *EntrepriseBase) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *EntrepriseBase) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetCodeNaf

`func (o *EntrepriseBase) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *EntrepriseBase) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *EntrepriseBase) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *EntrepriseBase) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *EntrepriseBase) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *EntrepriseBase) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *EntrepriseBase) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *EntrepriseBase) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetDomaineActivite

`func (o *EntrepriseBase) GetDomaineActivite() string`

GetDomaineActivite returns the DomaineActivite field if non-nil, zero value otherwise.

### GetDomaineActiviteOk

`func (o *EntrepriseBase) GetDomaineActiviteOk() (*string, bool)`

GetDomaineActiviteOk returns a tuple with the DomaineActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaineActivite

`func (o *EntrepriseBase) SetDomaineActivite(v string)`

SetDomaineActivite sets DomaineActivite field to given value.

### HasDomaineActivite

`func (o *EntrepriseBase) HasDomaineActivite() bool`

HasDomaineActivite returns a boolean if a field has been set.

### GetConventionsCollectives

`func (o *EntrepriseBase) GetConventionsCollectives() []EntrepriseBaseConventionsCollectivesInner`

GetConventionsCollectives returns the ConventionsCollectives field if non-nil, zero value otherwise.

### GetConventionsCollectivesOk

`func (o *EntrepriseBase) GetConventionsCollectivesOk() (*[]EntrepriseBaseConventionsCollectivesInner, bool)`

GetConventionsCollectivesOk returns a tuple with the ConventionsCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionsCollectives

`func (o *EntrepriseBase) SetConventionsCollectives(v []EntrepriseBaseConventionsCollectivesInner)`

SetConventionsCollectives sets ConventionsCollectives field to given value.

### HasConventionsCollectives

`func (o *EntrepriseBase) HasConventionsCollectives() bool`

HasConventionsCollectives returns a boolean if a field has been set.

### GetDateCreation

`func (o *EntrepriseBase) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *EntrepriseBase) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *EntrepriseBase) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *EntrepriseBase) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateCreationFormate

`func (o *EntrepriseBase) GetDateCreationFormate() string`

GetDateCreationFormate returns the DateCreationFormate field if non-nil, zero value otherwise.

### GetDateCreationFormateOk

`func (o *EntrepriseBase) GetDateCreationFormateOk() (*string, bool)`

GetDateCreationFormateOk returns a tuple with the DateCreationFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreationFormate

`func (o *EntrepriseBase) SetDateCreationFormate(v string)`

SetDateCreationFormate sets DateCreationFormate field to given value.

### HasDateCreationFormate

`func (o *EntrepriseBase) HasDateCreationFormate() bool`

HasDateCreationFormate returns a boolean if a field has been set.

### GetEntrepriseCessee

`func (o *EntrepriseBase) GetEntrepriseCessee() bool`

GetEntrepriseCessee returns the EntrepriseCessee field if non-nil, zero value otherwise.

### GetEntrepriseCesseeOk

`func (o *EntrepriseBase) GetEntrepriseCesseeOk() (*bool, bool)`

GetEntrepriseCesseeOk returns a tuple with the EntrepriseCessee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseCessee

`func (o *EntrepriseBase) SetEntrepriseCessee(v bool)`

SetEntrepriseCessee sets EntrepriseCessee field to given value.

### HasEntrepriseCessee

`func (o *EntrepriseBase) HasEntrepriseCessee() bool`

HasEntrepriseCessee returns a boolean if a field has been set.

### GetDateCessation

`func (o *EntrepriseBase) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *EntrepriseBase) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *EntrepriseBase) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *EntrepriseBase) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEntrepriseEmployeuse

`func (o *EntrepriseBase) GetEntrepriseEmployeuse() bool`

GetEntrepriseEmployeuse returns the EntrepriseEmployeuse field if non-nil, zero value otherwise.

### GetEntrepriseEmployeuseOk

`func (o *EntrepriseBase) GetEntrepriseEmployeuseOk() (*bool, bool)`

GetEntrepriseEmployeuseOk returns a tuple with the EntrepriseEmployeuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseEmployeuse

`func (o *EntrepriseBase) SetEntrepriseEmployeuse(v bool)`

SetEntrepriseEmployeuse sets EntrepriseEmployeuse field to given value.

### HasEntrepriseEmployeuse

`func (o *EntrepriseBase) HasEntrepriseEmployeuse() bool`

HasEntrepriseEmployeuse returns a boolean if a field has been set.

### GetSocieteAMission

`func (o *EntrepriseBase) GetSocieteAMission() bool`

GetSocieteAMission returns the SocieteAMission field if non-nil, zero value otherwise.

### GetSocieteAMissionOk

`func (o *EntrepriseBase) GetSocieteAMissionOk() (*bool, bool)`

GetSocieteAMissionOk returns a tuple with the SocieteAMission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieteAMission

`func (o *EntrepriseBase) SetSocieteAMission(v bool)`

SetSocieteAMission sets SocieteAMission field to given value.

### HasSocieteAMission

`func (o *EntrepriseBase) HasSocieteAMission() bool`

HasSocieteAMission returns a boolean if a field has been set.

### GetCategorieJuridique

`func (o *EntrepriseBase) GetCategorieJuridique() string`

GetCategorieJuridique returns the CategorieJuridique field if non-nil, zero value otherwise.

### GetCategorieJuridiqueOk

`func (o *EntrepriseBase) GetCategorieJuridiqueOk() (*string, bool)`

GetCategorieJuridiqueOk returns a tuple with the CategorieJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieJuridique

`func (o *EntrepriseBase) SetCategorieJuridique(v string)`

SetCategorieJuridique sets CategorieJuridique field to given value.

### HasCategorieJuridique

`func (o *EntrepriseBase) HasCategorieJuridique() bool`

HasCategorieJuridique returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *EntrepriseBase) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *EntrepriseBase) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *EntrepriseBase) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *EntrepriseBase) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetMicroEntreprise

`func (o *EntrepriseBase) GetMicroEntreprise() bool`

GetMicroEntreprise returns the MicroEntreprise field if non-nil, zero value otherwise.

### GetMicroEntrepriseOk

`func (o *EntrepriseBase) GetMicroEntrepriseOk() (*bool, bool)`

GetMicroEntrepriseOk returns a tuple with the MicroEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroEntreprise

`func (o *EntrepriseBase) SetMicroEntreprise(v bool)`

SetMicroEntreprise sets MicroEntreprise field to given value.

### HasMicroEntreprise

`func (o *EntrepriseBase) HasMicroEntreprise() bool`

HasMicroEntreprise returns a boolean if a field has been set.

### GetFormeExercice

`func (o *EntrepriseBase) GetFormeExercice() string`

GetFormeExercice returns the FormeExercice field if non-nil, zero value otherwise.

### GetFormeExerciceOk

`func (o *EntrepriseBase) GetFormeExerciceOk() (*string, bool)`

GetFormeExerciceOk returns a tuple with the FormeExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeExercice

`func (o *EntrepriseBase) SetFormeExercice(v string)`

SetFormeExercice sets FormeExercice field to given value.

### HasFormeExercice

`func (o *EntrepriseBase) HasFormeExercice() bool`

HasFormeExercice returns a boolean if a field has been set.

### GetEffectif

`func (o *EntrepriseBase) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *EntrepriseBase) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *EntrepriseBase) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *EntrepriseBase) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *EntrepriseBase) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *EntrepriseBase) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *EntrepriseBase) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *EntrepriseBase) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *EntrepriseBase) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *EntrepriseBase) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *EntrepriseBase) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *EntrepriseBase) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *EntrepriseBase) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *EntrepriseBase) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *EntrepriseBase) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *EntrepriseBase) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *EntrepriseBase) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *EntrepriseBase) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *EntrepriseBase) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *EntrepriseBase) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCapital

`func (o *EntrepriseBase) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *EntrepriseBase) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *EntrepriseBase) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *EntrepriseBase) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetStatutRcs

`func (o *EntrepriseBase) GetStatutRcs() string`

GetStatutRcs returns the StatutRcs field if non-nil, zero value otherwise.

### GetStatutRcsOk

`func (o *EntrepriseBase) GetStatutRcsOk() (*string, bool)`

GetStatutRcsOk returns a tuple with the StatutRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRcs

`func (o *EntrepriseBase) SetStatutRcs(v string)`

SetStatutRcs sets StatutRcs field to given value.

### HasStatutRcs

`func (o *EntrepriseBase) HasStatutRcs() bool`

HasStatutRcs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


