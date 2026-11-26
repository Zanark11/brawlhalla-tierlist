
package main

import (
	"fmt"
	"log"

	"brawlhalla-tierlist/api"
)

func main() {
	// Récupération des légendes officielles.
	officialLegends, err := api.GetLegends()
	if err != nil {
		log.Fatal(err)
	}

	// Récupération des statistiques Brawlmance.
	stats, err := api.GetLegendStats()
	if err != nil {
		log.Fatal(err)
	}

	// Création d'une liste des identifiants officiels.
	officialIDs := make(map[int]string)

	for _, legend := range officialLegends {
		officialIDs[legend.ID] = legend.Name
	}

	// Vérification des correspondances.
	matches := 0
	missing := 0

	fmt.Println("===== VÉRIFICATION DES IDENTIFIANTS =====")

	for _, legend := range stats {
		name, exists := officialIDs[legend.ID]

		if exists {
			matches++
			fmt.Printf("OK  | ID : %3d | %s\n", legend.ID, name)
		} else {
			missing++
			fmt.Printf("ABSENT | ID : %3d | %s\n", legend.ID, legend.Name)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Println("Statistiques Brawlmance :", len(stats))
	fmt.Println("Correspondances :", matches)
	fmt.Println("Sans correspondance :", missing)
}
