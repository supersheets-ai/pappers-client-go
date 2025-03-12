# Recherche200ResponseResultatsInner

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
**Dirigeants** | Pointer to [**[]RepresentantRecherche**](RepresentantRecherche.md) | Dirigeants de l&#39;entreprise qui correspondent à la recherche (si vous recherchez dans la base dirigeants). | [optional] 
**Beneficiaires** | Pointer to [**[]Beneficiaire**](Beneficiaire.md) | Bénéficiaires effectifs de l&#39;entreprise qui correspondent à la recherche (si vous recherchez dans la base bénéficiaires effectifs). | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) | Documents de l&#39;entreprise qui correspondent à la recherche (si vous recherchez dans la base documents). | [optional] 
**Publications** | Pointer to [**[]Recherche200ResponseResultatsInnerAllOfPublicationsInner**](Recherche200ResponseResultatsInnerAllOfPublicationsInner.md) | Publications de l&#39;entreprise qui correspondent à la recherche (si vous recherchez dans la base publications). | [optional] 
**NbDirigeantsTotal** | Pointer to **int32** | Nombre de dirigeants total de l&#39;entreprise. | [optional] 
**NbBeneficiairesTotal** | Pointer to **int32** | Nombre de bénéficiaires effectifs total de l&#39;entreprise. | [optional] 
**NbDocumentsAvecMentions** | Pointer to **int32** | Nombre de documents de l&#39;entreprise qui correspondent à la recherche. | [optional] 
**NbDocumentsTotal** | Pointer to **int32** | Nombre de documents total de l&#39;entreprise. | [optional] 
**NbPublicationsAvecMentions** | Pointer to **int32** | Nombre de publications de l&#39;entreprise qui correspondent à la recherche. | [optional] 
**NbPublicationsTotal** | Pointer to **int32** | Nombre de publications total de l&#39;entreprise. | [optional] 

## Methods

### NewRecherche200ResponseResultatsInner

`func NewRecherche200ResponseResultatsInner() *Recherche200ResponseResultatsInner`

NewRecherche200ResponseResultatsInner instantiates a new Recherche200ResponseResultatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecherche200ResponseResultatsInnerWithDefaults

`func NewRecherche200ResponseResultatsInnerWithDefaults() *Recherche200ResponseResultatsInner`

NewRecherche200ResponseResultatsInnerWithDefaults instantiates a new Recherche200ResponseResultatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *Recherche200ResponseResultatsInner) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *Recherche200ResponseResultatsInner) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *Recherche200ResponseResultatsInner) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *Recherche200ResponseResultatsInner) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSirenFormate

`func (o *Recherche200ResponseResultatsInner) GetSirenFormate() string`

GetSirenFormate returns the SirenFormate field if non-nil, zero value otherwise.

### GetSirenFormateOk

`func (o *Recherche200ResponseResultatsInner) GetSirenFormateOk() (*string, bool)`

GetSirenFormateOk returns a tuple with the SirenFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenFormate

`func (o *Recherche200ResponseResultatsInner) SetSirenFormate(v string)`

SetSirenFormate sets SirenFormate field to given value.

### HasSirenFormate

`func (o *Recherche200ResponseResultatsInner) HasSirenFormate() bool`

HasSirenFormate returns a boolean if a field has been set.

### GetOppositionUtilisationCommerciale

`func (o *Recherche200ResponseResultatsInner) GetOppositionUtilisationCommerciale() bool`

GetOppositionUtilisationCommerciale returns the OppositionUtilisationCommerciale field if non-nil, zero value otherwise.

### GetOppositionUtilisationCommercialeOk

`func (o *Recherche200ResponseResultatsInner) GetOppositionUtilisationCommercialeOk() (*bool, bool)`

GetOppositionUtilisationCommercialeOk returns a tuple with the OppositionUtilisationCommerciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositionUtilisationCommerciale

`func (o *Recherche200ResponseResultatsInner) SetOppositionUtilisationCommerciale(v bool)`

