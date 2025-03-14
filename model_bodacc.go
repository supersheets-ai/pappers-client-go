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

// checks if the Bodacc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bodacc{}

// Bodacc struct for Bodacc
type Bodacc struct {
	// Numéro de parution de la publication.
	NumeroParution *string `json:"numero_parution,omitempty"`
	// Date de la publication, au format AAAA-MM-JJ.
	Date *string `json:"date,omitempty"`
	// Numéro d'annonce de la publication.
	NumeroAnnonce *string `json:"numero_annonce,omitempty"`
	// Bodacc de la publication (A, B ou C).
	Bodacc *string `json:"bodacc,omitempty"`
	// Type de la publication parmi la liste suivante : Création, Immatriculation, Modification, Vente, Achat, Radiation, Procédure collective, Dépôt des comptes.
	Type *string `json:"type,omitempty"`
	// Greffe de publication.
	Greffe *string `json:"greffe,omitempty"`
}

// NewBodacc instantiates a new Bodacc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodacc() *Bodacc {
	this := Bodacc{}
	return &this
}

// NewBodaccWithDefaults instantiates a new Bodacc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodaccWithDefaults() *Bodacc {
	this := Bodacc{}
	return &this
}

// GetNumeroParution returns the NumeroParution field value if set, zero value otherwise.
func (o *Bodacc) GetNumeroParution() string {
	if o == nil || IsNil(o.NumeroParution) {
		var ret string
		return ret
	}
	return *o.NumeroParution
}

// GetNumeroParutionOk returns a tuple with the NumeroParution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetNumeroParutionOk() (*string, bool) {
	if o == nil || IsNil(o.NumeroParution) {
		return nil, false
	}
	return o.NumeroParution, true
}

// HasNumeroParution returns a boolean if a field has been set.
func (o *Bodacc) HasNumeroParution() bool {
	if o != nil && !IsNil(o.NumeroParution) {
		return true
	}

	return false
}

// SetNumeroParution gets a reference to the given string and assigns it to the NumeroParution field.
func (o *Bodacc) SetNumeroParution(v string) {
	o.NumeroParution = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Bodacc) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Bodacc) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Bodacc) SetDate(v string) {
	o.Date = &v
}

// GetNumeroAnnonce returns the NumeroAnnonce field value if set, zero value otherwise.
func (o *Bodacc) GetNumeroAnnonce() string {
	if o == nil || IsNil(o.NumeroAnnonce) {
		var ret string
		return ret
	}
	return *o.NumeroAnnonce
}

// GetNumeroAnnonceOk returns a tuple with the NumeroAnnonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetNumeroAnnonceOk() (*string, bool) {
	if o == nil || IsNil(o.NumeroAnnonce) {
		return nil, false
	}
	return o.NumeroAnnonce, true
}

// HasNumeroAnnonce returns a boolean if a field has been set.
func (o *Bodacc) HasNumeroAnnonce() bool {
	if o != nil && !IsNil(o.NumeroAnnonce) {
		return true
	}

	return false
}

// SetNumeroAnnonce gets a reference to the given string and assigns it to the NumeroAnnonce field.
func (o *Bodacc) SetNumeroAnnonce(v string) {
	o.NumeroAnnonce = &v
}

// GetBodacc returns the Bodacc field value if set, zero value otherwise.
func (o *Bodacc) GetBodacc() string {
	if o == nil || IsNil(o.Bodacc) {
		var ret string
		return ret
	}
	return *o.Bodacc
}

// GetBodaccOk returns a tuple with the Bodacc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetBodaccOk() (*string, bool) {
	if o == nil || IsNil(o.Bodacc) {
		return nil, false
	}
	return o.Bodacc, true
}

// HasBodacc returns a boolean if a field has been set.
func (o *Bodacc) HasBodacc() bool {
	if o != nil && !IsNil(o.Bodacc) {
		return true
	}

	return false
}

// SetBodacc gets a reference to the given string and assigns it to the Bodacc field.
func (o *Bodacc) SetBodacc(v string) {
	o.Bodacc = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Bodacc) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Bodacc) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Bodacc) SetType(v string) {
	o.Type = &v
}

// GetGreffe returns the Greffe field value if set, zero value otherwise.
func (o *Bodacc) GetGreffe() string {
	if o == nil || IsNil(o.Greffe) {
		var ret string
		return ret
	}
	return *o.Greffe
}

// GetGreffeOk returns a tuple with the Greffe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bodacc) GetGreffeOk() (*string, bool) {
	if o == nil || IsNil(o.Greffe) {
		return nil, false
	}
	return o.Greffe, true
}

// HasGreffe returns a boolean if a field has been set.
func (o *Bodacc) HasGreffe() bool {
	if o != nil && !IsNil(o.Greffe) {
		return true
	}

	return false
}

// SetGreffe gets a reference to the given string and assigns it to the Greffe field.
func (o *Bodacc) SetGreffe(v string) {
	o.Greffe = &v
}

func (o Bodacc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bodacc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumeroParution) {
		toSerialize["numero_parution"] = o.NumeroParution
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.NumeroAnnonce) {
		toSerialize["numero_annonce"] = o.NumeroAnnonce
	}
	if !IsNil(o.Bodacc) {
		toSerialize["bodacc"] = o.Bodacc
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Greffe) {
		toSerialize["greffe"] = o.Greffe
	}
	return toSerialize, nil
}

type NullableBodacc struct {
	value *Bodacc
	isSet bool
}

func (v NullableBodacc) Get() *Bodacc {
	return v.value
}

func (v *NullableBodacc) Set(val *Bodacc) {
	v.value = val
	v.isSet = true
}

func (v NullableBodacc) IsSet() bool {
	return v.isSet
}

func (v *NullableBodacc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodacc(val *Bodacc) *NullableBodacc {
	return &NullableBodacc{value: val, isSet: true}
}

func (v NullableBodacc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodacc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


