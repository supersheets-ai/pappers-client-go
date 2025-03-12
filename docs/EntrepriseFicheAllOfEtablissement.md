# EntrepriseFicheAllOfEtablissement

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

### NewEntrepriseFicheAllOfEtablissement

`func NewEntrepriseFicheAllOfEtablissement() *EntrepriseFicheAllOfEtablissement`

NewEntrepriseFicheAllOfEtablissement instantiates a new EntrepriseFicheAllOfEtablissement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfEtablissementWithDefaults

`func NewEntrepriseFicheAllOfEtablissementWithDefaults() *EntrepriseFicheAllOfEtablissement`

NewEntrepriseFicheAllOfEtablissementWithDefaults instantiates a new EntrepriseFicheAllOfEtablissement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *EntrepriseFicheAllOfEtablissement) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *EntrepriseFicheAllOfEtablissement) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *EntrepriseFicheAllOfEtablissement) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *EntrepriseFicheAllOfEtablissement) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetSiretFormate

`func (o *EntrepriseFicheAllOfEtablissement) GetSiretFormate() string`

GetSiretFormate returns the SiretFormate field if non-nil, zero value otherwise.

### GetSiretFormateOk

`func (o *EntrepriseFicheAllOfEtablissement) GetSiretFormateOk() (*string, bool)`

GetSiretFormateOk returns a tuple with the SiretFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiretFormate

`func (o *EntrepriseFicheAllOfEtablissement) SetSiretFormate(v string)`

SetSiretFormate sets SiretFormate field to given value.

### HasSiretFormate

`func (o *EntrepriseFicheAllOfEtablissement) HasSiretFormate() bool`

HasSiretFormate returns a boolean if a field has been set.

### GetDiffusionPartielle

`func (o *EntrepriseFicheAllOfEtablissement) GetDiffusionPartielle() bool`

GetDiffusionPartielle returns the DiffusionPartielle field if non-nil, zero value otherwise.

### GetDiffusionPartielleOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDiffusionPartielleOk() (*bool, bool)`

GetDiffusionPartielleOk returns a tuple with the DiffusionPartielle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffusionPartielle

`func (o *EntrepriseFicheAllOfEtablissement) SetDiffusionPartielle(v bool)`

SetDiffusionPartielle sets DiffusionPartielle field to given value.

### HasDiffusionPartielle

`func (o *EntrepriseFicheAllOfEtablissement) HasDiffusionPartielle() bool`

HasDiffusionPartielle returns a boolean if a field has been set.

### GetNic

`func (o *EntrepriseFicheAllOfEtablissement) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *EntrepriseFicheAllOfEtablissement) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *EntrepriseFicheAllOfEtablissement) SetNic(v string)`

SetNic sets Nic field to given value.

### HasNic

`func (o *EntrepriseFicheAllOfEtablissement) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetCodePostal

`func (o *EntrepriseFicheAllOfEtablissement) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *EntrepriseFicheAllOfEtablissement) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *EntrepriseFicheAllOfEtablissement) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *EntrepriseFicheAllOfEtablissement) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *EntrepriseFicheAllOfEtablissement) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *EntrepriseFicheAllOfEtablissement) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *EntrepriseFicheAllOfEtablissement) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetPays

`func (o *EntrepriseFicheAllOfEtablissement) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *EntrepriseFicheAllOfEtablissement) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *EntrepriseFicheAllOfEtablissement) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *EntrepriseFicheAllOfEtablissement) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *EntrepriseFicheAllOfEtablissement) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *EntrepriseFicheAllOfEtablissement) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *EntrepriseFicheAllOfEtablissement) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetLatitude

`func (o *EntrepriseFicheAllOfEtablissement) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EntrepriseFicheAllOfEtablissement) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *EntrepriseFicheAllOfEtablissement) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *EntrepriseFicheAllOfEtablissement) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EntrepriseFicheAllOfEtablissement) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *EntrepriseFicheAllOfEtablissement) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetEtablissementCesse

`func (o *EntrepriseFicheAllOfEtablissement) GetEtablissementCesse() bool`

GetEtablissementCesse returns the EtablissementCesse field if non-nil, zero value otherwise.

### GetEtablissementCesseOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEtablissementCesseOk() (*bool, bool)`

GetEtablissementCesseOk returns a tuple with the EtablissementCesse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissementCesse

`func (o *EntrepriseFicheAllOfEtablissement) SetEtablissementCesse(v bool)`

SetEtablissementCesse sets EtablissementCesse field to given value.

### HasEtablissementCesse

`func (o *EntrepriseFicheAllOfEtablissement) HasEtablissementCesse() bool`

HasEtablissementCesse returns a boolean if a field has been set.

### GetSiege

`func (o *EntrepriseFicheAllOfEtablissement) GetSiege() bool`

GetSiege returns the Siege field if non-nil, zero value otherwise.

### GetSiegeOk

`func (o *EntrepriseFicheAllOfEtablissement) GetSiegeOk() (*bool, bool)`

GetSiegeOk returns a tuple with the Siege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiege

`func (o *EntrepriseFicheAllOfEtablissement) SetSiege(v bool)`

SetSiege sets Siege field to given value.

### HasSiege

`func (o *EntrepriseFicheAllOfEtablissement) HasSiege() bool`

HasSiege returns a boolean if a field has been set.

### GetEtablissementEmployeur

`func (o *EntrepriseFicheAllOfEtablissement) GetEtablissementEmployeur() bool`

GetEtablissementEmployeur returns the EtablissementEmployeur field if non-nil, zero value otherwise.

### GetEtablissementEmployeurOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEtablissementEmployeurOk() (*bool, bool)`

GetEtablissementEmployeurOk returns a tuple with the EtablissementEmployeur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtablissementEmployeur

`func (o *EntrepriseFicheAllOfEtablissement) SetEtablissementEmployeur(v bool)`

SetEtablissementEmployeur sets EtablissementEmployeur field to given value.

### HasEtablissementEmployeur

`func (o *EntrepriseFicheAllOfEtablissement) HasEtablissementEmployeur() bool`

HasEtablissementEmployeur returns a boolean if a field has been set.

### GetEffectif

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectif() string`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectifOk() (*string, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *EntrepriseFicheAllOfEtablissement) SetEffectif(v string)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *EntrepriseFicheAllOfEtablissement) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetEffectifMin

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectifMin() int32`

GetEffectifMin returns the EffectifMin field if non-nil, zero value otherwise.

### GetEffectifMinOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectifMinOk() (*int32, bool)`

GetEffectifMinOk returns a tuple with the EffectifMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMin

`func (o *EntrepriseFicheAllOfEtablissement) SetEffectifMin(v int32)`

SetEffectifMin sets EffectifMin field to given value.

### HasEffectifMin

`func (o *EntrepriseFicheAllOfEtablissement) HasEffectifMin() bool`

HasEffectifMin returns a boolean if a field has been set.

### GetEffectifMax

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectifMax() int32`

GetEffectifMax returns the EffectifMax field if non-nil, zero value otherwise.

### GetEffectifMaxOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEffectifMaxOk() (*int32, bool)`

GetEffectifMaxOk returns a tuple with the EffectifMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectifMax

`func (o *EntrepriseFicheAllOfEtablissement) SetEffectifMax(v int32)`

SetEffectifMax sets EffectifMax field to given value.

### HasEffectifMax

`func (o *EntrepriseFicheAllOfEtablissement) HasEffectifMax() bool`

HasEffectifMax returns a boolean if a field has been set.

### GetTrancheEffectif

`func (o *EntrepriseFicheAllOfEtablissement) GetTrancheEffectif() string`

GetTrancheEffectif returns the TrancheEffectif field if non-nil, zero value otherwise.

