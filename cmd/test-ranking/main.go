package main

import (
	"fmt"
	"log"

	"brawlhalla-tierlist/api"
	"brawlhalla-tierlist/services"
)

func main() {
	// Récupération des statistiques Brawlmance.
	legends, err := api.GetLegendStats()
	if err != nil {
		log.Fatal(err)
	}

	// Calcul du classement.
	rankings := services.CalculateRanking(legends)

	// Affichage des 10 meilleures légendes.
	fmt.Println("==============================================")
	fmt.Println("          TOP 10 - TIERLIST BRAWLHALLA")
	fmt.Println("==============================================")

	for i, ranking := range rankings {
		if i >= 10 {
			break
		}

		fmt.Printf(
			"%2d. %-15s | Score : %6.2f | Winrate : %5.2f | Pickrate : %5.2f | Armes : %5.2f\n",
			ranking.Rank,
			ranking.Name,
			ranking.TotalScore,
			ranking.WinrateScore,
			ranking.PlayrateScore,
			ranking.WeaponPerformanceScore,
		)
	}
}
