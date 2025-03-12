# Suggestions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultatsNomEntreprise** | Pointer to [**[]Suggestions200ResponseResultatsNomEntrepriseInner**](Suggestions200ResponseResultatsNomEntrepriseInner.md) | Liste des entreprises dont le nom (dénomination ou nom/prénom) peut compléter la recherche textuelle. Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;nom_entreprise&#x60;. | [optional] 
**ResultatsDenomination** | Pointer to [**[]Suggestions200ResponseResultatsDenominationInner**](Suggestions200ResponseResultatsDenominationInner.md) | Liste des entreprises dont la dénomination peut compléter la recherche textuelle (concerne les personnes morales seulement). Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;denomination&#x60;. | [optional] 
**ResultatsNomComplet** | Pointer to [**[]Suggestions200ResponseResultatsNomCompletInner**](Suggestions200ResponseResultatsNomCompletInner.md) | Liste des entreprises dont le nom complet (nom + prénom ou prénom + nom) peut compléter la recherche textuelle (concerne les personnes physiques seulement). Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;nom_complet&#x60;. | [optional] 
**ResultatsRepresentant** | Pointer to [**[]Suggestions200ResponseResultatsRepresentantInner**](Suggestions200ResponseResultatsRepresentantInner.md) | Liste des représentants dont le nom complet (nom + prénom ou prénom + nom) peut compléter la recherche textuelle. Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;representant&#x60;. | [optional] 
**ResultatsSiren** | Pointer to [**[]Suggestions200ResponseResultatsSirenInner**](Suggestions200ResponseResultatsSirenInner.md) | Liste des entreprises dont le numéro SIREN peut compléter la recherche textuelle. Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;siren&#x60;. | [optional] 
**ResultatsSiret** | Pointer to [**[]Suggestions200ResponseResultatsSiretInner**](Suggestions200ResponseResultatsSiretInner.md) | Liste des entreprises dont le numéro SIRET peut compléter la recherche textuelle. Uniquement présent si le paramètre &#x60;cibles&#x60; contient &#x60;siret&#x60;. | [optional] 

## Methods

### NewSuggestions200Response

`func NewSuggestions200Response() *Suggestions200Response`

NewSuggestions200Response instantiates a new Suggestions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestions200ResponseWithDefaults

`func NewSuggestions200ResponseWithDefaults() *Suggestions200Response`

NewSuggestions200ResponseWithDefaults instantiates a new Suggestions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultatsNomEntreprise

`func (o *Suggestions200Response) GetResultatsNomEntreprise() []Suggestions200ResponseResultatsNomEntrepriseInner`

GetResultatsNomEntreprise returns the ResultatsNomEntreprise field if non-nil, zero value otherwise.

### GetResultatsNomEntrepriseOk

`func (o *Suggestions200Response) GetResultatsNomEntrepriseOk() (*[]Suggestions200ResponseResultatsNomEntrepriseInner, bool)`

GetResultatsNomEntrepriseOk returns a tuple with the ResultatsNomEntreprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsNomEntreprise

`func (o *Suggestions200Response) SetResultatsNomEntreprise(v []Suggestions200ResponseResultatsNomEntrepriseInner)`

SetResultatsNomEntreprise sets ResultatsNomEntreprise field to given value.

### HasResultatsNomEntreprise

`func (o *Suggestions200Response) HasResultatsNomEntreprise() bool`

HasResultatsNomEntreprise returns a boolean if a field has been set.

### GetResultatsDenomination

`func (o *Suggestions200Response) GetResultatsDenomination() []Suggestions200ResponseResultatsDenominationInner`

GetResultatsDenomination returns the ResultatsDenomination field if non-nil, zero value otherwise.

### GetResultatsDenominationOk

`func (o *Suggestions200Response) GetResultatsDenominationOk() (*[]Suggestions200ResponseResultatsDenominationInner, bool)`

