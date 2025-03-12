# ComptesAnnuels200ResponseValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateDepot** | Pointer to **string** | Date de dépôt des comptes, au format AAAA-MM-JJ. | [optional] 
**CodeGreffe** | Pointer to **string** | Code du greffe de dépôt des comptes. | [optional] 
**NumeroDepot** | Pointer to **string** | Numéro de dépôt des comptes au greffe. | [optional] 
**NumeroGestion** | Pointer to **string** | Numéro de gestion au greffe. | [optional] 
**DateCloture** | Pointer to **string** | Date de cloture des comptes, au format AAAA-MM-JJ. | [optional] 
**DateClotureN1** | Pointer to **string** | Date de cloture des comptes n-1, lorsque présents, au format AAAA-MM-JJ. | [optional] 
**DureeExerciceN** | Pointer to **int32** | Durée de l&#39;exercice déposé, en mois. | [optional] 
**DureeExerciceN1** | Pointer to **int32** | Durée de l&#39;exercice précédent, lorsque présent, en mois. | [optional] 
**TypeComptes** | Pointer to **string** | Type de comptes (C &#x3D; complets ; S &#x3D; simplifiés ; K &#x3D; consolidés ; CS &#x3D; mélange complets/simplifiés ; B &#x3D; banques ; A &#x3D; assurances). | [optional] 
**LibelleTypeComptes** | Pointer to **string** | Libellé du type de comptes. | [optional] 
**Devise** | Pointer to **string** | Devise des comptes. | [optional] 
**DeviseOrigine** | Pointer to **string** | Devise d&#39;origine en cas de conversion. | [optional] 
**Confidentialite** | Pointer to **bool** | Confidentialité totale des comptes. | [optional] 
**ConfidentialiteCompteDeResultat** | Pointer to **bool** | Confidentialité partielle des comptes (seul le compte de résultat est confidentiel, le reste des comptes sont disponibles). | [optional] 
**CoherenceComptable** | Pointer to **bool** | Vrai si les comptes sont cohérents d&#39;un point de vue comptable (équilibre du bilan par exemple). | [optional] 
**TypeSaisie** | Pointer to **string** | Description du type de saisie des comptes. | [optional] 
**InformationsTraitement** | Pointer to **[]string** | Informations complémentaires sur le traitement des comptes. | [optional] 
**Sections** | Pointer to [**[]ComptesAnnuels200ResponseValueInnerSectionsInner**](ComptesAnnuels200ResponseValueInnerSectionsInner.md) | Liste des sections de liasses fiscales. | [optional] 
**Ratios** | Pointer to [**Ratios**](Ratios.md) |  | [optional] 

## Methods

### NewComptesAnnuels200ResponseValueInner

`func NewComptesAnnuels200ResponseValueInner() *ComptesAnnuels200ResponseValueInner`

NewComptesAnnuels200ResponseValueInner instantiates a new ComptesAnnuels200ResponseValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComptesAnnuels200ResponseValueInnerWithDefaults

`func NewComptesAnnuels200ResponseValueInnerWithDefaults() *ComptesAnnuels200ResponseValueInner`

NewComptesAnnuels200ResponseValueInnerWithDefaults instantiates a new ComptesAnnuels200ResponseValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateDepot

`func (o *ComptesAnnuels200ResponseValueInner) GetDateDepot() string`

GetDateDepot returns the DateDepot field if non-nil, zero value otherwise.

### GetDateDepotOk

`func (o *ComptesAnnuels200ResponseValueInner) GetDateDepotOk() (*string, bool)`

GetDateDepotOk returns a tuple with the DateDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepot

`func (o *ComptesAnnuels200ResponseValueInner) SetDateDepot(v string)`

SetDateDepot sets DateDepot field to given value.

### HasDateDepot

`func (o *ComptesAnnuels200ResponseValueInner) HasDateDepot() bool`

HasDateDepot returns a boolean if a field has been set.

### GetCodeGreffe

`func (o *ComptesAnnuels200ResponseValueInner) GetCodeGreffe() string`

GetCodeGreffe returns the CodeGreffe field if non-nil, zero value otherwise.

### GetCodeGreffeOk

`func (o *ComptesAnnuels200ResponseValueInner) GetCodeGreffeOk() (*string, bool)`

GetCodeGreffeOk returns a tuple with the CodeGreffe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeGreffe

`func (o *ComptesAnnuels200ResponseValueInner) SetCodeGreffe(v string)`

SetCodeGreffe sets CodeGreffe field to given value.

### HasCodeGreffe

`func (o *ComptesAnnuels200ResponseValueInner) HasCodeGreffe() bool`

HasCodeGreffe returns a boolean if a field has been set.

### GetNumeroDepot

`func (o *ComptesAnnuels200ResponseValueInner) GetNumeroDepot() string`

GetNumeroDepot returns the NumeroDepot field if non-nil, zero value otherwise.

### GetNumeroDepotOk

`func (o *ComptesAnnuels200ResponseValueInner) GetNumeroDepotOk() (*string, bool)`

GetNumeroDepotOk returns a tuple with the NumeroDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroDepot

`func (o *ComptesAnnuels200ResponseValueInner) SetNumeroDepot(v string)`

SetNumeroDepot sets NumeroDepot field to given value.

### HasNumeroDepot

`func (o *ComptesAnnuels200ResponseValueInner) HasNumeroDepot() bool`

HasNumeroDepot returns a boolean if a field has been set.

### GetNumeroGestion

`func (o *ComptesAnnuels200ResponseValueInner) GetNumeroGestion() string`

GetNumeroGestion returns the NumeroGestion field if non-nil, zero value otherwise.

### GetNumeroGestionOk

`func (o *ComptesAnnuels200ResponseValueInner) GetNumeroGestionOk() (*string, bool)`

GetNumeroGestionOk returns a tuple with the NumeroGestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroGestion

`func (o *ComptesAnnuels200ResponseValueInner) SetNumeroGestion(v string)`

SetNumeroGestion sets NumeroGestion field to given value.

### HasNumeroGestion

`func (o *ComptesAnnuels200ResponseValueInner) HasNumeroGestion() bool`

HasNumeroGestion returns a boolean if a field has been set.

### GetDateCloture

`func (o *ComptesAnnuels200ResponseValueInner) GetDateCloture() string`

GetDateCloture returns the DateCloture field if non-nil, zero value otherwise.

### GetDateClotureOk

`func (o *ComptesAnnuels200ResponseValueInner) GetDateClotureOk() (*string, bool)`

GetDateClotureOk returns a tuple with the DateCloture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCloture

`func (o *ComptesAnnuels200ResponseValueInner) SetDateCloture(v string)`

SetDateCloture sets DateCloture field to given value.

### HasDateCloture

`func (o *ComptesAnnuels200ResponseValueInner) HasDateCloture() bool`

HasDateCloture returns a boolean if a field has been set.

### GetDateClotureN1

`func (o *ComptesAnnuels200ResponseValueInner) GetDateClotureN1() string`

GetDateClotureN1 returns the DateClotureN1 field if non-nil, zero value otherwise.

### GetDateClotureN1Ok

`func (o *ComptesAnnuels200ResponseValueInner) GetDateClotureN1Ok() (*string, bool)`

GetDateClotureN1Ok returns a tuple with the DateClotureN1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClotureN1

`func (o *ComptesAnnuels200ResponseValueInner) SetDateClotureN1(v string)`

SetDateClotureN1 sets DateClotureN1 field to given value.

### HasDateClotureN1

`func (o *ComptesAnnuels200ResponseValueInner) HasDateClotureN1() bool`

HasDateClotureN1 returns a boolean if a field has been set.

### GetDureeExerciceN

`func (o *ComptesAnnuels200ResponseValueInner) GetDureeExerciceN() int32`

GetDureeExerciceN returns the DureeExerciceN field if non-nil, zero value otherwise.

### GetDureeExerciceNOk

`func (o *ComptesAnnuels200ResponseValueInner) GetDureeExerciceNOk() (*int32, bool)`

GetDureeExerciceNOk returns a tuple with the DureeExerciceN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDureeExerciceN

`func (o *ComptesAnnuels200ResponseValueInner) SetDureeExerciceN(v int32)`

SetDureeExerciceN sets DureeExerciceN field to given value.

### HasDureeExerciceN

`func (o *ComptesAnnuels200ResponseValueInner) HasDureeExerciceN() bool`

HasDureeExerciceN returns a boolean if a field has been set.

### GetDureeExerciceN1

`func (o *ComptesAnnuels200ResponseValueInner) GetDureeExerciceN1() int32`

GetDureeExerciceN1 returns the DureeExerciceN1 field if non-nil, zero value otherwise.

### GetDureeExerciceN1Ok

`func (o *ComptesAnnuels200ResponseValueInner) GetDureeExerciceN1Ok() (*int32, bool)`

GetDureeExerciceN1Ok returns a tuple with the DureeExerciceN1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDureeExerciceN1

`func (o *ComptesAnnuels200ResponseValueInner) SetDureeExerciceN1(v int32)`

SetDureeExerciceN1 sets DureeExerciceN1 field to given value.

### HasDureeExerciceN1

`func (o *ComptesAnnuels200ResponseValueInner) HasDureeExerciceN1() bool`

HasDureeExerciceN1 returns a boolean if a field has been set.

### GetTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) GetTypeComptes() string`

GetTypeComptes returns the TypeComptes field if non-nil, zero value otherwise.

### GetTypeComptesOk

`func (o *ComptesAnnuels200ResponseValueInner) GetTypeComptesOk() (*string, bool)`

GetTypeComptesOk returns a tuple with the TypeComptes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) SetTypeComptes(v string)`

SetTypeComptes sets TypeComptes field to given value.

### HasTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) HasTypeComptes() bool`

