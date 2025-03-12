/*
Pappers API

L'API de Pappers vous permet de récupérer des informations et documents sur les entreprises françaises à partir de leur numéro SIREN ou SIRET.  Vous devez indiquer votre clé d'API dans les requêtes, soit via le header `api-key`, soit (déconseillé) en utilisant le paramètre query `api_token`.  L'URL d'accès à l'API est https://api.pappers.fr/v2/  Lien vers la documentation de la V3 : https://www.pappers.fr/api/documentation/v3  Lien vers la documentation de l'API internationale : https://www.pappers.in/api/documentation  L'historique des modifications (changelog) est accessible à l'url suivante : https://www.pappers.fr/api/changelog 

API version: 2.16.0
Contact: support@pappers.fr
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SuiviJetons200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuiviJetons200Response{}

// SuiviJetons200Response Suivi d'utilisation des jetons.
type SuiviJetons200Response struct {
	// Le nombre de jetons mensuels initiaux de votre abonnement.
	JetonsAbonnement *float32 `json:"jetons_abonnement,omitempty"`
	// Le nombre de jetons mensuels de votre abonnement que vous avez utilisés
	JetonsAbonnementUtilises *float32 `json:"jetons_abonnement_utilises,omitempty"`
	// Le nombre de jetons pay as you go qu'il vous reste.
	JetonsPayAsYouGoRestants *float32 `json:"jetons_pay_as_you_go_restants,omitempty"`
}

// NewSuiviJetons200Response instantiates a new SuiviJetons200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuiviJetons200Response() *SuiviJetons200Response {
	this := SuiviJetons200Response{}
	return &this
}

// NewSuiviJetons200ResponseWithDefaults instantiates a new SuiviJetons200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuiviJetons200ResponseWithDefaults() *SuiviJetons200Response {
	this := SuiviJetons200Response{}
	return &this
}

// GetJetonsAbonnement returns the JetonsAbonnement field value if set, zero value otherwise.
func (o *SuiviJetons200Response) GetJetonsAbonnement() float32 {
	if o == nil || IsNil(o.JetonsAbonnement) {
		var ret float32
		return ret
	}
	return *o.JetonsAbonnement
}

// GetJetonsAbonnementOk returns a tuple with the JetonsAbonnement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuiviJetons200Response) GetJetonsAbonnementOk() (*float32, bool) {
	if o == nil || IsNil(o.JetonsAbonnement) {
		return nil, false
	}
	return o.JetonsAbonnement, true
}

// HasJetonsAbonnement returns a boolean if a field has been set.
func (o *SuiviJetons200Response) HasJetonsAbonnement() bool {
	if o != nil && !IsNil(o.JetonsAbonnement) {
		return true
	}

	return false
}

// SetJetonsAbonnement gets a reference to the given float32 and assigns it to the JetonsAbonnement field.
func (o *SuiviJetons200Response) SetJetonsAbonnement(v float32) {
	o.JetonsAbonnement = &v
}

// GetJetonsAbonnementUtilises returns the JetonsAbonnementUtilises field value if set, zero value otherwise.
func (o *SuiviJetons200Response) GetJetonsAbonnementUtilises() float32 {
	if o == nil || IsNil(o.JetonsAbonnementUtilises) {
		var ret float32
		return ret
	}
	return *o.JetonsAbonnementUtilises
}

// GetJetonsAbonnementUtilisesOk returns a tuple with the JetonsAbonnementUtilises field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuiviJetons200Response) GetJetonsAbonnementUtilisesOk() (*float32, bool) {
	if o == nil || IsNil(o.JetonsAbonnementUtilises) {
		return nil, false
	}
	return o.JetonsAbonnementUtilises, true
}

// HasJetonsAbonnementUtilises returns a boolean if a field has been set.
func (o *SuiviJetons200Response) HasJetonsAbonnementUtilises() bool {
	if o != nil && !IsNil(o.JetonsAbonnementUtilises) {
		return true
	}

	return false
}

// SetJetonsAbonnementUtilises gets a reference to the given float32 and assigns it to the JetonsAbonnementUtilises field.
func (o *SuiviJetons200Response) SetJetonsAbonnementUtilises(v float32) {
	o.JetonsAbonnementUtilises = &v
}

// GetJetonsPayAsYouGoRestants returns the JetonsPayAsYouGoRestants field value if set, zero value otherwise.
func (o *SuiviJetons200Response) GetJetonsPayAsYouGoRestants() float32 {
	if o == nil || IsNil(o.JetonsPayAsYouGoRestants) {
		var ret float32
		return ret
	}
	return *o.JetonsPayAsYouGoRestants
}

// GetJetonsPayAsYouGoRestantsOk returns a tuple with the JetonsPayAsYouGoRestants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuiviJetons200Response) GetJetonsPayAsYouGoRestantsOk() (*float32, bool) {
	if o == nil || IsNil(o.JetonsPayAsYouGoRestants) {
		return nil, false
	}
	return o.JetonsPayAsYouGoRestants, true
}

// HasJetonsPayAsYouGoRestants returns a boolean if a field has been set.
func (o *SuiviJetons200Response) HasJetonsPayAsYouGoRestants() bool {
	if o != nil && !IsNil(o.JetonsPayAsYouGoRestants) {
		return true
	}

	return false
}

// SetJetonsPayAsYouGoRestants gets a reference to the given float32 and assigns it to the JetonsPayAsYouGoRestants field.
func (o *SuiviJetons200Response) SetJetonsPayAsYouGoRestants(v float32) {
	o.JetonsPayAsYouGoRestants = &v
}

func (o SuiviJetons200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuiviJetons200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JetonsAbonnement) {
		toSerialize["jetons_abonnement"] = o.JetonsAbonnement
	}
	if !IsNil(o.JetonsAbonnementUtilises) {
		toSerialize["jetons_abonnement_utilises"] = o.JetonsAbonnementUtilises
	}
	if !IsNil(o.JetonsPayAsYouGoRestants) {
		toSerialize["jetons_pay_as_you_go_restants"] = o.JetonsPayAsYouGoRestants
	}
	return toSerialize, nil
}

type NullableSuiviJetons200Response struct {
	value *SuiviJetons200Response
	isSet bool
}

func (v NullableSuiviJetons200Response) Get() *SuiviJetons200Response {
	return v.value
}

func (v *NullableSuiviJetons200Response) Set(val *SuiviJetons200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSuiviJetons200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSuiviJetons200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuiviJetons200Response(val *SuiviJetons200Response) *NullableSuiviJetons200Response {
	return &NullableSuiviJetons200Response{value: val, isSet: true}
}

func (v NullableSuiviJetons200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuiviJetons200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