### GetTrancheEffectifOk

`func (o *EntrepriseFicheAllOfEtablissement) GetTrancheEffectifOk() (*string, bool)`

GetTrancheEffectifOk returns a tuple with the TrancheEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrancheEffectif

`func (o *EntrepriseFicheAllOfEtablissement) SetTrancheEffectif(v string)`

SetTrancheEffectif sets TrancheEffectif field to given value.

### HasTrancheEffectif

`func (o *EntrepriseFicheAllOfEtablissement) HasTrancheEffectif() bool`

HasTrancheEffectif returns a boolean if a field has been set.

### GetAnneeEffectif

`func (o *EntrepriseFicheAllOfEtablissement) GetAnneeEffectif() int32`

GetAnneeEffectif returns the AnneeEffectif field if non-nil, zero value otherwise.

### GetAnneeEffectifOk

`func (o *EntrepriseFicheAllOfEtablissement) GetAnneeEffectifOk() (*int32, bool)`

GetAnneeEffectifOk returns a tuple with the AnneeEffectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnneeEffectif

`func (o *EntrepriseFicheAllOfEtablissement) SetAnneeEffectif(v int32)`

SetAnneeEffectif sets AnneeEffectif field to given value.

### HasAnneeEffectif

`func (o *EntrepriseFicheAllOfEtablissement) HasAnneeEffectif() bool`

HasAnneeEffectif returns a boolean if a field has been set.

### GetCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeNaf() string`

GetCodeNaf returns the CodeNaf field if non-nil, zero value otherwise.

### GetCodeNafOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeNafOk() (*string, bool)`

GetCodeNafOk returns a tuple with the CodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) SetCodeNaf(v string)`

SetCodeNaf sets CodeNaf field to given value.

### HasCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) HasCodeNaf() bool`

HasCodeNaf returns a boolean if a field has been set.

### GetLibelleCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleCodeNaf() string`

GetLibelleCodeNaf returns the LibelleCodeNaf field if non-nil, zero value otherwise.

### GetLibelleCodeNafOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleCodeNafOk() (*string, bool)`

GetLibelleCodeNafOk returns a tuple with the LibelleCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) SetLibelleCodeNaf(v string)`

SetLibelleCodeNaf sets LibelleCodeNaf field to given value.

### HasLibelleCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) HasLibelleCodeNaf() bool`

HasLibelleCodeNaf returns a boolean if a field has been set.

### GetNomenclatureCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) GetNomenclatureCodeNaf() string`

GetNomenclatureCodeNaf returns the NomenclatureCodeNaf field if non-nil, zero value otherwise.

### GetNomenclatureCodeNafOk

`func (o *EntrepriseFicheAllOfEtablissement) GetNomenclatureCodeNafOk() (*string, bool)`

GetNomenclatureCodeNafOk returns a tuple with the NomenclatureCodeNaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatureCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) SetNomenclatureCodeNaf(v string)`

SetNomenclatureCodeNaf sets NomenclatureCodeNaf field to given value.

### HasNomenclatureCodeNaf

`func (o *EntrepriseFicheAllOfEtablissement) HasNomenclatureCodeNaf() bool`

HasNomenclatureCodeNaf returns a boolean if a field has been set.

### GetDateDeCreation

`func (o *EntrepriseFicheAllOfEtablissement) GetDateDeCreation() string`

GetDateDeCreation returns the DateDeCreation field if non-nil, zero value otherwise.

### GetDateDeCreationOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDateDeCreationOk() (*string, bool)`

GetDateDeCreationOk returns a tuple with the DateDeCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeCreation

`func (o *EntrepriseFicheAllOfEtablissement) SetDateDeCreation(v string)`

SetDateDeCreation sets DateDeCreation field to given value.

### HasDateDeCreation

`func (o *EntrepriseFicheAllOfEtablissement) HasDateDeCreation() bool`

