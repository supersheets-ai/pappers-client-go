# Association

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsWaldec** | Pointer to **bool** | Détermine si l&#39;association possède un numéro WALDEC/RNA. | [optional] 
**IdAssociation** | Pointer to **string** | L&#39;identifiant l&#39;association au format Wxxxxxxxxx si WALDEC, xxxxxxxxxxxxxx sinon. | [optional] 
**IdExAssociation** | Pointer to **string** | Ancien numéro de l&#39;association. | [optional] 
**Denomination** | Pointer to **string** | Dénomination de l&#39;association. | [optional] 
**Siret** | Pointer to **string** | Numéro siret de l&#39;association au format xxxxxxxxxxxxxx. | [optional] 
**NumeroRup** | Pointer to **string** | Numéro de Reconnaissance d&#39;Utilité Publique. | [optional] 
**Objet** | Pointer to **string** | Objet de l&#39;association. | [optional] 
**ObjetSocial1** | Pointer to **string** | Objet social 1 de l&#39;association. | [optional] 
**CategorieSocial1** | Pointer to **string** | Libellé correspondant à l&#39;objet social 1. | [optional] 
**ObjetSocial2** | Pointer to **string** | Objet social 2 de l&#39;association. | [optional] 
**CategorieSocial2** | Pointer to **string** | Libellé correspondant à l&#39;objet social 1. | [optional] 
**DateCreation** | Pointer to **string** | Date de déclaration de création au format AAAA-MM-JJ. | [optional] 
**DateDerniereDeclaration** | Pointer to **string** | Date de dernière déclaration au format AAAA-MM-JJ. | [optional] 
**DatePublicationCreation** | Pointer to **string** | Date de publication du Journal Officiel de l&#39;avis de création au format AAAA-MM-JJ. | [optional] 
**DateDeclarationDissolution** | Pointer to **string** | Date de déclaration de dissolution au format AAAA-MM-JJ. | [optional] 
**Groupement** | Pointer to **string** | Groupement de l&#39;association. | [optional] 
**PositionActivite** | Pointer to **string** | Position d&#39;activité de l&#39;association. | [optional] 
**Nature** | Pointer to **string** | Nature de l&#39;association. | [optional] 
**SiteWeb** | Pointer to **string** | Site web de l&#39;association. | [optional] 
**Telephone** | Pointer to **string** | Numéro de téléphone de l&#39;association. | [optional] 
**Email** | Pointer to **string** | Email de l&#39;association. | [optional] 
**AdresseSiege** | Pointer to [**AssociationAdresseSiege**](AssociationAdresseSiege.md) |  | [optional] 
**AdresseGestionnaire** | Pointer to [**AssociationAdresseGestionnaire**](AssociationAdresseGestionnaire.md) |  | [optional] 
**Observation** | Pointer to **string** | Observation relative à l&#39;association. | [optional] 
**CodeGestion** | Pointer to **string** | Code du site gestionnaire (préfecture, sous-préfecture) de l&#39;association. | [optional] 
**DirigeantCivilite** | Pointer to **string** | Civilité du dirigeant. | [optional] 
**DerniereMaj** | Pointer to **string** | Date de la dernière mise à jour des informations au RNA au format AAAA-MM-JJ. | [optional] 
**PublicationsJoafe** | Pointer to [**[]AssociationPublicationsJoafeInner**](AssociationPublicationsJoafeInner.md) | Publications JOAFE. | [optional] 

## Methods

### NewAssociation

`func NewAssociation() *Association`

NewAssociation instantiates a new Association object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationWithDefaults

`func NewAssociationWithDefaults() *Association`

NewAssociationWithDefaults instantiates a new Association object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsWaldec

`func (o *Association) GetIsWaldec() bool`

GetIsWaldec returns the IsWaldec field if non-nil, zero value otherwise.

### GetIsWaldecOk

`func (o *Association) GetIsWaldecOk() (*bool, bool)`

GetIsWaldecOk returns a tuple with the IsWaldec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaldec

`func (o *Association) SetIsWaldec(v bool)`

SetIsWaldec sets IsWaldec field to given value.

### HasIsWaldec

`func (o *Association) HasIsWaldec() bool`

HasIsWaldec returns a boolean if a field has been set.

### GetIdAssociation

`func (o *Association) GetIdAssociation() string`

GetIdAssociation returns the IdAssociation field if non-nil, zero value otherwise.

### GetIdAssociationOk

`func (o *Association) GetIdAssociationOk() (*string, bool)`

GetIdAssociationOk returns a tuple with the IdAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdAssociation

`func (o *Association) SetIdAssociation(v string)`

SetIdAssociation sets IdAssociation field to given value.

