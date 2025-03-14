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

// checks if the AssociationAdresseGestionnaire type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociationAdresseGestionnaire{}

// AssociationAdresseGestionnaire struct for AssociationAdresseGestionnaire
type AssociationAdresseGestionnaire struct {
	// Nom du gestionnaire de l'association.
	Gestionnaire *string `json:"gestionnaire,omitempty"`
	// Code postal du gestionnaire.
	CodePostal *string `json:"code_postal,omitempty"`
	// Ville du gestionnaire.
	Ville *string `json:"ville,omitempty"`
	// Complément de distribution du gestionnaire.
	Distribution *string `json:"distribution,omitempty"`
	// Adresse complète du gestionnaire.
	AdresseLigne *string `json:"adresse_ligne,omitempty"`
	// Complément de l'adresse du gestionnaire.
	ComplementAdresse *string `json:"complement_adresse,omitempty"`
	// Indication supplémentaire à l'adresse du gestionnaire.
	Indication *string `json:"indication,omitempty"`
	// Pays du gestionnaire.
	Pays *string `json:"pays,omitempty"`
}

// NewAssociationAdresseGestionnaire instantiates a new AssociationAdresseGestionnaire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationAdresseGestionnaire() *AssociationAdresseGestionnaire {
	this := AssociationAdresseGestionnaire{}
	return &this
}

// NewAssociationAdresseGestionnaireWithDefaults instantiates a new AssociationAdresseGestionnaire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationAdresseGestionnaireWithDefaults() *AssociationAdresseGestionnaire {
	this := AssociationAdresseGestionnaire{}
	return &this
}

// GetGestionnaire returns the Gestionnaire field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetGestionnaire() string {
	if o == nil || IsNil(o.Gestionnaire) {
		var ret string
		return ret
	}
	return *o.Gestionnaire
}

// GetGestionnaireOk returns a tuple with the Gestionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetGestionnaireOk() (*string, bool) {
	if o == nil || IsNil(o.Gestionnaire) {
		return nil, false
	}
	return o.Gestionnaire, true
}

// HasGestionnaire returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasGestionnaire() bool {
	if o != nil && !IsNil(o.Gestionnaire) {
		return true
	}

	return false
}

// SetGestionnaire gets a reference to the given string and assigns it to the Gestionnaire field.
func (o *AssociationAdresseGestionnaire) SetGestionnaire(v string) {
	o.Gestionnaire = &v
}

// GetCodePostal returns the CodePostal field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetCodePostal() string {
	if o == nil || IsNil(o.CodePostal) {
		var ret string
		return ret
	}
	return *o.CodePostal
}

// GetCodePostalOk returns a tuple with the CodePostal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetCodePostalOk() (*string, bool) {
	if o == nil || IsNil(o.CodePostal) {
		return nil, false
	}
	return o.CodePostal, true
}

// HasCodePostal returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasCodePostal() bool {
	if o != nil && !IsNil(o.CodePostal) {
		return true
	}

	return false
}

// SetCodePostal gets a reference to the given string and assigns it to the CodePostal field.
func (o *AssociationAdresseGestionnaire) SetCodePostal(v string) {
	o.CodePostal = &v
}

// GetVille returns the Ville field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetVille() string {
	if o == nil || IsNil(o.Ville) {
		var ret string
		return ret
	}
	return *o.Ville
}

// GetVilleOk returns a tuple with the Ville field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetVilleOk() (*string, bool) {
	if o == nil || IsNil(o.Ville) {
		return nil, false
	}
	return o.Ville, true
}

// HasVille returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasVille() bool {
	if o != nil && !IsNil(o.Ville) {
		return true
	}

	return false
}

// SetVille gets a reference to the given string and assigns it to the Ville field.
func (o *AssociationAdresseGestionnaire) SetVille(v string) {
	o.Ville = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *AssociationAdresseGestionnaire) SetDistribution(v string) {
	o.Distribution = &v
}

// GetAdresseLigne returns the AdresseLigne field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetAdresseLigne() string {
	if o == nil || IsNil(o.AdresseLigne) {
		var ret string
		return ret
	}
	return *o.AdresseLigne
}

// GetAdresseLigneOk returns a tuple with the AdresseLigne field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetAdresseLigneOk() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne) {
		return nil, false
	}
	return o.AdresseLigne, true
}

