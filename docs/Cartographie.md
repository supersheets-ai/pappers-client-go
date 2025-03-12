# Cartographie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entreprises** | Pointer to [**[]CartographieEntreprisesInner**](CartographieEntreprisesInner.md) | Liste des noeuds entreprises. | [optional] 
**Personnes** | Pointer to [**[]CartographiePersonnesInner**](CartographiePersonnesInner.md) | Liste des noeuds personnes (dirigeants ou bénéficiaires effectifs). | [optional] 
**LiensEntreprisesPersonnes** | Pointer to [**[][]CartographieLiensEntreprisesPersonnesInnerInner**]([]CartographieLiensEntreprisesPersonnesInnerInner.md) | Liste des arêtes liant les noeuds entreprises avec des noeuds personnes. | [optional] 
**LiensEntreprisesEntreprises** | Pointer to **[][]string** | Liste des arêtes liant les noeuds entreprises avec des d&#39;autres noeuds entreprises. | [optional] 
**ModificationsEffectuees** | Pointer to **map[string]interface{}** | Description des paramètres ayant été automatiquement modifiés. | [optional] 

## Methods

### NewCartographie

`func NewCartographie() *Cartographie`

NewCartographie instantiates a new Cartographie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartographieWithDefaults

`func NewCartographieWithDefaults() *Cartographie`

NewCartographieWithDefaults instantiates a new Cartographie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntreprises

`func (o *Cartographie) GetEntreprises() []CartographieEntreprisesInner`

GetEntreprises returns the Entreprises field if non-nil, zero value otherwise.

### GetEntreprisesOk

`func (o *Cartographie) GetEntreprisesOk() (*[]CartographieEntreprisesInner, bool)`

GetEntreprisesOk returns a tuple with the Entreprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntreprises

`func (o *Cartographie) SetEntreprises(v []CartographieEntreprisesInner)`

SetEntreprises sets Entreprises field to given value.

### HasEntreprises

`func (o *Cartographie) HasEntreprises() bool`

HasEntreprises returns a boolean if a field has been set.

### GetPersonnes

`func (o *Cartographie) GetPersonnes() []CartographiePersonnesInner`

GetPersonnes returns the Personnes field if non-nil, zero value otherwise.

### GetPersonnesOk

`func (o *Cartographie) GetPersonnesOk() (*[]CartographiePersonnesInner, bool)`

GetPersonnesOk returns a tuple with the Personnes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonnes

`func (o *Cartographie) SetPersonnes(v []CartographiePersonnesInner)`

SetPersonnes sets Personnes field to given value.

### HasPersonnes

`func (o *Cartographie) HasPersonnes() bool`

HasPersonnes returns a boolean if a field has been set.

### GetLiensEntreprisesPersonnes

`func (o *Cartographie) GetLiensEntreprisesPersonnes() [][]CartographieLiensEntreprisesPersonnesInnerInner`

GetLiensEntreprisesPersonnes returns the LiensEntreprisesPersonnes field if non-nil, zero value otherwise.

### GetLiensEntreprisesPersonnesOk

`func (o *Cartographie) GetLiensEntreprisesPersonnesOk() (*[][]CartographieLiensEntreprisesPersonnesInnerInner, bool)`

GetLiensEntreprisesPersonnesOk returns a tuple with the LiensEntreprisesPersonnes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiensEntreprisesPersonnes

`func (o *Cartographie) SetLiensEntreprisesPersonnes(v [][]CartographieLiensEntreprisesPersonnesInnerInner)`

SetLiensEntreprisesPersonnes sets LiensEntreprisesPersonnes field to given value.

### HasLiensEntreprisesPersonnes

`func (o *Cartographie) HasLiensEntreprisesPersonnes() bool`

HasLiensEntreprisesPersonnes returns a boolean if a field has been set.

### GetLiensEntreprisesEntreprises

`func (o *Cartographie) GetLiensEntreprisesEntreprises() [][]string`

GetLiensEntreprisesEntreprises returns the LiensEntreprisesEntreprises field if non-nil, zero value otherwise.

### GetLiensEntreprisesEntreprisesOk

`func (o *Cartographie) GetLiensEntreprisesEntreprisesOk() (*[][]string, bool)`

GetLiensEntreprisesEntreprisesOk returns a tuple with the LiensEntreprisesEntreprises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiensEntreprisesEntreprises

`func (o *Cartographie) SetLiensEntreprisesEntreprises(v [][]string)`

SetLiensEntreprisesEntreprises sets LiensEntreprisesEntreprises field to given value.

### HasLiensEntreprisesEntreprises

`func (o *Cartographie) HasLiensEntreprisesEntreprises() bool`

HasLiensEntreprisesEntreprises returns a boolean if a field has been set.

### GetModificationsEffectuees

`func (o *Cartographie) GetModificationsEffectuees() map[string]interface{}`

GetModificationsEffectuees returns the ModificationsEffectuees field if non-nil, zero value otherwise.

### GetModificationsEffectueesOk

`func (o *Cartographie) GetModificationsEffectueesOk() (*map[string]interface{}, bool)`

GetModificationsEffectueesOk returns a tuple with the ModificationsEffectuees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationsEffectuees

`func (o *Cartographie) SetModificationsEffectuees(v map[string]interface{})`

SetModificationsEffectuees sets ModificationsEffectuees field to given value.

### HasModificationsEffectuees

`func (o *Cartographie) HasModificationsEffectuees() bool`

HasModificationsEffectuees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


