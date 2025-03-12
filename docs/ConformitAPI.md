# \ConformitAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Conformite**](ConformitAPI.md#Conformite) | **Get** /conformite/personne_physique | Vérifie le statut de personne politiquement exposée et la présence de sanctions internationales pour une personne physique.



## Conformite

> Conformite200Response Conformite(ctx).Nom(nom).Prenom(prenom).DateDeNaissance(dateDeNaissance).Execute()

Vérifie le statut de personne politiquement exposée et la présence de sanctions internationales pour une personne physique.



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
	nom := "Macron" // string | Nom de la personne physique
	prenom := "Emmanuel" // string | Prénom de la personne physique
	dateDeNaissance := "1977-12-21" // string | Date de naissance de la personne physique, au format AAAA-MM-JJ ou AAAA-MM

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConformitAPI.Conformite(context.Background()).Nom(nom).Prenom(prenom).DateDeNaissance(dateDeNaissance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConformitAPI.Conformite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Conformite`: Conformite200Response
	fmt.Fprintf(os.Stdout, "Response from `ConformitAPI.Conformite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConformiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nom** | **string** | Nom de la personne physique | 
 **prenom** | **string** | Prénom de la personne physique | 
 **dateDeNaissance** | **string** | Date de naissance de la personne physique, au format AAAA-MM-JJ ou AAAA-MM | 

### Return type

[**Conformite200Response**](Conformite200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

