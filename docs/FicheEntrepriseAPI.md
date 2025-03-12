# \FicheEntrepriseAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Entreprise**](FicheEntrepriseAPI.md#Entreprise) | **Get** /entreprise | Récupère l&#39;ensemble des informations disponibles sur une entreprise.



## Entreprise

> EntrepriseFiche Entreprise(ctx).Siren(siren).Siret(siret).IntegrerDiffusionsPartielles(integrerDiffusionsPartielles).FormatPublicationsBodacc(formatPublicationsBodacc).Marques(marques).ValiditeTvaIntracommunautaire(validiteTvaIntracommunautaire).PublicationsBodaccBrutes(publicationsBodaccBrutes).ChampsSupplementaires(champsSupplementaires).Execute()

Récupère l'ensemble des informations disponibles sur une entreprise.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	siren := "443061841" // string | SIREN de l'entreprise (optional)
	siret := "44306184100047" // string | SIRET de l'entreprise (optional)
	integrerDiffusionsPartielles := true // bool | Si vrai et si l'entreprise est en diffusion partielle, le retour renverra les informations partielles disponibles. Valeur par défaut : `false`. (optional)
	formatPublicationsBodacc := "objet" // string | Format attendu pour les publications BODACC. Valeur par défaut : `\"objet\"`. (optional)
	marques := true // bool | Si vrai, le retour inclura les marques éventuelles de l'entreprise. Valeur par défaut : `false`. (optional)
	validiteTvaIntracommunautaire := true // bool | Si vrai, le champ validite_tva_intracommunautaire du retour indiquera si le numéro de tva est valide auprès de la Commission européenne. Valeur par défaut : `false`. (optional)
	publicationsBodaccBrutes := true // bool | Pappers traite les publications BODACC afin de supprimer les publications périmée. Si vrai, le retour inclura les publications bodacc sans traitement. Valeur par défaut : `false`. (optional)
	champsSupplementaires := "site_internet,telephone" // string | Liste des champs supplémentaires à inclure dans le retour. Certains champs peuvent entraîner une consommation de jetons supplémentaires.  Champs supplémentaires disponibles : - `sites_internet` : 1 jeton supplémentaire - `telephone` : 1 jeton supplémentaire * - `email` : 1 jeton supplémentaire * - `enseigne_1` : gratuit - `enseigne_2` : gratuit - `enseigne_3` : gratuit - `distribution_speciale` : gratuit - `code_cedex` : gratuit - `libelle_cedex` : gratuit - `code_commune` : gratuit - `code_region` : gratuit - `region` : gratuit - `code_departement` : gratuit - `departement` : gratuit - `nomenclature_code_naf` : gratuit - `labels` : gratuit - `labels:orias` : 0.5 jeton supplémentaire - `labels:cci` : 0.5 jeton supplémentaire - `micro_entreprise` : gratuit - `sanctions` : 1 jeton supplémentaire - `personne_politiquement_exposee` : 1 jeton supplémentaire - `deces` : 0.5 jeton supplémentaire - `scoring_financier` : 30 jetons supplémentaires - `scoring_non_financier` : 30 jetons supplémentaires - `categorie_entreprise` : gratuit  \\* : le coût des champs `telephone` et `email` est de 1 jeton supplémentaire au total, même si les deux sont demandés.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FicheEntrepriseAPI.Entreprise(context.Background()).Siren(siren).Siret(siret).IntegrerDiffusionsPartielles(integrerDiffusionsPartielles).FormatPublicationsBodacc(formatPublicationsBodacc).Marques(marques).ValiditeTvaIntracommunautaire(validiteTvaIntracommunautaire).PublicationsBodaccBrutes(publicationsBodaccBrutes).ChampsSupplementaires(champsSupplementaires).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FicheEntrepriseAPI.Entreprise``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Entreprise`: EntrepriseFiche
	fmt.Fprintf(os.Stdout, "Response from `FicheEntrepriseAPI.Entreprise`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntrepriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **siret** | **string** | SIRET de l&#39;entreprise | 
 **integrerDiffusionsPartielles** | **bool** | Si vrai et si l&#39;entreprise est en diffusion partielle, le retour renverra les informations partielles disponibles. Valeur par défaut : &#x60;false&#x60;. | 
 **formatPublicationsBodacc** | **string** | Format attendu pour les publications BODACC. Valeur par défaut : &#x60;\&quot;objet\&quot;&#x60;. | 
 **marques** | **bool** | Si vrai, le retour inclura les marques éventuelles de l&#39;entreprise. Valeur par défaut : &#x60;false&#x60;. | 
 **validiteTvaIntracommunautaire** | **bool** | Si vrai, le champ validite_tva_intracommunautaire du retour indiquera si le numéro de tva est valide auprès de la Commission européenne. Valeur par défaut : &#x60;false&#x60;. | 
 **publicationsBodaccBrutes** | **bool** | Pappers traite les publications BODACC afin de supprimer les publications périmée. Si vrai, le retour inclura les publications bodacc sans traitement. Valeur par défaut : &#x60;false&#x60;. | 
 **champsSupplementaires** | **string** | Liste des champs supplémentaires à inclure dans le retour. Certains champs peuvent entraîner une consommation de jetons supplémentaires.  Champs supplémentaires disponibles : - &#x60;sites_internet&#x60; : 1 jeton supplémentaire - &#x60;telephone&#x60; : 1 jeton supplémentaire * - &#x60;email&#x60; : 1 jeton supplémentaire * - &#x60;enseigne_1&#x60; : gratuit - &#x60;enseigne_2&#x60; : gratuit - &#x60;enseigne_3&#x60; : gratuit - &#x60;distribution_speciale&#x60; : gratuit - &#x60;code_cedex&#x60; : gratuit - &#x60;libelle_cedex&#x60; : gratuit - &#x60;code_commune&#x60; : gratuit - &#x60;code_region&#x60; : gratuit - &#x60;region&#x60; : gratuit - &#x60;code_departement&#x60; : gratuit - &#x60;departement&#x60; : gratuit - &#x60;nomenclature_code_naf&#x60; : gratuit - &#x60;labels&#x60; : gratuit - &#x60;labels:orias&#x60; : 0.5 jeton supplémentaire - &#x60;labels:cci&#x60; : 0.5 jeton supplémentaire - &#x60;micro_entreprise&#x60; : gratuit - &#x60;sanctions&#x60; : 1 jeton supplémentaire - &#x60;personne_politiquement_exposee&#x60; : 1 jeton supplémentaire - &#x60;deces&#x60; : 0.5 jeton supplémentaire - &#x60;scoring_financier&#x60; : 30 jetons supplémentaires - &#x60;scoring_non_financier&#x60; : 30 jetons supplémentaires - &#x60;categorie_entreprise&#x60; : gratuit  \\* : le coût des champs &#x60;telephone&#x60; et &#x60;email&#x60; est de 1 jeton supplémentaire au total, même si les deux sont demandés.  | 

### Return type

[**EntrepriseFiche**](EntrepriseFiche.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

