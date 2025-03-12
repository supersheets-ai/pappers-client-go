# EntrepriseRecherche

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

## Methods

### NewEntrepriseRecherche

`func NewEntrepriseRecherche() *EntrepriseRecherche`

NewEntrepriseRecherche instantiates a new EntrepriseRecherche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseRechercheWithDefaults

`func NewEntrepriseRechercheWithDefaults() *EntrepriseRecherche`

NewEntrepriseRechercheWithDefaults instantiates a new EntrepriseRecherche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *EntrepriseRecherche) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *EntrepriseRecherche) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *EntrepriseRecherche) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *EntrepriseRecherche) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSirenFormate

`func (o *EntrepriseRecherche) GetSirenFormate() string`

GetSirenFormate returns the SirenFormate field if non-nil, zero value otherwise.

### GetSirenFormateOk

`func (o *EntrepriseRecherche) GetSirenFormateOk() (*string, bool)`

GetSirenFormateOk returns a tuple with the SirenFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenFormate

`func (o *EntrepriseRecherche) SetSirenFormate(v string)`

SetSirenFormate sets SirenFormate field to given value.

### HasSirenFormate

`func (o *EntrepriseRecherche) HasSirenFormate() bool`

HasSirenFormate returns a boolean if a field has been set.

### GetOppositionUtilisationCommerciale

`func (o *EntrepriseRecherche) GetOppositionUtilisationCommerciale() bool`

GetOppositionUtilisationCommerciale returns the OppositionUtilisationCommerciale field if non-nil, zero value otherwise.

### GetOppositionUtilisationCommercialeOk

`func (o *EntrepriseRecherche) GetOppositionUtilisationCommercialeOk() (*bool, bool)`

GetOppositionUtilisationCommercialeOk returns a tuple with the OppositionUtilisationCommerciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositionUtilisationCommerciale

`func (o *EntrepriseRecherche) SetOppositionUtilisationCommerciale(v bool)`

SetOppositionUtilisationCommerciale sets OppositionUtilisationCommerciale field to given value.

### HasOppositionUtilisationCommerciale

`func (o *EntrepriseRecherche) HasOppositionUtilisationCommerciale() bool`

HasOppositionUtilisationCommerciale returns a boolean if a field has been set.

### GetNomEntreprise

`func (o *EntrepriseRecherche) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *EntrepriseRecherche) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *EntrepriseRecherche) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *EntrepriseRecherche) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *EntrepriseRecherche) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *EntrepriseRecherche) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *EntrepriseRecherche) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *EntrepriseRecherche) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *EntrepriseRecherche) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *EntrepriseRecherche) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *EntrepriseRecherche) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *EntrepriseRecherche) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *EntrepriseRecherche) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *EntrepriseRecherche) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *EntrepriseRecherche) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *EntrepriseRecherche) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *EntrepriseRecherche) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *EntrepriseRecherche) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *EntrepriseRecherche) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *EntrepriseRecherche) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetSexe

`func (o *EntrepriseRecherche) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *EntrepriseRecherche) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *EntrepriseRecherche) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *EntrepriseRecherche) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetCodeNaf

