# \DocumentsAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentAvisSituationInsee**](DocumentsAPI.md#DocumentAvisSituationInsee) | **Get** /document/avis_situation_insee | Télécharge l&#39;avis de situation INSEE d&#39;une entreprise à partir de son SIREN ou SIRET.
[**DocumentBeneficiairesEffectifs**](DocumentsAPI.md#DocumentBeneficiairesEffectifs) | **Get** /document/declaration_beneficiaires_effectifs | Télécharge la déclaration des bénéficiaires effectifs d&#39;une entreprise à partir de son SIREN.
[**DocumentExtraitInpi**](DocumentsAPI.md#DocumentExtraitInpi) | **Get** /document/extrait_inpi | Télécharge l&#39;extrait INPI d&#39;une entreprise à partir de son SIREN ou SIRET.
[**DocumentExtraitPappers**](DocumentsAPI.md#DocumentExtraitPappers) | **Get** /document/extrait_pappers | Télécharge l&#39;extrait Pappers d&#39;une entreprise à partir de son SIREN ou SIRET.
[**DocumentScoringFinancier**](DocumentsAPI.md#DocumentScoringFinancier) | **Get** /document/rapport_financier | Télécharge le rapport financier d&#39;une entreprise à partir de son SIREN.
[**DocumentScoringNonFinancier**](DocumentsAPI.md#DocumentScoringNonFinancier) | **Get** /document/rapport_non_financier | Télécharge le rapport non financier d&#39;une entreprise à partir de son SIREN.
[**DocumentStatus**](DocumentsAPI.md#DocumentStatus) | **Get** /document/statuts | Télécharge les derniers statuts disponibles d&#39;une entreprise à partir de son SIREN ou SIRET.
[**DocumentTelechargement**](DocumentsAPI.md#DocumentTelechargement) | **Get** /document/telechargement | Télécharge un document PDF ou XLSX à partir de son token.



## DocumentAvisSituationInsee

> *os.File DocumentAvisSituationInsee(ctx).Siren(siren).Siret(siret).Execute()

Télécharge l'avis de situation INSEE d'une entreprise à partir de son SIREN ou SIRET.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentAvisSituationInsee(context.Background()).Siren(siren).Siret(siret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentAvisSituationInsee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentAvisSituationInsee`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentAvisSituationInsee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentAvisSituationInseeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **siret** | **string** | SIRET de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentBeneficiairesEffectifs

> *os.File DocumentBeneficiairesEffectifs(ctx).Siren(siren).Execute()

Télécharge la déclaration des bénéficiaires effectifs d'une entreprise à partir de son SIREN.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentBeneficiairesEffectifs(context.Background()).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentBeneficiairesEffectifs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentBeneficiairesEffectifs`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentBeneficiairesEffectifs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentBeneficiairesEffectifsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentExtraitInpi

> *os.File DocumentExtraitInpi(ctx).Siren(siren).Siret(siret).Execute()

Télécharge l'extrait INPI d'une entreprise à partir de son SIREN ou SIRET.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentExtraitInpi(context.Background()).Siren(siren).Siret(siret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentExtraitInpi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentExtraitInpi`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentExtraitInpi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentExtraitInpiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **siret** | **string** | SIRET de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentExtraitPappers

> *os.File DocumentExtraitPappers(ctx).Siren(siren).Siret(siret).Execute()

Télécharge l'extrait Pappers d'une entreprise à partir de son SIREN ou SIRET.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentExtraitPappers(context.Background()).Siren(siren).Siret(siret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentExtraitPappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentExtraitPappers`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentExtraitPappers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentExtraitPappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **siret** | **string** | SIRET de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentScoringFinancier

> *os.File DocumentScoringFinancier(ctx).Siren(siren).Execute()

Télécharge le rapport financier d'une entreprise à partir de son SIREN.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentScoringFinancier(context.Background()).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentScoringFinancier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentScoringFinancier`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentScoringFinancier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentScoringFinancierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentScoringNonFinancier

> *os.File DocumentScoringNonFinancier(ctx).Siren(siren).Execute()

Télécharge le rapport non financier d'une entreprise à partir de son SIREN.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentScoringNonFinancier(context.Background()).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentScoringNonFinancier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentScoringNonFinancier`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentScoringNonFinancier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentScoringNonFinancierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentStatus

> *os.File DocumentStatus(ctx).Siren(siren).Siret(siret).Execute()

Télécharge les derniers statuts disponibles d'une entreprise à partir de son SIREN ou SIRET.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentStatus(context.Background()).Siren(siren).Siret(siret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentStatus`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siren** | **string** | SIREN de l&#39;entreprise | 
 **siret** | **string** | SIRET de l&#39;entreprise | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocumentTelechargement

> *os.File DocumentTelechargement(ctx).Token(token).Execute()

Télécharge un document PDF ou XLSX à partir de son token.



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
	token := "QTQ0MzA2MTg0MTIwMjAwNDA3RMOpY2lzaW9uKHMpIGRlIGwnYXNzb2Npw6kgdW5pcXVl" // string | Token du document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DocumentTelechargement(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentTelechargement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentTelechargement`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DocumentTelechargement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentTelechargementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token du document | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

