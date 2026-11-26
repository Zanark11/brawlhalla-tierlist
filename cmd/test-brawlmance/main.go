package main

import (
	"fmt"
	"log"

	"brawlhalla-tierlist/api"
)

func main() {
	// Récupération des statistiques Brawlmance.
	legends, err := api.GetLegendStats()
	if err != nil {
		log.Fatal(err)
	}

	// Affichage des résultats.
	fmt.Println("======================================")
	fmt.Println("      TEST API BRAWLMANCE")
	fmt.Println("======================================")
	fmt.Println("Patch :", api.SelectedPatch)
	fmt.Println("Rang :", api.SelectedTier)
	fmt.Println("Nombre de légendes :", len(legends))
	fmt.Println("--------------------------------------")

	for _, legend := range legends {
		fmt.Printf(
			"%-12s | Parties : %-8d | Winrate : %.2f%% | Pickrate : %.2f%%\n",
			legend.Name,
			legend.Stats.Games,
			legend.Stats.Winrate,
			legend.Stats.Playrate,
		)
	}
}
