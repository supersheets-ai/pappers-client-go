# \SuggestionsAutocompleteAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Suggestions**](SuggestionsAutocompleteAPI.md#Suggestions) | **Get** /suggestions | Recherche les entreprises qui correspondent à un début de recherche textuelle (type autocomplete).



## Suggestions

> Suggestions200Response Suggestions(ctx).Q(q).Longueur(longueur).Cibles(cibles).Execute()

Recherche les entreprises qui correspondent à un début de recherche textuelle (type autocomplete).



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
	q := "googl" // string | Début de recherche textuelle
	longueur := int32(20) // int32 | Nombre de résultats. Maximum 100. Valeur par défaut : `10`. (optional)
	cibles := "nom_entreprise,siren,siret" // string | Cibles de la recherche, séparées par des virgules. Valeur par défaut : `\"nom_entreprise\"`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAutocompleteAPI.Suggestions(context.Background()).Q(q).Longueur(longueur).Cibles(cibles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAutocompleteAPI.Suggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Suggestions`: Suggestions200Response
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAutocompleteAPI.Suggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Début de recherche textuelle | 
 **longueur** | **int32** | Nombre de résultats. Maximum 100. Valeur par défaut : &#x60;10&#x60;. | 
 **cibles** | **string** | Cibles de la recherche, séparées par des virgules. Valeur par défaut : &#x60;\&quot;nom_entreprise\&quot;&#x60;. | 

### Return type

[**Suggestions200Response**](Suggestions200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

