# EtablissementFiche

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siret** | Pointer to **string** | Numéro siret de l&#39;établissement au format xxxxxxxxxxxxxx. | [optional] 
**SiretFormate** | Pointer to **string** | Numéro siret de l&#39;établissement au format xxx xxx xxx xxxxx. | [optional] 
**DiffusionPartielle** | Pointer to **bool** | Si vrai, l&#39;établissement est en diffusion partielle. Dans ce cas, tous les champs relatifs à son adresse - en dehors de la ville et du pays - sont à &#x60;null&#x60;. | [optional] 
**Nic** | Pointer to **string** | Numéro NIC de l&#39;établissement. | [optional] 
**CodePostal** | Pointer to **string** | Code postal de l&#39;établissement. | [optional] 
**Ville** | Pointer to **string** | Ville de l&#39;établissement. | [optional] 
**Pays** | Pointer to **string** | Pays de l&#39;établissement | [optional] 
**CodePays** | Pointer to **string** | Code du pays de l&#39;établissement | [optional] 
**Latitude** | Pointer to **float32** | Latitude de l&#39;établissement. | [optional] 
**Longitude** | Pointer to **float32** | Longitude de l&#39;établissement. | [optional] 
**EtablissementCesse** | Pointer to **bool** | Vrai si l&#39;établissement est cessé, faux si il est en activité. | [optional] 
**Siege** | Pointer to **bool** | Vrai si l&#39;établissement est siège, faux s&#39;il ne l&#39;est pas. | [optional] 
**EtablissementEmployeur** | Pointer to **bool** | Si vrai, l&#39;établissement a au moins un employé. | [optional] 
**Effectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;établissement. | [optional] 
**EffectifMin** | Pointer to **int32** | Effectif minimal de l&#39;établissement. | [optional] 
**EffectifMax** | Pointer to **int32** | Effectif maximal de l&#39;établissement. | [optional] 
**TrancheEffectif** | Pointer to **string** | Tranche d&#39;effectif de l&#39;établissement, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73). | [optional] 
**AnneeEffectif** | Pointer to **int32** | Année correspondante à la tranche d&#39;effectif de l&#39;établissement. | [optional] 
**CodeNaf** | Pointer to **string** | Code NAF de l&#39;établissement. | [optional] 
**LibelleCodeNaf** | Pointer to **string** | Libellé du code NAF de l&#39;établissement. | [optional] 
**NomenclatureCodeNaf** | Pointer to **string** | Nomenclature code NAF de l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**DateDeCreation** | Pointer to **string** |  | [optional] 
**NumeroVoie** | Pointer to **int32** | Numéro de voie de l&#39;établissement. | [optional] 
**IndiceRepetition** | Pointer to **string** | Indice de répétition de l&#39;établissement. | [optional] 
**TypeVoie** | Pointer to **string** | Type de voie de l&#39;établissement. | [optional] 
**LibelleVoie** | Pointer to **string** | Libellé de la voie de l&#39;établissement. | [optional] 
**ComplementAdresse** | Pointer to **string** | Complément d&#39;adresse de l&#39;établissement. | [optional] 
**AdresseLigne1** | Pointer to **string** | Première ligne de l&#39;adresse de l&#39;établissement. Correspond à l&#39;ensemble des données numero_voie, indice_repetition, type_voie et libelle_voie. | [optional] 
**AdresseLigne2** | Pointer to **string** | Seconde ligne de l&#39;adresse de l&#39;établissement. Est égal à complement_adresse | [optional] 
**DateCessation** | Pointer to **string** | Date de fermeture de l&#39;établissement | [optional] 
**Enseigne** | Pointer to **string** | Enseigne de l&#39;établissement | [optional] 
**NomCommercial** | Pointer to **string** | Nom commercial de l&#39;établissement | [optional] 
**Domiciliation** | Pointer to [**EtablissementFicheDomiciliation**](EtablissementFicheDomiciliation.md) |  | [optional] 
**Labels** | Pointer to [**[]LabelsBase**](LabelsBase.md) | Liste des labels de l&#39;entreprise. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Predecesseurs** | Pointer to [**[]LienSuccession**](LienSuccession.md) | Liste des prédécesseurs de l&#39;établissement. | [optional] 
**Successeurs** | Pointer to [**[]LienSuccession**](LienSuccession.md) | Liste des successeurs de l&#39;établissement. | [optional] 
**Enseigne1** | Pointer to **string** | Enseigne 1 de l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**Enseigne2** | Pointer to **string** | Enseigne 2 de l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**Enseigne3** | Pointer to **string** | Enseigne 3 de l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**DistributionSpeciale** | Pointer to **string** | Distribution spéciale de l&#39;établissement. La distribution spéciale reprend les éléments particuliers qui accompagnent une adresse de distribution spéciale. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**CodeCedex** | Pointer to **string** | Code cedex de l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**LibelleCedex** | Pointer to **string** | Ce champ indique le libellé correspondant au code cedex de l&#39;établissement. Si le code cedex est à &#39;null&#39;, ce champ est également à &#39;null&#39;. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**CodeCommune** | Pointer to **string** | Le code commune désigne le code de la commune de localisation de l&#39;établissement. Cette valeur est à &#39;null&#39; pour les entreprises à l&#39;étranger. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**Departement** | Pointer to **string** | Libellé du département. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**CodeDepartement** | Pointer to **string** | Code du département. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**Region** | Pointer to **string** | Région dans laquelle est situé l&#39;établissement. Uniquement présent si demandé dans les champs supplémentaires | [optional] 
**CodeRegion** | Pointer to **string** | Code région. Uniquement présent si demandé dans les champs supplémentaires | [optional] 