SetOppositionUtilisationCommerciale sets OppositionUtilisationCommerciale field to given value.

### HasOppositionUtilisationCommerciale

`func (o *Recherche200ResponseResultatsInner) HasOppositionUtilisationCommerciale() bool`

HasOppositionUtilisationCommerciale returns a boolean if a field has been set.

### GetNomEntreprise

`func (o *Recherche200ResponseResultatsInner) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *Recherche200ResponseResultatsInner) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *Recherche200ResponseResultatsInner) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *Recherche200ResponseResultatsInner) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *Recherche200ResponseResultatsInner) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *Recherche200ResponseResultatsInner) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *Recherche200ResponseResultatsInner) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *Recherche200ResponseResultatsInner) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *Recherche200ResponseResultatsInner) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *Recherche200ResponseResultatsInner) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *Recherche200ResponseResultatsInner) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *Recherche200ResponseResultatsInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *Recherche200ResponseResultatsInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Recherche200ResponseResultatsInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Recherche200ResponseResultatsInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Recherche200ResponseResultatsInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *Recherche200ResponseResultatsInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *Recherche200ResponseResultatsInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *Recherche200ResponseResultatsInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *Recherche200ResponseResultatsInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetSexe

`func (o *Recherche200ResponseResultatsInner) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *Recherche200ResponseResultatsInner) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *Recherche200ResponseResultatsInner) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *Recherche200ResponseResultatsInner) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetCodeNaf

`func (o *Recherche200ResponseResultatsInner) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *Recherche200ResponseResultatsInner) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *Recherche200ResponseResultatsInner) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *Recherche200ResponseResultatsInner) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *Recherche200ResponseResultatsInner) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *Recherche200ResponseResultatsInner) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *Recherche200ResponseResultatsInner) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *Recherche200ResponseResultatsInner) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetDomaineActivite

`func (o *Recherche200ResponseResultatsInner) GetDomaineActivite() string`

GetDomaineActivite returns the DomaineActivite field if non-nil, zero value otherwise.

### GetDomaineActiviteOk

`func (o *Recherche200ResponseResultatsInner) GetDomaineActiviteOk() (*string, bool)`

GetDomaineActiviteOk returns a tuple with the DomaineActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaineActivite

`func (o *Recherche200ResponseResultatsInner) SetDomaineActivite(v string)`

SetDomaineActivite sets DomaineActivite field to given value.

### HasDomaineActivite

`func (o *Recherche200ResponseResultatsInner) HasDomaineActivite() bool`

HasDomaineActivite returns a boolean if a field has been set.

### GetConventionsCollectives

`func (o *Recherche200ResponseResultatsInner) GetConventionsCollectives() []EntrepriseBaseConventionsCollectivesInner`

GetConventionsCollectives returns the ConventionsCollectives field if non-nil, zero value otherwise.

### GetConventionsCollectivesOk

`func (o *Recherche200ResponseResultatsInner) GetConventionsCollectivesOk() (*[]EntrepriseBaseConventionsCollectivesInner, bool)`

GetConventionsCollectivesOk returns a tuple with the ConventionsCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionsCollectives

`func (o *Recherche200ResponseResultatsInner) SetConventionsCollectives(v []EntrepriseBaseConventionsCollectivesInner)`

SetConventionsCollectives sets ConventionsCollectives field to given value.

### HasConventionsCollectives

`func (o *Recherche200ResponseResultatsInner) HasConventionsCollectives() bool`

HasConventionsCollectives returns a boolean if a field has been set.

### GetDateCreation

`func (o *Recherche200ResponseResultatsInner) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *Recherche200ResponseResultatsInner) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *Recherche200ResponseResultatsInner) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *Recherche200ResponseResultatsInner) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateCreationFormate

`func (o *Recherche200ResponseResultatsInner) GetDateCreationFormate() string`

