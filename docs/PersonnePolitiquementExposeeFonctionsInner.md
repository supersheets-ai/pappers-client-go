# PersonnePolitiquementExposeeFonctionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fonction** | Pointer to **string** | Nom de la fonction. | [optional] 
**Pays** | Pointer to **string** | Pays associé à la fonction. | [optional] 
**CodePays** | Pointer to **string** | Code pays associé à la fonction. | [optional] 
**EnCours** | Pointer to **bool** | Vaut vrai si la fonction est en cours. | [optional] 
**DateDebut** | Pointer to **string** | Date de début de la fonction. | [optional] 
**DateFin** | Pointer to **string** | Date de fin de la fonction. | [optional] 
**Sources** | Pointer to [**[]PersonnePolitiquementExposeeFonctionsInnerSourcesInner**](PersonnePolitiquementExposeeFonctionsInnerSourcesInner.md) | Liste des sources. | [optional] 

## Methods

### NewPersonnePolitiquementExposeeFonctionsInner

`func NewPersonnePolitiquementExposeeFonctionsInner() *PersonnePolitiquementExposeeFonctionsInner`

NewPersonnePolitiquementExposeeFonctionsInner instantiates a new PersonnePolitiquementExposeeFonctionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonnePolitiquementExposeeFonctionsInnerWithDefaults

`func NewPersonnePolitiquementExposeeFonctionsInnerWithDefaults() *PersonnePolitiquementExposeeFonctionsInner`

NewPersonnePolitiquementExposeeFonctionsInnerWithDefaults instantiates a new PersonnePolitiquementExposeeFonctionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFonction

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetFonction() string`

GetFonction returns the Fonction field if non-nil, zero value otherwise.

### GetFonctionOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetFonctionOk() (*string, bool)`

GetFonctionOk returns a tuple with the Fonction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonction

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetFonction(v string)`

SetFonction sets Fonction field to given value.

### HasFonction

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasFonction() bool`

HasFonction returns a boolean if a field has been set.

### GetPays

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasPays() bool`

HasPays returns a boolean if a field has been set.

### GetCodePays

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.

### GetEnCours

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetEnCours() bool`

GetEnCours returns the EnCours field if non-nil, zero value otherwise.

### GetEnCoursOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetEnCoursOk() (*bool, bool)`

GetEnCoursOk returns a tuple with the EnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnCours

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetEnCours(v bool)`

SetEnCours sets EnCours field to given value.

### HasEnCours

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasEnCours() bool`

HasEnCours returns a boolean if a field has been set.

### GetDateDebut

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetDateDebut() string`

GetDateDebut returns the DateDebut field if non-nil, zero value otherwise.

### GetDateDebutOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetDateDebutOk() (*string, bool)`

GetDateDebutOk returns a tuple with the DateDebut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDebut

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetDateDebut(v string)`

SetDateDebut sets DateDebut field to given value.

### HasDateDebut

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasDateDebut() bool`

HasDateDebut returns a boolean if a field has been set.

### GetDateFin

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetDateFin() string`

GetDateFin returns the DateFin field if non-nil, zero value otherwise.

### GetDateFinOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetDateFinOk() (*string, bool)`

GetDateFinOk returns a tuple with the DateFin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFin

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetDateFin(v string)`

SetDateFin sets DateFin field to given value.

### HasDateFin

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasDateFin() bool`

HasDateFin returns a boolean if a field has been set.

### GetSources

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetSources() []PersonnePolitiquementExposeeFonctionsInnerSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PersonnePolitiquementExposeeFonctionsInner) GetSourcesOk() (*[]PersonnePolitiquementExposeeFonctionsInnerSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PersonnePolitiquementExposeeFonctionsInner) SetSources(v []PersonnePolitiquementExposeeFonctionsInnerSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PersonnePolitiquementExposeeFonctionsInner) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


