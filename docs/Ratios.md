# Ratios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChiffreAffaires** | Pointer to **float32** | Chiffre d&#39;affaires de l&#39;entreprise. | [optional] 
**Resultat** | Pointer to **float32** | Résultat de l&#39;entreprise. | [optional] 
**Effectif** | Pointer to **int32** | Effectif de l&#39;entreprise. | [optional] 
**MargeBrute** | Pointer to **float32** | Marge brute de l&#39;entreprise. | [optional] 
**ExcedentBrutExploitation** | Pointer to **float32** | Excédent brut d&#39;exploitation (EBITDA) de l&#39;entreprise. | [optional] 
**ResultatExploitation** | Pointer to **float32** | Résultat d&#39;exploitation (EBIT) de l&#39;entreprise. | [optional] 
**TauxCroissanceChiffreAffaires** | Pointer to **float32** | Taux de croissance du chiffre d&#39;affaires (en %) de l&#39;entreprise. | [optional] 
**TauxMargeBrute** | Pointer to **float32** | Taux de marge brute (en %) de l&#39;entreprise. | [optional] 
**TauxMargeEBITDA** | Pointer to **float32** | Taux de marge d&#39;EBITDA (en %) de l&#39;entreprise. | [optional] 
**TauxMargeOperationnelle** | Pointer to **float32** | Taux de marge opérationnelle (EBIT) (en %) de l&#39;entreprise. | [optional] 
**BFR** | Pointer to **float32** | BFR (Besoin en fonds de roulement) de l&#39;entreprise. | [optional] 
**BFRExploitation** | Pointer to **float32** | BFR exploitation de l&#39;entreprise. | [optional] 
**BFRHorsExploitation** | Pointer to **float32** | BFR hors exploitation de l&#39;entreprise. | [optional] 
**BFRJoursCA** | Pointer to **float32** | BFR (en jours de CA) de l&#39;entreprise. | [optional] 
**BFRExploitationJoursCA** | Pointer to **float32** | BFR exploitation (en jours de CA) de l&#39;entreprise. | [optional] 
**BFRHorsExploitationJoursCA** | Pointer to **float32** | BFR hors exploitation (en jours de CA) de l&#39;entreprise. | [optional] 
**DelaiPaiementClientsJours** | Pointer to **float32** | Délai de paiement clients (en jours) de l&#39;entreprise. | [optional] 
**DelaiPaiementFournisseursJours** | Pointer to **float32** | Délai de paiement fournisseurs (en jours) de l&#39;entreprise. | [optional] 
**RatioStockCAJours** | Pointer to **float32** | Ratio des stocks / CA (en jours) de l&#39;entreprise. | [optional] 
**CapaciteAutofinancement** | Pointer to **float32** | Capacité d&#39;autofinancement de l&#39;entreprise. | [optional] 
**CapaciteAutofinancementCA** | Pointer to **float32** | Capacité d&#39;autofinancement / CA (en %) de l&#39;entreprise. | [optional] 
**FondsRoulementNetGlobal** | Pointer to **float32** | Fonds de roulement net global de l&#39;entreprise. | [optional] 
**CouvertureBFR** | Pointer to **float32** | Couverture du BFR de l&#39;entreprise. | [optional] 
**Tresorerie** | Pointer to **float32** | Trésorerie de l&#39;entreprise. | [optional] 
**DettesFinancieres** | Pointer to **float32** | Dettes financières de l&#39;entreprise. | [optional] 
**CapaciteRemboursement** | Pointer to **float32** | Capacité de remboursement de l&#39;entreprise. | [optional] 
**RatioEndettement** | Pointer to **float32** | Ratio d&#39;endettement (Gearing) de l&#39;entreprise. | [optional] 
**AutonomieFinanciere** | Pointer to **float32** | Autonomie financière (en %) de l&#39;entreprise. | [optional] 
**TauxLevier** | Pointer to **float32** | Taux de levier (DFN/EBITDA) de l&#39;entreprise. | [optional] 
**EtatDettes1AnAuPlus** | Pointer to **float32** | Etat des dettes à 1 an au plus de l&#39;entreprise. | [optional] 
**LiquiditeGenerale** | Pointer to **float32** | Liquidité générale de l&#39;entreprise. | [optional] 
**CouvertureDettes** | Pointer to **float32** | Couverture des dettes de l&#39;entreprise. | [optional] 
**MargeNette** | Pointer to **float32** | Marge nette (en %) de l&#39;entreprise. | [optional] 
**RentabiliteFondsPropres** | Pointer to **float32** | Rentabilité sur fonds propres (en %) de l&#39;entreprise. | [optional] 
**RentabiliteEconomique** | Pointer to **float32** | Rentabilité économique (en %) de l&#39;entreprise. | [optional] 
**ValeurAjoutee** | Pointer to **float32** | Valeur ajoutée de l&#39;entreprise. | [optional] 
**ValeurAjouteeCA** | Pointer to **float32** | Valeur ajoutée / CA (en %) de l&#39;entreprise. | [optional] 
**SalairesChargesSociales** | Pointer to **float32** | Salaires et charges sociales de l&#39;entreprise. | [optional] 
**SalairesCA** | Pointer to **float32** | Salaires / CA (en %) de l&#39;entreprise. | [optional] 
**ImpotsTaxes** | Pointer to **float32** | Impôts et taxes de l&#39;entreprise. | [optional] 