HasDateDeCreation returns a boolean if a field has been set.

### GetNumeroVoie

`func (o *EntrepriseFicheAllOfEtablissement) GetNumeroVoie() int32`

GetNumeroVoie returns the NumeroVoie field if non-nil, zero value otherwise.

### GetNumeroVoieOk

`func (o *EntrepriseFicheAllOfEtablissement) GetNumeroVoieOk() (*int32, bool)`

GetNumeroVoieOk returns a tuple with the NumeroVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroVoie

`func (o *EntrepriseFicheAllOfEtablissement) SetNumeroVoie(v int32)`

SetNumeroVoie sets NumeroVoie field to given value.

### HasNumeroVoie

`func (o *EntrepriseFicheAllOfEtablissement) HasNumeroVoie() bool`

HasNumeroVoie returns a boolean if a field has been set.

### GetIndiceRepetition

`func (o *EntrepriseFicheAllOfEtablissement) GetIndiceRepetition() string`

GetIndiceRepetition returns the IndiceRepetition field if non-nil, zero value otherwise.

### GetIndiceRepetitionOk

`func (o *EntrepriseFicheAllOfEtablissement) GetIndiceRepetitionOk() (*string, bool)`

GetIndiceRepetitionOk returns a tuple with the IndiceRepetition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndiceRepetition

`func (o *EntrepriseFicheAllOfEtablissement) SetIndiceRepetition(v string)`

SetIndiceRepetition sets IndiceRepetition field to given value.

### HasIndiceRepetition

`func (o *EntrepriseFicheAllOfEtablissement) HasIndiceRepetition() bool`

HasIndiceRepetition returns a boolean if a field has been set.

### GetTypeVoie

`func (o *EntrepriseFicheAllOfEtablissement) GetTypeVoie() string`

GetTypeVoie returns the TypeVoie field if non-nil, zero value otherwise.

### GetTypeVoieOk

`func (o *EntrepriseFicheAllOfEtablissement) GetTypeVoieOk() (*string, bool)`

GetTypeVoieOk returns a tuple with the TypeVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVoie

`func (o *EntrepriseFicheAllOfEtablissement) SetTypeVoie(v string)`

SetTypeVoie sets TypeVoie field to given value.

### HasTypeVoie

`func (o *EntrepriseFicheAllOfEtablissement) HasTypeVoie() bool`

HasTypeVoie returns a boolean if a field has been set.

### GetLibelleVoie

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleVoie() string`

GetLibelleVoie returns the LibelleVoie field if non-nil, zero value otherwise.

### GetLibelleVoieOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleVoieOk() (*string, bool)`

GetLibelleVoieOk returns a tuple with the LibelleVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleVoie

`func (o *EntrepriseFicheAllOfEtablissement) SetLibelleVoie(v string)`

SetLibelleVoie sets LibelleVoie field to given value.

### HasLibelleVoie

`func (o *EntrepriseFicheAllOfEtablissement) HasLibelleVoie() bool`

HasLibelleVoie returns a boolean if a field has been set.

### GetComplementAdresse

`func (o *EntrepriseFicheAllOfEtablissement) GetComplementAdresse() string`

GetComplementAdresse returns the ComplementAdresse field if non-nil, zero value otherwise.

### GetComplementAdresseOk

`func (o *EntrepriseFicheAllOfEtablissement) GetComplementAdresseOk() (*string, bool)`

GetComplementAdresseOk returns a tuple with the ComplementAdresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementAdresse

`func (o *EntrepriseFicheAllOfEtablissement) SetComplementAdresse(v string)`

SetComplementAdresse sets ComplementAdresse field to given value.

### HasComplementAdresse

`func (o *EntrepriseFicheAllOfEtablissement) HasComplementAdresse() bool`

