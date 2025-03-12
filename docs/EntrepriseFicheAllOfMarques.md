# EntrepriseFicheAllOfMarques

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numero** | Pointer to **string** | Numéro de la marque. | [optional] 
**DateEnregistrement** | Pointer to **string** | Date d&#39;enregistrement de la marque, au format AAAA-MM-JJ. | [optional] 
**DateExpiration** | Pointer to **string** | Date d&#39;expiration de la marque au format AAAA-MM-JJ. | [optional] 
**LieuEnregistrement** | Pointer to **string** | Lieu d&#39;enregistrement de la marque. | [optional] 
**Statut** | Pointer to **string** | Statut de la marque. La description des différents types est disponible en page 14 du document suivant : https://www.inpi.fr/sites/default/files/doctech_marques_v1.6.pdf. | [optional] 
**Texte** | Pointer to **string** | Texte de la marque. | [optional] 
**Type** | Pointer to **string** | Type de la marque. | [optional] 
**LienImage** | Pointer to **string** | Lien vers l&#39;image déposée. | [optional] 
**Descriptions** | Pointer to **[]string** | Liste des descriptions de la marque. | [optional] 
**Classes** | Pointer to [**[]EntrepriseFicheAllOfClasses**](EntrepriseFicheAllOfClasses.md) | Liste des classes (produits et services) de la marque. La liste des classes est disponible sur le document suivant : https://www.inpi.fr/sites/default/files/classification_nice_2021_0.pdf. | [optional] 
**Deposant** | Pointer to [**PersonneMarque**](PersonneMarque.md) |  | [optional] 
**Mandataire** | Pointer to [**PersonneMarque**](PersonneMarque.md) |  | [optional] 
**Evenements** | Pointer to [**[]EntrepriseFicheAllOfEvenements**](EntrepriseFicheAllOfEvenements.md) | Liste des événements associés à la marque. | [optional] 

## Methods

### NewEntrepriseFicheAllOfMarques

`func NewEntrepriseFicheAllOfMarques() *EntrepriseFicheAllOfMarques`

NewEntrepriseFicheAllOfMarques instantiates a new EntrepriseFicheAllOfMarques object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntrepriseFicheAllOfMarquesWithDefaults

`func NewEntrepriseFicheAllOfMarquesWithDefaults() *EntrepriseFicheAllOfMarques`

NewEntrepriseFicheAllOfMarquesWithDefaults instantiates a new EntrepriseFicheAllOfMarques object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumero

`func (o *EntrepriseFicheAllOfMarques) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *EntrepriseFicheAllOfMarques) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *EntrepriseFicheAllOfMarques) SetNumero(v string)`

SetNumero sets Numero field to given value.

### HasNumero

`func (o *EntrepriseFicheAllOfMarques) HasNumero() bool`

HasNumero returns a boolean if a field has been set.

### GetDateEnregistrement

`func (o *EntrepriseFicheAllOfMarques) GetDateEnregistrement() string`

GetDateEnregistrement returns the DateEnregistrement field if non-nil, zero value otherwise.

### GetDateEnregistrementOk

`func (o *EntrepriseFicheAllOfMarques) GetDateEnregistrementOk() (*string, bool)`

GetDateEnregistrementOk returns a tuple with the DateEnregistrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnregistrement

`func (o *EntrepriseFicheAllOfMarques) SetDateEnregistrement(v string)`

SetDateEnregistrement sets DateEnregistrement field to given value.

### HasDateEnregistrement

`func (o *EntrepriseFicheAllOfMarques) HasDateEnregistrement() bool`

HasDateEnregistrement returns a boolean if a field has been set.

### GetDateExpiration

`func (o *EntrepriseFicheAllOfMarques) GetDateExpiration() string`

GetDateExpiration returns the DateExpiration field if non-nil, zero value otherwise.

### GetDateExpirationOk

`func (o *EntrepriseFicheAllOfMarques) GetDateExpirationOk() (*string, bool)`

GetDateExpirationOk returns a tuple with the DateExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpiration

`func (o *EntrepriseFicheAllOfMarques) SetDateExpiration(v string)`

SetDateExpiration sets DateExpiration field to given value.

### HasDateExpiration

`func (o *EntrepriseFicheAllOfMarques) HasDateExpiration() bool`

HasDateExpiration returns a boolean if a field has been set.

### GetLieuEnregistrement

`func (o *EntrepriseFicheAllOfMarques) GetLieuEnregistrement() string`

GetLieuEnregistrement returns the LieuEnregistrement field if non-nil, zero value otherwise.

### GetLieuEnregistrementOk

`func (o *EntrepriseFicheAllOfMarques) GetLieuEnregistrementOk() (*string, bool)`

GetLieuEnregistrementOk returns a tuple with the LieuEnregistrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLieuEnregistrement

`func (o *EntrepriseFicheAllOfMarques) SetLieuEnregistrement(v string)`

SetLieuEnregistrement sets LieuEnregistrement field to given value.

### HasLieuEnregistrement