## Methods

### NewRatios

`func NewRatios() *Ratios`

NewRatios instantiates a new Ratios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatiosWithDefaults

`func NewRatiosWithDefaults() *Ratios`

NewRatiosWithDefaults instantiates a new Ratios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChiffreAffaires

`func (o *Ratios) GetChiffreAffaires() float32`

GetChiffreAffaires returns the ChiffreAffaires field if non-nil, zero value otherwise.

### GetChiffreAffairesOk

`func (o *Ratios) GetChiffreAffairesOk() (*float32, bool)`

GetChiffreAffairesOk returns a tuple with the ChiffreAffaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChiffreAffaires

`func (o *Ratios) SetChiffreAffaires(v float32)`

SetChiffreAffaires sets ChiffreAffaires field to given value.

### HasChiffreAffaires

`func (o *Ratios) HasChiffreAffaires() bool`

HasChiffreAffaires returns a boolean if a field has been set.

### GetResultat

`func (o *Ratios) GetResultat() float32`

GetResultat returns the Resultat field if non-nil, zero value otherwise.

### GetResultatOk

`func (o *Ratios) GetResultatOk() (*float32, bool)`

GetResultatOk returns a tuple with the Resultat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultat

`func (o *Ratios) SetResultat(v float32)`

SetResultat sets Resultat field to given value.

### HasResultat

`func (o *Ratios) HasResultat() bool`

HasResultat returns a boolean if a field has been set.

### GetEffectif

`func (o *Ratios) GetEffectif() int32`

GetEffectif returns the Effectif field if non-nil, zero value otherwise.

### GetEffectifOk

`func (o *Ratios) GetEffectifOk() (*int32, bool)`

GetEffectifOk returns a tuple with the Effectif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectif

`func (o *Ratios) SetEffectif(v int32)`

SetEffectif sets Effectif field to given value.

### HasEffectif

`func (o *Ratios) HasEffectif() bool`

HasEffectif returns a boolean if a field has been set.

### GetMargeBrute

`func (o *Ratios) GetMargeBrute() float32`

GetMargeBrute returns the MargeBrute field if non-nil, zero value otherwise.

### GetMargeBruteOk

`func (o *Ratios) GetMargeBruteOk() (*float32, bool)`

GetMargeBruteOk returns a tuple with the MargeBrute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargeBrute

`func (o *Ratios) SetMargeBrute(v float32)`

SetMargeBrute sets MargeBrute field to given value.

### HasMargeBrute

`func (o *Ratios) HasMargeBrute() bool`

HasMargeBrute returns a boolean if a field has been set.

### GetExcedentBrutExploitation

`func (o *Ratios) GetExcedentBrutExploitation() float32`

GetExcedentBrutExploitation returns the ExcedentBrutExploitation field if non-nil, zero value otherwise.

### GetExcedentBrutExploitationOk

`func (o *Ratios) GetExcedentBrutExploitationOk() (*float32, bool)`

GetExcedentBrutExploitationOk returns a tuple with the ExcedentBrutExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcedentBrutExploitation

`func (o *Ratios) SetExcedentBrutExploitation(v float32)`

SetExcedentBrutExploitation sets ExcedentBrutExploitation field to given value.

### HasExcedentBrutExploitation

`func (o *Ratios) HasExcedentBrutExploitation() bool`