GetResultatsDenominationOk returns a tuple with the ResultatsDenomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsDenomination

`func (o *Suggestions200Response) SetResultatsDenomination(v []Suggestions200ResponseResultatsDenominationInner)`

SetResultatsDenomination sets ResultatsDenomination field to given value.

### HasResultatsDenomination

`func (o *Suggestions200Response) HasResultatsDenomination() bool`

HasResultatsDenomination returns a boolean if a field has been set.

### GetResultatsNomComplet

`func (o *Suggestions200Response) GetResultatsNomComplet() []Suggestions200ResponseResultatsNomCompletInner`

GetResultatsNomComplet returns the ResultatsNomComplet field if non-nil, zero value otherwise.

### GetResultatsNomCompletOk

`func (o *Suggestions200Response) GetResultatsNomCompletOk() (*[]Suggestions200ResponseResultatsNomCompletInner, bool)`

GetResultatsNomCompletOk returns a tuple with the ResultatsNomComplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsNomComplet

`func (o *Suggestions200Response) SetResultatsNomComplet(v []Suggestions200ResponseResultatsNomCompletInner)`

SetResultatsNomComplet sets ResultatsNomComplet field to given value.

### HasResultatsNomComplet

`func (o *Suggestions200Response) HasResultatsNomComplet() bool`

HasResultatsNomComplet returns a boolean if a field has been set.

### GetResultatsRepresentant

`func (o *Suggestions200Response) GetResultatsRepresentant() []Suggestions200ResponseResultatsRepresentantInner`

GetResultatsRepresentant returns the ResultatsRepresentant field if non-nil, zero value otherwise.

### GetResultatsRepresentantOk

`func (o *Suggestions200Response) GetResultatsRepresentantOk() (*[]Suggestions200ResponseResultatsRepresentantInner, bool)`

GetResultatsRepresentantOk returns a tuple with the ResultatsRepresentant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsRepresentant

`func (o *Suggestions200Response) SetResultatsRepresentant(v []Suggestions200ResponseResultatsRepresentantInner)`

SetResultatsRepresentant sets ResultatsRepresentant field to given value.

### HasResultatsRepresentant

`func (o *Suggestions200Response) HasResultatsRepresentant() bool`

HasResultatsRepresentant returns a boolean if a field has been set.

### GetResultatsSiren

`func (o *Suggestions200Response) GetResultatsSiren() []Suggestions200ResponseResultatsSirenInner`

GetResultatsSiren returns the ResultatsSiren field if non-nil, zero value otherwise.

### GetResultatsSirenOk

`func (o *Suggestions200Response) GetResultatsSirenOk() (*[]Suggestions200ResponseResultatsSirenInner, bool)`

GetResultatsSirenOk returns a tuple with the ResultatsSiren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsSiren

`func (o *Suggestions200Response) SetResultatsSiren(v []Suggestions200ResponseResultatsSirenInner)`

SetResultatsSiren sets ResultatsSiren field to given value.

### HasResultatsSiren

`func (o *Suggestions200Response) HasResultatsSiren() bool`

HasResultatsSiren returns a boolean if a field has been set.

### GetResultatsSiret

`func (o *Suggestions200Response) GetResultatsSiret() []Suggestions200ResponseResultatsSiretInner`

GetResultatsSiret returns the ResultatsSiret field if non-nil, zero value otherwise.

### GetResultatsSiretOk

`func (o *Suggestions200Response) GetResultatsSiretOk() (*[]Suggestions200ResponseResultatsSiretInner, bool)`

GetResultatsSiretOk returns a tuple with the ResultatsSiret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatsSiret

`func (o *Suggestions200Response) SetResultatsSiret(v []Suggestions200ResponseResultatsSiretInner)`

SetResultatsSiret sets ResultatsSiret field to given value.

### HasResultatsSiret

`func (o *Suggestions200Response) HasResultatsSiret() bool`

HasResultatsSiret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


