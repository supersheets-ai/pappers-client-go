# \RechercheAPI

All URIs are relative to *https://api.pappers.fr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Recherche**](RechercheAPI.md#Recherche) | **Get** /recherche | Recherche les entreprises qui correspondent à des critères.
[**RechercheBeneficiaires**](RechercheAPI.md#RechercheBeneficiaires) | **Get** /recherche-beneficiaires | Recherche les bénéficiaires effectifs qui correspondent à des critères.
[**RechercheDirigeants**](RechercheAPI.md#RechercheDirigeants) | **Get** /recherche-dirigeants | Recherche les dirigeants qui correspondent à des critères.
[**RechercheDocuments**](RechercheAPI.md#RechercheDocuments) | **Get** /recherche-documents | Recherche les documents qui correspondent à des critères.
[**RecherchePublications**](RechercheAPI.md#RecherchePublications) | **Get** /recherche-publications | Recherche les publications BODACC qui correspondent à des critères.



## Recherche

> Recherche200Response Recherche(ctx).Page(page).ParPage(parPage).Curseur(curseur).ParCurseur(parCurseur).Bases(bases).Precision(precision).Q(q).Siege(siege).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Execute()

Recherche les entreprises qui correspondent à des critères.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	page := int32(1) // int32 | Page de résultats. Valeur par défaut : `1`. (optional)
	parPage := int32(20) // int32 | Nombre de résultats affichés sur une page. Valeur par défaut : `10`. (optional)
	curseur := "*" // string | Curseur servant à parcourir l'ensemble des résultats (alternativement à la pagination qui est limitée à 400 résultats maximum). Doit valoir `*` pour la première requête, et doit pour les requêtes suivantes reprendre la valeur `curseurSuivant` retournée par la dernière réponse. (optional)
	parCurseur := int32(20) // int32 | Nombre de résultats affichés par curseur. Valeur par défaut : `50`. Valeur minimale: `1`. Valeur maximale : `1000`. (optional)
	bases := "entreprises" // string | Bases de données dans lesquelles rechercher. Il est possible d'indiquer plusieurs bases en les séparant par des virgules. Valeur par défaut : `\"entreprises\"`. (optional)
	precision := "standard" // string | Niveau de précision de la recherche. Valeur par défaut : `\"standard\"`. (optional)
	q := "Google France" // string | Texte à rechercher. Dénomination pour une personne morale, nom et prénom pour une personne physique. Si vous recherchez dans plusieurs bases, ce paramètre sera utilisé pour rechercher dans toutes les bases. (optional)
	siege := "true" // string | Défini si la requête se base sur le siège (optional)
	codeNaf := "70.10Z" // string | Code NAF de l'entreprise. Il est possible d'indiquer plusieurs codes NAF en les séparant par des virgules. (optional)
	departement := "75" // string | Numéro de département de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs départements en les séparant par des virgules. (optional)
	region := "11" // string | Code de la région de l'un des établissements de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d'indiquer plusieurs codes régions en les séparant par des virgules. (optional)
	codePostal := "75009" // string | Code postal de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs codes postaux en les séparant par des virgules. (optional)
	conventionCollective := "1486" // string | Convention collective de l'entreprise. (optional)
	categorieJuridique := "5499" // string | Catégorie juridique de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l'INSEE, à l'exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. (optional)
	entrepriseCessee := false // bool | Activité de l'entreprise cessée ou non. (optional)
	statutRcs := "inscrit" // string | Statut au RCS (optional)
	objetSocial := "La conception de moteurs de recherche sur internet." // string | Objet social de l'entreprise déclaré au RCS. (optional)
	dateImmatriculationRcsMin := "15-05-2002" // string | Date d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateImmatriculationRcsMax := "15-05-2002" // string | d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMin := "15-05-2002" // string | Date de radiation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMax := "17-05-2002" // string | Date de radiation au RCS maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	capitalMin := "411016200" // string | Capital minimum de l'entreprise. (optional)
	capitalMax := "411016400" // string | Capital maximum de l'entreprise. (optional)
	chiffreAffairesMin := "411016200" // string | Chiffre d'affaires minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	chiffreAffairesMax := "411016400" // string | Chiffre d'affaires maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMin := "29327000" // string | Résultat minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMax := "29327100" // string | Résultat maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	dateCreationMin := "15-05-2002" // string | Date de création minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateCreationMax := "17-05-2002" // string | Date de création maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	trancheEffectifMin := "40" // string | Tranche d'effectifs minimale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	trancheEffectifMax := "42" // string | Tranche d'effectifs maximale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	typeDirigeant := "physique" // string | Type du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	qualiteDirigeant := "Administrateur" // string | Qualité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nationaliteDirigeant := "Française" // string | Nationalité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nomDirigeant := "Yi" // string | Nom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	prenomDirigeant := "Kenneth H." // string | Prénom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMin := int32(40) // int32 | Âge minimal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMax := int32(42) // int32 | Âge maximal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceDirigeantMin := time.Now() // string | Date de naissance minimale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceDirigeantMax := time.Now() // string | Date de naissance maximale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	ageBeneficiaireMin := int32(40) // int32 | Âge minimal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	ageBeneficiaireMax := int32(42) // int32 | Âge maximal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceBeneficiaireMin := time.Now() // string | Date de naissance minimale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceBeneficiaireMax := time.Now() // string | Date de naissance maximale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	nationaliteBeneficiaire := "Française" // string | Nationalité du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDepotDocumentMin := time.Now() // string | Date de dépôt minimale du document, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMax := time.Now() // string | Date de dépôt maximale du document, au format JJ-MM-AAAA. (optional)
	typePublication := "Procédure collective" // string | Type de publication (optional)
	datePublicationMin := time.Now() // string | Date publication minimale de la publication, au format JJ-MM-AAAA. (optional)
	datePublicationMax := time.Now() // string | Date de publication maximale de la publication, au format JJ-MM-AAAA. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RechercheAPI.Recherche(context.Background()).Page(page).ParPage(parPage).Curseur(curseur).ParCurseur(parCurseur).Bases(bases).Precision(precision).Q(q).Siege(siege).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RechercheAPI.Recherche``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Recherche`: Recherche200Response
	fmt.Fprintf(os.Stdout, "Response from `RechercheAPI.Recherche`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page de résultats. Valeur par défaut : &#x60;1&#x60;. | 
 **parPage** | **int32** | Nombre de résultats affichés sur une page. Valeur par défaut : &#x60;10&#x60;. | 
 **curseur** | **string** | Curseur servant à parcourir l&#39;ensemble des résultats (alternativement à la pagination qui est limitée à 400 résultats maximum). Doit valoir &#x60;*&#x60; pour la première requête, et doit pour les requêtes suivantes reprendre la valeur &#x60;curseurSuivant&#x60; retournée par la dernière réponse. | 
 **parCurseur** | **int32** | Nombre de résultats affichés par curseur. Valeur par défaut : &#x60;50&#x60;. Valeur minimale: &#x60;1&#x60;. Valeur maximale : &#x60;1000&#x60;. | 
 **bases** | **string** | Bases de données dans lesquelles rechercher. Il est possible d&#39;indiquer plusieurs bases en les séparant par des virgules. Valeur par défaut : &#x60;\&quot;entreprises\&quot;&#x60;. | 
 **precision** | **string** | Niveau de précision de la recherche. Valeur par défaut : &#x60;\&quot;standard\&quot;&#x60;. | 
 **q** | **string** | Texte à rechercher. Dénomination pour une personne morale, nom et prénom pour une personne physique. Si vous recherchez dans plusieurs bases, ce paramètre sera utilisé pour rechercher dans toutes les bases. | 
 **siege** | **string** | Défini si la requête se base sur le siège | 
 **codeNaf** | **string** | Code NAF de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes NAF en les séparant par des virgules. | 
 **departement** | **string** | Numéro de département de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs départements en les séparant par des virgules. | 
 **region** | **string** | Code de la région de l&#39;un des établissements de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d&#39;indiquer plusieurs codes régions en les séparant par des virgules. | 
 **codePostal** | **string** | Code postal de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes postaux en les séparant par des virgules. | 
 **conventionCollective** | **string** | Convention collective de l&#39;entreprise. | 
 **categorieJuridique** | **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | 
 **entrepriseCessee** | **bool** | Activité de l&#39;entreprise cessée ou non. | 
 **statutRcs** | **string** | Statut au RCS | 
 **objetSocial** | **string** | Objet social de l&#39;entreprise déclaré au RCS. | 
 **dateImmatriculationRcsMin** | **string** | Date d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateImmatriculationRcsMax** | **string** | d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMin** | **string** | Date de radiation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMax** | **string** | Date de radiation au RCS maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **capitalMin** | **string** | Capital minimum de l&#39;entreprise. | 
 **capitalMax** | **string** | Capital maximum de l&#39;entreprise. | 
 **chiffreAffairesMin** | **string** | Chiffre d&#39;affaires minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **chiffreAffairesMax** | **string** | Chiffre d&#39;affaires maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMin** | **string** | Résultat minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMax** | **string** | Résultat maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **dateCreationMin** | **string** | Date de création minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateCreationMax** | **string** | Date de création maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **trancheEffectifMin** | **string** | Tranche d&#39;effectifs minimale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **trancheEffectifMax** | **string** | Tranche d&#39;effectifs maximale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **typeDirigeant** | **string** | Type du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **qualiteDirigeant** | **string** | Qualité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nationaliteDirigeant** | **string** | Nationalité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nomDirigeant** | **string** | Nom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **prenomDirigeant** | **string** | Prénom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMin** | **int32** | Âge minimal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMax** | **int32** | Âge maximal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceDirigeantMin** | **string** | Date de naissance minimale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceDirigeantMax** | **string** | Date de naissance maximale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **ageBeneficiaireMin** | **int32** | Âge minimal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageBeneficiaireMax** | **int32** | Âge maximal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceBeneficiaireMin** | **string** | Date de naissance minimale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceBeneficiaireMax** | **string** | Date de naissance maximale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **nationaliteBeneficiaire** | **string** | Nationalité du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDepotDocumentMin** | **string** | Date de dépôt minimale du document, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMax** | **string** | Date de dépôt maximale du document, au format JJ-MM-AAAA. | 
 **typePublication** | **string** | Type de publication | 
 **datePublicationMin** | **string** | Date publication minimale de la publication, au format JJ-MM-AAAA. | 
 **datePublicationMax** | **string** | Date de publication maximale de la publication, au format JJ-MM-AAAA. | 

### Return type

[**Recherche200Response**](Recherche200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercheBeneficiaires

> RechercheBeneficiaires200Response RechercheBeneficiaires(ctx).ParPage(parPage).Page(page).Precision(precision).Q(q).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()

Recherche les bénéficiaires effectifs qui correspondent à des critères.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	parPage := int32(20) // int32 | Nombre de résultats affichés sur une page. Valeur par défaut : `10`. (optional)
	page := int32(1) // int32 | Page de résultats. Valeur par défaut : `1`. (optional)
	precision := "standard" // string | Niveau de précision de la recherche. Valeur par défaut : `\"standard\"`. (optional)
	q := "Xavier Niel" // string | Nom et/ou prénom du bénéficiaire effectif. (optional)
	ageBeneficiaireMin := int32(40) // int32 | Âge minimal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	ageBeneficiaireMax := int32(42) // int32 | Âge maximal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceBeneficiaireMin := time.Now() // string | Date de naissance minimale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceBeneficiaireMax := time.Now() // string | Date de naissance maximale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	nationaliteBeneficiaire := "Française" // string | Nationalité du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	codeNaf := "70.10Z" // string | Code NAF de l'entreprise. Il est possible d'indiquer plusieurs codes NAF en les séparant par des virgules. (optional)
	departement := "75" // string | Numéro de département de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs départements en les séparant par des virgules. (optional)
	region := "11" // string | Code de la région de l'un des établissements de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d'indiquer plusieurs codes régions en les séparant par des virgules. (optional)
	codePostal := "75009" // string | Code postal de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs codes postaux en les séparant par des virgules. (optional)
	conventionCollective := "1486" // string | Convention collective de l'entreprise. (optional)
	categorieJuridique := "5499" // string | Catégorie juridique de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l'INSEE, à l'exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. (optional)
	entrepriseCessee := false // bool | Activité de l'entreprise cessée ou non. (optional)
	statutRcs := "inscrit" // string | Statut au RCS (optional)
	objetSocial := "La conception de moteurs de recherche sur internet." // string | Objet social de l'entreprise déclaré au RCS. (optional)
	dateImmatriculationRcsMin := "15-05-2002" // string | Date d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateImmatriculationRcsMax := "15-05-2002" // string | d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMin := "15-05-2002" // string | Date de radiation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMax := "17-05-2002" // string | Date de radiation au RCS maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	capitalMin := "411016200" // string | Capital minimum de l'entreprise. (optional)
	capitalMax := "411016400" // string | Capital maximum de l'entreprise. (optional)
	chiffreAffairesMin := "411016200" // string | Chiffre d'affaires minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	chiffreAffairesMax := "411016400" // string | Chiffre d'affaires maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMin := "29327000" // string | Résultat minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMax := "29327100" // string | Résultat maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	dateCreationMin := "15-05-2002" // string | Date de création minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateCreationMax := "17-05-2002" // string | Date de création maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	trancheEffectifMin := "40" // string | Tranche d'effectifs minimale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	trancheEffectifMax := "42" // string | Tranche d'effectifs maximale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	typeDirigeant := "physique" // string | Type du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	qualiteDirigeant := "Administrateur" // string | Qualité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nationaliteDirigeant := "Française" // string | Nationalité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nomDirigeant := "Yi" // string | Nom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	prenomDirigeant := "Kenneth H." // string | Prénom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMin := int32(40) // int32 | Âge minimal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMax := int32(42) // int32 | Âge maximal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceDirigeantMin := time.Now() // string | Date de naissance minimale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceDirigeantMax := time.Now() // string | Date de naissance maximale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMin := time.Now() // string | Date de dépôt minimale du document, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMax := time.Now() // string | Date de dépôt maximale du document, au format JJ-MM-AAAA. (optional)
	typePublication := "Procédure collective" // string | Type de publication (optional)
	datePublicationMin := time.Now() // string | Date publication minimale de la publication, au format JJ-MM-AAAA. (optional)
	datePublicationMax := time.Now() // string | Date de publication maximale de la publication, au format JJ-MM-AAAA. (optional)
	siren := "siren_example" // string | SIREN de l'entreprise. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RechercheAPI.RechercheBeneficiaires(context.Background()).ParPage(parPage).Page(page).Precision(precision).Q(q).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RechercheAPI.RechercheBeneficiaires``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercheBeneficiaires`: RechercheBeneficiaires200Response
	fmt.Fprintf(os.Stdout, "Response from `RechercheAPI.RechercheBeneficiaires`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercheBeneficiairesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parPage** | **int32** | Nombre de résultats affichés sur une page. Valeur par défaut : &#x60;10&#x60;. | 
 **page** | **int32** | Page de résultats. Valeur par défaut : &#x60;1&#x60;. | 
 **precision** | **string** | Niveau de précision de la recherche. Valeur par défaut : &#x60;\&quot;standard\&quot;&#x60;. | 
 **q** | **string** | Nom et/ou prénom du bénéficiaire effectif. | 
 **ageBeneficiaireMin** | **int32** | Âge minimal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageBeneficiaireMax** | **int32** | Âge maximal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceBeneficiaireMin** | **string** | Date de naissance minimale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceBeneficiaireMax** | **string** | Date de naissance maximale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **nationaliteBeneficiaire** | **string** | Nationalité du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **codeNaf** | **string** | Code NAF de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes NAF en les séparant par des virgules. | 
 **departement** | **string** | Numéro de département de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs départements en les séparant par des virgules. | 
 **region** | **string** | Code de la région de l&#39;un des établissements de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d&#39;indiquer plusieurs codes régions en les séparant par des virgules. | 
 **codePostal** | **string** | Code postal de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes postaux en les séparant par des virgules. | 
 **conventionCollective** | **string** | Convention collective de l&#39;entreprise. | 
 **categorieJuridique** | **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | 
 **entrepriseCessee** | **bool** | Activité de l&#39;entreprise cessée ou non. | 
 **statutRcs** | **string** | Statut au RCS | 
 **objetSocial** | **string** | Objet social de l&#39;entreprise déclaré au RCS. | 
 **dateImmatriculationRcsMin** | **string** | Date d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateImmatriculationRcsMax** | **string** | d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMin** | **string** | Date de radiation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMax** | **string** | Date de radiation au RCS maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **capitalMin** | **string** | Capital minimum de l&#39;entreprise. | 
 **capitalMax** | **string** | Capital maximum de l&#39;entreprise. | 
 **chiffreAffairesMin** | **string** | Chiffre d&#39;affaires minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **chiffreAffairesMax** | **string** | Chiffre d&#39;affaires maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMin** | **string** | Résultat minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMax** | **string** | Résultat maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **dateCreationMin** | **string** | Date de création minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateCreationMax** | **string** | Date de création maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **trancheEffectifMin** | **string** | Tranche d&#39;effectifs minimale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **trancheEffectifMax** | **string** | Tranche d&#39;effectifs maximale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **typeDirigeant** | **string** | Type du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **qualiteDirigeant** | **string** | Qualité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nationaliteDirigeant** | **string** | Nationalité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nomDirigeant** | **string** | Nom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **prenomDirigeant** | **string** | Prénom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMin** | **int32** | Âge minimal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMax** | **int32** | Âge maximal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceDirigeantMin** | **string** | Date de naissance minimale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceDirigeantMax** | **string** | Date de naissance maximale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMin** | **string** | Date de dépôt minimale du document, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMax** | **string** | Date de dépôt maximale du document, au format JJ-MM-AAAA. | 
 **typePublication** | **string** | Type de publication | 
 **datePublicationMin** | **string** | Date publication minimale de la publication, au format JJ-MM-AAAA. | 
 **datePublicationMax** | **string** | Date de publication maximale de la publication, au format JJ-MM-AAAA. | 
 **siren** | **string** | SIREN de l&#39;entreprise. | 

### Return type

[**RechercheBeneficiaires200Response**](RechercheBeneficiaires200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercheDirigeants

> RechercheDirigeants200Response RechercheDirigeants(ctx).ParPage(parPage).Page(page).Precision(precision).Q(q).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()

Recherche les dirigeants qui correspondent à des critères.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	parPage := int32(20) // int32 | Nombre de résultats affichés sur une page. Valeur par défaut : `10`. (optional)
	page := int32(1) // int32 | Page de résultats. Valeur par défaut : `1`. (optional)
	precision := "standard" // string | Niveau de précision de la recherche. Valeur par défaut : `\"standard\"`. (optional)
	q := "Google France" // string | Texte à rechercher. Nom et prénom du dirigeant pour une personne physique, dénomination pour une personne morale. (optional)
	typeDirigeant := "physique" // string | Type du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	qualiteDirigeant := "Administrateur" // string | Qualité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nationaliteDirigeant := "Française" // string | Nationalité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nomDirigeant := "Yi" // string | Nom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	prenomDirigeant := "Kenneth H." // string | Prénom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMin := int32(40) // int32 | Âge minimal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMax := int32(42) // int32 | Âge maximal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceDirigeantMin := time.Now() // string | Date de naissance minimale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceDirigeantMax := time.Now() // string | Date de naissance maximale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	codeNaf := "70.10Z" // string | Code NAF de l'entreprise. Il est possible d'indiquer plusieurs codes NAF en les séparant par des virgules. (optional)
	departement := "75" // string | Numéro de département de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs départements en les séparant par des virgules. (optional)
	region := "11" // string | Code de la région de l'un des établissements de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d'indiquer plusieurs codes régions en les séparant par des virgules. (optional)
	codePostal := "75009" // string | Code postal de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs codes postaux en les séparant par des virgules. (optional)
	conventionCollective := "1486" // string | Convention collective de l'entreprise. (optional)
	categorieJuridique := "5499" // string | Catégorie juridique de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l'INSEE, à l'exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. (optional)
	entrepriseCessee := false // bool | Activité de l'entreprise cessée ou non. (optional)
	statutRcs := "inscrit" // string | Statut au RCS (optional)
	objetSocial := "La conception de moteurs de recherche sur internet." // string | Objet social de l'entreprise déclaré au RCS. (optional)
	dateImmatriculationRcsMin := "15-05-2002" // string | Date d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateImmatriculationRcsMax := "15-05-2002" // string | d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMin := "15-05-2002" // string | Date de radiation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMax := "17-05-2002" // string | Date de radiation au RCS maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	capitalMin := "411016200" // string | Capital minimum de l'entreprise. (optional)
	capitalMax := "411016400" // string | Capital maximum de l'entreprise. (optional)
	chiffreAffairesMin := "411016200" // string | Chiffre d'affaires minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	chiffreAffairesMax := "411016400" // string | Chiffre d'affaires maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMin := "29327000" // string | Résultat minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMax := "29327100" // string | Résultat maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	dateCreationMin := "15-05-2002" // string | Date de création minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateCreationMax := "17-05-2002" // string | Date de création maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	trancheEffectifMin := "40" // string | Tranche d'effectifs minimale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	trancheEffectifMax := "42" // string | Tranche d'effectifs maximale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	ageBeneficiaireMin := int32(40) // int32 | Âge minimal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	ageBeneficiaireMax := int32(42) // int32 | Âge maximal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceBeneficiaireMin := time.Now() // string | Date de naissance minimale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceBeneficiaireMax := time.Now() // string | Date de naissance maximale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	nationaliteBeneficiaire := "Française" // string | Nationalité du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDepotDocumentMin := time.Now() // string | Date de dépôt minimale du document, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMax := time.Now() // string | Date de dépôt maximale du document, au format JJ-MM-AAAA. (optional)
	typePublication := "Procédure collective" // string | Type de publication (optional)
	datePublicationMin := time.Now() // string | Date publication minimale de la publication, au format JJ-MM-AAAA. (optional)
	datePublicationMax := time.Now() // string | Date de publication maximale de la publication, au format JJ-MM-AAAA. (optional)
	siren := "siren_example" // string | SIREN de l'entreprise. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RechercheAPI.RechercheDirigeants(context.Background()).ParPage(parPage).Page(page).Precision(precision).Q(q).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RechercheAPI.RechercheDirigeants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercheDirigeants`: RechercheDirigeants200Response
	fmt.Fprintf(os.Stdout, "Response from `RechercheAPI.RechercheDirigeants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercheDirigeantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parPage** | **int32** | Nombre de résultats affichés sur une page. Valeur par défaut : &#x60;10&#x60;. | 
 **page** | **int32** | Page de résultats. Valeur par défaut : &#x60;1&#x60;. | 
 **precision** | **string** | Niveau de précision de la recherche. Valeur par défaut : &#x60;\&quot;standard\&quot;&#x60;. | 
 **q** | **string** | Texte à rechercher. Nom et prénom du dirigeant pour une personne physique, dénomination pour une personne morale. | 
 **typeDirigeant** | **string** | Type du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **qualiteDirigeant** | **string** | Qualité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nationaliteDirigeant** | **string** | Nationalité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nomDirigeant** | **string** | Nom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **prenomDirigeant** | **string** | Prénom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMin** | **int32** | Âge minimal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMax** | **int32** | Âge maximal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceDirigeantMin** | **string** | Date de naissance minimale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceDirigeantMax** | **string** | Date de naissance maximale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **codeNaf** | **string** | Code NAF de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes NAF en les séparant par des virgules. | 
 **departement** | **string** | Numéro de département de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs départements en les séparant par des virgules. | 
 **region** | **string** | Code de la région de l&#39;un des établissements de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d&#39;indiquer plusieurs codes régions en les séparant par des virgules. | 
 **codePostal** | **string** | Code postal de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes postaux en les séparant par des virgules. | 
 **conventionCollective** | **string** | Convention collective de l&#39;entreprise. | 
 **categorieJuridique** | **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | 
 **entrepriseCessee** | **bool** | Activité de l&#39;entreprise cessée ou non. | 
 **statutRcs** | **string** | Statut au RCS | 
 **objetSocial** | **string** | Objet social de l&#39;entreprise déclaré au RCS. | 
 **dateImmatriculationRcsMin** | **string** | Date d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateImmatriculationRcsMax** | **string** | d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMin** | **string** | Date de radiation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMax** | **string** | Date de radiation au RCS maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **capitalMin** | **string** | Capital minimum de l&#39;entreprise. | 
 **capitalMax** | **string** | Capital maximum de l&#39;entreprise. | 
 **chiffreAffairesMin** | **string** | Chiffre d&#39;affaires minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **chiffreAffairesMax** | **string** | Chiffre d&#39;affaires maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMin** | **string** | Résultat minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMax** | **string** | Résultat maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **dateCreationMin** | **string** | Date de création minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateCreationMax** | **string** | Date de création maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **trancheEffectifMin** | **string** | Tranche d&#39;effectifs minimale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **trancheEffectifMax** | **string** | Tranche d&#39;effectifs maximale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **ageBeneficiaireMin** | **int32** | Âge minimal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageBeneficiaireMax** | **int32** | Âge maximal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceBeneficiaireMin** | **string** | Date de naissance minimale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceBeneficiaireMax** | **string** | Date de naissance maximale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **nationaliteBeneficiaire** | **string** | Nationalité du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDepotDocumentMin** | **string** | Date de dépôt minimale du document, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMax** | **string** | Date de dépôt maximale du document, au format JJ-MM-AAAA. | 
 **typePublication** | **string** | Type de publication | 
 **datePublicationMin** | **string** | Date publication minimale de la publication, au format JJ-MM-AAAA. | 
 **datePublicationMax** | **string** | Date de publication maximale de la publication, au format JJ-MM-AAAA. | 
 **siren** | **string** | SIREN de l&#39;entreprise. | 

### Return type

[**RechercheDirigeants200Response**](RechercheDirigeants200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RechercheDocuments

> RechercheDocuments200Response RechercheDocuments(ctx).ParPage(parPage).Page(page).Precision(precision).Q(q).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()

Recherche les documents qui correspondent à des critères.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	parPage := int32(20) // int32 | Nombre de résultats affichés sur une page. Valeur par défaut : `10`. (optional)
	page := int32(1) // int32 | Page de résultats. Valeur par défaut : `1`. (optional)
	precision := "standard" // string | Niveau de précision de la recherche. Valeur par défaut : `\"standard\"`. (optional)
	q := "Rémunération" // string | Mot-clé à rechercher dans le contenu du document. (optional)
	dateDepotDocumentMin := time.Now() // string | Date de dépôt minimale du document, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMax := time.Now() // string | Date de dépôt maximale du document, au format JJ-MM-AAAA. (optional)
	codeNaf := "70.10Z" // string | Code NAF de l'entreprise. Il est possible d'indiquer plusieurs codes NAF en les séparant par des virgules. (optional)
	departement := "75" // string | Numéro de département de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs départements en les séparant par des virgules. (optional)
	region := "11" // string | Code de la région de l'un des établissements de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d'indiquer plusieurs codes régions en les séparant par des virgules. (optional)
	codePostal := "75009" // string | Code postal de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs codes postaux en les séparant par des virgules. (optional)
	conventionCollective := "1486" // string | Convention collective de l'entreprise. (optional)
	categorieJuridique := "5499" // string | Catégorie juridique de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l'INSEE, à l'exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. (optional)
	entrepriseCessee := false // bool | Activité de l'entreprise cessée ou non. (optional)
	statutRcs := "inscrit" // string | Statut au RCS (optional)
	objetSocial := "La conception de moteurs de recherche sur internet." // string | Objet social de l'entreprise déclaré au RCS. (optional)
	dateImmatriculationRcsMin := "15-05-2002" // string | Date d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateImmatriculationRcsMax := "15-05-2002" // string | d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMin := "15-05-2002" // string | Date de radiation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMax := "17-05-2002" // string | Date de radiation au RCS maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	capitalMin := "411016200" // string | Capital minimum de l'entreprise. (optional)
	capitalMax := "411016400" // string | Capital maximum de l'entreprise. (optional)
	chiffreAffairesMin := "411016200" // string | Chiffre d'affaires minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	chiffreAffairesMax := "411016400" // string | Chiffre d'affaires maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMin := "29327000" // string | Résultat minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMax := "29327100" // string | Résultat maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	dateCreationMin := "15-05-2002" // string | Date de création minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateCreationMax := "17-05-2002" // string | Date de création maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	trancheEffectifMin := "40" // string | Tranche d'effectifs minimale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	trancheEffectifMax := "42" // string | Tranche d'effectifs maximale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	typeDirigeant := "physique" // string | Type du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	qualiteDirigeant := "Administrateur" // string | Qualité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nationaliteDirigeant := "Française" // string | Nationalité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nomDirigeant := "Yi" // string | Nom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	prenomDirigeant := "Kenneth H." // string | Prénom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMin := int32(40) // int32 | Âge minimal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMax := int32(42) // int32 | Âge maximal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceDirigeantMin := time.Now() // string | Date de naissance minimale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceDirigeantMax := time.Now() // string | Date de naissance maximale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	ageBeneficiaireMin := int32(40) // int32 | Âge minimal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	ageBeneficiaireMax := int32(42) // int32 | Âge maximal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceBeneficiaireMin := time.Now() // string | Date de naissance minimale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceBeneficiaireMax := time.Now() // string | Date de naissance maximale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	nationaliteBeneficiaire := "Française" // string | Nationalité du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	typePublication := "Procédure collective" // string | Type de publication (optional)
	datePublicationMin := time.Now() // string | Date publication minimale de la publication, au format JJ-MM-AAAA. (optional)
	datePublicationMax := time.Now() // string | Date de publication maximale de la publication, au format JJ-MM-AAAA. (optional)
	siren := "siren_example" // string | SIREN de l'entreprise. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RechercheAPI.RechercheDocuments(context.Background()).ParPage(parPage).Page(page).Precision(precision).Q(q).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RechercheAPI.RechercheDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RechercheDocuments`: RechercheDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `RechercheAPI.RechercheDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRechercheDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parPage** | **int32** | Nombre de résultats affichés sur une page. Valeur par défaut : &#x60;10&#x60;. | 
 **page** | **int32** | Page de résultats. Valeur par défaut : &#x60;1&#x60;. | 
 **precision** | **string** | Niveau de précision de la recherche. Valeur par défaut : &#x60;\&quot;standard\&quot;&#x60;. | 
 **q** | **string** | Mot-clé à rechercher dans le contenu du document. | 
 **dateDepotDocumentMin** | **string** | Date de dépôt minimale du document, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMax** | **string** | Date de dépôt maximale du document, au format JJ-MM-AAAA. | 
 **codeNaf** | **string** | Code NAF de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes NAF en les séparant par des virgules. | 
 **departement** | **string** | Numéro de département de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs départements en les séparant par des virgules. | 
 **region** | **string** | Code de la région de l&#39;un des établissements de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d&#39;indiquer plusieurs codes régions en les séparant par des virgules. | 
 **codePostal** | **string** | Code postal de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes postaux en les séparant par des virgules. | 
 **conventionCollective** | **string** | Convention collective de l&#39;entreprise. | 
 **categorieJuridique** | **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | 
 **entrepriseCessee** | **bool** | Activité de l&#39;entreprise cessée ou non. | 
 **statutRcs** | **string** | Statut au RCS | 
 **objetSocial** | **string** | Objet social de l&#39;entreprise déclaré au RCS. | 
 **dateImmatriculationRcsMin** | **string** | Date d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateImmatriculationRcsMax** | **string** | d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMin** | **string** | Date de radiation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMax** | **string** | Date de radiation au RCS maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **capitalMin** | **string** | Capital minimum de l&#39;entreprise. | 
 **capitalMax** | **string** | Capital maximum de l&#39;entreprise. | 
 **chiffreAffairesMin** | **string** | Chiffre d&#39;affaires minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **chiffreAffairesMax** | **string** | Chiffre d&#39;affaires maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMin** | **string** | Résultat minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMax** | **string** | Résultat maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **dateCreationMin** | **string** | Date de création minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateCreationMax** | **string** | Date de création maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **trancheEffectifMin** | **string** | Tranche d&#39;effectifs minimale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **trancheEffectifMax** | **string** | Tranche d&#39;effectifs maximale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **typeDirigeant** | **string** | Type du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **qualiteDirigeant** | **string** | Qualité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nationaliteDirigeant** | **string** | Nationalité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nomDirigeant** | **string** | Nom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **prenomDirigeant** | **string** | Prénom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMin** | **int32** | Âge minimal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMax** | **int32** | Âge maximal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceDirigeantMin** | **string** | Date de naissance minimale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceDirigeantMax** | **string** | Date de naissance maximale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **ageBeneficiaireMin** | **int32** | Âge minimal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageBeneficiaireMax** | **int32** | Âge maximal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceBeneficiaireMin** | **string** | Date de naissance minimale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceBeneficiaireMax** | **string** | Date de naissance maximale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **nationaliteBeneficiaire** | **string** | Nationalité du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **typePublication** | **string** | Type de publication | 
 **datePublicationMin** | **string** | Date publication minimale de la publication, au format JJ-MM-AAAA. | 
 **datePublicationMax** | **string** | Date de publication maximale de la publication, au format JJ-MM-AAAA. | 
 **siren** | **string** | SIREN de l&#39;entreprise. | 

### Return type

[**RechercheDocuments200Response**](RechercheDocuments200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecherchePublications

> RecherchePublications200Response RecherchePublications(ctx).ParPage(parPage).Page(page).Precision(precision).Q(q).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()

Recherche les publications BODACC qui correspondent à des critères.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/supersheets-ai/pappers-client-go"
)

func main() {
	parPage := int32(20) // int32 | Nombre de résultats affichés sur une page. Valeur par défaut : `10`. (optional)
	page := int32(1) // int32 | Page de résultats. Valeur par défaut : `1`. (optional)
	precision := "standard" // string | Niveau de précision de la recherche. Valeur par défaut : `\"standard\"`. (optional)
	q := "Liquidation" // string | Mot-clé à rechercher dans le contenu de la publication. (optional)
	codeNaf := "70.10Z" // string | Code NAF de l'entreprise. Il est possible d'indiquer plusieurs codes NAF en les séparant par des virgules. (optional)
	departement := "75" // string | Numéro de département de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs départements en les séparant par des virgules. (optional)
	region := "11" // string | Code de la région de l'un des établissements de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d'indiquer plusieurs codes régions en les séparant par des virgules. (optional)
	codePostal := "75009" // string | Code postal de l'un des établissements de l'entreprise. Il est possible d'indiquer plusieurs codes postaux en les séparant par des virgules. (optional)
	conventionCollective := "1486" // string | Convention collective de l'entreprise. (optional)
	categorieJuridique := "5499" // string | Catégorie juridique de l'entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l'INSEE, à l'exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. (optional)
	entrepriseCessee := false // bool | Activité de l'entreprise cessée ou non. (optional)
	statutRcs := "inscrit" // string | Statut au RCS (optional)
	objetSocial := "La conception de moteurs de recherche sur internet." // string | Objet social de l'entreprise déclaré au RCS. (optional)
	dateImmatriculationRcsMin := "15-05-2002" // string | Date d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateImmatriculationRcsMax := "15-05-2002" // string | d'immatriculation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMin := "15-05-2002" // string | Date de radiation au RCS minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateRadiationRcsMax := "17-05-2002" // string | Date de radiation au RCS maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	capitalMin := "411016200" // string | Capital minimum de l'entreprise. (optional)
	capitalMax := "411016400" // string | Capital maximum de l'entreprise. (optional)
	chiffreAffairesMin := "411016200" // string | Chiffre d'affaires minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	chiffreAffairesMax := "411016400" // string | Chiffre d'affaires maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMin := "29327000" // string | Résultat minimum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	resultatMax := "29327100" // string | Résultat maximum de l'entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d'office toutes les entreprises dont les comptes ne sont pas publiés. (optional)
	dateCreationMin := "15-05-2002" // string | Date de création minimale de l'entreprise, au format JJ-MM-AAAA. (optional)
	dateCreationMax := "17-05-2002" // string | Date de création maximale de l'entreprise, au format JJ-MM-AAAA. (optional)
	trancheEffectifMin := "40" // string | Tranche d'effectifs minimale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	trancheEffectifMax := "42" // string | Tranche d'effectifs maximale de l'entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur (optional)
	typeDirigeant := "physique" // string | Type du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	qualiteDirigeant := "Administrateur" // string | Qualité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nationaliteDirigeant := "Française" // string | Nationalité du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	nomDirigeant := "Yi" // string | Nom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	prenomDirigeant := "Kenneth H." // string | Prénom du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMin := int32(40) // int32 | Âge minimal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	ageDirigeantMax := int32(42) // int32 | Âge maximal du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceDirigeantMin := time.Now() // string | Date de naissance minimale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceDirigeantMax := time.Now() // string | Date de naissance maximale du dirigeant (ou de l'un des dirigeants de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	ageBeneficiaireMin := int32(40) // int32 | Âge minimal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	ageBeneficiaireMax := int32(42) // int32 | Âge maximal du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDeNaissanceBeneficiaireMin := time.Now() // string | Date de naissance minimale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises), au format JJ-MM-AAAA. (optional)
	dateDeNaissanceBeneficiaireMax := time.Now() // string | Date de naissance maximale du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises) de l'entreprise, au format JJ-MM-AAAA. (optional)
	nationaliteBeneficiaire := "Française" // string | Nationalité du bénéficiaire effectif (ou de l'un des bénéficiaires effectifs de l'entreprise pour une recherche d'entreprises). (optional)
	dateDepotDocumentMin := time.Now() // string | Date de dépôt minimale du document, au format JJ-MM-AAAA. (optional)
	dateDepotDocumentMax := time.Now() // string | Date de dépôt maximale du document, au format JJ-MM-AAAA. (optional)
	typePublication := "Procédure collective" // string | Type de publication (optional)
	datePublicationMin := time.Now() // string | Date publication minimale de la publication, au format JJ-MM-AAAA. (optional)
	datePublicationMax := time.Now() // string | Date de publication maximale de la publication, au format JJ-MM-AAAA. (optional)
	siren := "siren_example" // string | SIREN de l'entreprise. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RechercheAPI.RecherchePublications(context.Background()).ParPage(parPage).Page(page).Precision(precision).Q(q).CodeNaf(codeNaf).Departement(departement).Region(region).CodePostal(codePostal).ConventionCollective(conventionCollective).CategorieJuridique(categorieJuridique).EntrepriseCessee(entrepriseCessee).StatutRcs(statutRcs).ObjetSocial(objetSocial).DateImmatriculationRcsMin(dateImmatriculationRcsMin).DateImmatriculationRcsMax(dateImmatriculationRcsMax).DateRadiationRcsMin(dateRadiationRcsMin).DateRadiationRcsMax(dateRadiationRcsMax).CapitalMin(capitalMin).CapitalMax(capitalMax).ChiffreAffairesMin(chiffreAffairesMin).ChiffreAffairesMax(chiffreAffairesMax).ResultatMin(resultatMin).ResultatMax(resultatMax).DateCreationMin(dateCreationMin).DateCreationMax(dateCreationMax).TrancheEffectifMin(trancheEffectifMin).TrancheEffectifMax(trancheEffectifMax).TypeDirigeant(typeDirigeant).QualiteDirigeant(qualiteDirigeant).NationaliteDirigeant(nationaliteDirigeant).NomDirigeant(nomDirigeant).PrenomDirigeant(prenomDirigeant).AgeDirigeantMin(ageDirigeantMin).AgeDirigeantMax(ageDirigeantMax).DateDeNaissanceDirigeantMin(dateDeNaissanceDirigeantMin).DateDeNaissanceDirigeantMax(dateDeNaissanceDirigeantMax).AgeBeneficiaireMin(ageBeneficiaireMin).AgeBeneficiaireMax(ageBeneficiaireMax).DateDeNaissanceBeneficiaireMin(dateDeNaissanceBeneficiaireMin).DateDeNaissanceBeneficiaireMax(dateDeNaissanceBeneficiaireMax).NationaliteBeneficiaire(nationaliteBeneficiaire).DateDepotDocumentMin(dateDepotDocumentMin).DateDepotDocumentMax(dateDepotDocumentMax).TypePublication(typePublication).DatePublicationMin(datePublicationMin).DatePublicationMax(datePublicationMax).Siren(siren).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RechercheAPI.RecherchePublications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecherchePublications`: RecherchePublications200Response
	fmt.Fprintf(os.Stdout, "Response from `RechercheAPI.RecherchePublications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecherchePublicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parPage** | **int32** | Nombre de résultats affichés sur une page. Valeur par défaut : &#x60;10&#x60;. | 
 **page** | **int32** | Page de résultats. Valeur par défaut : &#x60;1&#x60;. | 
 **precision** | **string** | Niveau de précision de la recherche. Valeur par défaut : &#x60;\&quot;standard\&quot;&#x60;. | 
 **q** | **string** | Mot-clé à rechercher dans le contenu de la publication. | 
 **codeNaf** | **string** | Code NAF de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes NAF en les séparant par des virgules. | 
 **departement** | **string** | Numéro de département de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs départements en les séparant par des virgules. | 
 **region** | **string** | Code de la région de l&#39;un des établissements de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/4316069#titre-bloc-18). Il est possible d&#39;indiquer plusieurs codes régions en les séparant par des virgules. | 
 **codePostal** | **string** | Code postal de l&#39;un des établissements de l&#39;entreprise. Il est possible d&#39;indiquer plusieurs codes postaux en les séparant par des virgules. | 
 **conventionCollective** | **string** | Convention collective de l&#39;entreprise. | 
 **categorieJuridique** | **string** | Catégorie juridique de l&#39;entreprise, selon la [nomenclature Insee](https://www.insee.fr/fr/information/2028129).   **Note** : Le code correspond à celui de l&#39;INSEE, à l&#39;exception des SASU qui auront comme code 5720 et les EURL qui auront comme code 5498. | 
 **entrepriseCessee** | **bool** | Activité de l&#39;entreprise cessée ou non. | 
 **statutRcs** | **string** | Statut au RCS | 
 **objetSocial** | **string** | Objet social de l&#39;entreprise déclaré au RCS. | 
 **dateImmatriculationRcsMin** | **string** | Date d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateImmatriculationRcsMax** | **string** | d&#39;immatriculation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMin** | **string** | Date de radiation au RCS minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateRadiationRcsMax** | **string** | Date de radiation au RCS maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **capitalMin** | **string** | Capital minimum de l&#39;entreprise. | 
 **capitalMax** | **string** | Capital maximum de l&#39;entreprise. | 
 **chiffreAffairesMin** | **string** | Chiffre d&#39;affaires minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **chiffreAffairesMax** | **string** | Chiffre d&#39;affaires maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMin** | **string** | Résultat minimum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **resultatMax** | **string** | Résultat maximum de l&#39;entreprise.  **Note** : Filtrer sur ce critère restreint énormément les entreprises retournées car cela élimine d&#39;office toutes les entreprises dont les comptes ne sont pas publiés. | 
 **dateCreationMin** | **string** | Date de création minimale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **dateCreationMax** | **string** | Date de création maximale de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **trancheEffectifMin** | **string** | Tranche d&#39;effectifs minimale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **trancheEffectifMax** | **string** | Tranche d&#39;effectifs maximale de l&#39;entreprise, selon la [nomenclature Sirene](https://www.sirene.fr/static-resources/documentation/v_sommaire_311.htm#73).  **Note** : 00 ou NN donneront les mêmes résultats et veulent dire non employeur | 
 **typeDirigeant** | **string** | Type du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **qualiteDirigeant** | **string** | Qualité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nationaliteDirigeant** | **string** | Nationalité du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **nomDirigeant** | **string** | Nom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **prenomDirigeant** | **string** | Prénom du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMin** | **int32** | Âge minimal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageDirigeantMax** | **int32** | Âge maximal du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceDirigeantMin** | **string** | Date de naissance minimale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceDirigeantMax** | **string** | Date de naissance maximale du dirigeant (ou de l&#39;un des dirigeants de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **ageBeneficiaireMin** | **int32** | Âge minimal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **ageBeneficiaireMax** | **int32** | Âge maximal du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDeNaissanceBeneficiaireMin** | **string** | Date de naissance minimale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises), au format JJ-MM-AAAA. | 
 **dateDeNaissanceBeneficiaireMax** | **string** | Date de naissance maximale du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises) de l&#39;entreprise, au format JJ-MM-AAAA. | 
 **nationaliteBeneficiaire** | **string** | Nationalité du bénéficiaire effectif (ou de l&#39;un des bénéficiaires effectifs de l&#39;entreprise pour une recherche d&#39;entreprises). | 
 **dateDepotDocumentMin** | **string** | Date de dépôt minimale du document, au format JJ-MM-AAAA. | 
 **dateDepotDocumentMax** | **string** | Date de dépôt maximale du document, au format JJ-MM-AAAA. | 
 **typePublication** | **string** | Type de publication | 
 **datePublicationMin** | **string** | Date publication minimale de la publication, au format JJ-MM-AAAA. | 
 **datePublicationMax** | **string** | Date de publication maximale de la publication, au format JJ-MM-AAAA. | 
 **siren** | **string** | SIREN de l&#39;entreprise. | 

### Return type

[**RecherchePublications200Response**](RecherchePublications200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

