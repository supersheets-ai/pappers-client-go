# EntrepriseFiche

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
**MicroEntreprise** | Pointer to **bool** | Si vrai, l&#39;entreprise possède le statut de micro-entrepreneur | [optional] 
**FormeExercice** | Pointer to **string** | Forme d&#39;exercice de l&#39;activité principale. | [optional] 
**Effectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;entreprise. | [optional] 
**EffectifMin** | Pointer to **int32** | Effectif minimal de l&#39;entreprise. | [optional] 
**EffectifMax** | Pointer to **int32** | Effectif maximal de l&#39;entreprise. | [optional] 
**TrancheEffectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73). | [optional] 
**AnneeEffectif** | Pointer to **int32** | Année de validité des variables effectif, effectif_min et effectif_max. | [optional] 
**Capital** | Pointer to **float32** | Capital de l&#39;entreprise. | [optional] 
**StatutRcs** | Pointer to **string** | Statut de l&#39;entreprise au RCS. | [optional] 
**Siege** | Pointer to [**EtablissementFiche**](EtablissementFiche.md) |  | [optional] 
**Diffusable** | Pointer to **bool** | Le statut de diffusion de l&#39;entreprise. Non diffusable correspond à une entreprise ayant demandé une diffusion partielle de ses données. Aucune information n&#39;est alors disponible, sauf si vous utilisez le paramètre &#x60;integrer_diffusions_partielles&#x60;. | [optional] 
**Sigle** | Pointer to **string** | Sigle de l&#39;entreprise si personne morale. | [optional] 
**ObjetSocial** | Pointer to **string** | Objet social de l&#39;entreprise. | [optional] 
**CapitalFormate** | Pointer to **string** | Capital l&#39;entreprise au format xx xxx €. | [optional] 
**CapitalActuelSiVariable** | Pointer to **string** | Capital actuel de l&#39;entreprise si variable. | [optional] 
**DeviseCapital** | Pointer to **string** | Devise de capital_formate et capital_actuel_si_variable. | [optional] 
**NumeroRcs** | Pointer to **string** | Numéro RCS de l&#39;entreprise. | [optional] 
**DateClotureExercice** | Pointer to **string** | Date de clôture d&#39;exercice de l&#39;entreprise. | [optional] 
**DateClotureExerciceExceptionnelle** | Pointer to **string** | Date de clôture d&#39;exercice exceptionnel de l&#39;entreprise. | [optional] 
**DateClotureExerciceExceptionnelleFormate** | Pointer to **string** | Date de clôture d&#39;exercice exceptionnel formatée de l&#39;entreprise. | [optional] 
**ProchaineDateClotureExercice** | Pointer to **string** | Prochaine date de clôture d&#39;exercice de l&#39;entreprise. | [optional] 
**ProchaineDateClotureExerciceFormate** | Pointer to **string** | Prochaine date de clôture d&#39;exercice formatée de l&#39;entreprise. | [optional] 
**EconomieSocialeSolidaire** | Pointer to **bool** | Vrai si l&#39;entreprise est une entreprise de l&#39;économie sociale et solidaire. | [optional] 
**DureePersonneMorale** | Pointer to **string** | Durée de la personne morale. | [optional] 
**DernierTraitement** | Pointer to **string** | Date de dernier traitement de l&#39;entreprise. | [optional] 
**DerniereMiseAJourSirene** | Pointer to **string** | Dernière mise à jour de la base de donnée sirène au format AAAA-MM-JJ. | [optional] 
**DerniereMiseAJourRcs** | Pointer to **string** | Dernière mise à jour de la base de donnée RCS au format AAAA-MM-JJ. | [optional] 
**Greffe** | Pointer to **string** | Greffe RCS de l&#39;entreprise. | [optional] 
**CodeGreffe** | Pointer to **string** | Code greffe RCS de l&#39;entreprise. | [optional] 
**DateImmatriculationRcs** | Pointer to **string** | Date d&#39;immatriculation de l&#39;entreprise au RCS. | [optional] 
**DatePremiereImmatriculationRcs** | Pointer to **string** | Date de la première immatriculation de l&#39;entreprise au RCS. | [optional] 
**DateDebutActivite** | Pointer to **string** | Date de début d&#39;activité de l&#39;entreprise. | [optional] 
**DateDebutPremiereActivite** | Pointer to **string** | Date de début d&#39;activité de l&#39;entreprise. | [optional] 
**DateRadiationRcs** | Pointer to **string** | Date de radiation de l&#39;entreprise au RCS. | [optional] 
**StatutRne** | Pointer to **string** | Statut de l&#39;entreprise au RNE. | [optional] 
**DateImmatriculationRne** | Pointer to **string** | Date d&#39;immatriculation de l&#39;entreprise au RNE. | [optional] 
**DateRadiationRne** | Pointer to **string** | Date de radiation de l&#39;entreprise au RNE. | [optional] 
**NumeroTvaIntracommunautaire** | Pointer to **string** | Numéro de TVA intracommunautaire de l&#39;entreprise. | [optional] 
**ValiditeTvaIntracommunautaire** | Pointer to **bool** | Présent uniquement si le paramètre validite_tva_intracommunautaire a été mis à vrai.  Si vrai, le numéro de TVA intracommunautaire est valide. Si faux, il est invalide. Si null, la validité n&#39;a pas pu être vérifiée. | [optional] 
**AssocieUnique** | Pointer to **bool** | Si vrai, l&#39;entreprise est à associé unique (notamment pour les SASU et les EURL). | [optional] 
**Etablissements** | Pointer to [**[]EtablissementFiche**](EtablissementFiche.md) | Liste des établissements de l&#39;entreprise. (uniquement si le paramètre &#x60;siren&#x60; est utilisé) | [optional] 
**Etablissement** | Pointer to [**EntrepriseFicheAllOfEtablissement**](EntrepriseFicheAllOfEtablissement.md) |  | [optional] 
**Finances** | Pointer to [**[]EntrepriseFicheAllOfFinances**](EntrepriseFicheAllOfFinances.md) | Liste des finances de l&#39;entreprise. | [optional] 
**Representants** | Pointer to [**[]RepresentantEntreprise**](RepresentantEntreprise.md) | Liste des représentants de l&#39;entreprise. | [optional] 
**BeneficiairesEffectifs** | Pointer to [**[]EntrepriseFicheAllOfBeneficiairesEffectifs**](EntrepriseFicheAllOfBeneficiairesEffectifs.md) | Liste des bénéficiaires effectifs de l&#39;entreprise (si disponibles). | [optional] 
**DepotsActes** | Pointer to [**[]EntrepriseFicheAllOfDepotsActes**](EntrepriseFicheAllOfDepotsActes.md) | Liste des actes de l&#39;entreprise. | [optional] 
**Comptes** | Pointer to [**[]EntrepriseFicheAllOfComptes**](EntrepriseFicheAllOfComptes.md) | Liste des comptes de l&#39;entreprise. | [optional] 
**PublicationsBodacc** | Pointer to [**[]Bodacc**](Bodacc.md) | Liste des publications au Bodacc de l&#39;entreprise. | [optional] 
**ProceduresCollectives** | Pointer to [**[]EntrepriseFicheAllOfProceduresCollectives**](EntrepriseFicheAllOfProceduresCollectives.md) | Liste des procédures collectives de l&#39;entreprise. | [optional] 
**ProcedureCollectiveExiste** | Pointer to **bool** | Vrai si l&#39;entreprise a des procédures collectives (en cours ou terminées), faux sinon. | [optional] 
**ProcedureCollectiveEnCours** | Pointer to **bool** | Vrai si l&#39;entreprise a des procédures collectives en cours, faux sinon. | [optional] 
**DerniersStatuts** | Pointer to [**EntrepriseFicheAllOfDerniersStatuts**](EntrepriseFicheAllOfDerniersStatuts.md) |  | [optional] 
**ExtraitImmatriculation** | Pointer to [**EntrepriseFicheAllOfExtraitImmatriculation**](EntrepriseFicheAllOfExtraitImmatriculation.md) |  | [optional] 
**Rnm** | Pointer to [**EntrepriseFicheAllOfRnm**](EntrepriseFicheAllOfRnm.md) |  | [optional] 
**Marques** | Pointer to [**[]EntrepriseFicheAllOfMarques**](EntrepriseFicheAllOfMarques.md) | Liste des marques françaises déposées par l&#39;entreprise. Uniquement présent si le paramètre \&quot;marques\&quot; a été mis à vrai. | [optional] 
**Association** | Pointer to [**Association**](Association.md) |  | [optional] 
**Labels** | Pointer to [**[]Labels**](Labels.md) | Liste des labels de l&#39;entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**SitesInternet** | Pointer to **[]string** | Liste des sites internet de l&#39;entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Telephone** | Pointer to **string** | Numéro de téléphone de l&#39;entreprise au format 0123456789. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Email** | Pointer to **string** | Adresse email de l&#39;entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**ScoringNonFinancier** | Pointer to [**ScoringNonFinancier**](ScoringNonFinancier.md) |  | [optional] 
**ScoringFinancier** | Pointer to [**ScoringFinancier**](ScoringFinancier.md) |  | [optional] 
**CategorieEntreprise** | Pointer to **string** | Catégorie de l&#39;entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**AnneeCategorieEntreprise** | Pointer to **int32** | Année de correspondance de la catégorie de l&#39;entreprise. Uniquement présent si le champ supplémentaire &#x60;categorie_entreprise&#x60; est demandé. | [optional] 

