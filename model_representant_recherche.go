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

// checks if the RepresentantRecherche type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepresentantRecherche{}

// RepresentantRecherche struct for RepresentantRecherche
type RepresentantRecherche struct {
	// Qualité du représentant.
	Qualite *string `json:"qualite,omitempty"`
	// Vrai si le représentant est une personne morale, faux si personne physique.
	PersonneMorale *bool `json:"personne_morale,omitempty"`
	// Date de prise de poste du représentant.
	DatePriseDePoste *string `json:"date_prise_de_poste,omitempty"`
	// Dénomination du représentant si personne morale.
	Denomination *string `json:"denomination,omitempty"`
	// Siren du représentant si personne morale.
	Siren *string `json:"siren,omitempty"`
	// Forme juridique du représentant dans le cas d'une personne morale.
	FormeJuridique *string `json:"forme_juridique,omitempty"`
	// Sexe supposé du représentant si personne physique. F pour féminin, M pour masculin. Ce champ est estimé à partir du prénom du représentant.
	Sexe *string `json:"sexe,omitempty"`
	// Nom du représentant.
	Nom *string `json:"nom,omitempty"`
	// Prénoms du représentant.
	Prenom *string `json:"prenom,omitempty"`
	// Prénom usuel du représentant.
	PrenomUsuel *string `json:"prenom_usuel,omitempty"`
	// Nom complet du représentant.
	NomComplet *string `json:"nom_complet,omitempty"`
	// Date de naissance du représentant.
	DateDeNaissance *string `json:"date_de_naissance,omitempty"`
	// Date de naissance formatée du représentant.
	DateDeNaissanceFormate *string `json:"date_de_naissance_formate,omitempty"`
	// Age du représentant.
	Age *int32 `json:"age,omitempty"`
	// Nationalité du représentant.
	Nationalite *string `json:"nationalite,omitempty"`
	// Code nationalité du représentant
	CodeNationalite *string `json:"code_nationalite,omitempty"`
	// Ville de naissance du représentant.
	VilleDeNaissance *string `json:"ville_de_naissance,omitempty"`
	// Pays de naissance du représentant.
	PaysDeNaissance *string `json:"pays_de_naissance,omitempty"`
	// Code du pays de naissance du représentant.
	CodePaysDeNaissance *string `json:"code_pays_de_naissance,omitempty"`
	// Première ligne de l'adresse du représentant.
	AdresseLigne1 *string `json:"adresse_ligne_1,omitempty"`
	// Deuxième ligne de l'adresse du représentant.
	AdresseLigne2 *string `json:"adresse_ligne_2,omitempty"`
	// Troisième ligne de l'adresse du représentant.
	AdresseLigne3 *string `json:"adresse_ligne_3,omitempty"`
	// Code postal du représentant.
	CodePostal *string `json:"code_postal,omitempty"`
	// Ville du représentant.
	Ville *string `json:"ville,omitempty"`
	// Pays du représentant.
	Pays *string `json:"pays,omitempty"`
	// Code du pays du représentant
	CodePays *string `json:"code_pays,omitempty"`
	// Vaut vrai si le représentant est toujours à son poste.
	Actuel *bool `json:"actuel,omitempty"`
	// Date de départ de poste dans le cas où le représentant n'est plus un représentant actuel, au format AAAA-MM-JJ.
	DateDepartDePoste *string `json:"date_depart_de_poste,omitempty"`
}

// NewRepresentantRecherche instantiates a new RepresentantRecherche object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepresentantRecherche() *RepresentantRecherche {
	this := RepresentantRecherche{}
	return &this
}

// NewRepresentantRechercheWithDefaults instantiates a new RepresentantRecherche object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepresentantRechercheWithDefaults() *RepresentantRecherche {
	this := RepresentantRecherche{}
	return &this
}

// GetQualite returns the Qualite field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetQualite() string {
	if o == nil || IsNil(o.Qualite) {
		var ret string
		return ret
	}
	return *o.Qualite
}

// GetQualiteOk returns a tuple with the Qualite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetQualiteOk() (*string, bool) {
	if o == nil || IsNil(o.Qualite) {
		return nil, false
	}
	return o.Qualite, true
}

