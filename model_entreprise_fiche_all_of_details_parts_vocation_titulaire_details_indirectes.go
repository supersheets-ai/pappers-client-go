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

// checks if the EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes{}

// EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes Détails des parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon indirecte.
type EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes struct {
	// Parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon indirecte par le biais d'une indivision, en pourcentage des parts totales.
	PourcentageEnIndivision *float32 `json:"pourcentage_en_indivision,omitempty"`
	// Parts dont le bénéficiaire effectif a vocation à devenir titulaire de façon indirecte par le biais d'une personne morale, en pourcentage des parts totales.
	PourcentageEnPersonneMorale *float32 `json:"pourcentage_en_personne_morale,omitempty"`
	DetailsEnIndivision *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision `json:"details_en_indivision,omitempty"`
	DetailsEnPersonneMorale *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale `json:"details_en_personne_morale,omitempty"`
}

// NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes instantiates a new EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes {
	this := EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes{}
	return &this
}

// NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesWithDefaults instantiates a new EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesWithDefaults() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes {
	this := EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes{}
	return &this
}

// GetPourcentageEnIndivision returns the PourcentageEnIndivision field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetPourcentageEnIndivision() float32 {
	if o == nil || IsNil(o.PourcentageEnIndivision) {
		var ret float32
		return ret
	}
	return *o.PourcentageEnIndivision
}

// GetPourcentageEnIndivisionOk returns a tuple with the PourcentageEnIndivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetPourcentageEnIndivisionOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentageEnIndivision) {
		return nil, false
	}
	return o.PourcentageEnIndivision, true
}

// HasPourcentageEnIndivision returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) HasPourcentageEnIndivision() bool {
	if o != nil && !IsNil(o.PourcentageEnIndivision) {
		return true
	}

	return false
}

// SetPourcentageEnIndivision gets a reference to the given float32 and assigns it to the PourcentageEnIndivision field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) SetPourcentageEnIndivision(v float32) {
	o.PourcentageEnIndivision = &v
}

// GetPourcentageEnPersonneMorale returns the PourcentageEnPersonneMorale field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetPourcentageEnPersonneMorale() float32 {
	if o == nil || IsNil(o.PourcentageEnPersonneMorale) {
		var ret float32
		return ret
	}
	return *o.PourcentageEnPersonneMorale
}

// GetPourcentageEnPersonneMoraleOk returns a tuple with the PourcentageEnPersonneMorale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetPourcentageEnPersonneMoraleOk() (*float32, bool) {
	if o == nil || IsNil(o.PourcentageEnPersonneMorale) {
		return nil, false
	}
	return o.PourcentageEnPersonneMorale, true
}

// HasPourcentageEnPersonneMorale returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) HasPourcentageEnPersonneMorale() bool {
	if o != nil && !IsNil(o.PourcentageEnPersonneMorale) {
		return true
	}

	return false
}

// SetPourcentageEnPersonneMorale gets a reference to the given float32 and assigns it to the PourcentageEnPersonneMorale field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) SetPourcentageEnPersonneMorale(v float32) {
	o.PourcentageEnPersonneMorale = &v
}

// GetDetailsEnIndivision returns the DetailsEnIndivision field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetDetailsEnIndivision() EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision {
	if o == nil || IsNil(o.DetailsEnIndivision) {
		var ret EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision
		return ret
	}
	return *o.DetailsEnIndivision
}

// GetDetailsEnIndivisionOk returns a tuple with the DetailsEnIndivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetDetailsEnIndivisionOk() (*EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision, bool) {
	if o == nil || IsNil(o.DetailsEnIndivision) {
		return nil, false
	}
	return o.DetailsEnIndivision, true
}

// HasDetailsEnIndivision returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) HasDetailsEnIndivision() bool {
	if o != nil && !IsNil(o.DetailsEnIndivision) {
		return true
	}

	return false
}

// SetDetailsEnIndivision gets a reference to the given EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision and assigns it to the DetailsEnIndivision field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) SetDetailsEnIndivision(v EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnIndivision) {
	o.DetailsEnIndivision = &v
}

// GetDetailsEnPersonneMorale returns the DetailsEnPersonneMorale field value if set, zero value otherwise.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetDetailsEnPersonneMorale() EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale {
	if o == nil || IsNil(o.DetailsEnPersonneMorale) {
		var ret EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale
		return ret
	}
	return *o.DetailsEnPersonneMorale
}

// GetDetailsEnPersonneMoraleOk returns a tuple with the DetailsEnPersonneMorale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) GetDetailsEnPersonneMoraleOk() (*EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale, bool) {
	if o == nil || IsNil(o.DetailsEnPersonneMorale) {
		return nil, false
	}
	return o.DetailsEnPersonneMorale, true
}

// HasDetailsEnPersonneMorale returns a boolean if a field has been set.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) HasDetailsEnPersonneMorale() bool {
	if o != nil && !IsNil(o.DetailsEnPersonneMorale) {
		return true
	}

	return false
}

// SetDetailsEnPersonneMorale gets a reference to the given EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale and assigns it to the DetailsEnPersonneMorale field.
func (o *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) SetDetailsEnPersonneMorale(v EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectesDetailsEnPersonneMorale) {
	o.DetailsEnPersonneMorale = &v
}

func (o EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PourcentageEnIndivision) {
		toSerialize["pourcentage_en_indivision"] = o.PourcentageEnIndivision
	}
	if !IsNil(o.PourcentageEnPersonneMorale) {
		toSerialize["pourcentage_en_personne_morale"] = o.PourcentageEnPersonneMorale
	}
	if !IsNil(o.DetailsEnIndivision) {
		toSerialize["details_en_indivision"] = o.DetailsEnIndivision
	}
	if !IsNil(o.DetailsEnPersonneMorale) {
		toSerialize["details_en_personne_morale"] = o.DetailsEnPersonneMorale
	}
	return toSerialize, nil
}

type NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes struct {
	value *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes
	isSet bool
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) Get() *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes {
	return v.value
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) Set(val *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes(val *EntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes {
	return &NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes{value: val, isSet: true}
}

func (v NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrepriseFicheAllOfDetailsPartsVocationTitulaireDetailsIndirectes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


