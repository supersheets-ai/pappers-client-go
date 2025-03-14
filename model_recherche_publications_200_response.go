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

// checks if the RecherchePublications200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecherchePublications200Response{}

// RecherchePublications200Response struct for RecherchePublications200Response
type RecherchePublications200Response struct {
	// Liste des publications qui correspondent à la recherche.
	Resultats []RecherchePublications200ResponseResultatsInner `json:"resultats,omitempty"`
	// Nombre de publications qui correspondent à la recherche.
	Total *int32 `json:"total,omitempty"`
	// Page actuelle.
	Page *int32 `json:"page,omitempty"`
}

// NewRecherchePublications200Response instantiates a new RecherchePublications200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecherchePublications200Response() *RecherchePublications200Response {
	this := RecherchePublications200Response{}
	return &this
}

// NewRecherchePublications200ResponseWithDefaults instantiates a new RecherchePublications200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecherchePublications200ResponseWithDefaults() *RecherchePublications200Response {
	this := RecherchePublications200Response{}
	return &this
}

// GetResultats returns the Resultats field value if set, zero value otherwise.
func (o *RecherchePublications200Response) GetResultats() []RecherchePublications200ResponseResultatsInner {
	if o == nil || IsNil(o.Resultats) {
		var ret []RecherchePublications200ResponseResultatsInner
		return ret
	}
	return o.Resultats
}

// GetResultatsOk returns a tuple with the Resultats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecherchePublications200Response) GetResultatsOk() ([]RecherchePublications200ResponseResultatsInner, bool) {
	if o == nil || IsNil(o.Resultats) {
		return nil, false
	}
	return o.Resultats, true
}

// HasResultats returns a boolean if a field has been set.
func (o *RecherchePublications200Response) HasResultats() bool {
	if o != nil && !IsNil(o.Resultats) {
		return true
	}

	return false
}

// SetResultats gets a reference to the given []RecherchePublications200ResponseResultatsInner and assigns it to the Resultats field.
func (o *RecherchePublications200Response) SetResultats(v []RecherchePublications200ResponseResultatsInner) {
	o.Resultats = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RecherchePublications200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecherchePublications200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RecherchePublications200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *RecherchePublications200Response) SetTotal(v int32) {
	o.Total = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *RecherchePublications200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecherchePublications200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *RecherchePublications200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *RecherchePublications200Response) SetPage(v int32) {
	o.Page = &v
}

func (o RecherchePublications200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecherchePublications200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resultats) {
		toSerialize["resultats"] = o.Resultats
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableRecherchePublications200Response struct {
	value *RecherchePublications200Response
	isSet bool
}

func (v NullableRecherchePublications200Response) Get() *RecherchePublications200Response {
	return v.value
}

func (v *NullableRecherchePublications200Response) Set(val *RecherchePublications200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRecherchePublications200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRecherchePublications200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecherchePublications200Response(val *RecherchePublications200Response) *NullableRecherchePublications200Response {
	return &NullableRecherchePublications200Response{value: val, isSet: true}
}

func (v NullableRecherchePublications200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecherchePublications200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