// HasQualite returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasQualite() bool {
	if o != nil && !IsNil(o.Qualite) {
		return true
	}

	return false
}

// SetQualite gets a reference to the given string and assigns it to the Qualite field.
func (o *RepresentantRecherche) SetQualite(v string) {
	o.Qualite = &v
}

// GetPersonneMorale returns the PersonneMorale field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetPersonneMorale() bool {
	if o == nil || IsNil(o.PersonneMorale) {
		var ret bool
		return ret
	}
	return *o.PersonneMorale
}

// GetPersonneMoraleOk returns a tuple with the PersonneMorale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetPersonneMoraleOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonneMorale) {
		return nil, false
	}
	return o.PersonneMorale, true
}

// HasPersonneMorale returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasPersonneMorale() bool {
	if o != nil && !IsNil(o.PersonneMorale) {
		return true
	}

	return false
}

// SetPersonneMorale gets a reference to the given bool and assigns it to the PersonneMorale field.
func (o *RepresentantRecherche) SetPersonneMorale(v bool) {
	o.PersonneMorale = &v
}

// GetDatePriseDePoste returns the DatePriseDePoste field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetDatePriseDePoste() string {
	if o == nil || IsNil(o.DatePriseDePoste) {
		var ret string
		return ret
	}
	return *o.DatePriseDePoste
}

// GetDatePriseDePosteOk returns a tuple with the DatePriseDePoste field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetDatePriseDePosteOk() (*string, bool) {
	if o == nil || IsNil(o.DatePriseDePoste) {
		return nil, false
	}
	return o.DatePriseDePoste, true
}

// HasDatePriseDePoste returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasDatePriseDePoste() bool {
	if o != nil && !IsNil(o.DatePriseDePoste) {
		return true
	}

	return false
}

// SetDatePriseDePoste gets a reference to the given string and assigns it to the DatePriseDePoste field.
func (o *RepresentantRecherche) SetDatePriseDePoste(v string) {
	o.DatePriseDePoste = &v
}

// GetDenomination returns the Denomination field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetDenomination() string {
	if o == nil || IsNil(o.Denomination) {
		var ret string
		return ret
	}
	return *o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetDenominationOk() (*string, bool) {
	if o == nil || IsNil(o.Denomination) {
		return nil, false
	}
	return o.Denomination, true
}

// HasDenomination returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasDenomination() bool {
	if o != nil && !IsNil(o.Denomination) {
		return true
	}

	return false
}

// SetDenomination gets a reference to the given string and assigns it to the Denomination field.
func (o *RepresentantRecherche) SetDenomination(v string) {
	o.Denomination = &v
}

// GetSiren returns the Siren field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetSiren() string {
	if o == nil || IsNil(o.Siren) {
		var ret string
		return ret
	}
	return *o.Siren
}

// GetSirenOk returns a tuple with the Siren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetSirenOk() (*string, bool) {
	if o == nil || IsNil(o.Siren) {
		return nil, false
	}
	return o.Siren, true
}

// HasSiren returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasSiren() bool {
	if o != nil && !IsNil(o.Siren) {
		return true
	}

	return false
}

// SetSiren gets a reference to the given string and assigns it to the Siren field.
func (o *RepresentantRecherche) SetSiren(v string) {
	o.Siren = &v
}

// GetFormeJuridique returns the FormeJuridique field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetFormeJuridique() string {
	if o == nil || IsNil(o.FormeJuridique) {
		var ret string
		return ret
	}
	return *o.FormeJuridique
}

// GetFormeJuridiqueOk returns a tuple with the FormeJuridique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetFormeJuridiqueOk() (*string, bool) {
	if o == nil || IsNil(o.FormeJuridique) {
		return nil, false
	}
	return o.FormeJuridique, true
}

// HasFormeJuridique returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasFormeJuridique() bool {
	if o != nil && !IsNil(o.FormeJuridique) {
		return true
	}

	return false
}

// SetFormeJuridique gets a reference to the given string and assigns it to the FormeJuridique field.
func (o *RepresentantRecherche) SetFormeJuridique(v string) {
	o.FormeJuridique = &v
}

// GetSexe returns the Sexe field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetSexe() string {
	if o == nil || IsNil(o.Sexe) {
		var ret string
		return ret
	}
	return *o.Sexe
}