HasComplementAdresse returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *EntrepriseFicheAllOfEtablissement) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *EntrepriseFicheAllOfEtablissement) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *EntrepriseFicheAllOfEtablissement) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *EntrepriseFicheAllOfEtablissement) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *EntrepriseFicheAllOfEtablissement) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *EntrepriseFicheAllOfEtablissement) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *EntrepriseFicheAllOfEtablissement) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *EntrepriseFicheAllOfEtablissement) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetDateCessation

`func (o *EntrepriseFicheAllOfEtablissement) GetDateCessation() string`

GetDateCessation returns the DateCessation field if non-nil, zero value otherwise.

### GetDateCessationOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDateCessationOk() (*string, bool)`

GetDateCessationOk returns a tuple with the DateCessation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessation

`func (o *EntrepriseFicheAllOfEtablissement) SetDateCessation(v string)`

SetDateCessation sets DateCessation field to given value.

### HasDateCessation

`func (o *EntrepriseFicheAllOfEtablissement) HasDateCessation() bool`

HasDateCessation returns a boolean if a field has been set.

### GetEnseigne

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne() string`

GetEnseigne returns the Enseigne field if non-nil, zero value otherwise.

### GetEnseigneOk

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigneOk() (*string, bool)`

GetEnseigneOk returns a tuple with the Enseigne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne

`func (o *EntrepriseFicheAllOfEtablissement) SetEnseigne(v string)`

SetEnseigne sets Enseigne field to given value.

### HasEnseigne

`func (o *EntrepriseFicheAllOfEtablissement) HasEnseigne() bool`

HasEnseigne returns a boolean if a field has been set.

### GetNomCommercial

`func (o *EntrepriseFicheAllOfEtablissement) GetNomCommercial() string`

GetNomCommercial returns the NomCommercial field if non-nil, zero value otherwise.

### GetNomCommercialOk

`func (o *EntrepriseFicheAllOfEtablissement) GetNomCommercialOk() (*string, bool)`

GetNomCommercialOk returns a tuple with the NomCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomCommercial

`func (o *EntrepriseFicheAllOfEtablissement) SetNomCommercial(v string)`

SetNomCommercial sets NomCommercial field to given value.

### HasNomCommercial

`func (o *EntrepriseFicheAllOfEtablissement) HasNomCommercial() bool`

HasNomCommercial returns a boolean if a field has been set.

### GetDomiciliation

`func (o *EntrepriseFicheAllOfEtablissement) GetDomiciliation() EtablissementFicheDomiciliation`

GetDomiciliation returns the Domiciliation field if non-nil, zero value otherwise.

### GetDomiciliationOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDomiciliationOk() (*EtablissementFicheDomiciliation, bool)`

GetDomiciliationOk returns a tuple with the Domiciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomiciliation

`func (o *EntrepriseFicheAllOfEtablissement) SetDomiciliation(v EtablissementFicheDomiciliation)`

SetDomiciliation sets Domiciliation field to given value.

### HasDomiciliation

`func (o *EntrepriseFicheAllOfEtablissement) HasDomiciliation() bool`

HasDomiciliation returns a boolean if a field has been set.

### GetLabels

`func (o *EntrepriseFicheAllOfEtablissement) GetLabels() []LabelsBase`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLabelsOk() (*[]LabelsBase, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EntrepriseFicheAllOfEtablissement) SetLabels(v []LabelsBase)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EntrepriseFicheAllOfEtablissement) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPredecesseurs

`func (o *EntrepriseFicheAllOfEtablissement) GetPredecesseurs() []LienSuccession`

GetPredecesseurs returns the Predecesseurs field if non-nil, zero value otherwise.

### GetPredecesseursOk

`func (o *EntrepriseFicheAllOfEtablissement) GetPredecesseursOk() (*[]LienSuccession, bool)`

GetPredecesseursOk returns a tuple with the Predecesseurs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecesseurs

`func (o *EntrepriseFicheAllOfEtablissement) SetPredecesseurs(v []LienSuccession)`

