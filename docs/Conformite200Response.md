# Conformite200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonnePolitiquementExposee** | Pointer to [**PersonnePolitiquementExposee**](PersonnePolitiquementExposee.md) |  | [optional] 
**SanctionsEnCours** | Pointer to **bool** | Vaut vrai si le bénéficiaire effectif est actuellement sous sanction. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 
**Sanctions** | Pointer to [**[]Sanction**](Sanction.md) | Liste des sanctions du bénéficiaire effectif. Uniquement présent si demandé dans les champs supplémentaires. | [optional] 

## Methods

### NewConformite200Response

`func NewConformite200Response() *Conformite200Response`

NewConformite200Response instantiates a new Conformite200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConformite200ResponseWithDefaults

`func NewConformite200ResponseWithDefaults() *Conformite200Response`

NewConformite200ResponseWithDefaults instantiates a new Conformite200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonnePolitiquementExposee

`func (o *Conformite200Response) GetPersonnePolitiquementExposee() PersonnePolitiquementExposee`

GetPersonnePolitiquementExposee returns the PersonnePolitiquementExposee field if non-nil, zero value otherwise.

### GetPersonnePolitiquementExposeeOk

`func (o *Conformite200Response) GetPersonnePolitiquementExposeeOk() (*PersonnePolitiquementExposee, bool)`

GetPersonnePolitiquementExposeeOk returns a tuple with the PersonnePolitiquementExposee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonnePolitiquementExposee

`func (o *Conformite200Response) SetPersonnePolitiquementExposee(v PersonnePolitiquementExposee)`

SetPersonnePolitiquementExposee sets PersonnePolitiquementExposee field to given value.

### HasPersonnePolitiquementExposee

`func (o *Conformite200Response) HasPersonnePolitiquementExposee() bool`

HasPersonnePolitiquementExposee returns a boolean if a field has been set.

### GetSanctionsEnCours

`func (o *Conformite200Response) GetSanctionsEnCours() bool`

GetSanctionsEnCours returns the SanctionsEnCours field if non-nil, zero value otherwise.

### GetSanctionsEnCoursOk

`func (o *Conformite200Response) GetSanctionsEnCoursOk() (*bool, bool)`

GetSanctionsEnCoursOk returns a tuple with the SanctionsEnCours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsEnCours

`func (o *Conformite200Response) SetSanctionsEnCours(v bool)`

SetSanctionsEnCours sets SanctionsEnCours field to given value.

### HasSanctionsEnCours

`func (o *Conformite200Response) HasSanctionsEnCours() bool`

HasSanctionsEnCours returns a boolean if a field has been set.

### GetSanctions

`func (o *Conformite200Response) GetSanctions() []Sanction`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *Conformite200Response) GetSanctionsOk() (*[]Sanction, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *Conformite200Response) SetSanctions(v []Sanction)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *Conformite200Response) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


