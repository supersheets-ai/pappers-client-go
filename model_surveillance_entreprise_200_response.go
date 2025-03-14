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

// checks if the SurveillanceEntreprise200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveillanceEntreprise200Response{}

// SurveillanceEntreprise200Response Nombre de notifications ajoutées.
type SurveillanceEntreprise200Response struct {
	// Le nombre d'entreprises ajoutées à votre liste de surveillance d'entreprises.
	NotificationsAjoutees *float32 `json:"notifications_ajoutees,omitempty"`
}

// NewSurveillanceEntreprise200Response instantiates a new SurveillanceEntreprise200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveillanceEntreprise200Response() *SurveillanceEntreprise200Response {
	this := SurveillanceEntreprise200Response{}
	var notificationsAjoutees float32 = 0
	this.NotificationsAjoutees = &notificationsAjoutees
	return &this
}

// NewSurveillanceEntreprise200ResponseWithDefaults instantiates a new SurveillanceEntreprise200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveillanceEntreprise200ResponseWithDefaults() *SurveillanceEntreprise200Response {
	this := SurveillanceEntreprise200Response{}
	var notificationsAjoutees float32 = 0
	this.NotificationsAjoutees = &notificationsAjoutees
	return &this
}

// GetNotificationsAjoutees returns the NotificationsAjoutees field value if set, zero value otherwise.
func (o *SurveillanceEntreprise200Response) GetNotificationsAjoutees() float32 {
	if o == nil || IsNil(o.NotificationsAjoutees) {
		var ret float32
		return ret
	}
	return *o.NotificationsAjoutees
}

// GetNotificationsAjouteesOk returns a tuple with the NotificationsAjoutees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveillanceEntreprise200Response) GetNotificationsAjouteesOk() (*float32, bool) {
	if o == nil || IsNil(o.NotificationsAjoutees) {
		return nil, false
	}
	return o.NotificationsAjoutees, true
}

// HasNotificationsAjoutees returns a boolean if a field has been set.
func (o *SurveillanceEntreprise200Response) HasNotificationsAjoutees() bool {
	if o != nil && !IsNil(o.NotificationsAjoutees) {
		return true
	}

	return false
}

// SetNotificationsAjoutees gets a reference to the given float32 and assigns it to the NotificationsAjoutees field.
func (o *SurveillanceEntreprise200Response) SetNotificationsAjoutees(v float32) {
	o.NotificationsAjoutees = &v
}

func (o SurveillanceEntreprise200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveillanceEntreprise200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationsAjoutees) {
		toSerialize["notifications_ajoutees"] = o.NotificationsAjoutees
	}
	return toSerialize, nil
}

type NullableSurveillanceEntreprise200Response struct {
	value *SurveillanceEntreprise200Response
	isSet bool
}

func (v NullableSurveillanceEntreprise200Response) Get() *SurveillanceEntreprise200Response {
	return v.value
}

func (v *NullableSurveillanceEntreprise200Response) Set(val *SurveillanceEntreprise200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveillanceEntreprise200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveillanceEntreprise200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveillanceEntreprise200Response(val *SurveillanceEntreprise200Response) *NullableSurveillanceEntreprise200Response {
	return &NullableSurveillanceEntreprise200Response{value: val, isSet: true}
}

func (v NullableSurveillanceEntreprise200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveillanceEntreprise200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


