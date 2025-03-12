# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type de document | [optional] 
**Token** | Pointer to **string** | Token du document | [optional] 
**DateDepot** | Pointer to **string** | Date de dépôt du document | [optional] 
**Mentions** | Pointer to **[]string** | Mentions de la recherche dans le document. | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Document) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Document) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Document) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Document) HasType() bool`

HasType returns a boolean if a field has been set.

### GetToken

`func (o *Document) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Document) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Document) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Document) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDateDepot

`func (o *Document) GetDateDepot() string`

GetDateDepot returns the DateDepot field if non-nil, zero value otherwise.

### GetDateDepotOk

`func (o *Document) GetDateDepotOk() (*string, bool)`

GetDateDepotOk returns a tuple with the DateDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDepot

`func (o *Document) SetDateDepot(v string)`

SetDateDepot sets DateDepot field to given value.

### HasDateDepot

`func (o *Document) HasDateDepot() bool`

HasDateDepot returns a boolean if a field has been set.

### GetMentions

`func (o *Document) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *Document) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *Document) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *Document) HasMentions() bool`

HasMentions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


