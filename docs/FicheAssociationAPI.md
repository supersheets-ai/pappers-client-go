# \FicheAssociationAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Association**](FicheAssociationAPI.md#Association) | **Get** /association | Récupère l&#39;ensemble des informations disponibles sur une association.



## Association

> Association Association(ctx).IdAssociation(idAssociation).Siret(siret).Siren(siren).Execute()

Récupère l'ensemble des informations disponibles sur une association.



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
	idAssociation := "W142005389" // string | Identifiant de l'association (optional)
	siret := "95037097300014" // string | SIRET de l'association (optional)
	siren := "950370973" // string | SIREN de l'association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FicheAssociationAPI.Association(context.Background()).IdAssociation(idAssociation).Siret(siret).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FicheAssociationAPI.Association``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Association`: Association
	fmt.Fprintf(os.Stdout, "Response from `FicheAssociationAPI.Association`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idAssociation** | **string** | Identifiant de l&#39;association | 
 **siret** | **string** | SIRET de l&#39;association | 
 **siren** | **string** | SIREN de l&#39;association | 

### Return type

[**Association**](Association.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