// GetSexeOk returns a tuple with the Sexe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetSexeOk() (*string, bool) {
	if o == nil || IsNil(o.Sexe) {
		return nil, false
	}
	return o.Sexe, true
}

// HasSexe returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasSexe() bool {
	if o != nil && !IsNil(o.Sexe) {
		return true
	}

	return false
}

// SetSexe gets a reference to the given string and assigns it to the Sexe field.
func (o *RepresentantRecherche) SetSexe(v string) {
	o.Sexe = &v
}

// GetNom returns the Nom field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetNom() string {
	if o == nil || IsNil(o.Nom) {
		var ret string
		return ret
	}
	return *o.Nom
}

// GetNomOk returns a tuple with the Nom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetNomOk() (*string, bool) {
	if o == nil || IsNil(o.Nom) {
		return nil, false
	}
	return o.Nom, true
}

// HasNom returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasNom() bool {
	if o != nil && !IsNil(o.Nom) {
		return true
	}

	return false
}

// SetNom gets a reference to the given string and assigns it to the Nom field.
func (o *RepresentantRecherche) SetNom(v string) {
	o.Nom = &v
}

// GetPrenom returns the Prenom field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetPrenom() string {
	if o == nil || IsNil(o.Prenom) {
		var ret string
		return ret
	}
	return *o.Prenom
}

// GetPrenomOk returns a tuple with the Prenom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetPrenomOk() (*string, bool) {
	if o == nil || IsNil(o.Prenom) {
		return nil, false
	}
	return o.Prenom, true
}

// HasPrenom returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasPrenom() bool {
	if o != nil && !IsNil(o.Prenom) {
		return true
	}

	return false
}

// SetPrenom gets a reference to the given string and assigns it to the Prenom field.
func (o *RepresentantRecherche) SetPrenom(v string) {
	o.Prenom = &v
}

// GetPrenomUsuel returns the PrenomUsuel field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetPrenomUsuel() string {
	if o == nil || IsNil(o.PrenomUsuel) {
		var ret string
		return ret
	}
	return *o.PrenomUsuel
}

// GetPrenomUsuelOk returns a tuple with the PrenomUsuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetPrenomUsuelOk() (*string, bool) {
	if o == nil || IsNil(o.PrenomUsuel) {
		return nil, false
	}
	return o.PrenomUsuel, true
}

// HasPrenomUsuel returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasPrenomUsuel() bool {
	if o != nil && !IsNil(o.PrenomUsuel) {
		return true
	}

	return false
}

// SetPrenomUsuel gets a reference to the given string and assigns it to the PrenomUsuel field.
func (o *RepresentantRecherche) SetPrenomUsuel(v string) {
	o.PrenomUsuel = &v
}

// GetNomComplet returns the NomComplet field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetNomComplet() string {
	if o == nil || IsNil(o.NomComplet) {
		var ret string
		return ret
	}
	return *o.NomComplet
}

// GetNomCompletOk returns a tuple with the NomComplet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetNomCompletOk() (*string, bool) {
	if o == nil || IsNil(o.NomComplet) {
		return nil, false
	}
	return o.NomComplet, true
}

// HasNomComplet returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasNomComplet() bool {
	if o != nil && !IsNil(o.NomComplet) {
		return true
	}

	return false
}

// SetNomComplet gets a reference to the given string and assigns it to the NomComplet field.
func (o *RepresentantRecherche) SetNomComplet(v string) {
	o.NomComplet = &v
}

// GetDateDeNaissance returns the DateDeNaissance field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetDateDeNaissance() string {
	if o == nil || IsNil(o.DateDeNaissance) {
		var ret string
		return ret
	}
	return *o.DateDeNaissance
}

// GetDateDeNaissanceOk returns a tuple with the DateDeNaissance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetDateDeNaissanceOk() (*string, bool) {
	if o == nil || IsNil(o.DateDeNaissance) {
		return nil, false
	}
	return o.DateDeNaissance, true
}

// HasDateDeNaissance returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasDateDeNaissance() bool {
	if o != nil && !IsNil(o.DateDeNaissance) {
		return true
	}

	return false
}

