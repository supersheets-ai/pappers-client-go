# LienSuccession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siret** | Pointer to **string** | Numéro siret de l&#39;établissement au format xxxxxxxxxxxxxx. | [optional] 
**Date** | Pointer to **string** | Date à laquelle la succession a eu lieu. | [optional] 
**TransfertSiege** | Pointer to **bool** | Vrai si le lien de succession concerne l&#39;établissement siège, faux sinon. | [optional] 
**ContinuiteEconomique** | Pointer to **bool** | Vrai s&#39;il y a [continuité économique](https://www.sirene.fr/sirene/public/variable/continuiteEconomique), faux sinon. | [optional] 

## Methods

### NewLienSuccession

`func NewLienSuccession() *LienSuccession`

NewLienSuccession instantiates a new LienSuccession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLienSuccessionWithDefaults

`func NewLienSuccessionWithDefaults() *LienSuccession`

NewLienSuccessionWithDefaults instantiates a new LienSuccession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *LienSuccession) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *LienSuccession) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *LienSuccession) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *LienSuccession) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetDate

`func (o *LienSuccession) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LienSuccession) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LienSuccession) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *LienSuccession) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTransfertSiege

`func (o *LienSuccession) GetTransfertSiege() bool`

GetTransfertSiege returns the TransfertSiege field if non-nil, zero value otherwise.

### GetTransfertSiegeOk

`func (o *LienSuccession) GetTransfertSiegeOk() (*bool, bool)`

GetTransfertSiegeOk returns a tuple with the TransfertSiege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfertSiege

`func (o *LienSuccession) SetTransfertSiege(v bool)`

SetTransfertSiege sets TransfertSiege field to given value.

### HasTransfertSiege

`func (o *LienSuccession) HasTransfertSiege() bool`

HasTransfertSiege returns a boolean if a field has been set.

### GetContinuiteEconomique

`func (o *LienSuccession) GetContinuiteEconomique() bool`

GetContinuiteEconomique returns the ContinuiteEconomique field if non-nil, zero value otherwise.

### GetContinuiteEconomiqueOk

`func (o *LienSuccession) GetContinuiteEconomiqueOk() (*bool, bool)`

GetContinuiteEconomiqueOk returns a tuple with the ContinuiteEconomique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuiteEconomique

`func (o *LienSuccession) SetContinuiteEconomique(v bool)`

SetContinuiteEconomique sets ContinuiteEconomique field to given value.

### HasContinuiteEconomique

`func (o *LienSuccession) HasContinuiteEconomique() bool`

HasContinuiteEconomique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


