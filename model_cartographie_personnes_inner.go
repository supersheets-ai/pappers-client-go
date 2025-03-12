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

// checks if the CartographiePersonnesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartographiePersonnesInner{}

// CartographiePersonnesInner struct for CartographiePersonnesInner
type CartographiePersonnesInner struct {
	// Un identifiant unique du noeud.
	Id *string `json:"id,omitempty"`
	// SIREN de l'entreprise.
	Prenom *string `json:"prenom,omitempty"`
	// Nom de l'entreprise.
	Nom *string `json:"nom,omitempty"`
	// Niveau du noeud. Le niveau 1 correspond aux dirigeants et bénéficiaires effectifs directement liés à l'entreprise recherchée. Le niveau 2 correspond aux autres.
	Niveau *int32 `json:"niveau,omitempty"`
}

// NewCartographiePersonnesInner instantiates a new CartographiePersonnesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartographiePersonnesInner() *CartographiePersonnesInner {
	this := CartographiePersonnesInner{}
	return &this
}

// NewCartographiePersonnesInnerWithDefaults instantiates a new CartographiePersonnesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartographiePersonnesInnerWithDefaults() *CartographiePersonnesInner {
	this := CartographiePersonnesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CartographiePersonnesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartographiePersonnesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CartographiePersonnesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CartographiePersonnesInner) SetId(v string) {
	o.Id = &v
}

// GetPrenom returns the Prenom field value if set, zero value otherwise.
func (o *CartographiePersonnesInner) GetPrenom() string {
	if o == nil || IsNil(o.Prenom) {
		var ret string
		return ret
	}
	return *o.Prenom
}

// GetPrenomOk returns a tuple with the Prenom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartographiePersonnesInner) GetPrenomOk() (*string, bool) {
	if o == nil || IsNil(o.Prenom) {
		return nil, false
	}
	return o.Prenom, true
}

// HasPrenom returns a boolean if a field has been set.
func (o *CartographiePersonnesInner) HasPrenom() bool {
	if o != nil && !IsNil(o.Prenom) {
		return true
	}

	return false
}

// SetPrenom gets a reference to the given string and assigns it to the Prenom field.
func (o *CartographiePersonnesInner) SetPrenom(v string) {
	o.Prenom = &v
}

// GetNom returns the Nom field value if set, zero value otherwise.
func (o *CartographiePersonnesInner) GetNom() string {
	if o == nil || IsNil(o.Nom) {
		var ret string
		return ret
	}
	return *o.Nom
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartographiePersonnesInner) GetNomOk() (*string, bool) {
	if o == nil || IsNil(o.Nom) {
		return nil, false
	}
	return o.Nom, true
}

// HasNom returns a boolean if a field has been set.
func (o *CartographiePersonnesInner) HasNom() bool {
	if o != nil && !IsNil(o.Nom) {
		return true
	}

	return false
}

// SetNom gets a reference to the given string and assigns it to the Nom field.
func (o *CartographiePersonnesInner) SetNom(v string) {
	o.Nom = &v
}

// GetNiveau returns the Niveau field value if set, zero value otherwise.
func (o *CartographiePersonnesInner) GetNiveau() int32 {
	if o == nil || IsNil(o.Niveau) {
		var ret int32
		return ret
	}
	return *o.Niveau
}

// GetNiveauOk returns a tuple with the Niveau field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartographiePersonnesInner) GetNiveauOk() (*int32, bool) {
	if o == nil || IsNil(o.Niveau) {
		return nil, false
	}
	return o.Niveau, true
}

// HasNiveau returns a boolean if a field has been set.
func (o *CartographiePersonnesInner) HasNiveau() bool {
	if o != nil && !IsNil(o.Niveau) {
		return true
	}

	return false
}

// SetNiveau gets a reference to the given int32 and assigns it to the Niveau field.
func (o *CartographiePersonnesInner) SetNiveau(v int32) {
	o.Niveau = &v
}

func (o CartographiePersonnesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartographiePersonnesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Prenom) {
		toSerialize["prenom"] = o.Prenom
	}
	if !IsNil(o.Nom) {
		toSerialize["nom"] = o.Nom
	}
	if !IsNil(o.Niveau) {
		toSerialize["niveau"] = o.Niveau
	}
	return toSerialize, nil
}

type NullableCartographiePersonnesInner struct {
	value *CartographiePersonnesInner
	isSet bool
}

func (v NullableCartographiePersonnesInner) Get() *CartographiePersonnesInner {
	return v.value
}

func (v *NullableCartographiePersonnesInner) Set(val *CartographiePersonnesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCartographiePersonnesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCartographiePersonnesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartographiePersonnesInner(val *CartographiePersonnesInner) *NullableCartographiePersonnesInner {
	return &NullableCartographiePersonnesInner{value: val, isSet: true}
}

func (v NullableCartographiePersonnesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartographiePersonnesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


