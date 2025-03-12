# \SuiviUtilisationAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SuiviJetons**](SuiviUtilisationAPI.md#SuiviJetons) | **Get** /suivi-jetons | Récupère le suivi d&#39;utilisation des jetons.



## SuiviJetons

> SuiviJetons200Response SuiviJetons(ctx).Execute()

Récupère le suivi d'utilisation des jetons.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuiviUtilisationAPI.SuiviJetons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuiviUtilisationAPI.SuiviJetons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuiviJetons`: SuiviJetons200Response
	fmt.Fprintf(os.Stdout, "Response from `SuiviUtilisationAPI.SuiviJetons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSuiviJetonsRequest struct via the builder pattern


### Return type

[**SuiviJetons200Response**](SuiviJetons200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