// SetDateDeNaissance gets a reference to the given string and assigns it to the DateDeNaissance field.
func (o *RepresentantRecherche) SetDateDeNaissance(v string) {
	o.DateDeNaissance = &v
}

// GetDateDeNaissanceFormate returns the DateDeNaissanceFormate field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetDateDeNaissanceFormate() string {
	if o == nil || IsNil(o.DateDeNaissanceFormate) {
		var ret string
		return ret
	}
	return *o.DateDeNaissanceFormate
}

// GetDateDeNaissanceFormateOk returns a tuple with the DateDeNaissanceFormate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetDateDeNaissanceFormateOk() (*string, bool) {
	if o == nil || IsNil(o.DateDeNaissanceFormate) {
		return nil, false
	}
	return o.DateDeNaissanceFormate, true
}

// HasDateDeNaissanceFormate returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasDateDeNaissanceFormate() bool {
	if o != nil && !IsNil(o.DateDeNaissanceFormate) {
		return true
	}

	return false
}

// SetDateDeNaissanceFormate gets a reference to the given string and assigns it to the DateDeNaissanceFormate field.
func (o *RepresentantRecherche) SetDateDeNaissanceFormate(v string) {
	o.DateDeNaissanceFormate = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetAge() int32 {
	if o == nil || IsNil(o.Age) {
		var ret int32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given int32 and assigns it to the Age field.
func (o *RepresentantRecherche) SetAge(v int32) {
	o.Age = &v
}

// GetNationalite returns the Nationalite field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetNationalite() string {
	if o == nil || IsNil(o.Nationalite) {
		var ret string
		return ret
	}
	return *o.Nationalite
}

// GetNationaliteOk returns a tuple with the Nationalite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetNationaliteOk() (*string, bool) {
	if o == nil || IsNil(o.Nationalite) {
		return nil, false
	}
	return o.Nationalite, true
}

// HasNationalite returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasNationalite() bool {
	if o != nil && !IsNil(o.Nationalite) {
		return true
	}

	return false
}

// SetNationalite gets a reference to the given string and assigns it to the Nationalite field.
func (o *RepresentantRecherche) SetNationalite(v string) {
	o.Nationalite = &v
}

// GetCodeNationalite returns the CodeNationalite field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetCodeNationalite() string {
	if o == nil || IsNil(o.CodeNationalite) {
		var ret string
		return ret
	}
	return *o.CodeNationalite
}

// GetCodeNationaliteOk returns a tuple with the CodeNationalite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetCodeNationaliteOk() (*string, bool) {
	if o == nil || IsNil(o.CodeNationalite) {
		return nil, false
	}
	return o.CodeNationalite, true
}

// HasCodeNationalite returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasCodeNationalite() bool {
	if o != nil && !IsNil(o.CodeNationalite) {
		return true
	}

	return false
}

// SetCodeNationalite gets a reference to the given string and assigns it to the CodeNationalite field.
func (o *RepresentantRecherche) SetCodeNationalite(v string) {
	o.CodeNationalite = &v
}

// GetVilleDeNaissance returns the VilleDeNaissance field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetVilleDeNaissance() string {
	if o == nil || IsNil(o.VilleDeNaissance) {
		var ret string
		return ret
	}
	return *o.VilleDeNaissance
}

// GetVilleDeNaissanceOk returns a tuple with the VilleDeNaissance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetVilleDeNaissanceOk() (*string, bool) {
	if o == nil || IsNil(o.VilleDeNaissance) {
		return nil, false
	}
	return o.VilleDeNaissance, true
}

// HasVilleDeNaissance returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasVilleDeNaissance() bool {
	if o != nil && !IsNil(o.VilleDeNaissance) {
		return true
	}

	return false
}

// SetVilleDeNaissance gets a reference to the given string and assigns it to the VilleDeNaissance field.
func (o *RepresentantRecherche) SetVilleDeNaissance(v string) {
	o.VilleDeNaissance = &v
}

// GetPaysDeNaissance returns the PaysDeNaissance field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetPaysDeNaissance() string {
	if o == nil || IsNil(o.PaysDeNaissance) {
		var ret string
		return ret
	}
	return *o.PaysDeNaissance
}