HasTypeComptes returns a boolean if a field has been set.

### GetLibelleTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) GetLibelleTypeComptes() string`

GetLibelleTypeComptes returns the LibelleTypeComptes field if non-nil, zero value otherwise.

### GetLibelleTypeComptesOk

`func (o *ComptesAnnuels200ResponseValueInner) GetLibelleTypeComptesOk() (*string, bool)`

GetLibelleTypeComptesOk returns a tuple with the LibelleTypeComptes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) SetLibelleTypeComptes(v string)`

SetLibelleTypeComptes sets LibelleTypeComptes field to given value.

### HasLibelleTypeComptes

`func (o *ComptesAnnuels200ResponseValueInner) HasLibelleTypeComptes() bool`

HasLibelleTypeComptes returns a boolean if a field has been set.

### GetDevise

`func (o *ComptesAnnuels200ResponseValueInner) GetDevise() string`

GetDevise returns the Devise field if non-nil, zero value otherwise.

### GetDeviseOk

`func (o *ComptesAnnuels200ResponseValueInner) GetDeviseOk() (*string, bool)`

GetDeviseOk returns a tuple with the Devise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevise

`func (o *ComptesAnnuels200ResponseValueInner) SetDevise(v string)`

SetDevise sets Devise field to given value.

### HasDevise

`func (o *ComptesAnnuels200ResponseValueInner) HasDevise() bool`

HasDevise returns a boolean if a field has been set.

### GetDeviseOrigine

`func (o *ComptesAnnuels200ResponseValueInner) GetDeviseOrigine() string`

GetDeviseOrigine returns the DeviseOrigine field if non-nil, zero value otherwise.

### GetDeviseOrigineOk

`func (o *ComptesAnnuels200ResponseValueInner) GetDeviseOrigineOk() (*string, bool)`

GetDeviseOrigineOk returns a tuple with the DeviseOrigine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviseOrigine

`func (o *ComptesAnnuels200ResponseValueInner) SetDeviseOrigine(v string)`

SetDeviseOrigine sets DeviseOrigine field to given value.

### HasDeviseOrigine

`func (o *ComptesAnnuels200ResponseValueInner) HasDeviseOrigine() bool`

HasDeviseOrigine returns a boolean if a field has been set.

### GetConfidentialite

`func (o *ComptesAnnuels200ResponseValueInner) GetConfidentialite() bool`

GetConfidentialite returns the Confidentialite field if non-nil, zero value otherwise.

### GetConfidentialiteOk

`func (o *ComptesAnnuels200ResponseValueInner) GetConfidentialiteOk() (*bool, bool)`

GetConfidentialiteOk returns a tuple with the Confidentialite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialite

`func (o *ComptesAnnuels200ResponseValueInner) SetConfidentialite(v bool)`

SetConfidentialite sets Confidentialite field to given value.

### HasConfidentialite

`func (o *ComptesAnnuels200ResponseValueInner) HasConfidentialite() bool`

HasConfidentialite returns a boolean if a field has been set.

### GetConfidentialiteCompteDeResultat

`func (o *ComptesAnnuels200ResponseValueInner) GetConfidentialiteCompteDeResultat() bool`

GetConfidentialiteCompteDeResultat returns the ConfidentialiteCompteDeResultat field if non-nil, zero value otherwise.

### GetConfidentialiteCompteDeResultatOk

`func (o *ComptesAnnuels200ResponseValueInner) GetConfidentialiteCompteDeResultatOk() (*bool, bool)`

GetConfidentialiteCompteDeResultatOk returns a tuple with the ConfidentialiteCompteDeResultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialiteCompteDeResultat

`func (o *ComptesAnnuels200ResponseValueInner) SetConfidentialiteCompteDeResultat(v bool)`

SetConfidentialiteCompteDeResultat sets ConfidentialiteCompteDeResultat field to given value.

### HasConfidentialiteCompteDeResultat

`func (o *ComptesAnnuels200ResponseValueInner) HasConfidentialiteCompteDeResultat() bool`

HasConfidentialiteCompteDeResultat returns a boolean if a field has been set.

### GetCoherenceComptable

`func (o *ComptesAnnuels200ResponseValueInner) GetCoherenceComptable() bool`

GetCoherenceComptable returns the CoherenceComptable field if non-nil, zero value otherwise.

### GetCoherenceComptableOk

`func (o *ComptesAnnuels200ResponseValueInner) GetCoherenceComptableOk() (*bool, bool)`

GetCoherenceComptableOk returns a tuple with the CoherenceComptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoherenceComptable

`func (o *ComptesAnnuels200ResponseValueInner) SetCoherenceComptable(v bool)`

SetCoherenceComptable sets CoherenceComptable field to given value.

### HasCoherenceComptable

`func (o *ComptesAnnuels200ResponseValueInner) HasCoherenceComptable() bool`

HasCoherenceComptable returns a boolean if a field has been set.

### GetTypeSaisie

`func (o *ComptesAnnuels200ResponseValueInner) GetTypeSaisie() string`

GetTypeSaisie returns the TypeSaisie field if non-nil, zero value otherwise.

### GetTypeSaisieOk

`func (o *ComptesAnnuels200ResponseValueInner) GetTypeSaisieOk() (*string, bool)`

GetTypeSaisieOk returns a tuple with the TypeSaisie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSaisie

`func (o *ComptesAnnuels200ResponseValueInner) SetTypeSaisie(v string)`

SetTypeSaisie sets TypeSaisie field to given value.

### HasTypeSaisie

`func (o *ComptesAnnuels200ResponseValueInner) HasTypeSaisie() bool`

HasTypeSaisie returns a boolean if a field has been set.

### GetInformationsTraitement

`func (o *ComptesAnnuels200ResponseValueInner) GetInformationsTraitement() []string`

GetInformationsTraitement returns the InformationsTraitement field if non-nil, zero value otherwise.

### GetInformationsTraitementOk

`func (o *ComptesAnnuels200ResponseValueInner) GetInformationsTraitementOk() (*[]string, bool)`

GetInformationsTraitementOk returns a tuple with the InformationsTraitement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationsTraitement

`func (o *ComptesAnnuels200ResponseValueInner) SetInformationsTraitement(v []string)`

SetInformationsTraitement sets InformationsTraitement field to given value.

### HasInformationsTraitement

`func (o *ComptesAnnuels200ResponseValueInner) HasInformationsTraitement() bool`

HasInformationsTraitement returns a boolean if a field has been set.

### GetSections

`func (o *ComptesAnnuels200ResponseValueInner) GetSections() []ComptesAnnuels200ResponseValueInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ComptesAnnuels200ResponseValueInner) GetSectionsOk() (*[]ComptesAnnuels200ResponseValueInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ComptesAnnuels200ResponseValueInner) SetSections(v []ComptesAnnuels200ResponseValueInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *ComptesAnnuels200ResponseValueInner) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetRatios

`func (o *ComptesAnnuels200ResponseValueInner) GetRatios() Ratios`

GetRatios returns the Ratios field if non-nil, zero value otherwise.

### GetRatiosOk

`func (o *ComptesAnnuels200ResponseValueInner) GetRatiosOk() (*Ratios, bool)`

GetRatiosOk returns a tuple with the Ratios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatios

`func (o *ComptesAnnuels200ResponseValueInner) SetRatios(v Ratios)`

SetRatios sets Ratios field to given value.

### HasRatios

`func (o *ComptesAnnuels200ResponseValueInner) HasRatios() bool`

HasRatios returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