`func (o *EntrepriseRecherche) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *EntrepriseRecherche) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *EntrepriseRecherche) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *EntrepriseRecherche) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *EntrepriseRecherche) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *EntrepriseRecherche) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *EntrepriseRecherche) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *EntrepriseRecherche) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetDomaineActivite

`func (o *EntrepriseRecherche) GetDomaineActivite() string`

GetDomaineActivite returns the DomaineActivite field if non-nil, zero value otherwise.

### GetDomaineActiviteOk

`func (o *EntrepriseRecherche) GetDomaineActiviteOk() (*string, bool)`

GetDomaineActiviteOk returns a tuple with the DomaineActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaineActivite

`func (o *EntrepriseRecherche) SetDomaineActivite(v string)`

SetDomaineActivite sets DomaineActivite field to given value.

### HasDomaineActivite

`func (o *EntrepriseRecherche) HasDomaineActivite() bool`

HasDomaineActivite returns a boolean if a field has been set.

### GetConventionsCollectives

`func (o *EntrepriseRecherche) GetConventionsCollectives() []EntrepriseBaseConventionsCollectivesInner`

GetConventionsCollectives returns the ConventionsCollectives field if non-nil, zero value otherwise.

### GetConventionsCollectivesOk

`func (o *EntrepriseRecherche) GetConventionsCollectivesOk() (*[]EntrepriseBaseConventionsCollectivesInner, bool)`

GetConventionsCollectivesOk returns a tuple with the ConventionsCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionsCollectives

`func (o *EntrepriseRecherche) SetConventionsCollectives(v []EntrepriseBaseConventionsCollectivesInner)`

SetConventionsCollectives sets ConventionsCollectives field to given value.

### HasConventionsCollectives

`func (o *EntrepriseRecherche) HasConventionsCollectives() bool`

HasConventionsCollectives returns a boolean if a field has been set.

### GetDateCreation

`func (o *EntrepriseRecherche) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *EntrepriseRecherche) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *EntrepriseRecherche) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *EntrepriseRecherche) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateCreationFormate

`func (o *EntrepriseRecherche) GetDateCreationFormate() string`

GetDateCreationFormate returns the DateCreationFormate field if non-nil, zero value otherwise.

### GetDateCreationFormateOk

`func (o *EntrepriseRecherche) GetDateCreationFormateOk() (*string, bool)`

GetDateCreationFormateOk returns a tuple with the DateCreationFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreationFormate

`func (o *EntrepriseRecherche) SetDateCreationFormate(v string)`

SetDateCreationFormate sets DateCreationFormate field to given value.

### HasDateCreationFormate

`func (o *EntrepriseRecherche) HasDateCreationFormate() bool`

HasDateCreationFormate returns a boolean if a field has been set.

### GetEntrepriseCessee

`func (o *EntrepriseRecherche) GetEntrepriseCessee() bool`

GetEntrepriseCessee returns the EntrepriseCessee field if non-nil, zero value otherwise.

### GetEntrepriseCesseeOk

`func (o *EntrepriseRecherche) GetEntrepriseCesseeOk() (*bool, bool)`

GetEntrepriseCesseeOk returns a tuple with the EntrepriseCessee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseCessee

`func (o *EntrepriseRecherche) SetEntrepriseCessee(v bool)`

SetEntrepriseCessee sets EntrepriseCessee field to given value.

### HasEntrepriseCessee

`func (o *EntrepriseRecherche) HasEntrepriseCessee() bool`

HasEntrepriseCessee returns a boolean if a field has been set.

### GetDateCessation