SetPredecesseurs sets Predecesseurs field to given value.

### HasPredecesseurs

`func (o *EntrepriseFicheAllOfEtablissement) HasPredecesseurs() bool`

HasPredecesseurs returns a boolean if a field has been set.

### GetSuccesseurs

`func (o *EntrepriseFicheAllOfEtablissement) GetSuccesseurs() []LienSuccession`

GetSuccesseurs returns the Successeurs field if non-nil, zero value otherwise.

### GetSuccesseursOk

`func (o *EntrepriseFicheAllOfEtablissement) GetSuccesseursOk() (*[]LienSuccession, bool)`

GetSuccesseursOk returns a tuple with the Successeurs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccesseurs

`func (o *EntrepriseFicheAllOfEtablissement) SetSuccesseurs(v []LienSuccession)`

SetSuccesseurs sets Successeurs field to given value.

### HasSuccesseurs

`func (o *EntrepriseFicheAllOfEtablissement) HasSuccesseurs() bool`

HasSuccesseurs returns a boolean if a field has been set.

### GetEnseigne1

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne1() string`

GetEnseigne1 returns the Enseigne1 field if non-nil, zero value otherwise.

### GetEnseigne1Ok

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne1Ok() (*string, bool)`

GetEnseigne1Ok returns a tuple with the Enseigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne1

`func (o *EntrepriseFicheAllOfEtablissement) SetEnseigne1(v string)`

SetEnseigne1 sets Enseigne1 field to given value.

### HasEnseigne1

`func (o *EntrepriseFicheAllOfEtablissement) HasEnseigne1() bool`

HasEnseigne1 returns a boolean if a field has been set.

### GetEnseigne2

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne2() string`

GetEnseigne2 returns the Enseigne2 field if non-nil, zero value otherwise.

### GetEnseigne2Ok

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne2Ok() (*string, bool)`

GetEnseigne2Ok returns a tuple with the Enseigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne2

`func (o *EntrepriseFicheAllOfEtablissement) SetEnseigne2(v string)`

SetEnseigne2 sets Enseigne2 field to given value.

### HasEnseigne2

`func (o *EntrepriseFicheAllOfEtablissement) HasEnseigne2() bool`

HasEnseigne2 returns a boolean if a field has been set.

### GetEnseigne3

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne3() string`

GetEnseigne3 returns the Enseigne3 field if non-nil, zero value otherwise.

### GetEnseigne3Ok

`func (o *EntrepriseFicheAllOfEtablissement) GetEnseigne3Ok() (*string, bool)`

GetEnseigne3Ok returns a tuple with the Enseigne3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnseigne3

`func (o *EntrepriseFicheAllOfEtablissement) SetEnseigne3(v string)`

SetEnseigne3 sets Enseigne3 field to given value.

### HasEnseigne3

`func (o *EntrepriseFicheAllOfEtablissement) HasEnseigne3() bool`

HasEnseigne3 returns a boolean if a field has been set.

### GetDistributionSpeciale

`func (o *EntrepriseFicheAllOfEtablissement) GetDistributionSpeciale() string`

GetDistributionSpeciale returns the DistributionSpeciale field if non-nil, zero value otherwise.

### GetDistributionSpecialeOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDistributionSpecialeOk() (*string, bool)`

GetDistributionSpecialeOk returns a tuple with the DistributionSpeciale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionSpeciale

`func (o *EntrepriseFicheAllOfEtablissement) SetDistributionSpeciale(v string)`

SetDistributionSpeciale sets DistributionSpeciale field to given value.

### HasDistributionSpeciale

`func (o *EntrepriseFicheAllOfEtablissement) HasDistributionSpeciale() bool`

HasDistributionSpeciale returns a boolean if a field has been set.

### GetCodeCedex

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeCedex() string`

GetCodeCedex returns the CodeCedex field if non-nil, zero value otherwise.

### GetCodeCedexOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeCedexOk() (*string, bool)`