## Methods

### NewEntrepriseFiche

`func NewEntrepriseFiche() *EntrepriseFiche`

NewEntrepriseFiche instantiates a new EntrepriseFiche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheWithDefaults

`func NewEntrepriseFicheWithDefaults() *EntrepriseFiche`

NewEntrepriseFicheWithDefaults instantiates a new EntrepriseFiche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *EntrepriseFiche) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *EntrepriseFiche) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *EntrepriseFiche) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *EntrepriseFiche) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetSirenFormate

`func (o *EntrepriseFiche) GetSirenFormate() string`

GetSirenFormate returns the SirenFormate field if non-nil, zero value otherwise.

### GetSirenFormateOk

`func (o *EntrepriseFiche) GetSirenFormateOk() (*string, bool)`

GetSirenFormateOk returns a tuple with the SirenFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenFormate

`func (o *EntrepriseFiche) SetSirenFormate(v string)`

SetSirenFormate sets SirenFormate field to given value.

### HasSirenFormate

`func (o *EntrepriseFiche) HasSirenFormate() bool`

HasSirenFormate returns a boolean if a field has been set.

### GetOppositionUtilisationCommerciale

`func (o *EntrepriseFiche) GetOppositionUtilisationCommerciale() bool`

GetOppositionUtilisationCommerciale returns the OppositionUtilisationCommerciale field if non-nil, zero value otherwise.

### GetOppositionUtilisationCommercialeOk

`func (o *EntrepriseFiche) GetOppositionUtilisationCommercialeOk() (*bool, bool)`

GetOppositionUtilisationCommercialeOk returns a tuple with the OppositionUtilisationCommerciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOppositionUtilisationCommerciale

`func (o *EntrepriseFiche) SetOppositionUtilisationCommerciale(v bool)`

SetOppositionUtilisationCommerciale sets OppositionUtilisationCommerciale field to given value.

### HasOppositionUtilisationCommerciale

`func (o *EntrepriseFiche) HasOppositionUtilisationCommerciale() bool`

HasOppositionUtilisationCommerciale returns a boolean if a field has been set.

### GetNomEntreprise

`func (o *EntrepriseFiche) GetNomEntreprise() string`

GetNomEntreprise returns the NomEntreprise field if non-nil, zero value otherwise.

### GetNomEntrepriseOk

`func (o *EntrepriseFiche) GetNomEntrepriseOk() (*string, bool)`

GetNomEntrepriseOk returns a tuple with the NomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomEntreprise

`func (o *EntrepriseFiche) SetNomEntreprise(v string)`

SetNomEntreprise sets NomEntreprise field to given value.

### HasNomEntreprise

`func (o *EntrepriseFiche) HasNomEntreprise() bool`

HasNomEntreprise returns a boolean if a field has been set.

### GetPersonneMorale

`func (o *EntrepriseFiche) GetPersonneMorale() bool`

GetPersonneMorale returns the PersonneMorale field if non-nil, zero value otherwise.

### GetPersonneMoraleOk

`func (o *EntrepriseFiche) GetPersonneMoraleOk() (*bool, bool)`

GetPersonneMoraleOk returns a tuple with the PersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonneMorale

`func (o *EntrepriseFiche) SetPersonneMorale(v bool)`

SetPersonneMorale sets PersonneMorale field to given value.

### HasPersonneMorale

`func (o *EntrepriseFiche) HasPersonneMorale() bool`

HasPersonneMorale returns a boolean if a field has been set.

### GetDenomination

`func (o *EntrepriseFiche) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *EntrepriseFiche) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *EntrepriseFiche) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *EntrepriseFiche) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *EntrepriseFiche) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *EntrepriseFiche) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *EntrepriseFiche) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *EntrepriseFiche) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *EntrepriseFiche) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *EntrepriseFiche) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *EntrepriseFiche) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *EntrepriseFiche) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetSexe

`func (o *EntrepriseFiche) GetSexe() string`

GetSexe returns the Sexe field if non-nil, zero value otherwise.

### GetSexeOk

`func (o *EntrepriseFiche) GetSexeOk() (*string, bool)`

GetSexeOk returns a tuple with the Sexe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSexe

`func (o *EntrepriseFiche) SetSexe(v string)`

SetSexe sets Sexe field to given value.

### HasSexe

`func (o *EntrepriseFiche) HasSexe() bool`

HasSexe returns a boolean if a field has been set.

### GetCodeNaf

`func (o *EntrepriseFiche) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *EntrepriseFiche) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *EntrepriseFiche) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *EntrepriseFiche) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *EntrepriseFiche) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *EntrepriseFiche) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *EntrepriseFiche) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *EntrepriseFiche) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetDomaineActivite

`func (o *EntrepriseFiche) GetDomaineActivite() string`

GetDomaineActivite returns the DomaineActivite field if non-nil, zero value otherwise.

### GetDomaineActiviteOk

`func (o *EntrepriseFiche) GetDomaineActiviteOk() (*string, bool)`

GetDomaineActiviteOk returns a tuple with the DomaineActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomaineActivite

`func (o *EntrepriseFiche) SetDomaineActivite(v string)`

SetDomaineActivite sets DomaineActivite field to given value.

### HasDomaineActivite

`func (o *EntrepriseFiche) HasDomaineActivite() bool`

HasDomaineActivite returns a boolean if a field has been set.

### GetConventionsCollectives

`func (o *EntrepriseFiche) GetConventionsCollectives() []EntrepriseBaseConventionsCollectivesInner`

GetConventionsCollectives returns the ConventionsCollectives field if non-nil, zero value otherwise.

### GetConventionsCollectivesOk

`func (o *EntrepriseFiche) GetConventionsCollectivesOk() (*[]EntrepriseBaseConventionsCollectivesInner, bool)`

GetConventionsCollectivesOk returns a tuple with the ConventionsCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConventionsCollectives

`func (o *EntrepriseFiche) SetConventionsCollectives(v []EntrepriseBaseConventionsCollectivesInner)`

SetConventionsCollectives sets ConventionsCollectives field to given value.

### HasConventionsCollectives

`func (o *EntrepriseFiche) HasConventionsCollectives() bool`

HasConventionsCollectives returns a boolean if a field has been set.

### GetDateCreation