GetDateCreationFormate returns the DateCreationFormate field if non-nil, zero value otherwise.

### GetDateCreationFormateOk

`func (o *Recherche200ResponseResultatsInner) GetDateCreationFormateOk() (*string, bool)`

GetDateCreationFormateOk returns a tuple with the DateCreationFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreationFormate

`func (o *Recherche200ResponseResultatsInner) SetDateCreationFormate(v string)`

SetDateCreationFormate sets DateCreationFormate field to given value.

### HasDateCreationFormate

`func (o *Recherche200ResponseResultatsInner) HasDateCreationFormate() bool`

HasDateCreationFormate returns a boolean if a field has been set.

### GetEntrepriseCessee

`func (o *Recherche200ResponseResultatsInner) GetEntrepriseCessee() bool`

GetEntrepriseCessee returns the EntrepriseCessee field if non-nil, zero value otherwise.

### GetEntrepriseCesseeOk

`func (o *Recherche200ResponseResultatsInner) GetEntrepriseCesseeOk() (*bool, bool)`

GetEntrepriseCesseeOk returns a tuple with the EntrepriseCessee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseCessee

`func (o *Recherche200ResponseResultatsInner) SetEntrepriseCessee(v bool)`

SetEntrepriseCessee sets EntrepriseCessee field to given value.

### HasEntrepriseCessee

`func (o *Recherche200ResponseResultatsInner) HasEntrepriseCessee() bool`

HasEntrepriseCessee returns a boolean if a field has been set.

### GetDateCessation

