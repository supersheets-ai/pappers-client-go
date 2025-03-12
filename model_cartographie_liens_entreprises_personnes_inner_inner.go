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
	"fmt"
	"gopkg.in/validator.v2"
)

// CartographieLiensEntreprisesPersonnesInnerInner - struct for CartographieLiensEntreprisesPersonnesInnerInner
type CartographieLiensEntreprisesPersonnesInnerInner struct {
	String *string
}

// stringAsCartographieLiensEntreprisesPersonnesInnerInner is a convenience function that returns string wrapped in CartographieLiensEntreprisesPersonnesInnerInner
func StringAsCartographieLiensEntreprisesPersonnesInnerInner(v *string) CartographieLiensEntreprisesPersonnesInnerInner {
	return CartographieLiensEntreprisesPersonnesInnerInner{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CartographieLiensEntreprisesPersonnesInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CartographieLiensEntreprisesPersonnesInnerInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CartographieLiensEntreprisesPersonnesInnerInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CartographieLiensEntreprisesPersonnesInnerInner) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CartographieLiensEntreprisesPersonnesInnerInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CartographieLiensEntreprisesPersonnesInnerInner) GetActualInstanceValue() (interface{}) {
	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCartographieLiensEntreprisesPersonnesInnerInner struct {
	value *CartographieLiensEntreprisesPersonnesInnerInner
	isSet bool
}

func (v NullableCartographieLiensEntreprisesPersonnesInnerInner) Get() *CartographieLiensEntreprisesPersonnesInnerInner {
	return v.value
}

func (v *NullableCartographieLiensEntreprisesPersonnesInnerInner) Set(val *CartographieLiensEntreprisesPersonnesInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCartographieLiensEntreprisesPersonnesInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCartographieLiensEntreprisesPersonnesInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartographieLiensEntreprisesPersonnesInnerInner(val *CartographieLiensEntreprisesPersonnesInnerInner) *NullableCartographieLiensEntreprisesPersonnesInnerInner {
	return &NullableCartographieLiensEntreprisesPersonnesInnerInner{value: val, isSet: true}
}

func (v NullableCartographieLiensEntreprisesPersonnesInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartographieLiensEntreprisesPersonnesInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