HasExcedentBrutExploitation returns a boolean if a field has been set.

### GetResultatExploitation

`func (o *Ratios) GetResultatExploitation() float32`

GetResultatExploitation returns the ResultatExploitation field if non-nil, zero value otherwise.

### GetResultatExploitationOk

`func (o *Ratios) GetResultatExploitationOk() (*float32, bool)`

GetResultatExploitationOk returns a tuple with the ResultatExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultatExploitation

`func (o *Ratios) SetResultatExploitation(v float32)`

SetResultatExploitation sets ResultatExploitation field to given value.

### HasResultatExploitation

`func (o *Ratios) HasResultatExploitation() bool`

HasResultatExploitation returns a boolean if a field has been set.

### GetTauxCroissanceChiffreAffaires

`func (o *Ratios) GetTauxCroissanceChiffreAffaires() float32`

GetTauxCroissanceChiffreAffaires returns the TauxCroissanceChiffreAffaires field if non-nil, zero value otherwise.

### GetTauxCroissanceChiffreAffairesOk

`func (o *Ratios) GetTauxCroissanceChiffreAffairesOk() (*float32, bool)`

GetTauxCroissanceChiffreAffairesOk returns a tuple with the TauxCroissanceChiffreAffaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxCroissanceChiffreAffaires

`func (o *Ratios) SetTauxCroissanceChiffreAffaires(v float32)`

SetTauxCroissanceChiffreAffaires sets TauxCroissanceChiffreAffaires field to given value.

### HasTauxCroissanceChiffreAffaires

`func (o *Ratios) HasTauxCroissanceChiffreAffaires() bool`

HasTauxCroissanceChiffreAffaires returns a boolean if a field has been set.

### GetTauxMargeBrute

`func (o *Ratios) GetTauxMargeBrute() float32`

GetTauxMargeBrute returns the TauxMargeBrute field if non-nil, zero value otherwise.

### GetTauxMargeBruteOk

`func (o *Ratios) GetTauxMargeBruteOk() (*float32, bool)`

GetTauxMargeBruteOk returns a tuple with the TauxMargeBrute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxMargeBrute

`func (o *Ratios) SetTauxMargeBrute(v float32)`

SetTauxMargeBrute sets TauxMargeBrute field to given value.

### HasTauxMargeBrute

`func (o *Ratios) HasTauxMargeBrute() bool`

HasTauxMargeBrute returns a boolean if a field has been set.

### GetTauxMargeEBITDA

`func (o *Ratios) GetTauxMargeEBITDA() float32`

GetTauxMargeEBITDA returns the TauxMargeEBITDA field if non-nil, zero value otherwise.

### GetTauxMargeEBITDAOk

`func (o *Ratios) GetTauxMargeEBITDAOk() (*float32, bool)`

GetTauxMargeEBITDAOk returns a tuple with the TauxMargeEBITDA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxMargeEBITDA

`func (o *Ratios) SetTauxMargeEBITDA(v float32)`

SetTauxMargeEBITDA sets TauxMargeEBITDA field to given value.

### HasTauxMargeEBITDA

`func (o *Ratios) HasTauxMargeEBITDA() bool`

HasTauxMargeEBITDA returns a boolean if a field has been set.

### GetTauxMargeOperationnelle

`func (o *Ratios) GetTauxMargeOperationnelle() float32`

GetTauxMargeOperationnelle returns the TauxMargeOperationnelle field if non-nil, zero value otherwise.

### GetTauxMargeOperationnelleOk

`func (o *Ratios) GetTauxMargeOperationnelleOk() (*float32, bool)`

GetTauxMargeOperationnelleOk returns a tuple with the TauxMargeOperationnelle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxMargeOperationnelle

`func (o *Ratios) SetTauxMargeOperationnelle(v float32)`

SetTauxMargeOperationnelle sets TauxMargeOperationnelle field to given value.

### HasTauxMargeOperationnelle

`func (o *Ratios) HasTauxMargeOperationnelle() bool`

HasTauxMargeOperationnelle returns a boolean if a field has been set.

### GetBFR

`func (o *Ratios) GetBFR() float32`

GetBFR returns the BFR field if non-nil, zero value otherwise.

### GetBFROk

`func (o *Ratios) GetBFROk() (*float32, bool)`

GetBFROk returns a tuple with the BFR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFR

`func (o *Ratios) SetBFR(v float32)`

