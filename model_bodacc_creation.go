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

// checks if the BodaccCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BodaccCreation{}

// BodaccCreation struct for BodaccCreation
type BodaccCreation struct {
	Bodacc
	// Nom de l'entreprise concernée par la publication. Correspond à la dénomination en cas de personne morale et à nom + prenom en cas de personne physique.
	NomEntreprise *string `json:"nom_entreprise,omitempty"`
	// Vrai si l'entreprise concernée par la publication est une personne morale, faux si c'est une personne physique.
	PersonneMorale *bool `json:"personne_morale,omitempty"`
	// Dénomination de l'entreprise concernée par la publication (uniquement en cas de personne morale).
	Denomination *string `json:"denomination,omitempty"`
	// Nom de la personne physique concernée par la publication (uniquement en cas de personne physique).
	Nom *string `json:"nom,omitempty"`
	// Prénom de la personne physique concernée par la publication (uniquement en cas de personne physique).
	Prenom *string `json:"prenom,omitempty"`
	// Administration (dans le cas d'une personne morale).
	Administration *string `json:"administration,omitempty"`
	// Adresse de l'entreprise concernée par la publication.
	Adresse *string `json:"adresse,omitempty"`
	// Capital de l'entreprise concernée par la publication.
	Capital *float32 `json:"capital,omitempty"`
	// Activité de l'entreprise concernée par la publication.
	Activite *string `json:"activite,omitempty"`
	// Date de début d'activité de l'entreprise concernée par la publication, au format AAAA-MM-JJ.
	DateDebutActivite *string `json:"date_debut_activite,omitempty"`
}

// NewBodaccCreation instantiates a new BodaccCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBodaccCreation() *BodaccCreation {
	this := BodaccCreation{}
	return &this
}

// NewBodaccCreationWithDefaults instantiates a new BodaccCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBodaccCreationWithDefaults() *BodaccCreation {
	this := BodaccCreation{}
	return &this
}

// GetNomEntreprise returns the NomEntreprise field value if set, zero value otherwise.
func (o *BodaccCreation) GetNomEntreprise() string {
	if o == nil || IsNil(o.NomEntreprise) {
		var ret string
		return ret
	}
	return *o.NomEntreprise
}

// GetNomEntrepriseOk returns a tuple with the NomEntreprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetNomEntrepriseOk() (*string, bool) {
	if o == nil || IsNil(o.NomEntreprise) {
		return nil, false
	}
	return o.NomEntreprise, true
}

// HasNomEntreprise returns a boolean if a field has been set.
func (o *BodaccCreation) HasNomEntreprise() bool {
	if o != nil && !IsNil(o.NomEntreprise) {
		return true
	}

	return false
}

// SetNomEntreprise gets a reference to the given string and assigns it to the NomEntreprise field.
func (o *BodaccCreation) SetNomEntreprise(v string) {
	o.NomEntreprise = &v
}

// GetPersonneMorale returns the PersonneMorale field value if set, zero value otherwise.
func (o *BodaccCreation) GetPersonneMorale() bool {
	if o == nil || IsNil(o.PersonneMorale) {
		var ret bool
		return ret
	}
	return *o.PersonneMorale
}

// GetPersonneMoraleOk returns a tuple with the PersonneMorale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetPersonneMoraleOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonneMorale) {
		return nil, false
	}
	return o.PersonneMorale, true
}

// HasPersonneMorale returns a boolean if a field has been set.
func (o *BodaccCreation) HasPersonneMorale() bool {
	if o != nil && !IsNil(o.PersonneMorale) {
		return true
	}

	return false
}

// SetPersonneMorale gets a reference to the given bool and assigns it to the PersonneMorale field.
func (o *BodaccCreation) SetPersonneMorale(v bool) {
	o.PersonneMorale = &v
}