`func (o *EntrepriseFiche) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *EntrepriseFiche) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *EntrepriseFiche) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *EntrepriseFiche) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateCreationFormate

`func (o *EntrepriseFiche) GetDateCreationFormate() string`

GetDateCreationFormate returns the DateCreationFormate field if non-nil, zero value otherwise.

### GetDateCreationFormateOk

`func (o *EntrepriseFiche) GetDateCreationFormateOk() (*string, bool)`

GetDateCreationFormateOk returns a tuple with the DateCreationFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreationFormate

`func (o *EntrepriseFiche) SetDateCreationFormate(v string)`

SetDateCreationFormate sets DateCreationFormate field to given value.

### HasDateCreationFormate

`func (o *EntrepriseFiche) HasDateCreationFormate() bool`

HasDateCreationFormate returns a boolean if a field has been set.

### GetEntrepriseCessee

`func (o *EntrepriseFiche) GetEntrepriseCessee() bool`

GetEntrepriseCessee returns the EntrepriseCessee field if non-nil, zero value otherwise.

### GetEntrepriseCesseeOk

`func (o *EntrepriseFiche) GetEntrepriseCesseeOk() (*bool, bool)`

GetEntrepriseCesseeOk returns a tuple with the EntrepriseCessee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseCessee

`func (o *EntrepriseFiche) SetEntrepriseCessee(v bool)`

SetEntrepriseCessee sets EntrepriseCessee field to given value.

### HasEntrepriseCessee

`func (o *EntrepriseFiche) HasEntrepriseCessee() bool`

HasEntrepriseCessee returns a boolean if a field has been set.

### GetDateCessation

`func (o *EntrepriseFiche) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *EntrepriseFiche) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *EntrepriseFiche) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *EntrepriseFiche) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEntrepriseEmployeuse

`func (o *EntrepriseFiche) GetEntrepriseEmployeuse() bool`

GetEntrepriseEmployeuse returns the EntrepriseEmployeuse field if non-nil, zero value otherwise.

### GetEntrepriseEmployeuseOk

`func (o *EntrepriseFiche) GetEntrepriseEmployeuseOk() (*bool, bool)`

GetEntrepriseEmployeuseOk returns a tuple with the EntrepriseEmployeuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrepriseEmployeuse

`func (o *EntrepriseFiche) SetEntrepriseEmployeuse(v bool)`

SetEntrepriseEmployeuse sets EntrepriseEmployeuse field to given value.

### HasEntrepriseEmployeuse

`func (o *EntrepriseFiche) HasEntrepriseEmployeuse() bool`

HasEntrepriseEmployeuse returns a boolean if a field has been set.

### GetSocieteAMission

`func (o *EntrepriseFiche) GetSocieteAMission() bool`

GetSocieteAMission returns the SocieteAMission field if non-nil, zero value otherwise.

### GetSocieteAMissionOk

`func (o *EntrepriseFiche) GetSocieteAMissionOk() (*bool, bool)`

GetSocieteAMissionOk returns a tuple with the SocieteAMission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocieteAMission

`func (o *EntrepriseFiche) SetSocieteAMission(v bool)`

SetSocieteAMission sets SocieteAMission field to given value.

### HasSocieteAMission

`func (o *EntrepriseFiche) HasSocieteAMission() bool`

HasSocieteAMission returns a boolean if a field has been set.

### GetCategorieJuridique

`func (o *EntrepriseFiche) GetCategorieJuridique() string`

GetCategorieJuridique returns the CategorieJuridique field if non-nil, zero value otherwise.

### GetCategorieJuridiqueOk

`func (o *EntrepriseFiche) GetCategorieJuridiqueOk() (*string, bool)`

GetCategorieJuridiqueOk returns a tuple with the CategorieJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieJuridique

`func (o *EntrepriseFiche) SetCategorieJuridique(v string)`

SetCategorieJuridique sets CategorieJuridique field to given value.

### HasCategorieJuridique

`func (o *EntrepriseFiche) HasCategorieJuridique() bool`

HasCategorieJuridique returns a boolean if a field has been set.

### GetFormeJuridique

`func (o *EntrepriseFiche) GetFormeJuridique() string`

GetFormeJuridique returns the FormeJuridique field if non-nil, zero value otherwise.

### GetFormeJuridiqueOk

`func (o *EntrepriseFiche) GetFormeJuridiqueOk() (*string, bool)`

GetFormeJuridiqueOk returns a tuple with the FormeJuridique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeJuridique

`func (o *EntrepriseFiche) SetFormeJuridique(v string)`

SetFormeJuridique sets FormeJuridique field to given value.

### HasFormeJuridique

`func (o *EntrepriseFiche) HasFormeJuridique() bool`

HasFormeJuridique returns a boolean if a field has been set.

### GetMicroEntreprise

`func (o *EntrepriseFiche) GetMicroEntreprise() bool`

GetMicroEntreprise returns the MicroEntreprise field if non-nil, zero value otherwise.

### GetMicroEntrepriseOk

`func (o *EntrepriseFiche) GetMicroEntrepriseOk() (*bool, bool)`

GetMicroEntrepriseOk returns a tuple with the MicroEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroEntreprise

`func (o *EntrepriseFiche) SetMicroEntreprise(v bool)`

SetMicroEntreprise sets MicroEntreprise field to given value.

### HasMicroEntreprise

`func (o *EntrepriseFiche) HasMicroEntreprise() bool`

HasMicroEntreprise returns a boolean if a field has been set.

### GetFormeExercice

`func (o *EntrepriseFiche) GetFormeExercice() string`

GetFormeExercice returns the FormeExercice field if non-nil, zero value otherwise.

### GetFormeExerciceOk

`func (o *EntrepriseFiche) GetFormeExerciceOk() (*string, bool)`

GetFormeExerciceOk returns a tuple with the FormeExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormeExercice

`func (o *EntrepriseFiche) SetFormeExercice(v string)`

SetFormeExercice sets FormeExercice field to given value.

### HasFormeExercice

`func (o *EntrepriseFiche) HasFormeExercice() bool`

HasFormeExercice returns a boolean if a field has been set.

### GetEffectif

