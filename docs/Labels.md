# Labels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | Pointer to **string** | Nom du label. | [optional] 
**Certificats** | Pointer to **[]string** | Label RGE seulement : Liste des certificats. | [optional] 
**Specialites** | Pointer to **[]string** | Label QUALIOPI seulement : Liste des spécialités. | [optional] 
**Notes** | Pointer to **[]string** | Label EGALITE seulement : Liste des notes. | [optional] 
**NumeroImmatriculation** | Pointer to **string** | Label ORIAS et CCI seulement : Numéro d&#39;immatriculation ORIAS ou CCI. Uniquement présent si demandé avec le champ supplémentaire &#x60;label:orias&#x60; ou &#x60;label:cci&#x60;. | [optional] 
**Inscriptions** | Pointer to [**[]LabelsBaseInscriptionsInner**](LabelsBaseInscriptionsInner.md) | Label ORIAS seulement : Liste des inscriptions ORIAS. Uniquement présent si demandé avec le champ supplémentaire &#x60;label:orias&#x60;. | [optional] 
**Mentions** | Pointer to **[]string** | Label CCI seulement : Liste des mentions. | [optional] 
**NbEtablissementsConcernes** | Pointer to **int32** | Nombre d&#39;établissements concernés par le label, dans le cas d&#39;un label lié aux établissements. Null sinon. | [optional] 

## Methods

### NewLabels

`func NewLabels() *Labels`

NewLabels instantiates a new Labels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsWithDefaults

`func NewLabelsWithDefaults() *Labels`

NewLabelsWithDefaults instantiates a new Labels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *Labels) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *Labels) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *Labels) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *Labels) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetCertificats

`func (o *Labels) GetCertificats() []string`

GetCertificats returns the Certificats field if non-nil, zero value otherwise.

### GetCertificatsOk

`func (o *Labels) GetCertificatsOk() (*[]string, bool)`

GetCertificatsOk returns a tuple with the Certificats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificats

`func (o *Labels) SetCertificats(v []string)`

SetCertificats sets Certificats field to given value.

### HasCertificats

`func (o *Labels) HasCertificats() bool`

HasCertificats returns a boolean if a field has been set.

### GetSpecialites

`func (o *Labels) GetSpecialites() []string`

GetSpecialites returns the Specialites field if non-nil, zero value otherwise.

### GetSpecialitesOk

`func (o *Labels) GetSpecialitesOk() (*[]string, bool)`

GetSpecialitesOk returns a tuple with the Specialites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialites

`func (o *Labels) SetSpecialites(v []string)`

SetSpecialites sets Specialites field to given value.

### HasSpecialites

`func (o *Labels) HasSpecialites() bool`

HasSpecialites returns a boolean if a field has been set.

### GetNotes

`func (o *Labels) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Labels) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Labels) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Labels) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumeroImmatriculation

`func (o *Labels) GetNumeroImmatriculation() string`

GetNumeroImmatriculation returns the NumeroImmatriculation field if non-nil, zero value otherwise.

### GetNumeroImmatriculationOk

`func (o *Labels) GetNumeroImmatriculationOk() (*string, bool)`

GetNumeroImmatriculationOk returns a tuple with the NumeroImmatriculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroImmatriculation

`func (o *Labels) SetNumeroImmatriculation(v string)`

SetNumeroImmatriculation sets NumeroImmatriculation field to given value.

### HasNumeroImmatriculation

`func (o *Labels) HasNumeroImmatriculation() bool`

HasNumeroImmatriculation returns a boolean if a field has been set.

### GetInscriptions

`func (o *Labels) GetInscriptions() []LabelsBaseInscriptionsInner`

GetInscriptions returns the Inscriptions field if non-nil, zero value otherwise.

### GetInscriptionsOk

`func (o *Labels) GetInscriptionsOk() (*[]LabelsBaseInscriptionsInner, bool)`

GetInscriptionsOk returns a tuple with the Inscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptions

`func (o *Labels) SetInscriptions(v []LabelsBaseInscriptionsInner)`

SetInscriptions sets Inscriptions field to given value.

### HasInscriptions

`func (o *Labels) HasInscriptions() bool`

HasInscriptions returns a boolean if a field has been set.

### GetMentions

`func (o *Labels) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *Labels) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *Labels) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *Labels) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetNbEtablissementsConcernes

`func (o *Labels) GetNbEtablissementsConcernes() int32`

GetNbEtablissementsConcernes returns the NbEtablissementsConcernes field if non-nil, zero value otherwise.

### GetNbEtablissementsConcernesOk

`func (o *Labels) GetNbEtablissementsConcernesOk() (*int32, bool)`

GetNbEtablissementsConcernesOk returns a tuple with the NbEtablissementsConcernes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbEtablissementsConcernes

`func (o *Labels) SetNbEtablissementsConcernes(v int32)`

SetNbEtablissementsConcernes sets NbEtablissementsConcernes field to given value.

### HasNbEtablissementsConcernes

`func (o *Labels) HasNbEtablissementsConcernes() bool`

HasNbEtablissementsConcernes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


