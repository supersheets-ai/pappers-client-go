# EtablissementRecherche

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Siret** | Pointer to **string** | Numéro siret de l&#39;établissement au format xxxxxxxxxxxxxx. | [optional] 
**SiretFormate** | Pointer to **string** | Numéro siret de l&#39;établissement au format xxx xxx xxx xxxxx. | [optional] 
**Nic** | Pointer to **string** | Numéro NIC de l&#39;établissement. | [optional] 
**NumeroVoie** | Pointer to **int32** | Numéro de voie de l&#39;établissement. | [optional] 
**IndiceRepetition** | Pointer to **string** | Indice de répétition de l&#39;établissement. | [optional] 
**TypeVoie** | Pointer to **string** | Type de voie de l&#39;établissement. | [optional] 
**LibelleVoie** | Pointer to **string** | Libellé de la voie de l&#39;établissement. | [optional] 
**ComplementAdresse** | Pointer to **string** | Complément d&#39;adresse de l&#39;établissement. | [optional] 
**AdresseLigne1** | Pointer to **string** | Première ligne de l&#39;adresse de l&#39;établissement. Correspond à l&#39;ensemble des données numero_voie, indice_repetition, type_voie et libelle_voie. | [optional] 
**AdresseLigne2** | Pointer to **string** | Seconde ligne de l&#39;adresse de l&#39;établissement. Est égal à complement_adresse | [optional] 
**CodePostal** | Pointer to **string** | Code postal de l&#39;établissement. | [optional] 
**Ville** | Pointer to **string** | Ville de l&#39;établissement. | [optional] 
**Latitude** | Pointer to **float32** | Latitude de l&#39;établissement. | [optional] 
**Longitude** | Pointer to **float32** | Longitude de l&#39;établissement. | [optional] 

## Methods

### NewEtablissementRecherche

`func NewEtablissementRecherche() *EtablissementRecherche`

NewEtablissementRecherche instantiates a new EtablissementRecherche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtablissementRechercheWithDefaults

`func NewEtablissementRechercheWithDefaults() *EtablissementRecherche`

NewEtablissementRechercheWithDefaults instantiates a new EtablissementRecherche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiret

`func (o *EtablissementRecherche) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *EtablissementRecherche) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *EtablissementRecherche) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *EtablissementRecherche) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetSiretFormate

`func (o *EtablissementRecherche) GetSiretFormate() string`

GetSiretFormate returns the SiretFormate field if non-nil, zero value otherwise.

### GetSiretFormateOk

`func (o *EtablissementRecherche) GetSiretFormateOk() (*string, bool)`

GetSiretFormateOk returns a tuple with the SiretFormate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiretFormate

`func (o *EtablissementRecherche) SetSiretFormate(v string)`

SetSiretFormate sets SiretFormate field to given value.

### HasSiretFormate

`func (o *EtablissementRecherche) HasSiretFormate() bool`

HasSiretFormate returns a boolean if a field has been set.

### GetNic

`func (o *EtablissementRecherche) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *EtablissementRecherche) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *EtablissementRecherche) SetNic(v string)`

SetNic sets Nic field to given value.

### HasNic

`func (o *EtablissementRecherche) HasNic() bool`

HasNic returns a boolean if a field has been set.

### GetNumeroVoie

`func (o *EtablissementRecherche) GetNumeroVoie() int32`

GetNumeroVoie returns the NumeroVoie field if non-nil, zero value otherwise.

### GetNumeroVoieOk

`func (o *EtablissementRecherche) GetNumeroVoieOk() (*int32, bool)`

GetNumeroVoieOk returns a tuple with the NumeroVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroVoie

`func (o *EtablissementRecherche) SetNumeroVoie(v int32)`

SetNumeroVoie sets NumeroVoie field to given value.

### HasNumeroVoie

`func (o *EtablissementRecherche) HasNumeroVoie() bool`

HasNumeroVoie returns a boolean if a field has been set.

### GetIndiceRepetition

`func (o *EtablissementRecherche) GetIndiceRepetition() string`

GetIndiceRepetition returns the IndiceRepetition field if non-nil, zero value otherwise.

### GetIndiceRepetitionOk

`func (o *EtablissementRecherche) GetIndiceRepetitionOk() (*string, bool)`

GetIndiceRepetitionOk returns a tuple with the IndiceRepetition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndiceRepetition

`func (o *EtablissementRecherche) SetIndiceRepetition(v string)`

SetIndiceRepetition sets IndiceRepetition field to given value.

### HasIndiceRepetition

`func (o *EtablissementRecherche) HasIndiceRepetition() bool`

HasIndiceRepetition returns a boolean if a field has been set.

### GetTypeVoie

`func (o *EtablissementRecherche) GetTypeVoie() string`

GetTypeVoie returns the TypeVoie field if non-nil, zero value otherwise.

### GetTypeVoieOk

`func (o *EtablissementRecherche) GetTypeVoieOk() (*string, bool)`

GetTypeVoieOk returns a tuple with the TypeVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVoie

`func (o *EtablissementRecherche) SetTypeVoie(v string)`

SetTypeVoie sets TypeVoie field to given value.

### HasTypeVoie

`func (o *EtablissementRecherche) HasTypeVoie() bool`

HasTypeVoie returns a boolean if a field has been set.

### GetLibelleVoie

`func (o *EtablissementRecherche) GetLibelleVoie() string`

GetLibelleVoie returns the LibelleVoie field if non-nil, zero value otherwise.

### GetLibelleVoieOk

`func (o *EtablissementRecherche) GetLibelleVoieOk() (*string, bool)`

GetLibelleVoieOk returns a tuple with the LibelleVoie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibelleVoie

`func (o *EtablissementRecherche) SetLibelleVoie(v string)`

SetLibelleVoie sets LibelleVoie field to given value.

### HasLibelleVoie

`func (o *EtablissementRecherche) HasLibelleVoie() bool`

HasLibelleVoie returns a boolean if a field has been set.

### GetComplementAdresse

`func (o *EtablissementRecherche) GetComplementAdresse() string`

GetComplementAdresse returns the ComplementAdresse field if non-nil, zero value otherwise.

### GetComplementAdresseOk

`func (o *EtablissementRecherche) GetComplementAdresseOk() (*string, bool)`

GetComplementAdresseOk returns a tuple with the ComplementAdresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementAdresse

`func (o *EtablissementRecherche) SetComplementAdresse(v string)`

SetComplementAdresse sets ComplementAdresse field to given value.

### HasComplementAdresse

`func (o *EtablissementRecherche) HasComplementAdresse() bool`

HasComplementAdresse returns a boolean if a field has been set.

### GetAdresseLigne1

`func (o *EtablissementRecherche) GetAdresseLigne1() string`

GetAdresseLigne1 returns the AdresseLigne1 field if non-nil, zero value otherwise.

### GetAdresseLigne1Ok

`func (o *EtablissementRecherche) GetAdresseLigne1Ok() (*string, bool)`

GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne1

`func (o *EtablissementRecherche) SetAdresseLigne1(v string)`

SetAdresseLigne1 sets AdresseLigne1 field to given value.

### HasAdresseLigne1

`func (o *EtablissementRecherche) HasAdresseLigne1() bool`

HasAdresseLigne1 returns a boolean if a field has been set.

### GetAdresseLigne2

`func (o *EtablissementRecherche) GetAdresseLigne2() string`

GetAdresseLigne2 returns the AdresseLigne2 field if non-nil, zero value otherwise.

### GetAdresseLigne2Ok

`func (o *EtablissementRecherche) GetAdresseLigne2Ok() (*string, bool)`

GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne2

`func (o *EtablissementRecherche) SetAdresseLigne2(v string)`

SetAdresseLigne2 sets AdresseLigne2 field to given value.

### HasAdresseLigne2

`func (o *EtablissementRecherche) HasAdresseLigne2() bool`

HasAdresseLigne2 returns a boolean if a field has been set.

### GetCodePostal

`func (o *EtablissementRecherche) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *EtablissementRecherche) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *EtablissementRecherche) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *EtablissementRecherche) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *EtablissementRecherche) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *EtablissementRecherche) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *EtablissementRecherche) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *EtablissementRecherche) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetLatitude

`func (o *EtablissementRecherche) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *EtablissementRecherche) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *EtablissementRecherche) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *EtablissementRecherche) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *EtablissementRecherche) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *EtablissementRecherche) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *EtablissementRecherche) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *EtablissementRecherche) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