`func (o *EntrepriseFiche) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *EntrepriseFiche) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *EntrepriseFiche) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *EntrepriseFiche) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *EntrepriseFiche) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *EntrepriseFiche) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *EntrepriseFiche) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *EntrepriseFiche) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *EntrepriseFiche) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *EntrepriseFiche) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *EntrepriseFiche) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *EntrepriseFiche) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *EntrepriseFiche) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *EntrepriseFiche) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *EntrepriseFiche) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *EntrepriseFiche) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *EntrepriseFiche) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *EntrepriseFiche) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *EntrepriseFiche) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *EntrepriseFiche) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCapital

`func (o *EntrepriseFiche) GetCapital() float32`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *EntrepriseFiche) GetCapitalOk() (*float32, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *EntrepriseFiche) SetCapital(v float32)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *EntrepriseFiche) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetStatutRcs

`func (o *EntrepriseFiche) GetStatutRcs() string`

GetStatutRcs returns the StatutRcs field if non-nil, zero value otherwise.

### GetStatutRcsOk

`func (o *EntrepriseFiche) GetStatutRcsOk() (*string, bool)`

GetStatutRcsOk returns a tuple with the StatutRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRcs

`func (o *EntrepriseFiche) SetStatutRcs(v string)`

SetStatutRcs sets StatutRcs field to given value.

### HasStatutRcs

`func (o *EntrepriseFiche) HasStatutRcs() bool`

HasStatutRcs returns a boolean if a field has been set.

### GetSiege

`func (o *EntrepriseFiche) GetSiege() EtablissementFiche`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *EntrepriseFiche) GetSiegeOk() (*EtablissementFiche, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *EntrepriseFiche) SetSiege(v EtablissementFiche)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *EntrepriseFiche) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetDiffusable

`func (o *EntrepriseFiche) GetDiffusable() bool`

GetDiffusable returns the Diffusable field if non-nil, zero value otherwise.

### GetDiffusableOk

`func (o *EntrepriseFiche) GetDiffusableOk() (*bool, bool)`

GetDiffusableOk returns a tuple with the Diffusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffusable

`func (o *EntrepriseFiche) SetDiffusable(v bool)`

SetDiffusable sets Diffusable field to given value.

### HasDiffusable

`func (o *EntrepriseFiche) HasDiffusable() bool`

HasDiffusable returns a boolean if a field has been set.

### GetSigle

`func (o *EntrepriseFiche) GetSigle() string`

GetSigle returns the Sigle field if non-nil, zero value otherwise.

### GetSigleOk

`func (o *EntrepriseFiche) GetSigleOk() (*string, bool)`

GetSigleOk returns a tuple with the Sigle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigle

`func (o *EntrepriseFiche) SetSigle(v string)`

SetSigle sets Sigle field to given value.

### HasSigle

`func (o *EntrepriseFiche) HasSigle() bool`

HasSigle returns a boolean if a field has been set.

### GetObjetSocial

`func (o *EntrepriseFiche) GetObjetSocial() string`

GetObjetSocial returns the ObjetSocial field if non-nil, zero value otherwise.

### GetObjetSocialOk

`func (o *EntrepriseFiche) GetObjetSocialOk() (*string, bool)`

GetObjetSocialOk returns a tuple with the ObjetSocial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjetSocial

`func (o *EntrepriseFiche) SetObjetSocial(v string)`

SetObjetSocial sets ObjetSocial field to given value.

### HasObjetSocial

`func (o *EntrepriseFiche) HasObjetSocial() bool`

HasObjetSocial returns a boolean if a field has been set.

### GetCapitalFormate

`func (o *EntrepriseFiche) GetCapitalFormate() string`

GetCapitalFormate returns the CapitalFormate field if non-nil, zero value otherwise.

### GetCapitalFormateOk

`func (o *EntrepriseFiche) GetCapitalFormateOk() (*string, bool)`

GetCapitalFormateOk returns a tuple with the CapitalFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalFormate

`func (o *EntrepriseFiche) SetCapitalFormate(v string)`

SetCapitalFormate sets CapitalFormate field to given value.

### HasCapitalFormate

`func (o *EntrepriseFiche) HasCapitalFormate() bool`

HasCapitalFormate returns a boolean if a field has been set.

### GetCapitalActuelSiVariable

`func (o *EntrepriseFiche) GetCapitalActuelSiVariable() string`

GetCapitalActuelSiVariable returns the CapitalActuelSiVariable field if non-nil, zero value otherwise.

### GetCapitalActuelSiVariableOk

`func (o *EntrepriseFiche) GetCapitalActuelSiVariableOk() (*string, bool)`

GetCapitalActuelSiVariableOk returns a tuple with the CapitalActuelSiVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalActuelSiVariable

`func (o *EntrepriseFiche) SetCapitalActuelSiVariable(v string)`

SetCapitalActuelSiVariable sets CapitalActuelSiVariable field to given value.

### HasCapitalActuelSiVariable

`func (o *EntrepriseFiche) HasCapitalActuelSiVariable() bool`

HasCapitalActuelSiVariable returns a boolean if a field has been set.

### GetDeviseCapital

`func (o *EntrepriseFiche) GetDeviseCapital() string`

GetDeviseCapital returns the DeviseCapital field if non-nil, zero value otherwise.

### GetDeviseCapitalOk

`func (o *EntrepriseFiche) GetDeviseCapitalOk() (*string, bool)`

GetDeviseCapitalOk returns a tuple with the DeviseCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviseCapital

`func (o *EntrepriseFiche) SetDeviseCapital(v string)`

SetDeviseCapital sets DeviseCapital field to given value.

### HasDeviseCapital

`func (o *EntrepriseFiche) HasDeviseCapital() bool`

HasDeviseCapital returns a boolean if a field has been set.

### GetNumeroRcs

`func (o *EntrepriseFiche) GetNumeroRcs() string`

GetNumeroRcs returns the NumeroRcs field if non-nil, zero value otherwise.

### GetNumeroRcsOk

`func (o *EntrepriseFiche) GetNumeroRcsOk() (*string, bool)`

GetNumeroRcsOk returns a tuple with the NumeroRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroRcs

`func (o *EntrepriseFiche) SetNumeroRcs(v string)`

SetNumeroRcs sets NumeroRcs field to given value.

### HasNumeroRcs

`func (o *EntrepriseFiche) HasNumeroRcs() bool`

HasNumeroRcs returns a boolean if a field has been set.

### GetDateClotureExercice

`func (o *EntrepriseFiche) GetDateClotureExercice() string`

GetDateClotureExercice returns the DateClotureExercice field if non-nil, zero value otherwise.

### GetDateClotureExerciceOk

`func (o *EntrepriseFiche) GetDateClotureExerciceOk() (*string, bool)`

GetDateClotureExerciceOk returns a tuple with the DateClotureExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClotureExercice

`func (o *EntrepriseFiche) SetDateClotureExercice(v string)`

SetDateClotureExercice sets DateClotureExercice field to given value.

### HasDateClotureExercice

`func (o *EntrepriseFiche) HasDateClotureExercice() bool`

HasDateClotureExercice returns a boolean if a field has been set.

### GetDateClotureExerciceExceptionnelle

`func (o *EntrepriseFiche) GetDateClotureExerciceExceptionnelle() string`

GetDateClotureExerciceExceptionnelle returns the DateClotureExerciceExceptionnelle field if non-nil, zero value otherwise.

### GetDateClotureExerciceExceptionnelleOk

`func (o *EntrepriseFiche) GetDateClotureExerciceExceptionnelleOk() (*string, bool)`

GetDateClotureExerciceExceptionnelleOk returns a tuple with the DateClotureExerciceExceptionnelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClotureExerciceExceptionnelle

`func (o *EntrepriseFiche) SetDateClotureExerciceExceptionnelle(v string)`

SetDateClotureExerciceExceptionnelle sets DateClotureExerciceExceptionnelle field to given value.

### HasDateClotureExerciceExceptionnelle

`func (o *EntrepriseFiche) HasDateClotureExerciceExceptionnelle() bool`

HasDateClotureExerciceExceptionnelle returns a boolean if a field has been set.

### GetDateClotureExerciceExceptionnelleFormate

`func (o *EntrepriseFiche) GetDateClotureExerciceExceptionnelleFormate() string`

GetDateClotureExerciceExceptionnelleFormate returns the DateClotureExerciceExceptionnelleFormate field if non-nil, zero value otherwise.

### GetDateClotureExerciceExceptionnelleFormateOk

`func (o *EntrepriseFiche) GetDateClotureExerciceExceptionnelleFormateOk() (*string, bool)`

GetDateClotureExerciceExceptionnelleFormateOk returns a tuple with the DateClotureExerciceExceptionnelleFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClotureExerciceExceptionnelleFormate

`func (o *EntrepriseFiche) SetDateClotureExerciceExceptionnelleFormate(v string)`

SetDateClotureExerciceExceptionnelleFormate sets DateClotureExerciceExceptionnelleFormate field to given value.

### HasDateClotureExerciceExceptionnelleFormate

`func (o *EntrepriseFiche) HasDateClotureExerciceExceptionnelleFormate() bool`

HasDateClotureExerciceExceptionnelleFormate returns a boolean if a field has been set.

### GetProchaineDateClotureExercice

`func (o *EntrepriseFiche) GetProchaineDateClotureExercice() string`

GetProchaineDateClotureExercice returns the ProchaineDateClotureExercice field if non-nil, zero value otherwise.

### GetProchaineDateClotureExerciceOk

`func (o *EntrepriseFiche) GetProchaineDateClotureExerciceOk() (*string, bool)`

GetProchaineDateClotureExerciceOk returns a tuple with the ProchaineDateClotureExercice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProchaineDateClotureExercice

`func (o *EntrepriseFiche) SetProchaineDateClotureExercice(v string)`

SetProchaineDateClotureExercice sets ProchaineDateClotureExercice field to given value.

### HasProchaineDateClotureExercice

`func (o *EntrepriseFiche) HasProchaineDateClotureExercice() bool`

HasProchaineDateClotureExercice returns a boolean if a field has been set.

### GetProchaineDateClotureExerciceFormate

`func (o *EntrepriseFiche) GetProchaineDateClotureExerciceFormate() string`

GetProchaineDateClotureExerciceFormate returns the ProchaineDateClotureExerciceFormate field if non-nil, zero value otherwise.

### GetProchaineDateClotureExerciceFormateOk

`func (o *EntrepriseFiche) GetProchaineDateClotureExerciceFormateOk() (*string, bool)`

GetProchaineDateClotureExerciceFormateOk returns a tuple with the ProchaineDateClotureExerciceFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProchaineDateClotureExerciceFormate

`func (o *EntrepriseFiche) SetProchaineDateClotureExerciceFormate(v string)`

SetProchaineDateClotureExerciceFormate sets ProchaineDateClotureExerciceFormate field to given value.

### HasProchaineDateClotureExerciceFormate

`func (o *EntrepriseFiche) HasProchaineDateClotureExerciceFormate() bool`

HasProchaineDateClotureExerciceFormate returns a boolean if a field has been set.

### GetEconomieSocialeSolidaire

`func (o *EntrepriseFiche) GetEconomieSocialeSolidaire() bool`

GetEconomieSocialeSolidaire returns the EconomieSocialeSolidaire field if non-nil, zero value otherwise.

### GetEconomieSocialeSolidaireOk

`func (o *EntrepriseFiche) GetEconomieSocialeSolidaireOk() (*bool, bool)`

GetEconomieSocialeSolidaireOk returns a tuple with the EconomieSocialeSolidaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomieSocialeSolidaire

`func (o *EntrepriseFiche) SetEconomieSocialeSolidaire(v bool)`

SetEconomieSocialeSolidaire sets EconomieSocialeSolidaire field to given value.

### HasEconomieSocialeSolidaire

`func (o *EntrepriseFiche) HasEconomieSocialeSolidaire() bool`

HasEconomieSocialeSolidaire returns a boolean if a field has been set.

### GetDureePersonneMorale

`func (o *EntrepriseFiche) GetDureePersonneMorale() string`

GetDureePersonneMorale returns the DureePersonneMorale field if non-nil, zero value otherwise.

### GetDureePersonneMoraleOk

`func (o *EntrepriseFiche) GetDureePersonneMoraleOk() (*string, bool)`

GetDureePersonneMoraleOk returns a tuple with the DureePersonneMorale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDureePersonneMorale

`func (o *EntrepriseFiche) SetDureePersonneMorale(v string)`

SetDureePersonneMorale sets DureePersonneMorale field to given value.

### HasDureePersonneMorale

`func (o *EntrepriseFiche) HasDureePersonneMorale() bool`

HasDureePersonneMorale returns a boolean if a field has been set.

### GetDernierTraitement

`func (o *EntrepriseFiche) GetDernierTraitement() string`

GetDernierTraitement returns the DernierTraitement field if non-nil, zero value otherwise.

### GetDernierTraitementOk

`func (o *EntrepriseFiche) GetDernierTraitementOk() (*string, bool)`

GetDernierTraitementOk returns a tuple with the DernierTraitement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDernierTraitement

`func (o *EntrepriseFiche) SetDernierTraitement(v string)`

SetDernierTraitement sets DernierTraitement field to given value.

### HasDernierTraitement

`func (o *EntrepriseFiche) HasDernierTraitement() bool`

HasDernierTraitement returns a boolean if a field has been set.

### GetDerniereMiseAJourSirene

`func (o *EntrepriseFiche) GetDerniereMiseAJourSirene() string`

GetDerniereMiseAJourSirene returns the DerniereMiseAJourSirene field if non-nil, zero value otherwise.

### GetDerniereMiseAJourSireneOk

`func (o *EntrepriseFiche) GetDerniereMiseAJourSireneOk() (*string, bool)`

GetDerniereMiseAJourSireneOk returns a tuple with the DerniereMiseAJourSirene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerniereMiseAJourSirene

`func (o *EntrepriseFiche) SetDerniereMiseAJourSirene(v string)`

SetDerniereMiseAJourSirene sets DerniereMiseAJourSirene field to given value.

### HasDerniereMiseAJourSirene

`func (o *EntrepriseFiche) HasDerniereMiseAJourSirene() bool`

HasDerniereMiseAJourSirene returns a boolean if a field has been set.

### GetDerniereMiseAJourRcs

`func (o *EntrepriseFiche) GetDerniereMiseAJourRcs() string`

GetDerniereMiseAJourRcs returns the DerniereMiseAJourRcs field if non-nil, zero value otherwise.

### GetDerniereMiseAJourRcsOk

`func (o *EntrepriseFiche) GetDerniereMiseAJourRcsOk() (*string, bool)`

GetDerniereMiseAJourRcsOk returns a tuple with the DerniereMiseAJourRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerniereMiseAJourRcs

`func (o *EntrepriseFiche) SetDerniereMiseAJourRcs(v string)`

SetDerniereMiseAJourRcs sets DerniereMiseAJourRcs field to given value.

### HasDerniereMiseAJourRcs

`func (o *EntrepriseFiche) HasDerniereMiseAJourRcs() bool`

HasDerniereMiseAJourRcs returns a boolean if a field has been set.

### GetGreffe

`func (o *EntrepriseFiche) GetGreffe() string`

GetGreffe returns the Greffe field if non-nil, zero value otherwise.

### GetGreffeOk

`func (o *EntrepriseFiche) GetGreffeOk() (*string, bool)`

GetGreffeOk returns a tuple with the Greffe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreffe

`func (o *EntrepriseFiche) SetGreffe(v string)`

SetGreffe sets Greffe field to given value.

### HasGreffe

`func (o *EntrepriseFiche) HasGreffe() bool`

HasGreffe returns a boolean if a field has been set.

### GetCodeGreffe

`func (o *EntrepriseFiche) GetCodeGreffe() string`

GetCodeGreffe returns the CodeGreffe field if non-nil, zero value otherwise.

### GetCodeGreffeOk

`func (o *EntrepriseFiche) GetCodeGreffeOk() (*string, bool)`

GetCodeGreffeOk returns a tuple with the CodeGreffe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeGreffe

`func (o *EntrepriseFiche) SetCodeGreffe(v string)`

SetCodeGreffe sets CodeGreffe field to given value.

### HasCodeGreffe

`func (o *EntrepriseFiche) HasCodeGreffe() bool`

HasCodeGreffe returns a boolean if a field has been set.

### GetDateImmatriculationRcs

`func (o *EntrepriseFiche) GetDateImmatriculationRcs() string`

GetDateImmatriculationRcs returns the DateImmatriculationRcs field if non-nil, zero value otherwise.

### GetDateImmatriculationRcsOk

`func (o *EntrepriseFiche) GetDateImmatriculationRcsOk() (*string, bool)`

GetDateImmatriculationRcsOk returns a tuple with the DateImmatriculationRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateImmatriculationRcs

`func (o *EntrepriseFiche) SetDateImmatriculationRcs(v string)`

SetDateImmatriculationRcs sets DateImmatriculationRcs field to given value.

### HasDateImmatriculationRcs

`func (o *EntrepriseFiche) HasDateImmatriculationRcs() bool`

HasDateImmatriculationRcs returns a boolean if a field has been set.

### GetDatePremiereImmatriculationRcs

`func (o *EntrepriseFiche) GetDatePremiereImmatriculationRcs() string`

GetDatePremiereImmatriculationRcs returns the DatePremiereImmatriculationRcs field if non-nil, zero value otherwise.

### GetDatePremiereImmatriculationRcsOk

`func (o *EntrepriseFiche) GetDatePremiereImmatriculationRcsOk() (*string, bool)`

GetDatePremiereImmatriculationRcsOk returns a tuple with the DatePremiereImmatriculationRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePremiereImmatriculationRcs

`func (o *EntrepriseFiche) SetDatePremiereImmatriculationRcs(v string)`

SetDatePremiereImmatriculationRcs sets DatePremiereImmatriculationRcs field to given value.

### HasDatePremiereImmatriculationRcs

`func (o *EntrepriseFiche) HasDatePremiereImmatriculationRcs() bool`

HasDatePremiereImmatriculationRcs returns a boolean if a field has been set.

### GetDateDebutActivite

`func (o *EntrepriseFiche) GetDateDebutActivite() string`

GetDateDebutActivite returns the DateDebutActivite field if non-nil, zero value otherwise.

### GetDateDebutActiviteOk

`func (o *EntrepriseFiche) GetDateDebutActiviteOk() (*string, bool)`

GetDateDebutActiviteOk returns a tuple with the DateDebutActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebutActivite

`func (o *EntrepriseFiche) SetDateDebutActivite(v string)`

SetDateDebutActivite sets DateDebutActivite field to given value.

### HasDateDebutActivite

`func (o *EntrepriseFiche) HasDateDebutActivite() bool`

HasDateDebutActivite returns a boolean if a field has been set.

### GetDateDebutPremiereActivite

`func (o *EntrepriseFiche) GetDateDebutPremiereActivite() string`

GetDateDebutPremiereActivite returns the DateDebutPremiereActivite field if non-nil, zero value otherwise.

### GetDateDebutPremiereActiviteOk

`func (o *EntrepriseFiche) GetDateDebutPremiereActiviteOk() (*string, bool)`

GetDateDebutPremiereActiviteOk returns a tuple with the DateDebutPremiereActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebutPremiereActivite

`func (o *EntrepriseFiche) SetDateDebutPremiereActivite(v string)`

SetDateDebutPremiereActivite sets DateDebutPremiereActivite field to given value.

### HasDateDebutPremiereActivite

`func (o *EntrepriseFiche) HasDateDebutPremiereActivite() bool`

HasDateDebutPremiereActivite returns a boolean if a field has been set.

### GetDateRadiationRcs

`func (o *EntrepriseFiche) GetDateRadiationRcs() string`

GetDateRadiationRcs returns the DateRadiationRcs field if non-nil, zero value otherwise.

### GetDateRadiationRcsOk

`func (o *EntrepriseFiche) GetDateRadiationRcsOk() (*string, bool)`

GetDateRadiationRcsOk returns a tuple with the DateRadiationRcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRadiationRcs

`func (o *EntrepriseFiche) SetDateRadiationRcs(v string)`

SetDateRadiationRcs sets DateRadiationRcs field to given value.

### HasDateRadiationRcs

`func (o *EntrepriseFiche) HasDateRadiationRcs() bool`

HasDateRadiationRcs returns a boolean if a field has been set.

### GetStatutRne

`func (o *EntrepriseFiche) GetStatutRne() string`

GetStatutRne returns the StatutRne field if non-nil, zero value otherwise.

### GetStatutRneOk

`func (o *EntrepriseFiche) GetStatutRneOk() (*string, bool)`

GetStatutRneOk returns a tuple with the StatutRne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatutRne

`func (o *EntrepriseFiche) SetStatutRne(v string)`

SetStatutRne sets StatutRne field to given value.

### HasStatutRne

`func (o *EntrepriseFiche) HasStatutRne() bool`

HasStatutRne returns a boolean if a field has been set.

### GetDateImmatriculationRne

`func (o *EntrepriseFiche) GetDateImmatriculationRne() string`

GetDateImmatriculationRne returns the DateImmatriculationRne field if non-nil, zero value otherwise.

### GetDateImmatriculationRneOk

`func (o *EntrepriseFiche) GetDateImmatriculationRneOk() (*string, bool)`

GetDateImmatriculationRneOk returns a tuple with the DateImmatriculationRne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateImmatriculationRne

`func (o *EntrepriseFiche) SetDateImmatriculationRne(v string)`

SetDateImmatriculationRne sets DateImmatriculationRne field to given value.

### HasDateImmatriculationRne

`func (o *EntrepriseFiche) HasDateImmatriculationRne() bool`

HasDateImmatriculationRne returns a boolean if a field has been set.

### GetDateRadiationRne

`func (o *EntrepriseFiche) GetDateRadiationRne() string`

GetDateRadiationRne returns the DateRadiationRne field if non-nil, zero value otherwise.

### GetDateRadiationRneOk

`func (o *EntrepriseFiche) GetDateRadiationRneOk() (*string, bool)`

GetDateRadiationRneOk returns a tuple with the DateRadiationRne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRadiationRne

`func (o *EntrepriseFiche) SetDateRadiationRne(v string)`

SetDateRadiationRne sets DateRadiationRne field to given value.

### HasDateRadiationRne

`func (o *EntrepriseFiche) HasDateRadiationRne() bool`

HasDateRadiationRne returns a boolean if a field has been set.

### GetNumeroTvaIntracommunautaire

`func (o *EntrepriseFiche) GetNumeroTvaIntracommunautaire() string`

GetNumeroTvaIntracommunautaire returns the NumeroTvaIntracommunautaire field if non-nil, zero value otherwise.

### GetNumeroTvaIntracommunautaireOk

`func (o *EntrepriseFiche) GetNumeroTvaIntracommunautaireOk() (*string, bool)`

GetNumeroTvaIntracommunautaireOk returns a tuple with the NumeroTvaIntracommunautaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroTvaIntracommunautaire

`func (o *EntrepriseFiche) SetNumeroTvaIntracommunautaire(v string)`

SetNumeroTvaIntracommunautaire sets NumeroTvaIntracommunautaire field to given value.

### HasNumeroTvaIntracommunautaire

`func (o *EntrepriseFiche) HasNumeroTvaIntracommunautaire() bool`

HasNumeroTvaIntracommunautaire returns a boolean if a field has been set.

### GetValiditeTvaIntracommunautaire

`func (o *EntrepriseFiche) GetValiditeTvaIntracommunautaire() bool`

GetValiditeTvaIntracommunautaire returns the ValiditeTvaIntracommunautaire field if non-nil, zero value otherwise.

### GetValiditeTvaIntracommunautaireOk

`func (o *EntrepriseFiche) GetValiditeTvaIntracommunautaireOk() (*bool, bool)`

GetValiditeTvaIntracommunautaireOk returns a tuple with the ValiditeTvaIntracommunautaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValiditeTvaIntracommunautaire

`func (o *EntrepriseFiche) SetValiditeTvaIntracommunautaire(v bool)`

SetValiditeTvaIntracommunautaire sets ValiditeTvaIntracommunautaire field to given value.

### HasValiditeTvaIntracommunautaire

`func (o *EntrepriseFiche) HasValiditeTvaIntracommunautaire() bool`

HasValiditeTvaIntracommunautaire returns a boolean if a field has been set.

### GetAssocieUnique

`func (o *EntrepriseFiche) GetAssocieUnique() bool`

GetAssocieUnique returns the AssocieUnique field if non-nil, zero value otherwise.

### GetAssocieUniqueOk

`func (o *EntrepriseFiche) GetAssocieUniqueOk() (*bool, bool)`

GetAssocieUniqueOk returns a tuple with the AssocieUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocieUnique

`func (o *EntrepriseFiche) SetAssocieUnique(v bool)`

SetAssocieUnique sets AssocieUnique field to given value.

### HasAssocieUnique

`func (o *EntrepriseFiche) HasAssocieUnique() bool`

HasAssocieUnique returns a boolean if a field has been set.

### GetEtablissements

`func (o *EntrepriseFiche) GetEtablissements() []EtablissementFiche`

GetEtablissements returns the Etablissements field if non-nil, zero value otherwise.

### GetEtablissementsOk

`func (o *EntrepriseFiche) GetEtablissementsOk() (*[]EtablissementFiche, bool)`

GetEtablissementsOk returns a tuple with the Etablissements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissements

`func (o *EntrepriseFiche) SetEtablissements(v []EtablissementFiche)`

SetEtablissements sets Etablissements field to given value.

### HasEtablissements

`func (o *EntrepriseFiche) HasEtablissements() bool`

HasEtablissements returns a boolean if a field has been set.

### GetEtablissement

`func (o *EntrepriseFiche) GetEtablissement() EntrepriseFicheAllOfEtablissement`

GetEtablissement returns the Etablissement field if non-nil, zero value otherwise.

### GetEtablissementOk

`func (o *EntrepriseFiche) GetEtablissementOk() (*EntrepriseFicheAllOfEtablissement, bool)`

GetEtablissementOk returns a tuple with the Etablissement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissement

`func (o *EntrepriseFiche) SetEtablissement(v EntrepriseFicheAllOfEtablissement)`

SetEtablissement sets Etablissement field to given value.

### HasEtablissement

`func (o *EntrepriseFiche) HasEtablissement() bool`

HasEtablissement returns a boolean if a field has been set.

### GetFinances

`func (o *EntrepriseFiche) GetFinances() []EntrepriseFicheAllOfFinances`

GetFinances returns the Finances field if non-nil, zero value otherwise.

### GetFinancesOk

`func (o *EntrepriseFiche) GetFinancesOk() (*[]EntrepriseFicheAllOfFinances, bool)`

GetFinancesOk returns a tuple with the Finances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinances

`func (o *EntrepriseFiche) SetFinances(v []EntrepriseFicheAllOfFinances)`

SetFinances sets Finances field to given value.

### HasFinances

`func (o *EntrepriseFiche) HasFinances() bool`

HasFinances returns a boolean if a field has been set.

### GetRepresentants

`func (o *EntrepriseFiche) GetRepresentants() []RepresentantEntreprise`

GetRepresentants returns the Representants field if non-nil, zero value otherwise.

### GetRepresentantsOk

`func (o *EntrepriseFiche) GetRepresentantsOk() (*[]RepresentantEntreprise, bool)`

GetRepresentantsOk returns a tuple with the Representants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentants

`func (o *EntrepriseFiche) SetRepresentants(v []RepresentantEntreprise)`

SetRepresentants sets Representants field to given value.

### HasRepresentants

`func (o *EntrepriseFiche) HasRepresentants() bool`

HasRepresentants returns a boolean if a field has been set.

### GetBeneficiairesEffectifs

`func (o *EntrepriseFiche) GetBeneficiairesEffectifs() []EntrepriseFicheAllOfBeneficiairesEffectifs`

GetBeneficiairesEffectifs returns the BeneficiairesEffectifs field if non-nil, zero value otherwise.

### GetBeneficiairesEffectifsOk

`func (o *EntrepriseFiche) GetBeneficiairesEffectifsOk() (*[]EntrepriseFicheAllOfBeneficiairesEffectifs, bool)`

GetBeneficiairesEffectifsOk returns a tuple with the BeneficiairesEffectifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiairesEffectifs

`func (o *EntrepriseFiche) SetBeneficiairesEffectifs(v []EntrepriseFicheAllOfBeneficiairesEffectifs)`

SetBeneficiairesEffectifs sets BeneficiairesEffectifs field to given value.

### HasBeneficiairesEffectifs

`func (o *EntrepriseFiche) HasBeneficiairesEffectifs() bool`

HasBeneficiairesEffectifs returns a boolean if a field has been set.

### GetDepotsActes

`func (o *EntrepriseFiche) GetDepotsActes() []EntrepriseFicheAllOfDepotsActes`

GetDepotsActes returns the DepotsActes field if non-nil, zero value otherwise.

### GetDepotsActesOk

`func (o *EntrepriseFiche) GetDepotsActesOk() (*[]EntrepriseFicheAllOfDepotsActes, bool)`

GetDepotsActesOk returns a tuple with the DepotsActes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotsActes

`func (o *EntrepriseFiche) SetDepotsActes(v []EntrepriseFicheAllOfDepotsActes)`

SetDepotsActes sets DepotsActes field to given value.

### HasDepotsActes

`func (o *EntrepriseFiche) HasDepotsActes() bool`

HasDepotsActes returns a boolean if a field has been set.

### GetComptes

`func (o *EntrepriseFiche) GetComptes() []EntrepriseFicheAllOfComptes`

GetComptes returns the Comptes field if non-nil, zero value otherwise.

### GetComptesOk

`func (o *EntrepriseFiche) GetComptesOk() (*[]EntrepriseFicheAllOfComptes, bool)`

GetComptesOk returns a tuple with the Comptes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComptes

`func (o *EntrepriseFiche) SetComptes(v []EntrepriseFicheAllOfComptes)`

SetComptes sets Comptes field to given value.

### HasComptes

`func (o *EntrepriseFiche) HasComptes() bool`

HasComptes returns a boolean if a field has been set.

### GetPublicationsBodacc

`func (o *EntrepriseFiche) GetPublicationsBodacc() []Bodacc`

GetPublicationsBodacc returns the PublicationsBodacc field if non-nil, zero value otherwise.

### GetPublicationsBodaccOk

`func (o *EntrepriseFiche) GetPublicationsBodaccOk() (*[]Bodacc, bool)`

GetPublicationsBodaccOk returns a tuple with the PublicationsBodacc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationsBodacc

`func (o *EntrepriseFiche) SetPublicationsBodacc(v []Bodacc)`

SetPublicationsBodacc sets PublicationsBodacc field to given value.

### HasPublicationsBodacc

`func (o *EntrepriseFiche) HasPublicationsBodacc() bool`

HasPublicationsBodacc returns a boolean if a field has been set.

### GetProceduresCollectives

`func (o *EntrepriseFiche) GetProceduresCollectives() []EntrepriseFicheAllOfProceduresCollectives`

GetProceduresCollectives returns the ProceduresCollectives field if non-nil, zero value otherwise.

### GetProceduresCollectivesOk

`func (o *EntrepriseFiche) GetProceduresCollectivesOk() (*[]EntrepriseFicheAllOfProceduresCollectives, bool)`

GetProceduresCollectivesOk returns a tuple with the ProceduresCollectives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceduresCollectives

`func (o *EntrepriseFiche) SetProceduresCollectives(v []EntrepriseFicheAllOfProceduresCollectives)`

SetProceduresCollectives sets ProceduresCollectives field to given value.

### HasProceduresCollectives

`func (o *EntrepriseFiche) HasProceduresCollectives() bool`

HasProceduresCollectives returns a boolean if a field has been set.

### GetProcedureCollectiveExiste

`func (o *EntrepriseFiche) GetProcedureCollectiveExiste() bool`

GetProcedureCollectiveExiste returns the ProcedureCollectiveExiste field if non-nil, zero value otherwise.

### GetProcedureCollectiveExisteOk

`func (o *EntrepriseFiche) GetProcedureCollectiveExisteOk() (*bool, bool)`

GetProcedureCollectiveExisteOk returns a tuple with the ProcedureCollectiveExiste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedureCollectiveExiste

`func (o *EntrepriseFiche) SetProcedureCollectiveExiste(v bool)`

SetProcedureCollectiveExiste sets ProcedureCollectiveExiste field to given value.

### HasProcedureCollectiveExiste

`func (o *EntrepriseFiche) HasProcedureCollectiveExiste() bool`

HasProcedureCollectiveExiste returns a boolean if a field has been set.

### GetProcedureCollectiveEnCours

`func (o *EntrepriseFiche) GetProcedureCollectiveEnCours() bool`

GetProcedureCollectiveEnCours returns the ProcedureCollectiveEnCours field if non-nil, zero value otherwise.

### GetProcedureCollectiveEnCoursOk

`func (o *EntrepriseFiche) GetProcedureCollectiveEnCoursOk() (*bool, bool)`

GetProcedureCollectiveEnCoursOk returns a tuple with the ProcedureCollectiveEnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedureCollectiveEnCours

`func (o *EntrepriseFiche) SetProcedureCollectiveEnCours(v bool)`

SetProcedureCollectiveEnCours sets ProcedureCollectiveEnCours field to given value.

### HasProcedureCollectiveEnCours

`func (o *EntrepriseFiche) HasProcedureCollectiveEnCours() bool`

HasProcedureCollectiveEnCours returns a boolean if a field has been set.

### GetDerniersStatuts

`func (o *EntrepriseFiche) GetDerniersStatuts() EntrepriseFicheAllOfDerniersStatuts`

GetDerniersStatuts returns the DerniersStatuts field if non-nil, zero value otherwise.

### GetDerniersStatutsOk

`func (o *EntrepriseFiche) GetDerniersStatutsOk() (*EntrepriseFicheAllOfDerniersStatuts, bool)`

GetDerniersStatutsOk returns a tuple with the DerniersStatuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerniersStatuts

`func (o *EntrepriseFiche) SetDerniersStatuts(v EntrepriseFicheAllOfDerniersStatuts)`

SetDerniersStatuts sets DerniersStatuts field to given value.

### HasDerniersStatuts

`func (o *EntrepriseFiche) HasDerniersStatuts() bool`

HasDerniersStatuts returns a boolean if a field has been set.

### GetExtraitImmatriculation

`func (o *EntrepriseFiche) GetExtraitImmatriculation() EntrepriseFicheAllOfExtraitImmatriculation`

GetExtraitImmatriculation returns the ExtraitImmatriculation field if non-nil, zero value otherwise.

### GetExtraitImmatriculationOk

`func (o *EntrepriseFiche) GetExtraitImmatriculationOk() (*EntrepriseFicheAllOfExtraitImmatriculation, bool)`

GetExtraitImmatriculationOk returns a tuple with the ExtraitImmatriculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraitImmatriculation

`func (o *EntrepriseFiche) SetExtraitImmatriculation(v EntrepriseFicheAllOfExtraitImmatriculation)`

SetExtraitImmatriculation sets ExtraitImmatriculation field to given value.

### HasExtraitImmatriculation

`func (o *EntrepriseFiche) HasExtraitImmatriculation() bool`

HasExtraitImmatriculation returns a boolean if a field has been set.

### GetRnm

`func (o *EntrepriseFiche) GetRnm() EntrepriseFicheAllOfRnm`

GetRnm returns the Rnm field if non-nil, zero value otherwise.

### GetRnmOk

`func (o *EntrepriseFiche) GetRnmOk() (*EntrepriseFicheAllOfRnm, bool)`

GetRnmOk returns a tuple with the Rnm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnm

`func (o *EntrepriseFiche) SetRnm(v EntrepriseFicheAllOfRnm)`

SetRnm sets Rnm field to given value.

### HasRnm

`func (o *EntrepriseFiche) HasRnm() bool`

HasRnm returns a boolean if a field has been set.

### GetMarques

`func (o *EntrepriseFiche) GetMarques() []EntrepriseFicheAllOfMarques`

GetMarques returns the Marques field if non-nil, zero value otherwise.

### GetMarquesOk

`func (o *EntrepriseFiche) GetMarquesOk() (*[]EntrepriseFicheAllOfMarques, bool)`

GetMarquesOk returns a tuple with the Marques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarques

`func (o *EntrepriseFiche) SetMarques(v []EntrepriseFicheAllOfMarques)`

SetMarques sets Marques field to given value.

### HasMarques

`func (o *EntrepriseFiche) HasMarques() bool`

HasMarques returns a boolean if a field has been set.

### GetAssociation

`func (o *EntrepriseFiche) GetAssociation() Association`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *EntrepriseFiche) GetAssociationOk() (*Association, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *EntrepriseFiche) SetAssociation(v Association)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *EntrepriseFiche) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetLabels

`func (o *EntrepriseFiche) GetLabels() []Labels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EntrepriseFiche) GetLabelsOk() (*[]Labels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EntrepriseFiche) SetLabels(v []Labels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EntrepriseFiche) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSitesInternet

`func (o *EntrepriseFiche) GetSitesInternet() []string`

GetSitesInternet returns the SitesInternet field if non-nil, zero value otherwise.

### GetSitesInternetOk

`func (o *EntrepriseFiche) GetSitesInternetOk() (*[]string, bool)`

GetSitesInternetOk returns a tuple with the SitesInternet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitesInternet

`func (o *EntrepriseFiche) SetSitesInternet(v []string)`

SetSitesInternet sets SitesInternet field to given value.

### HasSitesInternet

`func (o *EntrepriseFiche) HasSitesInternet() bool`

HasSitesInternet returns a boolean if a field has been set.

### GetTelephone

`func (o *EntrepriseFiche) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *EntrepriseFiche) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *EntrepriseFiche) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *EntrepriseFiche) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetEmail

`func (o *EntrepriseFiche) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EntrepriseFiche) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EntrepriseFiche) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EntrepriseFiche) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetScoringNonFinancier

`func (o *EntrepriseFiche) GetScoringNonFinancier() ScoringNonFinancier`

GetScoringNonFinancier returns the ScoringNonFinancier field if non-nil, zero value otherwise.

### GetScoringNonFinancierOk

`func (o *EntrepriseFiche) GetScoringNonFinancierOk() (*ScoringNonFinancier, bool)`

