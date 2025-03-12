# ScoringNonFinancier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** | Note du score non financier de l&#39;entreprise. | [optional] 
**Score** | Pointer to **int32** | Score non financier de l&#39;entreprise, entre 0 et 20. | [optional] 
**DateCalcul** | Pointer to **string** | Date de calcul du score non financier de l&#39;entreprise. | [optional] 
**Erreur** | Pointer to **string** | Erreur lors du calcul du score non financier de l&#39;entreprise. | [optional] 

## Methods

### NewScoringNonFinancier

`func NewScoringNonFinancier() *ScoringNonFinancier`

NewScoringNonFinancier instantiates a new ScoringNonFinancier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoringNonFinancierWithDefaults

`func NewScoringNonFinancierWithDefaults() *ScoringNonFinancier`

NewScoringNonFinancierWithDefaults instantiates a new ScoringNonFinancier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *ScoringNonFinancier) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ScoringNonFinancier) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ScoringNonFinancier) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ScoringNonFinancier) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetScore

`func (o *ScoringNonFinancier) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScoringNonFinancier) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScoringNonFinancier) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScoringNonFinancier) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDateCalcul

`func (o *ScoringNonFinancier) GetDateCalcul() string`

GetDateCalcul returns the DateCalcul field if non-nil, zero value otherwise.

### GetDateCalculOk

`func (o *ScoringNonFinancier) GetDateCalculOk() (*string, bool)`

GetDateCalculOk returns a tuple with the DateCalcul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCalcul

`func (o *ScoringNonFinancier) SetDateCalcul(v string)`

SetDateCalcul sets DateCalcul field to given value.

### HasDateCalcul

`func (o *ScoringNonFinancier) HasDateCalcul() bool`

HasDateCalcul returns a boolean if a field has been set.

### GetErreur

`func (o *ScoringNonFinancier) GetErreur() string`

GetErreur returns the Erreur field if non-nil, zero value otherwise.

### GetErreurOk

`func (o *ScoringNonFinancier) GetErreurOk() (*string, bool)`

GetErreurOk returns a tuple with the Erreur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErreur

`func (o *ScoringNonFinancier) SetErreur(v string)`

SetErreur sets Erreur field to given value.

### HasErreur

`func (o *ScoringNonFinancier) HasErreur() bool`

HasErreur returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