// HasAdresseLigne returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasAdresseLigne() bool {
	if o != nil && !IsNil(o.AdresseLigne) {
		return true
	}

	return false
}

// SetAdresseLigne gets a reference to the given string and assigns it to the AdresseLigne field.
func (o *AssociationAdresseGestionnaire) SetAdresseLigne(v string) {
	o.AdresseLigne = &v
}

// GetComplementAdresse returns the ComplementAdresse field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetComplementAdresse() string {
	if o == nil || IsNil(o.ComplementAdresse) {
		var ret string
		return ret
	}
	return *o.ComplementAdresse
}

// GetComplementAdresseOk returns a tuple with the ComplementAdresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetComplementAdresseOk() (*string, bool) {
	if o == nil || IsNil(o.ComplementAdresse) {
		return nil, false
	}
	return o.ComplementAdresse, true
}

// HasComplementAdresse returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasComplementAdresse() bool {
	if o != nil && !IsNil(o.ComplementAdresse) {
		return true
	}

	return false
}

// SetComplementAdresse gets a reference to the given string and assigns it to the ComplementAdresse field.
func (o *AssociationAdresseGestionnaire) SetComplementAdresse(v string) {
	o.ComplementAdresse = &v
}

// GetIndication returns the Indication field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetIndication() string {
	if o == nil || IsNil(o.Indication) {
		var ret string
		return ret
	}
	return *o.Indication
}

// GetIndicationOk returns a tuple with the Indication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetIndicationOk() (*string, bool) {
	if o == nil || IsNil(o.Indication) {
		return nil, false
	}
	return o.Indication, true
}

// HasIndication returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasIndication() bool {
	if o != nil && !IsNil(o.Indication) {
		return true
	}

	return false
}

// SetIndication gets a reference to the given string and assigns it to the Indication field.
func (o *AssociationAdresseGestionnaire) SetIndication(v string) {
	o.Indication = &v
}

// GetPays returns the Pays field value if set, zero value otherwise.
func (o *AssociationAdresseGestionnaire) GetPays() string {
	if o == nil || IsNil(o.Pays) {
		var ret string
		return ret
	}
	return *o.Pays
}

// GetPaysOk returns a tuple with the Pays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationAdresseGestionnaire) GetPaysOk() (*string, bool) {
	if o == nil || IsNil(o.Pays) {
		return nil, false
	}
	return o.Pays, true
}

// HasPays returns a boolean if a field has been set.
func (o *AssociationAdresseGestionnaire) HasPays() bool {
	if o != nil && !IsNil(o.Pays) {
		return true
	}

	return false
}

// SetPays gets a reference to the given string and assigns it to the Pays field.
func (o *AssociationAdresseGestionnaire) SetPays(v string) {
	o.Pays = &v
}

func (o AssociationAdresseGestionnaire) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociationAdresseGestionnaire) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gestionnaire) {
		toSerialize["gestionnaire"] = o.Gestionnaire
	}
	if !IsNil(o.CodePostal) {
		toSerialize["code_postal"] = o.CodePostal
	}
	if !IsNil(o.Ville) {
		toSerialize["ville"] = o.Ville
	}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	if !IsNil(o.AdresseLigne) {
		toSerialize["adresse_ligne"] = o.AdresseLigne
	}
	if !IsNil(o.ComplementAdresse) {
		toSerialize["complement_adresse"] = o.ComplementAdresse
	}
	if !IsNil(o.Indication) {
		toSerialize["indication"] = o.Indication
	}
	if !IsNil(o.Pays) {
		toSerialize["pays"] = o.Pays
	}
	return toSerialize, nil
}

type NullableAssociationAdresseGestionnaire struct {
	value *AssociationAdresseGestionnaire
	isSet bool
}

func (v NullableAssociationAdresseGestionnaire) Get() *AssociationAdresseGestionnaire {
	return v.value
}

func (v *NullableAssociationAdresseGestionnaire) Set(val *AssociationAdresseGestionnaire) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationAdresseGestionnaire) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationAdresseGestionnaire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationAdresseGestionnaire(val *AssociationAdresseGestionnaire) *NullableAssociationAdresseGestionnaire {
	return &NullableAssociationAdresseGestionnaire{value: val, isSet: true}
}

func (v NullableAssociationAdresseGestionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationAdresseGestionnaire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