`func (o *Recherche200ResponseResultatsInner) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *Recherche200ResponseResultatsInner) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *Recherche200ResponseResultatsInner) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *Recherche200ResponseResultatsInner) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEntrepriseEmployeuse

`func (o *Recherche200ResponseResultatsInner) GetEntrepriseEmployeuse() bool`

GetEntrepriseEmployeuse returns the EntrepriseEmployeuse field if non-nil, zero value otherwise.

### GetEntrepriseEmployeuseOk

`func (o *Recherche200ResponseResultatsInner) GetEntrepriseEmployeuseOk() (*bool, bool)`

GetEntrepriseEmployeuseOk returns a tuple with the EntrepriseEmployeuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseEmployeuse

`func (o *Recherche200ResponseResultatsInner) SetEntrepriseEmployeuse(v bool)`

SetEntrepriseEmployeuse sets EntrepriseEmployeuse field to given value.

### HasEntrepriseEmployeuse

`func (o *Recherche200ResponseResultatsInner) HasEntrepriseEmployeuse() bool`

HasEntrepriseEmployeuse returns a boolean if a field has been set.

### GetSocieteAMission

`func (o *Recherche200ResponseResultatsInner) GetSocieteAMission() bool`

GetSocieteAMission returns the SocieteAMission field if non-nil, zero value otherwise.

### GetSocieteAMissionOk

`func (o *Recherche200ResponseResultatsInner) GetSocieteAMissionOk() (*bool, bool)`

GetSocieteAMissionOk returns a tuple with the SocieteAMission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieteAMission

`func (o *Recherche200ResponseResultatsInner) SetSocieteAMission(v bool)`

SetSocieteAMission sets SocieteAMission field to given value.

### HasSocieteAMission

`func (o *Recherche200ResponseResultatsInner) HasSocieteAMission() bool`

HasSocieteAMission returns a boolean if a field has been set.

### GetCategorieJuridique

`func (o *Recherche200ResponseResultatsInner) GetCategorieJuridique() string`

GetCategorieJuridique returns the CategorieJuridique field if non-nil, zero value otherwise.

### GetCategorieJuridiqueOk

`func (o *Recherche200ResponseResultatsInner) GetCategorieJuridiqueOk() (*string, bool)`

GetCategorieJuridiqueOk returns a tuple with the CategorieJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieJuridique

`func (o *Recherche200ResponseResultatsInner) SetCategorieJuridique(v string)`

SetCategorieJuridique sets CategorieJuridique field to given value.

### HasCategorieJuridique

`func (o *Recherche200ResponseResultatsInner) HasCategorieJuridique() bool`

HasCategorieJuridique returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *Recherche200ResponseResultatsInner) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *Recherche200ResponseResultatsInner) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *Recherche200ResponseResultatsInner) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *Recherche200ResponseResultatsInner) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetMicroEntreprise

`func (o *Recherche200ResponseResultatsInner) GetMicroEntreprise() bool`

GetMicroEntreprise returns the MicroEntreprise field if non-nil, zero value otherwise.

### GetMicroEntrepriseOk

`func (o *Recherche200ResponseResultatsInner) GetMicroEntrepriseOk() (*bool, bool)`

GetMicroEntrepriseOk returns a tuple with the MicroEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroEntreprise

`func (o *Recherche200ResponseResultatsInner) SetMicroEntreprise(v bool)`

SetMicroEntreprise sets MicroEntreprise field to given value.

### HasMicroEntreprise

`func (o *Recherche200ResponseResultatsInner) HasMicroEntreprise() bool`

HasMicroEntreprise returns a boolean if a field has been set.

### GetFormeExercice

`func (o *Recherche200ResponseResultatsInner) GetFormeExercice() string`

GetFormeExercice returns the FormeExercice field if non-nil, zero value otherwise.

### GetFormeExerciceOk

`func (o *Recherche200ResponseResultatsInner) GetFormeExerciceOk() (*string, bool)`

GetFormeExerciceOk returns a tuple with the FormeExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeExercice

`func (o *Recherche200ResponseResultatsInner) SetFormeExercice(v string)`

SetFormeExercice sets FormeExercice field to given value.

### HasFormeExercice

`func (o *Recherche200ResponseResultatsInner) HasFormeExercice() bool`

HasFormeExercice returns a boolean if a field has been set.

### GetEffectif

`func (o *Recherche200ResponseResultatsInner) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *Recherche200ResponseResultatsInner) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *Recherche200ResponseResultatsInner) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *Recherche200ResponseResultatsInner) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *Recherche200ResponseResultatsInner) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *Recherche200ResponseResultatsInner) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *Recherche200ResponseResultatsInner) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *Recherche200ResponseResultatsInner) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *Recherche200ResponseResultatsInner) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *Recherche200ResponseResultatsInner) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *Recherche200ResponseResultatsInner) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *Recherche200ResponseResultatsInner) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *Recherche200ResponseResultatsInner) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *Recherche200ResponseResultatsInner) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *Recherche200ResponseResultatsInner) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *Recherche200ResponseResultatsInner) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *Recherche200ResponseResultatsInner) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *Recherche200ResponseResultatsInner) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *Recherche200ResponseResultatsInner) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *Recherche200ResponseResultatsInner) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCapital

