# SurveillanceDirigeantRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siren** | **string** | SIREN de la personne morale (si ajout d&#39;une personne morale) | 
**Denomination** | Pointer to **string** | Dénomination de la personne morale (si ajout d&#39;une personne morale) | [optional] 
**Nom** | Pointer to **string** | Nom de la personne physique (si ajout d&#39;une personne physique) | [optional] 
**Prenom** | Pointer to **string** | Prénom de la personne physique (si ajout d&#39;une personne physique) | [optional] 
**DateDeNaissance** | Pointer to **string** | Date de naissance de la personne physique au format AAAA-MM-JJ (si ajout d&#39;une personne physique) | [optional] 
**RechercheElargie** | Pointer to **bool** | Active la recherche élargie du dirigeant : &lt;ul&gt;&lt;li&gt;Pour une personne physique : c&#39;est un cas rare mais la date de naissance du dirigeant n&#39;est pas toujours connue de Pappers. Si vous activez ce filtre et que la date de naissance est inconnue, l&#39;alerte sera générée uniquement sur la base du nom et prénom.&lt;/li&gt;&lt;li&gt;Pour une personne morale : c&#39;est un cas rare mais le SIREN du dirigeant n&#39;est pas toujours connu de Pappers. Si vous activez ce filtre et que le SIREN est inconnu, l&#39;alerte sera générée uniquement sur la base de la dénomination. Activable uniquement si le champ denomination est présent.&lt;/li&gt;&lt;/ul&gt; | [optional] [default to false]

## Methods

### NewSurveillanceDirigeantRequestInner

`func NewSurveillanceDirigeantRequestInner(siren string, ) *SurveillanceDirigeantRequestInner`

NewSurveillanceDirigeantRequestInner instantiates a new SurveillanceDirigeantRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveillanceDirigeantRequestInnerWithDefaults

`func NewSurveillanceDirigeantRequestInnerWithDefaults() *SurveillanceDirigeantRequestInner`

NewSurveillanceDirigeantRequestInnerWithDefaults instantiates a new SurveillanceDirigeantRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiren

`func (o *SurveillanceDirigeantRequestInner) GetSiren() string`

GetSiren returns the Siren field if non-nil, zero value otherwise.

### GetSirenOk

`func (o *SurveillanceDirigeantRequestInner) GetSirenOk() (*string, bool)`

GetSirenOk returns a tuple with the Siren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiren

`func (o *SurveillanceDirigeantRequestInner) SetSiren(v string)`

SetSiren sets Siren field to given value.


### GetDenomination

`func (o *SurveillanceDirigeantRequestInner) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *SurveillanceDirigeantRequestInner) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *SurveillanceDirigeantRequestInner) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *SurveillanceDirigeantRequestInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetNom

`func (o *SurveillanceDirigeantRequestInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *SurveillanceDirigeantRequestInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *SurveillanceDirigeantRequestInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *SurveillanceDirigeantRequestInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetPrenom

`func (o *SurveillanceDirigeantRequestInner) GetPrenom() string`

GetPrenom returns the Prenom field if non-nil, zero value otherwise.

### GetPrenomOk

`func (o *SurveillanceDirigeantRequestInner) GetPrenomOk() (*string, bool)`

GetPrenomOk returns a tuple with the Prenom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrenom

`func (o *SurveillanceDirigeantRequestInner) SetPrenom(v string)`

SetPrenom sets Prenom field to given value.

### HasPrenom

`func (o *SurveillanceDirigeantRequestInner) HasPrenom() bool`

HasPrenom returns a boolean if a field has been set.

### GetDateDeNaissance

`func (o *SurveillanceDirigeantRequestInner) GetDateDeNaissance() string`

GetDateDeNaissance returns the DateDeNaissance field if non-nil, zero value otherwise.

### GetDateDeNaissanceOk

`func (o *SurveillanceDirigeantRequestInner) GetDateDeNaissanceOk() (*string, bool)`

GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeNaissance

`func (o *SurveillanceDirigeantRequestInner) SetDateDeNaissance(v string)`

SetDateDeNaissance sets DateDeNaissance field to given value.

### HasDateDeNaissance

`func (o *SurveillanceDirigeantRequestInner) HasDateDeNaissance() bool`

HasDateDeNaissance returns a boolean if a field has been set.

### GetRechercheElargie

`func (o *SurveillanceDirigeantRequestInner) GetRechercheElargie() bool`

GetRechercheElargie returns the RechercheElargie field if non-nil, zero value otherwise.

### GetRechercheElargieOk

`func (o *SurveillanceDirigeantRequestInner) GetRechercheElargieOk() (*bool, bool)`

GetRechercheElargieOk returns a tuple with the RechercheElargie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechercheElargie

`func (o *SurveillanceDirigeantRequestInner) SetRechercheElargie(v bool)`

SetRechercheElargie sets RechercheElargie field to given value.

### HasRechercheElargie

`func (o *SurveillanceDirigeantRequestInner) HasRechercheElargie() bool`

HasRechercheElargie returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


