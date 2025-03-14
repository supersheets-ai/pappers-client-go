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

// checks if the EtablissementRecherche type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EtablissementRecherche{}

// EtablissementRecherche struct for EtablissementRecherche
type EtablissementRecherche struct {
	// Numéro siret de l'établissement au format xxxxxxxxxxxxxx.
	Siret *string `json:"siret,omitempty"`
	// Numéro siret de l'établissement au format xxx xxx xxx xxxxx.
	SiretFormate *string `json:"siret_formate,omitempty"`
	// Numéro NIC de l'établissement.
	Nic *string `json:"nic,omitempty"`
	// Numéro de voie de l'établissement.
	NumeroVoie *int32 `json:"numero_voie,omitempty"`
	// Indice de répétition de l'établissement.
	IndiceRepetition *string `json:"indice_repetition,omitempty"`
	// Type de voie de l'établissement.
	TypeVoie *string `json:"type_voie,omitempty"`
	// Libellé de la voie de l'établissement.
	LibelleVoie *string `json:"libelle_voie,omitempty"`
	// Complément d'adresse de l'établissement.
	ComplementAdresse *string `json:"complement_adresse,omitempty"`
	// Première ligne de l'adresse de l'établissement. Correspond à l'ensemble des données numero_voie, indice_repetition, type_voie et libelle_voie.
	AdresseLigne1 *string `json:"adresse_ligne_1,omitempty"`
	// Seconde ligne de l'adresse de l'établissement. Est égal à complement_adresse
	AdresseLigne2 *string `json:"adresse_ligne_2,omitempty"`
	// Code postal de l'établissement.
	CodePostal *string `json:"code_postal,omitempty"`
	// Ville de l'établissement.
	Ville *string `json:"ville,omitempty"`
	// Latitude de l'établissement.
	Latitude *float32 `json:"latitude,omitempty"`
	// Longitude de l'établissement.
	Longitude *float32 `json:"longitude,omitempty"`
}

// NewEtablissementRecherche instantiates a new EtablissementRecherche object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtablissementRecherche() *EtablissementRecherche {
	this := EtablissementRecherche{}
	return &this
}

// NewEtablissementRechercheWithDefaults instantiates a new EtablissementRecherche object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtablissementRechercheWithDefaults() *EtablissementRecherche {
	this := EtablissementRecherche{}
	return &this
}

// GetSiret returns the Siret field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetSiret() string {
	if o == nil || IsNil(o.Siret) {
		var ret string
		return ret
	}
	return *o.Siret
}

// GetSiretOk returns a tuple with the Siret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetSiretOk() (*string, bool) {
	if o == nil || IsNil(o.Siret) {
		return nil, false
	}
	return o.Siret, true
}

// HasSiret returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasSiret() bool {
	if o != nil && !IsNil(o.Siret) {
		return true
	}

	return false
}

// SetSiret gets a reference to the given string and assigns it to the Siret field.
func (o *EtablissementRecherche) SetSiret(v string) {
	o.Siret = &v
}

// GetSiretFormate returns the SiretFormate field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetSiretFormate() string {
	if o == nil || IsNil(o.SiretFormate) {
		var ret string
		return ret
	}
	return *o.SiretFormate
}

// GetSiretFormateOk returns a tuple with the SiretFormate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetSiretFormateOk() (*string, bool) {
	if o == nil || IsNil(o.SiretFormate) {
		return nil, false
	}
	return o.SiretFormate, true
}

// HasSiretFormate returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasSiretFormate() bool {
	if o != nil && !IsNil(o.SiretFormate) {
		return true
	}

	return false
}