### HasIdAssociation

`func (o *Association) HasIdAssociation() bool`

HasIdAssociation returns a boolean if a field has been set.

### GetIdExAssociation

`func (o *Association) GetIdExAssociation() string`

GetIdExAssociation returns the IdExAssociation field if non-nil, zero value otherwise.

### GetIdExAssociationOk

`func (o *Association) GetIdExAssociationOk() (*string, bool)`

GetIdExAssociationOk returns a tuple with the IdExAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdExAssociation

`func (o *Association) SetIdExAssociation(v string)`

SetIdExAssociation sets IdExAssociation field to given value.

### HasIdExAssociation

`func (o *Association) HasIdExAssociation() bool`

HasIdExAssociation returns a boolean if a field has been set.

### GetDenomination

`func (o *Association) GetDenomination() string`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *Association) GetDenominationOk() (*string, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *Association) SetDenomination(v string)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *Association) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetSiret

`func (o *Association) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Association) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Association) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Association) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetNumeroRup

`func (o *Association) GetNumeroRup() string`

GetNumeroRup returns the NumeroRup field if non-nil, zero value otherwise.

### GetNumeroRupOk

`func (o *Association) GetNumeroRupOk() (*string, bool)`

GetNumeroRupOk returns a tuple with the NumeroRup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroRup

`func (o *Association) SetNumeroRup(v string)`

SetNumeroRup sets NumeroRup field to given value.

### HasNumeroRup

`func (o *Association) HasNumeroRup() bool`

HasNumeroRup returns a boolean if a field has been set.

### GetObjet

`func (o *Association) GetObjet() string`

GetObjet returns the Objet field if non-nil, zero value otherwise.

### GetObjetOk

`func (o *Association) GetObjetOk() (*string, bool)`

GetObjetOk returns a tuple with the Objet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjet

`func (o *Association) SetObjet(v string)`

SetObjet sets Objet field to given value.

### HasObjet

`func (o *Association) HasObjet() bool`

HasObjet returns a boolean if a field has been set.

### GetObjetSocial1

`func (o *Association) GetObjetSocial1() string`

GetObjetSocial1 returns the ObjetSocial1 field if non-nil, zero value otherwise.

### GetObjetSocial1Ok

`func (o *Association) GetObjetSocial1Ok() (*string, bool)`

GetObjetSocial1Ok returns a tuple with the ObjetSocial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjetSocial1

`func (o *Association) SetObjetSocial1(v string)`

SetObjetSocial1 sets ObjetSocial1 field to given value.

### HasObjetSocial1

`func (o *Association) HasObjetSocial1() bool`

HasObjetSocial1 returns a boolean if a field has been set.

### GetCategorieSocial1

`func (o *Association) GetCategorieSocial1() string`

GetCategorieSocial1 returns the CategorieSocial1 field if non-nil, zero value otherwise.

### GetCategorieSocial1Ok

`func (o *Association) GetCategorieSocial1Ok() (*string, bool)`

GetCategorieSocial1Ok returns a tuple with the CategorieSocial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieSocial1

`func (o *Association) SetCategorieSocial1(v string)`

SetCategorieSocial1 sets CategorieSocial1 field to given value.

### HasCategorieSocial1

`func (o *Association) HasCategorieSocial1() bool`

HasCategorieSocial1 returns a boolean if a field has been set.

### GetObjetSocial2

`func (o *Association) GetObjetSocial2() string`

GetObjetSocial2 returns the ObjetSocial2 field if non-nil, zero value otherwise.

### GetObjetSocial2Ok

`func (o *Association) GetObjetSocial2Ok() (*string, bool)`

GetObjetSocial2Ok returns a tuple with the ObjetSocial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjetSocial2

`func (o *Association) SetObjetSocial2(v string)`

SetObjetSocial2 sets ObjetSocial2 field to given value.

### HasObjetSocial2

`func (o *Association) HasObjetSocial2() bool`

HasObjetSocial2 returns a boolean if a field has been set.

### GetCategorieSocial2

`func (o *Association) GetCategorieSocial2() string`

GetCategorieSocial2 returns the CategorieSocial2 field if non-nil, zero value otherwise.

### GetCategorieSocial2Ok

`func (o *Association) GetCategorieSocial2Ok() (*string, bool)`

GetCategorieSocial2Ok returns a tuple with the CategorieSocial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorieSocial2

`func (o *Association) SetCategorieSocial2(v string)`

SetCategorieSocial2 sets CategorieSocial2 field to given value.

### HasCategorieSocial2

`func (o *Association) HasCategorieSocial2() bool`

HasCategorieSocial2 returns a boolean if a field has been set.

### GetDateCreation

`func (o *Association) GetDateCreation() string`

GetDateCreation returns the DateCreation field if non-nil, zero value otherwise.

### GetDateCreationOk

`func (o *Association) GetDateCreationOk() (*string, bool)`

GetDateCreationOk returns a tuple with the DateCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreation

`func (o *Association) SetDateCreation(v string)`

SetDateCreation sets DateCreation field to given value.

### HasDateCreation

`func (o *Association) HasDateCreation() bool`

HasDateCreation returns a boolean if a field has been set.

### GetDateDerniereDeclaration

`func (o *Association) GetDateDerniereDeclaration() string`

GetDateDerniereDeclaration returns the DateDerniereDeclaration field if non-nil, zero value otherwise.

### GetDateDerniereDeclarationOk

`func (o *Association) GetDateDerniereDeclarationOk() (*string, bool)`

GetDateDerniereDeclarationOk returns a tuple with the DateDerniereDeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDerniereDeclaration

`func (o *Association) SetDateDerniereDeclaration(v string)`

SetDateDerniereDeclaration sets DateDerniereDeclaration field to given value.

### HasDateDerniereDeclaration

`func (o *Association) HasDateDerniereDeclaration() bool`

HasDateDerniereDeclaration returns a boolean if a field has been set.

### GetDatePublicationCreation

`func (o *Association) GetDatePublicationCreation() string`

GetDatePublicationCreation returns the DatePublicationCreation field if non-nil, zero value otherwise.

### GetDatePublicationCreationOk

`func (o *Association) GetDatePublicationCreationOk() (*string, bool)`

GetDatePublicationCreationOk returns a tuple with the DatePublicationCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublicationCreation

`func (o *Association) SetDatePublicationCreation(v string)`

SetDatePublicationCreation sets DatePublicationCreation field to given value.

### HasDatePublicationCreation

`func (o *Association) HasDatePublicationCreation() bool`

HasDatePublicationCreation returns a boolean if a field has been set.

### GetDateDeclarationDissolution

`func (o *Association) GetDateDeclarationDissolution() string`

GetDateDeclarationDissolution returns the DateDeclarationDissolution field if non-nil, zero value otherwise.

### GetDateDeclarationDissolutionOk

`func (o *Association) GetDateDeclarationDissolutionOk() (*string, bool)`

GetDateDeclarationDissolutionOk returns a tuple with the DateDeclarationDissolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeclarationDissolution

`func (o *Association) SetDateDeclarationDissolution(v string)`

SetDateDeclarationDissolution sets DateDeclarationDissolution field to given value.

### HasDateDeclarationDissolution

`func (o *Association) HasDateDeclarationDissolution() bool`

HasDateDeclarationDissolution returns a boolean if a field has been set.

### GetGroupement

`func (o *Association) GetGroupement() string`

GetGroupement returns the Groupement field if non-nil, zero value otherwise.

### GetGroupementOk

`func (o *Association) GetGroupementOk() (*string, bool)`

GetGroupementOk returns a tuple with the Groupement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupement

`func (o *Association) SetGroupement(v string)`

SetGroupement sets Groupement field to given value.

### HasGroupement

`func (o *Association) HasGroupement() bool`

HasGroupement returns a boolean if a field has been set.

### GetPositionActivite

`func (o *Association) GetPositionActivite() string`

GetPositionActivite returns the PositionActivite field if non-nil, zero value otherwise.

### GetPositionActiviteOk

`func (o *Association) GetPositionActiviteOk() (*string, bool)`

GetPositionActiviteOk returns a tuple with the PositionActivite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionActivite

`func (o *Association) SetPositionActivite(v string)`

SetPositionActivite sets PositionActivite field to given value.

### HasPositionActivite

`func (o *Association) HasPositionActivite() bool`

HasPositionActivite returns a boolean if a field has been set.

### GetNature

`func (o *Association) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *Association) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *Association) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *Association) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetSiteWeb

`func (o *Association) GetSiteWeb() string`

GetSiteWeb returns the SiteWeb field if non-nil, zero value otherwise.

### GetSiteWebOk

`func (o *Association) GetSiteWebOk() (*string, bool)`

GetSiteWebOk returns a tuple with the SiteWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteWeb

`func (o *Association) SetSiteWeb(v string)`

SetSiteWeb sets SiteWeb field to given value.

### HasSiteWeb

`func (o *Association) HasSiteWeb() bool`

HasSiteWeb returns a boolean if a field has been set.

### GetTelephone

`func (o *Association) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *Association) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *Association) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *Association) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetEmail

`func (o *Association) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Association) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Association) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Association) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdresseSiege

`func (o *Association) GetAdresseSiege() AssociationAdresseSiege`

GetAdresseSiege returns the AdresseSiege field if non-nil, zero value otherwise.

### GetAdresseSiegeOk

`func (o *Association) GetAdresseSiegeOk() (*AssociationAdresseSiege, bool)`

GetAdresseSiegeOk returns a tuple with the AdresseSiege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseSiege

`func (o *Association) SetAdresseSiege(v AssociationAdresseSiege)`

SetAdresseSiege sets AdresseSiege field to given value.

### HasAdresseSiege

`func (o *Association) HasAdresseSiege() bool`