GetScoringNonFinancierOk returns a tuple with the ScoringNonFinancier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringNonFinancier

`func (o *EntrepriseFiche) SetScoringNonFinancier(v ScoringNonFinancier)`

SetScoringNonFinancier sets ScoringNonFinancier field to given value.

### HasScoringNonFinancier

`func (o *EntrepriseFiche) HasScoringNonFinancier() bool`

HasScoringNonFinancier returns a boolean if a field has been set.

### GetScoringFinancier

`func (o *EntrepriseFiche) GetScoringFinancier() ScoringFinancier`

GetScoringFinancier returns the ScoringFinancier field if non-nil, zero value otherwise.

### GetScoringFinancierOk

`func (o *EntrepriseFiche) GetScoringFinancierOk() (*ScoringFinancier, bool)`

GetScoringFinancierOk returns a tuple with the ScoringFinancier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoringFinancier

`func (o *EntrepriseFiche) SetScoringFinancier(v ScoringFinancier)`

SetScoringFinancier sets ScoringFinancier field to given value.

### HasScoringFinancier

`func (o *EntrepriseFiche) HasScoringFinancier() bool`

HasScoringFinancier returns a boolean if a field has been set.

### GetCategorieEntreprise

`func (o *EntrepriseFiche) GetCategorieEntreprise() string`

