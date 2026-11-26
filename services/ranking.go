package services

import (
	"sort"

	"brawlhalla-tierlist/models"
)

// Pondération des trois critères du classement.
const (
	WinrateWeight  = 0.50
	PlayrateWeight = 0.30
	WeaponWeight   = 0.20
)

// weaponEntry associe une arme à ses dégâts et à sa légende.
type weaponEntry struct {
	LegendIndex int
	Slot        int
	Damage      float64
}

// CalculateRanking calcule le score global de chaque légende.
func CalculateRanking(legends []models.LegendWithStats) []models.LegendRanking {

	rankings := make([]models.LegendRanking, len(legends))

	var winrates []float64
	var playrates []float64

	// Récupération des valeurs de winrate et de pickrate.
	for _, legend := range legends {
		winrates = append(winrates, legend.Stats.Winrate)
		playrates = append(playrates, legend.Stats.Playrate)
	}

	// Préparation des statistiques des armes.
	weaponEntries := make(map[string][]weaponEntry)
	weaponDamages := make(map[string][]float64)

	for i, legend := range legends {

		// Première arme.
		if legend.Weapon1 != "" {
			entry := weaponEntry{
				LegendIndex: i,
				Slot:        1,
				Damage:      legend.Stats.DamageWeaponOne,
			}

			weaponEntries[legend.Weapon1] = append(
				weaponEntries[legend.Weapon1],
				entry,
			)

			weaponDamages[legend.Weapon1] = append(
				weaponDamages[legend.Weapon1],
				entry.Damage,
			)
		}

		// Deuxième arme.
		if legend.Weapon2 != "" {
			entry := weaponEntry{
				LegendIndex: i,
				Slot:        2,
				Damage:      legend.Stats.DamageWeaponTwo,
			}

			weaponEntries[legend.Weapon2] = append(
				weaponEntries[legend.Weapon2],
				entry,
			)

			weaponDamages[legend.Weapon2] = append(
				weaponDamages[legend.Weapon2],
				entry.Damage,
			)
		}
	}

	// Calcul des scores individuels.
	for i, legend := range legends {

		ranking := models.LegendRanking{
			LegendWithStats: legend,
		}

		// Score de winrate et de pickrate.
		ranking.WinrateScore = percentileScore(
			legend.Stats.Winrate,
			winrates,
		)

		ranking.PlayrateScore = percentileScore(
			legend.Stats.Playrate,
			playrates,
		)

		// Calcul des scores de chaque arme.
		for _, entry := range weaponEntries[legend.Weapon1] {
			if entry.LegendIndex == i && entry.Slot == 1 {
				ranking.WeaponOneScore = percentileScore(
					entry.Damage,
					weaponDamages[legend.Weapon1],
				)
				break
			}
		}

		for _, entry := range weaponEntries[legend.Weapon2] {
			if entry.LegendIndex == i && entry.Slot == 2 {
				ranking.WeaponTwoScore = percentileScore(
					entry.Damage,
					weaponDamages[legend.Weapon2],
				)
				break
			}
		}

		// Moyenne des performances des deux armes.
		ranking.WeaponPerformanceScore = (ranking.WeaponOneScore + ranking.WeaponTwoScore) / 2

		// Calcul du score global sur 100.
		ranking.TotalScore =
			ranking.WinrateScore*WinrateWeight +
				ranking.PlayrateScore*PlayrateWeight +
				ranking.WeaponPerformanceScore*WeaponWeight

		rankings[i] = ranking
	}

	// Tri des légendes du meilleur score au moins bon.
	sort.SliceStable(rankings, func(i, j int) bool {
		return rankings[i].TotalScore > rankings[j].TotalScore
	})

	// Attribution des rangs, avec égalités possibles.
	for i := range rankings {
		rankings[i].Rank = i + 1

		if i > 0 && rankings[i].TotalScore == rankings[i-1].TotalScore {
			rankings[i].Rank = rankings[i-1].Rank
		}
	}

	return rankings
}

// percentileScore transforme une valeur en score comparatif sur 100.
// Une valeur supérieure obtient un meilleur score.
func percentileScore(value float64, values []float64) float64 {

	if len(values) <= 1 {
		return 50
	}

	var lower int
	var equal int

	for _, other := range values {
		if other < value {
			lower++
		} else if other == value {
			equal++
		}
	}

	// Le rang moyen permet de gérer les valeurs identiques.
	position := float64(lower) + float64(equal-1)/2

	return position / float64(len(values)-1) * 100
}