`func (o *Recherche200ResponseResultatsInner) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *Recherche200ResponseResultatsInner) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *Recherche200ResponseResultatsInner) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *Recherche200ResponseResultatsInner) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetStatutRcs

`func (o *Recherche200ResponseResultatsInner) GetStatutRcs() string`

GetStatutRcs returns the StatutRcs field if non-nil, zero value otherwise.

### GetStatutRcsOk

`func (o *Recherche200ResponseResultatsInner) GetStatutRcsOk() (*string, bool)`

GetStatutRcsOk returns a tuple with the StatutRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRcs

`func (o *Recherche200ResponseResultatsInner) SetStatutRcs(v string)`

SetStatutRcs sets StatutRcs field to given value.

### HasStatutRcs

`func (o *Recherche200ResponseResultatsInner) HasStatutRcs() bool`

HasStatutRcs returns a boolean if a field has been set.

### GetSiege

`func (o *Recherche200ResponseResultatsInner) GetSiege() EtablissementRecherche`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *Recherche200ResponseResultatsInner) GetSiegeOk() (*EtablissementRecherche, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *Recherche200ResponseResultatsInner) SetSiege(v EtablissementRecherche)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *Recherche200ResponseResultatsInner) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetVilles

`func (o *Recherche200ResponseResultatsInner) GetVilles() []string`

GetVilles returns the Villes field if non-nil, zero value otherwise.

### GetVillesOk

`func (o *Recherche200ResponseResultatsInner) GetVillesOk() (*[]string, bool)`

GetVillesOk returns a tuple with the Villes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVilles

`func (o *Recherche200ResponseResultatsInner) SetVilles(v []string)`

SetVilles sets Villes field to given value.

### HasVilles

`func (o *Recherche200ResponseResultatsInner) HasVilles() bool`

HasVilles returns a boolean if a field has been set.

### GetChiffreAffaires

`func (o *Recherche200ResponseResultatsInner) GetChiffreAffaires() int32`

GetChiffreAffaires returns the ChiffreAffaires field if non-nil, zero value otherwise.

### GetChiffreAffairesOk

`func (o *Recherche200ResponseResultatsInner) GetChiffreAffairesOk() (*int32, bool)`

GetChiffreAffairesOk returns a tuple with the ChiffreAffaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChiffreAffaires

`func (o *Recherche200ResponseResultatsInner) SetChiffreAffaires(v int32)`

SetChiffreAffaires sets ChiffreAffaires field to given value.

### HasChiffreAffaires

`func (o *Recherche200ResponseResultatsInner) HasChiffreAffaires() bool`

HasChiffreAffaires returns a boolean if a field has been set.

### GetResultat

`func (o *Recherche200ResponseResultatsInner) GetResultat() int32`

GetResultat returns the Resultat field if non-nil, zero value otherwise.

### GetResultatOk

`func (o *Recherche200ResponseResultatsInner) GetResultatOk() (*int32, bool)`

GetResultatOk returns a tuple with the Resultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultat

`func (o *Recherche200ResponseResultatsInner) SetResultat(v int32)`

SetResultat sets Resultat field to given value.

### HasResultat

`func (o *Recherche200ResponseResultatsInner) HasResultat() bool`

HasResultat returns a boolean if a field has been set.

### GetEffectifsFinances

`func (o *Recherche200ResponseResultatsInner) GetEffectifsFinances() int32`

GetEffectifsFinances returns the EffectifsFinances field if non-nil, zero value otherwise.

### GetEffectifsFinancesOk

`func (o *Recherche200ResponseResultatsInner) GetEffectifsFinancesOk() (*int32, bool)`

GetEffectifsFinancesOk returns a tuple with the EffectifsFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifsFinances

`func (o *Recherche200ResponseResultatsInner) SetEffectifsFinances(v int32)`

SetEffectifsFinances sets EffectifsFinances field to given value.

### HasEffectifsFinances

`func (o *Recherche200ResponseResultatsInner) HasEffectifsFinances() bool`

HasEffectifsFinances returns a boolean if a field has been set.

### GetAnneeFinances

`func (o *Recherche200ResponseResultatsInner) GetAnneeFinances() string`

GetAnneeFinances returns the AnneeFinances field if non-nil, zero value otherwise.

### GetAnneeFinancesOk

`func (o *Recherche200ResponseResultatsInner) GetAnneeFinancesOk() (*string, bool)`

GetAnneeFinancesOk returns a tuple with the AnneeFinances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeFinances

`func (o *Recherche200ResponseResultatsInner) SetAnneeFinances(v string)`

SetAnneeFinances sets AnneeFinances field to given value.

### HasAnneeFinances

`func (o *Recherche200ResponseResultatsInner) HasAnneeFinances() bool`

HasAnneeFinances returns a boolean if a field has been set.

### GetDirigeants

`func (o *Recherche200ResponseResultatsInner) GetDirigeants() []RepresentantRecherche`

GetDirigeants returns the Dirigeants field if non-nil, zero value otherwise.

### GetDirigeantsOk

`func (o *Recherche200ResponseResultatsInner) GetDirigeantsOk() (*[]RepresentantRecherche, bool)`

GetDirigeantsOk returns a tuple with the Dirigeants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirigeants

`func (o *Recherche200ResponseResultatsInner) SetDirigeants(v []RepresentantRecherche)`

SetDirigeants sets Dirigeants field to given value.

### HasDirigeants

`func (o *Recherche200ResponseResultatsInner) HasDirigeants() bool`

HasDirigeants returns a boolean if a field has been set.

### GetBeneficiaires

`func (o *Recherche200ResponseResultatsInner) GetBeneficiaires() []Beneficiaire`

GetBeneficiaires returns the Beneficiaires field if non-nil, zero value otherwise.

### GetBeneficiairesOk

`func (o *Recherche200ResponseResultatsInner) GetBeneficiairesOk() (*[]Beneficiaire, bool)`

GetBeneficiairesOk returns a tuple with the Beneficiaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaires

`func (o *Recherche200ResponseResultatsInner) SetBeneficiaires(v []Beneficiaire)`

SetBeneficiaires sets Beneficiaires field to given value.

### HasBeneficiaires

`func (o *Recherche200ResponseResultatsInner) HasBeneficiaires() bool`

HasBeneficiaires returns a boolean if a field has been set.

### GetDocuments

`func (o *Recherche200ResponseResultatsInner) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Recherche200ResponseResultatsInner) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Recherche200ResponseResultatsInner) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Recherche200ResponseResultatsInner) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetPublications

`func (o *Recherche200ResponseResultatsInner) GetPublications() []Recherche200ResponseResultatsInnerAllOfPublicationsInner`

GetPublications returns the Publications field if non-nil, zero value otherwise.

### GetPublicationsOk

`func (o *Recherche200ResponseResultatsInner) GetPublicationsOk() (*[]Recherche200ResponseResultatsInnerAllOfPublicationsInner, bool)`

GetPublicationsOk returns a tuple with the Publications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublications

`func (o *Recherche200ResponseResultatsInner) SetPublications(v []Recherche200ResponseResultatsInnerAllOfPublicationsInner)`

SetPublications sets Publications field to given value.

### HasPublications

`func (o *Recherche200ResponseResultatsInner) HasPublications() bool`

HasPublications returns a boolean if a field has been set.

### GetNbDirigeantsTotal

`func (o *Recherche200ResponseResultatsInner) GetNbDirigeantsTotal() int32`

GetNbDirigeantsTotal returns the NbDirigeantsTotal field if non-nil, zero value otherwise.

### GetNbDirigeantsTotalOk

`func (o *Recherche200ResponseResultatsInner) GetNbDirigeantsTotalOk() (*int32, bool)`

GetNbDirigeantsTotalOk returns a tuple with the NbDirigeantsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbDirigeantsTotal

`func (o *Recherche200ResponseResultatsInner) SetNbDirigeantsTotal(v int32)`

SetNbDirigeantsTotal sets NbDirigeantsTotal field to given value.

### HasNbDirigeantsTotal

`func (o *Recherche200ResponseResultatsInner) HasNbDirigeantsTotal() bool`

HasNbDirigeantsTotal returns a boolean if a field has been set.

### GetNbBeneficiairesTotal

`func (o *Recherche200ResponseResultatsInner) GetNbBeneficiairesTotal() int32`

GetNbBeneficiairesTotal returns the NbBeneficiairesTotal field if non-nil, zero value otherwise.

### GetNbBeneficiairesTotalOk

`func (o *Recherche200ResponseResultatsInner) GetNbBeneficiairesTotalOk() (*int32, bool)`

GetNbBeneficiairesTotalOk returns a tuple with the NbBeneficiairesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbBeneficiairesTotal

`func (o *Recherche200ResponseResultatsInner) SetNbBeneficiairesTotal(v int32)`