GetCategorieEntreprise returns the CategorieEntreprise field if non-nil, zero value otherwise.

### GetCategorieEntrepriseOk

`func (o *EntrepriseFiche) GetCategorieEntrepriseOk() (*string, bool)`

GetCategorieEntrepriseOk returns a tuple with the CategorieEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieEntreprise

`func (o *EntrepriseFiche) SetCategorieEntreprise(v string)`

SetCategorieEntreprise sets CategorieEntreprise field to given value.

### HasCategorieEntreprise

`func (o *EntrepriseFiche) HasCategorieEntreprise() bool`

HasCategorieEntreprise returns a boolean if a field has been set.

### GetAnneeCategorieEntreprise

`func (o *EntrepriseFiche) GetAnneeCategorieEntreprise() int32`

GetAnneeCategorieEntreprise returns the AnneeCategorieEntreprise field if non-nil, zero value otherwise.

### GetAnneeCategorieEntrepriseOk

`func (o *EntrepriseFiche) GetAnneeCategorieEntrepriseOk() (*int32, bool)`

GetAnneeCategorieEntrepriseOk returns a tuple with the AnneeCategorieEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeCategorieEntreprise

`func (o *EntrepriseFiche) SetAnneeCategorieEntreprise(v int32)`

SetAnneeCategorieEntreprise sets AnneeCategorieEntreprise field to given value.

### HasAnneeCategorieEntreprise

`func (o *EntrepriseFiche) HasAnneeCategorieEntreprise() bool`

HasAnneeCategorieEntreprise returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