// SetSiretFormate gets a reference to the given string and assigns it to the SiretFormate field.
func (o *EtablissementRecherche) SetSiretFormate(v string) {
	o.SiretFormate = &v
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetNic() string {
	if o == nil || IsNil(o.Nic) {
		var ret string
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetNicOk() (*string, bool) {
	if o == nil || IsNil(o.Nic) {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasNic() bool {
	if o != nil && !IsNil(o.Nic) {
		return true
	}

	return false
}

// SetNic gets a reference to the given string and assigns it to the Nic field.
func (o *EtablissementRecherche) SetNic(v string) {
	o.Nic = &v
}

// GetNumeroVoie returns the NumeroVoie field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetNumeroVoie() int32 {
	if o == nil || IsNil(o.NumeroVoie) {
		var ret int32
		return ret
	}
	return *o.NumeroVoie
}

// GetNumeroVoieOk returns a tuple with the NumeroVoie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetNumeroVoieOk() (*int32, bool) {
	if o == nil || IsNil(o.NumeroVoie) {
		return nil, false
	}
	return o.NumeroVoie, true
}

// HasNumeroVoie returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasNumeroVoie() bool {
	if o != nil && !IsNil(o.NumeroVoie) {
		return true
	}

	return false
}

// SetNumeroVoie gets a reference to the given int32 and assigns it to the NumeroVoie field.
func (o *EtablissementRecherche) SetNumeroVoie(v int32) {
	o.NumeroVoie = &v
}

// GetIndiceRepetition returns the IndiceRepetition field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetIndiceRepetition() string {
	if o == nil || IsNil(o.IndiceRepetition) {
		var ret string
		return ret
	}
	return *o.IndiceRepetition
}

// GetIndiceRepetitionOk returns a tuple with the IndiceRepetition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetIndiceRepetitionOk() (*string, bool) {
	if o == nil || IsNil(o.IndiceRepetition) {
		return nil, false
	}
	return o.IndiceRepetition, true
}

// HasIndiceRepetition returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasIndiceRepetition() bool {
	if o != nil && !IsNil(o.IndiceRepetition) {
		return true
	}

	return false
}

// SetIndiceRepetition gets a reference to the given string and assigns it to the IndiceRepetition field.
func (o *EtablissementRecherche) SetIndiceRepetition(v string) {
	o.IndiceRepetition = &v
}

// GetTypeVoie returns the TypeVoie field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetTypeVoie() string {
	if o == nil || IsNil(o.TypeVoie) {
		var ret string
		return ret
	}
	return *o.TypeVoie
}

// GetTypeVoieOk returns a tuple with the TypeVoie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetTypeVoieOk() (*string, bool) {
	if o == nil || IsNil(o.TypeVoie) {
		return nil, false
	}
	return o.TypeVoie, true
}

// HasTypeVoie returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasTypeVoie() bool {
	if o != nil && !IsNil(o.TypeVoie) {
		return true
	}

	return false
}

// SetTypeVoie gets a reference to the given string and assigns it to the TypeVoie field.
func (o *EtablissementRecherche) SetTypeVoie(v string) {
	o.TypeVoie = &v
}

// GetLibelleVoie returns the LibelleVoie field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetLibelleVoie() string {
	if o == nil || IsNil(o.LibelleVoie) {
		var ret string
		return ret
	}
	return *o.LibelleVoie
}

// GetLibelleVoieOk returns a tuple with the LibelleVoie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetLibelleVoieOk() (*string, bool) {
	if o == nil || IsNil(o.LibelleVoie) {
		return nil, false
	}
	return o.LibelleVoie, true
}

// HasLibelleVoie returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasLibelleVoie() bool {
	if o != nil && !IsNil(o.LibelleVoie) {
		return true
	}

	return false
}

// SetLibelleVoie gets a reference to the given string and assigns it to the LibelleVoie field.
func (o *EtablissementRecherche) SetLibelleVoie(v string) {
	o.LibelleVoie = &v
}

// GetComplementAdresse returns the ComplementAdresse field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetComplementAdresse() string {
	if o == nil || IsNil(o.ComplementAdresse) {
		var ret string
		return ret
	}
	return *o.ComplementAdresse
}

// GetComplementAdresseOk returns a tuple with the ComplementAdresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetComplementAdresseOk() (*string, bool) {
	if o == nil || IsNil(o.ComplementAdresse) {
		return nil, false
	}
	return o.ComplementAdresse, true
}

// HasComplementAdresse returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasComplementAdresse() bool {
	if o != nil && !IsNil(o.ComplementAdresse) {
		return true
	}

	return false
}

// SetComplementAdresse gets a reference to the given string and assigns it to the ComplementAdresse field.
func (o *EtablissementRecherche) SetComplementAdresse(v string) {
	o.ComplementAdresse = &v
}

// GetAdresseLigne1 returns the AdresseLigne1 field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetAdresseLigne1() string {
	if o == nil || IsNil(o.AdresseLigne1) {
		var ret string
		return ret
	}
	return *o.AdresseLigne1
}

// GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetAdresseLigne1Ok() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne1) {
		return nil, false
	}
	return o.AdresseLigne1, true
}

// HasAdresseLigne1 returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasAdresseLigne1() bool {
	if o != nil && !IsNil(o.AdresseLigne1) {
		return true
	}

	return false
}

// SetAdresseLigne1 gets a reference to the given string and assigns it to the AdresseLigne1 field.
func (o *EtablissementRecherche) SetAdresseLigne1(v string) {
	o.AdresseLigne1 = &v
}