// GetPaysDeNaissanceOk returns a tuple with the PaysDeNaissance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetPaysDeNaissanceOk() (*string, bool) {
	if o == nil || IsNil(o.PaysDeNaissance) {
		return nil, false
	}
	return o.PaysDeNaissance, true
}

// HasPaysDeNaissance returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasPaysDeNaissance() bool {
	if o != nil && !IsNil(o.PaysDeNaissance) {
		return true
	}

	return false
}

// SetPaysDeNaissance gets a reference to the given string and assigns it to the PaysDeNaissance field.
func (o *RepresentantRecherche) SetPaysDeNaissance(v string) {
	o.PaysDeNaissance = &v
}

// GetCodePaysDeNaissance returns the CodePaysDeNaissance field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetCodePaysDeNaissance() string {
	if o == nil || IsNil(o.CodePaysDeNaissance) {
		var ret string
		return ret
	}
	return *o.CodePaysDeNaissance
}

// GetCodePaysDeNaissanceOk returns a tuple with the CodePaysDeNaissance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetCodePaysDeNaissanceOk() (*string, bool) {
	if o == nil || IsNil(o.CodePaysDeNaissance) {
		return nil, false
	}
	return o.CodePaysDeNaissance, true
}

// HasCodePaysDeNaissance returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasCodePaysDeNaissance() bool {
	if o != nil && !IsNil(o.CodePaysDeNaissance) {
		return true
	}

	return false
}

// SetCodePaysDeNaissance gets a reference to the given string and assigns it to the CodePaysDeNaissance field.
func (o *RepresentantRecherche) SetCodePaysDeNaissance(v string) {
	o.CodePaysDeNaissance = &v
}

// GetAdresseLigne1 returns the AdresseLigne1 field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetAdresseLigne1() string {
	if o == nil || IsNil(o.AdresseLigne1) {
		var ret string
		return ret
	}
	return *o.AdresseLigne1
}

// GetAdresseLigne1Ok returns a tuple with the AdresseLigne1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetAdresseLigne1Ok() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne1) {
		return nil, false
	}
	return o.AdresseLigne1, true
}

// HasAdresseLigne1 returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasAdresseLigne1() bool {
	if o != nil && !IsNil(o.AdresseLigne1) {
		return true
	}

	return false
}

// SetAdresseLigne1 gets a reference to the given string and assigns it to the AdresseLigne1 field.
func (o *RepresentantRecherche) SetAdresseLigne1(v string) {
	o.AdresseLigne1 = &v
}

// GetAdresseLigne2 returns the AdresseLigne2 field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetAdresseLigne2() string {
	if o == nil || IsNil(o.AdresseLigne2) {
		var ret string
		return ret
	}
	return *o.AdresseLigne2
}

// GetAdresseLigne2Ok returns a tuple with the AdresseLigne2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetAdresseLigne2Ok() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne2) {
		return nil, false
	}
	return o.AdresseLigne2, true
}

// HasAdresseLigne2 returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasAdresseLigne2() bool {
	if o != nil && !IsNil(o.AdresseLigne2) {
		return true
	}

	return false
}

// SetAdresseLigne2 gets a reference to the given string and assigns it to the AdresseLigne2 field.
func (o *RepresentantRecherche) SetAdresseLigne2(v string) {
	o.AdresseLigne2 = &v
}

// GetAdresseLigne3 returns the AdresseLigne3 field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetAdresseLigne3() string {
	if o == nil || IsNil(o.AdresseLigne3) {
		var ret string
		return ret
	}
	return *o.AdresseLigne3
}

// GetAdresseLigne3Ok returns a tuple with the AdresseLigne3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetAdresseLigne3Ok() (*string, bool) {
	if o == nil || IsNil(o.AdresseLigne3) {
		return nil, false
	}
	return o.AdresseLigne3, true
}

// HasAdresseLigne3 returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasAdresseLigne3() bool {
	if o != nil && !IsNil(o.AdresseLigne3) {
		return true
	}

	return false
}

// SetAdresseLigne3 gets a reference to the given string and assigns it to the AdresseLigne3 field.
func (o *RepresentantRecherche) SetAdresseLigne3(v string) {
	o.AdresseLigne3 = &v
}

