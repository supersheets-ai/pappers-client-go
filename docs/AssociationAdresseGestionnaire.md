# AssociationAdresseGestionnaire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gestionnaire** | Pointer to **string** | Nom du gestionnaire de l&#39;association. | [optional] 
**CodePostal** | Pointer to **string** | Code postal du gestionnaire. | [optional] 
**Ville** | Pointer to **string** | Ville du gestionnaire. | [optional] 
**Distribution** | Pointer to **string** | Complément de distribution du gestionnaire. | [optional] 
**AdresseLigne** | Pointer to **string** | Adresse complète du gestionnaire. | [optional] 
**ComplementAdresse** | Pointer to **string** | Complément de l&#39;adresse du gestionnaire. | [optional] 
**Indication** | Pointer to **string** | Indication supplémentaire à l&#39;adresse du gestionnaire. | [optional] 
**Pays** | Pointer to **string** | Pays du gestionnaire. | [optional] 

## Methods

### NewAssociationAdresseGestionnaire

`func NewAssociationAdresseGestionnaire() *AssociationAdresseGestionnaire`

NewAssociationAdresseGestionnaire instantiates a new AssociationAdresseGestionnaire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationAdresseGestionnaireWithDefaults

`func NewAssociationAdresseGestionnaireWithDefaults() *AssociationAdresseGestionnaire`

NewAssociationAdresseGestionnaireWithDefaults instantiates a new AssociationAdresseGestionnaire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGestionnaire

`func (o *AssociationAdresseGestionnaire) GetGestionnaire() string`

GetGestionnaire returns the Gestionnaire field if non-nil, zero value otherwise.

### GetGestionnaireOk

`func (o *AssociationAdresseGestionnaire) GetGestionnaireOk() (*string, bool)`

GetGestionnaireOk returns a tuple with the Gestionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGestionnaire

`func (o *AssociationAdresseGestionnaire) SetGestionnaire(v string)`

SetGestionnaire sets Gestionnaire field to given value.

### HasGestionnaire

`func (o *AssociationAdresseGestionnaire) HasGestionnaire() bool`

HasGestionnaire returns a boolean if a field has been set.

### GetCodePostal

`func (o *AssociationAdresseGestionnaire) GetCodePostal() string`

GetCodePostal returns the CodePostal field if non-nil, zero value otherwise.

### GetCodePostalOk

`func (o *AssociationAdresseGestionnaire) GetCodePostalOk() (*string, bool)`

GetCodePostalOk returns a tuple with the CodePostal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePostal

`func (o *AssociationAdresseGestionnaire) SetCodePostal(v string)`

SetCodePostal sets CodePostal field to given value.

### HasCodePostal

`func (o *AssociationAdresseGestionnaire) HasCodePostal() bool`

HasCodePostal returns a boolean if a field has been set.

### GetVille

`func (o *AssociationAdresseGestionnaire) GetVille() string`

GetVille returns the Ville field if non-nil, zero value otherwise.

### GetVilleOk

`func (o *AssociationAdresseGestionnaire) GetVilleOk() (*string, bool)`

GetVilleOk returns a tuple with the Ville field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVille

`func (o *AssociationAdresseGestionnaire) SetVille(v string)`

SetVille sets Ville field to given value.

### HasVille

`func (o *AssociationAdresseGestionnaire) HasVille() bool`

HasVille returns a boolean if a field has been set.

### GetDistribution

`func (o *AssociationAdresseGestionnaire) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *AssociationAdresseGestionnaire) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *AssociationAdresseGestionnaire) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *AssociationAdresseGestionnaire) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetAdresseLigne

`func (o *AssociationAdresseGestionnaire) GetAdresseLigne() string`

GetAdresseLigne returns the AdresseLigne field if non-nil, zero value otherwise.

### GetAdresseLigneOk

`func (o *AssociationAdresseGestionnaire) GetAdresseLigneOk() (*string, bool)`

GetAdresseLigneOk returns a tuple with the AdresseLigne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseLigne

`func (o *AssociationAdresseGestionnaire) SetAdresseLigne(v string)`

SetAdresseLigne sets AdresseLigne field to given value.

### HasAdresseLigne

`func (o *AssociationAdresseGestionnaire) HasAdresseLigne() bool`

HasAdresseLigne returns a boolean if a field has been set.

### GetComplementAdresse

`func (o *AssociationAdresseGestionnaire) GetComplementAdresse() string`

GetComplementAdresse returns the ComplementAdresse field if non-nil, zero value otherwise.

### GetComplementAdresseOk

`func (o *AssociationAdresseGestionnaire) GetComplementAdresseOk() (*string, bool)`

GetComplementAdresseOk returns a tuple with the ComplementAdresse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementAdresse

`func (o *AssociationAdresseGestionnaire) SetComplementAdresse(v string)`

SetComplementAdresse sets ComplementAdresse field to given value.

### HasComplementAdresse

`func (o *AssociationAdresseGestionnaire) HasComplementAdresse() bool`

HasComplementAdresse returns a boolean if a field has been set.

### GetIndication

`func (o *AssociationAdresseGestionnaire) GetIndication() string`

GetIndication returns the Indication field if non-nil, zero value otherwise.

### GetIndicationOk

`func (o *AssociationAdresseGestionnaire) GetIndicationOk() (*string, bool)`

GetIndicationOk returns a tuple with the Indication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndication

`func (o *AssociationAdresseGestionnaire) SetIndication(v string)`

SetIndication sets Indication field to given value.

### HasIndication

`func (o *AssociationAdresseGestionnaire) HasIndication() bool`

HasIndication returns a boolean if a field has been set.

### GetPays

`func (o *AssociationAdresseGestionnaire) GetPays() string`

GetPays returns the Pays field if non-nil, zero value otherwise.

### GetPaysOk

`func (o *AssociationAdresseGestionnaire) GetPaysOk() (*string, bool)`

GetPaysOk returns a tuple with the Pays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPays

`func (o *AssociationAdresseGestionnaire) SetPays(v string)`

SetPays sets Pays field to given value.

### HasPays

`func (o *AssociationAdresseGestionnaire) HasPays() bool`

HasPays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