SetBFR sets BFR field to given value.

### HasBFR

`func (o *Ratios) HasBFR() bool`

HasBFR returns a boolean if a field has been set.

### GetBFRExploitation

`func (o *Ratios) GetBFRExploitation() float32`

GetBFRExploitation returns the BFRExploitation field if non-nil, zero value otherwise.

### GetBFRExploitationOk

`func (o *Ratios) GetBFRExploitationOk() (*float32, bool)`

GetBFRExploitationOk returns a tuple with the BFRExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFRExploitation

`func (o *Ratios) SetBFRExploitation(v float32)`

SetBFRExploitation sets BFRExploitation field to given value.

### HasBFRExploitation

`func (o *Ratios) HasBFRExploitation() bool`

HasBFRExploitation returns a boolean if a field has been set.

### GetBFRHorsExploitation

`func (o *Ratios) GetBFRHorsExploitation() float32`

GetBFRHorsExploitation returns the BFRHorsExploitation field if non-nil, zero value otherwise.

### GetBFRHorsExploitationOk

`func (o *Ratios) GetBFRHorsExploitationOk() (*float32, bool)`

GetBFRHorsExploitationOk returns a tuple with the BFRHorsExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFRHorsExploitation

`func (o *Ratios) SetBFRHorsExploitation(v float32)`

SetBFRHorsExploitation sets BFRHorsExploitation field to given value.

### HasBFRHorsExploitation

`func (o *Ratios) HasBFRHorsExploitation() bool`

HasBFRHorsExploitation returns a boolean if a field has been set.

### GetBFRJoursCA

`func (o *Ratios) GetBFRJoursCA() float32`

GetBFRJoursCA returns the BFRJoursCA field if non-nil, zero value otherwise.

### GetBFRJoursCAOk

`func (o *Ratios) GetBFRJoursCAOk() (*float32, bool)`

GetBFRJoursCAOk returns a tuple with the BFRJoursCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFRJoursCA

`func (o *Ratios) SetBFRJoursCA(v float32)`

SetBFRJoursCA sets BFRJoursCA field to given value.

### HasBFRJoursCA

`func (o *Ratios) HasBFRJoursCA() bool`

HasBFRJoursCA returns a boolean if a field has been set.

### GetBFRExploitationJoursCA

`func (o *Ratios) GetBFRExploitationJoursCA() float32`

GetBFRExploitationJoursCA returns the BFRExploitationJoursCA field if non-nil, zero value otherwise.

### GetBFRExploitationJoursCAOk

`func (o *Ratios) GetBFRExploitationJoursCAOk() (*float32, bool)`

GetBFRExploitationJoursCAOk returns a tuple with the BFRExploitationJoursCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFRExploitationJoursCA

`func (o *Ratios) SetBFRExploitationJoursCA(v float32)`

SetBFRExploitationJoursCA sets BFRExploitationJoursCA field to given value.

### HasBFRExploitationJoursCA

`func (o *Ratios) HasBFRExploitationJoursCA() bool`

HasBFRExploitationJoursCA returns a boolean if a field has been set.

### GetBFRHorsExploitationJoursCA

`func (o *Ratios) GetBFRHorsExploitationJoursCA() float32`

GetBFRHorsExploitationJoursCA returns the BFRHorsExploitationJoursCA field if non-nil, zero value otherwise.

### GetBFRHorsExploitationJoursCAOk

`func (o *Ratios) GetBFRHorsExploitationJoursCAOk() (*float32, bool)`

GetBFRHorsExploitationJoursCAOk returns a tuple with the BFRHorsExploitationJoursCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBFRHorsExploitationJoursCA

`func (o *Ratios) SetBFRHorsExploitationJoursCA(v float32)`

SetBFRHorsExploitationJoursCA sets BFRHorsExploitationJoursCA field to given value.

### HasBFRHorsExploitationJoursCA

`func (o *Ratios) HasBFRHorsExploitationJoursCA() bool`

HasBFRHorsExploitationJoursCA returns a boolean if a field has been set.

### GetDelaiPaiementClientsJours

`func (o *Ratios) GetDelaiPaiementClientsJours() float32`

GetDelaiPaiementClientsJours returns the DelaiPaiementClientsJours field if non-nil, zero value otherwise.

### GetDelaiPaiementClientsJoursOk

`func (o *Ratios) GetDelaiPaiementClientsJoursOk() (*float32, bool)`