`func (o *EntrepriseRecherche) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *EntrepriseRecherche) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *EntrepriseRecherche) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *EntrepriseRecherche) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEntrepriseEmployeuse

`func (o *EntrepriseRecherche) GetEntrepriseEmployeuse() bool`

GetEntrepriseEmployeuse returns the EntrepriseEmployeuse field if non-nil, zero value otherwise.

### GetEntrepriseEmployeuseOk

`func (o *EntrepriseRecherche) GetEntrepriseEmployeuseOk() (*bool, bool)`

GetEntrepriseEmployeuseOk returns a tuple with the EntrepriseEmployeuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseEmployeuse

`func (o *EntrepriseRecherche) SetEntrepriseEmployeuse(v bool)`

SetEntrepriseEmployeuse sets EntrepriseEmployeuse field to given value.

### HasEntrepriseEmployeuse

`func (o *EntrepriseRecherche) HasEntrepriseEmployeuse() bool`

HasEntrepriseEmployeuse returns a boolean if a field has been set.

### GetSocieteAMission

`func (o *EntrepriseRecherche) GetSocieteAMission() bool`

GetSocieteAMission returns the SocieteAMission field if non-nil, zero value otherwise.

### GetSocieteAMissionOk

`func (o *EntrepriseRecherche) GetSocieteAMissionOk() (*bool, bool)`

GetSocieteAMissionOk returns a tuple with the SocieteAMission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieteAMission

`func (o *EntrepriseRecherche) SetSocieteAMission(v bool)`

SetSocieteAMission sets SocieteAMission field to given value.

### HasSocieteAMission

`func (o *EntrepriseRecherche) HasSocieteAMission() bool`

HasSocieteAMission returns a boolean if a field has been set.

### GetCategorieJuridique

`func (o *EntrepriseRecherche) GetCategorieJuridique() string`

GetCategorieJuridique returns the CategorieJuridique field if non-nil, zero value otherwise.

### GetCategorieJuridiqueOk

`func (o *EntrepriseRecherche) GetCategorieJuridiqueOk() (*string, bool)`

GetCategorieJuridiqueOk returns a tuple with the CategorieJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieJuridique

`func (o *EntrepriseRecherche) SetCategorieJuridique(v string)`

SetCategorieJuridique sets CategorieJuridique field to given value.

### HasCategorieJuridique

`func (o *EntrepriseRecherche) HasCategorieJuridique() bool`

HasCategorieJuridique returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *EntrepriseRecherche) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *EntrepriseRecherche) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *EntrepriseRecherche) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *EntrepriseRecherche) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetMicroEntreprise

`func (o *EntrepriseRecherche) GetMicroEntreprise() bool`

GetMicroEntreprise returns the MicroEntreprise field if non-nil, zero value otherwise.

### GetMicroEntrepriseOk

`func (o *EntrepriseRecherche) GetMicroEntrepriseOk() (*bool, bool)`

GetMicroEntrepriseOk returns a tuple with the MicroEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroEntreprise

`func (o *EntrepriseRecherche) SetMicroEntreprise(v bool)`

SetMicroEntreprise sets MicroEntreprise field to given value.

### HasMicroEntreprise

`func (o *EntrepriseRecherche) HasMicroEntreprise() bool`

HasMicroEntreprise returns a boolean if a field has been set.

### GetFormeExercice

`func (o *EntrepriseRecherche) GetFormeExercice() string`

GetFormeExercice returns the FormeExercice field if non-nil, zero value otherwise.

### GetFormeExerciceOk

`func (o *EntrepriseRecherche) GetFormeExerciceOk() (*string, bool)`

GetFormeExerciceOk returns a tuple with the FormeExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeExercice

`func (o *EntrepriseRecherche) SetFormeExercice(v string)`

SetFormeExercice sets FormeExercice field to given value.

### HasFormeExercice

`func (o *EntrepriseRecherche) HasFormeExercice() bool`

HasFormeExercice returns a boolean if a field has been set.

### GetEffectif

`func (o *EntrepriseRecherche) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *EntrepriseRecherche) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *EntrepriseRecherche) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *EntrepriseRecherche) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *EntrepriseRecherche) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *EntrepriseRecherche) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *EntrepriseRecherche) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *EntrepriseRecherche) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *EntrepriseRecherche) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *EntrepriseRecherche) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *EntrepriseRecherche) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *EntrepriseRecherche) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *EntrepriseRecherche) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *EntrepriseRecherche) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *EntrepriseRecherche) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *EntrepriseRecherche) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *EntrepriseRecherche) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *EntrepriseRecherche) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *EntrepriseRecherche) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *EntrepriseRecherche) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCapital

