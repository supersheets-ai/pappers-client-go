# PersonneMarque

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | Pointer to **string** | Siren de l&#39;entité, dans le cas d&#39;une personne morale. | [optional] 
**EntiteLegale** | Pointer to **string** | Entité légale. | [optional] 
**Nom** | Pointer to **string** | Nom de l&#39;entité. | [optional] 
**Batiment** | Pointer to **string** | Bâtiment de l&#39;entité. | [optional] 
**Rue** | Pointer to **string** | Rue de l&#39;entité. | [optional] 
**Ville** | Pointer to **string** | Ville de l&#39;entité. | [optional] 
**BoitePostale** | Pointer to **string** | Boîte postale de l&#39;entité. | [optional] 
**CodePostal** | Pointer to **string** | Code postal de l&#39;entité. | [optional] 
**CodePays** | Pointer to **string** | Code pays de l&#39;entité. | [optional] 

## Methods

### NewPersonneMarque

`func NewPersonneMarque() *PersonneMarque`

NewPersonneMarque instantiates a new PersonneMarque object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonneMarqueWithDefaults

`func NewPersonneMarqueWithDefaults() *PersonneMarque`

NewPersonneMarqueWithDefaults instantiates a new PersonneMarque object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *PersonneMarque) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *PersonneMarque) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *PersonneMarque) SetSiren(v string)`

SetSiren sets Siren field to given value.

### HasSiren

`func (o *PersonneMarque) HasSiren() bool`

HasSiren returns a boolean if a field has been set.

### GetEntiteLegale

`func (o *PersonneMarque) GetEntiteLegale() string`

GetEntiteLegale returns the EntiteLegale field if non-nil, zero value otherwise.

### GetEntiteLegaleOk

`func (o *PersonneMarque) GetEntiteLegaleOk() (*string, bool)`

GetEntiteLegaleOk returns a tuple with the EntiteLegale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntiteLegale

`func (o *PersonneMarque) SetEntiteLegale(v string)`

SetEntiteLegale sets EntiteLegale field to given value.

### HasEntiteLegale

`func (o *PersonneMarque) HasEntiteLegale() bool`

HasEntiteLegale returns a boolean if a field has been set.

### GetNom

`func (o *PersonneMarque) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *PersonneMarque) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *PersonneMarque) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *PersonneMarque) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetBatiment

`func (o *PersonneMarque) GetBatiment() string`

GetBatiment returns the Batiment field if non-nil, zero value otherwise.

### GetBatimentOk

`func (o *PersonneMarque) GetBatimentOk() (*string, bool)`

GetBatimentOk returns a tuple with the Batiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatiment

`func (o *PersonneMarque) SetBatiment(v string)`

SetBatiment sets Batiment field to given value.

### HasBatiment

`func (o *PersonneMarque) HasBatiment() bool`

HasBatiment returns a boolean if a field has been set.

### GetRue

`func (o *PersonneMarque) GetRue() string`

GetRue returns the Rue field if non-nil, zero value otherwise.

### GetRueOk

`func (o *PersonneMarque) GetRueOk() (*string, bool)`

GetRueOk returns a tuple with the Rue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRue

`func (o *PersonneMarque) SetRue(v string)`

SetRue sets Rue field to given value.

### HasRue

`func (o *PersonneMarque) HasRue() bool`

HasRue returns a boolean if a field has been set.

### GetVille

`func (o *PersonneMarque) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *PersonneMarque) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *PersonneMarque) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *PersonneMarque) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetBoitePostale

`func (o *PersonneMarque) GetBoitePostale() string`

GetBoitePostale returns the BoitePostale field if non-nil, zero value otherwise.

### GetBoitePostaleOk

`func (o *PersonneMarque) GetBoitePostaleOk() (*string, bool)`

GetBoitePostaleOk returns a tuple with the BoitePostale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoitePostale

`func (o *PersonneMarque) SetBoitePostale(v string)`

SetBoitePostale sets BoitePostale field to given value.

### HasBoitePostale

`func (o *PersonneMarque) HasBoitePostale() bool`

HasBoitePostale returns a boolean if a field has been set.

### GetCodePostal

`func (o *PersonneMarque) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *PersonneMarque) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *PersonneMarque) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *PersonneMarque) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetCodePays

`func (o *PersonneMarque) GetCodePays() string`

GetCodePays returns the CodePays field if non-nil, zero value otherwise.

### GetCodePaysOk

`func (o *PersonneMarque) GetCodePaysOk() (*string, bool)`

GetCodePaysOk returns a tuple with the CodePays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePays

`func (o *PersonneMarque) SetCodePays(v string)`

SetCodePays sets CodePays field to given value.

### HasCodePays

`func (o *PersonneMarque) HasCodePays() bool`

HasCodePays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