// GetDenomination returns the Denomination field value if set, zero value otherwise.
func (o *BodaccCreation) GetDenomination() string {
	if o == nil || IsNil(o.Denomination) {
		var ret string
		return ret
	}
	return *o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetDenominationOk() (*string, bool) {
	if o == nil || IsNil(o.Denomination) {
		return nil, false
	}
	return o.Denomination, true
}

// HasDenomination returns a boolean if a field has been set.
func (o *BodaccCreation) HasDenomination() bool {
	if o != nil && !IsNil(o.Denomination) {
		return true
	}

	return false
}

// SetDenomination gets a reference to the given string and assigns it to the Denomination field.
func (o *BodaccCreation) SetDenomination(v string) {
	o.Denomination = &v
}

// GetNom returns the Nom field value if set, zero value otherwise.
func (o *BodaccCreation) GetNom() string {
	if o == nil || IsNil(o.Nom) {
		var ret string
		return ret
	}
	return *o.Nom
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetNomOk() (*string, bool) {
	if o == nil || IsNil(o.Nom) {
		return nil, false
	}
	return o.Nom, true
}

// HasNom returns a boolean if a field has been set.
func (o *BodaccCreation) HasNom() bool {
	if o != nil && !IsNil(o.Nom) {
		return true
	}

	return false
}

// SetNom gets a reference to the given string and assigns it to the Nom field.
func (o *BodaccCreation) SetNom(v string) {
	o.Nom = &v
}

// GetPrenom returns the Prenom field value if set, zero value otherwise.
func (o *BodaccCreation) GetPrenom() string {
	if o == nil || IsNil(o.Prenom) {
		var ret string
		return ret
	}
	return *o.Prenom
}

// GetPrenomOk returns a tuple with the Prenom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetPrenomOk() (*string, bool) {
	if o == nil || IsNil(o.Prenom) {
		return nil, false
	}
	return o.Prenom, true
}

// HasPrenom returns a boolean if a field has been set.
func (o *BodaccCreation) HasPrenom() bool {
	if o != nil && !IsNil(o.Prenom) {
		return true
	}

	return false
}

// SetPrenom gets a reference to the given string and assigns it to the Prenom field.
func (o *BodaccCreation) SetPrenom(v string) {
	o.Prenom = &v
}

// GetAdministration returns the Administration field value if set, zero value otherwise.
func (o *BodaccCreation) GetAdministration() string {
	if o == nil || IsNil(o.Administration) {
		var ret string
		return ret
	}
	return *o.Administration
}

// GetAdministrationOk returns a tuple with the Administration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetAdministrationOk() (*string, bool) {
	if o == nil || IsNil(o.Administration) {
		return nil, false
	}
	return o.Administration, true
}

// HasAdministration returns a boolean if a field has been set.
func (o *BodaccCreation) HasAdministration() bool {
	if o != nil && !IsNil(o.Administration) {
		return true
	}

	return false
}

// SetAdministration gets a reference to the given string and assigns it to the Administration field.
func (o *BodaccCreation) SetAdministration(v string) {
	o.Administration = &v
}

// GetAdresse returns the Adresse field value if set, zero value otherwise.
func (o *BodaccCreation) GetAdresse() string {
	if o == nil || IsNil(o.Adresse) {
		var ret string
		return ret
	}
	return *o.Adresse
}

// GetAdresseOk returns a tuple with the Adresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetAdresseOk() (*string, bool) {
	if o == nil || IsNil(o.Adresse) {
		return nil, false
	}
	return o.Adresse, true
}

// HasAdresse returns a boolean if a field has been set.
func (o *BodaccCreation) HasAdresse() bool {
	if o != nil && !IsNil(o.Adresse) {
		return true
	}

	return false
}

// SetAdresse gets a reference to the given string and assigns it to the Adresse field.
func (o *BodaccCreation) SetAdresse(v string) {
	o.Adresse = &v
}

// GetCapital returns the Capital field value if set, zero value otherwise.
func (o *BodaccCreation) GetCapital() float32 {
	if o == nil || IsNil(o.Capital) {
		var ret float32
		return ret
	}
	return *o.Capital
}

// GetCapitalOk returns a tuple with the Capital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetCapitalOk() (*float32, bool) {
	if o == nil || IsNil(o.Capital) {
		return nil, false
	}
	return o.Capital, true
}

