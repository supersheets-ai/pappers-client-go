# \ComptesAnnuelsAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComptesAnnuels**](ComptesAnnuelsAPI.md#ComptesAnnuels) | **Get** /entreprise/comptes | Récupère les comptes annuels publiés d&#39;une entreprise.



## ComptesAnnuels

> map[string][]ComptesAnnuels200ResponseValueInner ComptesAnnuels(ctx).Siren(siren).Annee(annee).Execute()

Récupère les comptes annuels publiés d'une entreprise.



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
	annee := "2016,2017,2018" // string | Année de cloture des comptes souhaités. Il est possible d'indiquer plusieurs années en les séparant par des virgules. Si le paramètre n'est pas fourni, toutes les années disponibles seront retournées. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComptesAnnuelsAPI.ComptesAnnuels(context.Background()).Siren(siren).Annee(annee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComptesAnnuelsAPI.ComptesAnnuels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComptesAnnuels`: map[string][]ComptesAnnuels200ResponseValueInner
	fmt.Fprintf(os.Stdout, "Response from `ComptesAnnuelsAPI.ComptesAnnuels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComptesAnnuelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **annee** | **string** | Année de cloture des comptes souhaités. Il est possible d&#39;indiquer plusieurs années en les séparant par des virgules. Si le paramètre n&#39;est pas fourni, toutes les années disponibles seront retournées. | 

### Return type

[**map[string][]ComptesAnnuels200ResponseValueInner**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

