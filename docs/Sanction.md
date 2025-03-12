# Sanction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description de la sanction. | [optional] 
**Autorite** | Pointer to **string** | Autorité ayant prononcé la sanction. | [optional] 
**Pays** | Pointer to **string** | Pays de la sanction. | [optional] 
**CodePays** | Pointer to **string** | Code du pays de la sanction. | [optional] 
**EnCours** | Pointer to **bool** | Vaut vrai si la sanction est en cours. | [optional] 
**DateDebut** | Pointer to **string** | Date de début de la sanction. | [optional] 
**DateFin** | Pointer to **string** | Date de fin de la sanction. | [optional] 
**Sources** | Pointer to [**[]PersonnePolitiquementExposeeFonctionsInnerSourcesInner**](PersonnePolitiquementExposeeFonctionsInnerSourcesInner.md) | Liste des sources. | [optional] 

## Methods

### NewSanction

`func NewSanction() *Sanction`

NewSanction instantiates a new Sanction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSanctionWithDefaults

`func NewSanctionWithDefaults() *Sanction`

NewSanctionWithDefaults instantiates a new Sanction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Sanction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Sanction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Sanction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Sanction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutorite

`func (o *Sanction) GetAutorite() string`

GetAutorite returns the Autorite field if non-nil, zero value otherwise.

### GetAutoriteOk

`func (o *Sanction) GetAutoriteOk() (*string, bool)`

GetAutoriteOk returns a tuple with the Autorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorite

`func (o *Sanction) SetAutorite(v string)`

SetAutorite sets Autorite field to given value.

### HasAutorite

`func (o *Sanction) HasAutorite() bool`

HasAutorite returns a boolean if a field has been set.

### GetPays

`func (o *Sanction) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *Sanction) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *Sanction) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *Sanction) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *Sanction) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *Sanction) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *Sanction) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *Sanction) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetEnCours

`func (o *Sanction) GetEnCours() bool`

GetEnCours returns the EnCours field if non-nil, zero value otherwise.

### GetEnCoursOk

`func (o *Sanction) GetEnCoursOk() (*bool, bool)`

GetEnCoursOk returns a tuple with the EnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnCours

`func (o *Sanction) SetEnCours(v bool)`

SetEnCours sets EnCours field to given value.

### HasEnCours

`func (o *Sanction) HasEnCours() bool`

HasEnCours returns a boolean if a field has been set.

### GetDateDebut

`func (o *Sanction) GetDateDebut() string`

GetDateDebut returns the DateDebut field if non-nil, zero value otherwise.

### GetDateDebutOk

`func (o *Sanction) GetDateDebutOk() (*string, bool)`

GetDateDebutOk returns a tuple with the DateDebut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebut

`func (o *Sanction) SetDateDebut(v string)`

SetDateDebut sets DateDebut field to given value.

### HasDateDebut

`func (o *Sanction) HasDateDebut() bool`

HasDateDebut returns a boolean if a field has been set.

### GetDateFin

`func (o *Sanction) GetDateFin() string`

GetDateFin returns the DateFin field if non-nil, zero value otherwise.

### GetDateFinOk

`func (o *Sanction) GetDateFinOk() (*string, bool)`

GetDateFinOk returns a tuple with the DateFin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFin

`func (o *Sanction) SetDateFin(v string)`

SetDateFin sets DateFin field to given value.

### HasDateFin

`func (o *Sanction) HasDateFin() bool`

HasDateFin returns a boolean if a field has been set.

### GetSources

`func (o *Sanction) GetSources() []PersonnePolitiquementExposeeFonctionsInnerSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Sanction) GetSourcesOk() (*[]PersonnePolitiquementExposeeFonctionsInnerSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Sanction) SetSources(v []PersonnePolitiquementExposeeFonctionsInnerSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Sanction) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