// GetCodePostal returns the CodePostal field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetCodePostal() string {
	if o == nil || IsNil(o.CodePostal) {
		var ret string
		return ret
	}
	return *o.CodePostal
}

// GetCodePostalOk returns a tuple with the CodePostal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetCodePostalOk() (*string, bool) {
	if o == nil || IsNil(o.CodePostal) {
		return nil, false
	}
	return o.CodePostal, true
}

// HasCodePostal returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasCodePostal() bool {
	if o != nil && !IsNil(o.CodePostal) {
		return true
	}

	return false
}

// SetCodePostal gets a reference to the given string and assigns it to the CodePostal field.
func (o *RepresentantRecherche) SetCodePostal(v string) {
	o.CodePostal = &v
}

// GetVille returns the Ville field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetVille() string {
	if o == nil || IsNil(o.Ville) {
		var ret string
		return ret
	}
	return *o.Ville
}

// GetVilleOk returns a tuple with the Ville field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetVilleOk() (*string, bool) {
	if o == nil || IsNil(o.Ville) {
		return nil, false
	}
	return o.Ville, true
}

// HasVille returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasVille() bool {
	if o != nil && !IsNil(o.Ville) {
		return true
	}

	return false
}

// SetVille gets a reference to the given string and assigns it to the Ville field.
func (o *RepresentantRecherche) SetVille(v string) {
	o.Ville = &v
}

// GetPays returns the Pays field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetPays() string {
	if o == nil || IsNil(o.Pays) {
		var ret string
		return ret
	}
	return *o.Pays
}

// GetPaysOk returns a tuple with the Pays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetPaysOk() (*string, bool) {
	if o == nil || IsNil(o.Pays) {
		return nil, false
	}
	return o.Pays, true
}

// HasPays returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasPays() bool {
	if o != nil && !IsNil(o.Pays) {
		return true
	}

	return false
}

// SetPays gets a reference to the given string and assigns it to the Pays field.
func (o *RepresentantRecherche) SetPays(v string) {
	o.Pays = &v
}

// GetCodePays returns the CodePays field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetCodePays() string {
	if o == nil || IsNil(o.CodePays) {
		var ret string
		return ret
	}
	return *o.CodePays
}

// GetCodePaysOk returns a tuple with the CodePays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetCodePaysOk() (*string, bool) {
	if o == nil || IsNil(o.CodePays) {
		return nil, false
	}
	return o.CodePays, true
}

// HasCodePays returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasCodePays() bool {
	if o != nil && !IsNil(o.CodePays) {
		return true
	}

	return false
}

// SetCodePays gets a reference to the given string and assigns it to the CodePays field.
func (o *RepresentantRecherche) SetCodePays(v string) {
	o.CodePays = &v
}

// GetActuel returns the Actuel field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetActuel() bool {
	if o == nil || IsNil(o.Actuel) {
		var ret bool
		return ret
	}
	return *o.Actuel
}

// GetActuelOk returns a tuple with the Actuel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetActuelOk() (*bool, bool) {
	if o == nil || IsNil(o.Actuel) {
		return nil, false
	}
	return o.Actuel, true
}

// HasActuel returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasActuel() bool {
	if o != nil && !IsNil(o.Actuel) {
		return true
	}

	return false
}

// SetActuel gets a reference to the given bool and assigns it to the Actuel field.
func (o *RepresentantRecherche) SetActuel(v bool) {
	o.Actuel = &v
}

// GetDateDepartDePoste returns the DateDepartDePoste field value if set, zero value otherwise.
func (o *RepresentantRecherche) GetDateDepartDePoste() string {
	if o == nil || IsNil(o.DateDepartDePoste) {
		var ret string
		return ret
	}
	return *o.DateDepartDePoste
}

// GetDateDepartDePosteOk returns a tuple with the DateDepartDePoste field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepresentantRecherche) GetDateDepartDePosteOk() (*string, bool) {
	if o == nil || IsNil(o.DateDepartDePoste) {
		return nil, false
	}
	return o.DateDepartDePoste, true
}

// HasDateDepartDePoste returns a boolean if a field has been set.
func (o *RepresentantRecherche) HasDateDepartDePoste() bool {
	if o != nil && !IsNil(o.DateDepartDePoste) {
		return true
	}

	return false
}