## Methods

### NewEtablissementFiche

`func NewEtablissementFiche() *EtablissementFiche`

NewEtablissementFiche instantiates a new EtablissementFiche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtablissementFicheWithDefaults

`func NewEtablissementFicheWithDefaults() *EtablissementFiche`

NewEtablissementFicheWithDefaults instantiates a new EtablissementFiche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *EtablissementFiche) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *EtablissementFiche) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *EtablissementFiche) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *EtablissementFiche) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetSiretFormate

`func (o *EtablissementFiche) GetSiretFormate() string`

GetSiretFormate returns the SiretFormate field if non-nil, zero value otherwise.

### GetSiretFormateOk

`func (o *EtablissementFiche) GetSiretFormateOk() (*string, bool)`

GetSiretFormateOk returns a tuple with the SiretFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiretFormate

`func (o *EtablissementFiche) SetSiretFormate(v string)`

SetSiretFormate sets SiretFormate field to given value.

### HasSiretFormate

`func (o *EtablissementFiche) HasSiretFormate() bool`

HasSiretFormate returns a boolean if a field has been set.

### GetDiffusionPartielle

`func (o *EtablissementFiche) GetDiffusionPartielle() bool`

GetDiffusionPartielle returns the DiffusionPartielle field if non-nil, zero value otherwise.

### GetDiffusionPartielleOk

`func (o *EtablissementFiche) GetDiffusionPartielleOk() (*bool, bool)`

GetDiffusionPartielleOk returns a tuple with the DiffusionPartielle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffusionPartielle

`func (o *EtablissementFiche) SetDiffusionPartielle(v bool)`

SetDiffusionPartielle sets DiffusionPartielle field to given value.

### HasDiffusionPartielle

`func (o *EtablissementFiche) HasDiffusionPartielle() bool`

HasDiffusionPartielle returns a boolean if a field has been set.

### GetNic

`func (o *EtablissementFiche) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *EtablissementFiche) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *EtablissementFiche) SetNic(v string)`

SetNic sets Nic field to given value.

### HasNic

`func (o *EtablissementFiche) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetCodePostal

`func (o *EtablissementFiche) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *EtablissementFiche) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *EtablissementFiche) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *EtablissementFiche) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *EtablissementFiche) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *EtablissementFiche) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *EtablissementFiche) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *EtablissementFiche) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *EtablissementFiche) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *EtablissementFiche) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *EtablissementFiche) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *EtablissementFiche) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *EtablissementFiche) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *EtablissementFiche) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *EtablissementFiche) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *EtablissementFiche) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetLatitude

`func (o *EtablissementFiche) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EtablissementFiche) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EtablissementFiche) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *EtablissementFiche) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *EtablissementFiche) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EtablissementFiche) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EtablissementFiche) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *EtablissementFiche) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetEtablissementCesse

`func (o *EtablissementFiche) GetEtablissementCesse() bool`

GetEtablissementCesse returns the EtablissementCesse field if non-nil, zero value otherwise.

### GetEtablissementCesseOk

`func (o *EtablissementFiche) GetEtablissementCesseOk() (*bool, bool)`

GetEtablissementCesseOk returns a tuple with the EtablissementCesse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissementCesse

`func (o *EtablissementFiche) SetEtablissementCesse(v bool)`

SetEtablissementCesse sets EtablissementCesse field to given value.

### HasEtablissementCesse

`func (o *EtablissementFiche) HasEtablissementCesse() bool`

HasEtablissementCesse returns a boolean if a field has been set.

### GetSiege

`func (o *EtablissementFiche) GetSiege() bool`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *EtablissementFiche) GetSiegeOk() (*bool, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *EtablissementFiche) SetSiege(v bool)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *EtablissementFiche) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetEtablissementEmployeur

`func (o *EtablissementFiche) GetEtablissementEmployeur() bool`

GetEtablissementEmployeur returns the EtablissementEmployeur field if non-nil, zero value otherwise.

### GetEtablissementEmployeurOk

`func (o *EtablissementFiche) GetEtablissementEmployeurOk() (*bool, bool)`

GetEtablissementEmployeurOk returns a tuple with the EtablissementEmployeur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissementEmployeur

`func (o *EtablissementFiche) SetEtablissementEmployeur(v bool)`

SetEtablissementEmployeur sets EtablissementEmployeur field to given value.

### HasEtablissementEmployeur

`func (o *EtablissementFiche) HasEtablissementEmployeur() bool`

HasEtablissementEmployeur returns a boolean if a field has been set.

### GetEffectif

`func (o *EtablissementFiche) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *EtablissementFiche) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *EtablissementFiche) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *EtablissementFiche) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *EtablissementFiche) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *EtablissementFiche) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *EtablissementFiche) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *EtablissementFiche) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *EtablissementFiche) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *EtablissementFiche) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *EtablissementFiche) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *EtablissementFiche) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *EtablissementFiche) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *EtablissementFiche) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *EtablissementFiche) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *EtablissementFiche) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *EtablissementFiche) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *EtablissementFiche) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *EtablissementFiche) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *EtablissementFiche) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCodeNaf

