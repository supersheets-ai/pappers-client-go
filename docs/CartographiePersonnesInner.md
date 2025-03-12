# CartographiePersonnesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Un identifiant unique du noeud. | [optional] 
**Prenom** | Pointer to **string** | SIREN de l&#39;entreprise. | [optional] 
**Nom** | Pointer to **string** | Nom de l&#39;entreprise. | [optional] 
**Niveau** | Pointer to **int32** | Niveau du noeud. Le niveau 1 correspond aux dirigeants et bénéficiaires effectifs directement liés à l&#39;entreprise recherchée. Le niveau 2 correspond aux autres. | [optional] 

## Methods

### NewCartographiePersonnesInner

`func NewCartographiePersonnesInner() *CartographiePersonnesInner`

NewCartographiePersonnesInner instantiates a new CartographiePersonnesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartographiePersonnesInnerWithDefaults

`func NewCartographiePersonnesInnerWithDefaults() *CartographiePersonnesInner`

NewCartographiePersonnesInnerWithDefaults instantiates a new CartographiePersonnesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartographiePersonnesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartographiePersonnesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartographiePersonnesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartographiePersonnesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrenom

`func (o *CartographiePersonnesInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *CartographiePersonnesInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *CartographiePersonnesInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *CartographiePersonnesInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetNom

`func (o *CartographiePersonnesInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *CartographiePersonnesInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *CartographiePersonnesInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *CartographiePersonnesInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetNiveau

`func (o *CartographiePersonnesInner) GetNiveau() int32`

GetNiveau returns the Niveau field if non-nil, zero value otherwise.

### GetNiveauOk

`func (o *CartographiePersonnesInner) GetNiveauOk() (*int32, bool)`

GetNiveauOk returns a tuple with the Niveau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiveau

`func (o *CartographiePersonnesInner) SetNiveau(v int32)`

SetNiveau sets Niveau field to given value.

### HasNiveau

`func (o *CartographiePersonnesInner) HasNiveau() bool`

HasNiveau returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


