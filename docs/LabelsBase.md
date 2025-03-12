# LabelsBase

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

## Methods

### NewLabelsBase

`func NewLabelsBase() *LabelsBase`

NewLabelsBase instantiates a new LabelsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsBaseWithDefaults

`func NewLabelsBaseWithDefaults() *LabelsBase`

NewLabelsBaseWithDefaults instantiates a new LabelsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *LabelsBase) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *LabelsBase) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *LabelsBase) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *LabelsBase) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetCertificats

`func (o *LabelsBase) GetCertificats() []string`

GetCertificats returns the Certificats field if non-nil, zero value otherwise.

### GetCertificatsOk

`func (o *LabelsBase) GetCertificatsOk() (*[]string, bool)`

GetCertificatsOk returns a tuple with the Certificats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificats

`func (o *LabelsBase) SetCertificats(v []string)`

SetCertificats sets Certificats field to given value.

### HasCertificats

`func (o *LabelsBase) HasCertificats() bool`

HasCertificats returns a boolean if a field has been set.

### GetSpecialites

`func (o *LabelsBase) GetSpecialites() []string`

GetSpecialites returns the Specialites field if non-nil, zero value otherwise.

### GetSpecialitesOk

`func (o *LabelsBase) GetSpecialitesOk() (*[]string, bool)`

GetSpecialitesOk returns a tuple with the Specialites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialites

`func (o *LabelsBase) SetSpecialites(v []string)`

SetSpecialites sets Specialites field to given value.

### HasSpecialites

`func (o *LabelsBase) HasSpecialites() bool`

HasSpecialites returns a boolean if a field has been set.

### GetNotes

`func (o *LabelsBase) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *LabelsBase) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *LabelsBase) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *LabelsBase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNumeroImmatriculation

`func (o *LabelsBase) GetNumeroImmatriculation() string`

GetNumeroImmatriculation returns the NumeroImmatriculation field if non-nil, zero value otherwise.

### GetNumeroImmatriculationOk

`func (o *LabelsBase) GetNumeroImmatriculationOk() (*string, bool)`

GetNumeroImmatriculationOk returns a tuple with the NumeroImmatriculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroImmatriculation

`func (o *LabelsBase) SetNumeroImmatriculation(v string)`

SetNumeroImmatriculation sets NumeroImmatriculation field to given value.

### HasNumeroImmatriculation

`func (o *LabelsBase) HasNumeroImmatriculation() bool`

HasNumeroImmatriculation returns a boolean if a field has been set.

### GetInscriptions

`func (o *LabelsBase) GetInscriptions() []LabelsBaseInscriptionsInner`

GetInscriptions returns the Inscriptions field if non-nil, zero value otherwise.

### GetInscriptionsOk

`func (o *LabelsBase) GetInscriptionsOk() (*[]LabelsBaseInscriptionsInner, bool)`

GetInscriptionsOk returns a tuple with the Inscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInscriptions

`func (o *LabelsBase) SetInscriptions(v []LabelsBaseInscriptionsInner)`

SetInscriptions sets Inscriptions field to given value.

### HasInscriptions

`func (o *LabelsBase) HasInscriptions() bool`

HasInscriptions returns a boolean if a field has been set.

### GetMentions

`func (o *LabelsBase) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *LabelsBase) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *LabelsBase) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *LabelsBase) HasMentions() bool`

HasMentions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


