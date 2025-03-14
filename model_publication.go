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

// checks if the Publication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Publication{}

// Publication struct for Publication
type Publication struct {
	// Type de publication
	Type *string `json:"type,omitempty"`
	// Date de publication, au format AAAA-MM-JJ.
	Date *string `json:"date,omitempty"`
	// Contenu de la publication, avec les mentions correspondant à la recherche encerclée par une balise HTML <em>.
	Contenu *string `json:"contenu,omitempty"`
}

// NewPublication instantiates a new Publication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublication() *Publication {
	this := Publication{}
	return &this
}

// NewPublicationWithDefaults instantiates a new Publication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicationWithDefaults() *Publication {
	this := Publication{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Publication) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Publication) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Publication) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Publication) SetType(v string) {
	o.Type = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Publication) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Publication) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Publication) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Publication) SetDate(v string) {
	o.Date = &v
}

// GetContenu returns the Contenu field value if set, zero value otherwise.
func (o *Publication) GetContenu() string {
	if o == nil || IsNil(o.Contenu) {
		var ret string
		return ret
	}
	return *o.Contenu
}

// GetContenuOk returns a tuple with the Contenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Publication) GetContenuOk() (*string, bool) {
	if o == nil || IsNil(o.Contenu) {
		return nil, false
	}
	return o.Contenu, true
}

// HasContenu returns a boolean if a field has been set.
func (o *Publication) HasContenu() bool {
	if o != nil && !IsNil(o.Contenu) {
		return true
	}

	return false
}

// SetContenu gets a reference to the given string and assigns it to the Contenu field.
func (o *Publication) SetContenu(v string) {
	o.Contenu = &v
}

func (o Publication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Publication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Contenu) {
		toSerialize["contenu"] = o.Contenu
	}
	return toSerialize, nil
}

type NullablePublication struct {
	value *Publication
	isSet bool
}

func (v NullablePublication) Get() *Publication {
	return v.value
}

func (v *NullablePublication) Set(val *Publication) {
	v.value = val
	v.isSet = true
}

func (v NullablePublication) IsSet() bool {
	return v.isSet
}

func (v *NullablePublication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublication(val *Publication) *NullablePublication {
	return &NullablePublication{value: val, isSet: true}
}

func (v NullablePublication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


