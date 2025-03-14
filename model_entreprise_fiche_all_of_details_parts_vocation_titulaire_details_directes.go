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

// checks if the EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes{}

// EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes Détails des parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon directe.
type EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes struct {
	// Parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon directe en pleine propriété, en pourcentage des parts totales.
	PourcentagePleinePropriete *float32 `json:"pourcentage_pleine_propriete,omitempty"`
	// Parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon directe en nue propriété, en pourcentage des parts totales.
	PourcentageNuePropriete *float32 `json:"pourcentage_nue_propriete,omitempty"`
}

// NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes instantiates a new EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes {
	this := EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes{}
	return &this
}

// NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectesWithDefaults instantiates a new EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectesWithDefaults() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes {
	this := EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes{}
	return &this
}

// GetPourcentagePleinePropriete returns the PourcentagePleinePropriete field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) GetPourcentagePleinePropriete() float32 {
	if o == nil || IsNil(o.PourcentagePleinePropriete) {
		var ret float32
		return ret
	}
	return *o.PourcentagePleinePropriete
}

// GetPourcentagePleineProprieteOk returns a tuple with the PourcentagePleinePropriete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) GetPourcentagePleineProprieteOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentagePleinePropriete) {
		return nil, false
	}
	return o.PourcentagePleinePropriete, true
}

// HasPourcentagePleinePropriete returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) HasPourcentagePleinePropriete() bool {
	if o != nil && !IsNil(o.PourcentagePleinePropriete) {
		return true
	}

	return false
}

// SetPourcentagePleinePropriete gets a reference to the given float32 and assigns it to the PourcentagePleinePropriete field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) SetPourcentagePleinePropriete(v float32) {
	o.PourcentagePleinePropriete = &v
}

// GetPourcentageNuePropriete returns the PourcentageNuePropriete field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) GetPourcentageNuePropriete() float32 {
	if o == nil || IsNil(o.PourcentageNuePropriete) {
		var ret float32
		return ret
	}
	return *o.PourcentageNuePropriete
}

// GetPourcentageNueProprieteOk returns a tuple with the PourcentageNuePropriete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) GetPourcentageNueProprieteOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentageNuePropriete) {
		return nil, false
	}
	return o.PourcentageNuePropriete, true
}

// HasPourcentageNuePropriete returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) HasPourcentageNuePropriete() bool {
	if o != nil && !IsNil(o.PourcentageNuePropriete) {
		return true
	}

	return false
}

// SetPourcentageNuePropriete gets a reference to the given float32 and assigns it to the PourcentageNuePropriete field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) SetPourcentageNuePropriete(v float32) {
	o.PourcentageNuePropriete = &v
}

func (o EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PourcentagePleinePropriete) {
		toSerialize["pourcentage_pleine_propriete"] = o.PourcentagePleinePropriete
	}
	if !IsNil(o.PourcentageNuePropriete) {
		toSerialize["pourcentage_nue_propriete"] = o.PourcentageNuePropriete
	}
	return toSerialize, nil
}

type NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes struct {
	value *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes
	isSet bool
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) Get() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes {
	return v.value
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) Set(val *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes(val *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes {
	return &NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes{value: val, isSet: true}
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsDirectes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