GetDelaiPaiementClientsJoursOk returns a tuple with the DelaiPaiementClientsJours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaiPaiementClientsJours

`func (o *Ratios) SetDelaiPaiementClientsJours(v float32)`

SetDelaiPaiementClientsJours sets DelaiPaiementClientsJours field to given value.

### HasDelaiPaiementClientsJours

`func (o *Ratios) HasDelaiPaiementClientsJours() bool`

HasDelaiPaiementClientsJours returns a boolean if a field has been set.

### GetDelaiPaiementFournisseursJours

`func (o *Ratios) GetDelaiPaiementFournisseursJours() float32`

GetDelaiPaiementFournisseursJours returns the DelaiPaiementFournisseursJours field if non-nil, zero value otherwise.

### GetDelaiPaiementFournisseursJoursOk

`func (o *Ratios) GetDelaiPaiementFournisseursJoursOk() (*float32, bool)`

GetDelaiPaiementFournisseursJoursOk returns a tuple with the DelaiPaiementFournisseursJours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaiPaiementFournisseursJours

`func (o *Ratios) SetDelaiPaiementFournisseursJours(v float32)`

SetDelaiPaiementFournisseursJours sets DelaiPaiementFournisseursJours field to given value.

### HasDelaiPaiementFournisseursJours

`func (o *Ratios) HasDelaiPaiementFournisseursJours() bool`

HasDelaiPaiementFournisseursJours returns a boolean if a field has been set.

### GetRatioStockCAJours

`func (o *Ratios) GetRatioStockCAJours() float32`

GetRatioStockCAJours returns the RatioStockCAJours field if non-nil, zero value otherwise.

### GetRatioStockCAJoursOk

`func (o *Ratios) GetRatioStockCAJoursOk() (*float32, bool)`

GetRatioStockCAJoursOk returns a tuple with the RatioStockCAJours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatioStockCAJours

`func (o *Ratios) SetRatioStockCAJours(v float32)`

SetRatioStockCAJours sets RatioStockCAJours field to given value.

### HasRatioStockCAJours

`func (o *Ratios) HasRatioStockCAJours() bool`

HasRatioStockCAJours returns a boolean if a field has been set.

### GetCapaciteAutofinancement

`func (o *Ratios) GetCapaciteAutofinancement() float32`

GetCapaciteAutofinancement returns the CapaciteAutofinancement field if non-nil, zero value otherwise.

### GetCapaciteAutofinancementOk

`func (o *Ratios) GetCapaciteAutofinancementOk() (*float32, bool)`

GetCapaciteAutofinancementOk returns a tuple with the CapaciteAutofinancement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapaciteAutofinancement

`func (o *Ratios) SetCapaciteAutofinancement(v float32)`

SetCapaciteAutofinancement sets CapaciteAutofinancement field to given value.

### HasCapaciteAutofinancement

`func (o *Ratios) HasCapaciteAutofinancement() bool`

HasCapaciteAutofinancement returns a boolean if a field has been set.

### GetCapaciteAutofinancementCA

`func (o *Ratios) GetCapaciteAutofinancementCA() float32`

GetCapaciteAutofinancementCA returns the CapaciteAutofinancementCA field if non-nil, zero value otherwise.

### GetCapaciteAutofinancementCAOk

`func (o *Ratios) GetCapaciteAutofinancementCAOk() (*float32, bool)`

GetCapaciteAutofinancementCAOk returns a tuple with the CapaciteAutofinancementCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapaciteAutofinancementCA

`func (o *Ratios) SetCapaciteAutofinancementCA(v float32)`

SetCapaciteAutofinancementCA sets CapaciteAutofinancementCA field to given value.

### HasCapaciteAutofinancementCA

`func (o *Ratios) HasCapaciteAutofinancementCA() bool`

HasCapaciteAutofinancementCA returns a boolean if a field has been set.

### GetFondsRoulementNetGlobal

`func (o *Ratios) GetFondsRoulementNetGlobal() float32`

GetFondsRoulementNetGlobal returns the FondsRoulementNetGlobal field if non-nil, zero value otherwise.

### GetFondsRoulementNetGlobalOk

`func (o *Ratios) GetFondsRoulementNetGlobalOk() (*float32, bool)`

GetFondsRoulementNetGlobalOk returns a tuple with the FondsRoulementNetGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFondsRoulementNetGlobal

`func (o *Ratios) SetFondsRoulementNetGlobal(v float32)`

SetFondsRoulementNetGlobal sets FondsRoulementNetGlobal field to given value.

### HasFondsRoulementNetGlobal

`func (o *Ratios) HasFondsRoulementNetGlobal() bool`

HasFondsRoulementNetGlobal returns a boolean if a field has been set.

### GetCouvertureBFR

`func (o *Ratios) GetCouvertureBFR() float32`

GetCouvertureBFR returns the CouvertureBFR field if non-nil, zero value otherwise.

### GetCouvertureBFROk

`func (o *Ratios) GetCouvertureBFROk() (*float32, bool)`

GetCouvertureBFROk returns a tuple with the CouvertureBFR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouvertureBFR

`func (o *Ratios) SetCouvertureBFR(v float32)`

SetCouvertureBFR sets CouvertureBFR field to given value.

### HasCouvertureBFR

`func (o *Ratios) HasCouvertureBFR() bool`

HasCouvertureBFR returns a boolean if a field has been set.

### GetTresorerie

`func (o *Ratios) GetTresorerie() float32`

GetTresorerie returns the Tresorerie field if non-nil, zero value otherwise.

### GetTresorerieOk

`func (o *Ratios) GetTresorerieOk() (*float32, bool)`

GetTresorerieOk returns a tuple with the Tresorerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresorerie

`func (o *Ratios) SetTresorerie(v float32)`

SetTresorerie sets Tresorerie field to given value.

### HasTresorerie

`func (o *Ratios) HasTresorerie() bool`

HasTresorerie returns a boolean if a field has been set.

### GetDettesFinancieres

`func (o *Ratios) GetDettesFinancieres() float32`

GetDettesFinancieres returns the DettesFinancieres field if non-nil, zero value otherwise.

### GetDettesFinancieresOk

`func (o *Ratios) GetDettesFinancieresOk() (*float32, bool)`

GetDettesFinancieresOk returns a tuple with the DettesFinancieres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDettesFinancieres

`func (o *Ratios) SetDettesFinancieres(v float32)`

SetDettesFinancieres sets DettesFinancieres field to given value.

### HasDettesFinancieres

`func (o *Ratios) HasDettesFinancieres() bool`

HasDettesFinancieres returns a boolean if a field has been set.

### GetCapaciteRemboursement

`func (o *Ratios) GetCapaciteRemboursement() float32`

GetCapaciteRemboursement returns the CapaciteRemboursement field if non-nil, zero value otherwise.

### GetCapaciteRemboursementOk

`func (o *Ratios) GetCapaciteRemboursementOk() (*float32, bool)`

GetCapaciteRemboursementOk returns a tuple with the CapaciteRemboursement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapaciteRemboursement

`func (o *Ratios) SetCapaciteRemboursement(v float32)`

SetCapaciteRemboursement sets CapaciteRemboursement field to given value.

### HasCapaciteRemboursement

`func (o *Ratios) HasCapaciteRemboursement() bool`

HasCapaciteRemboursement returns a boolean if a field has been set.

### GetRatioEndettement

`func (o *Ratios) GetRatioEndettement() float32`

GetRatioEndettement returns the RatioEndettement field if non-nil, zero value otherwise.

### GetRatioEndettementOk

`func (o *Ratios) GetRatioEndettementOk() (*float32, bool)`

GetRatioEndettementOk returns a tuple with the RatioEndettement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatioEndettement

`func (o *Ratios) SetRatioEndettement(v float32)`

SetRatioEndettement sets RatioEndettement field to given value.

### HasRatioEndettement

`func (o *Ratios) HasRatioEndettement() bool`

HasRatioEndettement returns a boolean if a field has been set.

### GetAutonomieFinanciere

`func (o *Ratios) GetAutonomieFinanciere() float32`

GetAutonomieFinanciere returns the AutonomieFinanciere field if non-nil, zero value otherwise.

### GetAutonomieFinanciereOk

`func (o *Ratios) GetAutonomieFinanciereOk() (*float32, bool)`

GetAutonomieFinanciereOk returns a tuple with the AutonomieFinanciere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomieFinanciere

`func (o *Ratios) SetAutonomieFinanciere(v float32)`

SetAutonomieFinanciere sets AutonomieFinanciere field to given value.

### HasAutonomieFinanciere