// GetAdresseLigne2 returns the AdresseLigne2 field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetAdresseLigne2() string {
	if o == nil || IsNil(o.AdresseLigne2) {
		var ret string
		return ret
	}
	return *o.AdresseLigne2
}

// GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetAdresseLigne2Ok() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne2) {
		return nil, false
	}
	return o.AdresseLigne2, true
}

// HasAdresseLigne2 returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasAdresseLigne2() bool {
	if o != nil && !IsNil(o.AdresseLigne2) {
		return true
	}

	return false
}

// SetAdresseLigne2 gets a reference to the given string and assigns it to the AdresseLigne2 field.
func (o *EtablissementRecherche) SetAdresseLigne2(v string) {
	o.AdresseLigne2 = &v
}

// GetCodePostal returns the CodePostal field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetCodePostal() string {
	if o == nil || IsNil(o.CodePostal) {
		var ret string
		return ret
	}
	return *o.CodePostal
}

// GetCodePostalOk returns a tuple with the CodePostal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetCodePostalOk() (*string, bool) {
	if o == nil || IsNil(o.CodePostal) {
		return nil, false
	}
	return o.CodePostal, true
}

// HasCodePostal returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasCodePostal() bool {
	if o != nil && !IsNil(o.CodePostal) {
		return true
	}

	return false
}

// SetCodePostal gets a reference to the given string and assigns it to the CodePostal field.
func (o *EtablissementRecherche) SetCodePostal(v string) {
	o.CodePostal = &v
}

// GetVille returns the Ville field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetVille() string {
	if o == nil || IsNil(o.Ville) {
		var ret string
		return ret
	}
	return *o.Ville
}

// GetVilleOk returns a tuple with the Ville field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetVilleOk() (*string, bool) {
	if o == nil || IsNil(o.Ville) {
		return nil, false
	}
	return o.Ville, true
}

// HasVille returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasVille() bool {
	if o != nil && !IsNil(o.Ville) {
		return true
	}

	return false
}

// SetVille gets a reference to the given string and assigns it to the Ville field.
func (o *EtablissementRecherche) SetVille(v string) {
	o.Ville = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *EtablissementRecherche) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *EtablissementRecherche) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtablissementRecherche) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *EtablissementRecherche) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *EtablissementRecherche) SetLongitude(v float32) {
	o.Longitude = &v
}

func (o EtablissementRecherche) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EtablissementRecherche) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Siret) {
		toSerialize["siret"] = o.Siret
	}
	if !IsNil(o.SiretFormate) {
		toSerialize["siret_formate"] = o.SiretFormate
	}
	if !IsNil(o.Nic) {
		toSerialize["nic"] = o.Nic
	}
	if !IsNil(o.NumeroVoie) {
		toSerialize["numero_voie"] = o.NumeroVoie
	}
	if !IsNil(o.IndiceRepetition) {
		toSerialize["indice_repetition"] = o.IndiceRepetition
	}
	if !IsNil(o.TypeVoie) {
		toSerialize["type_voie"] = o.TypeVoie
	}
	if !IsNil(o.LibelleVoie) {
		toSerialize["libelle_voie"] = o.LibelleVoie
	}
	if !IsNil(o.ComplementAdresse) {
		toSerialize["complement_adresse"] = o.ComplementAdresse
	}
	if !IsNil(o.AdresseLigne1) {
		toSerialize["adresse_ligne_1"] = o.AdresseLigne1
	}
	if !IsNil(o.AdresseLigne2) {
		toSerialize["adresse_ligne_2"] = o.AdresseLigne2
	}
	if !IsNil(o.CodePostal) {
		toSerialize["code_postal"] = o.CodePostal
	}
	if !IsNil(o.Ville) {
		toSerialize["ville"] = o.Ville
	}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	return toSerialize, nil
}

type NullableEtablissementRecherche struct {
	value *EtablissementRecherche
	isSet bool
}

func (v NullableEtablissementRecherche) Get() *EtablissementRecherche {
	return v.value
}

func (v *NullableEtablissementRecherche) Set(val *EtablissementRecherche) {
	v.value = val
	v.isSet = true
}

func (v NullableEtablissementRecherche) IsSet() bool {
	return v.isSet
}

func (v *NullableEtablissementRecherche) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtablissementRecherche(val *EtablissementRecherche) *NullableEtablissementRecherche {
	return &NullableEtablissementRecherche{value: val, isSet: true}
}

func (v NullableEtablissementRecherche) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtablissementRecherche) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


