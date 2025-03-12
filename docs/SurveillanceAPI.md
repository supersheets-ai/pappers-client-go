# \SurveillanceAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SurveillanceDirigeant**](SurveillanceAPI.md#SurveillanceDirigeant) | **Post** /liste/ | Ajoute un (ou plusieurs) dirigeant(s) à une liste de dirigeants.
[**SurveillanceEntreprise**](SurveillanceAPI.md#SurveillanceEntreprise) | **Post** /liste | Ajoute une (ou plusieurs) entreprise(s) à une liste d&#39;entreprises.
[**SurveillanceListeInformations**](SurveillanceAPI.md#SurveillanceListeInformations) | **Post** /liste-informations | Ajoute des informations à une (ou plusieurs) notification(s).
[**SurveillanceNotificationsDelete**](SurveillanceAPI.md#SurveillanceNotificationsDelete) | **Delete** /liste/ | Supprime une (ou plusieurs) notification(s) d&#39;une liste.



## SurveillanceDirigeant

> SurveillanceDirigeant200Response SurveillanceDirigeant(ctx).IdListe(idListe).SurveillanceDirigeantRequestInner(surveillanceDirigeantRequestInner).Execute()

Ajoute un (ou plusieurs) dirigeant(s) à une liste de dirigeants.



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
	idListe := "votre_identifiant_ici" // string | Identifiant unique de votre liste de surveillance de dirigeants
	surveillanceDirigeantRequestInner := []openapiclient.SurveillanceDirigeantRequestInner{*openapiclient.NewSurveillanceDirigeantRequestInner("Siren_example")} // []SurveillanceDirigeantRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveillanceAPI.SurveillanceDirigeant(context.Background()).IdListe(idListe).SurveillanceDirigeantRequestInner(surveillanceDirigeantRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveillanceAPI.SurveillanceDirigeant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveillanceDirigeant`: SurveillanceDirigeant200Response
	fmt.Fprintf(os.Stdout, "Response from `SurveillanceAPI.SurveillanceDirigeant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveillanceDirigeantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idListe** | **string** | Identifiant unique de votre liste de surveillance de dirigeants | 
 **surveillanceDirigeantRequestInner** | [**[]SurveillanceDirigeantRequestInner**](SurveillanceDirigeantRequestInner.md) |  | 

### Return type

[**SurveillanceDirigeant200Response**](SurveillanceDirigeant200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveillanceEntreprise

> SurveillanceEntreprise200Response SurveillanceEntreprise(ctx).IdListe(idListe).SurveillanceEntrepriseRequestInner(surveillanceEntrepriseRequestInner).Execute()

Ajoute une (ou plusieurs) entreprise(s) à une liste d'entreprises.



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
	idListe := "votre_identifiant_ici" // string | Identifiant unique de votre liste de surveillance d'entreprises
	surveillanceEntrepriseRequestInner := []openapiclient.SurveillanceEntrepriseRequestInner{*openapiclient.NewSurveillanceEntrepriseRequestInner("443061841")} // []SurveillanceEntrepriseRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveillanceAPI.SurveillanceEntreprise(context.Background()).IdListe(idListe).SurveillanceEntrepriseRequestInner(surveillanceEntrepriseRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveillanceAPI.SurveillanceEntreprise``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveillanceEntreprise`: SurveillanceEntreprise200Response
	fmt.Fprintf(os.Stdout, "Response from `SurveillanceAPI.SurveillanceEntreprise`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveillanceEntrepriseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idListe** | **string** | Identifiant unique de votre liste de surveillance d&#39;entreprises | 
 **surveillanceEntrepriseRequestInner** | [**[]SurveillanceEntrepriseRequestInner**](SurveillanceEntrepriseRequestInner.md) |  | 

### Return type

[**SurveillanceEntreprise200Response**](SurveillanceEntreprise200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveillanceListeInformations

> SurveillanceListeInformations(ctx).IdListe(idListe).SurveillanceListeInformationsRequest(surveillanceListeInformationsRequest).Execute()

Ajoute des informations à une (ou plusieurs) notification(s).



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
	idListe := "votre_identifiant_ici" // string | Identifiant unique de votre liste de surveillance d'entreprises
	surveillanceListeInformationsRequest := *openapiclient.NewSurveillanceListeInformationsRequest() // SurveillanceListeInformationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveillanceAPI.SurveillanceListeInformations(context.Background()).IdListe(idListe).SurveillanceListeInformationsRequest(surveillanceListeInformationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveillanceAPI.SurveillanceListeInformations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveillanceListeInformationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idListe** | **string** | Identifiant unique de votre liste de surveillance d&#39;entreprises | 
 **surveillanceListeInformationsRequest** | [**SurveillanceListeInformationsRequest**](SurveillanceListeInformationsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveillanceNotificationsDelete

> SurveillanceNotificationsDelete200Response SurveillanceNotificationsDelete(ctx).IdListe(idListe).Siren(siren).Id(id).SupprimerTotalite(supprimerTotalite).Execute()

Supprime une (ou plusieurs) notification(s) d'une liste.



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
	idListe := "votre_identifiant_ici" // string | Identifiant unique de votre liste de surveillance
	siren := "443061841,950370973" // string | Liste des sirens des notifications à supprimer, séparés par une virgule (optional)
	id := "ecfca3e5fb1ce07bb3fdfe8760a28df5aa617ed4951b7f61,a67b9de66189ba98eb4aede09f4b2b565b0f18fe694da7bc" // string | Liste des ids des notifications à supprimer, séparés par une virgule (optional)
	supprimerTotalite := true // bool | Suppression de toutes les notifications de la liste (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveillanceAPI.SurveillanceNotificationsDelete(context.Background()).IdListe(idListe).Siren(siren).Id(id).SupprimerTotalite(supprimerTotalite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveillanceAPI.SurveillanceNotificationsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveillanceNotificationsDelete`: SurveillanceNotificationsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `SurveillanceAPI.SurveillanceNotificationsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSurveillanceNotificationsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idListe** | **string** | Identifiant unique de votre liste de surveillance | 
 **siren** | **string** | Liste des sirens des notifications à supprimer, séparés par une virgule | 
 **id** | **string** | Liste des ids des notifications à supprimer, séparés par une virgule | 
 **supprimerTotalite** | **bool** | Suppression de toutes les notifications de la liste | 

### Return type

[**SurveillanceNotificationsDelete200Response**](SurveillanceNotificationsDelete200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

