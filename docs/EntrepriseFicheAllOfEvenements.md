# EntrepriseFicheAllOfEvenements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type d&#39;événement. | [optional] 
**IdentifiantEvenement** | Pointer to **string** | Identifiant de l&#39;événement. | [optional] 
**Reference** | Pointer to **string** | Référence de l&#39;événement. | [optional] 
**Date** | Pointer to **string** | Date de l&#39;événement, au format AAAA-MM-JJ. | [optional] 
**NumeroBopi** | Pointer to **string** | Numéro du BOPI dans lequel l&#39;événement a été publié. | [optional] 
**DateBopi** | Pointer to **string** | Date de publication du BOPI au format AAAA-MM-JJ. | [optional] 
**Beneficiaire** | Pointer to **string** | Bénéficiaire associé à l&#39;événement. | [optional] 

## Methods

### NewEntrepriseFicheAllOfEvenements

`func NewEntrepriseFicheAllOfEvenements() *EntrepriseFicheAllOfEvenements`

NewEntrepriseFicheAllOfEvenements instantiates a new EntrepriseFicheAllOfEvenements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfEvenementsWithDefaults

`func NewEntrepriseFicheAllOfEvenementsWithDefaults() *EntrepriseFicheAllOfEvenements`

NewEntrepriseFicheAllOfEvenementsWithDefaults instantiates a new EntrepriseFicheAllOfEvenements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntrepriseFicheAllOfEvenements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrepriseFicheAllOfEvenements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrepriseFicheAllOfEvenements) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntrepriseFicheAllOfEvenements) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdentifiantEvenement

`func (o *EntrepriseFicheAllOfEvenements) GetIdentifiantEvenement() string`

GetIdentifiantEvenement returns the IdentifiantEvenement field if non-nil, zero value otherwise.

### GetIdentifiantEvenementOk

`func (o *EntrepriseFicheAllOfEvenements) GetIdentifiantEvenementOk() (*string, bool)`

GetIdentifiantEvenementOk returns a tuple with the IdentifiantEvenement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiantEvenement

`func (o *EntrepriseFicheAllOfEvenements) SetIdentifiantEvenement(v string)`

SetIdentifiantEvenement sets IdentifiantEvenement field to given value.

### HasIdentifiantEvenement

`func (o *EntrepriseFicheAllOfEvenements) HasIdentifiantEvenement() bool`

HasIdentifiantEvenement returns a boolean if a field has been set.

### GetReference

`func (o *EntrepriseFicheAllOfEvenements) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EntrepriseFicheAllOfEvenements) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EntrepriseFicheAllOfEvenements) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EntrepriseFicheAllOfEvenements) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDate

`func (o *EntrepriseFicheAllOfEvenements) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EntrepriseFicheAllOfEvenements) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EntrepriseFicheAllOfEvenements) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EntrepriseFicheAllOfEvenements) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNumeroBopi

`func (o *EntrepriseFicheAllOfEvenements) GetNumeroBopi() string`

GetNumeroBopi returns the NumeroBopi field if non-nil, zero value otherwise.

### GetNumeroBopiOk

`func (o *EntrepriseFicheAllOfEvenements) GetNumeroBopiOk() (*string, bool)`

GetNumeroBopiOk returns a tuple with the NumeroBopi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroBopi

`func (o *EntrepriseFicheAllOfEvenements) SetNumeroBopi(v string)`

SetNumeroBopi sets NumeroBopi field to given value.

### HasNumeroBopi

`func (o *EntrepriseFicheAllOfEvenements) HasNumeroBopi() bool`

HasNumeroBopi returns a boolean if a field has been set.

### GetDateBopi

`func (o *EntrepriseFicheAllOfEvenements) GetDateBopi() string`

GetDateBopi returns the DateBopi field if non-nil, zero value otherwise.

### GetDateBopiOk

`func (o *EntrepriseFicheAllOfEvenements) GetDateBopiOk() (*string, bool)`

GetDateBopiOk returns a tuple with the DateBopi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBopi

`func (o *EntrepriseFicheAllOfEvenements) SetDateBopi(v string)`

SetDateBopi sets DateBopi field to given value.

### HasDateBopi

`func (o *EntrepriseFicheAllOfEvenements) HasDateBopi() bool`

HasDateBopi returns a boolean if a field has been set.

### GetBeneficiaire

`func (o *EntrepriseFicheAllOfEvenements) GetBeneficiaire() string`

GetBeneficiaire returns the Beneficiaire field if non-nil, zero value otherwise.

### GetBeneficiaireOk

`func (o *EntrepriseFicheAllOfEvenements) GetBeneficiaireOk() (*string, bool)`

GetBeneficiaireOk returns a tuple with the Beneficiaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaire

`func (o *EntrepriseFicheAllOfEvenements) SetBeneficiaire(v string)`

SetBeneficiaire sets Beneficiaire field to given value.

### HasBeneficiaire

`func (o *EntrepriseFicheAllOfEvenements) HasBeneficiaire() bool`

HasBeneficiaire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