`func (o *Ratios) HasAutonomieFinanciere() bool`

HasAutonomieFinanciere returns a boolean if a field has been set.

### GetTauxLevier

`func (o *Ratios) GetTauxLevier() float32`

GetTauxLevier returns the TauxLevier field if non-nil, zero value otherwise.

### GetTauxLevierOk

`func (o *Ratios) GetTauxLevierOk() (*float32, bool)`

GetTauxLevierOk returns a tuple with the TauxLevier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTauxLevier

`func (o *Ratios) SetTauxLevier(v float32)`

SetTauxLevier sets TauxLevier field to given value.

### HasTauxLevier

`func (o *Ratios) HasTauxLevier() bool`

HasTauxLevier returns a boolean if a field has been set.

### GetEtatDettes1AnAuPlus

`func (o *Ratios) GetEtatDettes1AnAuPlus() float32`

GetEtatDettes1AnAuPlus returns the EtatDettes1AnAuPlus field if non-nil, zero value otherwise.

### GetEtatDettes1AnAuPlusOk

`func (o *Ratios) GetEtatDettes1AnAuPlusOk() (*float32, bool)`

GetEtatDettes1AnAuPlusOk returns a tuple with the EtatDettes1AnAuPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtatDettes1AnAuPlus

`func (o *Ratios) SetEtatDettes1AnAuPlus(v float32)`

SetEtatDettes1AnAuPlus sets EtatDettes1AnAuPlus field to given value.

### HasEtatDettes1AnAuPlus

`func (o *Ratios) HasEtatDettes1AnAuPlus() bool`

HasEtatDettes1AnAuPlus returns a boolean if a field has been set.

### GetLiquiditeGenerale

`func (o *Ratios) GetLiquiditeGenerale() float32`

GetLiquiditeGenerale returns the LiquiditeGenerale field if non-nil, zero value otherwise.

### GetLiquiditeGeneraleOk

`func (o *Ratios) GetLiquiditeGeneraleOk() (*float32, bool)`

GetLiquiditeGeneraleOk returns a tuple with the LiquiditeGenerale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquiditeGenerale

`func (o *Ratios) SetLiquiditeGenerale(v float32)`

SetLiquiditeGenerale sets LiquiditeGenerale field to given value.

### HasLiquiditeGenerale

`func (o *Ratios) HasLiquiditeGenerale() bool`

HasLiquiditeGenerale returns a boolean if a field has been set.

### GetCouvertureDettes

`func (o *Ratios) GetCouvertureDettes() float32`

GetCouvertureDettes returns the CouvertureDettes field if non-nil, zero value otherwise.

### GetCouvertureDettesOk

`func (o *Ratios) GetCouvertureDettesOk() (*float32, bool)`

GetCouvertureDettesOk returns a tuple with the CouvertureDettes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouvertureDettes

`func (o *Ratios) SetCouvertureDettes(v float32)`

SetCouvertureDettes sets CouvertureDettes field to given value.

### HasCouvertureDettes

`func (o *Ratios) HasCouvertureDettes() bool`

HasCouvertureDettes returns a boolean if a field has been set.

### GetMargeNette

`func (o *Ratios) GetMargeNette() float32`

GetMargeNette returns the MargeNette field if non-nil, zero value otherwise.

### GetMargeNetteOk

`func (o *Ratios) GetMargeNetteOk() (*float32, bool)`

GetMargeNetteOk returns a tuple with the MargeNette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargeNette

`func (o *Ratios) SetMargeNette(v float32)`

SetMargeNette sets MargeNette field to given value.

### HasMargeNette

`func (o *Ratios) HasMargeNette() bool`

HasMargeNette returns a boolean if a field has been set.

### GetRentabiliteFondsPropres

`func (o *Ratios) GetRentabiliteFondsPropres() float32`

GetRentabiliteFondsPropres returns the RentabiliteFondsPropres field if non-nil, zero value otherwise.

### GetRentabiliteFondsPropresOk

`func (o *Ratios) GetRentabiliteFondsPropresOk() (*float32, bool)`

GetRentabiliteFondsPropresOk returns a tuple with the RentabiliteFondsPropres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentabiliteFondsPropres

`func (o *Ratios) SetRentabiliteFondsPropres(v float32)`

SetRentabiliteFondsPropres sets RentabiliteFondsPropres field to given value.

### HasRentabiliteFondsPropres

