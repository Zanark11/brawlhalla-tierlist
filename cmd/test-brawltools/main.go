package main

import (
	"fmt"
	"log"

	"brawlhalla-tierlist/api"
)

func main() {
	// Récupère les tournois européens officiels de 2026.
	tournaments, err := api.GetEUTournaments(2026)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nombre de tournois :", len(tournaments))

	// Affiche les informations de chaque tournoi.
	for _, tournament := range tournaments {
		fmt.Println("-----------------------------")
		fmt.Println("Nom :", tournament.TournamentName)
		fmt.Println("Année :", tournament.Year)
		fmt.Println("ID :", tournament.ID)
	}
}
