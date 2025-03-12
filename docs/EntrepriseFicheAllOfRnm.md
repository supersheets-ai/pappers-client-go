# EntrepriseFicheAllOfRnm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateImmatriculation** | Pointer to **string** | Date d&#39;immatriculation au Répertoire des Métiers, au format AAAA-MM-JJ. | [optional] 
**DateRadiation** | Pointer to **string** | Date de radiation du Répertoire des Métiers, le cas échéant, au format AAAA-MM-JJ. | [optional] 
**DateDebutActivite** | Pointer to **string** | Date de début d&#39;activé déclarée au Répertoire des Métiers, au format AAAA-MM-JJ. | [optional] 
**DateCessationActivite** | Pointer to **string** | Date de cessation d&#39;activité déclarée au Répertoire des Métiers, le cas échéant, au format AAAA-MM-JJ. | [optional] 
**ChambreDesMetiers** | Pointer to **string** | Chambre des métiers où l&#39;entreprise est immatriculée. | [optional] 
**Qualification** | Pointer to **string** | Qualification retenue par le Répertoire des Métiers. | [optional] 
**DerniereMiseAJour** | Pointer to **string** | Date de dernière mise à jour de l&#39;entreprise au Répertoire des Métiers, au format AAAA-MM-JJ. | [optional] 

## Methods

### NewEntrepriseFicheAllOfRnm

`func NewEntrepriseFicheAllOfRnm() *EntrepriseFicheAllOfRnm`

NewEntrepriseFicheAllOfRnm instantiates a new EntrepriseFicheAllOfRnm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfRnmWithDefaults

`func NewEntrepriseFicheAllOfRnmWithDefaults() *EntrepriseFicheAllOfRnm`

NewEntrepriseFicheAllOfRnmWithDefaults instantiates a new EntrepriseFicheAllOfRnm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateImmatriculation

`func (o *EntrepriseFicheAllOfRnm) GetDateImmatriculation() string`

GetDateImmatriculation returns the DateImmatriculation field if non-nil, zero value otherwise.

### GetDateImmatriculationOk

`func (o *EntrepriseFicheAllOfRnm) GetDateImmatriculationOk() (*string, bool)`

GetDateImmatriculationOk returns a tuple with the DateImmatriculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateImmatriculation

`func (o *EntrepriseFicheAllOfRnm) SetDateImmatriculation(v string)`

SetDateImmatriculation sets DateImmatriculation field to given value.

### HasDateImmatriculation

`func (o *EntrepriseFicheAllOfRnm) HasDateImmatriculation() bool`

HasDateImmatriculation returns a boolean if a field has been set.

### GetDateRadiation

`func (o *EntrepriseFicheAllOfRnm) GetDateRadiation() string`

GetDateRadiation returns the DateRadiation field if non-nil, zero value otherwise.

### GetDateRadiationOk

`func (o *EntrepriseFicheAllOfRnm) GetDateRadiationOk() (*string, bool)`

GetDateRadiationOk returns a tuple with the DateRadiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRadiation

`func (o *EntrepriseFicheAllOfRnm) SetDateRadiation(v string)`

SetDateRadiation sets DateRadiation field to given value.

### HasDateRadiation

`func (o *EntrepriseFicheAllOfRnm) HasDateRadiation() bool`

HasDateRadiation returns a boolean if a field has been set.

### GetDateDebutActivite

`func (o *EntrepriseFicheAllOfRnm) GetDateDebutActivite() string`

GetDateDebutActivite returns the DateDebutActivite field if non-nil, zero value otherwise.

### GetDateDebutActiviteOk

`func (o *EntrepriseFicheAllOfRnm) GetDateDebutActiviteOk() (*string, bool)`

GetDateDebutActiviteOk returns a tuple with the DateDebutActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebutActivite

`func (o *EntrepriseFicheAllOfRnm) SetDateDebutActivite(v string)`

SetDateDebutActivite sets DateDebutActivite field to given value.

### HasDateDebutActivite

`func (o *EntrepriseFicheAllOfRnm) HasDateDebutActivite() bool`

HasDateDebutActivite returns a boolean if a field has been set.

### GetDateCessationActivite

`func (o *EntrepriseFicheAllOfRnm) GetDateCessationActivite() string`

GetDateCessationActivite returns the DateCessationActivite field if non-nil, zero value otherwise.

### GetDateCessationActiviteOk

`func (o *EntrepriseFicheAllOfRnm) GetDateCessationActiviteOk() (*string, bool)`

GetDateCessationActiviteOk returns a tuple with the DateCessationActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCessationActivite

`func (o *EntrepriseFicheAllOfRnm) SetDateCessationActivite(v string)`

SetDateCessationActivite sets DateCessationActivite field to given value.

### HasDateCessationActivite

`func (o *EntrepriseFicheAllOfRnm) HasDateCessationActivite() bool`

HasDateCessationActivite returns a boolean if a field has been set.

### GetChambreDesMetiers

`func (o *EntrepriseFicheAllOfRnm) GetChambreDesMetiers() string`

GetChambreDesMetiers returns the ChambreDesMetiers field if non-nil, zero value otherwise.

### GetChambreDesMetiersOk

`func (o *EntrepriseFicheAllOfRnm) GetChambreDesMetiersOk() (*string, bool)`

GetChambreDesMetiersOk returns a tuple with the ChambreDesMetiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChambreDesMetiers

`func (o *EntrepriseFicheAllOfRnm) SetChambreDesMetiers(v string)`

SetChambreDesMetiers sets ChambreDesMetiers field to given value.

### HasChambreDesMetiers

`func (o *EntrepriseFicheAllOfRnm) HasChambreDesMetiers() bool`

HasChambreDesMetiers returns a boolean if a field has been set.

### GetQualification

`func (o *EntrepriseFicheAllOfRnm) GetQualification() string`

GetQualification returns the Qualification field if non-nil, zero value otherwise.

### GetQualificationOk

`func (o *EntrepriseFicheAllOfRnm) GetQualificationOk() (*string, bool)`

GetQualificationOk returns a tuple with the Qualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualification

`func (o *EntrepriseFicheAllOfRnm) SetQualification(v string)`

SetQualification sets Qualification field to given value.

### HasQualification

`func (o *EntrepriseFicheAllOfRnm) HasQualification() bool`

HasQualification returns a boolean if a field has been set.

### GetDerniereMiseAJour

`func (o *EntrepriseFicheAllOfRnm) GetDerniereMiseAJour() string`

GetDerniereMiseAJour returns the DerniereMiseAJour field if non-nil, zero value otherwise.

### GetDerniereMiseAJourOk

`func (o *EntrepriseFicheAllOfRnm) GetDerniereMiseAJourOk() (*string, bool)`

GetDerniereMiseAJourOk returns a tuple with the DerniereMiseAJour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerniereMiseAJour

`func (o *EntrepriseFicheAllOfRnm) SetDerniereMiseAJour(v string)`

SetDerniereMiseAJour sets DerniereMiseAJour field to given value.

### HasDerniereMiseAJour

`func (o *EntrepriseFicheAllOfRnm) HasDerniereMiseAJour() bool`

HasDerniereMiseAJour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


