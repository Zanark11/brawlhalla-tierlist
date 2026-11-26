package models

// LegendRanking représente le résultat du calcul du classement.
type LegendRanking struct {
	LegendWithStats

	// Scores normalisés sur 100 pour chaque critère.
	WinrateScore  float64
	PlayrateScore float64

	// Scores de performance des deux armes.
	WeaponOneScore float64
	WeaponTwoScore float64

	// Moyenne des performances des deux armes.
	WeaponPerformanceScore float64

	// Score final pondéré sur 100.
	TotalScore float64

	// Position de la légende dans le classement.
	Rank int
}