`func (o *EntrepriseRecherche) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *EntrepriseRecherche) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *EntrepriseRecherche) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *EntrepriseRecherche) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetStatutRcs

`func (o *EntrepriseRecherche) GetStatutRcs() string`

GetStatutRcs returns the StatutRcs field if non-nil, zero value otherwise.

### GetStatutRcsOk

`func (o *EntrepriseRecherche) GetStatutRcsOk() (*string, bool)`

GetStatutRcsOk returns a tuple with the StatutRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRcs

`func (o *EntrepriseRecherche) SetStatutRcs(v string)`

SetStatutRcs sets StatutRcs field to given value.

### HasStatutRcs

`func (o *EntrepriseRecherche) HasStatutRcs() bool`

HasStatutRcs returns a boolean if a field has been set.

### GetSiege

`func (o *EntrepriseRecherche) GetSiege() EtablissementRecherche`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *EntrepriseRecherche) GetSiegeOk() (*EtablissementRecherche, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *EntrepriseRecherche) SetSiege(v EtablissementRecherche)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *EntrepriseRecherche) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetVilles

`func (o *EntrepriseRecherche) GetVilles() []string`

GetVilles returns the Villes field if non-nil, zero value otherwise.

### GetVillesOk

`func (o *EntrepriseRecherche) GetVillesOk() (*[]string, bool)`

GetVillesOk returns a tuple with the Villes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilles

`func (o *EntrepriseRecherche) SetVilles(v []string)`

SetVilles sets Villes field to given value.

### HasVilles

`func (o *EntrepriseRecherche) HasVilles() bool`

HasVilles returns a boolean if a field has been set.

### GetChiffreAffaires

`func (o *EntrepriseRecherche) GetChiffreAffaires() int32`

GetChiffreAffaires returns the ChiffreAffaires field if non-nil, zero value otherwise.

### GetChiffreAffairesOk

`func (o *EntrepriseRecherche) GetChiffreAffairesOk() (*int32, bool)`

GetChiffreAffairesOk returns a tuple with the ChiffreAffaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChiffreAffaires

`func (o *EntrepriseRecherche) SetChiffreAffaires(v int32)`

SetChiffreAffaires sets ChiffreAffaires field to given value.

### HasChiffreAffaires

`func (o *EntrepriseRecherche) HasChiffreAffaires() bool`

HasChiffreAffaires returns a boolean if a field has been set.

### GetResultat

`func (o *EntrepriseRecherche) GetResultat() int32`

GetResultat returns the Resultat field if non-nil, zero value otherwise.

### GetResultatOk

`func (o *EntrepriseRecherche) GetResultatOk() (*int32, bool)`

GetResultatOk returns a tuple with the Resultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultat

`func (o *EntrepriseRecherche) SetResultat(v int32)`

SetResultat sets Resultat field to given value.

### HasResultat

`func (o *EntrepriseRecherche) HasResultat() bool`

HasResultat returns a boolean if a field has been set.

### GetEffectifsFinances

`func (o *EntrepriseRecherche) GetEffectifsFinances() int32`

GetEffectifsFinances returns the EffectifsFinances field if non-nil, zero value otherwise.

### GetEffectifsFinancesOk

`func (o *EntrepriseRecherche) GetEffectifsFinancesOk() (*int32, bool)`

GetEffectifsFinancesOk returns a tuple with the EffectifsFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifsFinances

`func (o *EntrepriseRecherche) SetEffectifsFinances(v int32)`

SetEffectifsFinances sets EffectifsFinances field to given value.

### HasEffectifsFinances

`func (o *EntrepriseRecherche) HasEffectifsFinances() bool`

HasEffectifsFinances returns a boolean if a field has been set.

### GetAnneeFinances

`func (o *EntrepriseRecherche) GetAnneeFinances() string`

GetAnneeFinances returns the AnneeFinances field if non-nil, zero value otherwise.

### GetAnneeFinancesOk

`func (o *EntrepriseRecherche) GetAnneeFinancesOk() (*string, bool)`

GetAnneeFinancesOk returns a tuple with the AnneeFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeFinances

`func (o *EntrepriseRecherche) SetAnneeFinances(v string)`

SetAnneeFinances sets AnneeFinances field to given value.

### HasAnneeFinances

`func (o *EntrepriseRecherche) HasAnneeFinances() bool`

HasAnneeFinances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


