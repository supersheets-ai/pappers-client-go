# \CartographieAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cartographie**](CartographieAPI.md#Cartographie) | **Get** /entreprise/cartographie | Récupère les données permettant d&#39;établir une cartographie d&#39;une entreprise.



## Cartographie

> Cartographie Cartographie(ctx).Siren(siren).InclureEntreprisesDirigees(inclureEntreprisesDirigees).InclureEntreprisesCitees(inclureEntreprisesCitees).InclureSci(inclureSci).AutoriserModifications(autoriserModifications).RejeterPremierDegre(rejeterPremierDegre).Degre(degre).Execute()

Récupère les données permettant d'établir une cartographie d'une entreprise.



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
	siren := "443061841" // string | SIREN de l'entreprise
	inclureEntreprisesDirigees := true // bool | Si vrai, la cartographie intègrera les entreprises dirigées par l'entreprise recherchée et les entreprises qui dirigent l'entreprise recherchée. Valeur par défaut : `true`. (optional)
	inclureEntreprisesCitees := true // bool | Si vrai, la cartographie intègrera les entreprises citées conjointement avec l'entreprise recherchée dans des actes et statuts. Valeur par défaut : `false`. (optional)
	inclureSci := true // bool | Si vrai, la cartographie intègrera les SCI. Valeur par défaut : `true`. (optional)
	autoriserModifications := true // bool | Si vrai, la cartographie pourra adapter automatiquement ses paramètres si ceux choisis manuellement ne sont pas idéaux. Valeur par défaut : `false`. (optional)
	rejeterPremierDegre := true // bool | Si vrai et que la cartographie ne fait apparaître que l'entreprise recherchée ainsi que ses dirigeants directs, une erreur 404 sera renvoyée et la requête ne sera pas comptabilisée dans le quota de jetons. Valeur par défaut : `false`. (optional)
	degre := int32(56) // int32 | Permet de choisir manuellement un degré pour la cartographie. Seuls deux états sont possibles : un nombre <= 2 ou bien un nombre > 2. Cela veut dire que 0, 1 ou 2 donneront la même cartographie, tout comme 3, 4 ou 5. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartographieAPI.Cartographie(context.Background()).Siren(siren).InclureEntreprisesDirigees(inclureEntreprisesDirigees).InclureEntreprisesCitees(inclureEntreprisesCitees).InclureSci(inclureSci).AutoriserModifications(autoriserModifications).RejeterPremierDegre(rejeterPremierDegre).Degre(degre).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartographieAPI.Cartographie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cartographie`: Cartographie
	fmt.Fprintf(os.Stdout, "Response from `CartographieAPI.Cartographie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartographieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **inclureEntreprisesDirigees** | **bool** | Si vrai, la cartographie intègrera les entreprises dirigées par l&#39;entreprise recherchée et les entreprises qui dirigent l&#39;entreprise recherchée. Valeur par défaut : &#x60;true&#x60;. | 
 **inclureEntreprisesCitees** | **bool** | Si vrai, la cartographie intègrera les entreprises citées conjointement avec l&#39;entreprise recherchée dans des actes et statuts. Valeur par défaut : &#x60;false&#x60;. | 
 **inclureSci** | **bool** | Si vrai, la cartographie intègrera les SCI. Valeur par défaut : &#x60;true&#x60;. | 
 **autoriserModifications** | **bool** | Si vrai, la cartographie pourra adapter automatiquement ses paramètres si ceux choisis manuellement ne sont pas idéaux. Valeur par défaut : &#x60;false&#x60;. | 
 **rejeterPremierDegre** | **bool** | Si vrai et que la cartographie ne fait apparaître que l&#39;entreprise recherchée ainsi que ses dirigeants directs, une erreur 404 sera renvoyée et la requête ne sera pas comptabilisée dans le quota de jetons. Valeur par défaut : &#x60;false&#x60;. | 
 **degre** | **int32** | Permet de choisir manuellement un degré pour la cartographie. Seuls deux états sont possibles : un nombre &lt;&#x3D; 2 ou bien un nombre &gt; 2. Cela veut dire que 0, 1 ou 2 donneront la même cartographie, tout comme 3, 4 ou 5. | 

### Return type

[**Cartographie**](Cartographie.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