HasAdresseSiege returns a boolean if a field has been set.

### GetAdresseGestionnaire

`func (o *Association) GetAdresseGestionnaire() AssociationAdresseGestionnaire`

GetAdresseGestionnaire returns the AdresseGestionnaire field if non-nil, zero value otherwise.

### GetAdresseGestionnaireOk

`func (o *Association) GetAdresseGestionnaireOk() (*AssociationAdresseGestionnaire, bool)`

GetAdresseGestionnaireOk returns a tuple with the AdresseGestionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdresseGestionnaire

`func (o *Association) SetAdresseGestionnaire(v AssociationAdresseGestionnaire)`

SetAdresseGestionnaire sets AdresseGestionnaire field to given value.

### HasAdresseGestionnaire

`func (o *Association) HasAdresseGestionnaire() bool`

HasAdresseGestionnaire returns a boolean if a field has been set.

### GetObservation

`func (o *Association) GetObservation() string`

GetObservation returns the Observation field if non-nil, zero value otherwise.

### GetObservationOk

`func (o *Association) GetObservationOk() (*string, bool)`

GetObservationOk returns a tuple with the Observation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservation

`func (o *Association) SetObservation(v string)`

SetObservation sets Observation field to given value.

### HasObservation

`func (o *Association) HasObservation() bool`

HasObservation returns a boolean if a field has been set.

### GetCodeGestion

`func (o *Association) GetCodeGestion() string`

GetCodeGestion returns the CodeGestion field if non-nil, zero value otherwise.

### GetCodeGestionOk

`func (o *Association) GetCodeGestionOk() (*string, bool)`

GetCodeGestionOk returns a tuple with the CodeGestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeGestion

`func (o *Association) SetCodeGestion(v string)`

SetCodeGestion sets CodeGestion field to given value.

### HasCodeGestion

`func (o *Association) HasCodeGestion() bool`

HasCodeGestion returns a boolean if a field has been set.

### GetDirigeantCivilite

`func (o *Association) GetDirigeantCivilite() string`

GetDirigeantCivilite returns the DirigeantCivilite field if non-nil, zero value otherwise.

### GetDirigeantCiviliteOk

`func (o *Association) GetDirigeantCiviliteOk() (*string, bool)`

GetDirigeantCiviliteOk returns a tuple with the DirigeantCivilite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirigeantCivilite

`func (o *Association) SetDirigeantCivilite(v string)`

SetDirigeantCivilite sets DirigeantCivilite field to given value.

### HasDirigeantCivilite

`func (o *Association) HasDirigeantCivilite() bool`

HasDirigeantCivilite returns a boolean if a field has been set.

### GetDerniereMaj

`func (o *Association) GetDerniereMaj() string`

GetDerniereMaj returns the DerniereMaj field if non-nil, zero value otherwise.

### GetDerniereMajOk

`func (o *Association) GetDerniereMajOk() (*string, bool)`

GetDerniereMajOk returns a tuple with the DerniereMaj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerniereMaj

`func (o *Association) SetDerniereMaj(v string)`

SetDerniereMaj sets DerniereMaj field to given value.

### HasDerniereMaj

`func (o *Association) HasDerniereMaj() bool`

HasDerniereMaj returns a boolean if a field has been set.

### GetPublicationsJoafe

`func (o *Association) GetPublicationsJoafe() []AssociationPublicationsJoafeInner`

GetPublicationsJoafe returns the PublicationsJoafe field if non-nil, zero value otherwise.

### GetPublicationsJoafeOk

`func (o *Association) GetPublicationsJoafeOk() (*[]AssociationPublicationsJoafeInner, bool)`

GetPublicationsJoafeOk returns a tuple with the PublicationsJoafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationsJoafe

`func (o *Association) SetPublicationsJoafe(v []AssociationPublicationsJoafeInner)`

SetPublicationsJoafe sets PublicationsJoafe field to given value.

### HasPublicationsJoafe

`func (o *Association) HasPublicationsJoafe() bool`

HasPublicationsJoafe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


