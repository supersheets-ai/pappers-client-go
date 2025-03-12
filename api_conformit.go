/*
Pappers API

L'API de Pappers vous permet de récupérer des informations et documents sur les entreprises françaises à partir de leur numéro SIREN ou SIRET.  Vous devez indiquer votre clé d'API dans les requêtes, soit via le header `api-key`, soit (déconseillé) en utilisant le paramètre query `api_token`.  L'URL d'accès à l'API est https://api.pappers.fr/v2/  Lien vers la documentation de la V3 : https://www.pappers.fr/api/documentation/v3  Lien vers la documentation de l'API internationale : https://www.pappers.in/api/documentation  L'historique des modifications (changelog) est accessible à l'url suivante : https://www.pappers.fr/api/changelog 

API version: 2.16.0
Contact: support@pappers.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ConformitAPIService ConformitAPI service
type ConformitAPIService service

type ApiConformiteRequest struct {
	ctx context.Context
	ApiService *ConformitAPIService
	nom *string
	prenom *string
	dateDeNaissance *string
}

// Nom de la personne physique
func (r ApiConformiteRequest) Nom(nom string) ApiConformiteRequest {
	r.nom = &nom
	return r
}

// Prénom de la personne physique
func (r ApiConformiteRequest) Prenom(prenom string) ApiConformiteRequest {
	r.prenom = &prenom
	return r
}

// Date de naissance de la personne physique, au format AAAA-MM-JJ ou AAAA-MM
func (r ApiConformiteRequest) DateDeNaissance(dateDeNaissance string) ApiConformiteRequest {
	r.dateDeNaissance = &dateDeNaissance
	return r
}

func (r ApiConformiteRequest) Execute() (*Conformite200Response, *http.Response, error) {
	return r.ApiService.ConformiteExecute(r)
}

/*
Conformite Vérifie le statut de personne politiquement exposée et la présence de sanctions internationales pour une personne physique.

Cette route vérifie le statut de personne politiquement exposée et la présence de sanctions internationales pour une personne physique.

Pour vérifier le statut de dirigeants et bénéficiaires effectifs d'entreprises, vous pouvez directement utiliser la route `/entreprise` avec les champs supplémentaires `personne_politiquement_exposee` et `sanctions`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConformiteRequest
*/
func (a *ConformitAPIService) Conformite(ctx context.Context) ApiConformiteRequest {
	return ApiConformiteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Conformite200Response
func (a *ConformitAPIService) ConformiteExecute(r ApiConformiteRequest) (*Conformite200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Conformite200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConformitAPIService.Conformite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conformite/personne_physique"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nom == nil {
		return localVarReturnValue, nil, reportError("nom is required and must be specified")
	}
	if r.prenom == nil {
		return localVarReturnValue, nil, reportError("prenom is required and must be specified")
	}
	if r.dateDeNaissance == nil {
		return localVarReturnValue, nil, reportError("dateDeNaissance is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "nom", r.nom, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "prenom", r.prenom, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "date_de_naissance", r.dateDeNaissance, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
