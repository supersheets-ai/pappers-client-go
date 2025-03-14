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

// checks if the Conformite200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conformite200Response{}

// Conformite200Response struct for Conformite200Response
type Conformite200Response struct {
	PersonnePolitiquementExposee *PersonnePolitiquementExposee `json:"personne_politiquement_exposee,omitempty"`
	// Vaut vrai si le bénéficiaire effectif est actuellement sous sanction. Uniquement présent si demandé dans les champs supplémentaires.
	SanctionsEnCours *bool `json:"sanctions_en_cours,omitempty"`
	// Liste des sanctions du bénéficiaire effectif. Uniquement présent si demandé dans les champs supplémentaires.
	Sanctions []Sanction `json:"sanctions,omitempty"`
}

// NewConformite200Response instantiates a new Conformite200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConformite200Response() *Conformite200Response {
	this := Conformite200Response{}
	return &this
}

// NewConformite200ResponseWithDefaults instantiates a new Conformite200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConformite200ResponseWithDefaults() *Conformite200Response {
	this := Conformite200Response{}
	return &this
}

// GetPersonnePolitiquementExposee returns the PersonnePolitiquementExposee field value if set, zero value otherwise.
func (o *Conformite200Response) GetPersonnePolitiquementExposee() PersonnePolitiquementExposee {
	if o == nil || IsNil(o.PersonnePolitiquementExposee) {
		var ret PersonnePolitiquementExposee
		return ret
	}
	return *o.PersonnePolitiquementExposee
}

// GetPersonnePolitiquementExposeeOk returns a tuple with the PersonnePolitiquementExposee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conformite200Response) GetPersonnePolitiquementExposeeOk() (*PersonnePolitiquementExposee, bool) {
	if o == nil || IsNil(o.PersonnePolitiquementExposee) {
		return nil, false
	}
	return o.PersonnePolitiquementExposee, true
}

// HasPersonnePolitiquementExposee returns a boolean if a field has been set.
func (o *Conformite200Response) HasPersonnePolitiquementExposee() bool {
	if o != nil && !IsNil(o.PersonnePolitiquementExposee) {
		return true
	}

	return false
}

// SetPersonnePolitiquementExposee gets a reference to the given PersonnePolitiquementExposee and assigns it to the PersonnePolitiquementExposee field.
func (o *Conformite200Response) SetPersonnePolitiquementExposee(v PersonnePolitiquementExposee) {
	o.PersonnePolitiquementExposee = &v
}

// GetSanctionsEnCours returns the SanctionsEnCours field value if set, zero value otherwise.
func (o *Conformite200Response) GetSanctionsEnCours() bool {
	if o == nil || IsNil(o.SanctionsEnCours) {
		var ret bool
		return ret
	}
	return *o.SanctionsEnCours
}

// GetSanctionsEnCoursOk returns a tuple with the SanctionsEnCours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conformite200Response) GetSanctionsEnCoursOk() (*bool, bool) {
	if o == nil || IsNil(o.SanctionsEnCours) {
		return nil, false
	}
	return o.SanctionsEnCours, true
}

// HasSanctionsEnCours returns a boolean if a field has been set.
func (o *Conformite200Response) HasSanctionsEnCours() bool {
	if o != nil && !IsNil(o.SanctionsEnCours) {
		return true
	}

	return false
}

// SetSanctionsEnCours gets a reference to the given bool and assigns it to the SanctionsEnCours field.
func (o *Conformite200Response) SetSanctionsEnCours(v bool) {
	o.SanctionsEnCours = &v
}

// GetSanctions returns the Sanctions field value if set, zero value otherwise.
func (o *Conformite200Response) GetSanctions() []Sanction {
	if o == nil || IsNil(o.Sanctions) {
		var ret []Sanction
		return ret
	}
	return o.Sanctions
}

// GetSanctionsOk returns a tuple with the Sanctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conformite200Response) GetSanctionsOk() ([]Sanction, bool) {
	if o == nil || IsNil(o.Sanctions) {
		return nil, false
	}
	return o.Sanctions, true
}

// HasSanctions returns a boolean if a field has been set.
func (o *Conformite200Response) HasSanctions() bool {
	if o != nil && !IsNil(o.Sanctions) {
		return true
	}

	return false
}

// SetSanctions gets a reference to the given []Sanction and assigns it to the Sanctions field.
func (o *Conformite200Response) SetSanctions(v []Sanction) {
	o.Sanctions = v
}

func (o Conformite200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conformite200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonnePolitiquementExposee) {
		toSerialize["personne_politiquement_exposee"] = o.PersonnePolitiquementExposee
	}
	if !IsNil(o.SanctionsEnCours) {
		toSerialize["sanctions_en_cours"] = o.SanctionsEnCours
	}
	if !IsNil(o.Sanctions) {
		toSerialize["sanctions"] = o.Sanctions
	}
	return toSerialize, nil
}

type NullableConformite200Response struct {
	value *Conformite200Response
	isSet bool
}

func (v NullableConformite200Response) Get() *Conformite200Response {
	return v.value
}

func (v *NullableConformite200Response) Set(val *Conformite200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableConformite200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableConformite200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConformite200Response(val *Conformite200Response) *NullableConformite200Response {
	return &NullableConformite200Response{value: val, isSet: true}
}

func (v NullableConformite200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConformite200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