SetNbBeneficiairesTotal sets NbBeneficiairesTotal field to given value.

### HasNbBeneficiairesTotal

`func (o *Recherche200ResponseResultatsInner) HasNbBeneficiairesTotal() bool`

HasNbBeneficiairesTotal returns a boolean if a field has been set.

### GetNbDocumentsAvecMentions

`func (o *Recherche200ResponseResultatsInner) GetNbDocumentsAvecMentions() int32`

GetNbDocumentsAvecMentions returns the NbDocumentsAvecMentions field if non-nil, zero value otherwise.

### GetNbDocumentsAvecMentionsOk

`func (o *Recherche200ResponseResultatsInner) GetNbDocumentsAvecMentionsOk() (*int32, bool)`

GetNbDocumentsAvecMentionsOk returns a tuple with the NbDocumentsAvecMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbDocumentsAvecMentions

`func (o *Recherche200ResponseResultatsInner) SetNbDocumentsAvecMentions(v int32)`

SetNbDocumentsAvecMentions sets NbDocumentsAvecMentions field to given value.

### HasNbDocumentsAvecMentions

`func (o *Recherche200ResponseResultatsInner) HasNbDocumentsAvecMentions() bool`

HasNbDocumentsAvecMentions returns a boolean if a field has been set.

### GetNbDocumentsTotal

`func (o *Recherche200ResponseResultatsInner) GetNbDocumentsTotal() int32`

GetNbDocumentsTotal returns the NbDocumentsTotal field if non-nil, zero value otherwise.

### GetNbDocumentsTotalOk

`func (o *Recherche200ResponseResultatsInner) GetNbDocumentsTotalOk() (*int32, bool)`

GetNbDocumentsTotalOk returns a tuple with the NbDocumentsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbDocumentsTotal

`func (o *Recherche200ResponseResultatsInner) SetNbDocumentsTotal(v int32)`

SetNbDocumentsTotal sets NbDocumentsTotal field to given value.

### HasNbDocumentsTotal

`func (o *Recherche200ResponseResultatsInner) HasNbDocumentsTotal() bool`

HasNbDocumentsTotal returns a boolean if a field has been set.

### GetNbPublicationsAvecMentions

`func (o *Recherche200ResponseResultatsInner) GetNbPublicationsAvecMentions() int32`

GetNbPublicationsAvecMentions returns the NbPublicationsAvecMentions field if non-nil, zero value otherwise.

### GetNbPublicationsAvecMentionsOk

`func (o *Recherche200ResponseResultatsInner) GetNbPublicationsAvecMentionsOk() (*int32, bool)`

GetNbPublicationsAvecMentionsOk returns a tuple with the NbPublicationsAvecMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPublicationsAvecMentions

`func (o *Recherche200ResponseResultatsInner) SetNbPublicationsAvecMentions(v int32)`

SetNbPublicationsAvecMentions sets NbPublicationsAvecMentions field to given value.

### HasNbPublicationsAvecMentions

`func (o *Recherche200ResponseResultatsInner) HasNbPublicationsAvecMentions() bool`

HasNbPublicationsAvecMentions returns a boolean if a field has been set.

### GetNbPublicationsTotal

`func (o *Recherche200ResponseResultatsInner) GetNbPublicationsTotal() int32`

GetNbPublicationsTotal returns the NbPublicationsTotal field if non-nil, zero value otherwise.

### GetNbPublicationsTotalOk

`func (o *Recherche200ResponseResultatsInner) GetNbPublicationsTotalOk() (*int32, bool)`

GetNbPublicationsTotalOk returns a tuple with the NbPublicationsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPublicationsTotal

`func (o *Recherche200ResponseResultatsInner) SetNbPublicationsTotal(v int32)`

SetNbPublicationsTotal sets NbPublicationsTotal field to given value.

### HasNbPublicationsTotal

`func (o *Recherche200ResponseResultatsInner) HasNbPublicationsTotal() bool`

HasNbPublicationsTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


