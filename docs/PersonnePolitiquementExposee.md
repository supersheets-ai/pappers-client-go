# PersonnePolitiquementExposee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnCours** | Pointer to **bool** | Vaut vrai si la personne est actuellement politiquement exposée. | [optional] 
**Fonctions** | Pointer to [**[]PersonnePolitiquementExposeeFonctionsInner**](PersonnePolitiquementExposeeFonctionsInner.md) | Liste des fonctions actuelles et passées de la personne politiquement exposée. | [optional] 

## Methods

### NewPersonnePolitiquementExposee

`func NewPersonnePolitiquementExposee() *PersonnePolitiquementExposee`

NewPersonnePolitiquementExposee instantiates a new PersonnePolitiquementExposee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonnePolitiquementExposeeWithDefaults

`func NewPersonnePolitiquementExposeeWithDefaults() *PersonnePolitiquementExposee`

NewPersonnePolitiquementExposeeWithDefaults instantiates a new PersonnePolitiquementExposee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnCours

`func (o *PersonnePolitiquementExposee) GetEnCours() bool`

GetEnCours returns the EnCours field if non-nil, zero value otherwise.

### GetEnCoursOk

`func (o *PersonnePolitiquementExposee) GetEnCoursOk() (*bool, bool)`

GetEnCoursOk returns a tuple with the EnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnCours

`func (o *PersonnePolitiquementExposee) SetEnCours(v bool)`

SetEnCours sets EnCours field to given value.

### HasEnCours

`func (o *PersonnePolitiquementExposee) HasEnCours() bool`

HasEnCours returns a boolean if a field has been set.

### GetFonctions

`func (o *PersonnePolitiquementExposee) GetFonctions() []PersonnePolitiquementExposeeFonctionsInner`

GetFonctions returns the Fonctions field if non-nil, zero value otherwise.

### GetFonctionsOk

`func (o *PersonnePolitiquementExposee) GetFonctionsOk() (*[]PersonnePolitiquementExposeeFonctionsInner, bool)`

GetFonctionsOk returns a tuple with the Fonctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonctions

`func (o *PersonnePolitiquementExposee) SetFonctions(v []PersonnePolitiquementExposeeFonctionsInner)`

SetFonctions sets Fonctions field to given value.

### HasFonctions

`func (o *PersonnePolitiquementExposee) HasFonctions() bool`

HasFonctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