`func (o *EntrepriseFicheAllOfMarques) HasLieuEnregistrement() bool`

HasLieuEnregistrement returns a boolean if a field has been set.

### GetStatut

`func (o *EntrepriseFicheAllOfMarques) GetStatut() string`

GetStatut returns the Statut field if non-nil, zero value otherwise.

### GetStatutOk

`func (o *EntrepriseFicheAllOfMarques) GetStatutOk() (*string, bool)`

GetStatutOk returns a tuple with the Statut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatut

`func (o *EntrepriseFicheAllOfMarques) SetStatut(v string)`

SetStatut sets Statut field to given value.

### HasStatut

`func (o *EntrepriseFicheAllOfMarques) HasStatut() bool`

HasStatut returns a boolean if a field has been set.

### GetTexte

`func (o *EntrepriseFicheAllOfMarques) GetTexte() string`

GetTexte returns the Texte field if non-nil, zero value otherwise.

### GetTexteOk

`func (o *EntrepriseFicheAllOfMarques) GetTexteOk() (*string, bool)`

GetTexteOk returns a tuple with the Texte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexte

`func (o *EntrepriseFicheAllOfMarques) SetTexte(v string)`

SetTexte sets Texte field to given value.

### HasTexte

`func (o *EntrepriseFicheAllOfMarques) HasTexte() bool`

HasTexte returns a boolean if a field has been set.

### GetType

`func (o *EntrepriseFicheAllOfMarques) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntrepriseFicheAllOfMarques) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntrepriseFicheAllOfMarques) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntrepriseFicheAllOfMarques) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLienImage

`func (o *EntrepriseFicheAllOfMarques) GetLienImage() string`

GetLienImage returns the LienImage field if non-nil, zero value otherwise.

### GetLienImageOk

`func (o *EntrepriseFicheAllOfMarques) GetLienImageOk() (*string, bool)`

GetLienImageOk returns a tuple with the LienImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLienImage

`func (o *EntrepriseFicheAllOfMarques) SetLienImage(v string)`

SetLienImage sets LienImage field to given value.

### HasLienImage

`func (o *EntrepriseFicheAllOfMarques) HasLienImage() bool`

HasLienImage returns a boolean if a field has been set.

### GetDescriptions

`func (o *EntrepriseFicheAllOfMarques) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *EntrepriseFicheAllOfMarques) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *EntrepriseFicheAllOfMarques) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *EntrepriseFicheAllOfMarques) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetClasses

`func (o *EntrepriseFicheAllOfMarques) GetClasses() []EntrepriseFicheAllOfClasses`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *EntrepriseFicheAllOfMarques) GetClassesOk() (*[]EntrepriseFicheAllOfClasses, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *EntrepriseFicheAllOfMarques) SetClasses(v []EntrepriseFicheAllOfClasses)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *EntrepriseFicheAllOfMarques) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetDeposant

`func (o *EntrepriseFicheAllOfMarques) GetDeposant() PersonneMarque`

GetDeposant returns the Deposant field if non-nil, zero value otherwise.

### GetDeposantOk

`func (o *EntrepriseFicheAllOfMarques) GetDeposantOk() (*PersonneMarque, bool)`

GetDeposantOk returns a tuple with the Deposant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposant

`func (o *EntrepriseFicheAllOfMarques) SetDeposant(v PersonneMarque)`

SetDeposant sets Deposant field to given value.

### HasDeposant

`func (o *EntrepriseFicheAllOfMarques) HasDeposant() bool`

HasDeposant returns a boolean if a field has been set.

### GetMandataire

`func (o *EntrepriseFicheAllOfMarques) GetMandataire() PersonneMarque`

GetMandataire returns the Mandataire field if non-nil, zero value otherwise.

### GetMandataireOk

`func (o *EntrepriseFicheAllOfMarques) GetMandataireOk() (*PersonneMarque, bool)`

GetMandataireOk returns a tuple with the Mandataire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandataire

`func (o *EntrepriseFicheAllOfMarques) SetMandataire(v PersonneMarque)`

SetMandataire sets Mandataire field to given value.

### HasMandataire

`func (o *EntrepriseFicheAllOfMarques) HasMandataire() bool`

HasMandataire returns a boolean if a field has been set.

### GetEvenements

`func (o *EntrepriseFicheAllOfMarques) GetEvenements() []EntrepriseFicheAllOfEvenements`

GetEvenements returns the Evenements field if non-nil, zero value otherwise.

### GetEvenementsOk

`func (o *EntrepriseFicheAllOfMarques) GetEvenementsOk() (*[]EntrepriseFicheAllOfEvenements, bool)`

GetEvenementsOk returns a tuple with the Evenements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvenements

`func (o *EntrepriseFicheAllOfMarques) SetEvenements(v []EntrepriseFicheAllOfEvenements)`

SetEvenements sets Evenements field to given value.

### HasEvenements

`func (o *EntrepriseFicheAllOfMarques) HasEvenements() bool`

HasEvenements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


