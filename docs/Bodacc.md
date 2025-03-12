# Bodacc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumeroParution** | Pointer to **string** | Numéro de parution de la publication. | [optional] 
**Date** | Pointer to **string** | Date de la publication, au format AAAA-MM-JJ. | [optional] 
**NumeroAnnonce** | Pointer to **string** | Numéro d&#39;annonce de la publication. | [optional] 
**Bodacc** | Pointer to **string** | Bodacc de la publication (A, B ou C). | [optional] 
**Type** | Pointer to **string** | Type de la publication parmi la liste suivante : Création, Immatriculation, Modification, Vente, Achat, Radiation, Procédure collective, Dépôt des comptes. | [optional] 
**Greffe** | Pointer to **string** | Greffe de publication. | [optional] 

## Methods

### NewBodacc

`func NewBodacc() *Bodacc`

NewBodacc instantiates a new Bodacc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodaccWithDefaults

`func NewBodaccWithDefaults() *Bodacc`

NewBodaccWithDefaults instantiates a new Bodacc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumeroParution

`func (o *Bodacc) GetNumeroParution() string`

GetNumeroParution returns the NumeroParution field if non-nil, zero value otherwise.

### GetNumeroParutionOk

`func (o *Bodacc) GetNumeroParutionOk() (*string, bool)`

GetNumeroParutionOk returns a tuple with the NumeroParution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroParution

`func (o *Bodacc) SetNumeroParution(v string)`

SetNumeroParution sets NumeroParution field to given value.

### HasNumeroParution

`func (o *Bodacc) HasNumeroParution() bool`

HasNumeroParution returns a boolean if a field has been set.

### GetDate

`func (o *Bodacc) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Bodacc) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Bodacc) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Bodacc) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNumeroAnnonce

`func (o *Bodacc) GetNumeroAnnonce() string`

GetNumeroAnnonce returns the NumeroAnnonce field if non-nil, zero value otherwise.

### GetNumeroAnnonceOk

`func (o *Bodacc) GetNumeroAnnonceOk() (*string, bool)`

GetNumeroAnnonceOk returns a tuple with the NumeroAnnonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroAnnonce

`func (o *Bodacc) SetNumeroAnnonce(v string)`

SetNumeroAnnonce sets NumeroAnnonce field to given value.

### HasNumeroAnnonce

`func (o *Bodacc) HasNumeroAnnonce() bool`

HasNumeroAnnonce returns a boolean if a field has been set.

### GetBodacc

`func (o *Bodacc) GetBodacc() string`

GetBodacc returns the Bodacc field if non-nil, zero value otherwise.

### GetBodaccOk

`func (o *Bodacc) GetBodaccOk() (*string, bool)`

GetBodaccOk returns a tuple with the Bodacc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodacc

`func (o *Bodacc) SetBodacc(v string)`

SetBodacc sets Bodacc field to given value.

### HasBodacc

`func (o *Bodacc) HasBodacc() bool`

HasBodacc returns a boolean if a field has been set.

### GetType

`func (o *Bodacc) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bodacc) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bodacc) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Bodacc) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGreffe

`func (o *Bodacc) GetGreffe() string`

GetGreffe returns the Greffe field if non-nil, zero value otherwise.

### GetGreffeOk

`func (o *Bodacc) GetGreffeOk() (*string, bool)`

GetGreffeOk returns a tuple with the Greffe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreffe

`func (o *Bodacc) SetGreffe(v string)`

SetGreffe sets Greffe field to given value.

### HasGreffe

`func (o *Bodacc) HasGreffe() bool`

HasGreffe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