// SetDateDepartDePoste gets a reference to the given string and assigns it to the DateDepartDePoste field.
func (o *RepresentantRecherche) SetDateDepartDePoste(v string) {
	o.DateDepartDePoste = &v
}

func (o RepresentantRecherche) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepresentantRecherche) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qualite) {
		toSerialize["qualite"] = o.Qualite
	}
	if !IsNil(o.PersonneMorale) {
		toSerialize["personne_morale"] = o.PersonneMorale
	}
	if !IsNil(o.DatePriseDePoste) {
		toSerialize["date_prise_de_poste"] = o.DatePriseDePoste
	}
	if !IsNil(o.Denomination) {
		toSerialize["denomination"] = o.Denomination
	}
	if !IsNil(o.Siren) {
		toSerialize["siren"] = o.Siren
	}
	if !IsNil(o.FormeJuridique) {
		toSerialize["forme_juridique"] = o.FormeJuridique
	}
	if !IsNil(o.Sexe) {
		toSerialize["sexe"] = o.Sexe
	}
	if !IsNil(o.Nom) {
		toSerialize["nom"] = o.Nom
	}
	if !IsNil(o.Prenom) {
		toSerialize["prenom"] = o.Prenom
	}
	if !IsNil(o.PrenomUsuel) {
		toSerialize["prenom_usuel"] = o.PrenomUsuel
	}
	if !IsNil(o.NomComplet) {
		toSerialize["nom_complet"] = o.NomComplet
	}
	if !IsNil(o.DateDeNaissance) {
		toSerialize["date_de_naissance"] = o.DateDeNaissance
	}
	if !IsNil(o.DateDeNaissanceFormate) {
		toSerialize["date_de_naissance_formate"] = o.DateDeNaissanceFormate
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.Nationalite) {
		toSerialize["nationalite"] = o.Nationalite
	}
	if !IsNil(o.CodeNationalite) {
		toSerialize["code_nationalite"] = o.CodeNationalite
	}
	if !IsNil(o.VilleDeNaissance) {
		toSerialize["ville_de_naissance"] = o.VilleDeNaissance
	}
	if !IsNil(o.PaysDeNaissance) {
		toSerialize["pays_de_naissance"] = o.PaysDeNaissance
	}
	if !IsNil(o.CodePaysDeNaissance) {
		toSerialize["code_pays_de_naissance"] = o.CodePaysDeNaissance
	}
	if !IsNil(o.AdresseLigne1) {
		toSerialize["adresse_ligne_1"] = o.AdresseLigne1
	}
	if !IsNil(o.AdresseLigne2) {
		toSerialize["adresse_ligne_2"] = o.AdresseLigne2
	}
	if !IsNil(o.AdresseLigne3) {
		toSerialize["adresse_ligne_3"] = o.AdresseLigne3
	}
	if !IsNil(o.CodePostal) {
		toSerialize["code_postal"] = o.CodePostal
	}
	if !IsNil(o.Ville) {
		toSerialize["ville"] = o.Ville
	}
	if !IsNil(o.Pays) {
		toSerialize["pays"] = o.Pays
	}
	if !IsNil(o.CodePays) {
		toSerialize["code_pays"] = o.CodePays
	}
	if !IsNil(o.Actuel) {
		toSerialize["actuel"] = o.Actuel
	}
	if !IsNil(o.DateDepartDePoste) {
		toSerialize["date_depart_de_poste"] = o.DateDepartDePoste
	}
	return toSerialize, nil
}

type NullableRepresentantRecherche struct {
	value *RepresentantRecherche
	isSet bool
}

func (v NullableRepresentantRecherche) Get() *RepresentantRecherche {
	return v.value
}

func (v *NullableRepresentantRecherche) Set(val *RepresentantRecherche) {
	v.value = val
	v.isSet = true
}

func (v NullableRepresentantRecherche) IsSet() bool {
	return v.isSet
}

func (v *NullableRepresentantRecherche) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepresentantRecherche(val *RepresentantRecherche) *NullableRepresentantRecherche {
	return &NullableRepresentantRecherche{value: val, isSet: true}
}

func (v NullableRepresentantRecherche) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepresentantRecherche) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


