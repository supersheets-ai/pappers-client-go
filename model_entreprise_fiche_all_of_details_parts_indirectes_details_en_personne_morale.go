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

// checks if the EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale{}

// EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale Détails des parts détenues de façon indirecte par le biais d'une personne morale par le bénéficiaire effectif.
type EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale struct {
	// Parts détenues de façon indirecte par le biais d'une personne morale en pleine propriété par le bénéficiaire effectif, en pourcentage des parts totales.
	PourcentagePleinePropriete *float32 `json:"pourcentage_pleine_propriete,omitempty"`
	// Parts détenues de façon indirecte par le biais d'une personne morale en nue propriété par le bénéficiaire effectif, en pourcentage des parts totales.
	PourcentageNuePropriete *float32 `json:"pourcentage_nue_propriete,omitempty"`
}

// NewEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale instantiates a new EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale() *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale {
	this := EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale{}
	return &this
}

// NewEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMoraleWithDefaults instantiates a new EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMoraleWithDefaults() *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale {
	this := EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale{}
	return &this
}

// GetPourcentagePleinePropriete returns the PourcentagePleinePropriete field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) GetPourcentagePleinePropriete() float32 {
	if o == nil || IsNil(o.PourcentagePleinePropriete) {
		var ret float32
		return ret
	}
	return *o.PourcentagePleinePropriete
}

// GetPourcentagePleineProprieteOk returns a tuple with the PourcentagePleinePropriete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) GetPourcentagePleineProprieteOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentagePleinePropriete) {
		return nil, false
	}
	return o.PourcentagePleinePropriete, true
}

// HasPourcentagePleinePropriete returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) HasPourcentagePleinePropriete() bool {
	if o != nil && !IsNil(o.PourcentagePleinePropriete) {
		return true
	}

	return false
}

// SetPourcentagePleinePropriete gets a reference to the given float32 and assigns it to the PourcentagePleinePropriete field.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) SetPourcentagePleinePropriete(v float32) {
	o.PourcentagePleinePropriete = &v
}

// GetPourcentageNuePropriete returns the PourcentageNuePropriete field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) GetPourcentageNuePropriete() float32 {
	if o == nil || IsNil(o.PourcentageNuePropriete) {
		var ret float32
		return ret
	}
	return *o.PourcentageNuePropriete
}

// GetPourcentageNueProprieteOk returns a tuple with the PourcentageNuePropriete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) GetPourcentageNueProprieteOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentageNuePropriete) {
		return nil, false
	}
	return o.PourcentageNuePropriete, true
}

// HasPourcentageNuePropriete returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) HasPourcentageNuePropriete() bool {
	if o != nil && !IsNil(o.PourcentageNuePropriete) {
		return true
	}

	return false
}

// SetPourcentageNuePropriete gets a reference to the given float32 and assigns it to the PourcentageNuePropriete field.
func (o *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) SetPourcentageNuePropriete(v float32) {
	o.PourcentageNuePropriete = &v
}

func (o EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PourcentagePleinePropriete) {
		toSerialize["pourcentage_pleine_propriete"] = o.PourcentagePleinePropriete
	}
	if !IsNil(o.PourcentageNuePropriete) {
		toSerialize["pourcentage_nue_propriete"] = o.PourcentageNuePropriete
	}
	return toSerialize, nil
}

type NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale struct {
	value *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale
	isSet bool
}

func (v NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) Get() *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale {
	return v.value
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) Set(val *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale(val *EntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) *NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale {
	return &NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale{value: val, isSet: true}
}

func (v NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsIndirectesDetailsEnPersonneMorale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