`func (o *EtablissementFiche) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *EtablissementFiche) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *EtablissementFiche) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *EtablissementFiche) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *EtablissementFiche) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *EtablissementFiche) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *EtablissementFiche) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *EtablissementFiche) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetNomenclatureCodeNaf

`func (o *EtablissementFiche) GetNomenclatureCodeNaf() string`

GetNomenclatureCodeNaf returns the NomenclatureCodeNaf field if non-nil, zero value otherwise.

### GetNomenclatureCodeNafOk

`func (o *EtablissementFiche) GetNomenclatureCodeNafOk() (*string, bool)`

GetNomenclatureCodeNafOk returns a tuple with the NomenclatureCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatureCodeNaf

`func (o *EtablissementFiche) SetNomenclatureCodeNaf(v string)`

SetNomenclatureCodeNaf sets NomenclatureCodeNaf field to given value.

### HasNomenclatureCodeNaf

`func (o *EtablissementFiche) HasNomenclatureCodeNaf() bool`

HasNomenclatureCodeNaf returns a boolean if a field has been set.

### GetDateDeCreation

`func (o *EtablissementFiche) GetDateDeCreation() string`

GetDateDeCreation returns the DateDeCreation field if non-nil, zero value otherwise.

### GetDateDeCreationOk

`func (o *EtablissementFiche) GetDateDeCreationOk() (*string, bool)`

GetDateDeCreationOk returns a tuple with the DateDeCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeCreation

`func (o *EtablissementFiche) SetDateDeCreation(v string)`

SetDateDeCreation sets DateDeCreation field to given value.

### HasDateDeCreation

`func (o *EtablissementFiche) HasDateDeCreation() bool`

HasDateDeCreation returns a boolean if a field has been set.

### GetNumeroVoie

`func (o *EtablissementFiche) GetNumeroVoie() int32`

GetNumeroVoie returns the NumeroVoie field if non-nil, zero value otherwise.

### GetNumeroVoieOk

`func (o *EtablissementFiche) GetNumeroVoieOk() (*int32, bool)`

GetNumeroVoieOk returns a tuple with the NumeroVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroVoie

`func (o *EtablissementFiche) SetNumeroVoie(v int32)`

SetNumeroVoie sets NumeroVoie field to given value.

### HasNumeroVoie

`func (o *EtablissementFiche) HasNumeroVoie() bool`

HasNumeroVoie returns a boolean if a field has been set.

### GetIndiceRepetition

`func (o *EtablissementFiche) GetIndiceRepetition() string`

GetIndiceRepetition returns the IndiceRepetition field if non-nil, zero value otherwise.

### GetIndiceRepetitionOk

`func (o *EtablissementFiche) GetIndiceRepetitionOk() (*string, bool)`

GetIndiceRepetitionOk returns a tuple with the IndiceRepetition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndiceRepetition

`func (o *EtablissementFiche) SetIndiceRepetition(v string)`

SetIndiceRepetition sets IndiceRepetition field to given value.

### HasIndiceRepetition

`func (o *EtablissementFiche) HasIndiceRepetition() bool`

HasIndiceRepetition returns a boolean if a field has been set.

### GetTypeVoie

`func (o *EtablissementFiche) GetTypeVoie() string`

GetTypeVoie returns the TypeVoie field if non-nil, zero value otherwise.

### GetTypeVoieOk

`func (o *EtablissementFiche) GetTypeVoieOk() (*string, bool)`

GetTypeVoieOk returns a tuple with the TypeVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVoie

`func (o *EtablissementFiche) SetTypeVoie(v string)`

SetTypeVoie sets TypeVoie field to given value.

### HasTypeVoie

`func (o *EtablissementFiche) HasTypeVoie() bool`

HasTypeVoie returns a boolean if a field has been set.

### GetLibelleVoie

`func (o *EtablissementFiche) GetLibelleVoie() string`

GetLibelleVoie returns the LibelleVoie field if non-nil, zero value otherwise.

### GetLibelleVoieOk

`func (o *EtablissementFiche) GetLibelleVoieOk() (*string, bool)`

GetLibelleVoieOk returns a tuple with the LibelleVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleVoie

`func (o *EtablissementFiche) SetLibelleVoie(v string)`

SetLibelleVoie sets LibelleVoie field to given value.

### HasLibelleVoie

`func (o *EtablissementFiche) HasLibelleVoie() bool`

HasLibelleVoie returns a boolean if a field has been set.

### GetComplementAdresse

`func (o *EtablissementFiche) GetComplementAdresse() string`

GetComplementAdresse returns the ComplementAdresse field if non-nil, zero value otherwise.

### GetComplementAdresseOk

`func (o *EtablissementFiche) GetComplementAdresseOk() (*string, bool)`

GetComplementAdresseOk returns a tuple with the ComplementAdresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementAdresse

`func (o *EtablissementFiche) SetComplementAdresse(v string)`

SetComplementAdresse sets ComplementAdresse field to given value.

### HasComplementAdresse

`func (o *EtablissementFiche) HasComplementAdresse() bool`

HasComplementAdresse returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *EtablissementFiche) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *EtablissementFiche) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *EtablissementFiche) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *EtablissementFiche) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *EtablissementFiche) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *EtablissementFiche) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *EtablissementFiche) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *EtablissementFiche) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetDateCessation

`func (o *EtablissementFiche) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *EtablissementFiche) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *EtablissementFiche) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *EtablissementFiche) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEnseigne

`func (o *EtablissementFiche) GetEnseigne() string`

GetEnseigne returns the Enseigne field if non-nil, zero value otherwise.

### GetEnseigneOk

`func (o *EtablissementFiche) GetEnseigneOk() (*string, bool)`

GetEnseigneOk returns a tuple with the Enseigne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne

`func (o *EtablissementFiche) SetEnseigne(v string)`

SetEnseigne sets Enseigne field to given value.

### HasEnseigne

`func (o *EtablissementFiche) HasEnseigne() bool`

HasEnseigne returns a boolean if a field has been set.

### GetNomCommercial

`func (o *EtablissementFiche) GetNomCommercial() string`

GetNomCommercial returns the NomCommercial field if non-nil, zero value otherwise.

### GetNomCommercialOk

`func (o *EtablissementFiche) GetNomCommercialOk() (*string, bool)`

GetNomCommercialOk returns a tuple with the NomCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomCommercial

`func (o *EtablissementFiche) SetNomCommercial(v string)`

SetNomCommercial sets NomCommercial field to given value.

### HasNomCommercial

`func (o *EtablissementFiche) HasNomCommercial() bool`

HasNomCommercial returns a boolean if a field has been set.

### GetDomiciliation

`func (o *EtablissementFiche) GetDomiciliation() EtablissementFicheDomiciliation`

GetDomiciliation returns the Domiciliation field if non-nil, zero value otherwise.

