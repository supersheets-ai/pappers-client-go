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

// checks if the EntrepriseFicheAllOfDetailsSocieteDeGestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntrepriseFicheAllOfDetailsSocieteDeGestion{}

// EntrepriseFicheAllOfDetailsSocieteDeGestion Détails sur la société de gestion, le cas échéant.
type EntrepriseFicheAllOfDetailsSocieteDeGestion struct {
	// Nom de la société de gestion.
	Nom *string `json:"nom,omitempty"`
	// SIREN de la société de gestion.
	Siren *string `json:"siren,omitempty"`
	// Greffe de la société de gestion.
	Greffe *string `json:"greffe,omitempty"`
	// Adresse de la société de gestion.
	Adresse *string `json:"adresse,omitempty"`
	// Code postal de la société de gestion.
	CodePostal *string `json:"code_postal,omitempty"`
	// Ville de la société de gestion.
	Ville *string `json:"ville,omitempty"`
}

// NewEntrepriseFicheAllOfDetailsSocieteDeGestion instantiates a new EntrepriseFicheAllOfDetailsSocieteDeGestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrepriseFicheAllOfDetailsSocieteDeGestion() *EntrepriseFicheAllOfDetailsSocieteDeGestion {
	this := EntrepriseFicheAllOfDetailsSocieteDeGestion{}
	return &this
}

// NewEntrepriseFicheAllOfDetailsSocieteDeGestionWithDefaults instantiates a new EntrepriseFicheAllOfDetailsSocieteDeGestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrepriseFicheAllOfDetailsSocieteDeGestionWithDefaults() *EntrepriseFicheAllOfDetailsSocieteDeGestion {
	this := EntrepriseFicheAllOfDetailsSocieteDeGestion{}
	return &this
}

// GetNom returns the Nom field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetNom() string {
	if o == nil || IsNil(o.Nom) {
		var ret string
		return ret
	}
	return *o.Nom
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetNomOk() (*string, bool) {
	if o == nil || IsNil(o.Nom) {
		return nil, false
	}
	return o.Nom, true
}

// HasNom returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasNom() bool {
	if o != nil && !IsNil(o.Nom) {
		return true
	}

	return false
}

// SetNom gets a reference to the given string and assigns it to the Nom field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetNom(v string) {
	o.Nom = &v
}

// GetSiren returns the Siren field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetSiren() string {
	if o == nil || IsNil(o.Siren) {
		var ret string
		return ret
	}
	return *o.Siren
}

// GetSirenOk returns a tuple with the Siren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetSirenOk() (*string, bool) {
	if o == nil || IsNil(o.Siren) {
		return nil, false
	}
	return o.Siren, true
}

// HasSiren returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasSiren() bool {
	if o != nil && !IsNil(o.Siren) {
		return true
	}

	return false
}

// SetSiren gets a reference to the given string and assigns it to the Siren field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetSiren(v string) {
	o.Siren = &v
}

// GetGreffe returns the Greffe field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetGreffe() string {
	if o == nil || IsNil(o.Greffe) {
		var ret string
		return ret
	}
	return *o.Greffe
}

// GetGreffeOk returns a tuple with the Greffe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetGreffeOk() (*string, bool) {
	if o == nil || IsNil(o.Greffe) {
		return nil, false
	}
	return o.Greffe, true
}

// HasGreffe returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasGreffe() bool {
	if o != nil && !IsNil(o.Greffe) {
		return true
	}

	return false
}

// SetGreffe gets a reference to the given string and assigns it to the Greffe field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetGreffe(v string) {
	o.Greffe = &v
}

// GetAdresse returns the Adresse field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetAdresse() string {
	if o == nil || IsNil(o.Adresse) {
		var ret string
		return ret
	}
	return *o.Adresse
}

// GetAdresseOk returns a tuple with the Adresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetAdresseOk() (*string, bool) {
	if o == nil || IsNil(o.Adresse) {
		return nil, false
	}
	return o.Adresse, true
}

// HasAdresse returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasAdresse() bool {
	if o != nil && !IsNil(o.Adresse) {
		return true
	}

	return false
}

// SetAdresse gets a reference to the given string and assigns it to the Adresse field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetAdresse(v string) {
	o.Adresse = &v
}

// GetCodePostal returns the CodePostal field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetCodePostal() string {
	if o == nil || IsNil(o.CodePostal) {
		var ret string
		return ret
	}
	return *o.CodePostal
}

// GetCodePostalOk returns a tuple with the CodePostal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetCodePostalOk() (*string, bool) {
	if o == nil || IsNil(o.CodePostal) {
		return nil, false
	}
	return o.CodePostal, true
}

// HasCodePostal returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasCodePostal() bool {
	if o != nil && !IsNil(o.CodePostal) {
		return true
	}

	return false
}

// SetCodePostal gets a reference to the given string and assigns it to the CodePostal field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetCodePostal(v string) {
	o.CodePostal = &v
}

// GetVille returns the Ville field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetVille() string {
	if o == nil || IsNil(o.Ville) {
		var ret string
		return ret
	}
	return *o.Ville
}

// GetVilleOk returns a tuple with the Ville field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) GetVilleOk() (*string, bool) {
	if o == nil || IsNil(o.Ville) {
		return nil, false
	}
	return o.Ville, true
}

// HasVille returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) HasVille() bool {
	if o != nil && !IsNil(o.Ville) {
		return true
	}

	return false
}

// SetVille gets a reference to the given string and assigns it to the Ville field.
func (o *EntrepriseFicheAllOfDetailsSocieteDeGestion) SetVille(v string) {
	o.Ville = &v
}

func (o EntrepriseFicheAllOfDetailsSocieteDeGestion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrepriseFicheAllOfDetailsSocieteDeGestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nom) {
		toSerialize["nom"] = o.Nom
	}
	if !IsNil(o.Siren) {
		toSerialize["siren"] = o.Siren
	}
	if !IsNil(o.Greffe) {
		toSerialize["greffe"] = o.Greffe
	}
	if !IsNil(o.Adresse) {
		toSerialize["adresse"] = o.Adresse
	}
	if !IsNil(o.CodePostal) {
		toSerialize["code_postal"] = o.CodePostal
	}
	if !IsNil(o.Ville) {
		toSerialize["ville"] = o.Ville
	}
	return toSerialize, nil
}

type NullableEntrepriseFicheAllOfDetailsSocieteDeGestion struct {
	value *EntrepriseFicheAllOfDetailsSocieteDeGestion
	isSet bool
}

func (v NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) Get() *EntrepriseFicheAllOfDetailsSocieteDeGestion {
	return v.value
}

func (v *NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) Set(val *EntrepriseFicheAllOfDetailsSocieteDeGestion) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrepriseFicheAllOfDetailsSocieteDeGestion(val *EntrepriseFicheAllOfDetailsSocieteDeGestion) *NullableEntrepriseFicheAllOfDetailsSocieteDeGestion {
	return &NullableEntrepriseFicheAllOfDetailsSocieteDeGestion{value: val, isSet: true}
}

func (v NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrepriseFicheAllOfDetailsSocieteDeGestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