GetCodeCedexOk returns a tuple with the CodeCedex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCedex

`func (o *EntrepriseFicheAllOfEtablissement) SetCodeCedex(v string)`

SetCodeCedex sets CodeCedex field to given value.

### HasCodeCedex

`func (o *EntrepriseFicheAllOfEtablissement) HasCodeCedex() bool`

HasCodeCedex returns a boolean if a field has been set.

### GetLibelleCedex

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleCedex() string`

GetLibelleCedex returns the LibelleCedex field if non-nil, zero value otherwise.

### GetLibelleCedexOk

`func (o *EntrepriseFicheAllOfEtablissement) GetLibelleCedexOk() (*string, bool)`

GetLibelleCedexOk returns a tuple with the LibelleCedex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleCedex

`func (o *EntrepriseFicheAllOfEtablissement) SetLibelleCedex(v string)`

SetLibelleCedex sets LibelleCedex field to given value.

### HasLibelleCedex

`func (o *EntrepriseFicheAllOfEtablissement) HasLibelleCedex() bool`

HasLibelleCedex returns a boolean if a field has been set.

### GetCodeCommune

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeCommune() string`

GetCodeCommune returns the CodeCommune field if non-nil, zero value otherwise.

### GetCodeCommuneOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeCommuneOk() (*string, bool)`

GetCodeCommuneOk returns a tuple with the CodeCommune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeCommune

`func (o *EntrepriseFicheAllOfEtablissement) SetCodeCommune(v string)`

SetCodeCommune sets CodeCommune field to given value.

### HasCodeCommune

`func (o *EntrepriseFicheAllOfEtablissement) HasCodeCommune() bool`

HasCodeCommune returns a boolean if a field has been set.

### GetDepartement

`func (o *EntrepriseFicheAllOfEtablissement) GetDepartement() string`

GetDepartement returns the Departement field if non-nil, zero value otherwise.

### GetDepartementOk

`func (o *EntrepriseFicheAllOfEtablissement) GetDepartementOk() (*string, bool)`

GetDepartementOk returns a tuple with the Departement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartement

`func (o *EntrepriseFicheAllOfEtablissement) SetDepartement(v string)`

SetDepartement sets Departement field to given value.

### HasDepartement

`func (o *EntrepriseFicheAllOfEtablissement) HasDepartement() bool`

HasDepartement returns a boolean if a field has been set.

### GetCodeDepartement

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeDepartement() string`

GetCodeDepartement returns the CodeDepartement field if non-nil, zero value otherwise.

### GetCodeDepartementOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeDepartementOk() (*string, bool)`

GetCodeDepartementOk returns a tuple with the CodeDepartement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeDepartement

`func (o *EntrepriseFicheAllOfEtablissement) SetCodeDepartement(v string)`

SetCodeDepartement sets CodeDepartement field to given value.

### HasCodeDepartement

`func (o *EntrepriseFicheAllOfEtablissement) HasCodeDepartement() bool`

HasCodeDepartement returns a boolean if a field has been set.

### GetRegion

`func (o *EntrepriseFicheAllOfEtablissement) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EntrepriseFicheAllOfEtablissement) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EntrepriseFicheAllOfEtablissement) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EntrepriseFicheAllOfEtablissement) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCodeRegion

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeRegion() string`

GetCodeRegion returns the CodeRegion field if non-nil, zero value otherwise.

### GetCodeRegionOk

`func (o *EntrepriseFicheAllOfEtablissement) GetCodeRegionOk() (*string, bool)`

GetCodeRegionOk returns a tuple with the CodeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeRegion

`func (o *EntrepriseFicheAllOfEtablissement) SetCodeRegion(v string)`

SetCodeRegion sets CodeRegion field to given value.

### HasCodeRegion

`func (o *EntrepriseFicheAllOfEtablissement) HasCodeRegion() bool`

HasCodeRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