### GetDomiciliationOk

`func (o *EtablissementFiche) GetDomiciliationOk() (*EtablissementFicheDomiciliation, bool)`

GetDomiciliationOk returns a tuple with the Domiciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomiciliation

`func (o *EtablissementFiche) SetDomiciliation(v EtablissementFicheDomiciliation)`

SetDomiciliation sets Domiciliation field to given value.

### HasDomiciliation

`func (o *EtablissementFiche) HasDomiciliation() bool`

HasDomiciliation returns a boolean if a field has been set.

### GetLabels

`func (o *EtablissementFiche) GetLabels() []LabelsBase`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EtablissementFiche) GetLabelsOk() (*[]LabelsBase, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EtablissementFiche) SetLabels(v []LabelsBase)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EtablissementFiche) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPredecesseurs

`func (o *EtablissementFiche) GetPredecesseurs() []LienSuccession`

GetPredecesseurs returns the Predecesseurs field if non-nil, zero value otherwise.

### GetPredecesseursOk

`func (o *EtablissementFiche) GetPredecesseursOk() (*[]LienSuccession, bool)`

GetPredecesseursOk returns a tuple with the Predecesseurs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecesseurs

`func (o *EtablissementFiche) SetPredecesseurs(v []LienSuccession)`

SetPredecesseurs sets Predecesseurs field to given value.

### HasPredecesseurs

`func (o *EtablissementFiche) HasPredecesseurs() bool`

HasPredecesseurs returns a boolean if a field has been set.

### GetSuccesseurs

`func (o *EtablissementFiche) GetSuccesseurs() []LienSuccession`

GetSuccesseurs returns the Successeurs field if non-nil, zero value otherwise.

### GetSuccesseursOk

`func (o *EtablissementFiche) GetSuccesseursOk() (*[]LienSuccession, bool)`

GetSuccesseursOk returns a tuple with the Successeurs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesseurs

`func (o *EtablissementFiche) SetSuccesseurs(v []LienSuccession)`

SetSuccesseurs sets Successeurs field to given value.

### HasSuccesseurs

`func (o *EtablissementFiche) HasSuccesseurs() bool`

HasSuccesseurs returns a boolean if a field has been set.

### GetEnseigne1

`func (o *EtablissementFiche) GetEnseigne1() string`

GetEnseigne1 returns the Enseigne1 field if non-nil, zero value otherwise.

### GetEnseigne1Ok

`func (o *EtablissementFiche) GetEnseigne1Ok() (*string, bool)`

GetEnseigne1Ok returns a tuple with the Enseigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne1

`func (o *EtablissementFiche) SetEnseigne1(v string)`

SetEnseigne1 sets Enseigne1 field to given value.

### HasEnseigne1

`func (o *EtablissementFiche) HasEnseigne1() bool`

HasEnseigne1 returns a boolean if a field has been set.

### GetEnseigne2

`func (o *EtablissementFiche) GetEnseigne2() string`

GetEnseigne2 returns the Enseigne2 field if non-nil, zero value otherwise.

### GetEnseigne2Ok

`func (o *EtablissementFiche) GetEnseigne2Ok() (*string, bool)`

GetEnseigne2Ok returns a tuple with the Enseigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne2

`func (o *EtablissementFiche) SetEnseigne2(v string)`

SetEnseigne2 sets Enseigne2 field to given value.

### HasEnseigne2

`func (o *EtablissementFiche) HasEnseigne2() bool`

HasEnseigne2 returns a boolean if a field has been set.

### GetEnseigne3

`func (o *EtablissementFiche) GetEnseigne3() string`

GetEnseigne3 returns the Enseigne3 field if non-nil, zero value otherwise.

### GetEnseigne3Ok

`func (o *EtablissementFiche) GetEnseigne3Ok() (*string, bool)`

GetEnseigne3Ok returns a tuple with the Enseigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne3

`func (o *EtablissementFiche) SetEnseigne3(v string)`

SetEnseigne3 sets Enseigne3 field to given value.

### HasEnseigne3

`func (o *EtablissementFiche) HasEnseigne3() bool`

HasEnseigne3 returns a boolean if a field has been set.

### GetDistributionSpeciale

`func (o *EtablissementFiche) GetDistributionSpeciale() string`

GetDistributionSpeciale returns the DistributionSpeciale field if non-nil, zero value otherwise.

### GetDistributionSpecialeOk

