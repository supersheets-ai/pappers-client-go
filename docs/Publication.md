# Publication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type de publication | [optional] 
**Date** | Pointer to **string** | Date de publication, au format AAAA-MM-JJ. | [optional] 
**Contenu** | Pointer to **string** | Contenu de la publication, avec les mentions correspondant à la recherche encerclée par une balise HTML &lt;em&gt;. | [optional] 

## Methods

### NewPublication

`func NewPublication() *Publication`

NewPublication instantiates a new Publication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicationWithDefaults

`func NewPublicationWithDefaults() *Publication`

NewPublicationWithDefaults instantiates a new Publication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Publication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Publication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Publication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Publication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDate

`func (o *Publication) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Publication) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Publication) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Publication) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetContenu

`func (o *Publication) GetContenu() string`

GetContenu returns the Contenu field if non-nil, zero value otherwise.

### GetContenuOk

`func (o *Publication) GetContenuOk() (*string, bool)`

GetContenuOk returns a tuple with the Contenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContenu

`func (o *Publication) SetContenu(v string)`

SetContenu sets Contenu field to given value.

### HasContenu

`func (o *Publication) HasContenu() bool`

HasContenu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


