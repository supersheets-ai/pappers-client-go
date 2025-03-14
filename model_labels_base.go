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

// checks if the LabelsBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelsBase{}

// LabelsBase struct for LabelsBase
type LabelsBase struct {
	// Nom du label.
	Nom *string `json:"nom,omitempty"`
	// Label RGE seulement : Liste des certificats.
	Certificats []string `json:"certificats,omitempty"`
	// Label QUALIOPI seulement : Liste des spécialités.
	Specialites []string `json:"specialites,omitempty"`
	// Label EGALITE seulement : Liste des notes.
	Notes []string `json:"notes,omitempty"`
	// Label ORIAS et CCI seulement : Numéro d'immatriculation ORIAS ou CCI. Uniquement présent si demandé avec le champ supplémentaire `label:orias` ou `label:cci`.
	NumeroImmatriculation *string `json:"numero_immatriculation,omitempty"`
	// Label ORIAS seulement : Liste des inscriptions ORIAS. Uniquement présent si demandé avec le champ supplémentaire `label:orias`.
	Inscriptions []LabelsBaseInscriptionsInner `json:"inscriptions,omitempty"`
	// Label CCI seulement : Liste des mentions.
	Mentions []string `json:"mentions,omitempty"`
}

// NewLabelsBase instantiates a new LabelsBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsBase() *LabelsBase {
	this := LabelsBase{}
	return &this
}

// NewLabelsBaseWithDefaults instantiates a new LabelsBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsBaseWithDefaults() *LabelsBase {
	this := LabelsBase{}
	return &this
}

// GetNom returns the Nom field value if set, zero value otherwise.
func (o *LabelsBase) GetNom() string {
	if o == nil || IsNil(o.Nom) {
		var ret string
		return ret
	}
	return *o.Nom
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetNomOk() (*string, bool) {
	if o == nil || IsNil(o.Nom) {
		return nil, false
	}
	return o.Nom, true
}

// HasNom returns a boolean if a field has been set.
func (o *LabelsBase) HasNom() bool {
	if o != nil && !IsNil(o.Nom) {
		return true
	}

	return false
}

// SetNom gets a reference to the given string and assigns it to the Nom field.
func (o *LabelsBase) SetNom(v string) {
	o.Nom = &v
}

// GetCertificats returns the Certificats field value if set, zero value otherwise.
func (o *LabelsBase) GetCertificats() []string {
	if o == nil || IsNil(o.Certificats) {
		var ret []string
		return ret
	}
	return o.Certificats
}

// GetCertificatsOk returns a tuple with the Certificats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetCertificatsOk() ([]string, bool) {
	if o == nil || IsNil(o.Certificats) {
		return nil, false
	}
	return o.Certificats, true
}

// HasCertificats returns a boolean if a field has been set.
func (o *LabelsBase) HasCertificats() bool {
	if o != nil && !IsNil(o.Certificats) {
		return true
	}

	return false
}

// SetCertificats gets a reference to the given []string and assigns it to the Certificats field.
func (o *LabelsBase) SetCertificats(v []string) {
	o.Certificats = v
}

// GetSpecialites returns the Specialites field value if set, zero value otherwise.
func (o *LabelsBase) GetSpecialites() []string {
	if o == nil || IsNil(o.Specialites) {
		var ret []string
		return ret
	}
	return o.Specialites
}

// GetSpecialitesOk returns a tuple with the Specialites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetSpecialitesOk() ([]string, bool) {
	if o == nil || IsNil(o.Specialites) {
		return nil, false
	}
	return o.Specialites, true
}

// HasSpecialites returns a boolean if a field has been set.
func (o *LabelsBase) HasSpecialites() bool {
	if o != nil && !IsNil(o.Specialites) {
		return true
	}

	return false
}

// SetSpecialites gets a reference to the given []string and assigns it to the Specialites field.
func (o *LabelsBase) SetSpecialites(v []string) {
	o.Specialites = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *LabelsBase) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *LabelsBase) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *LabelsBase) SetNotes(v []string) {
	o.Notes = v
}

// GetNumeroImmatriculation returns the NumeroImmatriculation field value if set, zero value otherwise.
func (o *LabelsBase) GetNumeroImmatriculation() string {
	if o == nil || IsNil(o.NumeroImmatriculation) {
		var ret string
		return ret
	}
	return *o.NumeroImmatriculation
}

// GetNumeroImmatriculationOk returns a tuple with the NumeroImmatriculation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetNumeroImmatriculationOk() (*string, bool) {
	if o == nil || IsNil(o.NumeroImmatriculation) {
		return nil, false
	}
	return o.NumeroImmatriculation, true
}

// HasNumeroImmatriculation returns a boolean if a field has been set.
func (o *LabelsBase) HasNumeroImmatriculation() bool {
	if o != nil && !IsNil(o.NumeroImmatriculation) {
		return true
	}

	return false
}

// SetNumeroImmatriculation gets a reference to the given string and assigns it to the NumeroImmatriculation field.
func (o *LabelsBase) SetNumeroImmatriculation(v string) {
	o.NumeroImmatriculation = &v
}

// GetInscriptions returns the Inscriptions field value if set, zero value otherwise.
func (o *LabelsBase) GetInscriptions() []LabelsBaseInscriptionsInner {
	if o == nil || IsNil(o.Inscriptions) {
		var ret []LabelsBaseInscriptionsInner
		return ret
	}
	return o.Inscriptions
}

// GetInscriptionsOk returns a tuple with the Inscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetInscriptionsOk() ([]LabelsBaseInscriptionsInner, bool) {
	if o == nil || IsNil(o.Inscriptions) {
		return nil, false
	}
	return o.Inscriptions, true
}

// HasInscriptions returns a boolean if a field has been set.
func (o *LabelsBase) HasInscriptions() bool {
	if o != nil && !IsNil(o.Inscriptions) {
		return true
	}

	return false
}

// SetInscriptions gets a reference to the given []LabelsBaseInscriptionsInner and assigns it to the Inscriptions field.
func (o *LabelsBase) SetInscriptions(v []LabelsBaseInscriptionsInner) {
	o.Inscriptions = v
}

// GetMentions returns the Mentions field value if set, zero value otherwise.
func (o *LabelsBase) GetMentions() []string {
	if o == nil || IsNil(o.Mentions) {
		var ret []string
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsBase) GetMentionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *LabelsBase) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given []string and assigns it to the Mentions field.
func (o *LabelsBase) SetMentions(v []string) {
	o.Mentions = v
}

func (o LabelsBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelsBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nom) {
		toSerialize["nom"] = o.Nom
	}
	if !IsNil(o.Certificats) {
		toSerialize["certificats"] = o.Certificats
	}
	if !IsNil(o.Specialites) {
		toSerialize["specialites"] = o.Specialites
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.NumeroImmatriculation) {
		toSerialize["numero_immatriculation"] = o.NumeroImmatriculation
	}
	if !IsNil(o.Inscriptions) {
		toSerialize["inscriptions"] = o.Inscriptions
	}
	if !IsNil(o.Mentions) {
		toSerialize["mentions"] = o.Mentions
	}
	return toSerialize, nil
}

type NullableLabelsBase struct {
	value *LabelsBase
	isSet bool
}

func (v NullableLabelsBase) Get() *LabelsBase {
	return v.value
}

func (v *NullableLabelsBase) Set(val *LabelsBase) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsBase) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsBase(val *LabelsBase) *NullableLabelsBase {
	return &NullableLabelsBase{value: val, isSet: true}
}

func (v NullableLabelsBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