`func (o *EtablissementFiche) GetDistributionSpecialeOk() (*string, bool)`

GetDistributionSpecialeOk returns a tuple with the DistributionSpeciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionSpeciale

`func (o *EtablissementFiche) SetDistributionSpeciale(v string)`

SetDistributionSpeciale sets DistributionSpeciale field to given value.

### HasDistributionSpeciale

`func (o *EtablissementFiche) HasDistributionSpeciale() bool`

HasDistributionSpeciale returns a boolean if a field has been set.

### GetCodeCedex

`func (o *EtablissementFiche) GetCodeCedex() string`

GetCodeCedex returns the CodeCedex field if non-nil, zero value otherwise.

### GetCodeCedexOk

`func (o *EtablissementFiche) GetCodeCedexOk() (*string, bool)`

GetCodeCedexOk returns a tuple with the CodeCedex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCedex

`func (o *EtablissementFiche) SetCodeCedex(v string)`

SetCodeCedex sets CodeCedex field to given value.

### HasCodeCedex

`func (o *EtablissementFiche) HasCodeCedex() bool`

HasCodeCedex returns a boolean if a field has been set.

### GetLibelleCedex

`func (o *EtablissementFiche) GetLibelleCedex() string`

GetLibelleCedex returns the LibelleCedex field if non-nil, zero value otherwise.

### GetLibelleCedexOk

`func (o *EtablissementFiche) GetLibelleCedexOk() (*string, bool)`

GetLibelleCedexOk returns a tuple with the LibelleCedex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCedex

`func (o *EtablissementFiche) SetLibelleCedex(v string)`

SetLibelleCedex sets LibelleCedex field to given value.

### HasLibelleCedex

`func (o *EtablissementFiche) HasLibelleCedex() bool`

HasLibelleCedex returns a boolean if a field has been set.

### GetCodeCommune

`func (o *EtablissementFiche) GetCodeCommune() string`

GetCodeCommune returns the CodeCommune field if non-nil, zero value otherwise.

### GetCodeCommuneOk

`func (o *EtablissementFiche) GetCodeCommuneOk() (*string, bool)`

GetCodeCommuneOk returns a tuple with the CodeCommune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCommune

`func (o *EtablissementFiche) SetCodeCommune(v string)`

SetCodeCommune sets CodeCommune field to given value.

### HasCodeCommune

`func (o *EtablissementFiche) HasCodeCommune() bool`

HasCodeCommune returns a boolean if a field has been set.

### GetDepartement

`func (o *EtablissementFiche) GetDepartement() string`

GetDepartement returns the Departement field if non-nil, zero value otherwise.

### GetDepartementOk

`func (o *EtablissementFiche) GetDepartementOk() (*string, bool)`

GetDepartementOk returns a tuple with the Departement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartement

`func (o *EtablissementFiche) SetDepartement(v string)`

SetDepartement sets Departement field to given value.

### HasDepartement

`func (o *EtablissementFiche) HasDepartement() bool`

HasDepartement returns a boolean if a field has been set.

### GetCodeDepartement

`func (o *EtablissementFiche) GetCodeDepartement() string`

GetCodeDepartement returns the CodeDepartement field if non-nil, zero value otherwise.

### GetCodeDepartementOk

`func (o *EtablissementFiche) GetCodeDepartementOk() (*string, bool)`

GetCodeDepartementOk returns a tuple with the CodeDepartement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeDepartement

`func (o *EtablissementFiche) SetCodeDepartement(v string)`

SetCodeDepartement sets CodeDepartement field to given value.

### HasCodeDepartement

`func (o *EtablissementFiche) HasCodeDepartement() bool`

HasCodeDepartement returns a boolean if a field has been set.

### GetRegion

`func (o *EtablissementFiche) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EtablissementFiche) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EtablissementFiche) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EtablissementFiche) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCodeRegion

`func (o *EtablissementFiche) GetCodeRegion() string`

GetCodeRegion returns the CodeRegion field if non-nil, zero value otherwise.

### GetCodeRegionOk

`func (o *EtablissementFiche) GetCodeRegionOk() (*string, bool)`

GetCodeRegionOk returns a tuple with the CodeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRegion

`func (o *EtablissementFiche) SetCodeRegion(v string)`

SetCodeRegion sets CodeRegion field to given value.

### HasCodeRegion

`func (o *EtablissementFiche) HasCodeRegion() bool`

HasCodeRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