// HasCapital returns a boolean if a field has been set.
func (o *BodaccCreation) HasCapital() bool {
	if o != nil && !IsNil(o.Capital) {
		return true
	}

	return false
}

// SetCapital gets a reference to the given float32 and assigns it to the Capital field.
func (o *BodaccCreation) SetCapital(v float32) {
	o.Capital = &v
}

// GetActivite returns the Activite field value if set, zero value otherwise.
func (o *BodaccCreation) GetActivite() string {
	if o == nil || IsNil(o.Activite) {
		var ret string
		return ret
	}
	return *o.Activite
}

// GetActiviteOk returns a tuple with the Activite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetActiviteOk() (*string, bool) {
	if o == nil || IsNil(o.Activite) {
		return nil, false
	}
	return o.Activite, true
}

// HasActivite returns a boolean if a field has been set.
func (o *BodaccCreation) HasActivite() bool {
	if o != nil && !IsNil(o.Activite) {
		return true
	}

	return false
}

// SetActivite gets a reference to the given string and assigns it to the Activite field.
func (o *BodaccCreation) SetActivite(v string) {
	o.Activite = &v
}

// GetDateDebutActivite returns the DateDebutActivite field value if set, zero value otherwise.
func (o *BodaccCreation) GetDateDebutActivite() string {
	if o == nil || IsNil(o.DateDebutActivite) {
		var ret string
		return ret
	}
	return *o.DateDebutActivite
}

// GetDateDebutActiviteOk returns a tuple with the DateDebutActivite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BodaccCreation) GetDateDebutActiviteOk() (*string, bool) {
	if o == nil || IsNil(o.DateDebutActivite) {
		return nil, false
	}
	return o.DateDebutActivite, true
}

// HasDateDebutActivite returns a boolean if a field has been set.
func (o *BodaccCreation) HasDateDebutActivite() bool {
	if o != nil && !IsNil(o.DateDebutActivite) {
		return true
	}

	return false
}

// SetDateDebutActivite gets a reference to the given string and assigns it to the DateDebutActivite field.
func (o *BodaccCreation) SetDateDebutActivite(v string) {
	o.DateDebutActivite = &v
}

func (o BodaccCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BodaccCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBodacc, errBodacc := json.Marshal(o.Bodacc)
	if errBodacc != nil {
		return map[string]interface{}{}, errBodacc
	}
	errBodacc = json.Unmarshal([]byte(serializedBodacc), &toSerialize)
	if errBodacc != nil {
		return map[string]interface{}{}, errBodacc
	}
	if !IsNil(o.NomEntreprise) {
		toSerialize["nom_entreprise"] = o.NomEntreprise
	}
	if !IsNil(o.PersonneMorale) {
		toSerialize["personne_morale"] = o.PersonneMorale
	}
	if !IsNil(o.Denomination) {
		toSerialize["denomination"] = o.Denomination
	}
	if !IsNil(o.Nom) {
		toSerialize["nom"] = o.Nom
	}
	if !IsNil(o.Prenom) {
		toSerialize["prenom"] = o.Prenom
	}
	if !IsNil(o.Administration) {
		toSerialize["administration"] = o.Administration
	}
	if !IsNil(o.Adresse) {
		toSerialize["adresse"] = o.Adresse
	}
	if !IsNil(o.Capital) {
		toSerialize["capital"] = o.Capital
	}
	if !IsNil(o.Activite) {
		toSerialize["activite"] = o.Activite
	}
	if !IsNil(o.DateDebutActivite) {
		toSerialize["date_debut_activite"] = o.DateDebutActivite
	}
	return toSerialize, nil
}

type NullableBodaccCreation struct {
	value *BodaccCreation
	isSet bool
}

func (v NullableBodaccCreation) Get() *BodaccCreation {
	return v.value
}

func (v *NullableBodaccCreation) Set(val *BodaccCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableBodaccCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableBodaccCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBodaccCreation(val *BodaccCreation) *NullableBodaccCreation {
	return &NullableBodaccCreation{value: val, isSet: true}
}

func (v NullableBodaccCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBodaccCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


