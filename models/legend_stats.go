package models

// LegendStats représente les statistiques d'une légende fournies par Brawlmance.
type LegendStats struct {
	Games              int     `json:"games"`
	Playrate           float64 `json:"playrate"`
	Winrate            float64 `json:"winrate"`
	DamageDealt        float64 `json:"damagedealt"`
	DamageWeaponOne    float64 `json:"damagedealt_weaponone"`
	DamageWeaponTwo    float64 `json:"damagedealt_weapontwo"`
	DamageUnarmed      float64 `json:"damagedealt_unarmed"`
	DamageGadgets      float64 `json:"damagedealt_gadgets"`
	DamageTaken        float64 `json:"damagetaken"`
	MatchTime          float64 `json:"matchtime"`
	MatchTimeWeaponOne float64 `json:"matchtime_weaponone"`
	MatchTimeWeaponTwo float64 `json:"matchtime_weapontwo"`
}

// LegendStatsResponse représente la réponse complète de l'API Brawlmance.
type LegendStatsResponse struct {
	Legends []LegendWithStats `json:"legends"`
}

// LegendWithStats associe une légende à ses statistiques.
type LegendWithStats struct {
	Legend
	Stats LegendStats `json:"stats"`
}
