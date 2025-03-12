# EntrepriseBaseConventionsCollectivesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nom** | Pointer to **string** | Nom de la convention collective. | [optional] 
**Idcc** | Pointer to **int32** | Code IDCC de l&#39;entreprise. | [optional] 
**Confirmee** | Pointer to **bool** | Si confirmée, l&#39;information est issue de la DSN de l&#39;entreprise et fournie par le ministère du Travail. Si non confirmée, ce n&#39;est qu&#39;une estimation à partir du code NAF de l&#39;entreprise. | [optional] 
**Pourcentage** | Pointer to **float32** | Pourcentage de fiabilité de l&#39;estimation. Si la convention est confirmée, vaut null. | [optional] 

## Methods

### NewEntrepriseBaseConventionsCollectivesInner

`func NewEntrepriseBaseConventionsCollectivesInner() *EntrepriseBaseConventionsCollectivesInner`

NewEntrepriseBaseConventionsCollectivesInner instantiates a new EntrepriseBaseConventionsCollectivesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseBaseConventionsCollectivesInnerWithDefaults

`func NewEntrepriseBaseConventionsCollectivesInnerWithDefaults() *EntrepriseBaseConventionsCollectivesInner`

NewEntrepriseBaseConventionsCollectivesInnerWithDefaults instantiates a new EntrepriseBaseConventionsCollectivesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNom

`func (o *EntrepriseBaseConventionsCollectivesInner) GetNom() string`

GetNom returns the Nom field if non-nil, zero value otherwise.

### GetNomOk

`func (o *EntrepriseBaseConventionsCollectivesInner) GetNomOk() (*string, bool)`

GetNomOk returns a tuple with the Nom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNom

`func (o *EntrepriseBaseConventionsCollectivesInner) SetNom(v string)`

SetNom sets Nom field to given value.

### HasNom

`func (o *EntrepriseBaseConventionsCollectivesInner) HasNom() bool`

HasNom returns a boolean if a field has been set.

### GetIdcc

`func (o *EntrepriseBaseConventionsCollectivesInner) GetIdcc() int32`

GetIdcc returns the Idcc field if non-nil, zero value otherwise.

### GetIdccOk

`func (o *EntrepriseBaseConventionsCollectivesInner) GetIdccOk() (*int32, bool)`

GetIdccOk returns a tuple with the Idcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdcc

`func (o *EntrepriseBaseConventionsCollectivesInner) SetIdcc(v int32)`

SetIdcc sets Idcc field to given value.

### HasIdcc

`func (o *EntrepriseBaseConventionsCollectivesInner) HasIdcc() bool`

HasIdcc returns a boolean if a field has been set.

### GetConfirmee

`func (o *EntrepriseBaseConventionsCollectivesInner) GetConfirmee() bool`

GetConfirmee returns the Confirmee field if non-nil, zero value otherwise.

### GetConfirmeeOk

`func (o *EntrepriseBaseConventionsCollectivesInner) GetConfirmeeOk() (*bool, bool)`

GetConfirmeeOk returns a tuple with the Confirmee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmee

`func (o *EntrepriseBaseConventionsCollectivesInner) SetConfirmee(v bool)`

SetConfirmee sets Confirmee field to given value.

### HasConfirmee

`func (o *EntrepriseBaseConventionsCollectivesInner) HasConfirmee() bool`

HasConfirmee returns a boolean if a field has been set.

### GetPourcentage

`func (o *EntrepriseBaseConventionsCollectivesInner) GetPourcentage() float32`

GetPourcentage returns the Pourcentage field if non-nil, zero value otherwise.

### GetPourcentageOk

`func (o *EntrepriseBaseConventionsCollectivesInner) GetPourcentageOk() (*float32, bool)`

GetPourcentageOk returns a tuple with the Pourcentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPourcentage

`func (o *EntrepriseBaseConventionsCollectivesInner) SetPourcentage(v float32)`

SetPourcentage sets Pourcentage field to given value.

### HasPourcentage

`func (o *EntrepriseBaseConventionsCollectivesInner) HasPourcentage() bool`

HasPourcentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