`func (o *Ratios) HasRentabiliteFondsPropres() bool`

HasRentabiliteFondsPropres returns a boolean if a field has been set.

### GetRentabiliteEconomique

`func (o *Ratios) GetRentabiliteEconomique() float32`

GetRentabiliteEconomique returns the RentabiliteEconomique field if non-nil, zero value otherwise.

### GetRentabiliteEconomiqueOk

`func (o *Ratios) GetRentabiliteEconomiqueOk() (*float32, bool)`

GetRentabiliteEconomiqueOk returns a tuple with the RentabiliteEconomique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentabiliteEconomique

`func (o *Ratios) SetRentabiliteEconomique(v float32)`

SetRentabiliteEconomique sets RentabiliteEconomique field to given value.

### HasRentabiliteEconomique

`func (o *Ratios) HasRentabiliteEconomique() bool`

HasRentabiliteEconomique returns a boolean if a field has been set.

### GetValeurAjoutee

`func (o *Ratios) GetValeurAjoutee() float32`

GetValeurAjoutee returns the ValeurAjoutee field if non-nil, zero value otherwise.

### GetValeurAjouteeOk

`func (o *Ratios) GetValeurAjouteeOk() (*float32, bool)`

GetValeurAjouteeOk returns a tuple with the ValeurAjoutee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurAjoutee

`func (o *Ratios) SetValeurAjoutee(v float32)`

SetValeurAjoutee sets ValeurAjoutee field to given value.

### HasValeurAjoutee

`func (o *Ratios) HasValeurAjoutee() bool`

HasValeurAjoutee returns a boolean if a field has been set.

### GetValeurAjouteeCA

`func (o *Ratios) GetValeurAjouteeCA() float32`

GetValeurAjouteeCA returns the ValeurAjouteeCA field if non-nil, zero value otherwise.

### GetValeurAjouteeCAOk

`func (o *Ratios) GetValeurAjouteeCAOk() (*float32, bool)`

GetValeurAjouteeCAOk returns a tuple with the ValeurAjouteeCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValeurAjouteeCA

`func (o *Ratios) SetValeurAjouteeCA(v float32)`

SetValeurAjouteeCA sets ValeurAjouteeCA field to given value.

### HasValeurAjouteeCA

`func (o *Ratios) HasValeurAjouteeCA() bool`

HasValeurAjouteeCA returns a boolean if a field has been set.

### GetSalairesChargesSociales

`func (o *Ratios) GetSalairesChargesSociales() float32`

GetSalairesChargesSociales returns the SalairesChargesSociales field if non-nil, zero value otherwise.

### GetSalairesChargesSocialesOk

`func (o *Ratios) GetSalairesChargesSocialesOk() (*float32, bool)`

GetSalairesChargesSocialesOk returns a tuple with the SalairesChargesSociales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalairesChargesSociales

`func (o *Ratios) SetSalairesChargesSociales(v float32)`

SetSalairesChargesSociales sets SalairesChargesSociales field to given value.

### HasSalairesChargesSociales

`func (o *Ratios) HasSalairesChargesSociales() bool`

HasSalairesChargesSociales returns a boolean if a field has been set.

### GetSalairesCA

`func (o *Ratios) GetSalairesCA() float32`

GetSalairesCA returns the SalairesCA field if non-nil, zero value otherwise.

### GetSalairesCAOk

`func (o *Ratios) GetSalairesCAOk() (*float32, bool)`

GetSalairesCAOk returns a tuple with the SalairesCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalairesCA

`func (o *Ratios) SetSalairesCA(v float32)`

SetSalairesCA sets SalairesCA field to given value.

### HasSalairesCA

`func (o *Ratios) HasSalairesCA() bool`

HasSalairesCA returns a boolean if a field has been set.

### GetImpotsTaxes

`func (o *Ratios) GetImpotsTaxes() float32`

GetImpotsTaxes returns the ImpotsTaxes field if non-nil, zero value otherwise.

### GetImpotsTaxesOk

`func (o *Ratios) GetImpotsTaxesOk() (*float32, bool)`

GetImpotsTaxesOk returns a tuple with the ImpotsTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpotsTaxes

`func (o *Ratios) SetImpotsTaxes(v float32)`

SetImpotsTaxes sets ImpotsTaxes field to given value.

### HasImpotsTaxes

`func (o *Ratios) HasImpotsTaxes() bool`

HasImpotsTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


